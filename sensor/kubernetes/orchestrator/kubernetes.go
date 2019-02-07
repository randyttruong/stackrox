package orchestrator

import (
	"errors"
	"fmt"
	"time"

	"github.com/stackrox/rox/pkg/benchmarks"
	"github.com/stackrox/rox/pkg/env"
	"github.com/stackrox/rox/pkg/errorhelpers"
	pkgKubernetes "github.com/stackrox/rox/pkg/kubernetes"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/orchestrators"
	v1beta12 "k8s.io/api/extensions/v1beta1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	"k8s.io/client-go/rest"
)

const (
	ownershipLabel = `owner.stackrox.io/sensor`

	namespace = "stackrox"
)

var (
	log = logging.LoggerForModule()
)

type kubernetesOrchestrator struct {
	client    *kubernetes.Clientset
	converter converter
	namespace string

	sensorInstanceID string
}

// MustCreate returns a new Kubernetes orchestrator client, or panics.
func MustCreate(sensorInstanceID string) orchestrators.Orchestrator {
	o, err := New(sensorInstanceID)
	if err != nil {
		panic(err)
	}
	return o
}

// New returns a new kubernetes orchestrator client.
func New(sensorInstanceID string) (orchestrators.Orchestrator, error) {
	c, err := setupClient()
	if err != nil {
		log.Errorf("unable to create kubernetes client: %s", err)
		return nil, err
	}
	return &kubernetesOrchestrator{
		client:           c,
		converter:        newConverter(),
		namespace:        namespace,
		sensorInstanceID: sensorInstanceID,
	}, nil
}

func setupClient() (client *kubernetes.Clientset, err error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return
	}

	return kubernetes.NewForConfig(config)
}

func (k *kubernetesOrchestrator) launch(setInterface v1beta1.DaemonSetInterface, ds *v1beta12.DaemonSet) (string, error) {
	for i := 0; i < 3; i++ {
		actual, err := k.client.ExtensionsV1beta1().DaemonSets(k.namespace).Create(ds)
		if err != nil {
			if statusErr, ok := err.(*k8sErrors.StatusError); ok && statusErr.Status().Reason == metav1.StatusReasonAlreadyExists {
				time.Sleep(10 * time.Second)
				continue
			}
			return "", err
		}
		return actual.Name, nil
	}
	return "", errors.New("unable to launch daemonset")
}

func (k *kubernetesOrchestrator) patchLabels(labels *map[string]string) {
	if *labels == nil {
		*labels = make(map[string]string)
	}
	(*labels)[ownershipLabel] = k.sensorInstanceID
}

func (k *kubernetesOrchestrator) Launch(service orchestrators.SystemService) (string, error) {
	if service.Global {
		ds := k.converter.asDaemonSet(k.newServiceWrap(service))
		k.patchLabels(&ds.Labels)
		launchedName, err := k.launch(k.client.ExtensionsV1beta1().DaemonSets(k.namespace), ds)
		if err != nil {
			log.Errorf("unable to create daemonset %s: %s", service.Name, err)
			return "", err
		}
		return launchedName, nil
	}

	deploy := k.converter.asDeployment(k.newServiceWrap(service))
	k.patchLabels(&deploy.Labels)
	actual, err := k.client.ExtensionsV1beta1().Deployments(k.namespace).Create(deploy)
	if err != nil {
		log.Errorf("unable to create deployment %s: %s", service.Name, err)
		return "", err
	}

	return actual.Name, nil
}

func (k *kubernetesOrchestrator) LaunchBenchmark(service orchestrators.SystemService) (string, error) {
	service.Command = []string{benchmarks.BenchmarkCommand}
	service.Envs = append(service.Envs, env.Combine(env.BenchmarkCompletion.EnvVar(), "true"))
	service.Mounts = benchmarks.BenchmarkMounts
	service.HostPID = true
	return k.Launch(service)
}

func (k *kubernetesOrchestrator) newServiceWrap(service orchestrators.SystemService) *serviceWrap {
	return &serviceWrap{
		SystemService: service,
		namespace:     k.namespace,
	}
}

func (k *kubernetesOrchestrator) Kill(name string) error {
	if ds, err := k.client.ExtensionsV1beta1().DaemonSets(k.namespace).Get(name, metav1.GetOptions{}); err == nil && ds != nil {
		if err := k.client.ExtensionsV1beta1().DaemonSets(k.namespace).Delete(name, pkgKubernetes.DeleteOption); err != nil {
			log.Errorf("unable to delete daemonset %s: %s", name, err)
			return err
		}
		return nil
	}

	if deploy, err := k.client.ExtensionsV1beta1().Deployments(k.namespace).Get(name, metav1.GetOptions{}); err == nil && deploy != nil {
		if err := k.client.ExtensionsV1beta1().Deployments(k.namespace).Delete(name, pkgKubernetes.DeleteOption); err != nil {
			log.Errorf("unable to delete deployment %s: %s", name, err)
			return err
		}
		return nil
	}

	err := fmt.Errorf("unable to delete service %s; service not found", name)
	log.Error(err)
	return err
}

// WaitForCompletion currently cannot be implemented in Kubernetes because DaemonSet Restart Policy must be always
func (k *kubernetesOrchestrator) WaitForCompletion(_ string, timeout time.Duration) error {
	time.Sleep(timeout)
	return nil
}

func (k *kubernetesOrchestrator) labelSelector(ownedByThisInstance bool) (labels.Selector, error) {
	hasLabelReq, err := labels.NewRequirement(ownershipLabel, selection.Exists, nil)
	if err != nil {
		return nil, err
	}
	var op selection.Operator
	if ownedByThisInstance {
		op = selection.Equals
	} else {
		op = selection.NotEquals
	}
	labelMatchesReq, err := labels.NewRequirement(ownershipLabel, op, []string{k.sensorInstanceID})
	if err != nil {
		return nil, err
	}
	return labels.NewSelector().Add(*hasLabelReq, *labelMatchesReq), nil
}

func (k *kubernetesOrchestrator) CleanUp(ownedByThisInstance bool) error {
	ls, err := k.labelSelector(ownedByThisInstance)
	if err != nil {
		return fmt.Errorf("creating label selector: %v", err)
	}
	listOpts := metav1.ListOptions{
		LabelSelector: ls.String(),
	}
	propagationPolicy := metav1.DeletePropagationBackground
	deleteOpts := &metav1.DeleteOptions{
		PropagationPolicy: &propagationPolicy,
	}

	var errList errorhelpers.ErrorList
	err = k.client.ExtensionsV1beta1().DaemonSets(k.namespace).DeleteCollection(deleteOpts, listOpts)
	if err != nil {
		errList.AddStringf("deleting daemonsets: %v", err)
	}
	err = k.client.ExtensionsV1beta1().Deployments(k.namespace).DeleteCollection(deleteOpts, listOpts)
	if err != nil {
		errList.AddStringf("deleting deployments: %v", err)
	}

	return errList.ToError()
}

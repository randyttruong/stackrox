package scan

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/roxctl/common"
	"github.com/stackrox/rox/roxctl/common/flags"
	"github.com/stackrox/rox/roxctl/common/util"
)

// Command checks the image against image build lifecycle policies
func Command() *cobra.Command {
	var (
		image          string
		force          bool
		includeSnoozed bool
	)
	c := &cobra.Command{
		Use: "scan",
		RunE: util.RunENoArgs(func(c *cobra.Command) error {
			if image == "" {
				return errors.New("--image must be set")
			}
			return scanImage(image, force, includeSnoozed, flags.Timeout(c))
		}),
	}

	c.Flags().StringVarP(&image, "image", "i", "", "image name and reference. (e.g. nginx:latest or nginx@sha256:...)")
	c.Flags().BoolVarP(&force, "force", "f", false, "the --force flag ignores Central's cache for the scan and forces a fresh re-pull from Scanner")
	c.Flags().BoolVarP(&includeSnoozed, "include-snoozed", "a", true, "the --include-snoozed flag returns both snoozed and unsnoozed CVEs if set to false")
	return c
}

func scanImage(image string, force, includeSnoozed bool, timeout time.Duration) error {
	// Create the connection to the central detection service.
	conn, err := common.GetGRPCConnection()
	if err != nil {
		return err
	}
	defer func() {
		_ = conn.Close()
	}()

	service := v1.NewImageServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	imageResult, err := service.ScanImage(ctx, &v1.ScanImageRequest{
		ImageName:      image,
		Force:          force,
		IncludeSnoozed: includeSnoozed,
	})
	if err != nil {
		return err
	}

	marshaler := &jsonpb.Marshaler{
		Indent: "  ",
	}
	result, err := marshaler.MarshalToString(imageResult)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

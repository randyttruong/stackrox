// Code generated by blevebindings generator. DO NOT EDIT.

package index

import (
	"bytes"
	"context"
	bleve "github.com/blevesearch/bleve"
	metrics "github.com/stackrox/rox/central/metrics"
	mappings "github.com/stackrox/rox/central/nodecomponentedge/mappings"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	batcher "github.com/stackrox/rox/pkg/batcher"
	ops "github.com/stackrox/rox/pkg/metrics"
	search "github.com/stackrox/rox/pkg/search"
	blevesearch "github.com/stackrox/rox/pkg/search/blevesearch"
	"time"
)

const batchSize = 5000

const resourceName = "NodeComponentEdge"

type indexerImpl struct {
	index bleve.Index
}

type nodeComponentEdgeWrapper struct {
	*storage.NodeComponentEdge `json:"node_component_edge"`
	Type                       string `json:"type"`
}

func (b *indexerImpl) AddNodeComponentEdge(nodecomponentedge *storage.NodeComponentEdge) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Add, "NodeComponentEdge")
	if err := b.index.Index(nodecomponentedge.GetId(), &nodeComponentEdgeWrapper{
		NodeComponentEdge: nodecomponentedge,
		Type:              v1.SearchCategory_NODE_COMPONENT_EDGE.String(),
	}); err != nil {
		return err
	}
	return nil
}

func (b *indexerImpl) AddNodeComponentEdges(nodecomponentedges []*storage.NodeComponentEdge) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.AddMany, "NodeComponentEdge")
	batchManager := batcher.New(len(nodecomponentedges), batchSize)
	for {
		start, end, ok := batchManager.Next()
		if !ok {
			break
		}
		if err := b.processBatch(nodecomponentedges[start:end]); err != nil {
			return err
		}
	}
	return nil
}

func (b *indexerImpl) processBatch(nodecomponentedges []*storage.NodeComponentEdge) error {
	batch := b.index.NewBatch()
	for _, nodecomponentedge := range nodecomponentedges {
		if err := batch.Index(nodecomponentedge.GetId(), &nodeComponentEdgeWrapper{
			NodeComponentEdge: nodecomponentedge,
			Type:              v1.SearchCategory_NODE_COMPONENT_EDGE.String(),
		}); err != nil {
			return err
		}
	}
	return b.index.Batch(batch)
}

func (b *indexerImpl) Count(ctx context.Context, q *v1.Query, opts ...blevesearch.SearchOption) (int, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Count, "NodeComponentEdge")
	return blevesearch.RunCountRequest(v1.SearchCategory_NODE_COMPONENT_EDGE, q, b.index, mappings.OptionsMap, opts...)
}

func (b *indexerImpl) DeleteNodeComponentEdge(id string) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Remove, "NodeComponentEdge")
	if err := b.index.Delete(id); err != nil {
		return err
	}
	return nil
}

func (b *indexerImpl) DeleteNodeComponentEdges(ids []string) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.RemoveMany, "NodeComponentEdge")
	batch := b.index.NewBatch()
	for _, id := range ids {
		batch.Delete(id)
	}
	if err := b.index.Batch(batch); err != nil {
		return err
	}
	return nil
}

func (b *indexerImpl) MarkInitialIndexingComplete() error {
	return b.index.SetInternal([]byte(resourceName), []byte("old"))
}

func (b *indexerImpl) NeedsInitialIndexing() (bool, error) {
	data, err := b.index.GetInternal([]byte(resourceName))
	if err != nil {
		return false, err
	}
	return !bytes.Equal([]byte("old"), data), nil
}

func (b *indexerImpl) Search(ctx context.Context, q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Search, "NodeComponentEdge")
	return blevesearch.RunSearchRequest(v1.SearchCategory_NODE_COMPONENT_EDGE, q, b.index, mappings.OptionsMap, opts...)
}

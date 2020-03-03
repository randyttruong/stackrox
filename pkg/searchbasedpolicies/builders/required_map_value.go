package builders

import (
	"fmt"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/searchbasedpolicies"
)

// RequiredMapValueQueryBuilder builds queries to check for the (absence of) a required map value.
type RequiredMapValueQueryBuilder struct {
	GetKeyValuePolicy func(*storage.PolicyFields) *storage.KeyValuePolicy
	FieldName         string
	FieldLabel        search.FieldLabel
}

// Query implements the PolicyQueryBuilder interface.
func (r RequiredMapValueQueryBuilder) Query(fields *storage.PolicyFields, optionsMap map[search.FieldLabel]*v1.SearchField) (*v1.Query, searchbasedpolicies.ViolationPrinter, error) {
	keyValuePolicy := r.GetKeyValuePolicy(fields)
	return mapKeyValueQuery(optionsMap, keyValuePolicy, r.FieldLabel, r.FieldName, r.Name(), search.ExactMatchString)
}

// Name implements the PolicyQueryBuilder interface.
func (r RequiredMapValueQueryBuilder) Name() string {
	return fmt.Sprintf("query builder for required %s", r.FieldName)
}

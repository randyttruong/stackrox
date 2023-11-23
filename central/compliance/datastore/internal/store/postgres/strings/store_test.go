// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/suite"
)

type ComplianceStringsStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestComplianceStringsStore(t *testing.T) {
	suite.Run(t, new(ComplianceStringsStoreSuite))
}

func (s *ComplianceStringsStoreSuite) SetupSuite() {

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *ComplianceStringsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE compliance_strings CASCADE")
	s.T().Log("compliance_strings", tag)
	s.store = New(s.testDB.DB)
	s.NoError(err)
}

func (s *ComplianceStringsStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *ComplianceStringsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	complianceStrings := &storage.ComplianceStrings{}
	s.NoError(testutils.FullInit(complianceStrings, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundComplianceStrings, exists, err := store.Get(ctx, complianceStrings.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundComplianceStrings)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, complianceStrings))
	foundComplianceStrings, exists, err = store.Get(ctx, complianceStrings.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(complianceStrings, foundComplianceStrings)

	complianceStringsCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(1, complianceStringsCount)
	complianceStringsCount, err = store.Count(withNoAccessCtx)
	s.NoError(err)
	s.Zero(complianceStringsCount)

	complianceStringsExists, err := store.Exists(ctx, complianceStrings.GetId())
	s.NoError(err)
	s.True(complianceStringsExists)
	s.NoError(store.Upsert(ctx, complianceStrings))
	s.ErrorIs(store.Upsert(withNoAccessCtx, complianceStrings), sac.ErrResourceAccessDenied)

	foundComplianceStrings, exists, err = store.Get(ctx, complianceStrings.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(complianceStrings, foundComplianceStrings)

	s.NoError(store.Delete(ctx, complianceStrings.GetId()))
	foundComplianceStrings, exists, err = store.Get(ctx, complianceStrings.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundComplianceStrings)
	s.NoError(store.Delete(withNoAccessCtx, complianceStrings.GetId()))

	var complianceStringss []*storage.ComplianceStrings
	var complianceStringsIDs []string
	for i := 0; i < 200; i++ {
		complianceStrings := &storage.ComplianceStrings{}
		s.NoError(testutils.FullInit(complianceStrings, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		complianceStringss = append(complianceStringss, complianceStrings)
		complianceStringsIDs = append(complianceStringsIDs, complianceStrings.GetId())
	}

	s.NoError(store.UpsertMany(ctx, complianceStringss))

	complianceStringsCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(200, complianceStringsCount)

	s.NoError(store.DeleteMany(ctx, complianceStringsIDs))

	complianceStringsCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(0, complianceStringsCount)
}

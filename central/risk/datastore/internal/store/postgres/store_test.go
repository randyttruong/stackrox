// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	storage "github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/suite"
)

type RiskStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
}

func TestRiskStore(t *testing.T) {
	suite.Run(t, new(RiskStoreSuite))
}

func (s *RiskStoreSuite) SetupTest() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}
}

func (s *RiskStoreSuite) TearDownTest() {
	s.envIsolator.RestoreAll()
}

func (s *RiskStoreSuite) TestStore() {
	ctx := context.Background()

	source := pgtest.GetConnectionString(s.T())
	config, err := pgxpool.ParseConfig(source)
	s.Require().NoError(err)
	pool, err := pgxpool.ConnectConfig(ctx, config)
	s.NoError(err)
	defer pool.Close()

	Destroy(ctx, pool)
	store := New(ctx, pool)

	risk := &storage.Risk{}
	s.NoError(testutils.FullInit(risk, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundRisk, exists, err := store.Get(ctx, risk.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundRisk)

	s.NoError(store.Upsert(ctx, risk))
	foundRisk, exists, err = store.Get(ctx, risk.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(risk, foundRisk)

	riskCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(riskCount, 1)

	riskExists, err := store.Exists(ctx, risk.GetId())
	s.NoError(err)
	s.True(riskExists)
	s.NoError(store.Upsert(ctx, risk))

	foundRisk, exists, err = store.Get(ctx, risk.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(risk, foundRisk)

	s.NoError(store.Delete(ctx, risk.GetId()))
	foundRisk, exists, err = store.Get(ctx, risk.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundRisk)

	var risks []*storage.Risk
	for i := 0; i < 200; i++ {
		risk := &storage.Risk{}
		s.NoError(testutils.FullInit(risk, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		risks = append(risks, risk)
	}

	s.NoError(store.UpsertMany(ctx, risks))

	riskCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(riskCount, 200)
}

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

type SecretsStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
}

func TestSecretsStore(t *testing.T) {
	suite.Run(t, new(SecretsStoreSuite))
}

func (s *SecretsStoreSuite) SetupTest() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}
}

func (s *SecretsStoreSuite) TearDownTest() {
	s.envIsolator.RestoreAll()
}

func (s *SecretsStoreSuite) TestStore() {
	ctx := context.Background()

	source := pgtest.GetConnectionString(s.T())
	config, err := pgxpool.ParseConfig(source)
	s.Require().NoError(err)
	pool, err := pgxpool.ConnectConfig(ctx, config)
	s.NoError(err)
	defer pool.Close()

	Destroy(ctx, pool)
	store := New(ctx, pool)

	secret := &storage.Secret{}
	s.NoError(testutils.FullInit(secret, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundSecret, exists, err := store.Get(ctx, secret.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundSecret)

	s.NoError(store.Upsert(ctx, secret))
	foundSecret, exists, err = store.Get(ctx, secret.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(secret, foundSecret)

	secretCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(secretCount, 1)

	secretExists, err := store.Exists(ctx, secret.GetId())
	s.NoError(err)
	s.True(secretExists)
	s.NoError(store.Upsert(ctx, secret))

	foundSecret, exists, err = store.Get(ctx, secret.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(secret, foundSecret)

	s.NoError(store.Delete(ctx, secret.GetId()))
	foundSecret, exists, err = store.Get(ctx, secret.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundSecret)

	var secrets []*storage.Secret
	for i := 0; i < 200; i++ {
		secret := &storage.Secret{}
		s.NoError(testutils.FullInit(secret, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		secrets = append(secrets, secret)
	}

	s.NoError(store.UpsertMany(ctx, secrets))

	secretCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(secretCount, 200)
}

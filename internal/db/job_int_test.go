//go:build integration
// +build integration

package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJobDatabsse(t *testing.T) {
	t.Run("test create job", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		job, err := db.AddJob(context.Background(), "test-job", []byte("test-payload"))

		assert.NoError(t, err)

		assert.Greater(t, job, uint64(0), "job should be greater than 0")
	})
}

//go:build e2e
// +build e2e

package tests

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestAddJob(t *testing.T) {
	t.Run("addJob", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetBody(`{"name":"job1","payload":"echo hello"}`).
			Post("http://localhost:8080/api/v1/job")

		assert.NoError(t, err)
		assert.Equal(t, 201, resp.StatusCode())
	})
}

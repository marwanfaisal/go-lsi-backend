package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyncProductNoError(t *testing.T) {
	// Execute
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, "http://127.0.0.1:8000/v1/product/sync", nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Validate
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, response.StatusCode)
}

func TestSourceProductNoError(t *testing.T) {
	// Execute
	response, err := http.Get("http://127.0.0.1:8000/v1/product/source")

	// Validate
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, response.StatusCode)
}

func TestDestinationProductNoError(t *testing.T) {
	// Execute
	response, err := http.Get("http://127.0.0.1:8000/v1/product/destination")

	// Validate
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, response.StatusCode)
}

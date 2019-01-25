package giantbomb

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`
			{
				"version": "42"
			}
		`))
	}))
	defer server.Close()

	api := GBClient{server.Client(), server.URL, ""}
	response, err := api.Search("test", 10, 1, []string{ResourceTypeGame})

	assert.Nil(t, err, "There must be no errors")
	assert.Equal(t, &Response{
		Version: "42",
	}, response, "Response should be same as expected")
}

func TestSearchWithError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	api := GBClient{server.Client(), server.URL, ""}
	response, err := api.Search("test", 10, 1, []string{ResourceTypeGame})

	assert.NotNil(t, err, "Error must be set")
	assert.Nil(t, response, "Response shouldn't be set")
}

func TestPlatforms(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`
			{
				"version": "42"
			}
		`))
	}))
	defer server.Close()

	api := GBClient{server.Client(), server.URL, ""}
	response, err := api.Platforms(10, 0)

	assert.Nil(t, err, "There must be no errors")
	assert.Equal(t, &Response{
		Version: "42",
	}, response, "Response should be same as expected")
}

func TestPlatformsWithError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	api := GBClient{server.Client(), server.URL, ""}
	response, err := api.Platforms(10, 0)

	assert.NotNil(t, err, "Error must be set")
	assert.Nil(t, response, "Response shouldn't be set")
}

func TestPlatformsWithInvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`
			this is not JSON...
		`))
	}))
	defer server.Close()

	api := GBClient{server.Client(), server.URL, ""}
	response, err := api.Platforms(10, 0)

	assert.NotNil(t, err, "Error must be set")
	assert.Nil(t, response, "Response shouldn't be set")
}

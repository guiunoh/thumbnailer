package httpclient_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"thumbnailer/pkg/httpclient"

	"github.com/stretchr/testify/assert"
)

func TestDoRequest(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}))
	defer server.Close()

	// Create a new httpclient
	client := httpclient.DefaultHttpClient()

	// Do a request to the test server
	req, err := http.NewRequest("GET", server.URL, nil)
	assert.NoError(t, err)

	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

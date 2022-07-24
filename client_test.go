package dictionaryapi

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

var (
	client *Client
)

func init() {
	client = NewClient()
}

func TestClient_NewClient(t *testing.T) {

	t.Run("default", func(t *testing.T) {
		client := NewClient()
		assert.NotNil(t, client)
	})

	t.Run("with custom httpClient", func(t *testing.T) {
		httpClient := resty.New()
		client := NewClient(WithHTTPClient(httpClient))
		assert.NotNil(t, client)
		assert.Exactly(t, httpClient, client.httpClient)
	})

	t.Run("with custom WordManager", func(t *testing.T) {
		wm := &wordManager{Client: client}
		client := NewClient(WithWordsManager(wm))
		assert.NotNil(t, client)
		assert.Exactly(t, wm, client.Word)
	})
}

package dictionaryapi

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/privatesquare/bkst-go-utils/utils/fileutils"
	"github.com/stretchr/testify/assert"
)

const (
	word                = "hello"
	successTestDataPath = "data/test/success.json"
	errorTestDataPath   = "data/test/error.json"
)

func TestWord_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		httpmock.ActivateNonDefault(client.httpClient.GetClient())
		defer httpmock.DeactivateAndReset()

		var out []Word
		loadTestData(t, successTestDataPath, &out)
		responder, err := httpmock.NewJsonResponder(http.StatusOK, out)
		assert.NoError(t, err)
		httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(wordEntriesApiPath, word), responder)

		result, restErr := client.Word.Get(word)
		assert.Nil(t, restErr)
		assert.NotNil(t, result)
		assert.Equal(t, word, result.Word)
		assert.Equal(t, 3, len(result.Meanings))
	})

	t.Run("no definitions found", func(t *testing.T) {
		httpmock.ActivateNonDefault(client.httpClient.GetClient())
		defer httpmock.DeactivateAndReset()

		var out Error
		loadTestData(t, errorTestDataPath, &out)
		responder, err := httpmock.NewJsonResponder(http.StatusNotFound, out)
		assert.NoError(t, err)
		httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(wordEntriesApiPath, word), responder)

		result, restErr := client.Word.Get(word)
		assert.Nil(t, result)
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusNotFound, restErr.StatusCode)
		assert.Equal(t, "Sorry pal, we couldn't find definitions for the word you were looking for.", restErr.Message)
		assert.Equal(t, "No Definitions Found", restErr.Error)
	})

	t.Run("http request error", func(t *testing.T) {
		httpmock.ActivateNonDefault(client.httpClient.GetClient())
		defer httpmock.DeactivateAndReset()

		result, restErr := client.Word.Get(word)
		assert.Nil(t, result)
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusInternalServerError, restErr.StatusCode)
	})
}

func loadTestData(t *testing.T, path string, out interface{}) {
	err := fileutils.ReadJsonFile(path, &out)
	assert.NoError(t, err)
}

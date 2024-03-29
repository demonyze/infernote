package fetcher

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/demonyze/infernote/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestGetLanguage(t *testing.T) {
	mockImport := model.LanguageImport{
		Language: "en",
		Types: map[string]string{
			"major": "Major",
			"minor": "Minor",
			"5":     "5th",
			"7":     "7th",
		},
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockImport)
	}))
	defer server.Close()

	t.Run("Success", func(t *testing.T) {

		langImport, err := GetLanguage("en", server.URL)

		assert.NoError(t, err)
		assert.Equal(t, mockImport, langImport)
	})

	t.Run("Error - Unsupported language", func(t *testing.T) {
		_, err := GetLanguage("unsupported", server.URL)
		assert.Error(t, err)
	})

	t.Run("Error - Unsupported url", func(t *testing.T) {
		_, err := GetLanguage("en", "")
		assert.Error(t, err)
	})
}

package fetcher

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetJson(t *testing.T) {
	type MockResponse struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	serverSuccess := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockResponse := MockResponse{
			Id:   1,
			Name: "Mock",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer serverSuccess.Close()

	serverError := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer serverError.Close()

	t.Run("Success", func(t *testing.T) {
		res, err := GetJson[MockResponse](serverSuccess.URL)
		assert.NoError(t, err)
		assert.Equal(t, "Mock", res.Name)
		assert.Equal(t, 1, res.Id)
	})

	t.Run("Error - Url", func(t *testing.T) {
		_, err := GetJson[MockResponse]("")
		assert.Error(t, err)
	})

	t.Run("Error - Response", func(t *testing.T) {
		_, err := GetJson[MockResponse](serverError.URL)
		assert.Error(t, err)
	})
}

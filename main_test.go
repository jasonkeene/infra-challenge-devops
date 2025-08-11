package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := healthHandler()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestInfoHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/info", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := infoHandler()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response InfoResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Go DevOps Challenge App", response.Message)
}

func TestGetEnv(t *testing.T) {
	// Test default value
	value := getEnv("NONEXISTENT", "default")
	assert.Equal(t, "default", value)

	// Test environment variable
	err := os.Setenv("TEST_VAR", "test_value")
	assert.NoError(t, err)
	value = getEnv("TEST_VAR", "default")
	assert.Equal(t, "test_value", value)
}

func TestNYTimeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/nytime", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := nyTimeHandler()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "ny_time")
	assert.Contains(t, response, "ny_timezone")
	assert.Contains(t, response, "timestamp")
	assert.Equal(t, "America/New_York", response["ny_timezone"])
}

func TestFetchHandler(t *testing.T) {
	// Test with default URL
	req, err := http.NewRequest("GET", "/fetch", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := fetchHandler("dummy-secret-key-12345")
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "status")
	assert.Contains(t, response, "target_url")
	assert.Contains(t, response, "timestamp")
	assert.Equal(t, "success", response["status"])
}

func TestFetchHandlerWithAPIKey(t *testing.T) {
	// Test that the API key is properly set in the request
	apiKey := "test-api-key-12345"

	req, err := http.NewRequest("GET", "/fetch", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := fetchHandler(apiKey)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response["status"])
	assert.Equal(t, "https://httpbin.org/json", response["target_url"])
	assert.Contains(t, response, "timestamp")
}

func TestGetEnvAPIKey(t *testing.T) {
	// Test API_KEY environment variable handling
	originalAPIKey := os.Getenv("API_KEY")

	// Test with API_KEY set
	err := os.Setenv("API_KEY", "custom-api-key")
	assert.NoError(t, err)
	value := getEnv("API_KEY", "default-key")
	assert.Equal(t, "custom-api-key", value)

	// Test with API_KEY not set (should return default)
	err = os.Unsetenv("API_KEY")
	assert.NoError(t, err)
	value = getEnv("API_KEY", "default-key")
	assert.Equal(t, "default-key", value)

	// Restore original API_KEY value
	if originalAPIKey != "" {
		err = os.Setenv("API_KEY", originalAPIKey)
		assert.NoError(t, err)
	} else {
		err = os.Unsetenv("API_KEY")
		assert.NoError(t, err)
	}
}

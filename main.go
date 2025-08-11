package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type HealthResponse struct {
	Status      string    `json:"status"`
	Timestamp   time.Time `json:"timestamp"`
	Version     string    `json:"version"`
	Environment string    `json:"environment"`
}

type InfoResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Hostname  string    `json:"hostname"`
}

func main() {
	// Get environment variables
	port := getEnv("PORT", "8080")
	environment := getEnv("ENVIRONMENT", "development")
	version := getEnv("VERSION", "1.0.0")
	apiKey := getEnv("API_KEY", "dummy-secret-key-12345")

	log.Printf("Starting server on port %s", port)
	log.Printf("Environment: %s", environment)
	log.Printf("Version: %s", version)

	http.HandleFunc("/health", healthHandler())
	http.HandleFunc("/info", infoHandler())
	http.HandleFunc("/nytime", nyTimeHandler())
	http.HandleFunc("/fetch", fetchHandler(apiKey))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func healthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}
}

func infoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}
		response := InfoResponse{
			Message:   "Go DevOps Challenge App",
			Timestamp: time.Now(),
			Hostname:  hostname,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Error encoding response: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}

func nyTimeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nyLocation, err := time.LoadLocation("America/New_York")
		if err != nil {
			log.Printf("Error loading timezone: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		nyTime := time.Now().In(nyLocation)
		response := map[string]interface{}{
			"ny_time":     nyTime.Format("2006-01-02 15:04:05 MST"),
			"ny_timezone": "America/New_York",
			"timestamp":   nyTime.Unix(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Error encoding response: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}

func fetchHandler(apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get URL from query parameter, default to a public API
		targetURL := "https://httpbin.org/json"

		// Create HTTP client with timeout
		client := &http.Client{
			Timeout: 10 * time.Second,
		}

		// Create request with custom headers including API key
		req, err := http.NewRequest("GET", targetURL, nil)
		if err != nil {
			http.Error(w, "Failed to create request: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Add API key and other headers
		req.Header.Set("X-API-Key", apiKey)

		// Make HTTP request
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Failed to fetch data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer func() {
			if closeErr := resp.Body.Close(); closeErr != nil {
				log.Printf("Error closing response body: %v", closeErr)
			}
		}()

		// Just return success status without the response data
		response := map[string]interface{}{
			"status":     "success",
			"target_url": targetURL,
			"timestamp":  time.Now().Unix(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Error encoding response: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}

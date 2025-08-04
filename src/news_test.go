// go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCallNewsAPI_Success(t *testing.T) {
	// Mock News API server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"articles": [
				{"title": "title1"},
				{"title": "title2"}
			]
		}`)
	}))
	defer server.Close()

	requestURL := server.URL
	NewsApiClient := &NewsApiClient{NewsHttpClient: &http.Client{}}
	result, err := NewsApiClient.CallNewsApi(requestURL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == "" || result == "---news---" {
		t.Errorf("expected news content, got %q", result)
	}
	if want := "title1"; !contains(result, want) {
		t.Errorf("expected result to contain %q, got %q", want, result)
	}
}

func TestCallNewsAPI_HTTPError(t *testing.T) {
	// Close server before request to simulate connection error
	server := httptest.NewServer(nil)
	server.Close()

	requestURL := server.URL
	NewsApiClient := &NewsApiClient{NewsHttpClient: &http.Client{}}
	_, err := NewsApiClient.CallNewsApi(requestURL)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestCallNewsAPI_JSONError(t *testing.T) {
	// Mock server returns invalid JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `invalid json`)
	}))
	defer server.Close()

	requestURL := server.URL
	NewsApiClient := &NewsApiClient{NewsHttpClient: &http.Client{}}
	_, err := NewsApiClient.CallNewsApi(requestURL)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

// Helper for substring check
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && (contains(s[1:], substr) || s[:len(substr)] == substr))
}

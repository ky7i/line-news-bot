//go:build testonly
// +build testonly

package main

import (
	"net/url"
	"testing"
)

func TestBuildNewsAPIURL_Success(t *testing.T) {
	fileName := "../testdate/newsParams-success.json"
	baseURL := "http://test.com"
	apiKey := "test"
	newsAPIURL, err := BuildNewsAPIURL(fileName, baseURL, apiKey)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	u, err := url.Parse(newsAPIURL)
	if err != nil {
		t.Fatalf("failed to parse built URL: %v", err)
	}
	q := u.Query()
	if q.Get("q") != "query" {
		t.Errorf("expected q=query, got %v", q.Get("q"))
	}
	if q.Get("sortBy") != "sortBy" {
		t.Errorf("expected sortBy=sortBy, got %v", q.Get("sortBy"))
	}
	if q.Get("pageSize") != "pageSize" {
		t.Errorf("expected pageSize=pageSize, got %v", q.Get("pageSize"))
	}
	if q.Get("language") != "language" {
		t.Errorf("expected language=language, got %v", q.Get("language"))
	}
	if q.Get("apiKey") != "apiKey" {
		t.Errorf("expected apiKey=apiKey, got %v", q.Get("apiKey"))
	}
	if u.Scheme != "http" || u.Host != "test.com" {
		t.Errorf("unexpected base url: %v", newsAPIURL)
	}
}

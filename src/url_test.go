package main

import (
	"net/url"
	"strings"
	"testing"
)

func TestBuildNewsAPIURL_Success(t *testing.T) {
	fileName := "../testdate/newsParams-success.json"
	baseURL := "http://test.com"
	apiKey := "apiKey"
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

func TestBuildNewsAPIURL_FileNotFound(t *testing.T) {
	fileName := "invalid.json"
	baseURL := "http://test.com"
	apiKey := "apiKey"
	_, err := BuildNewsAPIURL(fileName, baseURL, apiKey)
	if err == nil {
		t.Errorf("expected file not found error, got no errors")
	}
	if err != nil && !strings.Contains(err.Error(), "failed to open params file:") {
		t.Errorf("expected the message 'failed to open params file: <<error>>', got this : %v", err)
	}
}

func TestBuildNewsAPIURL_FileEmpty(t *testing.T) {
	fileName := "../testdate/newsParams-empty.json"
	baseURL := "http://test.com"
	apiKey := "apiKey"
	_, err := BuildNewsAPIURL(fileName, baseURL, apiKey)
	if err == nil {
		t.Errorf("expected file empty error, got no errors")
	}
	if err != nil && !strings.Contains(err.Error(), "params file is empty") {
		t.Errorf("expected the message 'params file is empty', got this : %v", err)
	}
}

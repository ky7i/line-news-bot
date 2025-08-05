package main

import (
	"testing"
)

func TestBuildNewsAPIURL_Success(t *testing.T) {
	newsAPIParams := NewsAPIParams{
		BaseURL:  "http://test.com",
		Query:    "query",
		SortBy:   "sortBy",
		PageSize: "pageSize",
		Language: "language",
		APIKey:   "apiKey",
	}
	newsAPIURL, err := BuildNewsAPIURL(newsAPIParams)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if newsAPIURL != "http://test.com?apiKey=apiKey&language=language&pageSize=pageSize&q=query&sortBy=sortBy" {
		t.Errorf("built URL mismatch. \r\nbuilt URL : %v", newsAPIURL)
	}
}

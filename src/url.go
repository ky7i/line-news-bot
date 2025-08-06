package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

type Param struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type NewsParamsFile struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Reference   string  `json:"reference"`
	Params      []Param `json:"params"`
}

// fileName: path to newsParams.json
// baseURL: NewsAPI endpoint base URL (e.g. https://newsapi.org/v2/everything)
// apiKey: API key string (should not be in json file)
func BuildNewsAPIURL(fileName, baseURL, apiKey string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to open params file: %w", err)
	}
	defer file.Close()

	var paramsFile NewsParamsFile
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&paramsFile); err != nil {
		return "", fmt.Errorf("failed to decode params json: %w", err)
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid baseURL: %w", err)
	}
	q := u.Query()
	for _, p := range paramsFile.Params {
		q.Set(p.Key, p.Value)
	}
	q.Set("apiKey", apiKey)
	u.RawQuery = q.Encode()
	return u.String(), nil
}

//	u, err := url.Parse(p.BaseURL)
//	if err != nil {
//		fmt.Println(err)
//		return "", err
//	}
//	q := u.Query()
//	q.Set("q", p.Query)
//	q.Set("sortBy", p.SortBy)
//	q.Set("pageSize", p.PageSize)
//	q.Set("language", p.Language)
//	q.Set("apiKey", p.APIKey)
//	u.RawQuery = q.Encode()
//	return u.String(), nil

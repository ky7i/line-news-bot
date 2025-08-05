package main

import (
	"fmt"
	"net/url"
)

type NewsAPIParams struct {
	BaseURL   string
	Query     string
	SortBy    string
	PageSize  string
	Language  string
	APIKey    string
}

func BuildNewsAPIURL(p NewsAPIParams) (string, error) {
	u, err := url.Parse(p.BaseURL)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	q := u.Query()
	q.Set("q", p.Query)
	q.Set("sortBy", p.SortBy)
	q.Set("pageSize", p.PageSize)
	q.Set("language", p.Language)
	q.Set("apiKey", p.APIKey)
	u.RawQuery = q.Encode()
	return u.String(), nil
}
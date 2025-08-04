package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type NewsCaller interface {
	CallNewsApi(requestURL string) (string, error)
}

type NewsHttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type NewsApiClient struct {
	NewsHttpClient NewsHttpClient
}

func (n *NewsApiClient) CallNewsApi(requestURL string) (string, error) {
	res, err := n.NewsHttpClient.Get(requestURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return "", err
	}

	// TODO: extracts into a function
	contents := "---news---"
	articles := result["articles"].([]interface{})

	for i := 0; i < len(articles); i++ {
		content := articles[i].(map[string]interface{})["title"]
		contentStr, ok := content.(string)
		if !ok {
			continue
		}
		contents = contents + "\r\n" + contentStr
	}
	return contents, nil
}

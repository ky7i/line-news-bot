package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type NewsHttpClient interface {
	Get(url string) (resp *http.Response, err error)	
}

type NewsApiClient struct {
	NewsHttpClient NewsHttpClient
}

func (n *NewsApiClient) CallNewsAPI(requestURL string) (string, error) {
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

	// NewsAPIのレスポンスの整形
	// TODO: 別関数への切り出し
	contents := "---news---"
	articles := result["articles"].([]interface{})

	for i := 0; i < len(articles); i++ {
		content := articles[i].(map[string]interface{})["title"]
		// titleのvalueがnullのケースへのバリデーション
		contentStr, ok := content.(string)
		if !ok {
			continue
		}
		contents = contents + "\r\n" + contentStr
	}
	return contents, nil
}

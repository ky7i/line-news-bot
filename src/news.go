package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func CallNewsAPI(requestURL string) (string, error) {
	fmt.Println("NewsAPI requestURL : ", requestURL)
	res, err := http.Get(requestURL)
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

	contents := "---news---"
	articles := result["articles"].([]interface{})

	for i := 0; i < len(articles); i++ {
		content := articles[i].(map[string]interface{})["title"]
		// titleのvalueがnullのケースへのバリデーション
		contentStr, ok := content.(string)
		if !ok {
			continue
		}

		// NewsAPIのcontentに改行文字が含まれている。
		// おそらく、contentは "Summary\r\nDetail" という仕様
		// Detailは文字数が多いため使用しない
		// TODO : 仕様の詳細を把握
		contentFormatted, _, _ := strings.Cut(contentStr, "\r\n")
		contents = contents + "\r\n" + contentFormatted
	}
	return contents, nil
}

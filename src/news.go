package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func CallNewsAPI(NEWS_API_URI string, NEWS_API_PARAMETER string, NEWS_API_KEY string) (string, error) {
	res, err := http.Get(NEWS_API_URI + NEWS_API_PARAMETER + NEWS_API_KEY)
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
		contentStr, err := content.(string)
		if err != false {
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

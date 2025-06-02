package main

import (
	"strings"
	"io"
	"os"
	"net/http"
	"encoding/json"
)

func CallNewsAPI() (string, error){
	// TODO : mainで取得し、引数で渡されるようにする
	// 環境変数の取得
	NEWS_API_URI := os.Getenv("NEWS_API_URI")
	NEWS_API_PARAMETER := os.Getenv("NEWS_API_PARAMETER")
	NEWS_API_KEY := os.Getenv("NEWS_API_KEY")

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
		content := articles[i].(map[string]interface{})["content"].(string)
		// NewsAPIのcontentに改行文字が含まれている。
		// おそらく、contentは "Summary\r\nDetail" という仕様
		// Detailは文字数が多いため使用しない
		// TODO : 仕様の詳細を把握
		contentFormatted, _, _ := strings.Cut(content, "\r\n")
		contents = contents + "\r\n" + contentFormatted
	}
	return contents, nil
}

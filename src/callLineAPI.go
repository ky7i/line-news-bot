package main

import (
	"os"
	"encoding/json"
	"net/http"
	"bytes"
)

func CallLineAPI(inputText string) error {
	// TODO : mainで取得し、引数で渡されるようにする
	// 環境変数の取得
	LINE_API_ACCESS_TOKEN := os.Getenv("LINE_API_ACCESS_TOKEN")
	LINE_API_USER_ID := os.Getenv("LINE_API_USER_ID")
	LINE_API_URI := os.Getenv("LINE_API_URI")

	// HTTPリクエストのBody部作成
	payload := map[string]interface{} {
		"to": LINE_API_USER_ID,
		"messages": []map[string]string {
			{
				"type": "text",
				"text": inputText,
			},
		},
	}

	jsonDate, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", LINE_API_URI, bytes.NewReader(jsonDate))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	// 認証情報はヘッダーに設定
	req.Header.Add("Authorization", "Bearer " + LINE_API_ACCESS_TOKEN)

	// 送信処理
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}
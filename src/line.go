package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func CallLineAPI(LINE_API_URI string, LINE_API_USER_ID string, LINE_API_ACCESS_TOKEN string, inputText string) error {
	// HTTPリクエストのBody部作成
	payload := map[string]interface{}{
		"to": LINE_API_USER_ID,
		"messages": []map[string]string{
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
	req.Header.Add("Authorization", "Bearer "+LINE_API_ACCESS_TOKEN)

	// 送信処理
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type RequestCreator interface {
	NewRequest(method, url string, body io.Reader) (*http.Request, error)
}

type DefaultRequestCreator struct{}

func (d *DefaultRequestCreator) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}

type LineApiClient struct {
	RequestCreator RequestCreator
	LineHttpClient *http.Client
}

func (l *LineApiClient) CallLineAPI(LINE_API_URI string, LINE_API_USER_ID string, LINE_API_ACCESS_TOKEN string, inputText string) error {
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

	req, err := l.RequestCreator.NewRequest("POST", LINE_API_URI, bytes.NewReader(jsonDate))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	// 認証情報はヘッダーに設定
	req.Header.Add("Authorization", "Bearer "+LINE_API_ACCESS_TOKEN)

	// 送信処理
	_, err = l.LineHttpClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}

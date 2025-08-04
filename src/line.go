package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// In tests, this interface make it easy to replace CallLineApi with mocks.
type LineCaller interface {
	CallLineApi(LINE_API_URI string, LINE_API_USER_ID string, LINE_API_ACCESS_TOKEN string, inputText string) error
}

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

// push <inputText> to LINE
func (l *LineApiClient) CallLineApi(LINE_API_URI string, LINE_API_USER_ID string, LINE_API_ACCESS_TOKEN string, inputText string) error {
	// request Body
	payload := CreateRequestBody(LINE_API_USER_ID, inputText)

	jsonDate, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := l.RequestCreator.NewRequest("POST", LINE_API_URI, bytes.NewReader(jsonDate))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+LINE_API_ACCESS_TOKEN)

	_, err = l.LineHttpClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func CreateRequestBody(LINE_API_USER_ID string, inputText string) map[string]any {
	return map[string]any{
		"to": LINE_API_USER_ID,
		"messages": []map[string]string{
			{
				"type": "text",
				"text": inputText,
			},
		},
	}
}

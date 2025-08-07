// go
package main

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

type MockNewsCaller struct {
	CallFunc func(requestURL string) (string, error)
}

func (n *MockNewsCaller) CallNewsApi(requestURL string) (string, error) {
	return n.CallFunc(requestURL)
}

type MockLineCaller struct {
	CallFunc func(LINE_API_URI string, LINE_API_USER_ID string, LINE_API_ACCESS_TOKEN string, inputText string) error
}

func (l *MockLineCaller) CallLineApi(LINE_API_URI string, LINE_API_USER_ID string, LINE_API_ACCESS_TOKEN string, inputText string) error {
	return l.CallFunc(LINE_API_URI, LINE_API_USER_ID, LINE_API_ACCESS_TOKEN, inputText)
}

func TestHandler_Success(t *testing.T) {
	mockNewsCaller := &MockNewsCaller{
		CallFunc: func(requestURL string) (string, error) {
			// mock of NewsAPI : Success
			return "news title", nil
		},
	}

	mockLineCaller := &MockLineCaller{
		CallFunc: func(LINE_API_URI string, LINE_API_USER_ID string, LINE_API_ACCESS_TOKEN string, inputText string) error {
			// mock of LineAPI : Success
			if inputText != "news title" {
				return errors.New("news content mismatch")
			}
			return nil
		},
	}

	resp, err := handlerWithDeps(context.Background(), events.APIGatewayProxyRequest{}, mockNewsCaller, mockLineCaller)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestHandler_CallNewsAPIError(t *testing.T) {
	mockNewsCaller := &MockNewsCaller{
		CallFunc: func(requestURL string) (string, error) {
			// error in NewsAPI
			return "", errors.New("news error in a mock server")
		},
	}
	mockLineCaller := &MockLineCaller{
		CallFunc: func(LINE_API_URI, LINE_API_USER_ID, LINE_API_ACCESS_TOKEN, inputText string) error {
			return nil
		},
	}
	resp, err := handlerWithDeps(context.Background(), events.APIGatewayProxyRequest{}, mockNewsCaller, mockLineCaller)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestHandler_CallLineAPIError(t *testing.T) {
	mockNewsCaller := &MockNewsCaller{
		CallFunc: func(requestURL string) (string, error) {
			return "news", nil
		},
	}
	mockLineCaller := &MockLineCaller{
		CallFunc: func(LINE_API_URI, LINE_API_USER_ID, LINE_API_ACCESS_TOKEN, inputText string) error {
			// error in Line API
			return errors.New("line error in a mock server")
		},
	}
	resp, err := handlerWithDeps(context.Background(), events.APIGatewayProxyRequest{}, mockNewsCaller, mockLineCaller)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

// Additional tests
func TestHandler_EmptyNews(t *testing.T) {
	mockNewsCaller := &MockNewsCaller{
		CallFunc: func(requestURL string) (string, error) {
			// value is empty
			return "", nil
		},
	}
	mockLineCaller := &MockLineCaller{
		CallFunc: func(LINE_API_URI, LINE_API_USER_ID, LINE_API_ACCESS_TOKEN, inputText string) error {
			return nil
		},
	}
	resp, err := handlerWithDeps(context.Background(), events.APIGatewayProxyRequest{}, mockNewsCaller, mockLineCaller)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestHandler_LongNews(t *testing.T) {
	longNews := ""
	for i := 0; i < 1000; i++ {
		longNews += "news "
	}
	mockNewsCaller := &MockNewsCaller{
		CallFunc: func(requestURL string) (string, error) {
			return longNews, nil
		},
	}
	mockLineCaller := &MockLineCaller{
		CallFunc: func(LINE_API_URI, LINE_API_USER_ID, LINE_API_ACCESS_TOKEN, inputText string) error {
			if inputText != longNews {
				return errors.New("news content mismatch")
			}
			return nil
		},
	}
	resp, err := handlerWithDeps(context.Background(), events.APIGatewayProxyRequest{}, mockNewsCaller, mockLineCaller)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

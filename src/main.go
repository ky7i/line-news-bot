package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "\"Hello from Lambda!\"",
	}

	// 環境変数
	// Googleカレンダー
	// AWS_SECRET_MANAGER_NAME := os.Getenv("AWS_SECRET_MANAGER_NAME")
	// AWS_SECRET_MANAGER_REGION := os.Getenv("AWS_SECRET_MANAGER_REGION")

	// LINE_API
	LINE_API_ACCESS_TOKEN := os.Getenv("LINE_API_ACCESS_TOKEN")
	LINE_API_USER_ID := os.Getenv("LINE_API_USER_ID")
	LINE_API_URI := os.Getenv("LINE_API_URI")

	// NewsAPI
	NEWS_API_URI := os.Getenv("NEWS_API_URI")
	NEWS_API_PARAMETER := os.Getenv("NEWS_API_PARAMETER")
	NEWS_API_KEY := os.Getenv("NEWS_API_KEY")

	news, err := CallNewsAPI(NEWS_API_URI, NEWS_API_PARAMETER, NEWS_API_KEY)
	if err != nil {
		fmt.Println(err)
		return response, err
	}

	err = CallLineAPI(LINE_API_URI, LINE_API_USER_ID, LINE_API_ACCESS_TOKEN, news)
	if err != nil {
		fmt.Println(err)
		return response, err
	}

	return response, nil
}

func main() {
	lambda.Start(handler)
}

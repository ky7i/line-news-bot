package main

import (
	"context"
	"fmt"
	"net/url"
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
	AWS_SECRET_MANAGER_NAME := os.Getenv("AWS_SECRET_MANAGER_NAME")
	AWS_SECRET_MANAGER_REGION := os.Getenv("AWS_SECRET_MANAGER_REGION")

	// LINE_API
	LINE_API_ACCESS_TOKEN := os.Getenv("LINE_API_ACCESS_TOKEN")
	LINE_API_USER_ID := os.Getenv("LINE_API_USER_ID")
	LINE_API_URI := os.Getenv("LINE_API_URI")

	// NewsAPI
	NEWS_API_BASE_URL := os.Getenv("NEWS_API_BASE_URL")

	credential, err := GetSecretString(AWS_SECRET_MANAGER_NAME, AWS_SECRET_MANAGER_REGION)
	if err != nil {
		fmt.Println(err)
		return response, err
	}
	schedule := getCalendar(credential)
	fmt.Println("schedule : ", schedule)

	params := url.Values{}
	params.Add("q", os.Getenv("NEWS_API_QUERY"))
	params.Add("sortBy", os.Getenv("NEWS_API_SORT_BY"))
	params.Add("pageSize", os.Getenv("NEWS_API_PAGE_SIZE"))
	params.Add("language", os.Getenv("NEWS_API_LANGUAGE"))
	params.Add("apiKey", os.Getenv("NEWS_API_KEY"))

	news, err := CallNewsAPI(NEWS_API_BASE_URL + "?" + params.Encode())
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

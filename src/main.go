package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	newsApiClient NewsApiClient
	lineApiClient LineApiClient
)

// initialize API clients
func init() {
	newsApiClient = NewsApiClient{NewsHttpClient: &http.Client{}}
	lineApiClient = LineApiClient{
		RequestCreator: &DefaultRequestCreator{},
		LineHttpClient: &http.Client{},
	}
}

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// In tests, replace the second and third arguments to mocks.
	// In production, clients initialized in init() are used
	return handlerWithDeps(ctx, event, &newsApiClient, &lineApiClient)
}

func handlerWithDeps(ctx context.Context, event events.APIGatewayProxyRequest, newsCaller NewsCaller, lineCaller LineCaller) (events.APIGatewayProxyResponse, error) {
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "\"Hello from Lambda!\"",
	}

	// environment variable

	// NewsAPI
	NEWS_API_BASE_URL := os.Getenv("NEWS_API_BASE_URL")
	NEWS_API_KEY := os.Getenv("NEWS_API_KEY")

	// GoogleCalendar
	// AWS_SECRET_MANAGER_NAME := os.Getenv("AWS_SECRET_MANAGER_NAME")
	// AWS_SECRET_MANAGER_REGION := os.Getenv("AWS_SECRET_MANAGER_REGION")

	// LINE_API
	LINE_API_ACCESS_TOKEN := os.Getenv("LINE_API_ACCESS_TOKEN")
	LINE_API_USER_ID := os.Getenv("LINE_API_USER_ID")
	LINE_API_URI := os.Getenv("LINE_API_URI")

	// credential, err := GetSecretString(AWS_SECRET_MANAGER_NAME, AWS_SECRET_MANAGER_REGION)
	// if err != nil {
	// fmt.Println(err)
	// return response, err
	// }
	// schedule := getCalendar(credential)
	// fmt.Println("schedule : ", schedule)

	//	newsAPIParams := NewsAPIParams{
	//		BaseURL:  os.Getenv("NEWS_API_BASE_URL"),
	//		Query:    os.Getenv("NEWS_API_QUERY"),
	//		SortBy:   os.Getenv("NEWS_API_SORT_BY"),
	//		PageSize: os.Getenv("NEWS_API_PAGE_SIZE"),
	//		Language: os.Getenv("NEWS_API_LANGUAGE"),
	//		APIKey:   os.Getenv("NEWS_API_KEY"),
	//	}

	// newsAPIURL, err := BuildNewsAPIURL(newsAPIParams)
	newsAPIURL, err := BuildNewsAPIURL("../newsParams.json", NEWS_API_BASE_URL, NEWS_API_KEY)
	if err != nil {
		fmt.Println(err)
	}

	news, err := newsCaller.CallNewsApi(newsAPIURL)
	if err != nil {
		fmt.Println(err)
		return response, err
	}

	err = lineCaller.CallLineApi(LINE_API_URI, LINE_API_USER_ID, LINE_API_ACCESS_TOKEN, news)
	if err != nil {
		fmt.Println(err)
		return response, err
	}
	return response, nil
}

func main() {
	lambda.Start(handler)
}

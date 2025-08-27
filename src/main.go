package main

import (
	"context"
	"log"
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

// TODO: consider this function signature
// reference: https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/golang-handler.html#golang-handler-signatures
func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// In tests, replace the second and third arguments to mocks.
	// In production, clients initialized in init() are used
	return handlerWithDeps(ctx, event, &newsApiClient, &lineApiClient)
}

func handlerWithDeps(ctx context.Context, event events.APIGatewayProxyRequest, newsCaller NewsCaller, lineCaller LineCaller) (events.APIGatewayProxyResponse, error) {
	// TODO: Add more response variation, like 5XX
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "\"Hello from Lambda!\"",
	}

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

	NEWS_PARAMETER_FILE := os.Getenv("NEWS_PARAMETER_FILE")
	newsAPIURL, err := BuildNewsAPIURL(NEWS_PARAMETER_FILE, NEWS_API_BASE_URL, NEWS_API_KEY)
	if err != nil {
		log.Println(err)
		return response, err
	}

	news, err := newsCaller.CallNewsApi(newsAPIURL)
	if err != nil {
		log.Println(err)
		return response, err
	}

	err = lineCaller.CallLineApi(LINE_API_URI, LINE_API_USER_ID, LINE_API_ACCESS_TOKEN, news)
	if err != nil {
		log.Println(err)
		return response, err
	}
	return response, nil
}

func main() {
	lambda.Start(handler)
}

package main

import (
	"fmt"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "\"Hello from Lambda!\"",
	}

	news, err := CallNewsAPI()
	if err != nil {
		fmt.Println(err)
		return response, err
	}

	err = CallLineAPI(news)
	if err != nil {
		fmt.Println(err)
		return response, err
	}

	return response, nil
}

func main() {
	lambda.Start(handler)
}
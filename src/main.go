package main

import (
	"fmt"
	// "github.com/joho/godotenv"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler() {
	// TODO : パスをハードコーディングしたくない
	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	news, err := CallNewsAPI()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = CallLineAPI(news)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	lambda.Start(handler)
}
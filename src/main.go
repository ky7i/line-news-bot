package main

import (
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	// TODO : パスをハードコーディングしたくない
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
		return
	}

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
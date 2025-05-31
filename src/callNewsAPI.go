package main

import (
	"strings"
	"io"
	"os"
	"fmt"
	"net/http"
	"encoding/json"
)

func CallNewsAPI() string {
	NEWS_API_URI := os.Getenv("NEWS_API_URI")
	NEWS_API_KEY := os.Getenv("NEWS_API_KEY")

	res, err := http.Get(NEWS_API_URI + NEWS_API_KEY)
	if err != nil {
		fmt.Println("error occured at a API request.", err)
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error occured at Read reaponse body.", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		fmt.Println("error occured at unmarshaling response body formatted by json.", err)
	}
	
	contents := "---news---"
	articles := result["articles"].([]interface{})

	for i := 0; i < len(articles); i++ {
		content := articles[i].(map[string]interface{})["content"].(string)
		contentFormatted, _, _ := strings.Cut(content, "\r\n")
		contents = contents + "\r\n" + contentFormatted
	}
	return contents
}

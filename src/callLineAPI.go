package main

import (
	"os"
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
)

func CallLineAPI(inputText string) {
	LINE_API_ACCESS_TOKEN := os.Getenv("LINE_API_ACCESS_TOKEN")
	LINE_API_USER_ID := os.Getenv("LINE_API_USER_ID")
	LINE_API_URI := os.Getenv("LINE_API_URI")

	payload := map[string]interface{} {
		"to": LINE_API_USER_ID,
		"messages": []map[string]string {
			{
				"type": "text",
				"text": inputText,
			},
		},
	}

	jsonDate, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("error occured at marcharing body of LINE API request.", err)
	}

	req, err := http.NewRequest("POST", LINE_API_URI, bytes.NewReader(jsonDate))
	if err != nil {
		fmt.Println("error occured ad creating LINE API request.", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + LINE_API_ACCESS_TOKEN)

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		fmt.Println("error occured at posting a LINE API request.", err)
	}
}
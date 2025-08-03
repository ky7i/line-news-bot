//go:build testonly
// +build testonly

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// Reference : https://developers.google.com/workspace/calendar/api/quickstart/go?hl=ja
// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	// tok, err := tokenFromFile(tokFile)
	if err != nil {
		// tok = getTokenFromWeb(config)
		// saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func getCalendar(credentialByte []byte) string {
	schedule := ""
	ctx := context.Background()
	// b, err := os.ReadFile("credentials.json")
	// if err != nil {
	// log.Fatalf("Unable to read client secret file: %v", err)
	// }

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(credentialByte, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := getClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	t := time.Now().Format(time.RFC3339)
	tomorrow := time.Now().AddDate(0, 0, 1).Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).TimeMax(tomorrow).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			// formatted in "hour:minute"
			t, err := time.Parse(time.RFC3339, date)
			if err != nil {
				fmt.Println("Time Parse Error: ", err)
			}
			hour, minute, _ := t.Clock()
			schedule += item.Summary + " " + fmt.Sprintf("%02d", hour) + ":" + fmt.Sprintf("%02d", minute) + "\n"
			fmt.Printf(schedule)
		}
	}
	return schedule
}

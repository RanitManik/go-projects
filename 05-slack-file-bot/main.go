package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load .env")
	}

	token := os.Getenv("SLACK_BOT_TOKEN")
	channel := os.Getenv("CHANNEL_ID")

	if token == "" || channel == "" {
		panic("Missing SLACK_BOT_TOKEN or CHANNEL_ID")
	}

	data, err := os.ReadFile("SampleFile.md")
	if err != nil {
		panic(err)
	}
	if len(data) == 0 {
		panic("File is empty")
	}

	api := slack.New(token)

	params := slack.UploadFileV2Parameters{
		Channel:  channel,
		Filename: "sample-file.md",
		Title:    "Same Markdown File Upload",
		Reader:   bytes.NewReader(data),
		FileSize: len(data),
	}

	file, err := api.UploadFileV2(params)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Uploaded\nID: %s\nTitle: %s\n", file.ID, file.Title)
}
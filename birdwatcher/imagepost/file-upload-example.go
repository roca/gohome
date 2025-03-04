package main

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	imagePath := "./sworks.jpeg"

	params := slack.UploadFileV2Parameters{
		// Channel: os.Getenv("CHANNEL_ID"),
		// Channel: "general",
		File:     imagePath,
		Title:    "Sworks",
		Filename: "sworks.jpeg",
		FileSize: 303178,
	}

	file,err := api.UploadFileV2(params)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("File ID:", file.ID)

}

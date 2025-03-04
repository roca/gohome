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

	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatalln(err)
	}
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()

	params := slack.UploadFileV2Parameters{
		Channel: os.Getenv("CHANNEL_ID"),
		File:     imagePath,
		Title:    "Sworks",
		Filename: "sworks.jpeg",
		FileSize: int(fileSize),
	}

	uploadedFileInfo, err := api.UploadFileV2(params)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("File ID:", uploadedFileInfo.ID)

}

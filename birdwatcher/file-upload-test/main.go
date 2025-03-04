package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	imagePath := "./sworks.jpeg"

	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatalln(err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	params := slack.UploadFileV2Parameters{
		Channel:  os.Getenv("CHANNEL_ID"),
		File:     imagePath,
		Title:    filepath.Base(imagePath),
		Filename: fileInfo.Name(),
		FileSize: int(fileInfo.Size()),
	}

	uploadedFileInfo, err := api.UploadFileV2(params)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("File ID:", uploadedFileInfo.ID)

}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/slack-go/slack"
)

func main() {
	imagePath :=flag.String("imagePath", "", "Path to the image file to upload")
	flag.Parse()
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	
	if imagePath == nil || *imagePath == "" {
		log.Fatalln("Please provide the path to the image file to upload")
	}

	file, err := os.Open(*imagePath)
	if err != nil {
		log.Fatalln(err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	params := slack.UploadFileV2Parameters{
		Channel:  os.Getenv("CHANNEL_ID"),
		File:     *imagePath,
		Title:    filepath.Base(*imagePath),
		Filename: fileInfo.Name(),
		FileSize: int(fileInfo.Size()),
	}

	uploadedFileInfo, err := api.UploadFileV2(params)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("File ID:", uploadedFileInfo.ID)

}

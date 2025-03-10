package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/slack-go/slack"
	"github.com/stianeikeland/go-rpio/v4"
)

var api *slack.Client
var slackBotToken string
var channelID string

func init() {
	slackBotToken, ok := os.LookupEnv("SLACK_BOT_TOKEN")
	if !ok {
		fmt.Fprintln(os.Stderr, "'SLACK_BOT_TOKEN' env var is required")
		os.Exit(1)
	}
	channelID, ok = os.LookupEnv("CHANNEL_ID")
	if !ok {
		fmt.Fprintln(os.Stderr, "'CHANNEL_ID' env var is required")
		os.Exit(1)
	}

	api = slack.New(slackBotToken)
}

func main() {

	pin := rpio.Pin(18)

	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer rpio.Close()

	pin.Input()
	pin.PullUp()
	pin.Detect(rpio.FallEdge)

	fmt.Println("Sensing Enabled.")

	motionDetected := false

	for range time.Tick(500 * time.Millisecond) {
		if pin.EdgeDetected() {
			if motionDetected {
				motionDetected = false
				continue
			}

			motionDetected = true
			fmt.Println("Motion Detected!")

			go captureSendImage()
		}
	}
}

func captureSendImage() {
	capturedImage, err := captureImage()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to capture image:", err)
		return
	}

	if err := sendImage(capturedImage, slackBotToken, channelID); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to send image:", err)
		return
	}

	fmt.Println("Image sent successfully to Slack channel ID", channelID)
	os.Remove(capturedImage)
}

func captureImage() (string, error) {
	output, err := os.CreateTemp("", "capture*.jpg")
	if err != nil {
		return "", err
	}
	output.Close()

	cmd := exec.Command("libcamera-still", "--width", "1024", "--height", "768", "-o", output.Name())
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to capture image: %w", err)
	}

	return output.Name(), nil
}

func sendImage(imagePath, slackBotToken, channelID string) error {

	file, err := os.Open(imagePath)
	if err != nil {
		return err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	params := slack.UploadFileV2Parameters{
		Channel:  channelID,
		File:     imagePath,
		Title:    filepath.Base(imagePath),
		Filename: fileInfo.Name(),
		FileSize: int(fileInfo.Size()),
	}

	_, err = api.UploadFileV2(params)
	if err != nil {
		return err
	}

	return nil
}

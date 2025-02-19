package main

import (
	"fmt"
	"log"
	"os"

	hue "github.com/collinux/gohue"
)

func main() {
	HUE_ID := os.Getenv("HUE_ID")
	HUE_IP_ADDRESS := os.Getenv("HUE_IP_ADDRESS")

	bridge, _ := hue.NewBridge(HUE_IP_ADDRESS)

	bridge.Login(HUE_ID)

	lightStrip, err := bridge.GetLightByName("Hue lightstrip plus 1")
	if err != nil {
		log.Fatal(err)
	}
	lightStrip.Toggle()
}

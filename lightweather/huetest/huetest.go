package main

import (
	"log"
	"os"

	hue "github.com/collinux/gohue"
)

func main() {
	bridge, err := hue.NewBridge(os.Getenv("HUE_IP_ADDRESS"))
	if err != nil {
		log.Fatalln("Error connecting to bridge")
	}
	err = bridge.Login(os.Getenv("HUE_ID"))
	if err != nil {
		log.Fatalln("Error logging in to bridge")
	}

	lightStrip, err := bridge.GetLightByName("HueLightstrip01")
	if err != nil {
		log.Fatal(err)
	}
	lightStrip.Toggle()
}

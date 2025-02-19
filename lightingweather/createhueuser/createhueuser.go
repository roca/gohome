package main

import (
	"fmt"
	"log"
	"os"

	hue "github.com/collinux/gohue"
)

func main() {

	// bridgesOnNetwork, _ := hue.FindBridges()
	// if len(bridgesOnNetwork) == 0 {
	// 	log.Fatalln("No bridges found on network")
	// }
	// fmt.Println(bridgesOnNetwork[0].IPAddress)
	// bridge := bridgesOnNetwork[0]
	// fmt.Println(bridge.Info)

	bridge, err := hue.NewBridge(os.Getenv("HUE_IP_ADDRESS"))
	if err != nil {
		log.Fatalln("Error connecting to bridge")
	}
	username, _ := bridge.CreateUser("gohomeuser")
	fmt.Println(username)
}

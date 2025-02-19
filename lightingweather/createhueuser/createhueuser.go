package main

import (
	"fmt"
	"os"

	hue "github.com/collinux/gohue"
)

func main() {

	// bridgesOnNetwork, _ := hue.FindBridges()
	// if len(bridgesOnNetwork) == 0 {
	// 	fmt.Println("No bridges found on network")
	// 	return
	// }
	// fmt.Println(bridgesOnNetwork[0].IPAddress)
	// bridge := bridgesOnNetwork[0]

	bridge, err := hue.NewBridge(os.Getenv("HUE_IP_ADDRESS"))
	if err != nil {
		fmt.Println("Error connecting to bridge")
		return
	}
	username, _ := bridge.CreateUser("gohomeuser")
	fmt.Println(username)
}

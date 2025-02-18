package main

import (
	"fmt"

	hue "github.com/collinux/gohue"
)

func main() {
	bridgesOnNetwork, _ := hue.FindBridges()
	bridge := bridgesOnNetwork[0]
	username, _ := bridge.CreateUser("gohomeuser")
	fmt.Println(username)
}

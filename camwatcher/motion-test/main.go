package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

/*
GPIO pin 18 on the Pi Zero W is used to detect motion.
The pin is connected to a PIR sensor.
We will print a message to the console when motion is detected.
*/

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

	for range time.Tick(500 * time.Millisecond) {
		if pin.EdgeDetected() {
			fmt.Println("Motion Detected!")
		}
	}
}

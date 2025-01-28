package main

import (
	"log"

	"github.com/stianeikeland/go-rpio/v4"
)

type state rpio.state

func (s state) String() string {
	if s == state(rpio.Low) {
		return "Open"
	}
	return "Closed"
}

func setupGPIO(pinNumber int) (rpio.Pin, error) {
	if err := rpio.Open(); err != nil {
		log.Println("Error opening GPIO:", err)
		return 0, err
	}
	pin := rpio.Pin(pinNumber)

	pin.Input()
	rpio.PullMode(pin, rpio.PullUp)

	return pin, nil
}


func getDoorState(pin rpio.Pin) state {
	return state(pin.Read())
}
package main

import (
    "log"
    "machine"
    "time"
)

func main() {
    led := machine.LED
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    for {
        led.Low()
        time.Sleep(time.Millisecond * 1000)

        led.High()
        time.Sleep(time.Millisecond * 1000)

        log.Println("Hello, World!")
        machine.Serial.Write([]byte("Hello, World!\n"))
    }
}
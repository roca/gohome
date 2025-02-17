package  main

import (
	"fmt"
	"log"
	"os"

	owm "github.com/briandowns/openweathermap"
)

func main() {
	w, err := owm.NewCurrent("F", "EN", os.Getenv("OWM_API_KEY"))
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByZipcode(os.Getenv("ZIP_CODE"), os.Getenv("COUNTRY_CODE"))
	var currentTemp = w.Main.Temp

	switch {
	case currentTemp < 51:
		fmt.Println("Blue")
	case currentTemp >= 51 && currentTemp < 66:
		fmt.Println("Yellow")
	case currentTemp >= 66 && currentTemp < 80:
		fmt.Println("Green")
	case currentTemp >= 80 && currentTemp < 90:
		fmt.Println("Orange")
	case currentTemp >= 90:
		fmt.Println("Red")
	}
}

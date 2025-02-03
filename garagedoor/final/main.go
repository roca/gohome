package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

type state rpio.State

func (s state) String() string {
	if s == state(rpio.Low) {
		return "Open"
	}
	return "Closed"
}

func setupGPIO(pinNumber int) (rpio.Pin, error) {
	if err := rpio.Open(); err != nil {
		log.Println("Error opening GPIO:", err)
		return rpio.Pin(0), err
	}

	pin := rpio.Pin(pinNumber)

	pin.Input()
	rpio.PullMode(pin, rpio.PullUp)

	return pin, nil
}

func getDoorState(pin rpio.Pin) state {
	return state(pin.Read())
}

func isPeriod(start, end time.Time) bool {
	cur := time.Now().Format("15:04")

	now, err := time.Parse("15:04", cur)
	if err != nil {
		log.Println(err)
		return false
	}

	if end.Before(start) {
		end = end.Add(24 * time.Hour)
	}

	if now.Before(start) {
		now = now.Add(24 * time.Hour)
	}

	return now.After(start) && now.Before(end)
}

func sendNotification(slackWebhook, message string) {
	u, err := url.Parse(slackWebhook)
	if err != nil {
		log.Println("Invalid Slack webhook URL", err)
		return
	}

	v := url.Values{}
	v.Set("wait", "true")
	u.RawQuery = v.Encode()

	payload := struct {
		Text string `json:"text"`
	}{
		Text: message,
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(payload); err != nil {
		log.Println("Error encoding JSON payload:", err)
		return
	}

	request, err := http.NewRequest(http.MethodPost, u.String(), &body)
	if err != nil {
		log.Println("Error creating Slack HTTP request:", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		log.Println("Error sending Slack notification:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Invalid response from Slack channel: %s", response.Status)
	}
}

func checkDoor(pin rpio.Pin, cfg *config, slackWebhookURL string) {
	for range time.Tick(1 * time.Minute) {
		doorState := getDoorState(pin)
		log.Println("Door state:", doorState)
		if doorState == state(rpio.Low) {
			if isPeriod(cfg.PeriodStart.t, cfg.PeriodEnd.t) {
				message := fmt.Sprint("Door open at period:", time.Now())
				log.Println(message)
				go sendNotification(slackWebhookURL, message)
			}
		}
	}
}

func doorStateHandler(pin rpio.Pin) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		doorState := getDoorState(pin)
		log.Println("Door state:", doorState)

		response := struct {
			DoorState    state  `json:"door_state"`
			DoorStatText string `json:"door_state_text"`
		}{
			DoorState:    doorState,
			DoorStatText: fmt.Sprint(doorState),
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Println("Error encoding door state JSON response:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}
}

func newMux(pin rpio.Pin) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "Garage status API running...")
	})
	mux.HandleFunc("GET /getdoor", doorStateHandler(pin))

	return mux
}

func main() {
	c := flag.String("c", "config.yml", "Path to configuration file")
	flag.Parse()

	cfg, err := newConfig(*c)
	if err != nil {
		log.Fatal("Error reading configuration file:", err)
	}

	slackWebhookURL, ok := os.LookupEnv("SLACK_WEBHOOK_URL")
	if !ok {
		log.Fatal("Slack webhook URL env var is required")
	}

	pin, err := setupGPIO(cfg.SwitchPinNumber)
	if err != nil {
		log.Fatal("Error setting up GPIO:", err)
	}
	defer rpio.Close()

	go checkDoor(pin, cfg, slackWebhookURL)

	s := &http.Server{
		Addr:         ":3060",
		Handler:      newMux(pin),
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Starting server on port 3060")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

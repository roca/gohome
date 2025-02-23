package main

import (
	"flag"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	owm "github.com/briandowns/openweathermap"
	hue "github.com/collinux/gohue"
)

//go:embed rootPage.html
var rootPageHTML []byte

func lightweather(cfg *config, chRefresh <-chan struct{}) {
	externalWeatherTemp := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "external_weather_temperature",
	})

	run := func() {
		log.Println("INFO: Gettinmg current temperature")
		currentTemp, err := getCurrentTemperature(cfg)
		if err != nil {
			log.Println("ERROR:", err)
		}

		externalWeatherTemp.Set(float64(currentTemp))

		log.Println("INFO: Setting light")
		if err := setLight(cfg, currentTemp); err != nil {
			log.Println("ERROR:", err)
		}

	}

	for {
		select {
		case <-chRefresh:
			run()
		case <-time.Tick(30 * time.Minute):
			run()
		}
	}
}

func newMux(cfg *config) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		log.Println("INFO: Received request root")
		w.Write(rootPageHTML)
	})

	chRefresh := make(chan struct{}, 2)

	mux.HandleFunc("POST /refresh", func(w http.ResponseWriter, _ *http.Request) {
		log.Println("INFO: Received refresh request")

		chRefresh <- struct{}{}
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Refresh request accepted"))
	})

	mux.Handle("/metrics", promhttp.Handler())

	return mux
}

func main() {
	c := flag.String("c", "config.yml", "Path to the configuration file")
	flag.Parse()

	cfg, err := newConfig(*c)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	s := &http.Server{
		Addr:         ":3040",
		Handler:      newMux(cfg),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalln("ERROR:", err)
	}
}

func getCurrentTemperature(cfg *config) (int, error) {
	w, err := owm.NewCurrent(cfg.Unit, cfg.Lang, cfg.OWMAPIKEY)
	if err != nil {
		return 0, err
	}

	err = w.CurrentByName(cfg.Location)
	return int(math.Round(w.Main.Temp)), err
}

func setLight(cfg *config, currentTemp int) error {
	bridge, err := hue.NewBridge(cfg.HueIPAddress)
	if err != nil {
		return err
	}

	// hue-id, _ := bridge.CreateUser("create-new-user")
	if err := bridge.Login(cfg.HueID); err != nil {
		return err
	}

	weatherLight, err := bridge.GetLightByName(cfg.LightName)
	if err != nil {
		return err
	}

	if err := weatherLight.SetColor(
		pickColor(cfg, currentTemp)); err != nil {
		return err
	}

	return nil
}

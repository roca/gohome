package main

import "github.com/prometheus/client_golang/prometheus"

//go:embed rootPage.html
var rootPageHTML []byte

func lightweather(cfg *config, chRefresh <-chan struct{}) {
	externalWeatherTemp := promauto.NewGuage(prometheus.GaugeOpts{
		Name: "external_weather_temperature",
	})
}

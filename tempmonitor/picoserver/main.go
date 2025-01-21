package main

import (
	"log/slog"
	"time"

	"machine"
)

const (
	connTimeout = 3 * time.Second
	maxconns    = 3
	tcpbufsize  = 2030
	hostname    = "picotemp"
	listenPort  = "80"
)

type temp struct {
	TempC float64 `json:"tempC"`
	TempF float64 `json:"tempF"`
}

var logger *slog.Logger

func init() {
	logger = slog.New(
		slog.NewTextHandler(
			machine.Serial,
			&slog.HandlerOptions{
				Level: slog.LevelInfo,
			},
		))
}

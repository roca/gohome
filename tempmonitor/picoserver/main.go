package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"machine"
	"net/netip"
	"time"

	"log/slog"

	"github.com/soypat/cyw43439"
	"github.com/soypat/cyw43439/examples/common"

	"github.com/soypat/seqs/httpx"
	"github.com/soypat/seqs/stacks"
)

const (
	connTimeout = 3 * time.Second
	maxconns    = 3
	tcpbufsize  = 4096
	hostname    = "picotemp"
	listenPort  = 80
)

type temp struct {
	TempC float64 `json:"tempC"`
	TempF float64 `json:"tempF"`
}

var logger *slog.Logger

func init() {
	logger = slog.New(
		slog.NewTextHandler(machine.Serial, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
}

func main() {
	stack, dev := setupDevice()
	conn := newConn(stack)

	blink := make(chan uint, 3)
	go blinkLED(dev, blink)
	go handleConnection(conn, blink)

	for {
		select {
		case <-time.After(1 * time.Minute):
			logger.Info("Waiting for connections...")
		}
	}
}
func handleConnection(conn *stacks.TCPConn, blink chan uint) {
	// Reuse the same buffers for each
	// connection to avoid heap allocations.
	var resp httpx.ResponseHeader
	buf := bufio.NewReaderSize(nil, 1024)

	const newISS = 1337

	err := conn.OpenListenTCP(listenPort, newISS+100)
	if err != nil {
		logger.Error("open close:", slog.String("err", err.Error()))
		time.Sleep(3000 * time.Millisecond)
	}
	logger.Info("Port opened")

	err = conn.SetDeadline(time.Now().Add(connTimeout))
	if err != nil {
		conn.Close()
		logger.Error(
			"conn set deadline:",
			slog.String("err", err.Error()),
		)
	}

	for {

		err := conn.FlushOutputBuffer()
		if err != nil {
			logger.Error("flush output buffer:", slog.String("err", err.Error()))
			break
		}
		buf.Reset(conn)
		resp.Reset()
		HTTPHandler(conn, &resp)

		time.Sleep(3000 * time.Millisecond)

		blink <- 5
	}
}

func changeLEDState(dev *cyw43439.Device, state bool) {
	if err := dev.GPIOSet(0, state); err != nil {
		logger.Error("failed to change LED state:",
			slog.String("err", err.Error()))
	}
}

func blinkLED(dev *cyw43439.Device, blink chan uint) {
	for {
		select {
		case n := <-blink:
			lastLedState := true
			if n == 0 {
				n = 5
			}
			for i := uint(0); i < n; i++ {
				lastLedState = !lastLedState
				changeLEDState(dev, lastLedState)
				time.Sleep(500 * time.Millisecond)
			}
			// Ensure LED is on at the end
			changeLEDState(dev, true)
		}
	}
}

func setupDevice() (*stacks.PortStack, *cyw43439.Device) {
	_, stack, dev, err := common.SetupWithDHCP(common.SetupConfig{
		Hostname:    hostname,
		RequestedIP: "192.168.68.136",
		Logger:      logger,
		TCPPorts:    1,
	})
	if err != nil {
		panic("setup DHCP:" + err.Error())
	}

	// Turn LED on
	changeLEDState(dev, true)

	return stack, dev
}

func newConn(stack *stacks.PortStack) *stacks.TCPConn {
	// Start TCP server.
	listenAddr := netip.AddrPortFrom(stack.Addr(), listenPort)
	conn, err := stacks.NewTCPConn(stack, stacks.TCPConnConfig{
		TxBufSize: tcpbufsize,
		RxBufSize: tcpbufsize,
	})
	if err != nil {
		panic("TCPConn create:" + err.Error())
	}

	logger.Info("listening",
		slog.String("addr", "http://"+listenAddr.String()+":"+fmt.Sprintf("%v", listenAddr.Port())),
	)

	return conn
}

func getTemperature() *temp {
	curTemp := machine.ReadTemperature()

	return &temp{
		TempC: float64(curTemp) / 1000,
		TempF: ((float64(curTemp) / 1000) * 9 / 5) + 32,
	}
}

func HTTPHandler(respWriter io.Writer, resp *httpx.ResponseHeader) {
	resp.SetConnectionClose()
	logger.Info("Got temperature request...")
	t := getTemperature()

	logger.Info(fmt.Sprintf("%v", t))

	body, err := json.Marshal(t)
	if err != nil {
		logger.Error(
			"temperature json:",
			slog.String("err", err.Error()),
		)
		resp.SetStatusCode(500)
	} else {
		resp.SetContentType("application/json")
		resp.SetContentLength(len(body))
	}
	_, err = respWriter.Write(resp.Header())
	if err != nil {
		logger.Error("Failed to write response Header:", err)
	}
	respWriter.Write(body)
}

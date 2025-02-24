package main

import (
	"errors"
	"fmt"
	"os"
	"sort"

	hue "github.com/collinux/gohue"
	"gopkg.in/yaml.v2"
)

var errInvalidColor = errors.New("invalid color")

type color struct {
	Color     string `yaml:"color"`
	Threshold int    `yaml:"threshold"`
}

var colorTranslate = map[string]*[2]float32{
	"blue":   hue.BLUE,
	"cyan":   hue.CYAN,
	"green":  hue.GREEN,
	"orange": hue.ORANGE,
	"pink":   hue.PINK,
	"purple": hue.PURPLE,
	"red":    hue.RED,
	"white":  hue.WHITE,
	"yellow": hue.YELLOW,
}

type config struct {
	Unit         string  `yaml:"unit"`
	Lang         string  `yaml:"lang"`
	Location     string  `yaml:"location"`
	HueID        string  `yaml:"hue_id"`
	HueIPAddress string  `yaml:"hue_ip_address"`
	OWMAPIKey    string  `yaml:"owm_api_key"`
	LightName    string  `yaml:"light_name"`
	MaxColor     string  `yaml:"max_color"`
	Colors       []color `yaml:"colors"`
	ZipCode     string  `yaml:"zipcode"`
	Country     string  `yaml:"country"`
}

func (cfg *config) sortColorRange() *config {
	sort.Slice(cfg.Colors, func(i, j int) bool {
		return cfg.Colors[i].Threshold < cfg.Colors[j].Threshold
	})
	return cfg
}

func newConfig(configFile string) (*config, error) {
	cf, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer cf.Close()

	var cfg config
	if err := yaml.NewDecoder(cf).Decode(&cfg); err != nil {
		return nil, err
	}

	for _, cl := range cfg.Colors {
		if _, ok := colorTranslate[cl.Color]; !ok {
			return nil, fmt.Errorf("%w, %s", errInvalidColor, cl.Color)
		}
	}

	// AllLow user to override OWM API Key, Hue ID, Hue IP address, zipcode and country with env vars
	if owmKey, ok := os.LookupEnv("OWM_API_KEY"); ok {
		cfg.OWMAPIKey = owmKey
	}

	if hueID, ok := os.LookupEnv("HUE_ID"); ok {
		cfg.HueID = hueID
	}

	if hueIP, ok := os.LookupEnv("HUE_IP_ADDRESS"); ok {
		cfg.HueIPAddress = hueIP
	}

	if zipCode, ok := os.LookupEnv("ZIPCODE"); ok {
		cfg.ZipCode = zipCode
	}

	if country, ok := os.LookupEnv("COUNTRY"); ok {
		cfg.Country = country
	}

	return cfg.sortColorRange(), nil
}

func pickColor(cfg *config, curtTemp int) *[2]float32 {
	for _, cl := range cfg.Colors {
		if curtTemp < cl.Threshold {
			return colorTranslate[cl.Color]
		}
	}
	return colorTranslate[cfg.MaxColor]
}

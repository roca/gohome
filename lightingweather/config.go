package main

type config struct {
	Unit         string
	Lang         string
	OWMAPIKEY    string
	Location     string
	HueIPAddress string
	HueID        string
	LightName    string
}

func newConfig(configFile string) (*config, error) {
	return &config{}, nil
}

func pickColor(cfg *config, curtTemp int) *[2]float32 {
	return &[2]float32{0.0, 0.0}
}

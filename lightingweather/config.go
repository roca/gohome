package main


type config struct {
	Unit string
	Lang string
	OWMAPIKEY string
	Location string
	HueIPAddress string
	HueID string
}

func newConfig(configFile string) (*config, error) {
	return &config{},nil
}
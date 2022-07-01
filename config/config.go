package config

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Title    string  `yaml:"title"`
	TileSize float64 `yaml:"tile-size"`
	Window   struct {
		Height int `yaml:"height"`
		Width  int `yaml:"width"`
	} `yaml:"window"`
	Server                  string `yaml:"server"`
	Port                    int    `yaml:"port"`
	ScaleFactor             int    `yaml:"scale-factor"`
	DebugEnabled            bool   `yaml:"debug"`
	DebugCollidablesEnabled bool   `yaml:"debug-collidables"`
	Engine                  uint   `yaml:"engine"`
	FPSEnabled              bool   `yaml:"fps-enabled"`
}

const (
	ScreenHeight = 128
	ScreenWidth  = 128
)

//go:embed config.yml
var configRaw []byte

var config *Config

func init() {
	Reset()
}

func (c *Config) Unmarshal(raw []byte) {
	yaml.Unmarshal(raw, c)
}

func Reset() {
	c := &Config{}
	c.Unmarshal(configRaw)
	config = c
}

func Get() *Config {
	return config
}

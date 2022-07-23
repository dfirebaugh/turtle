package config

import (
	_ "embed"
	"image/color"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Title    string  `yaml:"title"`
	TileSize float64 `yaml:"tile-size"`
	Window   struct {
		Height int `yaml:"height"`
		Width  int `yaml:"width"`
	} `yaml:"window"`
	Server       string `yaml:"server"`
	Port         int    `yaml:"port"`
	ScaleFactor  int    `yaml:"scale-factor"`
	DebugEnabled bool   `yaml:"debug"`
	FPSEnabled   bool   `yaml:"fps-enabled"`
}

// some things need the screen size as a const
// we can't dynamically load this from the yaml file
const (
	ScreenHeight = 128
	ScreenWidth  = 128
)

//go:embed config.yml
var configRaw []byte

var (
	config   *Config
	Pallette = []color.Color{
		color.Black,
		color.RGBA{127, 36, 84, 255},
		color.RGBA{28, 43, 83, 255},
		color.RGBA{0, 135, 81, 255},
		color.RGBA{171, 82, 54, 255},
		color.RGBA{96, 88, 79, 255},
		color.RGBA{195, 195, 198, 255},
		color.RGBA{255, 241, 233, 255},
		color.RGBA{237, 27, 81, 255},
		color.RGBA{250, 162, 27, 255},
		color.RGBA{247, 236, 47, 255},
		color.RGBA{93, 187, 77, 255},
		color.RGBA{81, 166, 220, 255},
		color.RGBA{131, 118, 156, 255},
		color.RGBA{241, 118, 166, 255},
		color.RGBA{252, 204, 171, 255},
	}
)

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

package config

import (
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Params struct {
	ModbusAddr        string
	WebDirPath        string
	SolcastApiKey     string
	SolcastPropertyID string
	Location          *time.Location
	Palette           *PaletteParams
}

type PaletteParams struct {
	Palette      string
	AutoColor    string
	AutoColorSrc string
	ConfigFile   string
}

var Config *Params

// we use viper get the config from flags or env
func Init() {

	pflag.String("modbus.addr", "192.168.1.128:1502", "Modbus address")
	pflag.String("web.dir", "./web", "Web directory path")
	pflag.String("solcast.api_key", "", "Solcast API key")
	pflag.String("solcast.property_id", "", "Solcast property ID")
	pflag.String("palette", "", "Color palette (e.g. c41b5c-08415c-6b818c-f1bf98)")
	pflag.String("auto-color", "", "Auto-fetch palette interval (e.g. 5m, 1h)")
	pflag.String("auto-color-source", "colormind", "Auto-color source: colormind, local, coolors, or hybrid")
	pflag.String("config", "config.json", "Config file path")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
	viper.AutomaticEnv()

	berlin, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		panic(err)
	}

	Config = &Params{
		ModbusAddr:        viper.GetString("modbus.addr"),
		WebDirPath:        viper.GetString("web.dir"),
		SolcastApiKey:     viper.GetString("solcast.api_key"),
		SolcastPropertyID: viper.GetString("solcast.property_id"),
		Location:          berlin,
		Palette: &PaletteParams{
			Palette:      viper.GetString("palette"),
			AutoColor:    viper.GetString("auto-color"),
			AutoColorSrc: viper.GetString("auto-color-source"),
			ConfigFile:   viper.GetString("config"),
		},
	}
}

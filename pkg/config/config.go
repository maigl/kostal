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
	Location *time.Location
}

var Config *Params

// we use viper get the config from flags or env
func Init() {

	pflag.String("modbus.addr", "192.168.1.128:1502", "Modbus address")
	pflag.String("web.dir", "./web", "Web directory path")
	pflag.String("solcast.api_key", "", "Solcast API key")
	pflag.String("solcast.property_id", "", "Solcast property ID")
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
	}
}

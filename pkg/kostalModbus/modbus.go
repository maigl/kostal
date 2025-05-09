package kostalModbus

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/goburrow/modbus"
	"github.com/maigl/kostal/pkg/config"
	"github.com/maigl/kostal/pkg/data"
	"github.com/maigl/kostal/pkg/solcast"
)

var client modbus.Client

type PowerItem struct {
	Label    string
	Unit     string
	Value    string
	Forecast string
}

func getModbusHandler() *modbus.TCPClientHandler {
	// Modbus TCP
	handler := modbus.NewTCPClientHandler(config.Config.ModbusAddr)
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 71
	// handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session
	err := handler.Connect()
	if err != nil {
		log.Fatal(err)
	}
	return handler
}

func GetPower() map[string]PowerItem {
	if client == nil {
		// TODO maybe cache power object and reduce number of modbus calls
		modbusHandler := getModbusHandler()
		client = modbus.NewClient(modbusHandler)
		// defer modbusHandler.Close()
	}

	br := data.Registers["514"]
	err := br.Read(client)
	if err != nil {
		log.Fatal(err)
	}

	battery := fmt.Sprintf("%d", br.Value)

	yr := data.Registers["260"]
	err = yr.Read(client)
	if err != nil {
		log.Fatal(err)
	}

	yield := yr.Float32()
	yr = data.Registers["270"]
	err = yr.Read(client)
	if err != nil {
		log.Fatal(err)
	}

	yield += yr.Float32()
	yield /= float32(1000)
	if yield < 0 {
		yield = 0
	}
	yieldString := fmt.Sprintf("%1.1f", yield)

	cr := data.Registers["106"]
	err = cr.Read(client)
	if err != nil {
		log.Fatal(err)
	}

	consumption := cr.Float32()
	cr = data.Registers["108"]
	err = cr.Read(client)
	if err != nil {
		log.Fatal(err)
	}

	consumption += cr.Float32()
	cr = data.Registers["116"]
	err = cr.Read(client)
	if err != nil {
		log.Fatal(err)
	}

	consumption += cr.Float32()
	consumption /= float32(1000)
	if consumption < 0 {
		consumption = 0
	}
	consumptionString := fmt.Sprintf("%1.1f", consumption)

	ir := data.Registers["575"]
	err = ir.Read(client)
	if err != nil {
		log.Fatal(err)
	}
	gridRaw := float32(ir.Uint16()) / 1000.
	// we seem to have an overflow here
	if gridRaw > 60 {
		gridRaw = 0
	}
	grid := gridRaw - consumption
	gridLabel := "to grid"
	if grid <= 0 {
		gridLabel = "from grid"
	}
	gridString := fmt.Sprintf("%1.1f", math.Abs(float64(grid)))

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, config.Config.Location)

	forecast, err := solcast.GetForcast(today)
	if err != nil {
		log.Printf("error getting forecast: %v", err)
	}

	return map[string]PowerItem{
		"battery":     {Label: "battery", Unit: "%", Value: battery, Forecast: forecast.Morning.Icon()},
		"consumption": {Label: "consumption", Unit: "kW", Value: consumptionString, Forecast: forecast.Noon.Icon()},
		"grid":        {Label: gridLabel, Unit: "kW", Value: gridString, Forecast: forecast.Afternoon.Icon()},
		"yield":       {Label: "yield", Unit: "kW", Value: yieldString, Forecast: forecast.Evening.Icon()},
	}
}

func PrintAllRegisters() { // nolint
	for _, r := range data.Registers {
		err := r.Read(client)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("r: %+v", r)
	}
}

package main

import (
	"fmt"
	"log"
	"maigl/kostal/data"
	"math"
	"net/http"
	"text/template"
	"time"

	"github.com/goburrow/modbus"
)

type PowerItem struct {
	Label string
	Unit  string
	Value string
}

func web(w http.ResponseWriter, r *http.Request) {

	power := getPower()

	//fp := path.Join("web", "index.html")
	//tmpl, err := template.ParseFiles(fp)

	tmpl, err := template.New("web").Parse(html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, power); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var client modbus.Client

func main() {

	http.HandleFunc("/", web)
	http.ListenAndServe(":8080", nil)

}

func getModbusHandler() *modbus.TCPClientHandler {
	addr := "192.168.0.38:1502"
	// Modbus TCP
	handler := modbus.NewTCPClientHandler(addr)
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 71
	//handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session
	err := handler.Connect()
	if err != nil {
		log.Fatal(err)
	}
	return handler
}

func getPower() map[string]PowerItem {

	if client == nil {
		//TODO maybe cache power object and reduce number of modbus calls
		modbusHandler := getModbusHandler()
		client = modbus.NewClient(modbusHandler)
		//defer modbusHandler.Close()
	}

	br := data.Registers["514"]
	br.Read(client)
	battery := fmt.Sprintf("%d", br.Value)

	yr := data.Registers["260"]
	yr.Read(client)
	yield := yr.Float32()
	yr = data.Registers["270"]
	yr.Read(client)
	yield += yr.Float32()
	yield /= float32(1000)
	if yield < 0 {
		yield = 0
	}
	yieldString := fmt.Sprintf("%1.1f", yield)

	cr := data.Registers["106"]
	cr.Read(client)
	consumption := cr.Float32()
	cr = data.Registers["108"]
	cr.Read(client)
	consumption += cr.Float32()
	cr = data.Registers["116"]
	cr.Read(client)
	consumption += cr.Float32()
	consumption /= float32(1000)
	if consumption < 0 {
		consumption = 0
	}
	consumptionString := fmt.Sprintf("%1.1f", consumption)

	ir := data.Registers["575"]
	ir.Read(client)
	grid := float32(ir.Uint16())/1000. - consumption
	gridLabel := "to grid"
	if grid <= 0 {
		gridLabel = "from grid"
	}
	gridString := fmt.Sprintf("%1.1f", math.Abs(float64(grid)))

	return map[string]PowerItem{
		"battery":     PowerItem{Label: "battery", Unit: "%", Value: battery},
		"yield":       PowerItem{Label: "yield", Unit: "kW", Value: yieldString},
		"consumption": PowerItem{Label: "consumption", Unit: "kW", Value: consumptionString},
		"grid":        PowerItem{Label: gridLabel, Unit: "kW", Value: gridString},
	}

}

func printAllRegisters() {
	for _, r := range data.Registers {
		r.Read(client)
		log.Printf("r: %+v", r)
	}
}

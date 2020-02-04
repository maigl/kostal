package main

import (
	"log"
	"maigl/kostal/data"
	"os"
	"time"

	"github.com/goburrow/modbus"
)

func main() {

	addr := "192.168.0.38:1502"

	// Modbus TCP
	handler := modbus.NewTCPClientHandler(addr)
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 71
	handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session
	err := handler.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer handler.Close()

	client := modbus.NewClient(handler)

	r := data.Register{Addr: 100, Description: "total power", Unit: "W", Length: 2, Format: "Float"}
	r.Read(client)
	log.Printf("r: %+v", r)
	r = data.Register{Addr: 512, Description: "battery gross capacity", Unit: "Ah", Length: 2, Format: "U32"}
	r.Read(client)
	log.Printf("r: %+v", r)
	r = data.Register{Addr: 531, Description: "inverter max power", Unit: "W", Length: 1, Format: "U16"}
	r.Read(client)
	log.Printf("r: %+v", r)
	r = data.Register{Addr: 112, Description: "home total consumption grid", Unit: "W", Length: 2, Format: "Float"}
	r.Read(client)
	log.Printf("r: %+v", r)

}

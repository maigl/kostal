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

	for _, r := range data.Registers {
		r.Read(client)
		log.Printf("r: %+v", r)
	}

}

package data

import (
	"encoding/binary"
	"math"

	"github.com/goburrow/modbus"
)

type Register struct {
	Addr        uint16
	Description string
	Unit        string
	Format      string
	Length      uint16
	Access      string
	Data        []byte
	Value       interface{}
}

func (r *Register) Read(client modbus.Client) error {
	v, err := client.ReadHoldingRegisters(r.Addr, r.Length)
	if err != nil {
		return err
	}
	r.Data = v
	r.Value = r.Get()
	return nil
}

func (r *Register) Uint32() (res uint32) {
	return binary.BigEndian.Uint32(r.Data)
}

func (r *Register) Uint16() (res uint16) {
	return binary.BigEndian.Uint16(r.Data)
}

func (r *Register) Float32() (res float32) {
	//so byte order is BigEndias .. but word order is Little Endian ?!
	tmp := []byte{r.Data[2], r.Data[3], r.Data[0], r.Data[1]}
	return math.Float32frombits(binary.BigEndian.Uint32(tmp))
}

func (r *Register) Get() interface{} {
	switch r.Format {
	case "Float":
		return r.Float32()
	case "U32":
		return r.Uint32()
	case "U16":
		return r.Uint16()
	default:
		return nil
	}
	return nil
}

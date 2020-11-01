package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"reflect"
)

func ConvLittleEndianNumericToSlice(value interface{}) []byte {
	var b []byte

	switch v := value.(type) {
	case uint8:
		b = []byte{byte(v)}
	case int8:
		b = []byte{byte(v)}
	case uint16:
		b = make([]byte, 2)
		binary.LittleEndian.PutUint16(b, v)
	case int16:
		b = make([]byte, 2)
		binary.LittleEndian.PutUint16(b, uint16(v))
	case uint32:
		b = make([]byte, 4)
		binary.LittleEndian.PutUint32(b, v)
	case int32:
		b = make([]byte, 4)
		binary.LittleEndian.PutUint32(b, uint32(v))
	case uint64:
		b = make([]byte, 8)
		binary.LittleEndian.PutUint64(b, v)
	case int64:
		b = make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(v))
	case float32:
		b = make([]byte, 4)
		binary.LittleEndian.PutUint32(b, math.Float32bits(v))
	case float64:
		b = make([]byte, 8)
		binary.LittleEndian.PutUint64(b, math.Float64bits(v))
	case string:
		b = []byte(v)

	}

	return b
}

type Ardrone3PilotingPCMDArguments struct {
	Flag               uint8
	Roll               int8
	Pitch              int8
	Yaw                int8
	Gaz                int8
	TimestampAndSeqNum uint32
}

func (a Ardrone3PilotingPCMDArguments) Encode() []byte {
	//TODO: .............

	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(a)
	log.Printf("valueOf: %#v\n", valueOf)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	log.Printf("valueOf.NumField(): %#v\n", valueOf.NumField())
	for i := 0; i < valueOf.NumField(); i++ {
		log.Printf("valueOf.Field(i): %T\n", valueOf.Field(i).Interface())

		b := ConvLittleEndianNumericToSlice(valueOf.Field(i).Interface())
		fmt.Printf("mySlice = %#v\n", b)

		log.Printf("b: %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

func main() {
	a := Ardrone3PilotingPCMDArguments{
		Gaz: 7,
	}

	b := a.Encode()

	fmt.Printf("%#v\n", b)
}

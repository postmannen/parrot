package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	size := uint32(7)
	var buf bytes.Buffer

	err := binary.Write(&buf, binary.LittleEndian, size)
	if err != nil {
		fmt.Printf("error: binary write failed: %v\n", err)
	}

	fmt.Printf("%v\n", buf.Bytes())
}

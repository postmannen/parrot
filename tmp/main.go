package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//2 127 75 23 0 0 0 - 1 4 6 0 137 146 75 60 159 186 239 58 74 247 29 192
//1, 4, 6, 0, 137, 146, 75, 60, 159, 186, 239, 58, 74, 247, 29, 192

func main() {
	data := []byte{1, 4, 6, 0, 137, 146, 75, 60, 159, 186, 239, 58, 74, 247, 29, 192}
	var d uint16

	r := bytes.NewReader(data[2:4])

	err := binary.Read(r, binary.LittleEndian, &d)
	if err != nil {
		fmt.Println("error: binary.Read failed: ", err)
	}

	myVar := int(d)

	fmt.Printf("%v, %b\n", myVar, myVar)
}

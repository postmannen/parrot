package main

import (
	"fmt"
	"reflect"
)

type data struct {
	a uint8
	b uint8
	c uint16
}

func main() {
	d := data{5, 6, 7}

	rv := reflect.ValueOf(d)

	for i := 0; i < rv.NumField(); i++ {
		v := rv.Field(i).Interface()
		fmt.Printf("%b, %T\n", v, v)
	}

}

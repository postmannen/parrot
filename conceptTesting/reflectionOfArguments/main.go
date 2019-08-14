package main

import (
	"fmt"
	"reflect"
)

type post struct {
	ArgumentHeight int     `type:"int"`
	ArgumenAngle   float64 `type:"float"`
}

func checkSwitch(d interface{}) {
	value := reflect.ValueOf(d)
	if value.Kind() != reflect.Ptr {
		panic("not a pointer")
	}

	elements := value.Elem()

	//this loops through the fields
	for i := 0; i < elements.NumField(); i++ { // iterates through every struct type field
		tag := elements.Type().Field(i).Tag // returns the tag string
		field := elements.Field(i)          // returns the content of the struct type field
		switch tag.Get("type") {
		case "int":
			field.SetInt(20)
		case "float":
			field.SetFloat(20.5)

		}
	}
}

func main() {
	p := &post{ArgumentHeight: 10, ArgumenAngle: 10.5}
	fmt.Println(p)
	checkSwitch(p)
	fmt.Println(p)
}

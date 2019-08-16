package main

import (
	"fmt"
)

type animalChopper func(string) string

func chopAnimal(s string) string {
	r := fmt.Sprintf("The %v have been chop'ed !\n", s)
	return r
}

func main() {
	chopMap := make(map[string]func(string) string)
	chopMap["monkey"] = func(s string) string { return s }

	s := chopMap["monkey"]

	fmt.Printf("%v, type = %T\n", s("monkey"), s)
	//fmt.Printf("%v, type = %T\n", v, v)
}

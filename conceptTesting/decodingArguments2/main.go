package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"reflect"
)

/*
u8 1 unsigned 8bit value
i8 1 signed 8bit value
u16 2 unsigned 16bit value
i16 2 signed 16bit value
u32 4 unsigned 32bit value
i32 4 signed 32bit value
u64 8 unsigned 64bit value
i64 8 signed 64bit value
float 4 IEEE-754 single precision
double 8 IEEE-754 double precision
string * Null terminated string (C-String)
(Variable size)
enum 4 Per command defined enum
*/

// ------------------------------------------------------------------------------------

// uint8Type makes a type for uint8 data.
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type uint8Type struct {
	value  uint8
	length int
}

var u8 = uint8Type{
	length: 1,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *uint8Type) getLength() int {
	return f.length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *uint8Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the uint8Type.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val uint8

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	f.value = val

	fmt.Printf("Content of f = %#v\n", *f)

	return f.value, nil
}

// ------------------------------------------------------------------------------------

// int8Type makes a type for int8 data.
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type int8Type struct {
	value  int8
	length int
}

var i8 = int8Type{
	length: 1,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *int8Type) getLength() int {
	return f.length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *int8Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the int8Type.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val int8

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	f.value = val

	fmt.Printf("Content of f = %#v\n", *f)

	return f.value, nil
}

// ------------------------------------------------------------------------------------

// uint16Type makes a type for int8 data.
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type uint16Type struct {
	value  uint16
	length int
}

var u16 = uint16Type{
	length: 2,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *uint16Type) getLength() int {
	return f.length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *uint16Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the uint16Type.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val uint16

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	f.value = val

	fmt.Printf("Content of f = %#v\n", *f)

	return f.value, nil
}

// ------------------------------------------------------------------------------------

// int16Type makes a type for int8 data.
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type int16Type struct {
	value  int16
	length int
}

var i16 = uint16Type{
	length: 2,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *int16Type) getLength() int {
	return f.length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *int16Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the int16Type.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val int16

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	f.value = val

	fmt.Printf("Content of f = %#v\n", *f)

	return f.value, nil
}

// ------------------------------------------------------------------------------------

// uint32Type
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type uint32Type struct {
	value  uint32
	length int
}

var u32 = uint32Type{
	length: 4,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *uint32Type) getLength() int {
	return f.length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *uint32Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the uint32Type.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val uint32

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	f.value = val

	fmt.Printf("Content of f = %#v\n", *f)

	return f.value, nil
}

// ------------------------------------------------------------------------------------

// int32Type
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type int32Type struct {
	value  int32
	length int
}

var i32 = uint32Type{
	length: 4,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *int32Type) getLength() int {
	return f.length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *int32Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the int32Type.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val int32

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	f.value = val

	fmt.Printf("Content of f = %#v\n", *f)

	return f.value, nil
}

// ------------------------------------------------------------------------------------

// uint64Type
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type uint64Type struct {
	value  uint64
	length int
}

var u64 = uint64Type{
	length: 8,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *uint64Type) getLength() int {
	return f.length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *uint64Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the uint64Type.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val uint64

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	f.value = val

	fmt.Printf("Content of f = %#v\n", *f)

	return f.value, nil
}

// ------------------------------------------------------------------------------------

// int64Type
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type int64Type struct {
	value  int64
	length int
}

var i64 = uint64Type{
	length: 8,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *int64Type) getLength() int {
	return f.length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *int64Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the int64Type.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val int64

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	f.value = val

	fmt.Printf("Content of f = %#v\n", *f)

	return f.value, nil
}

// ------------------------------------------------------------------------------------

// float32Type makes a type for int8 data.
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type float32Type struct {
	value  float32
	length int
}

// float makes a type for float32 data, and tells the length of bytes it
// is made of.
var float = float32Type{
	length: 4,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *float32Type) getLength() int {
	return f.length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *float32Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the float32.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val float32

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	f.value = val

	fmt.Printf("Content of f = %#v\n", *f)

	return f.value, nil
}

// ------------------------------------------------------------------------------------

// float64Type
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type float64Type struct {
	value  float64
	length int
}

// double makes a type for float64 data, and tells the length of bytes it
// is made of.
var double = float64Type{
	length: 8,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *float64Type) getLength() int {
	return f.length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *float64Type) argDecode(b []byte) (value interface{}, err error) {
	fmt.Printf("running the float64.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val float64

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return nil, err
	}

	f.value = val

	fmt.Printf("Content of f = %#v\n", *f)

	return f.value, nil
}

// ------------------------------------------------------------------------------------

// ------------------------------------------------------------------------------------

// checkSwitch takes the struct to fill as a pointer value, and the arguments as a
// slice of []interface{} as input.
func insertArgValueIntoStruct(argStruct interface{}, argValues []interface{}) {
	dataValue := reflect.ValueOf(argStruct)
	if dataValue.Kind() != reflect.Ptr {
		panic("not a pointer")
	}

	dataElements := dataValue.Elem()

	//this loops through the fields
	for i := 0; i < dataElements.NumField(); i++ { // iterates through every struct type field
		//k := elements.Kind()
		dataType := dataElements.Type().Field(i).Type // returns the tag string
		dataField := dataElements.Field(i)            // returns the content of the struct type field

		argVal := reflect.ValueOf(argValues[i])
		fmt.Printf("argVal = %+v, type = %T\n", argVal, argVal)

		// Check what the types it is, and when the correct type for the field is
		// found, insert the value into the struct field.
		switch dataType.String() {
		case "int":
			fmt.Printf("Reflecting INT\n")
			v := argVal.Int()
			//fmt.Printf("v = %+v, and type = %T\n", v(), v)
			dataField.SetInt(v)
		case "int8":
			fmt.Printf("Reflecting INT8\n")
			v := argVal.Int()
			//fmt.Printf("v = %+v, and type = %T\n", v(), v)
			dataField.SetInt(v)
		case "float64":
			fmt.Printf("Reflecting float64\n")
			v := argVal.Float()
			dataField.SetFloat(v)
		case "float32":
			fmt.Printf("Reflecting float32\n")
			v := argVal.Float()
			dataField.SetFloat(v)
		}
	}
}

// ------------------------------------------------------------------------------------
// argDecoder is an interface type which tells that any type that
// have the methods argDecode([]byte) error, and getLength() int
// is of the interface type argDecoder.
type argDecoder interface {
	argDecode([]byte) (interface{}, error)
	getLength() int
}

// argumentState is a type for keeping track of the start position of the
// data to read in a slice.
type argumentsState struct {
	position int
}

// argumentsToDecode takes a []byte and any number of the interface type argDecoder
// is input.
// It will loop through the argDecoder methods, and run the concrete types method,
// one by one until all methods are done.
// The method will use the getLength() method to know the size of the portion of data
// to work with, and increase the the position with the size of the last data read to
// know where the next piece of data starts.
// TODO: Make logic check if there are given the correct amount of argDecoders to
// handle the length of the data slice given as input, and return error if they don't
// match.
func (as *argumentsState) argumentsToDecode(argStruct interface{}, d []byte, a ...argDecoder) ([]interface{}, error) {
	as.position = 0
	argumentSlice := []interface{}{}
	for _, v := range a {
		fmt.Println("------------------Decoding byte or bytes----------------------")
		val, err := v.argDecode(d[as.position : as.position+v.getLength()])
		if err != nil {
			return nil, err
		}

		// testing putting the valius into a slice, to iterate later.
		argumentSlice = append(argumentSlice, val)
		fmt.Printf("val = %+v, type = %T\n", val, val)

		l := v.getLength()
		as.position += l
	}

	return argumentSlice, nil
}

func main() {
	//The data below should decode
	//bytes 1->4 = a float32,
	//byte 5 = an int8
	tmpData := []byte{154, 221, 45, 61, 83}

	droneArguments := &struct {
		SomeFloatValue float32
		SomeIntValue   int8
	}{}

	a := argumentsState{}
	argSlice, err := a.argumentsToDecode(droneArguments, tmpData, &float, &i8)
	if err != nil {
		fmt.Println("error: argumentsToDecode: failed looping over v ", err)
	}

	insertArgValueIntoStruct(droneArguments, argSlice)

	fmt.Printf("%+v\n", droneArguments)
}
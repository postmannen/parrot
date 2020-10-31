package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"reflect"
)

type ProjectDef uint8
type ClassDef uint8
type CmdDef uint16

type Command struct {
	Project ProjectDef
	Class   ClassDef
	Cmd     CmdDef
}

// All ARDrone3-only commands
const ProjectArdrone3 ProjectDef = 1

// All commands related to piloting the drone
const ClassPiloting ClassDef = 0

// title : Take off,
// desc : Ask the drone to take off.\n On the fixed wings (such as Disco): not used except to cancel a land.,
// support : 0901;090c;090e,
// result : On the quadcopters: the drone takes off if its [FlyingState](#1-4-1) was landed.\n On the fixed wings, the landing process is aborted if the [FlyingState](#1-4-1) was landing.\n Then, event [FlyingState](#1-4-1) is triggered.,
const CmdTakeOff CmdDef = 1

type Ardrone3PilotingTakeOff Command

type Ardrone3PilotingTakeOffArguments struct {
}

func (a Ardrone3PilotingTakeOff) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingTakeOffArguments{}
	// No arguments to decode here !!

	return arg
}
func (a Ardrone3PilotingTakeOff) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingTakeOff = Ardrone3PilotingTakeOff{
	Project: ProjectArdrone3,
	Class:   ClassPiloting,
	Cmd:     CmdTakeOff,
}

// title : Move the drone,
// desc : Move the drone.\n The libARController is sending the command each 50ms.\n\n **Please note that you should call setPilotingPCMD and not sendPilotingPCMD because the libARController is handling the periodicity and the buffer on which it is sent.**,
// support : 0901;090c;090e,
// result : The drone moves! Yaaaaay!\n Event [SpeedChanged](#1-4-5), [AttitudeChanged](#1-4-6) and [PositionChanged](#1-4-4) (only if gps of the drone has fixed) are triggered.,
const CmdPCMD CmdDef = 2

type Ardrone3PilotingPCMD Command

type Ardrone3PilotingPCMDArguments struct {
	Flag               uint8
	Roll               int8
	Pitch              int8
	Yaw                int8
	Gaz                int8
	TimestampAndSeqNum uint32
}

func (a Ardrone3PilotingPCMD) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingPCMDArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Flag)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Roll)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Pitch)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Yaw)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Gaz)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TimestampAndSeqNum)
	offset += 4

	return arg
}
func (a Ardrone3PilotingPCMD) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingPCMD = Ardrone3PilotingPCMD{
	Project: ProjectArdrone3,
	Class:   ClassPiloting,
	Cmd:     CmdPCMD,
}

// title : Land,
// desc : Land.\n Please note that on copters, if you put some positive gaz (in the [PilotingCommand](#1-0-2)) during the landing, it will cancel it.,
// support : 0901;090c;090e,
// result : On the copters, the drone lands if its [FlyingState](#1-4-1) was taking off, hovering or flying.\n On the fixed wings, the drone lands if its [FlyingState](#1-4-1) was hovering or flying.\n Then, event [FlyingState](#1-4-1) is triggered.,
const CmdLanding CmdDef = 3

type Ardrone3PilotingLanding Command

type Ardrone3PilotingLandingArguments struct {
}

func (a Ardrone3PilotingLanding) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingLandingArguments{}
	// No arguments to decode here !!

	return arg
}
func (a Ardrone3PilotingLanding) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingLanding = Ardrone3PilotingLanding{
	Project: ProjectArdrone3,
	Class:   ClassPiloting,
	Cmd:     CmdLanding,
}

// title : Cut out the motors,
// desc : Cut out the motors.\n This cuts immediatly the motors. The drone will fall.\n This command is sent on a dedicated high priority buffer which will infinitely retry to send it if the command is not delivered.,
// support : 0901;090c;090e,
// result : The drone immediatly cuts off its motors.\n Then, event [FlyingState](#1-4-1) is triggered.,
const CmdEmergency CmdDef = 4

type Ardrone3PilotingEmergency Command

type Ardrone3PilotingEmergencyArguments struct {
}

func (a Ardrone3PilotingEmergency) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingEmergencyArguments{}
	// No arguments to decode here !!

	return arg
}
func (a Ardrone3PilotingEmergency) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingEmergency = Ardrone3PilotingEmergency{
	Project: ProjectArdrone3,
	Class:   ClassPiloting,
	Cmd:     CmdEmergency,
}

// title : Return home,
// desc : Return home.\n Ask the drone to fly to its [HomePosition](#1-24-0).\n The availability of the return home can be get from [ReturnHomeState](#1-4-3).\n Please note that the drone will wait to be hovering to start its return home. This means that it will wait to have a [flag](#1-0-2) set at 0.,
// support : 0901;090c;090e,
// result : The drone will fly back to its home position.\n Then, event [ReturnHomeState](#1-4-3) is triggered.\n You can get a state pending if the drone is not ready to start its return home process but will do it as soon as it is possible.,
const CmdNavigateHome CmdDef = 5

type Ardrone3PilotingNavigateHome Command

type Ardrone3PilotingNavigateHomeArguments struct {
	Start uint8
}

func (a Ardrone3PilotingNavigateHome) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingNavigateHomeArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Start)
	offset++

	return arg
}
func (a Ardrone3PilotingNavigateHome) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingNavigateHome = Ardrone3PilotingNavigateHome{
	Project: ProjectArdrone3,
	Class:   ClassPiloting,
	Cmd:     CmdNavigateHome,
}

// title : Auto take off mode,
// desc : Auto take off mode.,
const CmdAutoTakeOffMode CmdDef = 6

type Ardrone3PilotingAutoTakeOffMode Command

type Ardrone3PilotingAutoTakeOffModeArguments struct {
	State uint8
}

func (a Ardrone3PilotingAutoTakeOffMode) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingAutoTakeOffModeArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.State)
	offset++

	return arg
}
func (a Ardrone3PilotingAutoTakeOffMode) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingAutoTakeOffMode = Ardrone3PilotingAutoTakeOffMode{
	Project: ProjectArdrone3,
	Class:   ClassPiloting,
	Cmd:     CmdAutoTakeOffMode,
}

// title : Move the drone to a relative position,
// desc : Move the drone to a relative position and rotate heading by a given angle.\n Moves are relative to the current drone orientation, (drone's reference).\n Also note that the given rotation will not modify the move (i.e. moves are always rectilinear).,
// support : 0901:3.3.0;090c:3.3.0,
// result : The drone will move of the given offsets.\n Then, event [RelativeMoveEnded](#1-34-0) is triggered.\n If you send a second relative move command, the drone will trigger a [RelativeMoveEnded](#1-34-0) with the offsets it managed to do before this new command and the value of error set to interrupted.,
const CmdMoveBy CmdDef = 7

type Ardrone3PilotingmoveBy Command

type Ardrone3PilotingmoveByArguments struct {
	DX   float32
	DY   float32
	DZ   float32
	DPsi float32
}

func (a Ardrone3PilotingmoveBy) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingmoveByArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.DX)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.DY)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.DZ)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.DPsi)
	offset += 4

	return arg
}
func (a Ardrone3PilotingmoveBy) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingmoveBy = Ardrone3PilotingmoveBy{
	Project: ProjectArdrone3,
	Class:   ClassPiloting,
	Cmd:     CmdMoveBy,
}

// title : Prepare the drone to take off,
// desc : Prepare the drone to take off.\n On copters: initiates the thrown takeoff. Note that the drone will do the thrown take off even if it is steady.\n On fixed wings: initiates the take off process on the fixed wings.\n\n Setting the state to 0 will cancel the preparation. You can cancel it before that the drone takes off.,
// support : 090e;090c:4.3.0,
// result : The drone will arm its motors if not already armed.\n Then, event [FlyingState](#1-4-1) is triggered with state set at motor ramping.\n Then, event [FlyingState](#1-4-1) is triggered with state set at userTakeOff.\n Then user can throw the drone to make it take off.,
const CmdUserTakeOff CmdDef = 8

type Ardrone3PilotingUserTakeOff Command

type Ardrone3PilotingUserTakeOffArguments struct {
	State uint8
}

func (a Ardrone3PilotingUserTakeOff) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingUserTakeOffArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.State)
	offset++

	return arg
}
func (a Ardrone3PilotingUserTakeOff) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingUserTakeOff = Ardrone3PilotingUserTakeOff{
	Project: ProjectArdrone3,
	Class:   ClassPiloting,
	Cmd:     CmdUserTakeOff,
}

// title : Circle,
// desc : Make the fixed wing circle.\n The circle will use the [CirclingAltitude](#1-6-14) and the [CirclingRadius](#1-6-13),
// support : 090e,
// result : The fixed wing will circle in the given direction.\n Then, event [FlyingState](#1-4-1) is triggered with state set at hovering.,
const CmdCircle CmdDef = 9

type Ardrone3PilotingCircle Command

type Ardrone3PilotingCircleArguments struct {
	Direction uint32
}

func (a Ardrone3PilotingCircle) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingCircleArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Direction)
	offset += 4

	return arg
}
func (a Ardrone3PilotingCircle) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingCircle = Ardrone3PilotingCircle{
	Project: ProjectArdrone3,
	Class:   ClassPiloting,
	Cmd:     CmdCircle,
}

// title : Move to a location,
// desc : Move the drone to a specified location.\n If a new command moveTo is sent, the drone will immediatly run it (no cancel will be issued).\n If a [CancelMoveTo](#1-0-11) command is sent, the moveTo is stopped.\n During the moveTo, all pitch, roll and gaz values of the piloting command will be ignored by the drone.\n However, the yaw value can be used.,
// support : 090c:4.3.0,
// result : Event [MovingTo](#1-4-12) is triggered with state running. Then, the drone will move to the given location.\n Then, event [MoveToChanged](#1-4-12) is triggered with state succeed.,
const CmdMoveTo CmdDef = 10

type Ardrone3PilotingmoveTo Command

type Ardrone3PilotingmoveToArguments struct {
	Latitude        float64
	Longitude       float64
	Altitude        float64
	Orientationmode uint32
	Heading         float32
}

func (a Ardrone3PilotingmoveTo) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingmoveToArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Latitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Longitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Altitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Orientationmode)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Heading)
	offset += 4

	return arg
}
func (a Ardrone3PilotingmoveTo) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingmoveTo = Ardrone3PilotingmoveTo{
	Project: ProjectArdrone3,
	Class:   ClassPiloting,
	Cmd:     CmdMoveTo,
}

// title : Cancel the moveTo,
// desc : Cancel the current moveTo.\n If there is no current moveTo, this command has no effect.,
// support : 090c:4.3.0,
// result : Event [MoveToChanged](#1-4-12) is triggered with state canceled.,
const CmdCancelMoveTo CmdDef = 11

type Ardrone3PilotingCancelMoveTo Command

type Ardrone3PilotingCancelMoveToArguments struct {
}

func (a Ardrone3PilotingCancelMoveTo) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingCancelMoveToArguments{}
	// No arguments to decode here !!

	return arg
}
func (a Ardrone3PilotingCancelMoveTo) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingCancelMoveTo = Ardrone3PilotingCancelMoveTo{
	Project: ProjectArdrone3,
	Class:   ClassPiloting,
	Cmd:     CmdCancelMoveTo,
}

// title : Start a piloted POI,
// desc : Start a piloted Point Of Interest.\n During a piloted POI, the drone will always look at the given POI but can be piloted normally. However, yaw value is ignored. Camera tilt and pan command is also ignored.\n Ignored if [PilotedPOI](#1-4-14) state is UNAVAILABLE.,
// support : 090c:4.3.0,
// result : If the drone is hovering, event [PilotedPOI](#1-4-14) is triggered with state RUNNING. If the drone is not hovering, event [PilotedPOI](#1-4-14) is triggered with state PENDING, waiting to hover. When the drone hovers, the state will change to RUNNING. If the drone does not hover for a given time, piloted POI is canceled by the drone and state will change to AVAILABLE. Then, the drone will look at the given location.,
const CmdStartPilotedPOI CmdDef = 12

type Ardrone3PilotingStartPilotedPOI Command

type Ardrone3PilotingStartPilotedPOIArguments struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

func (a Ardrone3PilotingStartPilotedPOI) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStartPilotedPOIArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Latitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Longitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Altitude)
	offset += 8

	return arg
}
func (a Ardrone3PilotingStartPilotedPOI) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStartPilotedPOI = Ardrone3PilotingStartPilotedPOI{
	Project: ProjectArdrone3,
	Class:   ClassPiloting,
	Cmd:     CmdStartPilotedPOI,
}

// title : Stop the piloted POI,
// desc : Stop the piloted Point Of Interest.\n If [PilotedPOI](#1-4-14) state is RUNNING or PENDING, stop it.,
// support : 090c:4.3.0,
// result : Event [PilotedPOI](#1-4-14) is triggered with state AVAILABLE.,
const CmdStopPilotedPOI CmdDef = 13

type Ardrone3PilotingStopPilotedPOI Command

type Ardrone3PilotingStopPilotedPOIArguments struct {
}

func (a Ardrone3PilotingStopPilotedPOI) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStopPilotedPOIArguments{}
	// No arguments to decode here !!

	return arg
}
func (a Ardrone3PilotingStopPilotedPOI) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStopPilotedPOI = Ardrone3PilotingStopPilotedPOI{
	Project: ProjectArdrone3,
	Class:   ClassPiloting,
	Cmd:     CmdStopPilotedPOI,
}

// title : Cancel the relative move,
// desc : Cancel the current relative move.\n If there is no current relative move, this command has no effect.,
// result : Event [RelativeMoveChanged](#1-4-16) is triggered with state canceled.,
const CmdCancelMoveBy CmdDef = 14

type Ardrone3PilotingCancelMoveBy Command

type Ardrone3PilotingCancelMoveByArguments struct {
}

func (a Ardrone3PilotingCancelMoveBy) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingCancelMoveByArguments{}
	// No arguments to decode here !!

	return arg
}
func (a Ardrone3PilotingCancelMoveBy) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingCancelMoveBy = Ardrone3PilotingCancelMoveBy{
	Project: ProjectArdrone3,
	Class:   ClassPiloting,
	Cmd:     CmdCancelMoveBy,
}

// Animation commands
const ClassAnimations ClassDef = 5

// title : Make a flip,
// desc : Make a flip.,
// support : 0901;090c,
// result : The drone will make a flip if it has enough battery.,
const CmdFlip CmdDef = 0

type Ardrone3AnimationsFlip Command

type Ardrone3AnimationsFlipArguments struct {
	Direction uint32
}

func (a Ardrone3AnimationsFlip) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3AnimationsFlipArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Direction)
	offset += 4

	return arg
}
func (a Ardrone3AnimationsFlip) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var AnimationsFlip = Ardrone3AnimationsFlip{
	Project: ProjectArdrone3,
	Class:   ClassAnimations,
	Cmd:     CmdFlip,
}

// Ask the drone to move camera
const ClassCamera ClassDef = 1

// title : Move the camera,
// desc : Move the camera.\n You can get min and max values for tilt and pan using [CameraInfo](#0-15-0).,
// support : 0901;090c;090e,
// result : The drone moves its camera.\n Then, event [CameraOrientation](#1-25-0) is triggered.,
const CmdOrientation CmdDef = 0

type Ardrone3CameraOrientation Command

type Ardrone3CameraOrientationArguments struct {
	Tilt int8
	Pan  int8
}

func (a Ardrone3CameraOrientation) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3CameraOrientationArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Tilt)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Pan)
	offset++

	return arg
}
func (a Ardrone3CameraOrientation) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CameraOrientation = Ardrone3CameraOrientation{
	Project: ProjectArdrone3,
	Class:   ClassCamera,
	Cmd:     CmdOrientation,
}

// title : Move the camera,
// desc : Move the camera.\n You can get min and max values for tilt and pan using [CameraInfo](#0-15-0).,
// support : 0901;090c;090e,
// result : The drone moves its camera.\n Then, event [CameraOrientationV2](#1-25-2) is triggered.,
const CmdOrientationV2 CmdDef = 1

type Ardrone3CameraOrientationV2 Command

type Ardrone3CameraOrientationV2Arguments struct {
	Tilt float32
	Pan  float32
}

func (a Ardrone3CameraOrientationV2) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3CameraOrientationV2Arguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Tilt)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Pan)
	offset += 4

	return arg
}
func (a Ardrone3CameraOrientationV2) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CameraOrientationV2 = Ardrone3CameraOrientationV2{
	Project: ProjectArdrone3,
	Class:   ClassCamera,
	Cmd:     CmdOrientationV2,
}

// title : Move the camera using velocity,
// desc : Move the camera given velocity consign.\n You can get min and max values for tilt and pan using [CameraVelocityRange](#1-25-4).,
// support : 0901;090c;090e,
// result : The drone moves its camera.\n Then, event [CameraOrientationV2](#1-25-2) is triggered.,
const CmdVelocity CmdDef = 2

type Ardrone3CameraVelocity Command

type Ardrone3CameraVelocityArguments struct {
	Tilt float32
	Pan  float32
}

func (a Ardrone3CameraVelocity) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3CameraVelocityArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Tilt)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Pan)
	offset += 4

	return arg
}
func (a Ardrone3CameraVelocity) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CameraVelocity = Ardrone3CameraVelocity{
	Project: ProjectArdrone3,
	Class:   ClassCamera,
	Cmd:     CmdVelocity,
}

// Media recording management
const ClassMediaRecord ClassDef = 7

// title : Take a picture,
// desc : Take a picture.,
const CmdPicture CmdDef = 0

type Ardrone3MediaRecordPicture Command

type Ardrone3MediaRecordPictureArguments struct {
	Massstorageid uint8
}

func (a Ardrone3MediaRecordPicture) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3MediaRecordPictureArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Massstorageid)
	offset++

	return arg
}
func (a Ardrone3MediaRecordPicture) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MediaRecordPicture = Ardrone3MediaRecordPicture{
	Project: ProjectArdrone3,
	Class:   ClassMediaRecord,
	Cmd:     CmdPicture,
}

// title : Record a video,
// desc : Record a video.,
const CmdVideo CmdDef = 1

type Ardrone3MediaRecordVideo Command

type Ardrone3MediaRecordVideoArguments struct {
	Record        uint32
	Massstorageid uint8
}

func (a Ardrone3MediaRecordVideo) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3MediaRecordVideoArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Record)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Massstorageid)
	offset++

	return arg
}
func (a Ardrone3MediaRecordVideo) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MediaRecordVideo = Ardrone3MediaRecordVideo{
	Project: ProjectArdrone3,
	Class:   ClassMediaRecord,
	Cmd:     CmdVideo,
}

// title : Take a picture,
// desc : Take a picture.\n The type of picture taken is related to the picture setting.\n You can set the picture format by sending the command [SetPictureFormat](#1-19-0). You can also get the current picture format with [PictureFormat](#1-20-0).\n Please note that the time required to take the picture is highly related to this format.\n\n You can check if the picture taking is available with [PictureState](#1-8-2).\n Also, please note that if your picture format is different from snapshot, picture taking will stop video recording (it will restart after that the picture has been taken).,
// support : 0901:2.0.1;090c;090e,
// result : Event [PictureState](#1-8-2) will be triggered with a state busy.\n The drone will take a picture.\n Then, when picture has been taken, notification [PictureEvent](#1-3-0) is triggered.\n And normally [PictureState](#1-8-2) will be triggered with a state ready.,
const CmdPictureV2 CmdDef = 2

type Ardrone3MediaRecordPictureV2 Command

type Ardrone3MediaRecordPictureV2Arguments struct {
}

func (a Ardrone3MediaRecordPictureV2) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3MediaRecordPictureV2Arguments{}
	// No arguments to decode here !!

	return arg
}
func (a Ardrone3MediaRecordPictureV2) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MediaRecordPictureV2 = Ardrone3MediaRecordPictureV2{
	Project: ProjectArdrone3,
	Class:   ClassMediaRecord,
	Cmd:     CmdPictureV2,
}

// title : Record a video,
// desc : Record a video (or start timelapse).\n You can check if the video recording is available with [VideoState](#1-8-3).\n This command can start a video (obvious huh?), but also a timelapse if the timelapse mode is set. You can check if the timelapse mode is set with the event [TimelapseMode](#1-20-4).\n Also, please note that if your picture format is different from snapshot, picture taking will stop video recording (it will restart after the picture has been taken).,
// support : 0901:2.0.1;090c;090e,
// result : The drone will begin or stop to record the video (or timelapse).\n Then, event [VideoState](#1-8-3) will be triggered. Also, notification [VideoEvent](#1-3-1) is triggered.,
const CmdVideoV2 CmdDef = 3

type Ardrone3MediaRecordVideoV2 Command

type Ardrone3MediaRecordVideoV2Arguments struct {
	Record uint32
}

func (a Ardrone3MediaRecordVideoV2) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3MediaRecordVideoV2Arguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Record)
	offset += 4

	return arg
}
func (a Ardrone3MediaRecordVideoV2) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MediaRecordVideoV2 = Ardrone3MediaRecordVideoV2{
	Project: ProjectArdrone3,
	Class:   ClassMediaRecord,
	Cmd:     CmdVideoV2,
}

// State of media recording
const ClassMediaRecordState ClassDef = 8

// title : Picture state,
// desc : Picture state.,
const CmdPictureStateChanged CmdDef = 0

type Ardrone3MediaRecordStatePictureStateChanged Command

type Ardrone3MediaRecordStatePictureStateChangedArguments struct {
	State         uint8
	Massstorageid uint8
}

func (a Ardrone3MediaRecordStatePictureStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3MediaRecordStatePictureStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.State)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Massstorageid)
	offset++

	return arg
}
func (a Ardrone3MediaRecordStatePictureStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MediaRecordStatePictureStateChanged = Ardrone3MediaRecordStatePictureStateChanged{
	Project: ProjectArdrone3,
	Class:   ClassMediaRecordState,
	Cmd:     CmdPictureStateChanged,
}

// title : Video record state,
// desc : Picture record state.,
const CmdVideoStateChanged CmdDef = 1

type Ardrone3MediaRecordStateVideoStateChanged Command

type Ardrone3MediaRecordStateVideoStateChangedArguments struct {
	State         uint32
	Massstorageid uint8
}

func (a Ardrone3MediaRecordStateVideoStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3MediaRecordStateVideoStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.State)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Massstorageid)
	offset++

	return arg
}
func (a Ardrone3MediaRecordStateVideoStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MediaRecordStateVideoStateChanged = Ardrone3MediaRecordStateVideoStateChanged{
	Project: ProjectArdrone3,
	Class:   ClassMediaRecordState,
	Cmd:     CmdVideoStateChanged,
}

// title : Picture state,
// desc : Picture state.,
// support : 0901:2.0.1;090c;090e,
// triggered : by [TakePicture](#1-7-2) or by a change in the picture state,
const CmdPictureStateChangedV2 CmdDef = 2

type Ardrone3MediaRecordStatePictureStateChangedV2 Command

type Ardrone3MediaRecordStatePictureStateChangedV2Arguments struct {
	State uint32
	Error uint32
}

func (a Ardrone3MediaRecordStatePictureStateChangedV2) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3MediaRecordStatePictureStateChangedV2Arguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.State)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Error)
	offset += 4

	return arg
}
func (a Ardrone3MediaRecordStatePictureStateChangedV2) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MediaRecordStatePictureStateChangedV2 = Ardrone3MediaRecordStatePictureStateChangedV2{
	Project: ProjectArdrone3,
	Class:   ClassMediaRecordState,
	Cmd:     CmdPictureStateChangedV2,
}

// title : Video record state,
// desc : Video record state.,
// support : 0901:2.0.1;090c;090e,
// triggered : by [RecordVideo](#1-7-3) or by a change in the video state,
const CmdVideoStateChangedV2 CmdDef = 3

type Ardrone3MediaRecordStateVideoStateChangedV2 Command

type Ardrone3MediaRecordStateVideoStateChangedV2Arguments struct {
	State uint32
	Error uint32
}

func (a Ardrone3MediaRecordStateVideoStateChangedV2) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3MediaRecordStateVideoStateChangedV2Arguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.State)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Error)
	offset += 4

	return arg
}
func (a Ardrone3MediaRecordStateVideoStateChangedV2) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MediaRecordStateVideoStateChangedV2 = Ardrone3MediaRecordStateVideoStateChangedV2{
	Project: ProjectArdrone3,
	Class:   ClassMediaRecordState,
	Cmd:     CmdVideoStateChangedV2,
}

// title : Video resolution,
// desc : Video resolution.\n Informs about streaming and recording video resolutions.\n Note that this is only an indication about what the resolution should be. To know the real resolution, you should get it from the frame.,
// support : none,
// triggered : when the resolution changes.,
const CmdVideoResolutionState CmdDef = 4

type Ardrone3MediaRecordStateVideoResolutionState Command

type Ardrone3MediaRecordStateVideoResolutionStateArguments struct {
	Streaming uint32
	Recording uint32
}

func (a Ardrone3MediaRecordStateVideoResolutionState) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3MediaRecordStateVideoResolutionStateArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Streaming)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Recording)
	offset += 4

	return arg
}
func (a Ardrone3MediaRecordStateVideoResolutionState) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MediaRecordStateVideoResolutionState = Ardrone3MediaRecordStateVideoResolutionState{
	Project: ProjectArdrone3,
	Class:   ClassMediaRecordState,
	Cmd:     CmdVideoResolutionState,
}

// Events of media recording
const ClassMediaRecordEvent ClassDef = 3

// title : Picture taken,
// desc : Picture taken.\n\n **This event is a notification, you can't retrieve it in the cache of the device controller.**,
// support : 0901:2.0.1;090c;090e,
// triggered : after a [TakePicture](#1-7-2), when the picture has been taken (or it has failed).,
const CmdPictureEventChanged CmdDef = 0

type Ardrone3MediaRecordEventPictureEventChanged Command

type Ardrone3MediaRecordEventPictureEventChangedArguments struct {
	Event uint32
	Error uint32
}

func (a Ardrone3MediaRecordEventPictureEventChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3MediaRecordEventPictureEventChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Event)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Error)
	offset += 4

	return arg
}
func (a Ardrone3MediaRecordEventPictureEventChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MediaRecordEventPictureEventChanged = Ardrone3MediaRecordEventPictureEventChanged{
	Project: ProjectArdrone3,
	Class:   ClassMediaRecordEvent,
	Cmd:     CmdPictureEventChanged,
}

// title : Video record notification,
// desc : Video record notification.\n\n **This event is a notification, you can't retrieve it in the cache of the device controller.**,
// support : 0901:2.0.1;090c;090e,
// triggered : by [RecordVideo](#1-7-3) or a change in the video state.,
const CmdVideoEventChanged CmdDef = 1

type Ardrone3MediaRecordEventVideoEventChanged Command

type Ardrone3MediaRecordEventVideoEventChangedArguments struct {
	Event uint32
	Error uint32
}

func (a Ardrone3MediaRecordEventVideoEventChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3MediaRecordEventVideoEventChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Event)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Error)
	offset += 4

	return arg
}
func (a Ardrone3MediaRecordEventVideoEventChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MediaRecordEventVideoEventChanged = Ardrone3MediaRecordEventVideoEventChanged{
	Project: ProjectArdrone3,
	Class:   ClassMediaRecordEvent,
	Cmd:     CmdVideoEventChanged,
}

// State from drone
const ClassPilotingState ClassDef = 4

// title : Flying state,
// desc : Flying state.,
// support : 0901;090c;090e,
// triggered : when the flying state changes.,
const CmdFlyingStateChanged CmdDef = 1

type Ardrone3PilotingStateFlyingStateChanged Command

type Ardrone3PilotingStateFlyingStateChangedArguments struct {
	State uint32
}

func (a Ardrone3PilotingStateFlyingStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStateFlyingStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.State)
	offset += 4

	return arg
}
func (a Ardrone3PilotingStateFlyingStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStateFlyingStateChanged = Ardrone3PilotingStateFlyingStateChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdFlyingStateChanged,
}

// title : Alert state,
// desc : Alert state.,
// support : 0901;090c;090e,
// triggered : when an alert happens on the drone.,
const CmdAlertStateChanged CmdDef = 2

type Ardrone3PilotingStateAlertStateChanged Command

type Ardrone3PilotingStateAlertStateChangedArguments struct {
	State uint32
}

func (a Ardrone3PilotingStateAlertStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStateAlertStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.State)
	offset += 4

	return arg
}
func (a Ardrone3PilotingStateAlertStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStateAlertStateChanged = Ardrone3PilotingStateAlertStateChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdAlertStateChanged,
}

// title : Return home state,
// desc : Return home state.\n Availability is related to gps fix, magnetometer calibration.,
// support : 0901;090c;090e,
// triggered : by [ReturnHome](#1-0-5) or when the state of the return home changes.,
const CmdNavigateHomeStateChanged CmdDef = 3

type Ardrone3PilotingStateNavigateHomeStateChanged Command

type Ardrone3PilotingStateNavigateHomeStateChangedArguments struct {
	State  uint32
	Reason uint32
}

func (a Ardrone3PilotingStateNavigateHomeStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStateNavigateHomeStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.State)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Reason)
	offset += 4

	return arg
}
func (a Ardrone3PilotingStateNavigateHomeStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStateNavigateHomeStateChanged = Ardrone3PilotingStateNavigateHomeStateChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdNavigateHomeStateChanged,
}

// title : Drone's position changed,
// desc : Drone's position changed.,
// support : 0901;090c;090e,
// triggered : regularly.,
const CmdPositionChanged CmdDef = 4

type Ardrone3PilotingStatePositionChanged Command

type Ardrone3PilotingStatePositionChangedArguments struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

func (a Ardrone3PilotingStatePositionChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStatePositionChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Latitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Longitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Altitude)
	offset += 8

	return arg
}
func (a Ardrone3PilotingStatePositionChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStatePositionChanged = Ardrone3PilotingStatePositionChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdPositionChanged,
}

// title : Drone's speed changed,
// desc : Drone's speed changed.\n Expressed in the NED referential (North-East-Down).,
// support : 0901;090c;090e,
// triggered : regularly.,
const CmdSpeedChanged CmdDef = 5

type Ardrone3PilotingStateSpeedChanged Command

type Ardrone3PilotingStateSpeedChangedArguments struct {
	SpeedX float32
	SpeedY float32
	SpeedZ float32
}

func (a Ardrone3PilotingStateSpeedChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStateSpeedChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.SpeedX)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.SpeedY)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.SpeedZ)
	offset += 4

	return arg
}
func (a Ardrone3PilotingStateSpeedChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStateSpeedChanged = Ardrone3PilotingStateSpeedChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdSpeedChanged,
}

// title : Drone's attitude changed,
// desc : Drone's attitude changed.,
// support : 0901;090c;090e,
// triggered : regularly.,
const CmdAttitudeChanged CmdDef = 6

type Ardrone3PilotingStateAttitudeChanged Command

type Ardrone3PilotingStateAttitudeChangedArguments struct {
	Roll  float32
	Pitch float32
	Yaw   float32
}

func (a Ardrone3PilotingStateAttitudeChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStateAttitudeChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Roll)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Pitch)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Yaw)
	offset += 4

	return arg
}
func (a Ardrone3PilotingStateAttitudeChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStateAttitudeChanged = Ardrone3PilotingStateAttitudeChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdAttitudeChanged,
}

// title : Auto takeoff mode,
// desc : Auto takeoff mode,
const CmdAutoTakeOffModeChanged CmdDef = 7

type Ardrone3PilotingStateAutoTakeOffModeChanged Command

type Ardrone3PilotingStateAutoTakeOffModeChangedArguments struct {
	State uint8
}

func (a Ardrone3PilotingStateAutoTakeOffModeChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStateAutoTakeOffModeChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.State)
	offset++

	return arg
}
func (a Ardrone3PilotingStateAutoTakeOffModeChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStateAutoTakeOffModeChanged = Ardrone3PilotingStateAutoTakeOffModeChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdAutoTakeOffModeChanged,
}

// title : Drone's altitude changed,
// desc : Drone's altitude changed.\n The altitude reported is the altitude above the take off point.\n To get the altitude above sea level, see [PositionChanged](#1-4-4).,
// support : 0901;090c;090e,
// triggered : regularly.,
const CmdAltitudeChanged CmdDef = 8

type Ardrone3PilotingStateAltitudeChanged Command

type Ardrone3PilotingStateAltitudeChangedArguments struct {
	Altitude float64
}

func (a Ardrone3PilotingStateAltitudeChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStateAltitudeChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Altitude)
	offset += 8

	return arg
}
func (a Ardrone3PilotingStateAltitudeChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStateAltitudeChanged = Ardrone3PilotingStateAltitudeChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdAltitudeChanged,
}

// title : Drone's location changed,
// desc : Drone's location changed.\n This event is meant to replace [PositionChanged](#1-4-4).,
// support : 0901:4.0.0;090c:4.0.0,
// triggered : regularly.,
const CmdGpsLocationChanged CmdDef = 9

type Ardrone3PilotingStateGpsLocationChanged Command

type Ardrone3PilotingStateGpsLocationChangedArguments struct {
	Latitude          float64
	Longitude         float64
	Altitude          float64
	Latitudeaccuracy  int8
	Longitudeaccuracy int8
	Altitudeaccuracy  int8
}

func (a Ardrone3PilotingStateGpsLocationChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStateGpsLocationChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Latitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Longitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Altitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Latitudeaccuracy)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Longitudeaccuracy)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Altitudeaccuracy)
	offset++

	return arg
}
func (a Ardrone3PilotingStateGpsLocationChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStateGpsLocationChanged = Ardrone3PilotingStateGpsLocationChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdGpsLocationChanged,
}

// title : Landing state,
// desc : Landing state.\n Only available for fixed wings (which have two landing modes).,
// support : 090e,
// triggered : when the landing state changes.,
const CmdLandingStateChanged CmdDef = 10

type Ardrone3PilotingStateLandingStateChanged Command

type Ardrone3PilotingStateLandingStateChangedArguments struct {
	State uint32
}

func (a Ardrone3PilotingStateLandingStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStateLandingStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.State)
	offset += 4

	return arg
}
func (a Ardrone3PilotingStateLandingStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStateLandingStateChanged = Ardrone3PilotingStateLandingStateChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdLandingStateChanged,
}

// title : Drone's air speed changed,
// desc : Drone's air speed changed\n Expressed in the drone's referential.,
// support : 090e:1.2.0,
// triggered : regularly.,
const CmdAirSpeedChanged CmdDef = 11

type Ardrone3PilotingStateAirSpeedChanged Command

type Ardrone3PilotingStateAirSpeedChangedArguments struct {
	AirSpeed float32
}

func (a Ardrone3PilotingStateAirSpeedChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStateAirSpeedChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.AirSpeed)
	offset += 4

	return arg
}
func (a Ardrone3PilotingStateAirSpeedChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStateAirSpeedChanged = Ardrone3PilotingStateAirSpeedChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdAirSpeedChanged,
}

// title : Move to changed,
// desc : The drone moves or moved to a given location.,
// support : 090c:4.3.0,
// triggered : by [MoveTo](#1-0-10) or when the drone did reach the given position.,
const CmdMoveToChanged CmdDef = 12

type Ardrone3PilotingStatemoveToChanged Command

type Ardrone3PilotingStatemoveToChangedArguments struct {
	Latitude        float64
	Longitude       float64
	Altitude        float64
	Orientationmode uint32
	Heading         float32
	Status          uint32
}

func (a Ardrone3PilotingStatemoveToChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStatemoveToChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Latitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Longitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Altitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Orientationmode)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Heading)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Status)
	offset += 4

	return arg
}
func (a Ardrone3PilotingStatemoveToChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStatemoveToChanged = Ardrone3PilotingStatemoveToChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdMoveToChanged,
}

// title : Motion state,
// desc : Motion state.\n If [MotionDetection](#1-6-16) is disabled, motion is steady.\n This information is only valid when the drone is not flying.,
// support : 090c:4.3.0,
// triggered : when the [FlyingState](#1-4-1) is landed and the [MotionDetection](#1-6-16) is enabled and the motion state changes.\n This event is triggered at a filtered rate.,
const CmdMotionState CmdDef = 13

type Ardrone3PilotingStateMotionState Command

type Ardrone3PilotingStateMotionStateArguments struct {
	State uint32
}

func (a Ardrone3PilotingStateMotionState) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStateMotionStateArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.State)
	offset += 4

	return arg
}
func (a Ardrone3PilotingStateMotionState) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStateMotionState = Ardrone3PilotingStateMotionState{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdMotionState,
}

// title : Piloted POI state,
// desc : Piloted POI state.,
// support : 090c:4.3.0,
// triggered : by [StartPilotedPOI](#1-0-12) or [StopPilotedPOI](#1-0-13) or when piloted POI becomes unavailable.,
const CmdPilotedPOI CmdDef = 14

type Ardrone3PilotingStatePilotedPOI Command

type Ardrone3PilotingStatePilotedPOIArguments struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
	Status    uint32
}

func (a Ardrone3PilotingStatePilotedPOI) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStatePilotedPOIArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Latitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Longitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Altitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Status)
	offset += 4

	return arg
}
func (a Ardrone3PilotingStatePilotedPOI) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStatePilotedPOI = Ardrone3PilotingStatePilotedPOI{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdPilotedPOI,
}

// title : Return home battery capacity,
// desc : Battery capacity status to return home.,
// support : 090c:4.3.0,
// triggered : when the status of the battery capacity to do a return home changes. This means that it is triggered either when the battery level changes, when the distance to the home changes or when the position of the home changes.,
const CmdReturnHomeBatteryCapacity CmdDef = 15

type Ardrone3PilotingStateReturnHomeBatteryCapacity Command

type Ardrone3PilotingStateReturnHomeBatteryCapacityArguments struct {
	Status uint32
}

func (a Ardrone3PilotingStateReturnHomeBatteryCapacity) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStateReturnHomeBatteryCapacityArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Status)
	offset += 4

	return arg
}
func (a Ardrone3PilotingStateReturnHomeBatteryCapacity) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStateReturnHomeBatteryCapacity = Ardrone3PilotingStateReturnHomeBatteryCapacity{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdReturnHomeBatteryCapacity,
}

// title : Relative move changed,
// desc : Relative move changed.,
// triggered : by [MoveRelatively](#1-0-7), or [CancelRelativeMove](#1-0-14) or when the drone's relative move state changes.,
const CmdMoveByChanged CmdDef = 16

type Ardrone3PilotingStatemoveByChanged Command

type Ardrone3PilotingStatemoveByChangedArguments struct {
	DXAsked   float32
	DYAsked   float32
	DZAsked   float32
	DPsiAsked float32
	DX        float32
	DY        float32
	DZ        float32
	DPsi      float32
	Status    uint32
}

func (a Ardrone3PilotingStatemoveByChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStatemoveByChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.DXAsked)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.DYAsked)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.DZAsked)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.DPsiAsked)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.DX)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.DY)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.DZ)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.DPsi)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Status)
	offset += 4

	return arg
}
func (a Ardrone3PilotingStatemoveByChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStatemoveByChanged = Ardrone3PilotingStatemoveByChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdMoveByChanged,
}

// title : Hovering warning,
// desc : Indicate that the drone may have difficulties to maintain a fix position when hovering.,
// support : 0915,
// triggered : at connection and on changes.,
const CmdHoveringWarning CmdDef = 17

type Ardrone3PilotingStateHoveringWarning Command

type Ardrone3PilotingStateHoveringWarningArguments struct {
	Nogpstoodark uint8
	Nogpstoohigh uint8
}

func (a Ardrone3PilotingStateHoveringWarning) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStateHoveringWarningArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Nogpstoodark)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Nogpstoohigh)
	offset++

	return arg
}
func (a Ardrone3PilotingStateHoveringWarning) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStateHoveringWarning = Ardrone3PilotingStateHoveringWarning{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdHoveringWarning,
}

// title : Landing auto trigger.,
// desc : Forced landing auto trigger information.,
// support : ,
// triggered : at connection, and when forced landing auto trigger information changes, then every seconds while `reason` is different from `none`.,
const CmdForcedLandingAutoTrigger CmdDef = 18

type Ardrone3PilotingStateForcedLandingAutoTrigger Command

type Ardrone3PilotingStateForcedLandingAutoTriggerArguments struct {
	Reason uint32
	Delay  uint32
}

func (a Ardrone3PilotingStateForcedLandingAutoTrigger) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStateForcedLandingAutoTriggerArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Reason)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Delay)
	offset += 4

	return arg
}
func (a Ardrone3PilotingStateForcedLandingAutoTrigger) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStateForcedLandingAutoTrigger = Ardrone3PilotingStateForcedLandingAutoTrigger{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdForcedLandingAutoTrigger,
}

// title : Wind state,
// desc : Wind state.,
// support : 0914,
// triggered : at connection and on changes.,
const CmdWindStateChanged CmdDef = 19

type Ardrone3PilotingStateWindStateChanged Command

type Ardrone3PilotingStateWindStateChangedArguments struct {
	State uint32
}

func (a Ardrone3PilotingStateWindStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingStateWindStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.State)
	offset += 4

	return arg
}
func (a Ardrone3PilotingStateWindStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingStateWindStateChanged = Ardrone3PilotingStateWindStateChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingState,
	Cmd:     CmdWindStateChanged,
}

// Events of Piloting
const ClassPilotingEvent ClassDef = 34

// title : Relative move ended,
// desc : Relative move ended.\n Informs about the move that the drone managed to do and why it stopped.,
// support : 0901:3.3.0;090c:3.3.0,
// triggered : when the drone reaches its target or when it is interrupted by another [moveBy command](#1-0-7) or when an error occurs.,
const CmdMoveByEnd CmdDef = 0

type Ardrone3PilotingEventmoveByEnd Command

type Ardrone3PilotingEventmoveByEndArguments struct {
	DX    float32
	DY    float32
	DZ    float32
	DPsi  float32
	Error uint32
}

func (a Ardrone3PilotingEventmoveByEnd) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingEventmoveByEndArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.DX)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.DY)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.DZ)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.DPsi)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Error)
	offset += 4

	return arg
}
func (a Ardrone3PilotingEventmoveByEnd) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingEventmoveByEnd = Ardrone3PilotingEventmoveByEnd{
	Project: ProjectArdrone3,
	Class:   ClassPilotingEvent,
	Cmd:     CmdMoveByEnd,
}

// Network related commands
const ClassNetwork ClassDef = 13

// title : Scan wifi network,
// desc : Scan wifi network to get a list of all networks found by the drone,
// support : 0901;090c;090e,
// result : Event [WifiScanResults](#1-14-0) is triggered with all networks found.\n When all networks have been sent, event [WifiScanEnded](#1-14-1) is triggered.,
const CmdWifiScan CmdDef = 0

type Ardrone3NetworkWifiScan Command

type Ardrone3NetworkWifiScanArguments struct {
	Band uint32
}

func (a Ardrone3NetworkWifiScan) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3NetworkWifiScanArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Band)
	offset += 4

	return arg
}
func (a Ardrone3NetworkWifiScan) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var NetworkWifiScan = Ardrone3NetworkWifiScan{
	Project: ProjectArdrone3,
	Class:   ClassNetwork,
	Cmd:     CmdWifiScan,
}

// title : Ask for available wifi channels,
// desc : Ask for available wifi channels.\n The list of available Wifi channels is related to the country of the drone. You can get this country from the event [CountryChanged](#0-3-6).,
// support : 0901;090c;090e,
// result : Event [AvailableWifiChannels](#1-14-2) is triggered with all available channels. When all channels have been sent, event [AvailableWifiChannelsCompleted](#1-14-3) is triggered.,
const CmdWifiAuthChannel CmdDef = 1

type Ardrone3NetworkWifiAuthChannel Command

type Ardrone3NetworkWifiAuthChannelArguments struct {
}

func (a Ardrone3NetworkWifiAuthChannel) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3NetworkWifiAuthChannelArguments{}
	// No arguments to decode here !!

	return arg
}
func (a Ardrone3NetworkWifiAuthChannel) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var NetworkWifiAuthChannel = Ardrone3NetworkWifiAuthChannel{
	Project: ProjectArdrone3,
	Class:   ClassNetwork,
	Cmd:     CmdWifiAuthChannel,
}

// Network state from Product
const ClassNetworkState ClassDef = 14

// title : Wifi scan results,
// desc : Wifi scan results.\n Please note that the list is not complete until you receive the event [WifiScanEnded](#1-14-1).,
// support : 0901;090c;090e,
// triggered : for each wifi network scanned after a [ScanWifi](#1-13-0),
const CmdWifiScanListChanged CmdDef = 0

type Ardrone3NetworkStateWifiScanListChanged Command

type Ardrone3NetworkStateWifiScanListChangedArguments struct {
	Ssid    string
	Rssi    int16
	Band    uint32
	Channel uint8
}

func (a Ardrone3NetworkStateWifiScanListChanged) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := Ardrone3NetworkStateWifiScanListChangedArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Ssid = string(b[offset : offset+stringEnd])
	offset += stringEnd
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.Rssi)
	offset += 2
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Band)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Channel)
	offset++

	return arg
}
func (a Ardrone3NetworkStateWifiScanListChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var NetworkStateWifiScanListChanged = Ardrone3NetworkStateWifiScanListChanged{
	Project: ProjectArdrone3,
	Class:   ClassNetworkState,
	Cmd:     CmdWifiScanListChanged,
}

// title : Wifi scan ended,
// desc : Wifi scan ended.\n When receiving this event, the list of [WifiScanResults](#1-14-0) is complete.,
// support : 0901;090c;090e,
// triggered : after the last [WifiScanResult](#1-14-0) has been sent.,
const CmdAllWifiScanChanged CmdDef = 1

type Ardrone3NetworkStateAllWifiScanChanged Command

type Ardrone3NetworkStateAllWifiScanChangedArguments struct {
}

func (a Ardrone3NetworkStateAllWifiScanChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3NetworkStateAllWifiScanChangedArguments{}
	// No arguments to decode here !!

	return arg
}
func (a Ardrone3NetworkStateAllWifiScanChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var NetworkStateAllWifiScanChanged = Ardrone3NetworkStateAllWifiScanChanged{
	Project: ProjectArdrone3,
	Class:   ClassNetworkState,
	Cmd:     CmdAllWifiScanChanged,
}

// title : Available wifi channels,
// desc : Available wifi channels.\n Please note that the list is not complete until you receive the event [AvailableWifiChannelsCompleted](#1-14-3).,
// support : 0901;090c;090e,
// triggered : for each available channel after a [GetAvailableWifiChannels](#1-13-1).,
const CmdWifiAuthChannelListChanged CmdDef = 2

type Ardrone3NetworkStateWifiAuthChannelListChanged Command

type Ardrone3NetworkStateWifiAuthChannelListChangedArguments struct {
	Band    uint32
	Channel uint8
	Inorout uint8
}

func (a Ardrone3NetworkStateWifiAuthChannelListChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3NetworkStateWifiAuthChannelListChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Band)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Channel)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Inorout)
	offset++

	return arg
}
func (a Ardrone3NetworkStateWifiAuthChannelListChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var NetworkStateWifiAuthChannelListChanged = Ardrone3NetworkStateWifiAuthChannelListChanged{
	Project: ProjectArdrone3,
	Class:   ClassNetworkState,
	Cmd:     CmdWifiAuthChannelListChanged,
}

// title : Available wifi channels completed,
// desc : Available wifi channels completed.\n When receiving this event, the list of [AvailableWifiChannels](#1-14-2) is complete.,
// support : 0901;090c;090e,
// triggered : after the last [AvailableWifiChannel](#1-14-2) has been sent.,
const CmdAllWifiAuthChannelChanged CmdDef = 3

type Ardrone3NetworkStateAllWifiAuthChannelChanged Command

type Ardrone3NetworkStateAllWifiAuthChannelChangedArguments struct {
}

func (a Ardrone3NetworkStateAllWifiAuthChannelChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3NetworkStateAllWifiAuthChannelChangedArguments{}
	// No arguments to decode here !!

	return arg
}
func (a Ardrone3NetworkStateAllWifiAuthChannelChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var NetworkStateAllWifiAuthChannelChanged = Ardrone3NetworkStateAllWifiAuthChannelChanged{
	Project: ProjectArdrone3,
	Class:   ClassNetworkState,
	Cmd:     CmdAllWifiAuthChannelChanged,
}

// Piloting Settings commands
const ClassPilotingSettings ClassDef = 2

// title : Set max altitude,
// desc : Set max altitude.\n The drone will not fly over this max altitude when it is in manual piloting.\n Please note that if you set a max altitude which is below the current drone altitude, the drone will not go to given max altitude.\n You can get the bounds in the event [MaxAltitude](#1-6-0).,
// support : 0901;090c;090e,
// result : The max altitude is set.\n Then, event [MaxAltitude](#1-6-0) is triggered.,
const CmdMaxAltitude CmdDef = 0

type Ardrone3PilotingSettingsMaxAltitude Command

type Ardrone3PilotingSettingsMaxAltitudeArguments struct {
	Current float32
}

func (a Ardrone3PilotingSettingsMaxAltitude) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsMaxAltitudeArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Current)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingsMaxAltitude) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsMaxAltitude = Ardrone3PilotingSettingsMaxAltitude{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettings,
	Cmd:     CmdMaxAltitude,
}

// title : Set max pitch/roll,
// desc : Set max pitch/roll.\n This represent the max inclination allowed by the drone.\n You can get the bounds with the commands [MaxPitchRoll](#1-6-1).,
// support : 0901;090c,
// result : The max pitch/roll is set.\n Then, event [MaxPitchRoll](#1-6-1) is triggered.,
const CmdMaxTilt CmdDef = 1

type Ardrone3PilotingSettingsMaxTilt Command

type Ardrone3PilotingSettingsMaxTiltArguments struct {
	Current float32
}

func (a Ardrone3PilotingSettingsMaxTilt) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsMaxTiltArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Current)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingsMaxTilt) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsMaxTilt = Ardrone3PilotingSettingsMaxTilt{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettings,
	Cmd:     CmdMaxTilt,
}

// title : Set absolut control,
// desc : Set absolut control.,
const CmdAbsolutControl CmdDef = 2

type Ardrone3PilotingSettingsAbsolutControl Command

type Ardrone3PilotingSettingsAbsolutControlArguments struct {
	On uint8
}

func (a Ardrone3PilotingSettingsAbsolutControl) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsAbsolutControlArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.On)
	offset++

	return arg
}
func (a Ardrone3PilotingSettingsAbsolutControl) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsAbsolutControl = Ardrone3PilotingSettingsAbsolutControl{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettings,
	Cmd:     CmdAbsolutControl,
}

// title : Set max distance,
// desc : Set max distance.\n You can get the bounds from the event [MaxDistance](#1-6-3).\n\n If [Geofence](#1-6-4) is activated, the drone won't fly over the given max distance.,
// support : 0901;090c;090e,
// result : The max distance is set.\n Then, event [MaxDistance](#1-6-3) is triggered.,
const CmdMaxDistance CmdDef = 3

type Ardrone3PilotingSettingsMaxDistance Command

type Ardrone3PilotingSettingsMaxDistanceArguments struct {
	Value float32
}

func (a Ardrone3PilotingSettingsMaxDistance) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsMaxDistanceArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingsMaxDistance) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsMaxDistance = Ardrone3PilotingSettingsMaxDistance{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettings,
	Cmd:     CmdMaxDistance,
}

// title : Enable geofence,
// desc : Enable geofence.\n If geofence is enabled, the drone won't fly over the given max distance.\n You can get the max distance from the event [MaxDistance](#1-6-3). \n For copters: the distance is computed from the controller position, if this position is not known, it will use the take off.\n For fixed wings: the distance is computed from the take off position.,
// support : 0901;090c;090e,
// result : Geofencing is enabled or disabled.\n Then, event [Geofencing](#1-6-4) is triggered.,
const CmdNoFlyOverMaxDistance CmdDef = 4

type Ardrone3PilotingSettingsNoFlyOverMaxDistance Command

type Ardrone3PilotingSettingsNoFlyOverMaxDistanceArguments struct {
	ShouldNotFlyOver uint8
}

func (a Ardrone3PilotingSettingsNoFlyOverMaxDistance) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsNoFlyOverMaxDistanceArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.ShouldNotFlyOver)
	offset++

	return arg
}
func (a Ardrone3PilotingSettingsNoFlyOverMaxDistance) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsNoFlyOverMaxDistance = Ardrone3PilotingSettingsNoFlyOverMaxDistance{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettings,
	Cmd:     CmdNoFlyOverMaxDistance,
}

// title : Set autonomous flight max horizontal speed,
// desc : Set autonomous flight max horizontal speed.\n This will only be used during autonomous flights such as moveBy.,
// support : 0901:3.3.0;090c:3.3.0,
// result : The max horizontal speed is set.\n Then, event [AutonomousFlightMaxHorizontalSpeed](#1-6-5) is triggered.,
const CmdSetAutonomousFlightMaxHorizontalSpeed CmdDef = 5

type Ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalSpeed Command

type Ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalSpeedArguments struct {
	Value float32
}

func (a Ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalSpeed) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalSpeedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalSpeed) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingssetAutonomousFlightMaxHorizontalSpeed = Ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalSpeed{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettings,
	Cmd:     CmdSetAutonomousFlightMaxHorizontalSpeed,
}

// title : Set autonomous flight max vertical speed,
// desc : Set autonomous flight max vertical speed.\n This will only be used during autonomous flights such as moveBy.,
// support : 0901:3.3.0;090c:3.3.0,
// result : The max vertical speed is set.\n Then, event [AutonomousFlightMaxVerticalSpeed](#1-6-6) is triggered.,
const CmdSetAutonomousFlightMaxVerticalSpeed CmdDef = 6

type Ardrone3PilotingSettingssetAutonomousFlightMaxVerticalSpeed Command

type Ardrone3PilotingSettingssetAutonomousFlightMaxVerticalSpeedArguments struct {
	Value float32
}

func (a Ardrone3PilotingSettingssetAutonomousFlightMaxVerticalSpeed) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingssetAutonomousFlightMaxVerticalSpeedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingssetAutonomousFlightMaxVerticalSpeed) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingssetAutonomousFlightMaxVerticalSpeed = Ardrone3PilotingSettingssetAutonomousFlightMaxVerticalSpeed{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettings,
	Cmd:     CmdSetAutonomousFlightMaxVerticalSpeed,
}

// title : Set autonomous flight max horizontal acceleration,
// desc : Set autonomous flight max horizontal acceleration.\n This will only be used during autonomous flights such as moveBy.,
// support : 0901:3.3.0;090c:3.3.0,
// result : The max horizontal acceleration is set.\n Then, event [AutonomousFlightMaxHorizontalAcceleration](#1-6-7) is triggered.,
const CmdSetAutonomousFlightMaxHorizontalAcceleration CmdDef = 7

type Ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalAcceleration Command

type Ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalAccelerationArguments struct {
	Value float32
}

func (a Ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalAcceleration) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalAccelerationArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalAcceleration) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingssetAutonomousFlightMaxHorizontalAcceleration = Ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalAcceleration{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettings,
	Cmd:     CmdSetAutonomousFlightMaxHorizontalAcceleration,
}

// title : Set autonomous flight max vertical acceleration,
// desc : Set autonomous flight max vertical acceleration.\n This will only be used during autonomous flights such as moveBy.,
// support : 0901:3.3.0;090c:3.3.0,
// result : The max vertical acceleration is set.\n Then, event [AutonomousFlightMaxVerticalAcceleration](#1-6-8) is triggered.,
const CmdSetAutonomousFlightMaxVerticalAcceleration CmdDef = 8

type Ardrone3PilotingSettingssetAutonomousFlightMaxVerticalAcceleration Command

type Ardrone3PilotingSettingssetAutonomousFlightMaxVerticalAccelerationArguments struct {
	Value float32
}

func (a Ardrone3PilotingSettingssetAutonomousFlightMaxVerticalAcceleration) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingssetAutonomousFlightMaxVerticalAccelerationArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingssetAutonomousFlightMaxVerticalAcceleration) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingssetAutonomousFlightMaxVerticalAcceleration = Ardrone3PilotingSettingssetAutonomousFlightMaxVerticalAcceleration{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettings,
	Cmd:     CmdSetAutonomousFlightMaxVerticalAcceleration,
}

// title : Set autonomous flight max rotation speed,
// desc : Set autonomous flight max rotation speed.\n This will only be used during autonomous flights such as moveBy.,
// support : 0901:3.3.0;090c:3.3.0,
// result : The max rotation speed is set.\n Then, event [AutonomousFlightMaxRotationSpeed](#1-6-9) is triggered.,
const CmdSetAutonomousFlightMaxRotationSpeed CmdDef = 9

type Ardrone3PilotingSettingssetAutonomousFlightMaxRotationSpeed Command

type Ardrone3PilotingSettingssetAutonomousFlightMaxRotationSpeedArguments struct {
	Value float32
}

func (a Ardrone3PilotingSettingssetAutonomousFlightMaxRotationSpeed) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingssetAutonomousFlightMaxRotationSpeedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingssetAutonomousFlightMaxRotationSpeed) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingssetAutonomousFlightMaxRotationSpeed = Ardrone3PilotingSettingssetAutonomousFlightMaxRotationSpeed{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettings,
	Cmd:     CmdSetAutonomousFlightMaxRotationSpeed,
}

// title : Set banked turn mode,
// desc : Set banked turn mode.\n When banked turn mode is enabled, the drone will use yaw values from the piloting command to infer with roll and pitch on the drone when its horizontal speed is not null.,
// support : 0901:3.2.0;090c:3.2.0,
// result : The banked turn mode is enabled or disabled.\n Then, event [BankedTurnMode](#1-6-10) is triggered.,
const CmdBankedTurn CmdDef = 10

type Ardrone3PilotingSettingsBankedTurn Command

type Ardrone3PilotingSettingsBankedTurnArguments struct {
	Value uint8
}

func (a Ardrone3PilotingSettingsBankedTurn) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsBankedTurnArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Value)
	offset++

	return arg
}
func (a Ardrone3PilotingSettingsBankedTurn) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsBankedTurn = Ardrone3PilotingSettingsBankedTurn{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettings,
	Cmd:     CmdBankedTurn,
}

// title : Set minimum altitude,
// desc : Set minimum altitude.\n Only available for fixed wings.,
// support : 090e,
// result : The minimum altitude is set.\n Then, event [MinimumAltitude](#1-6-11) is triggered.,
const CmdMinAltitude CmdDef = 11

type Ardrone3PilotingSettingsMinAltitude Command

type Ardrone3PilotingSettingsMinAltitudeArguments struct {
	Current float32
}

func (a Ardrone3PilotingSettingsMinAltitude) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsMinAltitudeArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Current)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingsMinAltitude) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsMinAltitude = Ardrone3PilotingSettingsMinAltitude{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettings,
	Cmd:     CmdMinAltitude,
}

// title : Set default circling direction,
// desc : Set default circling direction. This direction will be used when the drone use an automatic circling or when [CIRCLE](#1-0-9) is sent with direction *default*.\n Only available for fixed wings.,
// support : 090e,
// result : The circling direction is set.\n Then, event [DefaultCirclingDirection](#1-6-12) is triggered.,
const CmdCirclingDirection CmdDef = 12

type Ardrone3PilotingSettingsCirclingDirection Command

type Ardrone3PilotingSettingsCirclingDirectionArguments struct {
	Value uint32
}

func (a Ardrone3PilotingSettingsCirclingDirection) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsCirclingDirectionArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingsCirclingDirection) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsCirclingDirection = Ardrone3PilotingSettingsCirclingDirection{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettings,
	Cmd:     CmdCirclingDirection,
}

// title : Set circling radius,
// desc : Set circling radius.\n Only available for fixed wings.,
// support : none,
// result : The circling radius is set.\n Then, event [CirclingRadius](#1-6-13) is triggered.,
const CmdCirclingRadius CmdDef = 13

type Ardrone3PilotingSettingsCirclingRadius Command

type Ardrone3PilotingSettingsCirclingRadiusArguments struct {
	Value uint16
}

func (a Ardrone3PilotingSettingsCirclingRadius) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsCirclingRadiusArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.Value)
	offset += 2

	return arg
}
func (a Ardrone3PilotingSettingsCirclingRadius) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsCirclingRadius = Ardrone3PilotingSettingsCirclingRadius{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettings,
	Cmd:     CmdCirclingRadius,
}

// title : Set min circling altitude,
// desc : Set min circling altitude (not used during take off).\n Only available for fixed wings.,
// support : 090e,
// result : The circling altitude is set.\n Then, event [CirclingAltitude](#1-6-14) is triggered.,
const CmdCirclingAltitude CmdDef = 14

type Ardrone3PilotingSettingsCirclingAltitude Command

type Ardrone3PilotingSettingsCirclingAltitudeArguments struct {
	Value uint16
}

func (a Ardrone3PilotingSettingsCirclingAltitude) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsCirclingAltitudeArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.Value)
	offset += 2

	return arg
}
func (a Ardrone3PilotingSettingsCirclingAltitude) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsCirclingAltitude = Ardrone3PilotingSettingsCirclingAltitude{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettings,
	Cmd:     CmdCirclingAltitude,
}

// title : Set pitch mode,
// desc : Set pitch mode.\n Only available for fixed wings.,
// support : 090e,
// result : The pitch mode is set.\n Then, event [PitchMode](#1-6-15) is triggered.,
const CmdPitchMode CmdDef = 15

type Ardrone3PilotingSettingsPitchMode Command

type Ardrone3PilotingSettingsPitchModeArguments struct {
	Value uint32
}

func (a Ardrone3PilotingSettingsPitchMode) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsPitchModeArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingsPitchMode) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsPitchMode = Ardrone3PilotingSettingsPitchMode{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettings,
	Cmd:     CmdPitchMode,
}

// title : Enable/disable the motion detection,
// desc : Enable/disable the motion detection.\n If the motion detection is enabled, the drone will send its [MotionState](#1-4-13) when its [FlyingState](#1-4-1) is landed. If the motion detection is disabled, [MotionState](#1-4-13) is steady.,
// support : 090c:4.3.0,
// result : The motion detection is enabled or disabled.\n Then, event [MotionDetection](#1-6-16) is triggered. After that, if enabled and [FlyingState](#1-4-1) is landed, the [MotionState](#1-4-13) is triggered upon changes.,
const CmdSetMotionDetectionMode CmdDef = 16

type Ardrone3PilotingSettingsSetMotionDetectionMode Command

type Ardrone3PilotingSettingsSetMotionDetectionModeArguments struct {
	Enable uint8
}

func (a Ardrone3PilotingSettingsSetMotionDetectionMode) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsSetMotionDetectionModeArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Enable)
	offset++

	return arg
}
func (a Ardrone3PilotingSettingsSetMotionDetectionMode) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsSetMotionDetectionMode = Ardrone3PilotingSettingsSetMotionDetectionMode{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettings,
	Cmd:     CmdSetMotionDetectionMode,
}

// Piloting Settings state from product
const ClassPilotingSettingsState ClassDef = 6

// title : Max altitude,
// desc : Max altitude.\n The drone will not fly higher than this altitude (above take off point).,
// support : 0901;090c;090e,
// triggered : by [SetMaxAltitude](#1-2-0).,
const CmdMaxAltitudeChanged CmdDef = 0

type Ardrone3PilotingSettingsStateMaxAltitudeChanged Command

type Ardrone3PilotingSettingsStateMaxAltitudeChangedArguments struct {
	Current float32
	Min     float32
	Max     float32
}

func (a Ardrone3PilotingSettingsStateMaxAltitudeChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsStateMaxAltitudeChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Current)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Min)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Max)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingsStateMaxAltitudeChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsStateMaxAltitudeChanged = Ardrone3PilotingSettingsStateMaxAltitudeChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettingsState,
	Cmd:     CmdMaxAltitudeChanged,
}

// title : Max pitch/roll,
// desc : Max pitch/roll.\n The drone will not fly higher than this altitude (above take off point).,
// support : 0901;090c,
// triggered : by [SetMaxAltitude](#1-2-0).,
const CmdMaxTiltChanged CmdDef = 1

type Ardrone3PilotingSettingsStateMaxTiltChanged Command

type Ardrone3PilotingSettingsStateMaxTiltChangedArguments struct {
	Current float32
	Min     float32
	Max     float32
}

func (a Ardrone3PilotingSettingsStateMaxTiltChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsStateMaxTiltChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Current)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Min)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Max)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingsStateMaxTiltChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsStateMaxTiltChanged = Ardrone3PilotingSettingsStateMaxTiltChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettingsState,
	Cmd:     CmdMaxTiltChanged,
}

// title : Absolut control,
// desc : Absolut control.,
const CmdAbsolutControlChanged CmdDef = 2

type Ardrone3PilotingSettingsStateAbsolutControlChanged Command

type Ardrone3PilotingSettingsStateAbsolutControlChangedArguments struct {
	On uint8
}

func (a Ardrone3PilotingSettingsStateAbsolutControlChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsStateAbsolutControlChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.On)
	offset++

	return arg
}
func (a Ardrone3PilotingSettingsStateAbsolutControlChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsStateAbsolutControlChanged = Ardrone3PilotingSettingsStateAbsolutControlChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettingsState,
	Cmd:     CmdAbsolutControlChanged,
}

// title : Max distance,
// desc : Max distance.,
// support : 0901;090c;090e,
// triggered : by [SetMaxDistance](#1-2-3).,
const CmdMaxDistanceChanged CmdDef = 3

type Ardrone3PilotingSettingsStateMaxDistanceChanged Command

type Ardrone3PilotingSettingsStateMaxDistanceChangedArguments struct {
	Current float32
	Min     float32
	Max     float32
}

func (a Ardrone3PilotingSettingsStateMaxDistanceChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsStateMaxDistanceChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Current)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Min)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Max)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingsStateMaxDistanceChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsStateMaxDistanceChanged = Ardrone3PilotingSettingsStateMaxDistanceChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettingsState,
	Cmd:     CmdMaxDistanceChanged,
}

// title : Geofencing,
// desc : Geofencing.\n If set, the drone won't fly over the [MaxDistance](#1-6-3).,
// support : 0901;090c;090e,
// triggered : by [EnableGeofence](#1-2-4).,
const CmdNoFlyOverMaxDistanceChanged CmdDef = 4

type Ardrone3PilotingSettingsStateNoFlyOverMaxDistanceChanged Command

type Ardrone3PilotingSettingsStateNoFlyOverMaxDistanceChangedArguments struct {
	ShouldNotFlyOver uint8
}

func (a Ardrone3PilotingSettingsStateNoFlyOverMaxDistanceChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsStateNoFlyOverMaxDistanceChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.ShouldNotFlyOver)
	offset++

	return arg
}
func (a Ardrone3PilotingSettingsStateNoFlyOverMaxDistanceChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsStateNoFlyOverMaxDistanceChanged = Ardrone3PilotingSettingsStateNoFlyOverMaxDistanceChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettingsState,
	Cmd:     CmdNoFlyOverMaxDistanceChanged,
}

// title : Autonomous flight max horizontal speed,
// desc : Autonomous flight max horizontal speed.,
// support : 0901:3.3.0;090c:3.3.0,
// triggered : by [SetAutonomousFlightMaxHorizontalSpeed](#1-2-5).,
const CmdAutonomousFlightMaxHorizontalSpeed CmdDef = 5

type Ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalSpeed Command

type Ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalSpeedArguments struct {
	Value float32
}

func (a Ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalSpeed) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalSpeedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalSpeed) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsStateAutonomousFlightMaxHorizontalSpeed = Ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalSpeed{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettingsState,
	Cmd:     CmdAutonomousFlightMaxHorizontalSpeed,
}

// title : Autonomous flight max vertical speed,
// desc : Autonomous flight max vertical speed.,
// support : 0901:3.3.0;090c:3.3.0,
// triggered : by [SetAutonomousFlightMaxVerticalSpeed](#1-2-6).,
const CmdAutonomousFlightMaxVerticalSpeed CmdDef = 6

type Ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalSpeed Command

type Ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalSpeedArguments struct {
	Value float32
}

func (a Ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalSpeed) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalSpeedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalSpeed) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsStateAutonomousFlightMaxVerticalSpeed = Ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalSpeed{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettingsState,
	Cmd:     CmdAutonomousFlightMaxVerticalSpeed,
}

// title : Autonomous flight max horizontal acceleration,
// desc : Autonomous flight max horizontal acceleration.,
// support : 0901:3.3.0;090c:3.3.0,
// triggered : by [SetAutonomousFlightMaxHorizontalAcceleration](#1-2-7).,
const CmdAutonomousFlightMaxHorizontalAcceleration CmdDef = 7

type Ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration Command

type Ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalAccelerationArguments struct {
	Value float32
}

func (a Ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalAccelerationArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration = Ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettingsState,
	Cmd:     CmdAutonomousFlightMaxHorizontalAcceleration,
}

// title : Autonomous flight max vertical acceleration,
// desc : Autonomous flight max vertical acceleration.,
// support : 0901:3.3.0;090c:3.3.0,
// triggered : by [SetAutonomousFlightMaxVerticalAcceleration](#1-2-8).,
const CmdAutonomousFlightMaxVerticalAcceleration CmdDef = 8

type Ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalAcceleration Command

type Ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalAccelerationArguments struct {
	Value float32
}

func (a Ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalAcceleration) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalAccelerationArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalAcceleration) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsStateAutonomousFlightMaxVerticalAcceleration = Ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalAcceleration{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettingsState,
	Cmd:     CmdAutonomousFlightMaxVerticalAcceleration,
}

// title : Autonomous flight max rotation speed,
// desc : Autonomous flight max rotation speed.,
// support : 0901:3.3.0;090c:3.3.0,
// triggered : by [SetAutonomousFlightMaxRotationSpeed](#1-2-9).,
const CmdAutonomousFlightMaxRotationSpeed CmdDef = 9

type Ardrone3PilotingSettingsStateAutonomousFlightMaxRotationSpeed Command

type Ardrone3PilotingSettingsStateAutonomousFlightMaxRotationSpeedArguments struct {
	Value float32
}

func (a Ardrone3PilotingSettingsStateAutonomousFlightMaxRotationSpeed) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsStateAutonomousFlightMaxRotationSpeedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingsStateAutonomousFlightMaxRotationSpeed) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsStateAutonomousFlightMaxRotationSpeed = Ardrone3PilotingSettingsStateAutonomousFlightMaxRotationSpeed{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettingsState,
	Cmd:     CmdAutonomousFlightMaxRotationSpeed,
}

// title : Banked Turn mode,
// desc : Banked Turn mode.\n If banked turn mode is enabled, the drone will use yaw values from the piloting command to infer with roll and pitch on the drone when its horizontal speed is not null.,
// support : 0901:3.2.0;090c:3.2.0,
// triggered : by [SetBankedTurnMode](#1-2-10).,
const CmdBankedTurnChanged CmdDef = 10

type Ardrone3PilotingSettingsStateBankedTurnChanged Command

type Ardrone3PilotingSettingsStateBankedTurnChangedArguments struct {
	State uint8
}

func (a Ardrone3PilotingSettingsStateBankedTurnChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsStateBankedTurnChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.State)
	offset++

	return arg
}
func (a Ardrone3PilotingSettingsStateBankedTurnChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsStateBankedTurnChanged = Ardrone3PilotingSettingsStateBankedTurnChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettingsState,
	Cmd:     CmdBankedTurnChanged,
}

// title : Min altitude,
// desc : Min altitude.\n Only sent by fixed wings.,
// support : 090e,
// triggered : by [SetMinAltitude](#1-2-11).,
const CmdMinAltitudeChanged CmdDef = 11

type Ardrone3PilotingSettingsStateMinAltitudeChanged Command

type Ardrone3PilotingSettingsStateMinAltitudeChangedArguments struct {
	Current float32
	Min     float32
	Max     float32
}

func (a Ardrone3PilotingSettingsStateMinAltitudeChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsStateMinAltitudeChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Current)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Min)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Max)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingsStateMinAltitudeChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsStateMinAltitudeChanged = Ardrone3PilotingSettingsStateMinAltitudeChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettingsState,
	Cmd:     CmdMinAltitudeChanged,
}

// title : Circling direction,
// desc : Circling direction.\n Only sent by fixed wings.,
// support : 090e,
// triggered : by [SetCirclingDirection](#1-2-12).,
const CmdCirclingDirectionChanged CmdDef = 12

type Ardrone3PilotingSettingsStateCirclingDirectionChanged Command

type Ardrone3PilotingSettingsStateCirclingDirectionChangedArguments struct {
	Value uint32
}

func (a Ardrone3PilotingSettingsStateCirclingDirectionChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsStateCirclingDirectionChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingsStateCirclingDirectionChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsStateCirclingDirectionChanged = Ardrone3PilotingSettingsStateCirclingDirectionChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettingsState,
	Cmd:     CmdCirclingDirectionChanged,
}

// title : Circling radius,
// desc : Circling radius.\n Only sent by fixed wings.,
// support : none,
// triggered : by [SetCirclingRadius](#1-2-13).,
const CmdCirclingRadiusChanged CmdDef = 13

type Ardrone3PilotingSettingsStateCirclingRadiusChanged Command

type Ardrone3PilotingSettingsStateCirclingRadiusChangedArguments struct {
	Current uint16
	Min     uint16
	Max     uint16
}

func (a Ardrone3PilotingSettingsStateCirclingRadiusChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsStateCirclingRadiusChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.Current)
	offset += 2
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.Min)
	offset += 2
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.Max)
	offset += 2

	return arg
}
func (a Ardrone3PilotingSettingsStateCirclingRadiusChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsStateCirclingRadiusChanged = Ardrone3PilotingSettingsStateCirclingRadiusChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettingsState,
	Cmd:     CmdCirclingRadiusChanged,
}

// title : Circling altitude,
// desc : Circling altitude.\n Bounds will be automatically adjusted according to the [MaxAltitude](#1-6-0).\n Only sent by fixed wings.,
// support : 090e,
// triggered : by [SetCirclingRadius](#1-2-14) or when bounds change due to [SetMaxAltitude](#1-2-0).,
const CmdCirclingAltitudeChanged CmdDef = 14

type Ardrone3PilotingSettingsStateCirclingAltitudeChanged Command

type Ardrone3PilotingSettingsStateCirclingAltitudeChangedArguments struct {
	Current uint16
	Min     uint16
	Max     uint16
}

func (a Ardrone3PilotingSettingsStateCirclingAltitudeChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsStateCirclingAltitudeChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.Current)
	offset += 2
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.Min)
	offset += 2
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.Max)
	offset += 2

	return arg
}
func (a Ardrone3PilotingSettingsStateCirclingAltitudeChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsStateCirclingAltitudeChanged = Ardrone3PilotingSettingsStateCirclingAltitudeChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettingsState,
	Cmd:     CmdCirclingAltitudeChanged,
}

// title : Pitch mode,
// desc : Pitch mode.,
// support : 090e,
// triggered : by [SetPitchMode](#1-2-15).,
const CmdPitchModeChanged CmdDef = 15

type Ardrone3PilotingSettingsStatePitchModeChanged Command

type Ardrone3PilotingSettingsStatePitchModeChangedArguments struct {
	Value uint32
}

func (a Ardrone3PilotingSettingsStatePitchModeChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsStatePitchModeChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4

	return arg
}
func (a Ardrone3PilotingSettingsStatePitchModeChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsStatePitchModeChanged = Ardrone3PilotingSettingsStatePitchModeChanged{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettingsState,
	Cmd:     CmdPitchModeChanged,
}

// title : State of the motion detection,
// desc : State of the motion detection.,
// support : 090c:4.3.0,
// triggered : by [SetMotionDetectionMode](#1-2-16),
const CmdMotionDetection CmdDef = 16

type Ardrone3PilotingSettingsStateMotionDetection Command

type Ardrone3PilotingSettingsStateMotionDetectionArguments struct {
	Enabled uint8
}

func (a Ardrone3PilotingSettingsStateMotionDetection) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PilotingSettingsStateMotionDetectionArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Enabled)
	offset++

	return arg
}
func (a Ardrone3PilotingSettingsStateMotionDetection) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PilotingSettingsStateMotionDetection = Ardrone3PilotingSettingsStateMotionDetection{
	Project: ProjectArdrone3,
	Class:   ClassPilotingSettingsState,
	Cmd:     CmdMotionDetection,
}

// Speed Settings commands
const ClassSpeedSettings ClassDef = 11

// title : Set max vertical speed,
// desc : Set max vertical speed.,
// support : 0901;090c,
// result : The max vertical speed is set.\n Then, event [MaxVerticalSpeed](#1-12-0) is triggered.,
const CmdMaxVerticalSpeed CmdDef = 0

type Ardrone3SpeedSettingsMaxVerticalSpeed Command

type Ardrone3SpeedSettingsMaxVerticalSpeedArguments struct {
	Current float32
}

func (a Ardrone3SpeedSettingsMaxVerticalSpeed) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3SpeedSettingsMaxVerticalSpeedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Current)
	offset += 4

	return arg
}
func (a Ardrone3SpeedSettingsMaxVerticalSpeed) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SpeedSettingsMaxVerticalSpeed = Ardrone3SpeedSettingsMaxVerticalSpeed{
	Project: ProjectArdrone3,
	Class:   ClassSpeedSettings,
	Cmd:     CmdMaxVerticalSpeed,
}

// title : Set max rotation speed,
// desc : Set max rotation speed.,
// support : 0901;090c,
// result : The max rotation speed is set.\n Then, event [MaxRotationSpeed](#1-12-1) is triggered.,
const CmdMaxRotationSpeed CmdDef = 1

type Ardrone3SpeedSettingsMaxRotationSpeed Command

type Ardrone3SpeedSettingsMaxRotationSpeedArguments struct {
	Current float32
}

func (a Ardrone3SpeedSettingsMaxRotationSpeed) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3SpeedSettingsMaxRotationSpeedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Current)
	offset += 4

	return arg
}
func (a Ardrone3SpeedSettingsMaxRotationSpeed) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SpeedSettingsMaxRotationSpeed = Ardrone3SpeedSettingsMaxRotationSpeed{
	Project: ProjectArdrone3,
	Class:   ClassSpeedSettings,
	Cmd:     CmdMaxRotationSpeed,
}

// title : Set the presence of hull protection,
// desc : Set the presence of hull protection.,
// support : 0901;090c,
// result : The drone knows that it has a hull protection.\n Then, event [HullProtection](#1-12-2) is triggered.,
const CmdHullProtection CmdDef = 2

type Ardrone3SpeedSettingsHullProtection Command

type Ardrone3SpeedSettingsHullProtectionArguments struct {
	Present uint8
}

func (a Ardrone3SpeedSettingsHullProtection) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3SpeedSettingsHullProtectionArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Present)
	offset++

	return arg
}
func (a Ardrone3SpeedSettingsHullProtection) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SpeedSettingsHullProtection = Ardrone3SpeedSettingsHullProtection{
	Project: ProjectArdrone3,
	Class:   ClassSpeedSettings,
	Cmd:     CmdHullProtection,
}

// title : Set outdoor mode,
// desc : Set outdoor mode.,
const CmdOutdoor CmdDef = 3

type Ardrone3SpeedSettingsOutdoor Command

type Ardrone3SpeedSettingsOutdoorArguments struct {
	Outdoor uint8
}

func (a Ardrone3SpeedSettingsOutdoor) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3SpeedSettingsOutdoorArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Outdoor)
	offset++

	return arg
}
func (a Ardrone3SpeedSettingsOutdoor) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SpeedSettingsOutdoor = Ardrone3SpeedSettingsOutdoor{
	Project: ProjectArdrone3,
	Class:   ClassSpeedSettings,
	Cmd:     CmdOutdoor,
}

// title : Set max pitch/roll rotation speed,
// desc : Set max pitch/roll rotation speed.,
// support : 0901;090c,
// result : The max pitch/roll rotation speed is set.\n Then, event [MaxPitchRollRotationSpeed](#1-12-4) is triggered.,
const CmdMaxPitchRollRotationSpeed CmdDef = 4

type Ardrone3SpeedSettingsMaxPitchRollRotationSpeed Command

type Ardrone3SpeedSettingsMaxPitchRollRotationSpeedArguments struct {
	Current float32
}

func (a Ardrone3SpeedSettingsMaxPitchRollRotationSpeed) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3SpeedSettingsMaxPitchRollRotationSpeedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Current)
	offset += 4

	return arg
}
func (a Ardrone3SpeedSettingsMaxPitchRollRotationSpeed) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SpeedSettingsMaxPitchRollRotationSpeed = Ardrone3SpeedSettingsMaxPitchRollRotationSpeed{
	Project: ProjectArdrone3,
	Class:   ClassSpeedSettings,
	Cmd:     CmdMaxPitchRollRotationSpeed,
}

// Speed Settings state from product
const ClassSpeedSettingsState ClassDef = 12

// title : Max vertical speed,
// desc : Max vertical speed.,
// support : 0901;090c,
// triggered : by [SetMaxVerticalSpeed](#1-11-0).,
const CmdMaxVerticalSpeedChanged CmdDef = 0

type Ardrone3SpeedSettingsStateMaxVerticalSpeedChanged Command

type Ardrone3SpeedSettingsStateMaxVerticalSpeedChangedArguments struct {
	Current float32
	Min     float32
	Max     float32
}

func (a Ardrone3SpeedSettingsStateMaxVerticalSpeedChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3SpeedSettingsStateMaxVerticalSpeedChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Current)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Min)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Max)
	offset += 4

	return arg
}
func (a Ardrone3SpeedSettingsStateMaxVerticalSpeedChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SpeedSettingsStateMaxVerticalSpeedChanged = Ardrone3SpeedSettingsStateMaxVerticalSpeedChanged{
	Project: ProjectArdrone3,
	Class:   ClassSpeedSettingsState,
	Cmd:     CmdMaxVerticalSpeedChanged,
}

// title : Max rotation speed,
// desc : Max rotation speed.,
// support : 0901;090c,
// triggered : by [SetMaxRotationSpeed](#1-11-1).,
const CmdMaxRotationSpeedChanged CmdDef = 1

type Ardrone3SpeedSettingsStateMaxRotationSpeedChanged Command

type Ardrone3SpeedSettingsStateMaxRotationSpeedChangedArguments struct {
	Current float32
	Min     float32
	Max     float32
}

func (a Ardrone3SpeedSettingsStateMaxRotationSpeedChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3SpeedSettingsStateMaxRotationSpeedChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Current)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Min)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Max)
	offset += 4

	return arg
}
func (a Ardrone3SpeedSettingsStateMaxRotationSpeedChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SpeedSettingsStateMaxRotationSpeedChanged = Ardrone3SpeedSettingsStateMaxRotationSpeedChanged{
	Project: ProjectArdrone3,
	Class:   ClassSpeedSettingsState,
	Cmd:     CmdMaxRotationSpeedChanged,
}

// title : Presence of hull protection,
// desc : Presence of hull protection.,
// support : 0901;090c,
// triggered : by [SetHullProtectionPresence](#1-11-2).,
const CmdHullProtectionChanged CmdDef = 2

type Ardrone3SpeedSettingsStateHullProtectionChanged Command

type Ardrone3SpeedSettingsStateHullProtectionChangedArguments struct {
	Present uint8
}

func (a Ardrone3SpeedSettingsStateHullProtectionChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3SpeedSettingsStateHullProtectionChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Present)
	offset++

	return arg
}
func (a Ardrone3SpeedSettingsStateHullProtectionChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SpeedSettingsStateHullProtectionChanged = Ardrone3SpeedSettingsStateHullProtectionChanged{
	Project: ProjectArdrone3,
	Class:   ClassSpeedSettingsState,
	Cmd:     CmdHullProtectionChanged,
}

// title : Outdoor mode,
// desc : Outdoor mode.,
const CmdOutdoorChanged CmdDef = 3

type Ardrone3SpeedSettingsStateOutdoorChanged Command

type Ardrone3SpeedSettingsStateOutdoorChangedArguments struct {
	Outdoor uint8
}

func (a Ardrone3SpeedSettingsStateOutdoorChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3SpeedSettingsStateOutdoorChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Outdoor)
	offset++

	return arg
}
func (a Ardrone3SpeedSettingsStateOutdoorChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SpeedSettingsStateOutdoorChanged = Ardrone3SpeedSettingsStateOutdoorChanged{
	Project: ProjectArdrone3,
	Class:   ClassSpeedSettingsState,
	Cmd:     CmdOutdoorChanged,
}

// title : Max pitch/roll rotation speed,
// desc : Max pitch/roll rotation speed.,
// support : 0901;090c,
// triggered : by [SetMaxPitchRollRotationSpeed](#1-11-4).,
const CmdMaxPitchRollRotationSpeedChanged CmdDef = 4

type Ardrone3SpeedSettingsStateMaxPitchRollRotationSpeedChanged Command

type Ardrone3SpeedSettingsStateMaxPitchRollRotationSpeedChangedArguments struct {
	Current float32
	Min     float32
	Max     float32
}

func (a Ardrone3SpeedSettingsStateMaxPitchRollRotationSpeedChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3SpeedSettingsStateMaxPitchRollRotationSpeedChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Current)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Min)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Max)
	offset += 4

	return arg
}
func (a Ardrone3SpeedSettingsStateMaxPitchRollRotationSpeedChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SpeedSettingsStateMaxPitchRollRotationSpeedChanged = Ardrone3SpeedSettingsStateMaxPitchRollRotationSpeedChanged{
	Project: ProjectArdrone3,
	Class:   ClassSpeedSettingsState,
	Cmd:     CmdMaxPitchRollRotationSpeedChanged,
}

// Network settings commands
const ClassNetworkSettings ClassDef = 9

// title : Select Wifi,
// desc : Select or auto-select channel of choosen band.,
// support : 0901;090c;090e,
// result : The wifi channel changes according to given parameters. Watch out, a disconnection might appear.\n Then, event [WifiSelection](#1-10-0) is triggered.,
const CmdWifiSelection CmdDef = 0

type Ardrone3NetworkSettingsWifiSelection Command

type Ardrone3NetworkSettingsWifiSelectionArguments struct {
	TypeX   uint32
	Band    uint32
	Channel uint8
}

func (a Ardrone3NetworkSettingsWifiSelection) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3NetworkSettingsWifiSelectionArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TypeX)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Band)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Channel)
	offset++

	return arg
}
func (a Ardrone3NetworkSettingsWifiSelection) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var NetworkSettingsWifiSelection = Ardrone3NetworkSettingsWifiSelection{
	Project: ProjectArdrone3,
	Class:   ClassNetworkSettings,
	Cmd:     CmdWifiSelection,
}

// title : Set wifi security type,
// desc : Set wifi security type.\n The security will be changed on the next restart,
// support : 0901;090c;090e,
// result : The wifi security is set (but not applied until next restart).\n Then, event [WifiSecurityType](#1-10-2) is triggered.,
const CmdWifiSecurity CmdDef = 1

type Ardrone3NetworkSettingswifiSecurity Command

type Ardrone3NetworkSettingswifiSecurityArguments struct {
	TypeX   uint32
	Key     string
	KeyType uint32
}

func (a Ardrone3NetworkSettingswifiSecurity) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := Ardrone3NetworkSettingswifiSecurityArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TypeX)
	offset += 4

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Key = string(b[offset : offset+stringEnd])
	offset += stringEnd
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.KeyType)
	offset += 4

	return arg
}
func (a Ardrone3NetworkSettingswifiSecurity) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var NetworkSettingswifiSecurity = Ardrone3NetworkSettingswifiSecurity{
	Project: ProjectArdrone3,
	Class:   ClassNetworkSettings,
	Cmd:     CmdWifiSecurity,
}

// Network settings state from product
const ClassNetworkSettingsState ClassDef = 10

// title : Wifi selection,
// desc : Wifi selection.,
// support : 0901;090c;090e,
// triggered : by [SelectWifi](#1-9-0).,
const CmdWifiSelectionChanged CmdDef = 0

type Ardrone3NetworkSettingsStateWifiSelectionChanged Command

type Ardrone3NetworkSettingsStateWifiSelectionChangedArguments struct {
	TypeX   uint32
	Band    uint32
	Channel uint8
}

func (a Ardrone3NetworkSettingsStateWifiSelectionChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3NetworkSettingsStateWifiSelectionChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TypeX)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Band)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Channel)
	offset++

	return arg
}
func (a Ardrone3NetworkSettingsStateWifiSelectionChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var NetworkSettingsStateWifiSelectionChanged = Ardrone3NetworkSettingsStateWifiSelectionChanged{
	Project: ProjectArdrone3,
	Class:   ClassNetworkSettingsState,
	Cmd:     CmdWifiSelectionChanged,
}

// title : Wifi security type,
// desc : Wifi security type.,
const CmdWifiSecurityChanged CmdDef = 1

type Ardrone3NetworkSettingsStatewifiSecurityChanged Command

type Ardrone3NetworkSettingsStatewifiSecurityChangedArguments struct {
	TypeX uint32
}

func (a Ardrone3NetworkSettingsStatewifiSecurityChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3NetworkSettingsStatewifiSecurityChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TypeX)
	offset += 4

	return arg
}
func (a Ardrone3NetworkSettingsStatewifiSecurityChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var NetworkSettingsStatewifiSecurityChanged = Ardrone3NetworkSettingsStatewifiSecurityChanged{
	Project: ProjectArdrone3,
	Class:   ClassNetworkSettingsState,
	Cmd:     CmdWifiSecurityChanged,
}

// title : Wifi security type,
// desc : Wifi security type.,
// support : 0901;090c;090e,
// triggered : by [SetWifiSecurityType](#1-9-1).,
const CmdWifiSecurityDUPLICATE CmdDef = 2

type Ardrone3NetworkSettingsStatewifiSecurity Command

type Ardrone3NetworkSettingsStatewifiSecurityArguments struct {
	TypeX   uint32
	Key     string
	KeyType uint32
}

func (a Ardrone3NetworkSettingsStatewifiSecurity) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := Ardrone3NetworkSettingsStatewifiSecurityArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TypeX)
	offset += 4

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Key = string(b[offset : offset+stringEnd])
	offset += stringEnd
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.KeyType)
	offset += 4

	return arg
}
func (a Ardrone3NetworkSettingsStatewifiSecurity) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var NetworkSettingsStatewifiSecurity = Ardrone3NetworkSettingsStatewifiSecurity{
	Project: ProjectArdrone3,
	Class:   ClassNetworkSettingsState,
	Cmd:     CmdWifiSecurity,
}

// Settings state from product
const ClassSettingsState ClassDef = 16

// title : Motor version,
// desc : Motor version.,
const CmdProductMotorVersionListChanged CmdDef = 0

type Ardrone3SettingsStateProductMotorVersionListChanged Command

type Ardrone3SettingsStateProductMotorVersionListChangedArguments struct {
	Motornumber uint8
	TypeX       string
	Software    string
	Hardware    string
}

func (a Ardrone3SettingsStateProductMotorVersionListChanged) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := Ardrone3SettingsStateProductMotorVersionListChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Motornumber)
	offset++

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.TypeX = string(b[offset : offset+stringEnd])
	offset += stringEnd

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Software = string(b[offset : offset+stringEnd])
	offset += stringEnd

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Hardware = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a Ardrone3SettingsStateProductMotorVersionListChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsStateProductMotorVersionListChanged = Ardrone3SettingsStateProductMotorVersionListChanged{
	Project: ProjectArdrone3,
	Class:   ClassSettingsState,
	Cmd:     CmdProductMotorVersionListChanged,
}

// title : GPS version,
// desc : GPS version.,
// support : 0901;090c;090e,
// triggered : at connection.,
const CmdProductGPSVersionChanged CmdDef = 1

type Ardrone3SettingsStateProductGPSVersionChanged Command

type Ardrone3SettingsStateProductGPSVersionChangedArguments struct {
	Software string
	Hardware string
}

func (a Ardrone3SettingsStateProductGPSVersionChanged) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := Ardrone3SettingsStateProductGPSVersionChangedArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Software = string(b[offset : offset+stringEnd])
	offset += stringEnd

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Hardware = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a Ardrone3SettingsStateProductGPSVersionChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsStateProductGPSVersionChanged = Ardrone3SettingsStateProductGPSVersionChanged{
	Project: ProjectArdrone3,
	Class:   ClassSettingsState,
	Cmd:     CmdProductGPSVersionChanged,
}

// title : Motor error,
// desc : Motor error.\n This event is sent back to *noError* as soon as the motor error disappear. To get the last motor error, see [LastMotorError](#1-16-5),
// support : 0901;090c;090e,
// triggered : when a motor error occurs.,
const CmdMotorErrorStateChanged CmdDef = 2

type Ardrone3SettingsStateMotorErrorStateChanged Command

type Ardrone3SettingsStateMotorErrorStateChangedArguments struct {
	MotorIds   uint8
	MotorError uint32
}

func (a Ardrone3SettingsStateMotorErrorStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3SettingsStateMotorErrorStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.MotorIds)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.MotorError)
	offset += 4

	return arg
}
func (a Ardrone3SettingsStateMotorErrorStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsStateMotorErrorStateChanged = Ardrone3SettingsStateMotorErrorStateChanged{
	Project: ProjectArdrone3,
	Class:   ClassSettingsState,
	Cmd:     CmdMotorErrorStateChanged,
}

// title : Motor version,
// desc : Motor version.,
const CmdMotorSoftwareVersionChanged CmdDef = 3

type Ardrone3SettingsStateMotorSoftwareVersionChanged Command

type Ardrone3SettingsStateMotorSoftwareVersionChangedArguments struct {
	Version string
}

func (a Ardrone3SettingsStateMotorSoftwareVersionChanged) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := Ardrone3SettingsStateMotorSoftwareVersionChangedArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Version = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a Ardrone3SettingsStateMotorSoftwareVersionChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsStateMotorSoftwareVersionChanged = Ardrone3SettingsStateMotorSoftwareVersionChanged{
	Project: ProjectArdrone3,
	Class:   ClassSettingsState,
	Cmd:     CmdMotorSoftwareVersionChanged,
}

// title : Motor flight status,
// desc : Motor flight status.,
// support : 0901;090c;090e,
// triggered : at connection.,
const CmdMotorFlightsStatusChanged CmdDef = 4

type Ardrone3SettingsStateMotorFlightsStatusChanged Command

type Ardrone3SettingsStateMotorFlightsStatusChangedArguments struct {
	NbFlights           uint16
	LastFlightDuration  uint16
	TotalFlightDuration uint32
}

func (a Ardrone3SettingsStateMotorFlightsStatusChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3SettingsStateMotorFlightsStatusChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.NbFlights)
	offset += 2
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.LastFlightDuration)
	offset += 2
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TotalFlightDuration)
	offset += 4

	return arg
}
func (a Ardrone3SettingsStateMotorFlightsStatusChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsStateMotorFlightsStatusChanged = Ardrone3SettingsStateMotorFlightsStatusChanged{
	Project: ProjectArdrone3,
	Class:   ClassSettingsState,
	Cmd:     CmdMotorFlightsStatusChanged,
}

// title : Last motor error,
// desc : Last motor error.\n This is a reminder of the last error. To know if a motor error is currently happening, see [MotorError](#1-16-2).,
// support : 0901;090c;090e,
// triggered : at connection and when an error occurs.,
const CmdMotorErrorLastErrorChanged CmdDef = 5

type Ardrone3SettingsStateMotorErrorLastErrorChanged Command

type Ardrone3SettingsStateMotorErrorLastErrorChangedArguments struct {
	MotorError uint32
}

func (a Ardrone3SettingsStateMotorErrorLastErrorChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3SettingsStateMotorErrorLastErrorChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.MotorError)
	offset += 4

	return arg
}
func (a Ardrone3SettingsStateMotorErrorLastErrorChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsStateMotorErrorLastErrorChanged = Ardrone3SettingsStateMotorErrorLastErrorChanged{
	Project: ProjectArdrone3,
	Class:   ClassSettingsState,
	Cmd:     CmdMotorErrorLastErrorChanged,
}

// title : P7ID,
// desc : P7ID.,
const CmdP7ID CmdDef = 6

type Ardrone3SettingsStateP7ID Command

type Ardrone3SettingsStateP7IDArguments struct {
	SerialID string
}

func (a Ardrone3SettingsStateP7ID) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := Ardrone3SettingsStateP7IDArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.SerialID = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a Ardrone3SettingsStateP7ID) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsStateP7ID = Ardrone3SettingsStateP7ID{
	Project: ProjectArdrone3,
	Class:   ClassSettingsState,
	Cmd:     CmdP7ID,
}

const CmdCPUID CmdDef = 7

type Ardrone3SettingsStateCPUID Command

type Ardrone3SettingsStateCPUIDArguments struct {
	Id string
}

func (a Ardrone3SettingsStateCPUID) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := Ardrone3SettingsStateCPUIDArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Id = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a Ardrone3SettingsStateCPUID) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsStateCPUID = Ardrone3SettingsStateCPUID{
	Project: ProjectArdrone3,
	Class:   ClassSettingsState,
	Cmd:     CmdCPUID,
}

// Photo settings chosen by the user
const ClassPictureSettings ClassDef = 19

// title : Set picture format,
// desc : Set picture format.\n Please note that the time required to take the picture is highly related to this format.\n Also, please note that if your picture format is different from snapshot, picture taking will stop video recording (it will restart after the picture has been taken).,
// support : 0901;090c;090e,
// result : The picture format is set.\n Then, event [PictureFormat](#1-20-0) is triggered.,
const CmdPictureFormatSelection CmdDef = 0

type Ardrone3PictureSettingsPictureFormatSelection Command

type Ardrone3PictureSettingsPictureFormatSelectionArguments struct {
	TypeX uint32
}

func (a Ardrone3PictureSettingsPictureFormatSelection) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsPictureFormatSelectionArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TypeX)
	offset += 4

	return arg
}
func (a Ardrone3PictureSettingsPictureFormatSelection) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsPictureFormatSelection = Ardrone3PictureSettingsPictureFormatSelection{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettings,
	Cmd:     CmdPictureFormatSelection,
}

// title : Set White Balance mode,
// desc : Set White Balance mode.,
// support : 0901;090c;090e,
// result : The white balance mode is set.\n Then, event [WhiteBalanceMode](#1-20-1) is triggered.,
const CmdAutoWhiteBalanceSelection CmdDef = 1

type Ardrone3PictureSettingsAutoWhiteBalanceSelection Command

type Ardrone3PictureSettingsAutoWhiteBalanceSelectionArguments struct {
	TypeX uint32
}

func (a Ardrone3PictureSettingsAutoWhiteBalanceSelection) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsAutoWhiteBalanceSelectionArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TypeX)
	offset += 4

	return arg
}
func (a Ardrone3PictureSettingsAutoWhiteBalanceSelection) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsAutoWhiteBalanceSelection = Ardrone3PictureSettingsAutoWhiteBalanceSelection{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettings,
	Cmd:     CmdAutoWhiteBalanceSelection,
}

// title : Set image exposure,
// desc : Set image exposure.,
// support : 0901;090c;090e,
// result : The exposure is set.\n Then, event [ImageExposure](#1-20-2) is triggered.,
const CmdExpositionSelection CmdDef = 2

type Ardrone3PictureSettingsExpositionSelection Command

type Ardrone3PictureSettingsExpositionSelectionArguments struct {
	Value float32
}

func (a Ardrone3PictureSettingsExpositionSelection) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsExpositionSelectionArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4

	return arg
}
func (a Ardrone3PictureSettingsExpositionSelection) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsExpositionSelection = Ardrone3PictureSettingsExpositionSelection{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettings,
	Cmd:     CmdExpositionSelection,
}

// title : Set image saturation,
// desc : Set image saturation.,
// support : 0901;090c;090e,
// result : The saturation is set.\n Then, event [ImageSaturation](#1-20-3) is triggered.,
const CmdSaturationSelection CmdDef = 3

type Ardrone3PictureSettingsSaturationSelection Command

type Ardrone3PictureSettingsSaturationSelectionArguments struct {
	Value float32
}

func (a Ardrone3PictureSettingsSaturationSelection) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsSaturationSelectionArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4

	return arg
}
func (a Ardrone3PictureSettingsSaturationSelection) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsSaturationSelection = Ardrone3PictureSettingsSaturationSelection{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettings,
	Cmd:     CmdSaturationSelection,
}

// title : Set timelapse mode,
// desc : Set timelapse mode.\n If timelapse mode is set, instead of taking a video, the drone will take picture regularly.\n Watch out, this command only configure the timelapse mode. Once it is configured, you can start/stop the timelapse with the [RecordVideo](#1-7-3) command.,
// support : 0901;090c;090e,
// result : The timelapse mode is set (but not started).\n Then, event [TimelapseMode](#1-20-4) is triggered.,
const CmdTimelapseSelection CmdDef = 4

type Ardrone3PictureSettingsTimelapseSelection Command

type Ardrone3PictureSettingsTimelapseSelectionArguments struct {
	Enabled  uint8
	Interval float32
}

func (a Ardrone3PictureSettingsTimelapseSelection) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsTimelapseSelectionArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Enabled)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Interval)
	offset += 4

	return arg
}
func (a Ardrone3PictureSettingsTimelapseSelection) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsTimelapseSelection = Ardrone3PictureSettingsTimelapseSelection{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettings,
	Cmd:     CmdTimelapseSelection,
}

// title : Set video autorecord mode,
// desc : Set video autorecord mode.\n If autorecord is set, video record will be automatically started when the drone takes off and stopped slightly after landing.,
// support : 0901;090c;090e,
// result : The autorecord mode is set.\n Then, event [AutorecordMode](#1-20-5) is triggered.,
const CmdVideoAutorecordSelection CmdDef = 5

type Ardrone3PictureSettingsVideoAutorecordSelection Command

type Ardrone3PictureSettingsVideoAutorecordSelectionArguments struct {
	Enabled       uint8
	Massstorageid uint8
}

func (a Ardrone3PictureSettingsVideoAutorecordSelection) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsVideoAutorecordSelectionArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Enabled)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Massstorageid)
	offset++

	return arg
}
func (a Ardrone3PictureSettingsVideoAutorecordSelection) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsVideoAutorecordSelection = Ardrone3PictureSettingsVideoAutorecordSelection{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettings,
	Cmd:     CmdVideoAutorecordSelection,
}

// title : Set video stabilization mode,
// desc : Set video stabilization mode.,
// support : 0901:3.4.0;090c:3.4.0;090e,
// result : The video stabilization mode is set.\n Then, event [VideoStabilizationMode](#1-20-6) is triggered.,
const CmdVideoStabilizationMode CmdDef = 6

type Ardrone3PictureSettingsVideoStabilizationMode Command

type Ardrone3PictureSettingsVideoStabilizationModeArguments struct {
	Mode uint32
}

func (a Ardrone3PictureSettingsVideoStabilizationMode) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsVideoStabilizationModeArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Mode)
	offset += 4

	return arg
}
func (a Ardrone3PictureSettingsVideoStabilizationMode) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsVideoStabilizationMode = Ardrone3PictureSettingsVideoStabilizationMode{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettings,
	Cmd:     CmdVideoStabilizationMode,
}

// title : Set video recording mode,
// desc : Set video recording mode.,
// support : 0901:3.4.0;090c:3.4.0;090e,
// result : The video recording mode is set.\n Then, event [VideoRecordingMode](#1-20-7) is triggered.,
const CmdVideoRecordingMode CmdDef = 7

type Ardrone3PictureSettingsVideoRecordingMode Command

type Ardrone3PictureSettingsVideoRecordingModeArguments struct {
	Mode uint32
}

func (a Ardrone3PictureSettingsVideoRecordingMode) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsVideoRecordingModeArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Mode)
	offset += 4

	return arg
}
func (a Ardrone3PictureSettingsVideoRecordingMode) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsVideoRecordingMode = Ardrone3PictureSettingsVideoRecordingMode{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettings,
	Cmd:     CmdVideoRecordingMode,
}

// title : Set video framerate,
// desc : Set video framerate.,
// support : 0901:3.4.0;090c:3.4.0;090e,
// result : The video framerate is set.\n Then, event [VideoFramerate](#1-20-8) is triggered.,
const CmdVideoFramerate CmdDef = 8

type Ardrone3PictureSettingsVideoFramerate Command

type Ardrone3PictureSettingsVideoFramerateArguments struct {
	Framerate uint32
}

func (a Ardrone3PictureSettingsVideoFramerate) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsVideoFramerateArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Framerate)
	offset += 4

	return arg
}
func (a Ardrone3PictureSettingsVideoFramerate) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsVideoFramerate = Ardrone3PictureSettingsVideoFramerate{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettings,
	Cmd:     CmdVideoFramerate,
}

// title : Set video resolutions,
// desc : Set video streaming and recording resolutions.,
// support : 0901:3.4.0;090c:3.4.0;090e,
// result : The video resolutions is set.\n Then, event [VideoResolutions](#1-20-9) is triggered.,
const CmdVideoResolutions CmdDef = 9

type Ardrone3PictureSettingsVideoResolutions Command

type Ardrone3PictureSettingsVideoResolutionsArguments struct {
	TypeX uint32
}

func (a Ardrone3PictureSettingsVideoResolutions) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsVideoResolutionsArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TypeX)
	offset += 4

	return arg
}
func (a Ardrone3PictureSettingsVideoResolutions) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsVideoResolutions = Ardrone3PictureSettingsVideoResolutions{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettings,
	Cmd:     CmdVideoResolutions,
}

// Photo settings state from product
const ClassPictureSettingsState ClassDef = 20

// title : Picture format,
// desc : Picture format.,
// support : 0901;090c;090e,
// triggered : by [SetPictureFormat](#1-19-0).,
const CmdPictureFormatChanged CmdDef = 0

type Ardrone3PictureSettingsStatePictureFormatChanged Command

type Ardrone3PictureSettingsStatePictureFormatChangedArguments struct {
	TypeX uint32
}

func (a Ardrone3PictureSettingsStatePictureFormatChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsStatePictureFormatChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TypeX)
	offset += 4

	return arg
}
func (a Ardrone3PictureSettingsStatePictureFormatChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsStatePictureFormatChanged = Ardrone3PictureSettingsStatePictureFormatChanged{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettingsState,
	Cmd:     CmdPictureFormatChanged,
}

// title : White balance mode,
// desc : White balance mode.,
// support : 0901;090c;090e,
// triggered : by [SetWhiteBalanceMode](#1-19-1).,
const CmdAutoWhiteBalanceChanged CmdDef = 1

type Ardrone3PictureSettingsStateAutoWhiteBalanceChanged Command

type Ardrone3PictureSettingsStateAutoWhiteBalanceChangedArguments struct {
	TypeX uint32
}

func (a Ardrone3PictureSettingsStateAutoWhiteBalanceChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsStateAutoWhiteBalanceChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TypeX)
	offset += 4

	return arg
}
func (a Ardrone3PictureSettingsStateAutoWhiteBalanceChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsStateAutoWhiteBalanceChanged = Ardrone3PictureSettingsStateAutoWhiteBalanceChanged{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettingsState,
	Cmd:     CmdAutoWhiteBalanceChanged,
}

// title : Image exposure,
// desc : Image exposure.,
// support : 0901;090c;090e,
// triggered : by [SetImageExposure](#1-19-2).,
const CmdExpositionChanged CmdDef = 2

type Ardrone3PictureSettingsStateExpositionChanged Command

type Ardrone3PictureSettingsStateExpositionChangedArguments struct {
	Value float32
	Min   float32
	Max   float32
}

func (a Ardrone3PictureSettingsStateExpositionChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsStateExpositionChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Min)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Max)
	offset += 4

	return arg
}
func (a Ardrone3PictureSettingsStateExpositionChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsStateExpositionChanged = Ardrone3PictureSettingsStateExpositionChanged{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettingsState,
	Cmd:     CmdExpositionChanged,
}

// title : Image saturation,
// desc : Image saturation.,
// support : 0901;090c;090e,
// triggered : by [SetImageSaturation](#1-19-3).,
const CmdSaturationChanged CmdDef = 3

type Ardrone3PictureSettingsStateSaturationChanged Command

type Ardrone3PictureSettingsStateSaturationChangedArguments struct {
	Value float32
	Min   float32
	Max   float32
}

func (a Ardrone3PictureSettingsStateSaturationChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsStateSaturationChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Min)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Max)
	offset += 4

	return arg
}
func (a Ardrone3PictureSettingsStateSaturationChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsStateSaturationChanged = Ardrone3PictureSettingsStateSaturationChanged{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettingsState,
	Cmd:     CmdSaturationChanged,
}

// title : Timelapse mode,
// desc : Timelapse mode.,
// support : 0901;090c;090e,
// triggered : by [SetTimelapseMode](#1-19-4).,
const CmdTimelapseChanged CmdDef = 4

type Ardrone3PictureSettingsStateTimelapseChanged Command

type Ardrone3PictureSettingsStateTimelapseChangedArguments struct {
	Enabled     uint8
	Interval    float32
	MinInterval float32
	MaxInterval float32
}

func (a Ardrone3PictureSettingsStateTimelapseChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsStateTimelapseChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Enabled)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Interval)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.MinInterval)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.MaxInterval)
	offset += 4

	return arg
}
func (a Ardrone3PictureSettingsStateTimelapseChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsStateTimelapseChanged = Ardrone3PictureSettingsStateTimelapseChanged{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettingsState,
	Cmd:     CmdTimelapseChanged,
}

// title : Video Autorecord mode,
// desc : Video Autorecord mode.,
// support : 0901;090c;090e,
// triggered : by [SetVideoAutorecordMode](#1-19-5).,
const CmdVideoAutorecordChanged CmdDef = 5

type Ardrone3PictureSettingsStateVideoAutorecordChanged Command

type Ardrone3PictureSettingsStateVideoAutorecordChangedArguments struct {
	Enabled       uint8
	Massstorageid uint8
}

func (a Ardrone3PictureSettingsStateVideoAutorecordChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsStateVideoAutorecordChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Enabled)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Massstorageid)
	offset++

	return arg
}
func (a Ardrone3PictureSettingsStateVideoAutorecordChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsStateVideoAutorecordChanged = Ardrone3PictureSettingsStateVideoAutorecordChanged{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettingsState,
	Cmd:     CmdVideoAutorecordChanged,
}

// title : Video stabilization mode,
// desc : Video stabilization mode.,
// support : 0901:3.4.0;090c:3.4.0;090e,
// triggered : by [SetVideoStabilizationMode](#1-19-6).,
const CmdVideoStabilizationModeChanged CmdDef = 6

type Ardrone3PictureSettingsStateVideoStabilizationModeChanged Command

type Ardrone3PictureSettingsStateVideoStabilizationModeChangedArguments struct {
	Mode uint32
}

func (a Ardrone3PictureSettingsStateVideoStabilizationModeChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsStateVideoStabilizationModeChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Mode)
	offset += 4

	return arg
}
func (a Ardrone3PictureSettingsStateVideoStabilizationModeChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsStateVideoStabilizationModeChanged = Ardrone3PictureSettingsStateVideoStabilizationModeChanged{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettingsState,
	Cmd:     CmdVideoStabilizationModeChanged,
}

// title : Video recording mode,
// desc : Video recording mode.,
// support : 0901:3.4.0;090c:3.4.0;090e,
// triggered : by [SetVideoRecordingMode](#1-19-7).,
const CmdVideoRecordingModeChanged CmdDef = 7

type Ardrone3PictureSettingsStateVideoRecordingModeChanged Command

type Ardrone3PictureSettingsStateVideoRecordingModeChangedArguments struct {
	Mode uint32
}

func (a Ardrone3PictureSettingsStateVideoRecordingModeChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsStateVideoRecordingModeChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Mode)
	offset += 4

	return arg
}
func (a Ardrone3PictureSettingsStateVideoRecordingModeChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsStateVideoRecordingModeChanged = Ardrone3PictureSettingsStateVideoRecordingModeChanged{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettingsState,
	Cmd:     CmdVideoRecordingModeChanged,
}

// title : Video framerate,
// desc : Video framerate.,
// support : 0901:3.4.0;090c:3.4.0;090e,
// triggered : by [SetVideoFramerateMode](#1-19-8).,
const CmdVideoFramerateChanged CmdDef = 8

type Ardrone3PictureSettingsStateVideoFramerateChanged Command

type Ardrone3PictureSettingsStateVideoFramerateChangedArguments struct {
	Framerate uint32
}

func (a Ardrone3PictureSettingsStateVideoFramerateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsStateVideoFramerateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Framerate)
	offset += 4

	return arg
}
func (a Ardrone3PictureSettingsStateVideoFramerateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsStateVideoFramerateChanged = Ardrone3PictureSettingsStateVideoFramerateChanged{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettingsState,
	Cmd:     CmdVideoFramerateChanged,
}

// title : Video resolutions,
// desc : Video resolutions.\n This event informs about the recording AND streaming resolutions.,
// support : 0901:3.4.0;090c:3.4.0;090e,
// triggered : by [SetVideResolutions](#1-19-9).,
const CmdVideoResolutionsChanged CmdDef = 9

type Ardrone3PictureSettingsStateVideoResolutionsChanged Command

type Ardrone3PictureSettingsStateVideoResolutionsChangedArguments struct {
	TypeX uint32
}

func (a Ardrone3PictureSettingsStateVideoResolutionsChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PictureSettingsStateVideoResolutionsChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TypeX)
	offset += 4

	return arg
}
func (a Ardrone3PictureSettingsStateVideoResolutionsChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PictureSettingsStateVideoResolutionsChanged = Ardrone3PictureSettingsStateVideoResolutionsChanged{
	Project: ProjectArdrone3,
	Class:   ClassPictureSettingsState,
	Cmd:     CmdVideoResolutionsChanged,
}

// Control media streaming behavior.
const ClassMediaStreaming ClassDef = 21

// title : Enable/disable video streaming,
// desc : Enable/disable video streaming.,
// support : 0901;090c;090e,
// result : The video stream is started or stopped.\n Then, event [VideoStreamState](#1-22-0) is triggered.,
const CmdVideoEnable CmdDef = 0

type Ardrone3MediaStreamingVideoEnable Command

type Ardrone3MediaStreamingVideoEnableArguments struct {
	Enable uint8
}

func (a Ardrone3MediaStreamingVideoEnable) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3MediaStreamingVideoEnableArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Enable)
	offset++

	return arg
}
func (a Ardrone3MediaStreamingVideoEnable) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MediaStreamingVideoEnable = Ardrone3MediaStreamingVideoEnable{
	Project: ProjectArdrone3,
	Class:   ClassMediaStreaming,
	Cmd:     CmdVideoEnable,
}

// title : Set the stream mode,
// desc : Set the stream mode.,
// support : 0901;090c;090e,
// result : The stream mode is set.\n Then, event [VideoStreamMode](#1-22-1) is triggered.,
const CmdVideoStreamMode CmdDef = 1

type Ardrone3MediaStreamingVideoStreamMode Command

type Ardrone3MediaStreamingVideoStreamModeArguments struct {
	Mode uint32
}

func (a Ardrone3MediaStreamingVideoStreamMode) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3MediaStreamingVideoStreamModeArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Mode)
	offset += 4

	return arg
}
func (a Ardrone3MediaStreamingVideoStreamMode) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MediaStreamingVideoStreamMode = Ardrone3MediaStreamingVideoStreamMode{
	Project: ProjectArdrone3,
	Class:   ClassMediaStreaming,
	Cmd:     CmdVideoStreamMode,
}

// Media streaming status.
const ClassMediaStreamingState ClassDef = 22

// title : Video stream state,
// desc : Video stream state.,
// support : 0901;090c;090e,
// triggered : by [EnableOrDisableVideoStream](#1-21-0).,
const CmdVideoEnableChanged CmdDef = 0

type Ardrone3MediaStreamingStateVideoEnableChanged Command

type Ardrone3MediaStreamingStateVideoEnableChangedArguments struct {
	Enabled uint32
}

func (a Ardrone3MediaStreamingStateVideoEnableChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3MediaStreamingStateVideoEnableChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Enabled)
	offset += 4

	return arg
}
func (a Ardrone3MediaStreamingStateVideoEnableChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MediaStreamingStateVideoEnableChanged = Ardrone3MediaStreamingStateVideoEnableChanged{
	Project: ProjectArdrone3,
	Class:   ClassMediaStreamingState,
	Cmd:     CmdVideoEnableChanged,
}

const CmdVideoStreamModeChanged CmdDef = 1

type Ardrone3MediaStreamingStateVideoStreamModeChanged Command

type Ardrone3MediaStreamingStateVideoStreamModeChangedArguments struct {
	Mode uint32
}

func (a Ardrone3MediaStreamingStateVideoStreamModeChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3MediaStreamingStateVideoStreamModeChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Mode)
	offset += 4

	return arg
}
func (a Ardrone3MediaStreamingStateVideoStreamModeChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MediaStreamingStateVideoStreamModeChanged = Ardrone3MediaStreamingStateVideoStreamModeChanged{
	Project: ProjectArdrone3,
	Class:   ClassMediaStreamingState,
	Cmd:     CmdVideoStreamModeChanged,
}

// GPS settings
const ClassGPSSettings ClassDef = 23

// title : Set home position,
// desc : Set home position.,
const CmdSetHome CmdDef = 0

type Ardrone3GPSSettingsSetHome Command

type Ardrone3GPSSettingsSetHomeArguments struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

func (a Ardrone3GPSSettingsSetHome) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3GPSSettingsSetHomeArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Latitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Longitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Altitude)
	offset += 8

	return arg
}
func (a Ardrone3GPSSettingsSetHome) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var GPSSettingsSetHome = Ardrone3GPSSettingsSetHome{
	Project: ProjectArdrone3,
	Class:   ClassGPSSettings,
	Cmd:     CmdSetHome,
}

// title : Reset home position,
// desc : Reset home position.,
// support : 0901;090c,
// result : The home position is reset.\n Then, event [HomeLocationReset](#1-24-1) is triggered.,
const CmdResetHome CmdDef = 1

type Ardrone3GPSSettingsResetHome Command

type Ardrone3GPSSettingsResetHomeArguments struct {
}

func (a Ardrone3GPSSettingsResetHome) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3GPSSettingsResetHomeArguments{}
	// No arguments to decode here !!

	return arg
}
func (a Ardrone3GPSSettingsResetHome) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var GPSSettingsResetHome = Ardrone3GPSSettingsResetHome{
	Project: ProjectArdrone3,
	Class:   ClassGPSSettings,
	Cmd:     CmdResetHome,
}

// title : Set controller gps location,
// desc : Set controller gps location.\n The user location might be used in case of return home, according to the home type and the accuracy of the given position. You can get the current home type with the event [HomeType](#1-24-4).,
// support : 0901;090c;090e,
// result : The controller position is known by the drone.\n Then, event [HomeLocation](#1-24-2) is triggered.,
const CmdSendControllerGPS CmdDef = 2

type Ardrone3GPSSettingsSendControllerGPS Command

type Ardrone3GPSSettingsSendControllerGPSArguments struct {
	Latitude           float64
	Longitude          float64
	Altitude           float64
	HorizontalAccuracy float64
	VerticalAccuracy   float64
}

func (a Ardrone3GPSSettingsSendControllerGPS) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3GPSSettingsSendControllerGPSArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Latitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Longitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Altitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.HorizontalAccuracy)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.VerticalAccuracy)
	offset += 8

	return arg
}
func (a Ardrone3GPSSettingsSendControllerGPS) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var GPSSettingsSendControllerGPS = Ardrone3GPSSettingsSendControllerGPS{
	Project: ProjectArdrone3,
	Class:   ClassGPSSettings,
	Cmd:     CmdSendControllerGPS,
}

// title : Set the preferred home type,
// desc : Set the preferred home type.\n Please note that this is only a preference. The actual type chosen is given by the event [HomeType](#1-31-2).\n You can get the currently available types with the event [HomeTypeAvailability](#1-31-1).,
// support : 0901;090c;090e,
// result : The user choice is known by the drone.\n Then, event [PreferredHomeType](#1-24-4) is triggered.,
const CmdHomeType CmdDef = 3

type Ardrone3GPSSettingsHomeType Command

type Ardrone3GPSSettingsHomeTypeArguments struct {
	TypeX uint32
}

func (a Ardrone3GPSSettingsHomeType) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3GPSSettingsHomeTypeArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TypeX)
	offset += 4

	return arg
}
func (a Ardrone3GPSSettingsHomeType) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var GPSSettingsHomeType = Ardrone3GPSSettingsHomeType{
	Project: ProjectArdrone3,
	Class:   ClassGPSSettings,
	Cmd:     CmdHomeType,
}

// title : Set the return home delay,
// desc : Set the delay after which the drone will automatically try to return home after a disconnection.,
// support : 0901;090c;090e,
// result : The delay of the return home is set.\n Then, event [ReturnHomeDelay](#1-24-5) is triggered.,
const CmdReturnHomeDelay CmdDef = 4

type Ardrone3GPSSettingsReturnHomeDelay Command

type Ardrone3GPSSettingsReturnHomeDelayArguments struct {
	Delay uint16
}

func (a Ardrone3GPSSettingsReturnHomeDelay) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3GPSSettingsReturnHomeDelayArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.Delay)
	offset += 2

	return arg
}
func (a Ardrone3GPSSettingsReturnHomeDelay) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var GPSSettingsReturnHomeDelay = Ardrone3GPSSettingsReturnHomeDelay{
	Project: ProjectArdrone3,
	Class:   ClassGPSSettings,
	Cmd:     CmdReturnHomeDelay,
}

// title : Set the return home min altitude,
// desc : Set the return home minimum altitude. If the drone is below this altitude when starting its return home, it will first reach the minimum altitude. If it is higher than this minimum altitude, it will operate its return home at its actual altitude.,
// support : ,
// result : The minimum altitude for the return home is set.\n Then, event [ReturnHomeMinAltitude](#1-24-7) is triggered.,
const CmdReturnHomeMinAltitude CmdDef = 5

type Ardrone3GPSSettingsReturnHomeMinAltitude Command

type Ardrone3GPSSettingsReturnHomeMinAltitudeArguments struct {
	Value float32
}

func (a Ardrone3GPSSettingsReturnHomeMinAltitude) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3GPSSettingsReturnHomeMinAltitudeArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4

	return arg
}
func (a Ardrone3GPSSettingsReturnHomeMinAltitude) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var GPSSettingsReturnHomeMinAltitude = Ardrone3GPSSettingsReturnHomeMinAltitude{
	Project: ProjectArdrone3,
	Class:   ClassGPSSettings,
	Cmd:     CmdReturnHomeMinAltitude,
}

// GPS settings state
const ClassGPSSettingsState ClassDef = 24

// title : Home location,
// desc : Home location.,
// support : 0901;090c;090e,
// triggered : when [HomeType](#1-31-2) changes. Or by [SetHomeLocation](#1-23-2) when [HomeType](#1-31-2) is Pilot. Or regularly after [SetControllerGPS](#140-1) when [HomeType](#1-31-2) is FollowMeTarget. Or at take off [HomeType](#1-31-2) is Takeoff. Or when the first fix occurs and the [HomeType](#1-31-2) is FirstFix.,
const CmdHomeChanged CmdDef = 0

type Ardrone3GPSSettingsStateHomeChanged Command

type Ardrone3GPSSettingsStateHomeChangedArguments struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

func (a Ardrone3GPSSettingsStateHomeChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3GPSSettingsStateHomeChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Latitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Longitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Altitude)
	offset += 8

	return arg
}
func (a Ardrone3GPSSettingsStateHomeChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var GPSSettingsStateHomeChanged = Ardrone3GPSSettingsStateHomeChanged{
	Project: ProjectArdrone3,
	Class:   ClassGPSSettingsState,
	Cmd:     CmdHomeChanged,
}

// title : Home location has been reset,
// desc : Home location has been reset.,
// support : 0901;090c,
// triggered : by [ResetHomeLocation](#1-23-1).,
const CmdResetHomeChanged CmdDef = 1

type Ardrone3GPSSettingsStateResetHomeChanged Command

type Ardrone3GPSSettingsStateResetHomeChangedArguments struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

func (a Ardrone3GPSSettingsStateResetHomeChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3GPSSettingsStateResetHomeChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Latitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Longitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Altitude)
	offset += 8

	return arg
}
func (a Ardrone3GPSSettingsStateResetHomeChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var GPSSettingsStateResetHomeChanged = Ardrone3GPSSettingsStateResetHomeChanged{
	Project: ProjectArdrone3,
	Class:   ClassGPSSettingsState,
	Cmd:     CmdResetHomeChanged,
}

// title : Gps fix info,
// desc : Gps fix info.,
// support : 0901;090c;090e,
// triggered : on change.,
const CmdGPSFixStateChanged CmdDef = 2

type Ardrone3GPSSettingsStateGPSFixStateChanged Command

type Ardrone3GPSSettingsStateGPSFixStateChangedArguments struct {
	Fixed uint8
}

func (a Ardrone3GPSSettingsStateGPSFixStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3GPSSettingsStateGPSFixStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Fixed)
	offset++

	return arg
}
func (a Ardrone3GPSSettingsStateGPSFixStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var GPSSettingsStateGPSFixStateChanged = Ardrone3GPSSettingsStateGPSFixStateChanged{
	Project: ProjectArdrone3,
	Class:   ClassGPSSettingsState,
	Cmd:     CmdGPSFixStateChanged,
}

// title : Gps update state,
// desc : Gps update state.,
// support : 0901;090c;090e,
// triggered : on change.,
const CmdGPSUpdateStateChanged CmdDef = 3

type Ardrone3GPSSettingsStateGPSUpdateStateChanged Command

type Ardrone3GPSSettingsStateGPSUpdateStateChangedArguments struct {
	State uint32
}

func (a Ardrone3GPSSettingsStateGPSUpdateStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3GPSSettingsStateGPSUpdateStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.State)
	offset += 4

	return arg
}
func (a Ardrone3GPSSettingsStateGPSUpdateStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var GPSSettingsStateGPSUpdateStateChanged = Ardrone3GPSSettingsStateGPSUpdateStateChanged{
	Project: ProjectArdrone3,
	Class:   ClassGPSSettingsState,
	Cmd:     CmdGPSUpdateStateChanged,
}

// title : Preferred home type,
// desc : User preference for the home type.\n See [HomeType](#1-31-2) to get the drone actual home type.,
// support : 0901;090c;090e,
// triggered : by [SetPreferredHomeType](#1-23-3).,
const CmdHomeTypeChanged CmdDef = 4

type Ardrone3GPSSettingsStateHomeTypeChanged Command

type Ardrone3GPSSettingsStateHomeTypeChangedArguments struct {
	TypeX uint32
}

func (a Ardrone3GPSSettingsStateHomeTypeChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3GPSSettingsStateHomeTypeChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TypeX)
	offset += 4

	return arg
}
func (a Ardrone3GPSSettingsStateHomeTypeChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var GPSSettingsStateHomeTypeChanged = Ardrone3GPSSettingsStateHomeTypeChanged{
	Project: ProjectArdrone3,
	Class:   ClassGPSSettingsState,
	Cmd:     CmdHomeTypeChanged,
}

// title : Return home delay,
// desc : Return home trigger delay. This delay represents the time after which the return home is automatically triggered after a disconnection.,
// support : 0901;090c;090e,
// triggered : by [SetReturnHomeDelay](#1-23-4).,
const CmdReturnHomeDelayChanged CmdDef = 5

type Ardrone3GPSSettingsStateReturnHomeDelayChanged Command

type Ardrone3GPSSettingsStateReturnHomeDelayChangedArguments struct {
	Delay uint16
}

func (a Ardrone3GPSSettingsStateReturnHomeDelayChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3GPSSettingsStateReturnHomeDelayChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.Delay)
	offset += 2

	return arg
}
func (a Ardrone3GPSSettingsStateReturnHomeDelayChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var GPSSettingsStateReturnHomeDelayChanged = Ardrone3GPSSettingsStateReturnHomeDelayChanged{
	Project: ProjectArdrone3,
	Class:   ClassGPSSettingsState,
	Cmd:     CmdReturnHomeDelayChanged,
}

// title : Geofence center,
// desc : Geofence center location. This location represents the center of the geofence zone. This is updated at a maximum frequency of 1 Hz.,
// triggered : when [HomeChanged](#1-24-0) and when [GpsLocationChanged](#1-4-9) before takeoff.,
const CmdGeofenceCenterChanged CmdDef = 6

type Ardrone3GPSSettingsStateGeofenceCenterChanged Command

type Ardrone3GPSSettingsStateGeofenceCenterChangedArguments struct {
	Latitude  float64
	Longitude float64
}

func (a Ardrone3GPSSettingsStateGeofenceCenterChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3GPSSettingsStateGeofenceCenterChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Latitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Longitude)
	offset += 8

	return arg
}
func (a Ardrone3GPSSettingsStateGeofenceCenterChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var GPSSettingsStateGeofenceCenterChanged = Ardrone3GPSSettingsStateGeofenceCenterChanged{
	Project: ProjectArdrone3,
	Class:   ClassGPSSettingsState,
	Cmd:     CmdGeofenceCenterChanged,
}

// title : Return home min altitude,
// desc : Minumum altitude for return home changed.,
// triggered : by [SetReturnHomeMinAltitude](#1-23-5).,
const CmdReturnHomeMinAltitudeChanged CmdDef = 7

type Ardrone3GPSSettingsStateReturnHomeMinAltitudeChanged Command

type Ardrone3GPSSettingsStateReturnHomeMinAltitudeChangedArguments struct {
	Value float32
	Min   float32
	Max   float32
}

func (a Ardrone3GPSSettingsStateReturnHomeMinAltitudeChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3GPSSettingsStateReturnHomeMinAltitudeChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Value)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Min)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Max)
	offset += 4

	return arg
}
func (a Ardrone3GPSSettingsStateReturnHomeMinAltitudeChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var GPSSettingsStateReturnHomeMinAltitudeChanged = Ardrone3GPSSettingsStateReturnHomeMinAltitudeChanged{
	Project: ProjectArdrone3,
	Class:   ClassGPSSettingsState,
	Cmd:     CmdReturnHomeMinAltitudeChanged,
}

// Camera state
const ClassCameraState ClassDef = 25

// title : Camera orientation,
// desc : Camera orientation.,
// support : 0901;090c;090e,
// triggered : by [SetCameraOrientation](#1-1-0).,
const CmdOrientationDUPLICATE CmdDef = 0

type Ardrone3CameraStateOrientation Command

type Ardrone3CameraStateOrientationArguments struct {
	Tilt int8
	Pan  int8
}

func (a Ardrone3CameraStateOrientation) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3CameraStateOrientationArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Tilt)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Pan)
	offset++

	return arg
}
func (a Ardrone3CameraStateOrientation) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CameraStateOrientation = Ardrone3CameraStateOrientation{
	Project: ProjectArdrone3,
	Class:   ClassCameraState,
	Cmd:     CmdOrientation,
}

// title : Orientation of the camera center,
// desc : Orientation of the center of the camera.\n This is the value to send when you want to center the camera.,
// support : 0901;090c;090e,
// triggered : at connection.,
const CmdDefaultCameraOrientation CmdDef = 1

type Ardrone3CameraStatedefaultCameraOrientation Command

type Ardrone3CameraStatedefaultCameraOrientationArguments struct {
	Tilt int8
	Pan  int8
}

func (a Ardrone3CameraStatedefaultCameraOrientation) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3CameraStatedefaultCameraOrientationArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Tilt)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Pan)
	offset++

	return arg
}
func (a Ardrone3CameraStatedefaultCameraOrientation) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CameraStatedefaultCameraOrientation = Ardrone3CameraStatedefaultCameraOrientation{
	Project: ProjectArdrone3,
	Class:   ClassCameraState,
	Cmd:     CmdDefaultCameraOrientation,
}

// title : Camera orientation,
// desc : Camera orientation with float arguments.,
// support : 0901;090c;090e,
// triggered : by [SetCameraOrientationV2](#1-1-1),
const CmdOrientationV2DUPLICATE CmdDef = 2

type Ardrone3CameraStateOrientationV2 Command

type Ardrone3CameraStateOrientationV2Arguments struct {
	Tilt float32
	Pan  float32
}

func (a Ardrone3CameraStateOrientationV2) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3CameraStateOrientationV2Arguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Tilt)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Pan)
	offset += 4

	return arg
}
func (a Ardrone3CameraStateOrientationV2) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CameraStateOrientationV2 = Ardrone3CameraStateOrientationV2{
	Project: ProjectArdrone3,
	Class:   ClassCameraState,
	Cmd:     CmdOrientationV2,
}

// title : Orientation of the camera center,
// desc : Orientation of the center of the camera.\n This is the value to send when you want to center the camera.,
// support : 0901;090c;090e,
// triggered : at connection.,
const CmdDefaultCameraOrientationV2 CmdDef = 3

type Ardrone3CameraStatedefaultCameraOrientationV2 Command

type Ardrone3CameraStatedefaultCameraOrientationV2Arguments struct {
	Tilt float32
	Pan  float32
}

func (a Ardrone3CameraStatedefaultCameraOrientationV2) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3CameraStatedefaultCameraOrientationV2Arguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Tilt)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Pan)
	offset += 4

	return arg
}
func (a Ardrone3CameraStatedefaultCameraOrientationV2) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CameraStatedefaultCameraOrientationV2 = Ardrone3CameraStatedefaultCameraOrientationV2{
	Project: ProjectArdrone3,
	Class:   ClassCameraState,
	Cmd:     CmdDefaultCameraOrientationV2,
}

// title : Camera velocity range,
// desc : Camera Orientation velocity limits.,
// support : 0901;090c;090e,
// triggered : at connection.,
const CmdVelocityRange CmdDef = 4

type Ardrone3CameraStateVelocityRange Command

type Ardrone3CameraStateVelocityRangeArguments struct {
	Maxtilt float32
	Maxpan  float32
}

func (a Ardrone3CameraStateVelocityRange) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3CameraStateVelocityRangeArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Maxtilt)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Maxpan)
	offset += 4

	return arg
}
func (a Ardrone3CameraStateVelocityRange) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CameraStateVelocityRange = Ardrone3CameraStateVelocityRange{
	Project: ProjectArdrone3,
	Class:   ClassCameraState,
	Cmd:     CmdVelocityRange,
}

// Anti-flickering related commands
const ClassAntiflickering ClassDef = 29

// title : Set the electric frequency,
// desc : Set the electric frequency of the surrounding lights.\n This is used to avoid the video flickering in auto mode. You can get the current antiflickering mode with the event [AntiflickeringModeChanged](#1-30-1).,
// support : 0901;090c,
// result : The electric frequency is set.\n Then, event [ElectricFrequency](#1-30-0) is triggered.,
const CmdElectricFrequency CmdDef = 0

type Ardrone3AntiflickeringelectricFrequency Command

type Ardrone3AntiflickeringelectricFrequencyArguments struct {
	Frequency uint32
}

func (a Ardrone3AntiflickeringelectricFrequency) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3AntiflickeringelectricFrequencyArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Frequency)
	offset += 4

	return arg
}
func (a Ardrone3AntiflickeringelectricFrequency) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var AntiflickeringelectricFrequency = Ardrone3AntiflickeringelectricFrequency{
	Project: ProjectArdrone3,
	Class:   ClassAntiflickering,
	Cmd:     CmdElectricFrequency,
}

// title : Set the antiflickering mode,
// desc : Set the antiflickering mode.\n If auto, the drone will detect when flickers appears on the video and trigger the antiflickering.\n In this case, this electric frequency it will use will be the one specified in the event [ElectricFrequency](#1-29-0).\n Forcing the antiflickering (FixedFiftyHertz or FixedFiftyHertz) can reduce luminosity of the video.,
// support : 0901;090c,
// result : The antiflickering mode is set.\n Then, event [AntiflickeringMode](#1-30-1) is triggered.,
const CmdSetMode CmdDef = 1

type Ardrone3AntiflickeringsetMode Command

type Ardrone3AntiflickeringsetModeArguments struct {
	Mode uint32
}

func (a Ardrone3AntiflickeringsetMode) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3AntiflickeringsetModeArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Mode)
	offset += 4

	return arg
}
func (a Ardrone3AntiflickeringsetMode) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var AntiflickeringsetMode = Ardrone3AntiflickeringsetMode{
	Project: ProjectArdrone3,
	Class:   ClassAntiflickering,
	Cmd:     CmdSetMode,
}

// Anti-flickering related states
const ClassAntiflickeringState ClassDef = 30

// title : Electric frequency,
// desc : Electric frequency.\n This piece of information is used for the antiflickering when the [AntiflickeringMode](#1-30-1) is set to *auto*.,
// support : 0901;090c,
// triggered : by [SetElectricFrequency](#1-29-0).,
const CmdElectricFrequencyChanged CmdDef = 0

type Ardrone3AntiflickeringStateelectricFrequencyChanged Command

type Ardrone3AntiflickeringStateelectricFrequencyChangedArguments struct {
	Frequency uint32
}

func (a Ardrone3AntiflickeringStateelectricFrequencyChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3AntiflickeringStateelectricFrequencyChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Frequency)
	offset += 4

	return arg
}
func (a Ardrone3AntiflickeringStateelectricFrequencyChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var AntiflickeringStateelectricFrequencyChanged = Ardrone3AntiflickeringStateelectricFrequencyChanged{
	Project: ProjectArdrone3,
	Class:   ClassAntiflickeringState,
	Cmd:     CmdElectricFrequencyChanged,
}

// title : Antiflickering mode,
// desc : Antiflickering mode.,
// support : 0901;090c,
// triggered : by [SetAntiflickeringMode](#1-29-1).,
const CmdModeChanged CmdDef = 1

type Ardrone3AntiflickeringStatemodeChanged Command

type Ardrone3AntiflickeringStatemodeChangedArguments struct {
	Mode uint32
}

func (a Ardrone3AntiflickeringStatemodeChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3AntiflickeringStatemodeChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Mode)
	offset += 4

	return arg
}
func (a Ardrone3AntiflickeringStatemodeChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var AntiflickeringStatemodeChanged = Ardrone3AntiflickeringStatemodeChanged{
	Project: ProjectArdrone3,
	Class:   ClassAntiflickeringState,
	Cmd:     CmdModeChanged,
}

// GPS related States
const ClassGPSState ClassDef = 31

// title : Number of GPS satellites,
// desc : Number of GPS satellites.,
// support : 0901;090c;090e,
// triggered : on change.,
const CmdNumberOfSatelliteChanged CmdDef = 0

type Ardrone3GPSStateNumberOfSatelliteChanged Command

type Ardrone3GPSStateNumberOfSatelliteChangedArguments struct {
	NumberOfSatellite uint8
}

func (a Ardrone3GPSStateNumberOfSatelliteChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3GPSStateNumberOfSatelliteChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.NumberOfSatellite)
	offset++

	return arg
}
func (a Ardrone3GPSStateNumberOfSatelliteChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var GPSStateNumberOfSatelliteChanged = Ardrone3GPSStateNumberOfSatelliteChanged{
	Project: ProjectArdrone3,
	Class:   ClassGPSState,
	Cmd:     CmdNumberOfSatelliteChanged,
}

// title : Home type availability,
// desc : Home type availability.,
// support : 0901;090c;090e,
// triggered : when the availability of, at least, one type changes.\n This might be due to controller position availability, gps fix before take off or other reason.,
const CmdHomeTypeAvailabilityChanged CmdDef = 1

type Ardrone3GPSStateHomeTypeAvailabilityChanged Command

type Ardrone3GPSStateHomeTypeAvailabilityChangedArguments struct {
	TypeX     uint32
	Available uint8
}

func (a Ardrone3GPSStateHomeTypeAvailabilityChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3GPSStateHomeTypeAvailabilityChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TypeX)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Available)
	offset++

	return arg
}
func (a Ardrone3GPSStateHomeTypeAvailabilityChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var GPSStateHomeTypeAvailabilityChanged = Ardrone3GPSStateHomeTypeAvailabilityChanged{
	Project: ProjectArdrone3,
	Class:   ClassGPSState,
	Cmd:     CmdHomeTypeAvailabilityChanged,
}

// title : Home type,
// desc : Home type.\n This choice is made by the drone, according to the [PreferredHomeType](#1-24-4) and the [HomeTypeAvailability](#1-31-1). The drone will choose the type matching with the user preference only if this type is available. If not, it will chose a type in this order:\n FOLLOWEE ; TAKEOFF ; PILOT ; FIRST_FIX,
// support : 0901;090c;090e,
// triggered : when the return home type chosen by the drone changes.\n This might be produced by a user preference triggered by [SetPreferedHomeType](#1-23-3) or by a change in the [HomeTypesAvailabilityChanged](#1-31-1).,
const CmdHomeTypeChosenChanged CmdDef = 2

type Ardrone3GPSStateHomeTypeChosenChanged Command

type Ardrone3GPSStateHomeTypeChosenChangedArguments struct {
	TypeX uint32
}

func (a Ardrone3GPSStateHomeTypeChosenChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3GPSStateHomeTypeChosenChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TypeX)
	offset += 4

	return arg
}
func (a Ardrone3GPSStateHomeTypeChosenChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var GPSStateHomeTypeChosenChanged = Ardrone3GPSStateHomeTypeChosenChanged{
	Project: ProjectArdrone3,
	Class:   ClassGPSState,
	Cmd:     CmdHomeTypeChosenChanged,
}

// Pro features enabled on the Bebop
const ClassPROState ClassDef = 32

// title : Pro features,
// desc : Pro features.,
const CmdFeatures CmdDef = 0

type Ardrone3PROStateFeatures Command

type Ardrone3PROStateFeaturesArguments struct {
	Features uint64
}

func (a Ardrone3PROStateFeatures) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3PROStateFeaturesArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Features)
	offset += 8

	return arg
}
func (a Ardrone3PROStateFeatures) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var PROStateFeatures = Ardrone3PROStateFeatures{
	Project: ProjectArdrone3,
	Class:   ClassPROState,
	Cmd:     CmdFeatures,
}

// Information about the connected accessories
const ClassAccessoryState ClassDef = 33

// title : List of connected accessories,
// desc : List of all connected accessories. This event presents the list of all connected accessories. To actually use the component, use the component dedicated feature.,
// support : 090e:1.5.0,
// triggered : at connection or when an accessory is connected.,
const CmdConnectedAccessories CmdDef = 0

type Ardrone3AccessoryStateConnectedAccessories Command

type Ardrone3AccessoryStateConnectedAccessoriesArguments struct {
	Id            uint8
	Accessorytype uint32
	Uid           string
	SwVersion     string
	Listflags     uint8
}

func (a Ardrone3AccessoryStateConnectedAccessories) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := Ardrone3AccessoryStateConnectedAccessoriesArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Id)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Accessorytype)
	offset += 4

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Uid = string(b[offset : offset+stringEnd])
	offset += stringEnd

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.SwVersion = string(b[offset : offset+stringEnd])
	offset += stringEnd
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Listflags)
	offset++

	return arg
}
func (a Ardrone3AccessoryStateConnectedAccessories) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var AccessoryStateConnectedAccessories = Ardrone3AccessoryStateConnectedAccessories{
	Project: ProjectArdrone3,
	Class:   ClassAccessoryState,
	Cmd:     CmdConnectedAccessories,
}

// title : Connected accessories battery,
// desc : Connected accessories battery.,
// support : none,
const CmdBattery CmdDef = 1

type Ardrone3AccessoryStateBattery Command

type Ardrone3AccessoryStateBatteryArguments struct {
	Id           uint8
	BatteryLevel uint8
	Listflags    uint8
}

func (a Ardrone3AccessoryStateBattery) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3AccessoryStateBatteryArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Id)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.BatteryLevel)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Listflags)
	offset++

	return arg
}
func (a Ardrone3AccessoryStateBattery) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var AccessoryStateBattery = Ardrone3AccessoryStateBattery{
	Project: ProjectArdrone3,
	Class:   ClassAccessoryState,
	Cmd:     CmdBattery,
}

// Sounds related commands
const ClassSound ClassDef = 35

// title : Start alert sound,
// desc : Start the alert sound. The alert sound can only be started when the drone is not flying.,
// support : none,
// result : The drone makes a sound and send back [AlertSoundState](#1-36-0) with state playing.,
const CmdStartAlertSound CmdDef = 0

type Ardrone3SoundStartAlertSound Command

type Ardrone3SoundStartAlertSoundArguments struct {
}

func (a Ardrone3SoundStartAlertSound) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3SoundStartAlertSoundArguments{}
	// No arguments to decode here !!

	return arg
}
func (a Ardrone3SoundStartAlertSound) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SoundStartAlertSound = Ardrone3SoundStartAlertSound{
	Project: ProjectArdrone3,
	Class:   ClassSound,
	Cmd:     CmdStartAlertSound,
}

// title : Stop alert sound,
// desc : Stop the alert sound.,
// support : none,
// result : The drone stops its alert sound and send back [AlertSoundState](#1-36-0) with state stopped.,
const CmdStopAlertSound CmdDef = 1

type Ardrone3SoundStopAlertSound Command

type Ardrone3SoundStopAlertSoundArguments struct {
}

func (a Ardrone3SoundStopAlertSound) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3SoundStopAlertSoundArguments{}
	// No arguments to decode here !!

	return arg
}
func (a Ardrone3SoundStopAlertSound) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SoundStopAlertSound = Ardrone3SoundStopAlertSound{
	Project: ProjectArdrone3,
	Class:   ClassSound,
	Cmd:     CmdStopAlertSound,
}

// Sounds related events
const ClassSoundState ClassDef = 36

// title : Alert sound state,
// desc : Alert sound state.,
// support : none,
// triggered : by [StartAlertSound](#1-35-0) or [StopAlertSound](#1-35-1) or when the drone starts or stops to play an alert sound by itself.,
const CmdAlertSound CmdDef = 0

type Ardrone3SoundStateAlertSound Command

type Ardrone3SoundStateAlertSoundArguments struct {
	State uint32
}

func (a Ardrone3SoundStateAlertSound) Decode(b []byte) interface{} {
	//TODO: .............
	arg := Ardrone3SoundStateAlertSoundArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.State)
	offset += 4

	return arg
}
func (a Ardrone3SoundStateAlertSound) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SoundStateAlertSound = Ardrone3SoundStateAlertSound{
	Project: ProjectArdrone3,
	Class:   ClassSoundState,
	Cmd:     CmdAlertSound,
}

// All common commands shared between all projects
const ProjectCommon ProjectDef = 0

// Network related commands
const ClassNetworkDUPLICATE ClassDef = 0

// title : Signals the remote that the host will disconnect,
// desc : Signals the remote that the host will disconnect.\n,
// support : none,
// result : None,
const CmdDisconnect CmdDef = 0

type CommonNetworkDisconnect Command

type CommonNetworkDisconnectArguments struct {
}

func (a CommonNetworkDisconnect) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonNetworkDisconnectArguments{}
	// No arguments to decode here !!

	return arg
}
func (a CommonNetworkDisconnect) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var NetworkDisconnect = CommonNetworkDisconnect{
	Project: ProjectCommon,
	Class:   ClassNetworkDUPLICATE,
	Cmd:     CmdDisconnect,
}

// Network Event from product
const ClassNetworkEvent ClassDef = 1

// title : Drone will disconnect,
// desc : Drone will disconnect.\n This event is mainly triggered when the user presses on the power button of the product.\n\n **This event is a notification, you can't retrieve it in the cache of the device controller.**,
// support : 0901;090c,
// triggered : mainly when the user presses the power button of the drone.,
const CmdDisconnection CmdDef = 0

type CommonNetworkEventDisconnection Command

type CommonNetworkEventDisconnectionArguments struct {
	Cause uint32
}

func (a CommonNetworkEventDisconnection) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonNetworkEventDisconnectionArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Cause)
	offset += 4

	return arg
}
func (a CommonNetworkEventDisconnection) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var NetworkEventDisconnection = CommonNetworkEventDisconnection{
	Project: ProjectCommon,
	Class:   ClassNetworkEvent,
	Cmd:     CmdDisconnection,
}

// Settings commands
const ClassSettings ClassDef = 2

// title : Ask for all settings,
// desc : Ask for all settings.\n\n **Please note that you should not send this command if you are using the\n libARController API as this library is handling the connection process for you.**,
// support : drones,
// result : The product will trigger all settings events (such as [CameraSettings](#0-15-0), or product specific settings as the [MaxAltitude](#1-6-0) for the Bebop).\n Then, it will trigger [AllSettingsEnd](#0-3-0).,
const CmdAllSettings CmdDef = 0

type CommonSettingsAllSettings Command

type CommonSettingsAllSettingsArguments struct {
}

func (a CommonSettingsAllSettings) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonSettingsAllSettingsArguments{}
	// No arguments to decode here !!

	return arg
}
func (a CommonSettingsAllSettings) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsAllSettings = CommonSettingsAllSettings{
	Project: ProjectCommon,
	Class:   ClassSettings,
	Cmd:     CmdAllSettings,
}

// title : Reset all settings,
// desc : Reset all settings.,
// support : drones,
// result : It will trigger [ResetChanged](#0-3-1).\n Then, the product will trigger all settings events (such as [CameraSettings](#0-15-0), or product specific settings as the [MaxAltitude](#1-6-0) for the Bebop) with factory values.,
const CmdReset CmdDef = 1

type CommonSettingsReset Command

type CommonSettingsResetArguments struct {
}

func (a CommonSettingsReset) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonSettingsResetArguments{}
	// No arguments to decode here !!

	return arg
}
func (a CommonSettingsReset) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsReset = CommonSettingsReset{
	Project: ProjectCommon,
	Class:   ClassSettings,
	Cmd:     CmdReset,
}

// title : Set product name,
// desc : Set the product name.\n It also sets the name of the SSID for Wifi products and advertisement name for BLE products (changed after a reboot of the product).,
// support : drones,
// result : Name is changed.\n Then, it will trigger [NameChanged](#0-3-2).,
const CmdProductName CmdDef = 2

type CommonSettingsProductName Command

type CommonSettingsProductNameArguments struct {
	Name string
}

func (a CommonSettingsProductName) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonSettingsProductNameArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Name = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonSettingsProductName) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsProductName = CommonSettingsProductName{
	Project: ProjectCommon,
	Class:   ClassSettings,
	Cmd:     CmdProductName,
}

// title : Set the country,
// desc : Set the country for Wifi products.\n This can modify Wifi band and/or channel.\n **Please note that you might be disconnected from the product after changing the country as it changes Wifi parameters.**,
// support : 0901;0902;0905;0906;090c;090e,
// result : The country is set.\n Then, it will trigger [CountryChanged](#0-3-6).,
const CmdCountry CmdDef = 3

type CommonSettingsCountry Command

type CommonSettingsCountryArguments struct {
	Code string
}

func (a CommonSettingsCountry) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonSettingsCountryArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Code = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonSettingsCountry) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsCountry = CommonSettingsCountry{
	Project: ProjectCommon,
	Class:   ClassSettings,
	Cmd:     CmdCountry,
}

// title : Enable auto-country,
// desc : Enable auto-country.\n If auto-country is set, the drone will guess its Wifi country by itself by checking other Wifi country around it.\n **Please note that you might be disconnected from the product after changing the country as it changes Wifi parameters.**,
// support : 0901;0902;0905;0906;090c;090e,
// result : The auto-country of the product is changed.\n Then, it will trigger [AutoCountryChanged](#0-3-7) and [CountryChanged](#0-3-6).,
const CmdAutoCountry CmdDef = 4

type CommonSettingsAutoCountry Command

type CommonSettingsAutoCountryArguments struct {
	Automatic uint8
}

func (a CommonSettingsAutoCountry) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonSettingsAutoCountryArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Automatic)
	offset++

	return arg
}
func (a CommonSettingsAutoCountry) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsAutoCountry = CommonSettingsAutoCountry{
	Project: ProjectCommon,
	Class:   ClassSettings,
	Cmd:     CmdAutoCountry,
}

// Settings state from product
const ClassSettingsStateDUPLICATE ClassDef = 3

// title : All settings have been sent,
// desc : All settings have been sent.\n\n **Please note that you should not care about this event if you are using the libARController API as this library is handling the connection process for you.**,
// support : drones,
// triggered : when all settings values have been sent.,
const CmdAllSettingsChanged CmdDef = 0

type CommonSettingsStateAllSettingsChanged Command

type CommonSettingsStateAllSettingsChangedArguments struct {
}

func (a CommonSettingsStateAllSettingsChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonSettingsStateAllSettingsChangedArguments{}
	// No arguments to decode here !!

	return arg
}
func (a CommonSettingsStateAllSettingsChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsStateAllSettingsChanged = CommonSettingsStateAllSettingsChanged{
	Project: ProjectCommon,
	Class:   ClassSettingsStateDUPLICATE,
	Cmd:     CmdAllSettingsChanged,
}

// title : All settings have been reset,
// desc : All settings have been reset.,
// support : drones,
// triggered : by [ResetSettings](#0-2-1).,
const CmdResetChanged CmdDef = 1

type CommonSettingsStateResetChanged Command

type CommonSettingsStateResetChangedArguments struct {
}

func (a CommonSettingsStateResetChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonSettingsStateResetChangedArguments{}
	// No arguments to decode here !!

	return arg
}
func (a CommonSettingsStateResetChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsStateResetChanged = CommonSettingsStateResetChanged{
	Project: ProjectCommon,
	Class:   ClassSettingsState,
	Cmd:     CmdResetChanged,
}

// title : Product name changed,
// desc : Product name changed.,
// support : drones,
// triggered : by [SetProductName](#0-2-2).,
const CmdProductNameChanged CmdDef = 2

type CommonSettingsStateProductNameChanged Command

type CommonSettingsStateProductNameChangedArguments struct {
	Name string
}

func (a CommonSettingsStateProductNameChanged) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonSettingsStateProductNameChangedArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Name = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonSettingsStateProductNameChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsStateProductNameChanged = CommonSettingsStateProductNameChanged{
	Project: ProjectCommon,
	Class:   ClassSettingsState,
	Cmd:     CmdProductNameChanged,
}

// title : Product version,
// desc : Product version.,
// support : drones,
// triggered : during the connection process.,
const CmdProductVersionChanged CmdDef = 3

type CommonSettingsStateProductVersionChanged Command

type CommonSettingsStateProductVersionChangedArguments struct {
	Software string
	Hardware string
}

func (a CommonSettingsStateProductVersionChanged) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonSettingsStateProductVersionChangedArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Software = string(b[offset : offset+stringEnd])
	offset += stringEnd

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Hardware = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonSettingsStateProductVersionChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsStateProductVersionChanged = CommonSettingsStateProductVersionChanged{
	Project: ProjectCommon,
	Class:   ClassSettingsState,
	Cmd:     CmdProductVersionChanged,
}

// title : Product serial (1st part),
// desc : Product serial (1st part).,
// support : drones,
// triggered : during the connection process.,
const CmdProductSerialHighChanged CmdDef = 4

type CommonSettingsStateProductSerialHighChanged Command

type CommonSettingsStateProductSerialHighChangedArguments struct {
	High string
}

func (a CommonSettingsStateProductSerialHighChanged) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonSettingsStateProductSerialHighChangedArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.High = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonSettingsStateProductSerialHighChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsStateProductSerialHighChanged = CommonSettingsStateProductSerialHighChanged{
	Project: ProjectCommon,
	Class:   ClassSettingsState,
	Cmd:     CmdProductSerialHighChanged,
}

// title : Product serial (2nd part),
// desc : Product serial (2nd part).,
// support : drones,
// triggered : during the connection process.,
const CmdProductSerialLowChanged CmdDef = 5

type CommonSettingsStateProductSerialLowChanged Command

type CommonSettingsStateProductSerialLowChangedArguments struct {
	Low string
}

func (a CommonSettingsStateProductSerialLowChanged) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonSettingsStateProductSerialLowChangedArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Low = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonSettingsStateProductSerialLowChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsStateProductSerialLowChanged = CommonSettingsStateProductSerialLowChanged{
	Project: ProjectCommon,
	Class:   ClassSettingsState,
	Cmd:     CmdProductSerialLowChanged,
}

// title : Country changed,
// desc : Country changed.,
// support : drones,
// triggered : by [SetCountry](#0-2-3).,
const CmdCountryChanged CmdDef = 6

type CommonSettingsStateCountryChanged Command

type CommonSettingsStateCountryChangedArguments struct {
	Code string
}

func (a CommonSettingsStateCountryChanged) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonSettingsStateCountryChangedArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Code = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonSettingsStateCountryChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsStateCountryChanged = CommonSettingsStateCountryChanged{
	Project: ProjectCommon,
	Class:   ClassSettingsState,
	Cmd:     CmdCountryChanged,
}

// title : Auto-country changed,
// desc : Auto-country changed.,
// support : drones,
// triggered : by [SetAutoCountry](#0-2-4).,
const CmdAutoCountryChanged CmdDef = 7

type CommonSettingsStateAutoCountryChanged Command

type CommonSettingsStateAutoCountryChangedArguments struct {
	Automatic uint8
}

func (a CommonSettingsStateAutoCountryChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonSettingsStateAutoCountryChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Automatic)
	offset++

	return arg
}
func (a CommonSettingsStateAutoCountryChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsStateAutoCountryChanged = CommonSettingsStateAutoCountryChanged{
	Project: ProjectCommon,
	Class:   ClassSettingsState,
	Cmd:     CmdAutoCountryChanged,
}

// title : Board id,
// desc : Board id.,
// support : drones,
// triggered : during the connection process.,
const CmdBoardIdChanged CmdDef = 8

type CommonSettingsStateBoardIdChanged Command

type CommonSettingsStateBoardIdChangedArguments struct {
	Id string
}

func (a CommonSettingsStateBoardIdChanged) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonSettingsStateBoardIdChangedArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Id = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonSettingsStateBoardIdChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var SettingsStateBoardIdChanged = CommonSettingsStateBoardIdChanged{
	Project: ProjectCommon,
	Class:   ClassSettingsState,
	Cmd:     CmdBoardIdChanged,
}

// Common commands
const ClassCommon ClassDef = 4

// title : Ask for all states,
// desc : Ask for all states.\n\n **Please note that you should not send this command if you are using the\n libARController API as this library is handling the connection process for you.**,
// support : drones,
// result : The product will trigger all states events (such as [FlyingState](#1-4-1) for the Bebop).\n Then, it will trigger [AllStatesEnd](#0-5-0).,
const CmdAllStates CmdDef = 0

type CommonCommonAllStates Command

type CommonCommonAllStatesArguments struct {
}

func (a CommonCommonAllStates) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCommonAllStatesArguments{}
	// No arguments to decode here !!

	return arg
}
func (a CommonCommonAllStates) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonAllStates = CommonCommonAllStates{
	Project: ProjectCommon,
	Class:   ClassCommon,
	Cmd:     CmdAllStates,
}

// title : Set the date,
// desc : Set the date.\n This date is taken by the drone as its own date.\n So medias and other files will be dated from this date\n\n **Please note that you should not send this command if you are using the\n libARController API as this library is handling the connection process for you.**,
// support : drones,
// result : The date of the product is set.\n Then, it will trigger [DateChanged](#0-5-4).,
const CmdCurrentDate CmdDef = 1

type CommonCommonCurrentDate Command

type CommonCommonCurrentDateArguments struct {
	Date string
}

func (a CommonCommonCurrentDate) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonCommonCurrentDateArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Date = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonCommonCurrentDate) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonCurrentDate = CommonCommonCurrentDate{
	Project: ProjectCommon,
	Class:   ClassCommon,
	Cmd:     CmdCurrentDate,
}

// title : Set the time,
// desc : Set the time.\n This time is taken by the drone as its own time.\n So medias and other files will be dated from this time\n\n **Please note that you should not send this command if you are using the\n libARController API as this library is handling the connection process for you.**,
// support : drones,
// result : The time of the product is set.\n Then, it will trigger [TimeChanged](#0-5-5).,
const CmdCurrentTime CmdDef = 2

type CommonCommonCurrentTime Command

type CommonCommonCurrentTimeArguments struct {
	Time string
}

func (a CommonCommonCurrentTime) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonCommonCurrentTimeArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Time = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonCommonCurrentTime) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonCurrentTime = CommonCommonCurrentTime{
	Project: ProjectCommon,
	Class:   ClassCommon,
	Cmd:     CmdCurrentTime,
}

// title : Reboot,
// desc : Reboot the product.\n The product will accept this command only if is not flying.,
// support : drones,
// result : The product will reboot if it can.,
const CmdReboot CmdDef = 3

type CommonCommonReboot Command

type CommonCommonRebootArguments struct {
}

func (a CommonCommonReboot) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCommonRebootArguments{}
	// No arguments to decode here !!

	return arg
}
func (a CommonCommonReboot) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonReboot = CommonCommonReboot{
	Project: ProjectCommon,
	Class:   ClassCommon,
	Cmd:     CmdReboot,
}

// title : Set the datetime,
// desc : Set both the date and the time with only one command.\n If using this command, do not use [CurrentDate](#0-4-1) and [CurrentTime](#0-4-2) commands.\n This datetime is taken by the drone as its own datetime.\n So medias and other files will be dated from this datetime\n\n **Please note that you should not send this command if you are using the\n libARController API as this library is handling the connection process for you.**,
// support : 0914,
// result : The datetime of the product is set.\n Then, it will trigger [CurrentDateTimeChanged](#0-5-15).,
const CmdCurrentDateTime CmdDef = 4

type CommonCommonCurrentDateTime Command

type CommonCommonCurrentDateTimeArguments struct {
	Datetime string
}

func (a CommonCommonCurrentDateTime) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonCommonCurrentDateTimeArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Datetime = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonCommonCurrentDateTime) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonCurrentDateTime = CommonCommonCurrentDateTime{
	Project: ProjectCommon,
	Class:   ClassCommon,
	Cmd:     CmdCurrentDateTime,
}

// Common state from product
const ClassCommonState ClassDef = 5

// title : All states have been sent,
// desc : All states have been sent.\n\n **Please note that you should not care about this event if you are using the libARController API as this library is handling the connection process for you.**,
// support : drones,
// triggered : when all states values have been sent.,
const CmdAllStatesChanged CmdDef = 0

type CommonCommonStateAllStatesChanged Command

type CommonCommonStateAllStatesChangedArguments struct {
}

func (a CommonCommonStateAllStatesChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCommonStateAllStatesChangedArguments{}
	// No arguments to decode here !!

	return arg
}
func (a CommonCommonStateAllStatesChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonStateAllStatesChanged = CommonCommonStateAllStatesChanged{
	Project: ProjectCommon,
	Class:   ClassCommonState,
	Cmd:     CmdAllStatesChanged,
}

// title : Battery state,
// desc : Battery state.,
// support : drones,
// triggered : when the battery level changes.,
const CmdBatteryStateChanged CmdDef = 1

type CommonCommonStateBatteryStateChanged Command

type CommonCommonStateBatteryStateChangedArguments struct {
	Percent uint8
}

func (a CommonCommonStateBatteryStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCommonStateBatteryStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Percent)
	offset++

	return arg
}
func (a CommonCommonStateBatteryStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonStateBatteryStateChanged = CommonCommonStateBatteryStateChanged{
	Project: ProjectCommon,
	Class:   ClassCommonState,
	Cmd:     CmdBatteryStateChanged,
}

// title : Mass storage state list,
// desc : Mass storage state list.,
// support : drones,
// triggered : when a mass storage is inserted or ejected.,
const CmdMassStorageStateListChanged CmdDef = 2

type CommonCommonStateMassStorageStateListChanged Command

type CommonCommonStateMassStorageStateListChangedArguments struct {
	Massstorageid uint8
	Name          string
}

func (a CommonCommonStateMassStorageStateListChanged) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonCommonStateMassStorageStateListChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Massstorageid)
	offset++

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Name = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonCommonStateMassStorageStateListChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonStateMassStorageStateListChanged = CommonCommonStateMassStorageStateListChanged{
	Project: ProjectCommon,
	Class:   ClassCommonState,
	Cmd:     CmdMassStorageStateListChanged,
}

// title : Mass storage info state list,
// desc : Mass storage info state list.,
// support : drones,
// triggered : when a mass storage info changes.,
const CmdMassStorageInfoStateListChanged CmdDef = 3

type CommonCommonStateMassStorageInfoStateListChanged Command

type CommonCommonStateMassStorageInfoStateListChangedArguments struct {
	Massstorageid uint8
	Size          uint32
	Usedsize      uint32
	Plugged       uint8
	Full          uint8
	Internal      uint8
}

func (a CommonCommonStateMassStorageInfoStateListChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCommonStateMassStorageInfoStateListChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Massstorageid)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Size)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Usedsize)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Plugged)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Full)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Internal)
	offset++

	return arg
}
func (a CommonCommonStateMassStorageInfoStateListChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonStateMassStorageInfoStateListChanged = CommonCommonStateMassStorageInfoStateListChanged{
	Project: ProjectCommon,
	Class:   ClassCommonState,
	Cmd:     CmdMassStorageInfoStateListChanged,
}

// title : Date changed,
// desc : Date changed.\n Corresponds to the latest date set on the drone.\n\n **Please note that you should not care about this event if you are using the libARController API as this library is handling the connection process for you.**,
// support : drones,
// triggered : by [SetDate](#0-4-1).,
const CmdCurrentDateChanged CmdDef = 4

type CommonCommonStateCurrentDateChanged Command

type CommonCommonStateCurrentDateChangedArguments struct {
	Date string
}

func (a CommonCommonStateCurrentDateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonCommonStateCurrentDateChangedArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Date = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonCommonStateCurrentDateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonStateCurrentDateChanged = CommonCommonStateCurrentDateChanged{
	Project: ProjectCommon,
	Class:   ClassCommonState,
	Cmd:     CmdCurrentDateChanged,
}

// title : Time changed,
// desc : Time changed.\n Corresponds to the latest time set on the drone.\n\n **Please note that you should not care about this event if you are using the libARController API as this library is handling the connection process for you.**,
// support : drones,
// triggered : by [SetTime](#0-4-2).,
const CmdCurrentTimeChanged CmdDef = 5

type CommonCommonStateCurrentTimeChanged Command

type CommonCommonStateCurrentTimeChangedArguments struct {
	Time string
}

func (a CommonCommonStateCurrentTimeChanged) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonCommonStateCurrentTimeChangedArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Time = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonCommonStateCurrentTimeChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonStateCurrentTimeChanged = CommonCommonStateCurrentTimeChanged{
	Project: ProjectCommon,
	Class:   ClassCommonState,
	Cmd:     CmdCurrentTimeChanged,
}

// title : Mass storage remaining data list,
// desc : Mass storage remaining data list.,
const CmdMassStorageInfoRemainingListChanged CmdDef = 6

type CommonCommonStateMassStorageInfoRemainingListChanged Command

type CommonCommonStateMassStorageInfoRemainingListChangedArguments struct {
	Freespace      uint32
	Rectime        uint16
	Photoremaining uint32
}

func (a CommonCommonStateMassStorageInfoRemainingListChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCommonStateMassStorageInfoRemainingListChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Freespace)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.Rectime)
	offset += 2
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Photoremaining)
	offset += 4

	return arg
}
func (a CommonCommonStateMassStorageInfoRemainingListChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonStateMassStorageInfoRemainingListChanged = CommonCommonStateMassStorageInfoRemainingListChanged{
	Project: ProjectCommon,
	Class:   ClassCommonState,
	Cmd:     CmdMassStorageInfoRemainingListChanged,
}

// title : Rssi changed,
// desc : Rssi (Wifi Signal between controller and product) changed.,
// support : 0901;0902;0905;0906;090c;090e,
// triggered : regularly.,
const CmdWifiSignalChanged CmdDef = 7

type CommonCommonStateWifiSignalChanged Command

type CommonCommonStateWifiSignalChangedArguments struct {
	Rssi int16
}

func (a CommonCommonStateWifiSignalChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCommonStateWifiSignalChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.Rssi)
	offset += 2

	return arg
}
func (a CommonCommonStateWifiSignalChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonStateWifiSignalChanged = CommonCommonStateWifiSignalChanged{
	Project: ProjectCommon,
	Class:   ClassCommonState,
	Cmd:     CmdWifiSignalChanged,
}

// title : Sensors state list,
// desc : Sensors state list.,
// support : 0901:2.0.3;0902;0905;0906;0907;0909;090a;090c;090e,
// triggered : at connection and when a sensor state changes.,
const CmdSensorsStatesListChanged CmdDef = 8

type CommonCommonStateSensorsStatesListChanged Command

type CommonCommonStateSensorsStatesListChangedArguments struct {
	SensorName  uint32
	SensorState uint8
}

func (a CommonCommonStateSensorsStatesListChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCommonStateSensorsStatesListChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.SensorName)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.SensorState)
	offset++

	return arg
}
func (a CommonCommonStateSensorsStatesListChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonStateSensorsStatesListChanged = CommonCommonStateSensorsStatesListChanged{
	Project: ProjectCommon,
	Class:   ClassCommonState,
	Cmd:     CmdSensorsStatesListChanged,
}

// title : Product sub-model,
// desc : Product sub-model.\n This can be used to customize the UI depending on the product.,
// support : 0905;0906;0907;0909,
// triggered : at connection.,
const CmdProductModel CmdDef = 9

type CommonCommonStateProductModel Command

type CommonCommonStateProductModelArguments struct {
	Model uint32
}

func (a CommonCommonStateProductModel) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCommonStateProductModelArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Model)
	offset += 4

	return arg
}
func (a CommonCommonStateProductModel) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonStateProductModel = CommonCommonStateProductModel{
	Project: ProjectCommon,
	Class:   ClassCommonState,
	Cmd:     CmdProductModel,
}

// title : Country list,
// desc : List of countries known by the drone.,
const CmdCountryListKnown CmdDef = 10

type CommonCommonStateCountryListKnown Command

type CommonCommonStateCountryListKnownArguments struct {
	ListFlags    uint8
	CountryCodes string
}

func (a CommonCommonStateCountryListKnown) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonCommonStateCountryListKnownArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.ListFlags)
	offset++

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.CountryCodes = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonCommonStateCountryListKnown) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonStateCountryListKnown = CommonCommonStateCountryListKnown{
	Project: ProjectCommon,
	Class:   ClassCommonState,
	Cmd:     CmdCountryListKnown,
}

// title : Mass storage content changed,
// desc : Mass storage content changed.,
const CmdDeprecatedMassStorageContentChanged CmdDef = 11

type CommonCommonStateDeprecatedMassStorageContentChanged Command

type CommonCommonStateDeprecatedMassStorageContentChangedArguments struct {
	Massstorageid uint8
	NbPhotos      uint16
	NbVideos      uint16
	NbPuds        uint16
	NbCrashLogs   uint16
}

func (a CommonCommonStateDeprecatedMassStorageContentChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCommonStateDeprecatedMassStorageContentChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Massstorageid)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.NbPhotos)
	offset += 2
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.NbVideos)
	offset += 2
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.NbPuds)
	offset += 2
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.NbCrashLogs)
	offset += 2

	return arg
}
func (a CommonCommonStateDeprecatedMassStorageContentChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonStateDeprecatedMassStorageContentChanged = CommonCommonStateDeprecatedMassStorageContentChanged{
	Project: ProjectCommon,
	Class:   ClassCommonState,
	Cmd:     CmdDeprecatedMassStorageContentChanged,
}

// title : Mass storage content,
// desc : Mass storage content.,
// support : 090c:4.0.0;090e:4.0.0,
// triggered : when the content of the mass storage changes.,
const CmdMassStorageContent CmdDef = 12

type CommonCommonStateMassStorageContent Command

type CommonCommonStateMassStorageContentArguments struct {
	Massstorageid uint8
	NbPhotos      uint16
	NbVideos      uint16
	NbPuds        uint16
	NbCrashLogs   uint16
	NbRawPhotos   uint16
}

func (a CommonCommonStateMassStorageContent) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCommonStateMassStorageContentArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Massstorageid)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.NbPhotos)
	offset += 2
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.NbVideos)
	offset += 2
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.NbPuds)
	offset += 2
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.NbCrashLogs)
	offset += 2
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.NbRawPhotos)
	offset += 2

	return arg
}
func (a CommonCommonStateMassStorageContent) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonStateMassStorageContent = CommonCommonStateMassStorageContent{
	Project: ProjectCommon,
	Class:   ClassCommonState,
	Cmd:     CmdMassStorageContent,
}

// title : Mass storage content for current run,
// desc : Mass storage content for current run.\n Only counts the files related to the current run (see [RunId](#0-30-0)),
// support : 090c:4.0.0;090e:4.0.0,
// triggered : when the content of the mass storage changes and this content is related to the current run.,
const CmdMassStorageContentForCurrentRun CmdDef = 13

type CommonCommonStateMassStorageContentForCurrentRun Command

type CommonCommonStateMassStorageContentForCurrentRunArguments struct {
	Massstorageid uint8
	NbPhotos      uint16
	NbVideos      uint16
	NbRawPhotos   uint16
}

func (a CommonCommonStateMassStorageContentForCurrentRun) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCommonStateMassStorageContentForCurrentRunArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Massstorageid)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.NbPhotos)
	offset += 2
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.NbVideos)
	offset += 2
	ConvLittleEndianSliceToNumeric(b[offset:offset+2], &arg.NbRawPhotos)
	offset += 2

	return arg
}
func (a CommonCommonStateMassStorageContentForCurrentRun) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonStateMassStorageContentForCurrentRun = CommonCommonStateMassStorageContentForCurrentRun{
	Project: ProjectCommon,
	Class:   ClassCommonState,
	Cmd:     CmdMassStorageContentForCurrentRun,
}

// title : Video recording timestamp,
// desc : Current or last video recording timestamp.\n Timestamp in milliseconds since 00:00:00 UTC on 1 January 1970.\n **Please note that values don't persist after drone reboot**,
// triggered : on video recording start and video recording stop or \n after that the date/time of the drone changed.,
const CmdVideoRecordingTimestamp CmdDef = 14

type CommonCommonStateVideoRecordingTimestamp Command

type CommonCommonStateVideoRecordingTimestampArguments struct {
	StartTimestamp uint64
	StopTimestamp  uint64
}

func (a CommonCommonStateVideoRecordingTimestamp) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCommonStateVideoRecordingTimestampArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.StartTimestamp)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.StopTimestamp)
	offset += 8

	return arg
}
func (a CommonCommonStateVideoRecordingTimestamp) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonStateVideoRecordingTimestamp = CommonCommonStateVideoRecordingTimestamp{
	Project: ProjectCommon,
	Class:   ClassCommonState,
	Cmd:     CmdVideoRecordingTimestamp,
}

// title : Datetime changed,
// desc : Both date and time changed.\n Corresponds to the latest datetime set on the drone.\n\n **Please note that you should not care about this event if you are using the libARController API as this library is handling the connection process for you.**,
// support : 0914,
// triggered : by [CurrentDateTime](#0-4-4).,
const CmdCurrentDateTimeChanged CmdDef = 15

type CommonCommonStateCurrentDateTimeChanged Command

type CommonCommonStateCurrentDateTimeChangedArguments struct {
	Datetime string
}

func (a CommonCommonStateCurrentDateTimeChanged) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonCommonStateCurrentDateTimeChangedArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Datetime = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonCommonStateCurrentDateTimeChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonStateCurrentDateTimeChanged = CommonCommonStateCurrentDateTimeChanged{
	Project: ProjectCommon,
	Class:   ClassCommonState,
	Cmd:     CmdCurrentDateTimeChanged,
}

// title : Link signal quality,
// desc : Link signal quality. Gives a overal indication of the radio link quality,
// support : 0914,
// triggered : when the link signal quality changes.,
const CmdLinkSignalQuality CmdDef = 16

type CommonCommonStateLinkSignalQuality Command

type CommonCommonStateLinkSignalQualityArguments struct {
	Value uint8
}

func (a CommonCommonStateLinkSignalQuality) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCommonStateLinkSignalQualityArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Value)
	offset++

	return arg
}
func (a CommonCommonStateLinkSignalQuality) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonStateLinkSignalQuality = CommonCommonStateLinkSignalQuality{
	Project: ProjectCommon,
	Class:   ClassCommonState,
	Cmd:     CmdLinkSignalQuality,
}

// title : Current Drone Boot id,
// desc : Current Drone Boot id.\n A Boot Id identifies a drone session and do not change between drone power on and power off.\n Also, each medias contains the Boot Id.,
// support : 0914,
// triggered : At connection.,
const CmdBootId CmdDef = 17

type CommonCommonStateBootId Command

type CommonCommonStateBootIdArguments struct {
	BootId string
}

func (a CommonCommonStateBootId) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonCommonStateBootIdArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.BootId = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonCommonStateBootId) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CommonStateBootId = CommonCommonStateBootId{
	Project: ProjectCommon,
	Class:   ClassCommonState,
	Cmd:     CmdBootId,
}

// Over heat commands
const ClassOverHeat ClassDef = 6

// title : Switch off after an overheat,
// desc : Switch off after an overheat.,
// support : none,
// result : None,
const CmdSwitchOff CmdDef = 0

type CommonOverHeatSwitchOff Command

type CommonOverHeatSwitchOffArguments struct {
}

func (a CommonOverHeatSwitchOff) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonOverHeatSwitchOffArguments{}
	// No arguments to decode here !!

	return arg
}
func (a CommonOverHeatSwitchOff) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var OverHeatSwitchOff = CommonOverHeatSwitchOff{
	Project: ProjectCommon,
	Class:   ClassOverHeat,
	Cmd:     CmdSwitchOff,
}

// title : Ventilate after an overheat,
// desc : Ventilate after an overheat.,
// support : none,
// result : None,
const CmdVentilate CmdDef = 1

type CommonOverHeatVentilate Command

type CommonOverHeatVentilateArguments struct {
}

func (a CommonOverHeatVentilate) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonOverHeatVentilateArguments{}
	// No arguments to decode here !!

	return arg
}
func (a CommonOverHeatVentilate) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var OverHeatVentilate = CommonOverHeatVentilate{
	Project: ProjectCommon,
	Class:   ClassOverHeat,
	Cmd:     CmdVentilate,
}

// Overheat state from product
const ClassOverHeatState ClassDef = 7

// title : Overheat,
// desc : Overheat temperature reached.,
const CmdOverHeatChanged CmdDef = 0

type CommonOverHeatStateOverHeatChanged Command

type CommonOverHeatStateOverHeatChangedArguments struct {
}

func (a CommonOverHeatStateOverHeatChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonOverHeatStateOverHeatChangedArguments{}
	// No arguments to decode here !!

	return arg
}
func (a CommonOverHeatStateOverHeatChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var OverHeatStateOverHeatChanged = CommonOverHeatStateOverHeatChanged{
	Project: ProjectCommon,
	Class:   ClassOverHeatState,
	Cmd:     CmdOverHeatChanged,
}

// title : Overheat regulation type,
// desc : Overheat regulation type.,
const CmdOverHeatRegulationChanged CmdDef = 1

type CommonOverHeatStateOverHeatRegulationChanged Command

type CommonOverHeatStateOverHeatRegulationChangedArguments struct {
	RegulationType uint8
}

func (a CommonOverHeatStateOverHeatRegulationChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonOverHeatStateOverHeatRegulationChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.RegulationType)
	offset++

	return arg
}
func (a CommonOverHeatStateOverHeatRegulationChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var OverHeatStateOverHeatRegulationChanged = CommonOverHeatStateOverHeatRegulationChanged{
	Project: ProjectCommon,
	Class:   ClassOverHeatState,
	Cmd:     CmdOverHeatRegulationChanged,
}

// Notify the device about the state of the controller application.
const ClassController ClassDef = 8

// title : Inform about hud entering,
// desc : Inform about hud entering.\n Tell the drone that the controller enters/leaves the piloting hud.\n On a non-flying products it is used to know when a run begins.,
// support : drones,
// result : If yes, the product will begin a new session (so it should send a new [runId](#0-30-0)).\n Also, on the JumpingSumos, if the video is in autorecord mode, it will start recording.,
const CmdIsPiloting CmdDef = 0

type CommonControllerisPiloting Command

type CommonControllerisPilotingArguments struct {
	Piloting uint8
}

func (a CommonControllerisPiloting) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonControllerisPilotingArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Piloting)
	offset++

	return arg
}
func (a CommonControllerisPiloting) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var ControllerisPiloting = CommonControllerisPiloting{
	Project: ProjectCommon,
	Class:   ClassController,
	Cmd:     CmdIsPiloting,
}

// title : A SDK peer has connected/disconnected,
// desc : A SDK peer (ie FreeFlight) has connected/disconnected to the Skycontroller.\n This is only meant to be sent by the Skycontroller, as it is acting as a proxy.,
// support : 0918,
// triggered : at connection and when the peer state changes.,
const CmdPeerStateChanged CmdDef = 1

type CommonControllerPeerStateChanged Command

type CommonControllerPeerStateChangedArguments struct {
	State    uint32
	TypeX    uint32
	PeerName string
	PeerId   string
	PeerType string
}

func (a CommonControllerPeerStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonControllerPeerStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.State)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TypeX)
	offset += 4

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.PeerName = string(b[offset : offset+stringEnd])
	offset += stringEnd

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.PeerId = string(b[offset : offset+stringEnd])
	offset += stringEnd

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.PeerType = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonControllerPeerStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var ControllerPeerStateChanged = CommonControllerPeerStateChanged{
	Project: ProjectCommon,
	Class:   ClassController,
	Cmd:     CmdPeerStateChanged,
}

// Wifi settings commands
const ClassWifiSettings ClassDef = 9

// title : Set wifi outdoor mode,
// desc : Set wifi indoor/outdoor mode.\n **Please note that you might be disconnected from the product after changing the indoor/outdoor setting as it changes Wifi parameters.**,
// support : 0901;0902;0905;0906;090c;090e,
// result : The product change its indoor/outdoor wifi settings.\n Then, it will trigger [WifiOutdoorMode](#0-10-0).,
const CmdOutdoorSetting CmdDef = 0

type CommonWifiSettingsOutdoorSetting Command

type CommonWifiSettingsOutdoorSettingArguments struct {
	Outdoor uint8
}

func (a CommonWifiSettingsOutdoorSetting) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonWifiSettingsOutdoorSettingArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Outdoor)
	offset++

	return arg
}
func (a CommonWifiSettingsOutdoorSetting) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var WifiSettingsOutdoorSetting = CommonWifiSettingsOutdoorSetting{
	Project: ProjectCommon,
	Class:   ClassWifiSettings,
	Cmd:     CmdOutdoorSetting,
}

// Wifi settings state from product
const ClassWifiSettingsState ClassDef = 10

// title : Wifi outdoor mode,
// desc : Wifi outdoor mode.,
// support : 0901;0902;0905;0906;090c;090e,
// triggered : by [SetWifiOutdoorMode](#0-9-0).,
const CmdOutdoorSettingsChanged CmdDef = 0

type CommonWifiSettingsStateoutdoorSettingsChanged Command

type CommonWifiSettingsStateoutdoorSettingsChangedArguments struct {
	Outdoor uint8
}

func (a CommonWifiSettingsStateoutdoorSettingsChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonWifiSettingsStateoutdoorSettingsChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Outdoor)
	offset++

	return arg
}
func (a CommonWifiSettingsStateoutdoorSettingsChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var WifiSettingsStateoutdoorSettingsChanged = CommonWifiSettingsStateoutdoorSettingsChanged{
	Project: ProjectCommon,
	Class:   ClassWifiSettingsState,
	Cmd:     CmdOutdoorSettingsChanged,
}

// Mavlink flight plans commands
const ClassMavlink ClassDef = 11

// title : Start a FlightPlan,
// desc : Start a FlightPlan based on a mavlink file existing on the drone.\n\n Requirements are:\n * Product is calibrated\n * Product should be in outdoor mode\n * Product has fixed its GPS\n,
// support : 0901:2.0.29;090c;090e,
// result : If the FlightPlan has been started, event [FlightPlanPlayingStateChanged](#0-12-0) is triggered with param state set to *playing*.\n Otherwise, event [FlightPlanPlayingStateChanged](#0-12-0) is triggered with param state set to stopped and event [MavlinkPlayErrorStateChanged](#0-12-1) is triggered with an explanation of the error.,
const CmdStart CmdDef = 0

type CommonMavlinkStart Command

type CommonMavlinkStartArguments struct {
	Filepath string
	TypeX    uint32
}

func (a CommonMavlinkStart) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonMavlinkStartArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Filepath = string(b[offset : offset+stringEnd])
	offset += stringEnd
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TypeX)
	offset += 4

	return arg
}
func (a CommonMavlinkStart) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MavlinkStart = CommonMavlinkStart{
	Project: ProjectCommon,
	Class:   ClassMavlink,
	Cmd:     CmdStart,
}

// title : Pause a FlightPlan,
// desc : Pause a FlightPlan that was playing.\n To unpause a FlightPlan, see [StartFlightPlan](#0-11-0)\n,
// support : 0901:2.0.29;090c;090e,
// result : The currently playing FlightPlan will be paused. Then, event [FlightPlanPlayingStateChanged](#0-12-0) is triggered with param state set to the current state of the FlightPlan (should be *paused* if everything went well).,
const CmdPause CmdDef = 1

type CommonMavlinkPause Command

type CommonMavlinkPauseArguments struct {
}

func (a CommonMavlinkPause) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonMavlinkPauseArguments{}
	// No arguments to decode here !!

	return arg
}
func (a CommonMavlinkPause) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MavlinkPause = CommonMavlinkPause{
	Project: ProjectCommon,
	Class:   ClassMavlink,
	Cmd:     CmdPause,
}

// title : Stop a FlightPlan,
// desc : Stop a FlightPlan that was playing.\n,
// support : 0901:2.0.29;090c;090e,
// result : The currently playing FlightPlan will be stopped. Then, event [FlightPlanPlayingStateChanged](#0-12-0) is triggered with param state set to the current state of the FlightPlan (should be *stopped* if everything went well).,
const CmdStop CmdDef = 2

type CommonMavlinkStop Command

type CommonMavlinkStopArguments struct {
}

func (a CommonMavlinkStop) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonMavlinkStopArguments{}
	// No arguments to decode here !!

	return arg
}
func (a CommonMavlinkStop) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MavlinkStop = CommonMavlinkStop{
	Project: ProjectCommon,
	Class:   ClassMavlink,
	Cmd:     CmdStop,
}

// Mavlink flight plans states commands
const ClassMavlinkState ClassDef = 12

// title : Playing state of a FlightPlan,
// desc : Playing state of a FlightPlan.,
// support : 0901:2.0.29;090c;090e,
// triggered : by [StartFlightPlan](#0-11-0), [PauseFlightPlan](#0-11-1) or [StopFlightPlan](#0-11-2).,
const CmdMavlinkFilePlayingStateChanged CmdDef = 0

type CommonMavlinkStateMavlinkFilePlayingStateChanged Command

type CommonMavlinkStateMavlinkFilePlayingStateChangedArguments struct {
	State    uint32
	Filepath string
	TypeX    uint32
}

func (a CommonMavlinkStateMavlinkFilePlayingStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonMavlinkStateMavlinkFilePlayingStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.State)
	offset += 4

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Filepath = string(b[offset : offset+stringEnd])
	offset += stringEnd
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TypeX)
	offset += 4

	return arg
}
func (a CommonMavlinkStateMavlinkFilePlayingStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MavlinkStateMavlinkFilePlayingStateChanged = CommonMavlinkStateMavlinkFilePlayingStateChanged{
	Project: ProjectCommon,
	Class:   ClassMavlinkState,
	Cmd:     CmdMavlinkFilePlayingStateChanged,
}

// title : FlightPlan error,
// desc : FlightPlan error.,
// support : 0901:2.0.29;090c;090e,
// triggered : by [StartFlightPlan](#0-11-0) if an error occurs.,
const CmdMavlinkPlayErrorStateChanged CmdDef = 1

type CommonMavlinkStateMavlinkPlayErrorStateChanged Command

type CommonMavlinkStateMavlinkPlayErrorStateChangedArguments struct {
	Error uint32
}

func (a CommonMavlinkStateMavlinkPlayErrorStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonMavlinkStateMavlinkPlayErrorStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Error)
	offset += 4

	return arg
}
func (a CommonMavlinkStateMavlinkPlayErrorStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MavlinkStateMavlinkPlayErrorStateChanged = CommonMavlinkStateMavlinkPlayErrorStateChanged{
	Project: ProjectCommon,
	Class:   ClassMavlinkState,
	Cmd:     CmdMavlinkPlayErrorStateChanged,
}

// title : Mission item executed,
// desc : Mission item has been executed.,
// support : 090c:4.2.0;090e:1.4.0,
// triggered : when a mission item has been executed during a flight plan.,
const CmdMissionItemExecuted CmdDef = 2

type CommonMavlinkStateMissionItemExecuted Command

type CommonMavlinkStateMissionItemExecutedArguments struct {
	Idx uint32
}

func (a CommonMavlinkStateMissionItemExecuted) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonMavlinkStateMissionItemExecutedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Idx)
	offset += 4

	return arg
}
func (a CommonMavlinkStateMissionItemExecuted) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var MavlinkStateMissionItemExecuted = CommonMavlinkStateMissionItemExecuted{
	Project: ProjectCommon,
	Class:   ClassMavlinkState,
	Cmd:     CmdMissionItemExecuted,
}

const ClassFlightPlanSettings ClassDef = 32

// title : Set ReturnHome behavior during FlightPlan,
// desc : Set ReturnHome behavior during FlightPlan\n When set, drone will return home, after return home delay, if a disconnection occurs during execution of FlightPlan,
// support : 0901:4.1.0;090c:4.1.0;090e:1.4.0,
// result : The return home mode is enabled or disabled.\n Then, event [ReturnHomeOnDisconnectionChanged](#0-33-0) is triggered.,
const CmdReturnHomeOnDisconnect CmdDef = 0

type CommonFlightPlanSettingsReturnHomeOnDisconnect Command

type CommonFlightPlanSettingsReturnHomeOnDisconnectArguments struct {
	Value uint8
}

func (a CommonFlightPlanSettingsReturnHomeOnDisconnect) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonFlightPlanSettingsReturnHomeOnDisconnectArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Value)
	offset++

	return arg
}
func (a CommonFlightPlanSettingsReturnHomeOnDisconnect) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var FlightPlanSettingsReturnHomeOnDisconnect = CommonFlightPlanSettingsReturnHomeOnDisconnect{
	Project: ProjectCommon,
	Class:   ClassFlightPlanSettings,
	Cmd:     CmdReturnHomeOnDisconnect,
}

const ClassFlightPlanSettingsState ClassDef = 33

// title : ReturnHome behavior during FlightPlan,
// desc : Define behavior of drone when disconnection occurs during a flight plan,
// support : 0901:4.1.0;090c:4.1.0;090e:1.4.0,
// triggered : by [setReturnHomeOnDisconnectMode](#0-32-0).,
const CmdReturnHomeOnDisconnectChanged CmdDef = 0

type CommonFlightPlanSettingsStateReturnHomeOnDisconnectChanged Command

type CommonFlightPlanSettingsStateReturnHomeOnDisconnectChangedArguments struct {
	State      uint8
	IsReadOnly uint8
}

func (a CommonFlightPlanSettingsStateReturnHomeOnDisconnectChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonFlightPlanSettingsStateReturnHomeOnDisconnectChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.State)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.IsReadOnly)
	offset++

	return arg
}
func (a CommonFlightPlanSettingsStateReturnHomeOnDisconnectChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var FlightPlanSettingsStateReturnHomeOnDisconnectChanged = CommonFlightPlanSettingsStateReturnHomeOnDisconnectChanged{
	Project: ProjectCommon,
	Class:   ClassFlightPlanSettingsState,
	Cmd:     CmdReturnHomeOnDisconnectChanged,
}

// Calibration commands
const ClassCalibration ClassDef = 13

// title : Start/Abort magnetometer calibration,
// desc : Start or abort magnetometer calibration process.\n,
// support : 0901;090c;090e,
// result : The magnetometer calibration process is started or aborted. Then, event [MagnetoCalibrationStartedChanged](#0-14-3) is triggered.\n If started, event [MagnetoCalibrationStateChanged](#0-14-3) is triggered with the current calibration state: a list of all axis and their calibration states.\n It will also trigger [MagnetoCalibrationAxisToCalibrateChanged](#0-14-2), that will inform the controller about the current axis to calibrate.,
const CmdMagnetoCalibration CmdDef = 0

type CommonCalibrationMagnetoCalibration Command

type CommonCalibrationMagnetoCalibrationArguments struct {
	Calibrate uint8
}

func (a CommonCalibrationMagnetoCalibration) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCalibrationMagnetoCalibrationArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Calibrate)
	offset++

	return arg
}
func (a CommonCalibrationMagnetoCalibration) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CalibrationMagnetoCalibration = CommonCalibrationMagnetoCalibration{
	Project: ProjectCommon,
	Class:   ClassCalibration,
	Cmd:     CmdMagnetoCalibration,
}

// title : Start/Abort Pitot calibration,
// desc : Start or abort Pitot tube calibration process.\n,
// support : 090e:1.1.0,
// result : The pitot calibration process is started or aborted. Then, event [PitotCalibrationStateChanged](#0-14-4) is triggered with the current calibration state.,
const CmdPitotCalibration CmdDef = 1

type CommonCalibrationPitotCalibration Command

type CommonCalibrationPitotCalibrationArguments struct {
	Calibrate uint8
}

func (a CommonCalibrationPitotCalibration) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCalibrationPitotCalibrationArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Calibrate)
	offset++

	return arg
}
func (a CommonCalibrationPitotCalibration) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CalibrationPitotCalibration = CommonCalibrationPitotCalibration{
	Project: ProjectCommon,
	Class:   ClassCalibration,
	Cmd:     CmdPitotCalibration,
}

// Status of the calibration
const ClassCalibrationState ClassDef = 14

// title : Magneto calib process axis state,
// desc : Magneto calib process axis state.,
// support : 0901;090c;090e,
// triggered : when the calibration process is started with [StartOrAbortMagnetoCalib](#0-13-0) and each time an axis calibration state changes.,
const CmdMagnetoCalibrationStateChanged CmdDef = 0

type CommonCalibrationStateMagnetoCalibrationStateChanged Command

type CommonCalibrationStateMagnetoCalibrationStateChangedArguments struct {
	XAxisCalibration  uint8
	YAxisCalibration  uint8
	ZAxisCalibration  uint8
	CalibrationFailed uint8
}

func (a CommonCalibrationStateMagnetoCalibrationStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCalibrationStateMagnetoCalibrationStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.XAxisCalibration)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.YAxisCalibration)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.ZAxisCalibration)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.CalibrationFailed)
	offset++

	return arg
}
func (a CommonCalibrationStateMagnetoCalibrationStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CalibrationStateMagnetoCalibrationStateChanged = CommonCalibrationStateMagnetoCalibrationStateChanged{
	Project: ProjectCommon,
	Class:   ClassCalibrationState,
	Cmd:     CmdMagnetoCalibrationStateChanged,
}

// title : Calibration required,
// desc : Calibration required.,
// support : 0901;090c;090e,
// triggered : when the calibration requirement changes.,
const CmdMagnetoCalibrationRequiredState CmdDef = 1

type CommonCalibrationStateMagnetoCalibrationRequiredState Command

type CommonCalibrationStateMagnetoCalibrationRequiredStateArguments struct {
	Required uint8
}

func (a CommonCalibrationStateMagnetoCalibrationRequiredState) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCalibrationStateMagnetoCalibrationRequiredStateArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Required)
	offset++

	return arg
}
func (a CommonCalibrationStateMagnetoCalibrationRequiredState) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CalibrationStateMagnetoCalibrationRequiredState = CommonCalibrationStateMagnetoCalibrationRequiredState{
	Project: ProjectCommon,
	Class:   ClassCalibrationState,
	Cmd:     CmdMagnetoCalibrationRequiredState,
}

// title : Axis to calibrate during calibration process,
// desc : Axis to calibrate during calibration process.,
// support : 0901;090c;090e,
// triggered : during the calibration process when the axis to calibrate changes.,
const CmdMagnetoCalibrationAxisToCalibrateChanged CmdDef = 2

type CommonCalibrationStateMagnetoCalibrationAxisToCalibrateChanged Command

type CommonCalibrationStateMagnetoCalibrationAxisToCalibrateChangedArguments struct {
	Axis uint32
}

func (a CommonCalibrationStateMagnetoCalibrationAxisToCalibrateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCalibrationStateMagnetoCalibrationAxisToCalibrateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Axis)
	offset += 4

	return arg
}
func (a CommonCalibrationStateMagnetoCalibrationAxisToCalibrateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CalibrationStateMagnetoCalibrationAxisToCalibrateChanged = CommonCalibrationStateMagnetoCalibrationAxisToCalibrateChanged{
	Project: ProjectCommon,
	Class:   ClassCalibrationState,
	Cmd:     CmdMagnetoCalibrationAxisToCalibrateChanged,
}

// title : Calibration process state,
// desc : Calibration process state.,
// support : 0901;090c;090e,
// triggered : by [StartOrAbortMagnetoCalib](#0-13-0) or when the process ends because it succeeded.,
const CmdMagnetoCalibrationStartedChanged CmdDef = 3

type CommonCalibrationStateMagnetoCalibrationStartedChanged Command

type CommonCalibrationStateMagnetoCalibrationStartedChangedArguments struct {
	Started uint8
}

func (a CommonCalibrationStateMagnetoCalibrationStartedChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCalibrationStateMagnetoCalibrationStartedChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Started)
	offset++

	return arg
}
func (a CommonCalibrationStateMagnetoCalibrationStartedChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CalibrationStateMagnetoCalibrationStartedChanged = CommonCalibrationStateMagnetoCalibrationStartedChanged{
	Project: ProjectCommon,
	Class:   ClassCalibrationState,
	Cmd:     CmdMagnetoCalibrationStartedChanged,
}

const CmdPitotCalibrationStateChanged CmdDef = 4

type CommonCalibrationStatePitotCalibrationStateChanged Command

type CommonCalibrationStatePitotCalibrationStateChangedArguments struct {
	State     uint32
	LastError uint8
}

func (a CommonCalibrationStatePitotCalibrationStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCalibrationStatePitotCalibrationStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.State)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.LastError)
	offset++

	return arg
}
func (a CommonCalibrationStatePitotCalibrationStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CalibrationStatePitotCalibrationStateChanged = CommonCalibrationStatePitotCalibrationStateChanged{
	Project: ProjectCommon,
	Class:   ClassCalibrationState,
	Cmd:     CmdPitotCalibrationStateChanged,
}

// Status of the camera settings
const ClassCameraSettingsState ClassDef = 15

// title : Camera info,
// desc : Camera info.,
// support : 0901;090c;090e,
// triggered : at connection.,
const CmdCameraSettingsChanged CmdDef = 0

type CommonCameraSettingsStateCameraSettingsChanged Command

type CommonCameraSettingsStateCameraSettingsChangedArguments struct {
	Fov     float32
	PanMax  float32
	PanMin  float32
	TiltMax float32
	TiltMin float32
}

func (a CommonCameraSettingsStateCameraSettingsChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonCameraSettingsStateCameraSettingsChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Fov)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.PanMax)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.PanMin)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TiltMax)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.TiltMin)
	offset += 4

	return arg
}
func (a CommonCameraSettingsStateCameraSettingsChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var CameraSettingsStateCameraSettingsChanged = CommonCameraSettingsStateCameraSettingsChanged{
	Project: ProjectCommon,
	Class:   ClassCameraSettingsState,
	Cmd:     CmdCameraSettingsChanged,
}

// GPS related commands
const ClassGPS ClassDef = 16

// title : Set the position of a run,
// desc : Set the position of a run.\n This will let the product know the controller location for the flight/run. The location is typically used to geotag medias.\n Only used on products that have no gps.\n Watch out, this command is not used by BLE products.,
// support : 0902;0905;0906,
// result : The position is set.,
const CmdControllerPositionForRun CmdDef = 0

type CommonGPSControllerPositionForRun Command

type CommonGPSControllerPositionForRunArguments struct {
	Latitude  float64
	Longitude float64
}

func (a CommonGPSControllerPositionForRun) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonGPSControllerPositionForRunArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Latitude)
	offset += 8
	ConvLittleEndianSliceToNumeric(b[offset:offset+8], &arg.Longitude)
	offset += 8

	return arg
}
func (a CommonGPSControllerPositionForRun) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var GPSControllerPositionForRun = CommonGPSControllerPositionForRun{
	Project: ProjectCommon,
	Class:   ClassGPS,
	Cmd:     CmdControllerPositionForRun,
}

// FlightPlan state commands
const ClassFlightPlanState ClassDef = 17

// title : FlightPlan availability,
// desc : FlightPlan availability.\n Availability is linked to GPS fix, magnetometer calibration, sensor states...,
// support : 0901:2.0.29;090c;090e,
// triggered : on change.,
const CmdAvailabilityStateChanged CmdDef = 0

type CommonFlightPlanStateAvailabilityStateChanged Command

type CommonFlightPlanStateAvailabilityStateChangedArguments struct {
	AvailabilityState uint8
}

func (a CommonFlightPlanStateAvailabilityStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonFlightPlanStateAvailabilityStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.AvailabilityState)
	offset++

	return arg
}
func (a CommonFlightPlanStateAvailabilityStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var FlightPlanStateAvailabilityStateChanged = CommonFlightPlanStateAvailabilityStateChanged{
	Project: ProjectCommon,
	Class:   ClassFlightPlanState,
	Cmd:     CmdAvailabilityStateChanged,
}

// title : FlightPlan components state list,
// desc : FlightPlan components state list.,
// support : 0901:2.0.29;090c;090e,
// triggered : when the state of required components changes. \n GPS component is triggered when the availability of the GPS of the drone changes. \n Calibration component is triggered when the calibration state of the drone sensors changes \n Mavlink_File component is triggered when the command [StartFlightPlan](#0-11-0) is received. \n Takeoff component is triggered when the drone needs to take-off to continue the FlightPlan. \n WaypointsBeyondGeofence component is triggered when the command [StartFlightPlan](#0-11-0) is received.,
const CmdComponentStateListChanged CmdDef = 1

type CommonFlightPlanStateComponentStateListChanged Command

type CommonFlightPlanStateComponentStateListChangedArguments struct {
	Component uint32
	State     uint8
}

func (a CommonFlightPlanStateComponentStateListChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonFlightPlanStateComponentStateListChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Component)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.State)
	offset++

	return arg
}
func (a CommonFlightPlanStateComponentStateListChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var FlightPlanStateComponentStateListChanged = CommonFlightPlanStateComponentStateListChanged{
	Project: ProjectCommon,
	Class:   ClassFlightPlanState,
	Cmd:     CmdComponentStateListChanged,
}

// title : FlightPlan lock state,
// desc : FlightPlan lock state.\n Represents the fact that the controller is able or not to stop or pause a playing FlightPlan,
// support : 0901:2.0.29;090c;090e,
// triggered : when the lock changes.,
const CmdLockStateChanged CmdDef = 2

type CommonFlightPlanStateLockStateChanged Command

type CommonFlightPlanStateLockStateChangedArguments struct {
	LockState uint8
}

func (a CommonFlightPlanStateLockStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonFlightPlanStateLockStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.LockState)
	offset++

	return arg
}
func (a CommonFlightPlanStateLockStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var FlightPlanStateLockStateChanged = CommonFlightPlanStateLockStateChanged{
	Project: ProjectCommon,
	Class:   ClassFlightPlanState,
	Cmd:     CmdLockStateChanged,
}

// FlightPlan Event commands
const ClassFlightPlanEvent ClassDef = 19

// title : FlightPlan start error,
// desc : FlightPlan start error.\n\n **This event is a notification, you can't retrieve it in the cache of the device controller.**,
// support : 0901:2.0.29;090c;090e,
// triggered : on an error after a [StartFlightPlan](#0-11-0).,
const CmdStartingErrorEvent CmdDef = 0

type CommonFlightPlanEventStartingErrorEvent Command

type CommonFlightPlanEventStartingErrorEventArguments struct {
}

func (a CommonFlightPlanEventStartingErrorEvent) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonFlightPlanEventStartingErrorEventArguments{}
	// No arguments to decode here !!

	return arg
}
func (a CommonFlightPlanEventStartingErrorEvent) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var FlightPlanEventStartingErrorEvent = CommonFlightPlanEventStartingErrorEvent{
	Project: ProjectCommon,
	Class:   ClassFlightPlanEvent,
	Cmd:     CmdStartingErrorEvent,
}

// title : FlightPlan speed clamping,
// desc : FlightPlan speed clamping.\n Sent when a speed specified in the FlightPlan file is considered too high by the drone.\n\n **This event is a notification, you can't retrieve it in the cache of the device controller.**,
// support : none,
// triggered : on an speed related clamping after a [StartFlightPlan](#0-11-0).,
const CmdSpeedBridleEvent CmdDef = 1

type CommonFlightPlanEventSpeedBridleEvent Command

type CommonFlightPlanEventSpeedBridleEventArguments struct {
}

func (a CommonFlightPlanEventSpeedBridleEvent) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonFlightPlanEventSpeedBridleEventArguments{}
	// No arguments to decode here !!

	return arg
}
func (a CommonFlightPlanEventSpeedBridleEvent) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var FlightPlanEventSpeedBridleEvent = CommonFlightPlanEventSpeedBridleEvent{
	Project: ProjectCommon,
	Class:   ClassFlightPlanEvent,
	Cmd:     CmdSpeedBridleEvent,
}

// ARlibs Versions Commands
const ClassARLibsVersionsState ClassDef = 18
const CmdControllerLibARCommandsVersion CmdDef = 0

type CommonARLibsVersionsStateControllerLibARCommandsVersion Command

type CommonARLibsVersionsStateControllerLibARCommandsVersionArguments struct {
	Version string
}

func (a CommonARLibsVersionsStateControllerLibARCommandsVersion) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonARLibsVersionsStateControllerLibARCommandsVersionArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Version = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonARLibsVersionsStateControllerLibARCommandsVersion) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var ARLibsVersionsStateControllerLibARCommandsVersion = CommonARLibsVersionsStateControllerLibARCommandsVersion{
	Project: ProjectCommon,
	Class:   ClassARLibsVersionsState,
	Cmd:     CmdControllerLibARCommandsVersion,
}

const CmdSkyControllerLibARCommandsVersion CmdDef = 1

type CommonARLibsVersionsStateSkyControllerLibARCommandsVersion Command

type CommonARLibsVersionsStateSkyControllerLibARCommandsVersionArguments struct {
	Version string
}

func (a CommonARLibsVersionsStateSkyControllerLibARCommandsVersion) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonARLibsVersionsStateSkyControllerLibARCommandsVersionArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Version = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonARLibsVersionsStateSkyControllerLibARCommandsVersion) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var ARLibsVersionsStateSkyControllerLibARCommandsVersion = CommonARLibsVersionsStateSkyControllerLibARCommandsVersion{
	Project: ProjectCommon,
	Class:   ClassARLibsVersionsState,
	Cmd:     CmdSkyControllerLibARCommandsVersion,
}

const CmdDeviceLibARCommandsVersion CmdDef = 2

type CommonARLibsVersionsStateDeviceLibARCommandsVersion Command

type CommonARLibsVersionsStateDeviceLibARCommandsVersionArguments struct {
	Version string
}

func (a CommonARLibsVersionsStateDeviceLibARCommandsVersion) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonARLibsVersionsStateDeviceLibARCommandsVersionArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.Version = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonARLibsVersionsStateDeviceLibARCommandsVersion) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var ARLibsVersionsStateDeviceLibARCommandsVersion = CommonARLibsVersionsStateDeviceLibARCommandsVersion{
	Project: ProjectCommon,
	Class:   ClassARLibsVersionsState,
	Cmd:     CmdDeviceLibARCommandsVersion,
}

// Audio-related commands.
const ClassAudio ClassDef = 20

// title : Set audio stream direction,
// desc : Set audio stream direction.,
// support : 0905;0906,
// result : The audio stream direction is set.\n Then, event [AudioStreamDirection](#0-21-0) is triggered.,
const CmdControllerReadyForStreaming CmdDef = 0

type CommonAudioControllerReadyForStreaming Command

type CommonAudioControllerReadyForStreamingArguments struct {
	Ready uint8
}

func (a CommonAudioControllerReadyForStreaming) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonAudioControllerReadyForStreamingArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Ready)
	offset++

	return arg
}
func (a CommonAudioControllerReadyForStreaming) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var AudioControllerReadyForStreaming = CommonAudioControllerReadyForStreaming{
	Project: ProjectCommon,
	Class:   ClassAudio,
	Cmd:     CmdControllerReadyForStreaming,
}

// Audio-related state updates.
const ClassAudioState ClassDef = 21

// title : Audio stream direction,
// desc : Audio stream direction.,
// support : 0905;0906,
// triggered : by [SetAudioStreamDirection](#0-20-0).,
const CmdAudioStreamingRunning CmdDef = 0

type CommonAudioStateAudioStreamingRunning Command

type CommonAudioStateAudioStreamingRunningArguments struct {
	Running uint8
}

func (a CommonAudioStateAudioStreamingRunning) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonAudioStateAudioStreamingRunningArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Running)
	offset++

	return arg
}
func (a CommonAudioStateAudioStreamingRunning) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var AudioStateAudioStreamingRunning = CommonAudioStateAudioStreamingRunning{
	Project: ProjectCommon,
	Class:   ClassAudioState,
	Cmd:     CmdAudioStreamingRunning,
}

// Controls the headlight LEDs of the Evo variants.
const ClassHeadlights ClassDef = 22

// title : Set LEDs intensity,
// desc : Set lighting LEDs intensity.,
// support : 0905;0906;0907,
// result : The intensity of the LEDs is changed.\n Then, event [LedIntensity](#0-23-0) is triggered.,
const CmdIntensity CmdDef = 0

type CommonHeadlightsintensity Command

type CommonHeadlightsintensityArguments struct {
	Left  uint8
	Right uint8
}

func (a CommonHeadlightsintensity) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonHeadlightsintensityArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Left)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Right)
	offset++

	return arg
}
func (a CommonHeadlightsintensity) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var Headlightsintensity = CommonHeadlightsintensity{
	Project: ProjectCommon,
	Class:   ClassHeadlights,
	Cmd:     CmdIntensity,
}

// Get information about the state of the Evo variants' LEDs.
const ClassHeadlightsState ClassDef = 23

// title : LEDs intensity,
// desc : Lighting LEDs intensity.,
// support : 0905;0906;0907,
// triggered : by [SetLedsIntensity](#0-22-0).,
const CmdIntensityChanged CmdDef = 0

type CommonHeadlightsStateintensityChanged Command

type CommonHeadlightsStateintensityChangedArguments struct {
	Left  uint8
	Right uint8
}

func (a CommonHeadlightsStateintensityChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonHeadlightsStateintensityChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Left)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Right)
	offset++

	return arg
}
func (a CommonHeadlightsStateintensityChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var HeadlightsStateintensityChanged = CommonHeadlightsStateintensityChanged{
	Project: ProjectCommon,
	Class:   ClassHeadlightsState,
	Cmd:     CmdIntensityChanged,
}

// Animations-related commands.
const ClassAnimationsDUPLICATE ClassDef = 24

// title : Start an animation,
// desc : Start a paramaterless animation.\n List of available animations can be retrieved from [AnimationsStateList](#0-25-0).,
// support : 0902;0905;0906;0907;0909,
// result : If possible, the product starts the requested animation. Then, event [AnimationsStateList](#0-25-0) is triggered.,
const CmdStartAnimation CmdDef = 0

type CommonAnimationsStartAnimation Command

type CommonAnimationsStartAnimationArguments struct {
	Anim uint32
}

func (a CommonAnimationsStartAnimation) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonAnimationsStartAnimationArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Anim)
	offset += 4

	return arg
}
func (a CommonAnimationsStartAnimation) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var AnimationsStartAnimation = CommonAnimationsStartAnimation{
	Project: ProjectCommon,
	Class:   ClassAnimationsDUPLICATE,
	Cmd:     CmdStartAnimation,
}

// title : Stop an animation,
// desc : Stop a paramaterless animation.\n List of running animations can be retrieved from [AnimationsStateList](#0-25-0).,
// support : 0902;0905;0906;0907;0909,
// result : If the requested animation was running, it will be stopped.\n Then, event [AnimationsStateList](#0-25-0) is triggered.,
const CmdStopAnimation CmdDef = 1

type CommonAnimationsStopAnimation Command

type CommonAnimationsStopAnimationArguments struct {
	Anim uint32
}

func (a CommonAnimationsStopAnimation) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonAnimationsStopAnimationArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Anim)
	offset += 4

	return arg
}
func (a CommonAnimationsStopAnimation) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var AnimationsStopAnimation = CommonAnimationsStopAnimation{
	Project: ProjectCommon,
	Class:   ClassAnimations,
	Cmd:     CmdStopAnimation,
}

// title : Stop all animations,
// desc : Stop all running paramaterless animations.\n List of running animations can be retrieved from [AnimationsStateList](#0-25-0).,
// support : 0902;0905;0906;0907;0909,
// result : All running animations are stopped.\n Then, event [AnimationsStateList](#0-25-0) is triggered.,
const CmdStopAllAnimations CmdDef = 2

type CommonAnimationsStopAllAnimations Command

type CommonAnimationsStopAllAnimationsArguments struct {
}

func (a CommonAnimationsStopAllAnimations) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonAnimationsStopAllAnimationsArguments{}
	// No arguments to decode here !!

	return arg
}
func (a CommonAnimationsStopAllAnimations) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var AnimationsStopAllAnimations = CommonAnimationsStopAllAnimations{
	Project: ProjectCommon,
	Class:   ClassAnimations,
	Cmd:     CmdStopAllAnimations,
}

// Animations-related notification/feedback commands.
const ClassAnimationsState ClassDef = 25

// title : Animation state list,
// desc : Paramaterless animations state list.,
// support : 0902;0905;0906;0907;0909,
// triggered : when the list of available animations changes and also when an animation state changes (can be triggered by [StartAnim](#0-24-0), [StopAnim](#0-24-1) or [StopAllAnims](#0-24-2).,
const CmdList CmdDef = 0

type CommonAnimationsStateList Command

type CommonAnimationsStateListArguments struct {
	Anim uint32
}

func (a CommonAnimationsStateList) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonAnimationsStateListArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Anim)
	offset += 4

	return arg
}
func (a CommonAnimationsStateList) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var AnimationsStateList = CommonAnimationsStateList{
	Project: ProjectCommon,
	Class:   ClassAnimationsState,
	Cmd:     CmdList,
}

// Accessories-related commands.
const ClassAccessory ClassDef = 26

// title : Declare an accessory,
// desc : Declare an accessory.\n You can choose the accessory between all accessible for this product.\n You can get this list through event [SupportedAccessories](#0-27-0).\n\n You can only set the accessory when the modification is enabled.\n You can know if it possible with the event [AccessoryDeclarationAvailability](#0-27-2).,
// support : 0902;0905;0906;0907;0909;090a,
// result : The product knows which accessory it is wearing.\n Then, event [AccessoryConfigChanged](#0-27-1) is triggered.,
const CmdConfig CmdDef = 0

type CommonAccessoryConfig Command

type CommonAccessoryConfigArguments struct {
	Accessory uint32
}

func (a CommonAccessoryConfig) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonAccessoryConfigArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Accessory)
	offset += 4

	return arg
}
func (a CommonAccessoryConfig) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var AccessoryConfig = CommonAccessoryConfig{
	Project: ProjectCommon,
	Class:   ClassAccessory,
	Cmd:     CmdConfig,
}

// Accessories-related commands.
const ClassAccessoryStateDUPLICATE ClassDef = 27

// title : Supported accessories list,
// desc : Supported accessories list.,
// support : 0902;0905;0906;0907;0909;090a,
// triggered : at connection.,
const CmdSupportedAccessoriesListChanged CmdDef = 0

type CommonAccessoryStateSupportedAccessoriesListChanged Command

type CommonAccessoryStateSupportedAccessoriesListChangedArguments struct {
	Accessory uint32
}

func (a CommonAccessoryStateSupportedAccessoriesListChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonAccessoryStateSupportedAccessoriesListChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Accessory)
	offset += 4

	return arg
}
func (a CommonAccessoryStateSupportedAccessoriesListChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var AccessoryStateSupportedAccessoriesListChanged = CommonAccessoryStateSupportedAccessoriesListChanged{
	Project: ProjectCommon,
	Class:   ClassAccessoryStateDUPLICATE,
	Cmd:     CmdSupportedAccessoriesListChanged,
}

// title : Accessory config,
// desc : Accessory config.,
// support : 0902;0905;0906;0907;0909;090a,
// triggered : by [DeclareAccessory](#0-26-0).,
const CmdAccessoryConfigChanged CmdDef = 1

type CommonAccessoryStateAccessoryConfigChanged Command

type CommonAccessoryStateAccessoryConfigChangedArguments struct {
	NewAccessory uint32
	Error        uint32
}

func (a CommonAccessoryStateAccessoryConfigChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonAccessoryStateAccessoryConfigChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.NewAccessory)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Error)
	offset += 4

	return arg
}
func (a CommonAccessoryStateAccessoryConfigChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var AccessoryStateAccessoryConfigChanged = CommonAccessoryStateAccessoryConfigChanged{
	Project: ProjectCommon,
	Class:   ClassAccessoryState,
	Cmd:     CmdAccessoryConfigChanged,
}

// title : Accessory declaration availability,
// desc : Availability to declare or not an accessory.,
// support : 0902;0905;0906;0907;0909;090a,
// triggered : when the availability changes.,
const CmdAccessoryConfigModificationEnabled CmdDef = 2

type CommonAccessoryStateAccessoryConfigModificationEnabled Command

type CommonAccessoryStateAccessoryConfigModificationEnabledArguments struct {
	Enabled uint8
}

func (a CommonAccessoryStateAccessoryConfigModificationEnabled) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonAccessoryStateAccessoryConfigModificationEnabledArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Enabled)
	offset++

	return arg
}
func (a CommonAccessoryStateAccessoryConfigModificationEnabled) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var AccessoryStateAccessoryConfigModificationEnabled = CommonAccessoryStateAccessoryConfigModificationEnabled{
	Project: ProjectCommon,
	Class:   ClassAccessoryState,
	Cmd:     CmdAccessoryConfigModificationEnabled,
}

// Commands sent by the controller to set charger parameters.
const ClassCharger ClassDef = 28

// title : Set max charge rate,
// desc : The product will inform itself the controller about its charging type (see [ChargingInfoChanged](#0-29-3)).,
// support : none,
// result : None.,
const CmdSetMaxChargeRate CmdDef = 0

type CommonChargerSetMaxChargeRate Command

type CommonChargerSetMaxChargeRateArguments struct {
	Rate uint32
}

func (a CommonChargerSetMaxChargeRate) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonChargerSetMaxChargeRateArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Rate)
	offset += 4

	return arg
}
func (a CommonChargerSetMaxChargeRate) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var ChargerSetMaxChargeRate = CommonChargerSetMaxChargeRate{
	Project: ProjectCommon,
	Class:   ClassCharger,
	Cmd:     CmdSetMaxChargeRate,
}

// Commands sent by the firmware to advertise the charger status.
const ClassChargerState ClassDef = 29

// title : Max charge rate,
// desc : Max charge rate.,
const CmdMaxChargeRateChanged CmdDef = 0

type CommonChargerStateMaxChargeRateChanged Command

type CommonChargerStateMaxChargeRateChangedArguments struct {
	Rate uint32
}

func (a CommonChargerStateMaxChargeRateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonChargerStateMaxChargeRateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Rate)
	offset += 4

	return arg
}
func (a CommonChargerStateMaxChargeRateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var ChargerStateMaxChargeRateChanged = CommonChargerStateMaxChargeRateChanged{
	Project: ProjectCommon,
	Class:   ClassChargerState,
	Cmd:     CmdMaxChargeRateChanged,
}

// title : Current charge state,
// desc : Current charge state.,
const CmdCurrentChargeStateChanged CmdDef = 1

type CommonChargerStateCurrentChargeStateChanged Command

type CommonChargerStateCurrentChargeStateChangedArguments struct {
	Status uint32
	Phase  uint32
}

func (a CommonChargerStateCurrentChargeStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonChargerStateCurrentChargeStateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Status)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Phase)
	offset += 4

	return arg
}
func (a CommonChargerStateCurrentChargeStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var ChargerStateCurrentChargeStateChanged = CommonChargerStateCurrentChargeStateChanged{
	Project: ProjectCommon,
	Class:   ClassChargerState,
	Cmd:     CmdCurrentChargeStateChanged,
}

// title : Last charge rate,
// desc : Last charge rate.,
const CmdLastChargeRateChanged CmdDef = 2

type CommonChargerStateLastChargeRateChanged Command

type CommonChargerStateLastChargeRateChangedArguments struct {
	Rate uint32
}

func (a CommonChargerStateLastChargeRateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonChargerStateLastChargeRateChangedArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Rate)
	offset += 4

	return arg
}
func (a CommonChargerStateLastChargeRateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var ChargerStateLastChargeRateChanged = CommonChargerStateLastChargeRateChanged{
	Project: ProjectCommon,
	Class:   ClassChargerState,
	Cmd:     CmdLastChargeRateChanged,
}

// title : Charging information,
// desc : Charging information.,
// support : 0905;0906;0907;0909;090a,
// triggered : when the product is charging or when the charging state changes.,
const CmdChargingInfo CmdDef = 3

type CommonChargerStateChargingInfo Command

type CommonChargerStateChargingInfoArguments struct {
	Phase            uint32
	Rate             uint32
	Intensity        uint8
	FullChargingTime uint8
}

func (a CommonChargerStateChargingInfo) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonChargerStateChargingInfoArguments{}
	var offset = 0
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Phase)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Rate)
	offset += 4
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.Intensity)
	offset++
	ConvLittleEndianSliceToNumeric(b[offset:offset+1], &arg.FullChargingTime)
	offset++

	return arg
}
func (a CommonChargerStateChargingInfo) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var ChargerStateChargingInfo = CommonChargerStateChargingInfo{
	Project: ProjectCommon,
	Class:   ClassChargerState,
	Cmd:     CmdChargingInfo,
}

// Commands sent by the drone to inform about the run or flight state
const ClassRunState ClassDef = 30

// title : Current run id,
// desc : Current run id.\n A run id is uniquely identifying a run or a flight.\n For each run is generated on the drone a file which can be used by Academy to sum up the run.\n Also, each medias taken during a run has a filename containing the run id.,
// support : 0901:3.0.1;090c;090e,
// triggered : when the drone generates a new run id (generally right after a take off).,
const CmdRunIdChanged CmdDef = 0

type CommonRunStateRunIdChanged Command

type CommonRunStateRunIdChangedArguments struct {
	RunId string
}

func (a CommonRunStateRunIdChanged) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonRunStateRunIdChangedArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.RunId = string(b[offset : offset+stringEnd])
	offset += stringEnd

	return arg
}
func (a CommonRunStateRunIdChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var RunStateRunIdChanged = CommonRunStateRunIdChanged{
	Project: ProjectCommon,
	Class:   ClassRunState,
	Cmd:     CmdRunIdChanged,
}

// Factory reset commands
const ClassFactory ClassDef = 31

// title : Reset the product to its factory settings,
// desc : This command will request a factory reset from the prodcut. *The factory reset procedure implies an automatic reboot*, which will be done immediately after receiving this command.,
// result : The product will reboot, all settings will be reset to their default values. All data on the product will also be erased.,
const CmdResetDUPLICATE CmdDef = 0

type CommonFactoryReset Command

type CommonFactoryResetArguments struct {
}

func (a CommonFactoryReset) Decode(b []byte) interface{} {
	//TODO: .............
	arg := CommonFactoryResetArguments{}
	// No arguments to decode here !!

	return arg
}
func (a CommonFactoryReset) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var FactoryReset = CommonFactoryReset{
	Project: ProjectCommon,
	Class:   ClassFactory,
	Cmd:     CmdReset,
}

// Update related commands
const ClassUpdateState ClassDef = 34

// title : Software update status,
// desc : Status of the latest software update,
// support : 0914,
// triggered : at connection during the first boot after a firmware update.,
const CmdUpdateStateChanged CmdDef = 0

type CommonUpdateStateUpdateStateChanged Command

type CommonUpdateStateUpdateStateChangedArguments struct {
	SourceVersion string
	TargetVersion string
	Status        uint32
}

func (a CommonUpdateStateUpdateStateChanged) Decode(b []byte) interface{} {
	//TODO: .............
	var stringEnd int
	var err error
	arg := CommonUpdateStateUpdateStateChangedArguments{}
	var offset = 0

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.SourceVersion = string(b[offset : offset+stringEnd])
	offset += stringEnd

	stringEnd, err = getLengthOfStringData(b[offset:])
	if err != nil {
		log.Println("error: ", err)
	}
	arg.TargetVersion = string(b[offset : offset+stringEnd])
	offset += stringEnd
	ConvLittleEndianSliceToNumeric(b[offset:offset+4], &arg.Status)
	offset += 4

	return arg
}
func (a CommonUpdateStateUpdateStateChanged) Encode(commandStruct interface{}) []byte {
	//TODO: .............

	var bs []byte
	valueOf := reflect.ValueOf(commandStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := ConvLittleEndianNumericToSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

var UpdateStateUpdateStateChanged = CommonUpdateStateUpdateStateChanged{
	Project: ProjectCommon,
	Class:   ClassUpdateState,
	Cmd:     CmdUpdateStateChanged,
}

type Decoder interface {
	Decode([]byte) interface{}
}

var CommandMap = map[Command]Decoder{
	Command(PilotingTakeOff):                                                PilotingTakeOff,
	Command(PilotingPCMD):                                                   PilotingPCMD,
	Command(PilotingLanding):                                                PilotingLanding,
	Command(PilotingEmergency):                                              PilotingEmergency,
	Command(PilotingNavigateHome):                                           PilotingNavigateHome,
	Command(PilotingAutoTakeOffMode):                                        PilotingAutoTakeOffMode,
	Command(PilotingmoveBy):                                                 PilotingmoveBy,
	Command(PilotingUserTakeOff):                                            PilotingUserTakeOff,
	Command(PilotingCircle):                                                 PilotingCircle,
	Command(PilotingmoveTo):                                                 PilotingmoveTo,
	Command(PilotingCancelMoveTo):                                           PilotingCancelMoveTo,
	Command(PilotingStartPilotedPOI):                                        PilotingStartPilotedPOI,
	Command(PilotingStopPilotedPOI):                                         PilotingStopPilotedPOI,
	Command(PilotingCancelMoveBy):                                           PilotingCancelMoveBy,
	Command(AnimationsFlip):                                                 AnimationsFlip,
	Command(CameraOrientation):                                              CameraOrientation,
	Command(CameraOrientationV2):                                            CameraOrientationV2,
	Command(CameraVelocity):                                                 CameraVelocity,
	Command(MediaRecordPicture):                                             MediaRecordPicture,
	Command(MediaRecordVideo):                                               MediaRecordVideo,
	Command(MediaRecordPictureV2):                                           MediaRecordPictureV2,
	Command(MediaRecordVideoV2):                                             MediaRecordVideoV2,
	Command(MediaRecordStatePictureStateChanged):                            MediaRecordStatePictureStateChanged,
	Command(MediaRecordStateVideoStateChanged):                              MediaRecordStateVideoStateChanged,
	Command(MediaRecordStatePictureStateChangedV2):                          MediaRecordStatePictureStateChangedV2,
	Command(MediaRecordStateVideoStateChangedV2):                            MediaRecordStateVideoStateChangedV2,
	Command(MediaRecordStateVideoResolutionState):                           MediaRecordStateVideoResolutionState,
	Command(MediaRecordEventPictureEventChanged):                            MediaRecordEventPictureEventChanged,
	Command(MediaRecordEventVideoEventChanged):                              MediaRecordEventVideoEventChanged,
	Command(PilotingStateFlyingStateChanged):                                PilotingStateFlyingStateChanged,
	Command(PilotingStateAlertStateChanged):                                 PilotingStateAlertStateChanged,
	Command(PilotingStateNavigateHomeStateChanged):                          PilotingStateNavigateHomeStateChanged,
	Command(PilotingStatePositionChanged):                                   PilotingStatePositionChanged,
	Command(PilotingStateSpeedChanged):                                      PilotingStateSpeedChanged,
	Command(PilotingStateAttitudeChanged):                                   PilotingStateAttitudeChanged,
	Command(PilotingStateAutoTakeOffModeChanged):                            PilotingStateAutoTakeOffModeChanged,
	Command(PilotingStateAltitudeChanged):                                   PilotingStateAltitudeChanged,
	Command(PilotingStateGpsLocationChanged):                                PilotingStateGpsLocationChanged,
	Command(PilotingStateLandingStateChanged):                               PilotingStateLandingStateChanged,
	Command(PilotingStateAirSpeedChanged):                                   PilotingStateAirSpeedChanged,
	Command(PilotingStatemoveToChanged):                                     PilotingStatemoveToChanged,
	Command(PilotingStateMotionState):                                       PilotingStateMotionState,
	Command(PilotingStatePilotedPOI):                                        PilotingStatePilotedPOI,
	Command(PilotingStateReturnHomeBatteryCapacity):                         PilotingStateReturnHomeBatteryCapacity,
	Command(PilotingStatemoveByChanged):                                     PilotingStatemoveByChanged,
	Command(PilotingStateHoveringWarning):                                   PilotingStateHoveringWarning,
	Command(PilotingStateForcedLandingAutoTrigger):                          PilotingStateForcedLandingAutoTrigger,
	Command(PilotingStateWindStateChanged):                                  PilotingStateWindStateChanged,
	Command(PilotingEventmoveByEnd):                                         PilotingEventmoveByEnd,
	Command(NetworkWifiScan):                                                NetworkWifiScan,
	Command(NetworkWifiAuthChannel):                                         NetworkWifiAuthChannel,
	Command(NetworkStateWifiScanListChanged):                                NetworkStateWifiScanListChanged,
	Command(NetworkStateAllWifiScanChanged):                                 NetworkStateAllWifiScanChanged,
	Command(NetworkStateWifiAuthChannelListChanged):                         NetworkStateWifiAuthChannelListChanged,
	Command(NetworkStateAllWifiAuthChannelChanged):                          NetworkStateAllWifiAuthChannelChanged,
	Command(PilotingSettingsMaxAltitude):                                    PilotingSettingsMaxAltitude,
	Command(PilotingSettingsMaxTilt):                                        PilotingSettingsMaxTilt,
	Command(PilotingSettingsAbsolutControl):                                 PilotingSettingsAbsolutControl,
	Command(PilotingSettingsMaxDistance):                                    PilotingSettingsMaxDistance,
	Command(PilotingSettingsNoFlyOverMaxDistance):                           PilotingSettingsNoFlyOverMaxDistance,
	Command(PilotingSettingssetAutonomousFlightMaxHorizontalSpeed):          PilotingSettingssetAutonomousFlightMaxHorizontalSpeed,
	Command(PilotingSettingssetAutonomousFlightMaxVerticalSpeed):            PilotingSettingssetAutonomousFlightMaxVerticalSpeed,
	Command(PilotingSettingssetAutonomousFlightMaxHorizontalAcceleration):   PilotingSettingssetAutonomousFlightMaxHorizontalAcceleration,
	Command(PilotingSettingssetAutonomousFlightMaxVerticalAcceleration):     PilotingSettingssetAutonomousFlightMaxVerticalAcceleration,
	Command(PilotingSettingssetAutonomousFlightMaxRotationSpeed):            PilotingSettingssetAutonomousFlightMaxRotationSpeed,
	Command(PilotingSettingsBankedTurn):                                     PilotingSettingsBankedTurn,
	Command(PilotingSettingsMinAltitude):                                    PilotingSettingsMinAltitude,
	Command(PilotingSettingsCirclingDirection):                              PilotingSettingsCirclingDirection,
	Command(PilotingSettingsCirclingRadius):                                 PilotingSettingsCirclingRadius,
	Command(PilotingSettingsCirclingAltitude):                               PilotingSettingsCirclingAltitude,
	Command(PilotingSettingsPitchMode):                                      PilotingSettingsPitchMode,
	Command(PilotingSettingsSetMotionDetectionMode):                         PilotingSettingsSetMotionDetectionMode,
	Command(PilotingSettingsStateMaxAltitudeChanged):                        PilotingSettingsStateMaxAltitudeChanged,
	Command(PilotingSettingsStateMaxTiltChanged):                            PilotingSettingsStateMaxTiltChanged,
	Command(PilotingSettingsStateAbsolutControlChanged):                     PilotingSettingsStateAbsolutControlChanged,
	Command(PilotingSettingsStateMaxDistanceChanged):                        PilotingSettingsStateMaxDistanceChanged,
	Command(PilotingSettingsStateNoFlyOverMaxDistanceChanged):               PilotingSettingsStateNoFlyOverMaxDistanceChanged,
	Command(PilotingSettingsStateAutonomousFlightMaxHorizontalSpeed):        PilotingSettingsStateAutonomousFlightMaxHorizontalSpeed,
	Command(PilotingSettingsStateAutonomousFlightMaxVerticalSpeed):          PilotingSettingsStateAutonomousFlightMaxVerticalSpeed,
	Command(PilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration): PilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration,
	Command(PilotingSettingsStateAutonomousFlightMaxVerticalAcceleration):   PilotingSettingsStateAutonomousFlightMaxVerticalAcceleration,
	Command(PilotingSettingsStateAutonomousFlightMaxRotationSpeed):          PilotingSettingsStateAutonomousFlightMaxRotationSpeed,
	Command(PilotingSettingsStateBankedTurnChanged):                         PilotingSettingsStateBankedTurnChanged,
	Command(PilotingSettingsStateMinAltitudeChanged):                        PilotingSettingsStateMinAltitudeChanged,
	Command(PilotingSettingsStateCirclingDirectionChanged):                  PilotingSettingsStateCirclingDirectionChanged,
	Command(PilotingSettingsStateCirclingRadiusChanged):                     PilotingSettingsStateCirclingRadiusChanged,
	Command(PilotingSettingsStateCirclingAltitudeChanged):                   PilotingSettingsStateCirclingAltitudeChanged,
	Command(PilotingSettingsStatePitchModeChanged):                          PilotingSettingsStatePitchModeChanged,
	Command(PilotingSettingsStateMotionDetection):                           PilotingSettingsStateMotionDetection,
	Command(SpeedSettingsMaxVerticalSpeed):                                  SpeedSettingsMaxVerticalSpeed,
	Command(SpeedSettingsMaxRotationSpeed):                                  SpeedSettingsMaxRotationSpeed,
	Command(SpeedSettingsHullProtection):                                    SpeedSettingsHullProtection,
	Command(SpeedSettingsOutdoor):                                           SpeedSettingsOutdoor,
	Command(SpeedSettingsMaxPitchRollRotationSpeed):                         SpeedSettingsMaxPitchRollRotationSpeed,
	Command(SpeedSettingsStateMaxVerticalSpeedChanged):                      SpeedSettingsStateMaxVerticalSpeedChanged,
	Command(SpeedSettingsStateMaxRotationSpeedChanged):                      SpeedSettingsStateMaxRotationSpeedChanged,
	Command(SpeedSettingsStateHullProtectionChanged):                        SpeedSettingsStateHullProtectionChanged,
	Command(SpeedSettingsStateOutdoorChanged):                               SpeedSettingsStateOutdoorChanged,
	Command(SpeedSettingsStateMaxPitchRollRotationSpeedChanged):             SpeedSettingsStateMaxPitchRollRotationSpeedChanged,
	Command(NetworkSettingsWifiSelection):                                   NetworkSettingsWifiSelection,
	Command(NetworkSettingswifiSecurity):                                    NetworkSettingswifiSecurity,
	Command(NetworkSettingsStateWifiSelectionChanged):                       NetworkSettingsStateWifiSelectionChanged,
	Command(NetworkSettingsStatewifiSecurityChanged):                        NetworkSettingsStatewifiSecurityChanged,
	Command(NetworkSettingsStatewifiSecurity):                               NetworkSettingsStatewifiSecurity,
	Command(SettingsStateProductMotorVersionListChanged):                    SettingsStateProductMotorVersionListChanged,
	Command(SettingsStateProductGPSVersionChanged):                          SettingsStateProductGPSVersionChanged,
	Command(SettingsStateMotorErrorStateChanged):                            SettingsStateMotorErrorStateChanged,
	Command(SettingsStateMotorSoftwareVersionChanged):                       SettingsStateMotorSoftwareVersionChanged,
	Command(SettingsStateMotorFlightsStatusChanged):                         SettingsStateMotorFlightsStatusChanged,
	Command(SettingsStateMotorErrorLastErrorChanged):                        SettingsStateMotorErrorLastErrorChanged,
	Command(SettingsStateP7ID):                                              SettingsStateP7ID,
	Command(SettingsStateCPUID):                                             SettingsStateCPUID,
	Command(PictureSettingsPictureFormatSelection):                          PictureSettingsPictureFormatSelection,
	Command(PictureSettingsAutoWhiteBalanceSelection):                       PictureSettingsAutoWhiteBalanceSelection,
	Command(PictureSettingsExpositionSelection):                             PictureSettingsExpositionSelection,
	Command(PictureSettingsSaturationSelection):                             PictureSettingsSaturationSelection,
	Command(PictureSettingsTimelapseSelection):                              PictureSettingsTimelapseSelection,
	Command(PictureSettingsVideoAutorecordSelection):                        PictureSettingsVideoAutorecordSelection,
	Command(PictureSettingsVideoStabilizationMode):                          PictureSettingsVideoStabilizationMode,
	Command(PictureSettingsVideoRecordingMode):                              PictureSettingsVideoRecordingMode,
	Command(PictureSettingsVideoFramerate):                                  PictureSettingsVideoFramerate,
	Command(PictureSettingsVideoResolutions):                                PictureSettingsVideoResolutions,
	Command(PictureSettingsStatePictureFormatChanged):                       PictureSettingsStatePictureFormatChanged,
	Command(PictureSettingsStateAutoWhiteBalanceChanged):                    PictureSettingsStateAutoWhiteBalanceChanged,
	Command(PictureSettingsStateExpositionChanged):                          PictureSettingsStateExpositionChanged,
	Command(PictureSettingsStateSaturationChanged):                          PictureSettingsStateSaturationChanged,
	Command(PictureSettingsStateTimelapseChanged):                           PictureSettingsStateTimelapseChanged,
	Command(PictureSettingsStateVideoAutorecordChanged):                     PictureSettingsStateVideoAutorecordChanged,
	Command(PictureSettingsStateVideoStabilizationModeChanged):              PictureSettingsStateVideoStabilizationModeChanged,
	Command(PictureSettingsStateVideoRecordingModeChanged):                  PictureSettingsStateVideoRecordingModeChanged,
	Command(PictureSettingsStateVideoFramerateChanged):                      PictureSettingsStateVideoFramerateChanged,
	Command(PictureSettingsStateVideoResolutionsChanged):                    PictureSettingsStateVideoResolutionsChanged,
	Command(MediaStreamingVideoEnable):                                      MediaStreamingVideoEnable,
	Command(MediaStreamingVideoStreamMode):                                  MediaStreamingVideoStreamMode,
	Command(MediaStreamingStateVideoEnableChanged):                          MediaStreamingStateVideoEnableChanged,
	Command(MediaStreamingStateVideoStreamModeChanged):                      MediaStreamingStateVideoStreamModeChanged,
	Command(GPSSettingsSetHome):                                             GPSSettingsSetHome,
	Command(GPSSettingsResetHome):                                           GPSSettingsResetHome,
	Command(GPSSettingsSendControllerGPS):                                   GPSSettingsSendControllerGPS,
	Command(GPSSettingsHomeType):                                            GPSSettingsHomeType,
	Command(GPSSettingsReturnHomeDelay):                                     GPSSettingsReturnHomeDelay,
	Command(GPSSettingsReturnHomeMinAltitude):                               GPSSettingsReturnHomeMinAltitude,
	Command(GPSSettingsStateHomeChanged):                                    GPSSettingsStateHomeChanged,
	Command(GPSSettingsStateResetHomeChanged):                               GPSSettingsStateResetHomeChanged,
	Command(GPSSettingsStateGPSFixStateChanged):                             GPSSettingsStateGPSFixStateChanged,
	Command(GPSSettingsStateGPSUpdateStateChanged):                          GPSSettingsStateGPSUpdateStateChanged,
	Command(GPSSettingsStateHomeTypeChanged):                                GPSSettingsStateHomeTypeChanged,
	Command(GPSSettingsStateReturnHomeDelayChanged):                         GPSSettingsStateReturnHomeDelayChanged,
	Command(GPSSettingsStateGeofenceCenterChanged):                          GPSSettingsStateGeofenceCenterChanged,
	Command(GPSSettingsStateReturnHomeMinAltitudeChanged):                   GPSSettingsStateReturnHomeMinAltitudeChanged,
	Command(CameraStateOrientation):                                         CameraStateOrientation,
	Command(CameraStatedefaultCameraOrientation):                            CameraStatedefaultCameraOrientation,
	Command(CameraStateOrientationV2):                                       CameraStateOrientationV2,
	Command(CameraStatedefaultCameraOrientationV2):                          CameraStatedefaultCameraOrientationV2,
	Command(CameraStateVelocityRange):                                       CameraStateVelocityRange,
	Command(AntiflickeringelectricFrequency):                                AntiflickeringelectricFrequency,
	Command(AntiflickeringsetMode):                                          AntiflickeringsetMode,
	Command(AntiflickeringStateelectricFrequencyChanged):                    AntiflickeringStateelectricFrequencyChanged,
	Command(AntiflickeringStatemodeChanged):                                 AntiflickeringStatemodeChanged,
	Command(GPSStateNumberOfSatelliteChanged):                               GPSStateNumberOfSatelliteChanged,
	Command(GPSStateHomeTypeAvailabilityChanged):                            GPSStateHomeTypeAvailabilityChanged,
	Command(GPSStateHomeTypeChosenChanged):                                  GPSStateHomeTypeChosenChanged,
	Command(PROStateFeatures):                                               PROStateFeatures,
	Command(AccessoryStateConnectedAccessories):                             AccessoryStateConnectedAccessories,
	Command(AccessoryStateBattery):                                          AccessoryStateBattery,
	Command(SoundStartAlertSound):                                           SoundStartAlertSound,
	Command(SoundStopAlertSound):                                            SoundStopAlertSound,
	Command(SoundStateAlertSound):                                           SoundStateAlertSound,
	Command(NetworkDisconnect):                                              NetworkDisconnect,
	Command(NetworkEventDisconnection):                                      NetworkEventDisconnection,
	Command(SettingsAllSettings):                                            SettingsAllSettings,
	Command(SettingsReset):                                                  SettingsReset,
	Command(SettingsProductName):                                            SettingsProductName,
	Command(SettingsCountry):                                                SettingsCountry,
	Command(SettingsAutoCountry):                                            SettingsAutoCountry,
	Command(SettingsStateAllSettingsChanged):                                SettingsStateAllSettingsChanged,
	Command(SettingsStateResetChanged):                                      SettingsStateResetChanged,
	Command(SettingsStateProductNameChanged):                                SettingsStateProductNameChanged,
	Command(SettingsStateProductVersionChanged):                             SettingsStateProductVersionChanged,
	Command(SettingsStateProductSerialHighChanged):                          SettingsStateProductSerialHighChanged,
	Command(SettingsStateProductSerialLowChanged):                           SettingsStateProductSerialLowChanged,
	Command(SettingsStateCountryChanged):                                    SettingsStateCountryChanged,
	Command(SettingsStateAutoCountryChanged):                                SettingsStateAutoCountryChanged,
	Command(SettingsStateBoardIdChanged):                                    SettingsStateBoardIdChanged,
	Command(CommonAllStates):                                                CommonAllStates,
	Command(CommonCurrentDate):                                              CommonCurrentDate,
	Command(CommonCurrentTime):                                              CommonCurrentTime,
	Command(CommonReboot):                                                   CommonReboot,
	Command(CommonCurrentDateTime):                                          CommonCurrentDateTime,
	Command(CommonStateAllStatesChanged):                                    CommonStateAllStatesChanged,
	Command(CommonStateBatteryStateChanged):                                 CommonStateBatteryStateChanged,
	Command(CommonStateMassStorageStateListChanged):                         CommonStateMassStorageStateListChanged,
	Command(CommonStateMassStorageInfoStateListChanged):                     CommonStateMassStorageInfoStateListChanged,
	Command(CommonStateCurrentDateChanged):                                  CommonStateCurrentDateChanged,
	Command(CommonStateCurrentTimeChanged):                                  CommonStateCurrentTimeChanged,
	Command(CommonStateMassStorageInfoRemainingListChanged):                 CommonStateMassStorageInfoRemainingListChanged,
	Command(CommonStateWifiSignalChanged):                                   CommonStateWifiSignalChanged,
	Command(CommonStateSensorsStatesListChanged):                            CommonStateSensorsStatesListChanged,
	Command(CommonStateProductModel):                                        CommonStateProductModel,
	Command(CommonStateCountryListKnown):                                    CommonStateCountryListKnown,
	Command(CommonStateDeprecatedMassStorageContentChanged):                 CommonStateDeprecatedMassStorageContentChanged,
	Command(CommonStateMassStorageContent):                                  CommonStateMassStorageContent,
	Command(CommonStateMassStorageContentForCurrentRun):                     CommonStateMassStorageContentForCurrentRun,
	Command(CommonStateVideoRecordingTimestamp):                             CommonStateVideoRecordingTimestamp,
	Command(CommonStateCurrentDateTimeChanged):                              CommonStateCurrentDateTimeChanged,
	Command(CommonStateLinkSignalQuality):                                   CommonStateLinkSignalQuality,
	Command(CommonStateBootId):                                              CommonStateBootId,
	Command(OverHeatSwitchOff):                                              OverHeatSwitchOff,
	Command(OverHeatVentilate):                                              OverHeatVentilate,
	Command(OverHeatStateOverHeatChanged):                                   OverHeatStateOverHeatChanged,
	Command(OverHeatStateOverHeatRegulationChanged):                         OverHeatStateOverHeatRegulationChanged,
	Command(ControllerisPiloting):                                           ControllerisPiloting,
	Command(ControllerPeerStateChanged):                                     ControllerPeerStateChanged,
	Command(WifiSettingsOutdoorSetting):                                     WifiSettingsOutdoorSetting,
	Command(WifiSettingsStateoutdoorSettingsChanged):                        WifiSettingsStateoutdoorSettingsChanged,
	Command(MavlinkStart):                                                   MavlinkStart,
	Command(MavlinkPause):                                                   MavlinkPause,
	Command(MavlinkStop):                                                    MavlinkStop,
	Command(MavlinkStateMavlinkFilePlayingStateChanged):                     MavlinkStateMavlinkFilePlayingStateChanged,
	Command(MavlinkStateMavlinkPlayErrorStateChanged):                       MavlinkStateMavlinkPlayErrorStateChanged,
	Command(MavlinkStateMissionItemExecuted):                                MavlinkStateMissionItemExecuted,
	Command(FlightPlanSettingsReturnHomeOnDisconnect):                       FlightPlanSettingsReturnHomeOnDisconnect,
	Command(FlightPlanSettingsStateReturnHomeOnDisconnectChanged):           FlightPlanSettingsStateReturnHomeOnDisconnectChanged,
	Command(CalibrationMagnetoCalibration):                                  CalibrationMagnetoCalibration,
	Command(CalibrationPitotCalibration):                                    CalibrationPitotCalibration,
	Command(CalibrationStateMagnetoCalibrationStateChanged):                 CalibrationStateMagnetoCalibrationStateChanged,
	Command(CalibrationStateMagnetoCalibrationRequiredState):                CalibrationStateMagnetoCalibrationRequiredState,
	Command(CalibrationStateMagnetoCalibrationAxisToCalibrateChanged):       CalibrationStateMagnetoCalibrationAxisToCalibrateChanged,
	Command(CalibrationStateMagnetoCalibrationStartedChanged):               CalibrationStateMagnetoCalibrationStartedChanged,
	Command(CalibrationStatePitotCalibrationStateChanged):                   CalibrationStatePitotCalibrationStateChanged,
	Command(CameraSettingsStateCameraSettingsChanged):                       CameraSettingsStateCameraSettingsChanged,
	Command(GPSControllerPositionForRun):                                    GPSControllerPositionForRun,
	Command(FlightPlanStateAvailabilityStateChanged):                        FlightPlanStateAvailabilityStateChanged,
	Command(FlightPlanStateComponentStateListChanged):                       FlightPlanStateComponentStateListChanged,
	Command(FlightPlanStateLockStateChanged):                                FlightPlanStateLockStateChanged,
	Command(FlightPlanEventStartingErrorEvent):                              FlightPlanEventStartingErrorEvent,
	Command(FlightPlanEventSpeedBridleEvent):                                FlightPlanEventSpeedBridleEvent,
	Command(ARLibsVersionsStateControllerLibARCommandsVersion):              ARLibsVersionsStateControllerLibARCommandsVersion,
	Command(ARLibsVersionsStateSkyControllerLibARCommandsVersion):           ARLibsVersionsStateSkyControllerLibARCommandsVersion,
	Command(ARLibsVersionsStateDeviceLibARCommandsVersion):                  ARLibsVersionsStateDeviceLibARCommandsVersion,
	Command(AudioControllerReadyForStreaming):                               AudioControllerReadyForStreaming,
	Command(AudioStateAudioStreamingRunning):                                AudioStateAudioStreamingRunning,
	Command(Headlightsintensity):                                            Headlightsintensity,
	Command(HeadlightsStateintensityChanged):                                HeadlightsStateintensityChanged,
	Command(AnimationsStartAnimation):                                       AnimationsStartAnimation,
	Command(AnimationsStopAnimation):                                        AnimationsStopAnimation,
	Command(AnimationsStopAllAnimations):                                    AnimationsStopAllAnimations,
	Command(AnimationsStateList):                                            AnimationsStateList,
	Command(AccessoryConfig):                                                AccessoryConfig,
	Command(AccessoryStateSupportedAccessoriesListChanged):                  AccessoryStateSupportedAccessoriesListChanged,
	Command(AccessoryStateAccessoryConfigChanged):                           AccessoryStateAccessoryConfigChanged,
	Command(AccessoryStateAccessoryConfigModificationEnabled):               AccessoryStateAccessoryConfigModificationEnabled,
	Command(ChargerSetMaxChargeRate):                                        ChargerSetMaxChargeRate,
	Command(ChargerStateMaxChargeRateChanged):                               ChargerStateMaxChargeRateChanged,
	Command(ChargerStateCurrentChargeStateChanged):                          ChargerStateCurrentChargeStateChanged,
	Command(ChargerStateLastChargeRateChanged):                              ChargerStateLastChargeRateChanged,
	Command(ChargerStateChargingInfo):                                       ChargerStateChargingInfo,
	Command(RunStateRunIdChanged):                                           RunStateRunIdChanged,
	Command(FactoryReset):                                                   FactoryReset,
	Command(UpdateStateUpdateStateChanged):                                  UpdateStateUpdateStateChanged,
}

// lenStringData takes a []byte which is the data for the arguments, and returns
// the position of the 0 terminator for the string.
// The []byte given as input will start looking from the beginning of the slice,
// so the input slice should be sliced to start from the offset of the string.
func lenStringData(b []byte) (int, error) {
	// Figure out the length of the string
	for i := 0; i < cap(b); i++ {
		//fmt.Printf("%+v, of type %T\n", b[i], b[i])

		//fmt.Println("i = ", i)
		if b[i] == 0 {
			//fmt.Println("lengthString = ", i)

			// add 1 to jump to the 0
			return i + 1, nil
		}

	}

	err := fmt.Errorf("no string bytes found, returning 0")
	return 0, err
}

func getLengthOfStringData(b []byte) (int, error) {
	// Figure out the length of the string
	for i := 0; i < cap(b); i++ {
		//fmt.Printf("%+v, of type %T\n", b[i], b[i])

		//fmt.Println("i = ", i)
		if b[i] == 0 {
			//fmt.Println("lengthString = ", i)

			// add 1 to jump to the 0
			return i + 1, nil
		}

	}

	err := fmt.Errorf("no string bytes found, returning 0")
	return 0, err
}

// ConvLittleEndianSliceToNumeric takes a []byte, and an *out variable of type
// uint8/int8/uint16/int16/uint32/int32/uint64/int64/float32/float64
// and convert the []byte, and places the result into the *out variable.
func ConvLittleEndianSliceToNumeric(in []byte, out interface{}) {
	switch out := out.(type) {
	case *uint8:
		*out = uint8(in[0])
	case *int8:
		*out = int8(in[0])
	case *uint16:
		*out = binary.LittleEndian.Uint16(in)
	case *int16:
		*out = int16(binary.LittleEndian.Uint16(in))
	case *uint32:
		*out = binary.LittleEndian.Uint32(in)
	case *int32:
		*out = int32(binary.LittleEndian.Uint32(in))
	case *uint64:
		*out = binary.LittleEndian.Uint64(in)
	case *int64:
		*out = int64(binary.LittleEndian.Uint32(in))
	case *float32:
		bits := binary.LittleEndian.Uint32(in)
		*out = math.Float32frombits(bits)
	case *float64:
		bits := binary.LittleEndian.Uint64(in)
		*out = math.Float64frombits(bits)
	case *string:
		*out = string(in)
	}
}

// ConvLittleEndianNumericToSlice takes a a value of any of the standard types
// uint8/int8/uint16/int16/uint32/int32/uint64/int64/float32/float64
// and convert to a []byte.
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

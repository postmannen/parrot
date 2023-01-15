package parrot

import (
	"log"
	"os"

	"github.com/eiannone/keyboard"
)

type inputAction int

const (
	// Standard actions.
	//
	ActionPcmdFlag                       inputAction = iota
	ActionPcmdRollLeft                   inputAction = iota
	ActionPcmdRollRight                  inputAction = iota
	ActionPcmdPitchForward               inputAction = iota
	ActionPcmdPitchBackward              inputAction = iota
	ActionPcmdYawClockwise               inputAction = iota
	ActionPcmdYawCounterClockwise        inputAction = iota
	ActionPcmdHover                      inputAction = iota
	ActionPcmdGazInc                     inputAction = iota
	ActionPcmdGazDec                     inputAction = iota
	ActionPcmdRepeatLastCmd              inputAction = iota
	ActionTakeoff                        inputAction = iota
	ActionLanding                        inputAction = iota
	ActionEmergency                      inputAction = iota
	ActionNavigateHomeStart              inputAction = iota // Check how to implement it in xml line 153
	ActionNavigateHomeStop               inputAction = iota // Check how to implement it in xml line 153
	ActionMoveBy                         inputAction = iota // Check how to implement it in xml line 181
	ActionUserTakeoff                    inputAction = iota
	ActionMoveTo                         inputAction = iota // Check how to implement it in xml line 259
	ActionCancelMoveTo                   inputAction = iota
	ActionStartPilotedPOI                inputAction = iota
	ActionStopPilotedPOI                 inputAction = iota
	ActionCancelMoveBy                   inputAction = iota
	ActionMoveToSetLatInc                inputAction = iota // Direction North
	ActionMoveToSetLatDec                inputAction = iota // Direction South
	ActionMoveToSetLonInc                inputAction = iota // Direction East
	ActionMoveToSetLonDec                inputAction = iota // Direction West
	ActionMoveToExecute                  inputAction = iota // Execute moveTo next waypoint
	ActionMoveToCancel                   inputAction = iota // Cancel all moveTo operation
	ActionMoveToSetBufferCurrentPosition inputAction = iota // Set buffer to current position

	// Custom actions.
	//
	ActionHow inputAction = iota
	// Flattrim should be performed before a takeoff
	// to calibrate the drone.
	ActionFlatTrim inputAction = iota
	// TODO: Also check out the <class name="PilotingSettings" id="2">"
	// starting at line 1400 in the ardrone3.xml document, for more
	// commands to eventually implement.
)

// readKeyBoardEvent will read keys pressed on the keyboard,
// and pass on the correct action to be executed.
//
// TODO: Make more source to create inputActions than keyboard...
// Geofencing ?
// Map route ?
func (d *Drone) readKeyBoardEvent(packetCreator *udpPacketCreator) {

	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := keyboard.Close()
		if err != nil {
			log.Printf("error: failed to close keyboard: %v\n", err)
		}
	}()

	for {
		select {
		case event := <-keysEvents:

			if event.Err != nil {
				panic(event.Err)
			}

			switch {
			case event.Key == keyboard.KeyCtrlC:
				log.Printf("info: ctrl+c pressed, forcing shutdown of all threads\n")
				os.Exit(1)
			case event.Key == keyboard.KeyEsc:
				d.quitCh <- struct{}{}
			case event.Rune == 'q':
				// Initiate a reconnect of the network.
				select {
				case d.networkReconnectCh <- struct{}{}:
				default:
				}
			case event.Rune == '2':
				d.handleInputAction(packetCreator, ActionTakeoff)
			case event.Rune == '1':
				d.handleInputAction(packetCreator, ActionLanding)
			case event.Rune == 'h':
				d.handleInputAction(packetCreator, ActionNavigateHomeStart)
			case event.Rune == 'H':
				d.handleInputAction(packetCreator, ActionNavigateHomeStop)

			case event.Rune == 'w':
				d.handleInputAction(packetCreator, ActionPcmdGazInc)
			case event.Rune == 's':
				d.handleInputAction(packetCreator, ActionPcmdGazDec)
			case event.Rune == 'a':
				d.handleInputAction(packetCreator, ActionPcmdYawCounterClockwise)
			case event.Rune == 'd':
				d.handleInputAction(packetCreator, ActionPcmdYawClockwise)

			case event.Key == keyboard.KeyArrowUp:
				d.handleInputAction(packetCreator, ActionPcmdPitchForward)
			case event.Key == keyboard.KeyArrowDown:
				d.handleInputAction(packetCreator, ActionPcmdPitchBackward)
			case event.Key == keyboard.KeyArrowLeft:
				d.handleInputAction(packetCreator, ActionPcmdRollLeft)
			case event.Key == keyboard.KeyArrowRight:
				d.handleInputAction(packetCreator, ActionPcmdRollRight)
			case event.Key == keyboard.KeySpace:
				d.handleInputAction(packetCreator, ActionPcmdRepeatLastCmd)

			case event.Key == keyboard.KeyCtrlW:
				d.handleInputAction(packetCreator, ActionMoveToSetLatInc)
			case event.Key == keyboard.KeyCtrlS:
				d.handleInputAction(packetCreator, ActionMoveToSetLatDec)
			case event.Key == keyboard.KeyCtrlA:
				d.handleInputAction(packetCreator, ActionMoveToSetLonDec)
			case event.Key == keyboard.KeyCtrlD:
				d.handleInputAction(packetCreator, ActionMoveToSetLonInc)
			case event.Key == keyboard.KeyCtrlX:
				d.handleInputAction(packetCreator, ActionMoveToSetBufferCurrentPosition)
			case event.Key == keyboard.KeyCtrlSpace:
				d.handleInputAction(packetCreator, ActionMoveToExecute)
			case event.Key == keyboard.KeyCtrlQ:
				d.handleInputAction(packetCreator, ActionMoveToCancel)

			case event.Rune == 'z':
				d.handleInputAction(packetCreator, ActionPcmdHover)

			}
		}

	}

}

// handleInputAction is where we specify what package to send to the drone
// based on what action came out of the readKeyboardEvent method.
//
// The reason we have this function and don't encode the packets directly
// in readKeyBoardEvent, is that we might want to have other input methods
// then the keyboard to control the drone.
// This function will execute the commands that arrives on the d.inputActionsCh.
func (d *Drone) handleInputAction(packetCreator *udpPacketCreator, action inputAction) {

	// --------------Standard actions
	switch action {
	case ActionTakeoff:
		p := packetCreator.encodeCmd(Command(PilotingTakeOff), &Ardrone3PilotingTakeOffArguments{})
		d.pcmdPacketSchedulerCh <- p
	case ActionLanding:
		p := packetCreator.encodeCmd(Command(PilotingLanding), &Ardrone3PilotingLandingArguments{})
		d.pcmdPacketSchedulerCh <- p
	case ActionNavigateHomeStart:
		p := packetCreator.encodeCmd(Command(PilotingNavigateHome), &Ardrone3PilotingNavigateHomeArguments{Start: 1})
		d.pcmdPacketSchedulerCh <- p
	case ActionNavigateHomeStop:
		p := packetCreator.encodeCmd(Command(PilotingNavigateHome), &Ardrone3PilotingNavigateHomeArguments{Start: 0})
		d.pcmdPacketSchedulerCh <- p

	// --------------emulation of rc-controller sticks
	// using a,w,s,d and arrow keys.
	case ActionPcmdGazInc:
		if d.pcmd.Gaz < 0 {
			d.pcmd.Gaz = 0
		}
		d.pcmd.Flag = 1
		d.pcmd.Gaz++
		d.pcmd.Gaz = d.CheckLimitPcmdField(d.pcmd.Gaz)
		arg := &Ardrone3PilotingPCMDArguments{
			Flag: 1,
			Gaz:  d.pcmd.Gaz,
		}
		d.pcmdPacketSchedulerCh <- packetCreator.encodeCmd(Command(PilotingPCMD), arg)
	case ActionPcmdGazDec:
		if d.pcmd.Gaz > 0 {
			d.pcmd.Gaz = 0
		}
		d.pcmd.Flag = 1
		d.pcmd.Gaz--
		d.pcmd.Gaz = d.CheckLimitPcmdField(d.pcmd.Gaz)
		arg := &Ardrone3PilotingPCMDArguments{
			Flag: 1,
			Gaz:  d.pcmd.Gaz,
		}
		d.pcmdPacketSchedulerCh <- packetCreator.encodeCmd(Command(PilotingPCMD), arg)

	case ActionPcmdYawCounterClockwise:
		if d.pcmd.Yaw > 0 {
			d.pcmd.Yaw = 0
		}
		d.pcmd.Flag = 1
		d.pcmd.Yaw--
		d.pcmd.Yaw = d.CheckLimitPcmdField(d.pcmd.Yaw)
		arg := &Ardrone3PilotingPCMDArguments{
			Flag: 1,
			Yaw:  d.pcmd.Yaw,
		}
		d.pcmdPacketSchedulerCh <- packetCreator.encodeCmd(Command(PilotingPCMD), arg)
	case ActionPcmdYawClockwise:
		if d.pcmd.Yaw < 0 {
			d.pcmd.Yaw = 0
		}
		d.pcmd.Flag = 1
		d.pcmd.Yaw++
		d.pcmd.Yaw = d.CheckLimitPcmdField(d.pcmd.Yaw)
		arg := &Ardrone3PilotingPCMDArguments{
			Flag: 1,
			Yaw:  d.pcmd.Yaw,
		}
		d.pcmdPacketSchedulerCh <- packetCreator.encodeCmd(Command(PilotingPCMD), arg)

	case ActionPcmdHover:
		d.pcmd = Ardrone3PilotingPCMDArguments{
			Flag:               0, // TODO: maybe set this one to ZERO ?
			Gaz:                0,
			Pitch:              0,
			Roll:               0,
			TimestampAndSeqNum: 0,
			Yaw:                0,
		}

		arg := d.pcmd
		d.pcmdPacketSchedulerCh <- packetCreator.encodeCmd(Command(PilotingPCMD), arg)

	case ActionPcmdPitchForward:
		if d.pcmd.Pitch < 0 {
			d.pcmd.Pitch = 0
		}
		d.pcmd.Flag = 1
		d.pcmd.Pitch++
		d.pcmd.Pitch = d.CheckLimitPcmdField(d.pcmd.Pitch)
		arg := &Ardrone3PilotingPCMDArguments{
			Flag:  1,
			Pitch: d.pcmd.Pitch,
		}
		d.pcmdPacketSchedulerCh <- packetCreator.encodeCmd(Command(PilotingPCMD), arg)
	case ActionPcmdPitchBackward:
		if d.pcmd.Pitch > 0 {
			d.pcmd.Pitch = 0
		}
		d.pcmd.Flag = 1
		d.pcmd.Pitch--
		d.pcmd.Pitch = d.CheckLimitPcmdField(d.pcmd.Pitch)
		arg := &Ardrone3PilotingPCMDArguments{
			Flag:  1,
			Pitch: d.pcmd.Pitch,
		}
		d.pcmdPacketSchedulerCh <- packetCreator.encodeCmd(Command(PilotingPCMD), arg)

	case ActionPcmdRollLeft:
		if d.pcmd.Roll > 0 {
			d.pcmd.Roll = 0
		}
		d.pcmd.Flag = 1
		d.pcmd.Roll--
		d.pcmd.Roll = d.CheckLimitPcmdField(d.pcmd.Roll)
		arg := &Ardrone3PilotingPCMDArguments{
			Flag: 1,
			Roll: d.pcmd.Roll,
		}
		d.pcmdPacketSchedulerCh <- packetCreator.encodeCmd(Command(PilotingPCMD), arg)
	case ActionPcmdRollRight:
		if d.pcmd.Roll < 0 {
			d.pcmd.Roll = 0
		}
		d.pcmd.Flag = 1
		d.pcmd.Roll++
		d.pcmd.Roll = d.CheckLimitPcmdField(d.pcmd.Roll)
		arg := &Ardrone3PilotingPCMDArguments{
			Flag: 1,
			Roll: d.pcmd.Roll,
		}
		d.pcmdPacketSchedulerCh <- packetCreator.encodeCmd(Command(PilotingPCMD), arg)
	case ActionPcmdRepeatLastCmd:
		d.pcmdPacketSchedulerCh <- packetCreator.encodeCmd(Command(PilotingPCMD), d.pcmd)

	// --------------moveTo
	// The commands below is a bit overly complicated to use, but they
	// are implemented to manually be able to test out the moveTo feature.
	case ActionMoveToSetLatInc:
		if d.gps.latitudeMoveTo != 500 {
			d.gps.latitudeMoveTo = d.gps.latitudeMoveTo + 0.00001
			log.Printf("moveTo: %#v\n", d.gps)
		} else {
			log.Printf("ActionMoveToLatInc: failed, no connection with GPS: %v\n", d.gps.latitude)
		}
	case ActionMoveToSetLatDec:
		if d.gps.latitudeMoveTo != 500 {
			d.gps.latitudeMoveTo = d.gps.latitudeMoveTo - 0.00001
			log.Printf("moveTo: %#v\n", d.gps)
		} else {
			log.Printf("ActionMoveToLatDec: failed, no connection with GPS: %v\n", d.gps.latitude)
		}
	case ActionMoveToSetLonDec:
		if d.gps.longitudeMoveTo != 500 {
			d.gps.longitudeMoveTo = d.gps.longitudeMoveTo - 0.00001
			log.Printf("moveTo: %#v\n", d.gps)
		} else {
			log.Printf("ActionMoveToLonDec: failed, no connection with GPS: %v\n", d.gps.longitude)
		}
	case ActionMoveToSetLonInc:
		if d.gps.longitudeMoveTo != 500 {
			d.gps.longitudeMoveTo = d.gps.longitudeMoveTo + 0.00001
			log.Printf("moveTo: %#v\n", d.gps)
		} else {
			log.Printf("ActionMoveToLonInc: failed, no connection with GPS: %v\n", d.gps.longitude)
		}
	case ActionMoveToSetBufferCurrentPosition:
		if d.gps.latitude != 500 || d.gps.longitude != 500 {
			d.gps.latitudeMoveTo = d.gps.latitude
			d.gps.longitudeMoveTo = d.gps.longitude
		} else {
			log.Printf("ActionMoveToSetBufferCurrentPosition: failed, no connection with GPS: %v\n", d.gps.latitude)
		}
	case ActionMoveToExecute:
		// TODO:
		// The idea here is to use this action with a moveTo command to the drone,
		// and giving the current moveTo variables as arguments to the moveTo
		// command.

		d.gps.doingMoveTo = true
		d.gps.moveToExecuteCh <- struct{}{}
		// TODO: send the moveTo command here!!!
		log.Printf("*************************************************************\n")
		log.Printf("ActionMoveToExecute: current value of buffer: %#v\n", d.gps)
		log.Printf("*************************************************************\n")
	case ActionMoveToCancel:
		d.gps.doingMoveTo = false
		d.gps.moveToCancelCh <- struct{}{}
	}

}

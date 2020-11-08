package parrotbebop

import (
	"context"
	"log"

	"github.com/eiannone/keyboard"
)

type inputAction int

const (
	// Standard actions.
	//
	ActionPcmdFlag                inputAction = iota
	ActionPcmdRollLeft            inputAction = iota
	ActionPcmdRollRight           inputAction = iota
	ActionPcmdPitchForward        inputAction = iota
	ActionPcmdPitchBackward       inputAction = iota
	ActionPcmdYawClockwise        inputAction = iota
	ActionPcmdYawCounterClockwise inputAction = iota
	ActionPcmdHover               inputAction = iota
	ActionPcmdGazInc              inputAction = iota
	ActionPcmdGazDec              inputAction = iota
	ActionPcmdRepeatLastCmd       inputAction = iota
	ActionTakeoff                 inputAction = iota
	ActionLanding                 inputAction = iota
	ActionEmergency               inputAction = iota
	ActionNavigateHomeStart       inputAction = iota // Check how to implement it in xml line 153
	ActionNavigateHomeStop        inputAction = iota // Check how to implement it in xml line 153
	ActionMoveBy                  inputAction = iota // Check how to implement it in xml line 181
	ActionUserTakeoff             inputAction = iota
	ActionMoveTo                  inputAction = iota // Check how to implement it in xml line 259
	ActionCancelMoveTo            inputAction = iota
	ActionStartPilotedPOI         inputAction = iota
	ActionStopPilotedPOI          inputAction = iota
	ActionCancelMoveBy            inputAction = iota

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
func (d *Drone) readKeyBoardEvent() {

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

	// Since we are resetting the the go routines and thus are stopping
	// the listener of the input action channel we need to check if it
	// is a listener available to avoid deadlock.
	// checkChOpen is a little helper function for just that.
	//
	// NB: This function will drop input actions given if the channel is
	// closed. A benefit of doing just that is that we avoid any commands
	// that might have been given while the connection was gone to
	// suddenly be executed when the connection comes back, but this also
	// implies that we have mechanism's in place to handle continous
	// flight of the drone incase there is a drop, or the connection have
	// to be re-established for some reason
	checkChOpen := func(ch chan inputAction, ia inputAction) {
		select {
		case ch <- ia:
		default:
		}
	}

	for {
		select {
		case event := <-keysEvents:

			if event.Err != nil {
				panic(event.Err)
			}

			switch {
			case event.Key == keyboard.KeyEsc:
				d.chQuit <- struct{}{}
			case event.Rune == 'q':
				// Initiate a reconnect of the network.
				select {
				case d.chNetworkConnect <- struct{}{}:
				default:
				}
			case event.Rune == 't':
				checkChOpen(d.chInputActions, ActionTakeoff)
			case event.Rune == 'l':
				checkChOpen(d.chInputActions, ActionLanding)
			case event.Rune == 'r':
				checkChOpen(d.chInputActions, ActionNavigateHomeStart)
			case event.Rune == 'R':

				checkChOpen(d.chInputActions, ActionNavigateHomeStop)
			case event.Rune == 'w':
				checkChOpen(d.chInputActions, ActionPcmdGazInc)
			case event.Rune == 's':
				checkChOpen(d.chInputActions, ActionPcmdGazDec)
			case event.Rune == 'a':
				checkChOpen(d.chInputActions, ActionPcmdYawCounterClockwise)
			case event.Rune == 'd':
				checkChOpen(d.chInputActions, ActionPcmdYawClockwise)

			case event.Rune == 'h':
				checkChOpen(d.chInputActions, ActionPcmdHover)

			case event.Key == keyboard.KeyArrowUp:
				checkChOpen(d.chInputActions, ActionPcmdPitchForward)
			case event.Key == keyboard.KeyArrowDown:
				checkChOpen(d.chInputActions, ActionPcmdPitchBackward)
			case event.Key == keyboard.KeyArrowLeft:
				checkChOpen(d.chInputActions, ActionPcmdRollLeft)
			case event.Key == keyboard.KeyArrowRight:
				checkChOpen(d.chInputActions, ActionPcmdRollRight)
			case event.Key == keyboard.KeySpace:
				checkChOpen(d.chInputActions, ActionPcmdRepeatLastCmd)
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
// This function will execute the commands that arrives on the d.chInputActions.
func (d *Drone) handleInputAction(packetCreator udpPacketCreator, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("info: exiting handleInputAction")
			return

		case action := <-d.chInputActions:
			switch action {
			case ActionTakeoff:
				p := packetCreator.encodeCmd(Command(PilotingTakeOff), &Ardrone3PilotingTakeOffArguments{})
				d.chSendingUDPPacket <- p
			case ActionLanding:
				p := packetCreator.encodeCmd(Command(PilotingLanding), &Ardrone3PilotingLandingArguments{})
				d.chSendingUDPPacket <- p
			case ActionNavigateHomeStart:
				p := packetCreator.encodeCmd(Command(PilotingNavigateHome), &Ardrone3PilotingNavigateHomeArguments{Start: 1})
				d.chSendingUDPPacket <- p
			case ActionNavigateHomeStop:
				p := packetCreator.encodeCmd(Command(PilotingNavigateHome), &Ardrone3PilotingNavigateHomeArguments{Start: 0})
				d.chSendingUDPPacket <- p

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
				d.chPcmdPacketScheduler <- packetCreator.encodeCmd(Command(PilotingPCMD), arg)
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
				d.chPcmdPacketScheduler <- packetCreator.encodeCmd(Command(PilotingPCMD), arg)

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
				d.chPcmdPacketScheduler <- packetCreator.encodeCmd(Command(PilotingPCMD), arg)
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
				d.chPcmdPacketScheduler <- packetCreator.encodeCmd(Command(PilotingPCMD), arg)

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
				d.chPcmdPacketScheduler <- packetCreator.encodeCmd(Command(PilotingPCMD), arg)

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
				d.chPcmdPacketScheduler <- packetCreator.encodeCmd(Command(PilotingPCMD), arg)
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
				d.chPcmdPacketScheduler <- packetCreator.encodeCmd(Command(PilotingPCMD), arg)

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
				d.chPcmdPacketScheduler <- packetCreator.encodeCmd(Command(PilotingPCMD), arg)
			case ActionPcmdRollRight:
				if d.pcmd.Roll < 0 {
					d.pcmd.Roll = 0
				}
				d.pcmd.Flag = 1
				d.pcmd.Roll--
				d.pcmd.Roll = d.CheckLimitPcmdField(d.pcmd.Roll)
				arg := &Ardrone3PilotingPCMDArguments{
					Flag: 1,
					Roll: d.pcmd.Roll,
				}
				d.chPcmdPacketScheduler <- packetCreator.encodeCmd(Command(PilotingPCMD), arg)

			case ActionPcmdRepeatLastCmd:
				d.chPcmdPacketScheduler <- packetCreator.encodeCmd(Command(PilotingPCMD), d.pcmd)
			}
		}

	}
}

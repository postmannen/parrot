package parrotbebop

import (
	"fmt"
)

// Try to figure out what kind of command that where received.
// Based on the type of cmdArgs we can execute som action.
func (d *Drone) checkCmdFromDrone(cmd protocolARCommands, cmdArgs interface{}) {
	fmt.Printf("----------COMMAND-------------------------------------------\r\n")
	fmt.Printf("-- cmd = %+v\r\n", cmd)
	fmt.Printf("-- Value of cmdArgs = %+v\r\n", cmdArgs)
	fmt.Printf("-- Type of cmdArgs = %+T\r\n", cmdArgs)
	switch cmdArgs := cmdArgs.(type) {
	case Ardrone3CameraStateOrientationArguments:
		//log.Printf("** EXECUTING ACTION FOR TYPE, Ardrone3CameraStateOrientationArguments ...........\r\n")
	case Ardrone3PilotingStateAttitudeChangedArguments:
		//log.Printf("** EXECUTING ACTION FOR TYPE, Ardrone3PilotingStateAttitudeChangedArguments\r\n")
	case Ardrone3PilotingStateGpsLocationChangedArguments:
		d.gps.chCurrentLocation <- gpsLatLonAlt{
			latitude:  cmdArgs.Latitude,
			longitude: cmdArgs.Longitude,
			altitude:  cmdArgs.Altitude,
		}
	case Ardrone3PilotingStatemoveToChanged:
		// Indicated that the drone have moved to the asked position.
		// We send a signal to the moveTo handling here to indicate
		// that it can pick the next available position in the buffer.
		d.gps.chMoveToPositionDone <- struct{}{}
	}
	fmt.Printf("-----------------------------------------------------------\r\n")

}

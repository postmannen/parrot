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
	}
	fmt.Printf("-----------------------------------------------------------\r\n")

}

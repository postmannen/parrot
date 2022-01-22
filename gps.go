package parrot

import (
	"log"
)

// gpsLatLongAlt is used for messaging position data between
// go routines.
type gpsLatLonAlt struct {
	latitude float64
	// Longitude East/West
	longitude float64
	// Altitude height in meters above sea level
	altitude float64
}

// GPS will hold all the current values of the current
// gps location, and also the coordinate to move to
// next if moveTo action have been issued.
type GPS struct {
	currentLocationCh chan gpsLatLonAlt
	// connected ?
	connected bool
	// latitude North/South
	latitude float64
	// Longitude East/West
	longitude float64
	// Altitude height in meters above sea level
	altitude float64

	// latitude North/South
	latitudeMoveTo float64
	// Longitude East/West
	longitudeMoveTo float64
	// Altitude height in meters above sea level
	altitudeMoveto float64
	// Are the drone currently in a moveTo action ?
	// This value should be set to true when a moveTo are started,
	// and it should be set to false when a message from the drone
	// of type Ardrone3PilotingStatemoveToChanged are received.
	doingMoveTo bool
	// Initiate an execution of a moveTo to the next position in buffer.
	moveToExecuteCh chan struct{}
	// Cancel the execution of a moveTo command
	moveToCancelCh chan struct{}
	// When a moveTo command is succesful a Ardrone3PilotingStatePositionChanged
	// command is sent from the drone. In the actionsD2C we will check
	// for such commands and send a signal here, so we know that we
	// can pull the next waypoint.
	movedToPositionCh chan struct{}
}

// StartHandling, start handling incomming gps packages, and fill
// the registers with the current location values.
func (g *GPS) StartReadingPosition() {
	for v := range g.currentLocationCh {
		if v.latitude == 500 || v.longitude == 500 || v.altitude == 500 {
			g.connected = false
		}
		g.latitude = v.latitude
		g.longitude = v.longitude
		g.altitude = v.altitude

		log.Printf("gps location data: %#v\n", g)
	}
}

// for {
// 	wp, err := d.waypointBuffer.pullWayPointNext()
// 	if err != nil {
// 		log.Printf("info: waypointBufferEmpty, breaking out\n")
// 		break
// 	}
//
// 	arg := &Ardrone3PilotingmoveToArguments{
// 		Latitude:  wp.latitude,
// 		Longitude: wp.longitude,
// 		Altitude:  wp.altitude,
// 	}
//
// 	p := packetCreator.encodeCmd(Command(PilotingmoveTo), arg)
// 	d.chSendingUDPPacket <- p
// }

//------------------------------------------

// for {
// 	select {
// 	case <-d.gps.chMoveToCancel:
//
// 		p := packetCreator.encodeCmd(Command(PilotingCancelMoveTo), &// Ardrone3PilotingCancelMoveToArguments{})
// 		d.chSendingUDPPacket <- p
// 		log.Printf("*************************************************************\n")
// 		log.Printf("startMoveToExecutor: chMoveToCancel received\n")
// 		log.Printf("*************************************************************\n")
// 	case <-d.gps.chMoveToExecute:
// 		// TODO:
// 		log.Printf("*************************************************************\n")
// 		log.Printf("startMoveToExecutor: chMoveToExecute received\n")
// 		log.Printf("*************************************************************\n")
// 	}
// }

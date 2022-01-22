package parrot

import (
	"io"
	"log"
	"time"
)

// waypointBuffer holds the buffer of all the waypoints
// and the logic to receive, push and pull waypoints.
type waypointBuffer struct {
	// all the waypoints registered
	waypoints     []gpsLatLonAlt
	waypointInCh  chan gpsLatLonAlt
	waypointOutCh chan gpsLatLonAlt
}

// newMoveToBuffer is a push/pop storage for values for where to
// move to.
func newWaypointBuffer() *waypointBuffer {
	wpBuffer := waypointBuffer{
		waypointInCh: make(chan gpsLatLonAlt),
	}

	// Start the waypointBuffer listener, which basically will start
	// listening on the channel for moveTo messages, and add them
	// to the moveTo buffer
	go wpBuffer.start()

	go func() {
		for {
			wp, err := wpBuffer.pullNext()
			if err != nil {
				log.Printf("info: no way point in buffer, waiting 1 sec, and continue\n")
				time.Sleep(time.Second * 1)
				continue
			}

			// TODO: Might need to add a select with default here
			// incase the channel is not listening
			// or..maybe not since that would cause the wp to be dropped.
			// Need to check this out.
			wpBuffer.waypointOutCh <- wp
		}
	}()

	return &wpBuffer
}

// startWayPointReceiver will check if the wp received
// are within the allowed limits. If OK put it on the
// waypoint buffer, if not we just discard the value
// and wait for the next one.
func (s *waypointBuffer) start() {
	for {
		wp := <-s.waypointInCh
		// Check if the values are to big, which means no GPS connection
		// where available for calculation, and drop the data if it is
		// an not allowed value
		switch {
		case wp.latitude > 91 || wp.latitude < -91:
			log.Printf("waypointBuffer: not allowed value received: %v\n", wp)
			continue
		case wp.longitude > 181 || wp.longitude < -181:
			log.Printf("waypointBuffer: not allowed value received: %v\n", wp)
			continue
		}
		s.pushNew(wp)
	}
}

// push will add another item to the end of the buffer with a normal append
func (s *waypointBuffer) pushNew(d gpsLatLonAlt) {
	s.waypoints = append(s.waypoints, d)
}

// pop will remove and return the first element of the buffer,
// and will return io.EOF if buffer is empty.
func (s *waypointBuffer) pullNext() (gpsLatLonAlt, error) {
	if len(s.waypoints) == 0 {
		return gpsLatLonAlt{}, io.EOF
	}

	v := s.waypoints[0]
	s.waypoints = append(s.waypoints[0:0], s.waypoints[1:]...)

	return v, nil
}

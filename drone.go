// The latest version of the ardrone3.xml document can be found at
// https://github.com/Parrot-Developers/arsdk-xml/tree/master/xml

package parrot

import (
	"context"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

// Drone holds the data and methods specific for the drone.
type Drone struct {
	// The ip address of the drone
	ipAddress string
	// Used for initializing the connection to the drone over TCP.
	portDiscover string
	// Controller to drone, port the controller wil send the drone messages on.
	portC2D string
	// Drone to controller, port the controller will listen on for drone messages.
	portD2C        string
	portRTPStream  string
	portRTPControl string
	// Sending to this channel will quit the controller program.
	quitCh chan struct{}
	// Sending to this channel will disconnect all network related
	// go routines, and then reconnect to the drone.
	networkReconnectCh chan struct{}
	// pcmdPacketSchedulerCh is used to set the frequency of PcmdPacket's
	// that will be sent from the controller to the drone.
	// All Pcmd packets from the controller should go through here to not
	// overwhelm the drone with to many commands which can interupt
	// other commands.
	pcmdPacketSchedulerCh chan networkUDPPacket
	// The conn object for the UDP network listener
	connUDPRead net.PacketConn
	// The conn object for the UDP connection to send commands to
	// the drone.
	connUDPWrite *net.UDPConn
	// Piloting Command
	pcmd Ardrone3PilotingPCMDArguments
	// gps Data
	gps GPS
	// waypointBuffer is a FIFO buffer for storing the gps positions
	// of the route to fly.
	waypointBuffer *waypointBuffer
}

// TODO:
// We can send moveTo messages telling the location to move to with the
// Ardrone3Pilotingmoveto command.
// We can check when that request is fullfilled by checking for a message
// of type Ardrone3PilotingStatemoveToChanged. Maybe need timeout for this ??
// We can then issue the next command.
// For now it seems like we will need a buffer for the moveTo commands, so it
// will pick the next when the previous is done.
// Pressing space should add the next moveTo command to the buffer.
// moveTo paths should be able to be read from file, or other API ? Geofencing ?

// NewDrone will initalize all the variables needed for a drone,
// like ports used, ip adresses, etc.
func NewDrone() *Drone {
	d := &Drone{
		ipAddress:    "192.168.42.1",
		portDiscover: "44444",
		//portC2D:        "54321", // This one is now assigned via discovery
		portD2C:        "43210",
		portRTPStream:  "55004",
		portRTPControl: "55005",

		quitCh:                make(chan struct{}),
		networkReconnectCh:    make(chan struct{}),
		pcmdPacketSchedulerCh: make(chan networkUDPPacket),

		pcmd: Ardrone3PilotingPCMDArguments{
			Flag:               0,
			Roll:               0,
			Pitch:              0,
			Yaw:                0,
			Gaz:                0,
			TimestampAndSeqNum: 0,
		},

		// The default gps values received from the drone when not
		// connected is 500. We set all the values 500 and check
		// later in the code for that value, so we for example don't
		// initiate a moveTo when there is no connection, or add some
		// lat/lon distance if the current register value are 500.
		gps: GPS{
			currentLocationCh: make(chan gpsLatLonAlt),
			connected:         false,
			latitude:          500,
			longitude:         500,
			altitude:          500,
			latitudeMoveTo:    500,
			longitudeMoveTo:   500,
			altitudeMoveto:    500,
		},

		waypointBuffer: newWaypointBuffer(),
	}

	go func() {
		<-d.quitCh
		log.Printf("Operator asked to stop driver.\n")
		os.Exit(0)
	}()

	return d
}

// startMoveToExecutor
// The plan here is to receive a signal for when to execute a
// moveTo command to the drone, or to cancel it.
//
// When a moveto signal is reveived we will pull one waypoint
// at a time from the moveTo buffer, but before pulling a new
// waypoint we will wait for a positiosChanged command from
// the drone, since that will indicate that the last moveTo
// command was executed and done by the drone, and we can pull
// a new value and send another  moveTo package to the drone.
//
// When a cancel signal is received we should immediately send
// a moveTo cancel package to the drone, and also stop any moveTo
// processes.
func (d *Drone) startMoveToExecutor(packetCreator *udpPacketCreator, ctx context.Context) {
	for {
		<-d.gps.moveToExecuteCh
		ctx, cancel := context.WithCancel(ctx)
		var wg sync.WaitGroup
		wg.Add(1)

		go func(ctx context.Context) {
			for {

				ticker := time.NewTicker(time.Second * 5)

				select {
				case <-ctx.Done():
					return
				case <-d.gps.moveToCancelCh:
					p := packetCreator.encodeCmd(Command(PilotingCancelMoveTo), &Ardrone3PilotingCancelMoveToArguments{})
					d.pcmdPacketSchedulerCh <- p
					wg.Done()
				case wp := <-d.waypointBuffer.waypointOutCh:
					// Get a new wp, create the argument, and send the udp packet.
					arg := &Ardrone3PilotingmoveToArguments{
						Latitude:  wp.latitude,
						Longitude: wp.longitude,
						Altitude:  wp.altitude,
					}

					p := packetCreator.encodeCmd(Command(PilotingmoveTo), arg)
					d.pcmdPacketSchedulerCh <- p

					// Check if the waypoint was reached, and we got a confirmation
					// from the drone. If a waypoint is not received we break out,
					// loop and pick a new waypoint.
					select {
					case <-d.gps.movedToPositionCh:
						log.Printf("moveToPositionDone received, breaking out and looping")
						break
					case <-ticker.C:
						log.Printf("moveToPositionDone not received, ticker occured, looping")
						break
					}

				}
			}
		}(ctx)

		wg.Wait()
		cancel()

	}
}

func (d *Drone) Start() {
	// Since we need to use individual sequence number counters for each
	// buffer a udpPacketCreator will keep track of them, and increment
	// the currect buffer sequence number when a new package are created.
	// All UDP packet encoding methods are tied to this type.
	packetCreator := newUdpPacketCreator()

	// Check for keyboard press, and generate appropriate inputActions's.
	go d.readKeyBoardEvent(packetCreator)

	// Start handling incomming gps packages, and fill the registers with
	// the current location values.
	go d.gps.StartReadingPosition()

	for {
		var err error

		ctx, cancel := context.WithCancel(context.Background())

		// Initialize the network connection to the drone.
		// If the connection fails retry 20 times before giving up.
		//
		// TODO:
		// Make it call return-home if unable to initialize.
		log.Println("Initializing the traffic with the drone, and starting controller UDP listener.")
		for {
			err := d.Discover()
			if err != nil {
				log.Printf("error: client Discover failed: %v\n", err)
				time.Sleep(time.Second * 2)
				continue
			}

			// Connection ok, break out of loop.
			break
		}

		// create an 'empty' UDP listener for receiving data from the drone.
		d.connUDPRead, err = net.ListenPacket("udp", ":"+d.portD2C)
		if err != nil {
			log.Println("error: failed to start listener", err)
		}

		// Start the reading of whole UDP packets from the network.
		go d.readNetworkUDPPacketsD2C(ctx, packetCreator)

		// Prepare and dial the UDP connection from controller to drone.
		udpAddr, err := net.ResolveUDPAddr("udp", d.ipAddress+":"+d.portC2D)
		if err != nil {
			log.Printf("error: failed to resolveUDPAddr: %v", err)
		}
		d.connUDPWrite, err = net.DialUDP("udp", nil, udpAddr)
		if err != nil {
			log.Printf("error: failed to DialUDP: %v", err)
		}

		// Start the scheduler for writing UDP packets which will make
		// sure that if there are Pcmd packets to be sent they are only
		// sent at a fixed 25 milli second interval.
		go d.PcmdPacketScheduler(ctx)

		go d.startMoveToExecutor(packetCreator, ctx)

		// Wait here until receiving on quit channel. Trigger by pressing
		// 'q' on the keyboard, or for reconnect to drone.
		<-d.networkReconnectCh
		cancel()
		time.Sleep(time.Second * 3)
		continue
	}
}

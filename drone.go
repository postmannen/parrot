// The latest version of the ardrone3.xml document can be found at
// https://github.com/Parrot-Developers/arsdk-xml/tree/master/xml

package parrotbebop

import (
	"context"
	"log"
	"net"
	"os"
	"time"
)

// Drone holds the data and methods specific for the drone.
type Drone struct {
	// The ip address of the drone
	addressDrone string
	// Used for initializing the connection to the drone over TCP.
	portDiscover string
	// Controller to drone, port the controller wil send the drone messages on.
	portC2D string
	// Drone to controller, port the controller will listen on for drone messages.
	portD2C        string
	portRTPStream  string
	portRTPControl string
	// Channel to put the raw UDP packages from the drone.
	chReceivedUDPPacket chan networkUDPPacket
	// Channel to put the raw UDP packages to be sent to the drone.
	chSendingUDPPacket chan networkUDPPacket
	// Channel to put the inputAction type send to the drone when
	// for example a key is pressed on the keyboard.
	chInputActions chan inputAction
	// Sending to this channel will quit the controller program.
	chQuit chan struct{}
	// Sending to this channel will disconnect all network related
	// go routines, and then reconnect to the drone.
	chNetworkConnect chan struct{}
	// chPcmdPacketScheduler is used to set the frequency of PcmdPacket's
	// that will be sent from the controller to the drone.
	// All Pcmd packets from the controller should go through here to not
	// overwhelm the drone with to many commands which can interupt
	// other commands.
	chPcmdPacketScheduler chan networkUDPPacket
	// The conn object for the UDP network listener
	connUDPRead net.PacketConn
	// The conn object for the UDP connection to send commands to
	// the drone.
	connUDPWrite *net.UDPConn
	// Piloting Command
	pcmd Ardrone3PilotingPCMDArguments
	// gps Data
	gps GPS
}

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
	chCurrentLocation chan gpsLatLonAlt
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
}

// StartHandling, start handling incomming gps packages, and fill
// the registers with the current location values.
func (g *GPS) StartHandling() {
	for v := range g.chCurrentLocation {
		if v.latitude == 500 || v.longitude == 500 || v.altitude == 500 {
			g.connected = false
		}
		g.latitude = v.latitude
		g.longitude = v.longitude
		g.altitude = v.altitude

		log.Printf("gps location data: %#v\n", g)
	}
}

// NewDrone will initalize all the variables needed for a drone,
// like ports used, ip adresses, etc.
func NewDrone() *Drone {
	d := &Drone{
		addressDrone: "192.168.42.1",
		portDiscover: "44444",
		//portC2D:        "54321", // This one is now assigned via discovery
		portD2C:        "43210",
		portRTPStream:  "55004",
		portRTPControl: "55005",

		chReceivedUDPPacket:   make(chan networkUDPPacket),
		chSendingUDPPacket:    make(chan networkUDPPacket),
		chInputActions:        make(chan inputAction),
		chQuit:                make(chan struct{}),
		chNetworkConnect:      make(chan struct{}),
		chPcmdPacketScheduler: make(chan networkUDPPacket),

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
			chCurrentLocation: make(chan gpsLatLonAlt),
			connected:         false,
			latitude:          500,
			longitude:         500,
			altitude:          500,
			latitudeMoveTo:    500,
			longitudeMoveTo:   500,
			altitudeMoveto:    500,
		},
	}

	go func() {
		<-d.chQuit
		log.Printf("Operator asked to stop driver.\n")
		os.Exit(0)
	}()

	return d
}

func (d *Drone) Start() {
	// Check for keyboard press, and generate appropriate inputActions's.
	go d.readKeyBoardEvent()

	// Start handling incomming gps packages, and fill the registers with
	// the current location values.
	go d.gps.StartHandling()

	for {
		var err error

		// Since we need to use individual sequence number counters for each
		// buffer a udpPacketCreator will keep track of them, and increment
		// the currect buffer sequence number when a new package are created.
		// All UDP packet encoding methods are tied to this type.
		packetCreator := newUdpPacketCreator()

		ctxBg := context.Background()
		ctx, cancel := context.WithCancel(ctxBg)

		// Will handle all the events generated by input actions from keyboard etc.
		go d.handleInputAction(*packetCreator, ctx)

		// Initialize the network connection to the drone.
		// If the connection fails retry 20 times before giving up.
		//
		// TODO:
		// Make it call return-home if unable to initialize.
		log.Println("Initializing the traffic with the drone, and starting controller UDP listener.")
		for i := 0; i < 20; i++ {
			err := d.Discover()
			if err != nil {
				log.Printf("error: client Discover failed: %v\n", err)
				time.Sleep(time.Second * 2)
				continue
			}

			break
		}

		// create an 'empty' UDP listener.
		d.connUDPRead, err = net.ListenPacket("udp", ":"+d.portD2C)
		if err != nil {
			log.Println("error: failed to start listener", err)
		}

		// Start the reading of whole UDP packets from the network,
		// and put them on the Drone.chReceivedUDPPacket channel.
		go d.readNetworkUDPPacketsD2C(ctx)

		// Prepare and dial the UDP connection from controller to drone.
		udpAddr, err := net.ResolveUDPAddr("udp", d.addressDrone+":"+d.portC2D)
		if err != nil {
			log.Printf("error: failed to resolveUDPAddr: %v", err)
		}
		d.connUDPWrite, err = net.DialUDP("udp", nil, udpAddr)
		if err != nil {
			log.Printf("error: failed to DialUDP: %v", err)
		}

		// Start the scheduler which will make sure that if there are
		// Pcmd packets to be sent, they are only sent at a fixed 50
		// milli second interval.
		go d.PcmdPacketScheduler(ctx)

		// Start the sender of UDP packets,
		// will send UDP packets received at the Drone.chSendingUDPPacket
		// channel.
		go d.writeNetworkUDPPacketsC2D(ctx)

		go d.handleReadPackages(packetCreator, ctx)

		// Wait here until receiving on quit channel. Trigger by pressing
		// 'q' on the keyboard.
		<-d.chNetworkConnect
		cancel()
		time.Sleep(time.Second * 3)
		continue

	}
}

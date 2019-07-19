package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

//Drone holds the data and methods specific for the drone
type Drone struct {
	droneAddress     string //The ip address of the drone
	portDiscover     string //Used for initializing the connection to the drone over TCP
	portD2C          string //Drone to controller, port the controller will listen on for drone messages
	portRTPStream    string
	portRTPControl   string
	chReceivedPacket chan networkPacket
}

//NewDrone will initalize all the variables needed for a drone,
// like ports used, ip adresses, etc.
func NewDrone() *Drone {
	return &Drone{
		droneAddress:   "192.168.42.1",
		portDiscover:   "44444",
		portD2C:        "43210",
		portRTPStream:  "55004",
		portRTPControl: "55005",

		chReceivedPacket: make(chan networkPacket),
	}
}

//Discover will initalize the connection with the drone.
func (d *Drone) Discover() error {
	//const addr = "192.168.42.1:44444"

	discoverConn, err := net.Dial("tcp", d.droneAddress+":"+d.portDiscover)

	if err != nil {
		return err
	}

	_, err = discoverConn.Write(
		[]byte(
			fmt.Sprintf(`{
						"controller_type": "computer",
						"controller_name": "go-bebop",
						"d2c_port": "%s",
						"arstream2_client_stream_port": "%s",
						"arstream2_client_control_port": "%s",
						}`,
				d.portD2C,
				d.portRTPStream,
				d.portRTPControl),
		),
	)
	if err != nil {
		log.Println("error: Discover, discoveryClient.Write: ", err)
	}

	data := make([]byte, 10240)

	_, err = discoverConn.Read(data)

	if err != nil {
		return err
	}

	return discoverConn.Close()
}

//networkPacket (ARNetworkAL), a network packet is the main UDP packet read from the network.
// A network packet can contain multiple ARNetwork frames.
type networkPacket struct {
	size int
	data []byte
}

//getNetworkPacketsD2C gets the raw UDP packets from the drone sent to the controller.
// Will read the raw UDP packets from the network, and put them on a channel to be
// picked up by the frame decoder.
func (d *Drone) readNetworkPacketsD2C() {
	//create an 'empty' UDP listener.
	localConn, err := net.ListenPacket("udp", ":43210")
	if err != nil {
		log.Println("error: failed to start listener", err)
	}

	defer localConn.Close()

	for {
		buf := make([]byte, 1024)

		n, addr, err := localConn.ReadFrom(buf)
		if err != nil {
			log.Printf("error: failed ReadFrom: %v\n", err)
		}
		log.Printf("info from readNetworkPacketsD2C: read %v bytes, from %v\n", n, addr)

		p := networkPacket{
			size: n,
			data: buf,
		}

		//send the packet received over a channel to later parse out frames.
		d.chReceivedPacket <- p

	}
}

//networkFrame
// A network frame (ARNetworkAL)looks like this, and in the following order :
// dataType 1 byte,
// targetBufferID 1 byte,
// sequeneNumber 1 Byte,
// frameSize 4 Bytes (little endian),
// data n bytes (this is the actual drone data ARNetwork),
type networkFrame struct {
	dataType       int
	targetBufferID int
	sequenceNR     int
	size           int
	data           []byte
}

//decodeNetworkFrame will decode one single frame of the ARNetworkAL protocol.
// NB: If the packet consists of more than one frame eg. the length of the
// packet is longer than the length of the frame, then the function must be
// run again
func decodeNetworkFrame(buf []byte, startPosition int) networkFrame {
	frame := networkFrame{
		dataType:       int(buf[0]),
		targetBufferID: int(buf[1]),
		sequenceNR:     int(buf[2]),
		data:           []byte{},
	}

	var size uint32
	err := binary.Read(bytes.NewReader(buf[3:7]), binary.LittleEndian, &size)
	if err != nil {
		log.Println("error: NewNetworkFrame, binary.Read: ", err)
	}
	frame.size = int(size)
	fmt.Println("Info:decodeNetworkFrame: size of whole frame = ", frame.size+7)

	frame.data = buf[7:frame.size]

	return frame
}

func main() {
	drone := NewDrone()

	err := drone.Discover()
	if err != nil {
		log.Println("error: client Discover failed:", err)
	}

	//Will start the reading of whole UDP packets from the network,
	// and put them on the chReceivedPacket channel.
	go drone.readNetworkPacketsD2C()

	for {
		packet := <-drone.chReceivedPacket
		fmt.Println("info: main: packet size = ", packet.size)
		//fmt.Println("info: main: packet data = ", packet.data)

	}

	//time.Sleep(time.Second * 2)
}

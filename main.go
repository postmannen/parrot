package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
)

//Drone holds the data and methods specific for the drone
type Drone struct {
	addressDrone        string //The ip address of the drone
	portDiscover        string //Used for initializing the connection to the drone over TCP
	portC2D             string //TODO: Make this one to be filled with port value from discover.
	portD2C             string //Drone to controller, port the controller will listen on for drone messages
	portRTPStream       string
	portRTPControl      string
	chReceivedUDPPacket chan networkUDPPacket
	chSendingUDPPacket  chan networkUDPPacket
}

//NewDrone will initalize all the variables needed for a drone,
// like ports used, ip adresses, etc.
func NewDrone() *Drone {
	return &Drone{
		addressDrone: "192.168.42.1",
		portDiscover: "44444",
		//portC2D:        "54321", //This one is now assigned via discovery
		portD2C:        "43210",
		portRTPStream:  "55004",
		portRTPControl: "55005",

		chReceivedUDPPacket: make(chan networkUDPPacket),
		chSendingUDPPacket:  make(chan networkUDPPacket)}
}

//Discover will initalize the connection with the drone.
// A discover will JSON formated data like :
//
// { "status": 0, "c2d_port": 54321, "c2d_update_port": 51, "c2d_user_port": 21, "qos_mode": 0, "arstream2_server_stream_port": 5004, "arstream2_server_control_port": 5005 }
func (d *Drone) Discover() error {
	//const addr = "192.168.42.1:44444"

	discoverConn, err := net.Dial("tcp", d.addressDrone+":"+d.portDiscover)

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

	data := make([]byte, 1024) //not quite sure about the size here...

	_, err = discoverConn.Read(data)
	if err != nil {
		return err
	}
	fmt.Printf("*** Discovery data \n %v \n\n, Size of data = %v\n", string(data), len(data))

	//Using anonymous struct just for unmarshalling the discoveryData
	discoverData := struct {
		Status                     int `json:"status"`
		C2dPort                    int `json:"c2d_port"`
		C2dUpdate                  int `json:"c2d_update_port"`
		C2dUserPort                int `json:"c2d_user_port"`
		QosMode                    int `json:"qos_mode"`
		Arstream2ServerStreamPort  int `json:"arstream2_server_stream_port"`
		Arstream2ServerControlPort int `json:"arstream2_server_control_port"`
	}{}

	//Remove all the zero allocations in the byte slice, else unmarshal will fail.
	data = bytes.Trim(data, "\x00")

	if err := json.Unmarshal(data, &discoverData); err != nil {
		log.Println("error:Umarshal discovery data: ", err)
	}
	fmt.Println("Unmarshaled : ", discoverData)

	//if the status !=0 the disovery failed.
	if discoverData.Status != 0 {
		log.Fatal("DISCOVERY FAILED") //TODO: put in a timer and make it retry after a few seconds.
	}

	d.portC2D = strconv.Itoa(discoverData.C2dPort)
	fmt.Printf("portC2D is of type = %T, and value = %s \n", d.portC2D, d.portC2D)

	return discoverConn.Close()
}

//getNetworkPacketsD2C gets the raw UDP packets from the drone sent to the controller.
// Will read the raw UDP packets from the network, and put them on a channel to be
// picked up by the frame decoder.
func (d *Drone) readNetworkUDPPacketsD2C() {
	//create an 'empty' UDP listener.
	localConn, err := net.ListenPacket("udp", ":43210")
	if err != nil {
		log.Println("error: failed to start listener", err)
	}

	defer localConn.Close()

	for {
		buf := make([]byte, 16384) //NB: buf might be to small ?

		n, addr, err := localConn.ReadFrom(buf)
		if err != nil {
			log.Printf("error: failed ReadFrom: %v\n", err)
		}
		log.Printf("info from readNetworkPacketsD2C: read %v bytes, from %v\n", n, addr)

		p := networkUDPPacket{
			size: n,
			data: buf,
			//Since this is a new UDP packet, and we want to start reading
			// the first frame from the start we set the start position to 0.
			framePos: 0,
		}

		//send the packet received over a channel to later parse out ARNetworkAL/frames.
		d.chReceivedUDPPacket <- p

	}
}

//writeNetworkPacketsC2D writes the raw UDP packets from the controller to the drone.
// Will receive []byte packet to write on an incomming channel for the function.
func (d *Drone) writeNetworkPacketsC2D() {
	//TODO:
	//Have a channel waiting for incomming network packets to be written to the drone
	// .........
}

//networkFrame
// A network frame (ARNetworkAL)looks like this, and in the following order :
// - dataType 1 byte,
// - targetBufferID 1 byte,
// - sequeneNumber 1 Byte,
// - frameSize 4 Bytes (little endian) for the whole ARNetworkAL frame including 7bit header,
// - data n bytes (this is the actual drone data ARNetwork),
//
//	Example of size:
//	01 ba 27 08000000 42, 02 0b c3 0b000000 12345678
//  --size 0x08=8byte---, --size 0x0b=11byte--------
type networkFrame struct {
	//Data types
	//• Ack(1): Acknowledgment of previously received data
	//  To Ack a frame, set type to 1,
	//  add 128 to the value of the bufferID of the package that requires Ack,
	//	new unique sequence nr. for the ack buffer,
	//
	//• Data(2): Normal data (no ack requested)
	//• Low latency data(3): Treated as normal data on the network, but are
	//  given higher priority internally
	//• Data with ack(4): Data requesting an ack. The receiver must send an
	//  ack for this data !
	dataType int
	//• [0; 9]: Reserved values for ARNetwork internal use.
	//• [10; 127]: Data buffers.
	//• [128; 255]: Acknowledge buffers.
	targetBufferID int
	sequenceNR     int
	size           int
	dataARNetwork  []byte
}

//networkPacket is the main UDP packet read from the network.
// A network packet can contain multiple ARNetworkAL/frames.
type networkUDPPacket struct {
	//The total size of the UDP packet
	size int
	//The actual UDP data
	data []byte
	//Where to start reading. If there is only one ARNetworkAL frame in the
	// UDP packet this value will be 0. If there are more than one frame in
	// the packet the value will be set to the start position of the next
	// frame in the slice.
	framePos int
}

func encodeNetworkFrame(dataType int, targetBufferID int, sequenceNR int, size int, dataARNetworkAL []byte) {
	// TODO:.........................
}

//decodeARNetworkALpacket will decode a whole UDP packet given as input,
// and return a frame of the ARNetworkAL protocol, it will return error==
// io.EOF when decoding of the whole packet is done.
// If the there are more than one ARNetworkAL frame in the UDP packet the
// method will return error == nil, and the method should be run over again
// until io.EOF is received.
func (packet *networkUDPPacket) decode() (networkFrame, error) {
	frame := networkFrame{
		dataType:       int(packet.data[packet.framePos+0]),
		targetBufferID: int(packet.data[packet.framePos+1]),
		sequenceNR:     int(packet.data[packet.framePos+2]),
		dataARNetwork:  []byte{},
	}

	//Get the size of the ARNetworkAL frame. Size includes the header of 7bytes.
	var size uint32
	err := binary.Read(bytes.NewReader(packet.data[packet.framePos+3:packet.framePos+7]), binary.LittleEndian, &size)
	if err != nil {
		log.Println("error: NewNetworkFrame, binary.Read: ", err)
	}
	frame.size = int(size)
	frame.dataARNetwork = packet.data[packet.framePos+7 : packet.framePos+frame.size]

	//Figure out if there are another frame after this one.
	// This can be checked if there are a complete header
	// of 7bytes following directly afte the current frame.
	const headerSize = 7
	fmt.Println("---- frameStartPos+frame.size+headerSize = ", packet.framePos+frame.size+headerSize)
	fmt.Println("---- packet.size = ", packet.size)
	if packet.framePos+frame.size+headerSize <= packet.size {
		fmt.Println("----### ANOTHER PACKAGE FOLLOWS at position = ", packet.framePos+frame.size)
		fmt.Println("---- Next package is = ", packet.data[15:packet.size])
		packet.framePos = packet.framePos + frame.size

		return frame, nil

	}

	return frame, io.EOF
}

func main() {
	drone := NewDrone()

	err := drone.Discover()
	if err != nil {
		log.Println("error: client Discover failed:", err)
	}

	//Will start the reading of whole UDP packets from the network,
	// and put them on the chReceivedPacket channel.
	go drone.readNetworkUDPPacketsD2C()

	for {
		//Get a packet
		packet := <-drone.chReceivedUDPPacket
		fmt.Println("-----------------------------------------------------------")
		fmt.Println("info: main: packet size = ", packet.size)
		fmt.Println("info: main: packet data ARNetworkAL= ", packet.data[:packet.size])

		for {
			frame, err := packet.decode()
			//• Ack(1): Acknowledgment of previously received data
			//• Data(2): Normal data (no ack requested)
			//• Low latency data(3): Treated as normal data on the network, but are
			//  given higher priority internally
			//• Data with ack(4): Data requesting an ack. The receiver must send an
			//  ack for this data !
			fmt.Println("info: main: frame: data type: ", frame.dataType)
			fmt.Println("info: main: frame: target buffer id: ", frame.targetBufferID)
			fmt.Println("info: main: frame: size of current frame = ", frame.size)
			fmt.Println("info: main: frame: data_ARNetwork = ", frame.dataARNetwork)
			fmt.Println("-----------------------------------------------------------")

			if err == io.EOF {
				break
			}
		}
	}

	//time.Sleep(time.Second * 2)
}

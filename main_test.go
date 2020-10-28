package main

import (
	"io"
	"log"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	// TODO: Rewrite this one to reflect how the main code works now.
	// As it is now it is just a copy of how main used to be, and that
	// gives no value

	for n := 0; n < b.N; n++ {
		// Get a packet
		buf := []byte{
			2, 127, 8, 23, 0, 0, 0, 1, 4, 6, 0, 154, 221, 45, 61, 44, 209, 73, 188, 121, 230, 52, 64}

		packet := networkUDPPacket{
			size: len(buf),
			data: buf,
			// Since this is a new UDP packet, and we want to start reading
			// the first frame from the start we set the start position to 0.
			framePos: 0,
		}

		//fmt.Println("-----------------------------------------------------------")
		//fmt.Println("info: main: packet size = ", packet.size)
		//fmt.Println("info: main: packet data ARNetworkAL= ", packet.data[:packet.size])

		var lastFrame bool
		for {
			frameARNetworkAL, err := packet.decode()
			// Check if it was the last frame in the UDP packet.
			if err == io.EOF {
				lastFrame = true
			}
			// • Ack(1): Acknowledgment of previously received data
			// • Data(2): Normal data (no ack requested)
			// • Low latency data(3): Treated as normal data on the network, but are
			//   given higher priority internally
			// • Data with ack(4): Data requesting an ack. The receiver must send an
			//   ack for this data !
			//fmt.Println("info: main: frame: data type: ", frameARNetworkAL.dataType)
			//fmt.Println("info: main: frame: target buffer id: ", frameARNetworkAL.targetBufferID)
			//fmt.Println("info: main: frame: size of current frame = ", frameARNetworkAL.size)
			//fmt.Println("info: main: frame: data_ARNetwork = ", frameARNetworkAL.dataARNetwork)
			//fmt.Println("-----------------------------------------------------------")

			// TODO: Putting in a continue here to skip decoding of ping packets
			// which have a buffer ID of 0 or 1.
			// Replace this with a proper ping detection later.
			pingDetected := false
			if frameARNetworkAL.targetBufferID == 0 || frameARNetworkAL.targetBufferID == 1 {
				//	fmt.Println("PING DETECTED, PING DETECTED, PING DETECTED,PING DETECTED,")
				//	fmt.Println("NOT DECODING THE CONTENT OF THE FRAME")
				pingDetected = true
			}

			// If the package was not a ping package, then decode the ARCommand
			// from it.
			if !pingDetected {
				// TODO: Put in a select here on the cmd type to do some further processing on the packages
				_, _, err := frameARNetworkAL.decode()
				if err != nil {
					log.Println("error: frame.decode: ", err)
					break
				}
				//fmt.Println("----------COMMAND-------------------------------------------")
				//fmt.Printf("%+v\n", cmd)
				//fmt.Println("-----------------------------------------------------------")

			}

			// If no more frames, break out.
			if lastFrame {
				break
			}

			//	fmt.Println("-------------Working the next frame in the UDP package-----")
		}
	}

	//time.Sleep(time.Second * 2)

}

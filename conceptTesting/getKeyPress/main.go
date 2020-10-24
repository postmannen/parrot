package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	// fd 0 is stdin
	state, err := terminal.MakeRaw(0)
	if err != nil {
		log.Fatalln("setting stdin to raw:", err)
	}
	defer func() {
		if err := terminal.Restore(0, state); err != nil {
			log.Println("warning, failed to restore terminal:", err)
		}
	}()

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err != nil {
			log.Println("stdin:", err)
			break
		}

		if r == '\x1b' {
			r, _, _ := in.ReadRune()
			if r == '[' {
				r, _, _ := in.ReadRune()
				switch r {
				case 'A':
					fmt.Printf("UP ARROW\r\n")
				case 'B':
					fmt.Printf("DOWN ARROW\r\n")
				case 'C':
					fmt.Printf("RIGHT ARROW\r\n")
				case 'D':
					fmt.Printf("LEFT ARROW\r\n")
				}

			}
		}

		fmt.Printf("read rune %q\r\n", r)
		if r == 'q' {
			break
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"go.bug.st/serial"
)

// This program simulates a serial device that sends frames of data at regular intervals.
// It listens for user input to exit the loop gracefully.
// To run this program, you need to have a virtual serial port set up.
// You can use `socat` to create virtual serial ports:
// `socat -dd pty,rawer,echo=0,link=/tmp/ttyV0 pty,rawer,echo=0,link=/tmp/ttyV1`.
func main() {
	device := "/tmp/ttyV0"
	datasets := []string{
		"\nADCO 123456789000 D\r",
		"\nOPTARIF BBR( S\r",
		"\nISOUSC 30 9\r",
		"\nBBRHCJB 015736771 B\r",
		"\nBBRHPJB 002110855 @\r",
		"\nBBRHCJW 000242993 O\r",
		"\nBBRHPJW 000436619 \\\r",
		"\nBBRHCJR 000080750 A\r",
		"\nBBRHPJR 000067067 T\r",
		"\nPTEC HPJB P\r",
		"\nDEMAIN ---- \"\r",
		"\nIINST 001 X\r",
		"\nIMAX 090 H\r",
		"\nPAPP 00280 +\r",
		"\nHHPHC A ,\r",
		"\nMOTDETAT 000000 B\r",
	}

	// Open the serial port
	mode := &serial.Mode{
		BaudRate: 1200,
		Parity:   serial.EvenParity,
		DataBits: 7,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open(device, mode)
	if err != nil {
		log.Fatalf("Failed to open the serial port %q [%s]\n=> Launch `socat -dd pty,rawer,echo=0,link=/tmp/ttyV0 pty,rawer,echo=0,link=/tmp/ttyV1` to create the virtual serial ports.", device, err)
	}
	defer port.Close()

	quit := make(chan bool)

	// Goroutine to watch for 'Q' enter key to quit
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if strings.EqualFold(input, "Q") {
				quit <- true
				return
			}
		}
	}()

	fmt.Println("Loop is running. Enter 'Q' to exit!")

	i := 0
	_, err = port.Write([]byte{0x02}) // Send "Start TeXt" STX (0x02) to start the frame
	if err != nil {
		log.Fatal(err)
	}

	for {
		if i >= len(datasets) {
			i = 0                                   // Reset index to loop through frames again
			_, err = port.Write([]byte{0x03, 0x02}) // Send "End TeXt" ETX (0x03) to stop the frame and STX char to start the next frame
			if err != nil {
				log.Fatal(err)
			}
		}

		select {
		case <-quit:
			fmt.Println("Exiting...")
			return
		default:
			_, err := port.Write([]byte(datasets[i]))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(250 * time.Millisecond)
			i++
		}
	}
}

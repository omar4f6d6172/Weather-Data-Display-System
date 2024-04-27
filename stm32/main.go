package main

import (
    "machine"
    "time"
)

// City holds temperature and location information
type City struct {
    Name        string
    Latitude    float64
    Longitude   float64
    Temperature float64
}

// Define global variables for the SPI and GPIO pins
var (
    spi    *machine.SPI
    rs     machine.Pin
    csb    machine.Pin
)

func main() {
    // Initialize the SPI
    spi = machine.SPI1
    spi.Configure(machine.SPIConfig{
        SCK: machine.PA5,
        SDI: machine.PA7,
    })

    // Setup GPIO pins for RS and CSB
    rs = machine.PA9
    csb = machine.PA8
    rs.Configure(machine.PinConfig{Mode: machine.PinOutput})
    csb.Configure(machine.PinConfig{Mode: machine.PinOutput})

    // Initialize the display
    initDisplay()

    // Print "Hello, World!" every one second
    for {
        writeString("Hello, World!")
        time.Sleep(time.Second)
    }
}

// initDisplay sends commands to the display to initialize it
func initDisplay() {
    csb.Low()
    rs.Low() // Command mode
    spi.Tx([]byte{0x38}, nil) // Function Set
    spi.Tx([]byte{0x0C}, nil) // Display ON
}
// writeString sends a string to the display
func writeString(message string) {
	csb.Low()
	for _, char := range message {
		rs.High() // Set RS to data mode
		spi.Tx([]byte{byte(char)}, nil)
		time.Sleep(2 * time.Millisecond) // Wait for character to be written
	}
	csb.High()
}

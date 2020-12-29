package main

import (
	"github.com/tarm/serial"
	"testing"
)

func Test_Write(t *testing.T) {
	display := Display{
		//Config: serial.Config{Name: "/dev/ttyUSB0", Baud: 19200}, // Linux
		Config: serial.Config{Name: "COM6", Baud: 19200}, // Windows
	}

	display.Open()
	display.ClearDisplay()
	display.Write("LCD2021")
	display.Close()
}

func Test_CursorBlinkOff(t *testing.T) {
	display := Display{
		//Config: serial.Config{Name: "/dev/ttyUSB0", Baud: 19200}, // Linux
		Config: serial.Config{Name: "COM6", Baud: 19200}, // Windows
	}

	display.Open()
	display.ClearDisplay()
	display.CursorBlinkOff()
	display.Write("LCD2021")
	display.Close()
}

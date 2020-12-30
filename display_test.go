package main

import (
	"github.com/tarm/serial"
	"testing"
	"time"
)

func Test_Write(t *testing.T) {
	display := Display{
		//Config: serial.Config{Name: "/dev/ttyUSB0", Baud: 19200}, // Linux
		Config: serial.Config{Name: "COM6", Baud: 19200}, // Windows
	}

	display.Open()
	defer display.Close()
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
	defer display.Close()
	display.ClearDisplay()
	display.CursorBlinkOff()
	display.Write("LCD2021")
}

func Test_Clock(t *testing.T) {
	display := Display{
		//Config: serial.Config{Name: "/dev/ttyUSB0", Baud: 19200}, // Linux
		Config: serial.Config{Name: "COM6", Baud: 19200}, // Windows
	}

	display.Open()
	defer display.Close()
	display.ClearDisplay()
	display.CursorBlinkOff()
	for true {
		display.ClearDisplay()
		clock := time.Now()
		display.Write(clock.Format("2006-01-02 15:04:05"))
		time.Sleep(250 * time.Millisecond)
	}
}

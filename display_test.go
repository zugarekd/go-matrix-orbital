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
	display.Clear()
	//display.BlinkOff()
	//display.BackLightOn(0)
	display.Write("DanZ")
	//display.BlinkOn()
	//display.HorizontalBarGraph()
	//display.Contrast(50)
	display.Close()
}

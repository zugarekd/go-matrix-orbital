package main

import (
	"encoding/hex"
	"github.com/tarm/serial"
)

func main() {
	display := Display{
		Config: serial.Config{Name: "/dev/ttyUSB0", Baud: 19200},
	}

	display.Open()
	display.Clear()
	display.BlinkOff()
	//display.BackLightOn(0)
	display.Write("DanZ")
	display.BlinkOn()
	display.HorizontalBarGraph()
	//display.Contrast(50)
	display.Close()
}

const (
	CommandPrefix = "FE"
)

type Display struct {
	Config serial.Config
	port   serial.Port
}

func (d *Display) Open() error {
	a, err := serial.OpenPort(&d.Config)
	d.port = *a
	return err
}

func (d *Display) Close() error {
	err := d.port.Close()
	return err
}

func (d *Display) Write(data string) error {
	_, err := d.port.Write([]byte(data))
	return err
}

func (d *Display) Clear() error {
	command, _ := hex.DecodeString(CommandPrefix)
	command = append(command, []byte("X")...)
	_, err := d.port.Write([]byte(command))
	return err
}

func (d *Display) BackLightOn(time byte) error {
	command, _ := hex.DecodeString(CommandPrefix)
	command = append(command, []byte("B")...)
	command = append(command, time)
	_, err := d.port.Write([]byte(command))
	return err
}

func (d *Display) Contrast(contrast byte) error {
	command, _ := hex.DecodeString(CommandPrefix)
	command = append(command, []byte("P")...)
	command = append(command, contrast)
	_, err := d.port.Write([]byte(command))
	return err
}

func (d *Display) CursorOff() error {
	command, _ := hex.DecodeString(CommandPrefix)
	command = append(command, []byte("K")...)
	_, err := d.port.Write([]byte(command))
	return err
}

func (d *Display) CursorOn() error {
	command, _ := hex.DecodeString(CommandPrefix)
	command = append(command, []byte("J")...)
	_, err := d.port.Write([]byte(command))
	return err
}

func (d *Display) BlinkOff() error {
	command, _ := hex.DecodeString(CommandPrefix)
	command = append(command, []byte("T")...)
	_, err := d.port.Write([]byte(command))
	return err
}

func (d *Display) BlinkOn() error {
	command, _ := hex.DecodeString(CommandPrefix)
	command = append(command, []byte("S")...)
	_, err := d.port.Write([]byte(command))
	return err
}

func (d *Display) HorizontalBarGraph() error {
	command, _ := hex.DecodeString(CommandPrefix)
	command = append(command, []byte("h")...)
	_, err := d.port.Write([]byte(command))

	command, _ = hex.DecodeString(CommandPrefix + "7C0701")
	command = append(command, 1)
	length, _ := hex.DecodeString("64")
	command = append(command, length...)
	_, err = d.port.Write([]byte(command))

	command, _ = hex.DecodeString(CommandPrefix + "7C1002")
	command = append(command, 1)
	length, _ = hex.DecodeString("64")
	command = append(command, length...)
	_, err = d.port.Write([]byte(command))

	return err
}

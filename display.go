package main

import (
	"github.com/tarm/serial"
)

const (
	CommandPrefix   = "\xFE"
	AutoLineWrapOn  = CommandPrefix + "\x43"
	AutoLineWrapOff = CommandPrefix + "\x44"
	AutoScrollOn    = CommandPrefix + "\x51"
	AutoScrollOff   = CommandPrefix + "\x52"
	BacklightOn     = CommandPrefix + "\x42"
	BacklightOff    = CommandPrefix + "\x46"
	ClearDisplay    = CommandPrefix + "\x58"
	Contrast        = CommandPrefix + "\x50"
	CursorOn        = CommandPrefix + "\x4A"
	CursorOff       = CommandPrefix + "\x4B"
	CursorLeft      = CommandPrefix + "\x4C"
	CursorRight     = CommandPrefix + "\x4D"
	CursorBlinkOn   = CommandPrefix + "\x53"
	CursorBlinkOff  = CommandPrefix + "\x54"
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

// The current cursor position will be incremented for each character received.
// Please note that unless line wrap is turned on, the text will follow the memory map of the module.
func (d *Display) Write(data string) error {
	_, err := d.port.Write([]byte(data))
	return err
}

// This command enables the automatic line wrap function. Transmitted characters which overrun the width of the display
// will automatically wrap to the next line. The bottom line wraps to line 1 of the display. To carry out this command the user
// must send a command prefix followed the character 'C'.
func (d *Display) AutoLineWrapOn() error {
	_, err := d.port.Write([]byte(AutoLineWrapOn))
	return err
}

// This command disables the automatic line wrapping function. To turn off the auto line wrapping, send a command prefix
// followed by the character 'D'.
func (d *Display) AutoLineWrapOff() error {
	_, err := d.port.Write([]byte(AutoLineWrapOff))
	return err
}

// To turn the automatic line scrolling on, send a command prefix followed by the character 'Q'. In combination with the
// “Auto Line Wrap” command the text will automatically wrap and scroll up.
// Note: “Auto Line Wrap” must be enabled for “Auto Scroll” to work properly.
func (d *Display) AutoScrollOn() error {
	_, err := d.port.Write([]byte(AutoScrollOn))
	return err
}

// To turn the automatic line scrolling off, send a command prefix followed by the character 'R'.
func (d *Display) AutoScrollOff() error {
	_, err := d.port.Write([]byte(AutoScrollOff))
	return err
}

// To turn the backlight on, send a command prefix followed by the character 'B' as well as the number of minutes for the
// backlight to be activated. If <minutes> is sent as zero then the backlight will remain on indefinitely. The maximum value
// for <minutes> is 100.
func (d *Display) BacklightOn(time byte) error {
	command := []byte(BacklightOn)
	command = append(command, time)
	_, err := d.port.Write([]byte(command))
	return err
}

// To turn the backlight off, send a command prefix followed by the character 'F'.
func (d *Display) BacklightOff() error {
	_, err := d.port.Write([]byte(BacklightOff))
	return err
}

// This command clears any text and graphics off the display. To clear the display, send a command prefix followed the
// character 'X'.
func (d *Display) ClearDisplay() error {
	_, err := d.port.Write([]byte(ClearDisplay))
	return err
}

// This command allows you to set the display contrast to a level between 0(light) and 256(dark). To execute this command,
// send a command prefix followed by the character 'P' and a hex value between 0x00 and 0xFF. Different displays and lighting
// conditions will affect the actual value used. Different modules have different power up contrast settings
func (d *Display) Contrast(contrast byte) error {
	command := []byte(Contrast)
	command = append(command, contrast)
	_, err := d.port.Write([]byte(command))
	return err
}

// To turn the cursor on at the current position, send a command prefix followed by the character 'J'. Note cursor is on by
// default at power up.
func (d *Display) CursorOn() error {
	_, err := d.port.Write([]byte(CursorOn))
	return err
}

// To turn the cursor off at the current position , send a command prefix followed by the character 'K'.
func (d *Display) CursorOff() error {
	_, err := d.port.Write([]byte(CursorOff))
	return err
}

// To move the cursor one space to the left of current position, send a command prefix followed by the character 'L'.
func (d *Display) CursorLeft() error {
	_, err := d.port.Write([]byte(CursorLeft))
	return err
}

// To move the cursor one space to the right of the current position, send a command prefix followed by the character 'M'.
func (d *Display) CursorRight() error {
	_, err := d.port.Write([]byte(CursorRight))
	return err
}

// To turn on the blinking cursor at the current position, send a command prefix followed by the character 'S'. Please note that
// the blinking cursor is on by default at power up.
func (d *Display) CursorBlinkOn() error {
	_, err := d.port.Write([]byte(CursorBlinkOn))
	return err
}

// To turn off the blinking cursor at the current position send a command prefix followed by the character 'T'.
func (d *Display) CursorBlinkOff() error {
	_, err := d.port.Write([]byte(CursorBlinkOff))
	return err
}

/*
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
*/

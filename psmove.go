package psmove

/*
#cgo pkg-config: psmoveapi
#include "psmove.h"
*/
import "C"

import (
	"errors"
)

type Move struct {
	pointer *C.PSMove
}

func Connect() (Move, error) {
	move := Move{C.psmove_connect()}
	if move.pointer == nil {
		return move, errors.New("Could not connect to default Move controller")
	}
	return move, nil
}

func (move Move) Disconnect() {
	C.psmove_disconnect(move.pointer)
}

type ConnectionType uint

const (
	ConnectionTypeBluetooth ConnectionType = iota
	ConnectionTypeUSB
	ConnectionTypeUnknown
)

func (move Move) ConnectionType() ConnectionType {
	switch C.psmove_connection_type(move.pointer) {
	case C.Conn_Bluetooth:
		return ConnectionTypeBluetooth
	case C.Conn_USB:
		return ConnectionTypeUSB
	default:
		return ConnectionTypeUnknown
	}
}

func (move Move) SetLEDs(r, g, b byte) {
	C.psmove_set_leds(move.pointer, C.uchar(r), C.uchar(g), C.uchar(b))
}

func (move Move) SetRumble(value byte) {
	C.psmove_set_rumble(move.pointer, C.uchar(value))
}

func (move Move) UpdateLEDs() {
	C.psmove_update_leds(move.pointer)
}

func (move Move) Poll() bool {
	return C.psmove_poll(move.pointer) > 0
}

func (move Move) getButtons() C.uint {
	bits := C.psmove_get_buttons(move.pointer)
	return bits
}

func (move Move) IsTriggerPressed() bool {
	return (move.getButtons() & C.Btn_T) != 0
}

func (move Move) IsSquarePressed() bool {
	return (move.getButtons() & C.Btn_SQUARE) != 0
}

func (move Move) IsTrianglePressed() bool {
	return (move.getButtons() & C.Btn_TRIANGLE) != 0
}

func (move Move) IsCrossPressed() bool {
	return (move.getButtons() & C.Btn_CROSS) != 0
}

func (move Move) IsCirclePressed() bool {
	return (move.getButtons() & C.Btn_CIRCLE) != 0
}

func (move Move) IsPSButtonPressed() bool {
	return (move.getButtons() & C.Btn_PS) != 0
}

func (move Move) GetTriggerValue() byte {
	return byte(C.psmove_get_trigger(move.pointer))
}

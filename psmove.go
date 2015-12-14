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

func (move Move) ConnectionType() string {
	switch C.psmove_connection_type(move.pointer) {
	case C.Conn_Bluetooth:
		return "Bluetooth"
	case C.Conn_USB:
		return "USB"
	default:
		return "unknown"
	}
}

func (move Move) SetLEDs(r, g, b int) {
	C.psmove_set_leds(move.pointer, C.uchar(r), C.uchar(g), C.uchar(b))
}

func (move Move) UpdateLEDs() {
	C.psmove_update_leds(move.pointer)
}

package main

import (
	"unicode"
)

// Coord is a coordinate object.
type Coord struct {
	x, y int
}

// MakeCoord creates a new coordinate object.
func MakeCoord(x rune, y int) Coord {
	return Coord{y, int(unicode.ToUpper(x)) - int('A')}
}

// Slot represents the board, each slot having a reference to the ship placed
// upon it.
type Slot struct {
	shipRef *Ship
}

// Board is the game board which houses all of the ships.
type Board struct {
	slots [8][8]*Slot
}

// NewBoard creates a new board pointer
func NewBoard() *Board {
	b := new(Board)
	for _, arr := range b.slots {
		for i := range arr {
			arr[i] = new(Slot)
		}
	}
	return b
}

// SetSlot sets the ship at the current slot.
func (b *Board) SetSlot(slot Coord, shipRef *Ship) {
	temp := new(Slot)
	temp.shipRef = shipRef
	b.slots[slot.x][slot.y] = temp
}

// PingSlot pings the slot to see if there is a ship to hit.
func (b *Board) PingSlot(target Coord) bool {
	slot := b.slots[target.x][target.y]
	if slot != nil {
		pingedShip := slot.shipRef
		if pingedShip != nil {
			pingedShip.Hit()
			slot.shipRef = nil
			return true
		}
	}
	return false
}

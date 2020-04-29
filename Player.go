package main

// Player is an object which represents the user of the game.
type Player struct {
	name   string
	ship   *Ship
	board  *Board
	target Coord
}

// NewPlayer creates a new player with a given name.
func NewPlayer(name string) *Player {
	return &Player{name, NewShip(3), NewBoard(), MakeCoord('A', -1)}
}

// GetName gets the player's name.
func (p *Player) GetName() string {
	return p.name
}

// GetShip get the player's ship.
func (p *Player) GetShip() *Ship {
	return p.ship
}

// GetBoard get the player's ship.
func (p *Player) GetBoard() *Board {
	return p.board
}

// PlaceShip wraps the board set slot method to place the ship on the given
// coordinate.
func (p *Player) PlaceShip(slot Coord) {
	p.GetBoard().SetSlot(slot, p.ship)
}

// CheckForHit wraps the board ping method to check and report back if a shot
// hit or not.
func (p *Player) CheckForHit(target Coord) bool {
	return p.GetBoard().PingSlot(target)
}

// CheckIfSunk checks to see if the player's ship is still alive.
func (p *Player) CheckIfSunk() bool {
	return p.GetShip().GetHitpoints() == 0
}

// SetTarget sets the coordinates at which the player wishes to shoot.
func (p *Player) SetTarget(target Coord) {
	p.target = target
}

// Shoot fires at the player's target.
func (p *Player) Shoot() Coord {
	return p.target
}

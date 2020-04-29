package main

// Ship is an object owned by a player.
type Ship struct {
	hitpoints int
}

// NewShip creates a new ship pointer.
func NewShip(health int) *Ship {
	return &Ship{health}
}

// Hit registers a hit on the ship.
func (s *Ship) Hit() {
	s.hitpoints--
}

// GetHitpoints returns the total hitpoints remaining of the ship.
func (s *Ship) GetHitpoints() int {
	return s.hitpoints
}

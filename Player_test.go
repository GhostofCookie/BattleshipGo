package main

import "testing"

// TestShootAt tests the player's ability to shoot at a slot.
func TestShootAt(t *testing.T) {
	setup := func() (*Player, *Player) {
		p1 := NewPlayer("Tester1")
		p2 := NewPlayer("Tester2")

		p1.PlaceShip(MakeCoord('A', 0))
		p1.PlaceShip(MakeCoord('A', 1))
		p1.PlaceShip(MakeCoord('A', 2))

		return p1, p2
	}
	t.Run("TestMiss", func(t *testing.T) {
		p1, p2 := setup()
		p2.SetTarget(MakeCoord('B', 0))
		hitResult := p1.CheckForHit(p2.Shoot())
		if hitResult {
			t.Fail()
		}
	})
	t.Run("TestHit", func(t *testing.T) {
		p1, p2 := setup()
		p2.SetTarget(MakeCoord('A', 0))
		hitResult := p1.CheckForHit(p2.Shoot())
		if !hitResult {
			t.Fail()
		}
	})
	t.Run("TestSink", func(t *testing.T) {
		p1, p2 := setup()
		for i := 0; i < 3; i++ {
			p2.SetTarget(MakeCoord('A', i))
			hitResult := p1.CheckForHit(p2.Shoot())
			if !hitResult {
				t.Fail()
			}
		}
		sinkResult := p1.CheckIfSunk()
		if !sinkResult {
			t.Fail()
		}
	})
}

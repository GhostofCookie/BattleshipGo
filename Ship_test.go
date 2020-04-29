package main

import "testing"

func TestHit(t *testing.T) {
	testHealth := 3
	s := NewShip(testHealth)

	if s.GetHitpoints() != testHealth {
		t.Fail()
	}
	s.Hit()
	if s.GetHitpoints() == testHealth {
		t.Fail()
	}
}

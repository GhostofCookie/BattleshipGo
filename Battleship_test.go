package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestEnterNames(t *testing.T) {
	testGameInput := "T1\nT2\nA\n0\nA 0\n0\nA 1\n0\nB 1\n1\nA 1\nB 0\nB 1\nA 0\nC 1\n"
	temp := os.Stdout
	os.Stdout = nil
	g := new(Battleship)
	g.reader = bufio.NewReader(strings.NewReader(testGameInput))
	g.EnterNames()
	g.PlaceShips()
	g.Run()
	os.Stdout = temp

	if g.currentPlayer.GetName() != "T1" {
		t.Fail()
	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
	"unicode"
)

func main() {
	game := new(Battleship)
	game.reader = bufio.NewReader(os.Stdin)
	ClearScreen()
	game.EnterNames()
	ClearScreen()
	game.PlaceShips()
	game.Run()
}

// Battleship is an object which represents the game. It has a reader member
// which allows for IOC to make the class testable.
type Battleship struct {
	reader                    io.Reader
	currentPlayer, nextPlayer *Player
}

// EnterNames creates a dialog for creating the players
func (game *Battleship) EnterNames() {
	var name string
	fmt.Print("|-> Enter Player 1 name: ")
	fmt.Fscanf(game.reader, "%s\n", &name)
	game.currentPlayer = NewPlayer(name)
	fmt.Print("|-> Enter Player 2 name: ")
	fmt.Fscanf(game.reader, "%s\n", &name)
	game.nextPlayer = NewPlayer(name)
}

// PlaceShips runs through each player and asks them where to place their ships.
func (game *Battleship) PlaceShips() {
	/* Place ships */
	for i := 0; i < 2; i++ {
		fmt.Println(game.currentPlayer.GetBoard().Output(true))
		fmt.Printf("|-> %s place your ship (e.g. A 0): ", game.currentPlayer.GetName())
		var (
			x            rune
			y            int
			isHorizontal bool
		)
		_, err := fmt.Fscanf(game.reader, "%c %d\n", &x, &y)
		x = unicode.ToUpper(x)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("|-> Vertically or Horizontally (0 or 1): ")
		_, err = fmt.Fscanf(game.reader, "%t\n", &isHorizontal)
		if err != nil ||
			y < 0 || y > 7 || x < 'A' || x > 'H' ||
			(isHorizontal && x == 'A') || (!isHorizontal && y == 0) ||
			(isHorizontal && x == 'H') || (!isHorizontal && y == 7) {
			fmt.Println("<!> Invalid placement. Ships are place from the center out")
			i--
		} else {
			if isHorizontal {
				for i := 0; i < game.currentPlayer.GetShip().GetHitpoints(); i++ {
					game.currentPlayer.PlaceShip(MakeCoord(rune(i+int(x-'A')-1)+'A', y))
				}
			} else {
				for i := 0; i < game.currentPlayer.GetShip().GetHitpoints(); i++ {
					game.currentPlayer.PlaceShip(MakeCoord(x, (i + y - 1)))
				}
			}
			game.EndTurn()
		}
	}
}

// Run is the looping method that runs the game of shooting turn by turn.
func (game *Battleship) Run() {
	// Run the game
	for {
		// Start turn
		var (
			y int
			x rune
		)
		// Output boards
		fmt.Println(game.nextPlayer.GetBoard().Output(false))
		fmt.Println(game.currentPlayer.GetBoard().Output(true))

		fmt.Print("|-> " + game.currentPlayer.GetName() + " select your target: ")
		fmt.Fscanf(game.reader, "%c %d\n", &x, &y)

		game.currentPlayer.SetTarget(MakeCoord(x, y))
		if game.nextPlayer.CheckForHit(game.currentPlayer.Shoot()) {
			fmt.Println("<!> " + game.nextPlayer.GetName() + " reported a HIT")
		} else {
			fmt.Println("<!> " + game.nextPlayer.GetName() + " reported a MISS")
		}
		if game.nextPlayer.CheckIfSunk() {
			ClearScreen()
			fmt.Println(game.nextPlayer.GetBoard().Output(false))
			fmt.Println(game.currentPlayer.GetBoard().Output(true))
			break
		}
		time.Sleep(3 * time.Second)
		ClearScreen()
		fmt.Printf("<!> Pass control over to %s", game.nextPlayer.GetName())

		/* End the turn */
		time.Sleep(2 * time.Second)
		game.EndTurn()
	}

	fmt.Println("\n<!> " + game.nextPlayer.GetName() + " reported ship SUNK !")
	fmt.Println("=== " + game.currentPlayer.GetName() + " wins ===")
}

// EndTurn swaps the players.
func (game *Battleship) EndTurn() {
	temp := *game.currentPlayer
	*game.currentPlayer = *game.nextPlayer
	*game.nextPlayer = temp

	ClearScreen()
}

// ClearScreen clears the screen and re-prints the title.
func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("=== Battleship ===")
}

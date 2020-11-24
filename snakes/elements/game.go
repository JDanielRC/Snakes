package elements

import (
	"github.com/hajimehoshi/ebiten"
)

// Game contains everything needed for the game
type Game struct {
	gui   *GUI
	score int
}

func start() Game {
	game := Game{
		score: 0,
	}
	game.gui = initializeGUI()
	return game
}

// Update proceeds the game state.
func (g *Game) Update() error {
	// Write your game's logical update.
	return nil
}

// Draw interface
func (g *Game) Draw(screen *ebiten.Image) error {
	if err := g.gui.Draw(screen); err != nil {
		return err
	}
	return nil
}

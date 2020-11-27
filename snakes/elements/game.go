package elements

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

// Game contains everything needed for the game
type Game struct {
	gui        *GUI
	score      int
	foods      []*Food
	foodAmount int
	dotTime    int
}

// Start will begin all variables needed for the game
func Start() Game {
	game := Game{
		score:      0,
		foodAmount: 5,
		dotTime:    0,
	}
	foodArray := make([]*Food, game.foodAmount)
	for i := 0; i < game.foodAmount; i++ {
		foodArray[i] = RandomFood(&game)
		time.Sleep(10)
	}
	game.foods = foodArray
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
	for i := 0; i < len(g.foods); i++ {
		if err := g.foods[i].Draw(screen, g.dotTime); err != nil {
			return err
		}
	}
	return nil
}

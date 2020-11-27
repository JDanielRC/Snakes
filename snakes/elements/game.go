package elements

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
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
func Start(nFood int) Game {
	game := Game{
		score:      0,
		foodAmount: nFood,
		dotTime:    0,
	}
	foodArray := make([]*Food, game.foodAmount)
	for i := 0; i < game.foodAmount; i++ {
		foodArray[i] = RandomFood(&game)
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

// Draw interface, this follows a hierarchy, so snakes have to go at last, while background would have to be the first
func (g *Game) Draw(screen *ebiten.Image) error {
	drawer := &ebiten.DrawImageOptions{}
	drawer.GeoM.Translate(0, 0)
	background, _, _ := ebitenutil.NewImageFromFile("files/background.png", ebiten.FilterDefault)
	screen.DrawImage(background, drawer)
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

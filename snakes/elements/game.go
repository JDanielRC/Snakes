package elements

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Game contains everything needed for the game
type Game struct {
	player        *Player
	playerChannel chan int
	gui           *GUI
	score         int
	foods         []*Food
	foodAmount    int
	dotTime       int
	alive         bool
}

// Start will begin all variables needed for the game
func Start(nFood int) Game {
	game := Game{
		score:      0,
		foodAmount: nFood,
		dotTime:    0,
		alive:      true,
	}
	foodArray := make([]*Food, game.foodAmount)
	for i := 0; i < game.foodAmount; i++ {
		foodArray[i] = RandomFood(&game)
	}
	game.foods = foodArray
	game.player = Spawn(&game)
	game.playerChannel = make(chan int)
	go game.player.Behaviour()
	game.gui = initializeGUI()
	return game
}

// Update proceeds the game state.
func (g *Game) Update() error {
	if g.alive {
		if g.foodAmount == 0 {
			g.alive = false
		}
		g.dotTime = (g.dotTime + 1) % 10 //game speed

		if err := g.player.Update(g.dotTime); err != nil {
			g.playerChannel <- g.dotTime
		}

		x, y := g.player.getHead()

		for i := 0; i < len(g.foods); i++ {
			if x == g.foods[i].foodX && y == g.foods[i].foodY {
				g.foods[i].foodX = -20
				g.foods[i].foodY = -20 //set them out of screen
				g.gui.ateFood()
				g.foodAmount--
				g.player.ateFood()
				break
			}
		}

	}
	for i := 0; i < g.foodAmount; i++ {
		if err := g.foods[i].Update(g.dotTime); err != nil {
			return err
		}
	}
	// Write your game's logical update.
	return nil
}

// GameOver ends the game
func (g *Game) GameOver() {
	fmt.Printf("Moriste! \n")
	g.alive = false
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

	if err := g.player.Draw(screen, g.dotTime); err != nil {
		return err
	}

	for i := 0; i < len(g.foods); i++ {
		if err := g.foods[i].Draw(screen, g.dotTime); err != nil {
			return err
		}
	}

	return nil
}

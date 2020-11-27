package main

import (
	"elements/elements"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
)

var game elements.Game

func init() {
	game = elements.Start()
}

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	if err := game.Update(); err != nil {
		return err
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	if err := game.Draw(screen); err != nil {
		fmt.Println(err)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 720, 600
}

func main() {
	game := &Game{}

	ebiten.SetWindowSize(720, 600)
	ebiten.SetWindowTitle("Snake")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}

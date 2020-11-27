package elements

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font/inconsolata"
)

// GUI contains all info for the screen
type GUI struct {
	score            int
	remainingEnemies int
	mainSnakeHealth  int
}

func initializeGUI() *GUI {
	gui := GUI{
		score:            0,
		remainingEnemies: 0,
		mainSnakeHealth:  0,
	}
	return &gui
}

func (gui *GUI) ateFood() {
	gui.score++
}

// Draw GUI
func (gui *GUI) Draw(screen *ebiten.Image) error {
	text.Draw(screen, "Current Score: "+strconv.Itoa(gui.score), inconsolata.Bold8x16, 90, 20, color.Black)
	text.Draw(screen, "Remaining Health: "+strconv.Itoa(gui.mainSnakeHealth), inconsolata.Bold8x16, 410, 20, color.Black)
	text.Draw(screen, "Remaining enemies: "+strconv.Itoa(gui.remainingEnemies), inconsolata.Bold8x16, 780, 20, color.Black)
	return nil
}

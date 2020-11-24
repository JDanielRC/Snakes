package elements

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font/basicfont"
)

// GUI contains all info for the screen
type GUI struct {
	score int
}

func initializeGUI() *GUI {
	gui := GUI{
		score: 0,
	}
	return &gui
}

// Draw GUI
func (g *GUI) Draw(screen *ebiten.Image) error {
	text.Draw(screen, "Score: "+strconv.Itoa(g.score), basicfont.Face7x13, 20, 20, color.White)
	return nil
}

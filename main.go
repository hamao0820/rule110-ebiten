package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 720
)

type Game struct {
	cells []*Cell
}

func newGame() *Game {
	cells := make([]*Cell, 0)
	for i := 1; i < ScreenWidth; i += cellSize {
		for j := 1; j < ScreenHeight; j += cellSize {
			cells = append(cells, newCell(i, j))
		}
	}
	return &Game{
		cells: cells,
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, cell := range g.cells {
		cell.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	g := newGame()
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

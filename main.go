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
	for y := 1; y < ScreenHeight; y += cellSize {
		for x := 1; x < ScreenWidth; x += cellSize {
			cell := newCell(y, x)
			if y == 1 {
				cell.canClick = true
			}
			cells = append(cells, cell)
		}
	}
	return &Game{
		cells: cells,
	}
}

func (g *Game) Update() error {
	for _, cell := range g.cells {
		cell.Update()
	}
	ebiten.SetCursorShape(ebiten.CursorShapeNotAllowed)
	for _, cell := range g.cells {
		if cell.canClick && cell.hovered {
			ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
			break
		}
	}
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

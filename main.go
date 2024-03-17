package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 720
)

const cellSize = 20

type Cell struct {
	x      int
	y      int
	width  int
	height int
}

func newCell(x, y int) *Cell {
	return &Cell{
		x:      x,
		y:      y,
		width:  cellSize,
		height: cellSize,
	}

}

func (c *Cell) Draw(screen *ebiten.Image) {
	// 四辺を描画
	vector.StrokeLine(screen, float32(c.x), float32(c.y), float32(c.x+c.width), float32(c.y), 1, color.White, false)
	vector.StrokeLine(screen, float32(c.x), float32(c.y), float32(c.x), float32(c.y+c.height), 1, color.White, false)
	vector.StrokeLine(screen, float32(c.x+c.width), float32(c.y), float32(c.x+c.width), float32(c.y+c.height), 1, color.White, false)
	vector.StrokeLine(screen, float32(c.x), float32(c.y+c.height), float32(c.x+c.width), float32(c.y+c.height), 1, color.White, false)
}

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

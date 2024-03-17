package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

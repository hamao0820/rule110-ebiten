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
	bg     color.Color
}

func newCell(x, y int) *Cell {
	return &Cell{
		x:      x,
		y:      y,
		width:  cellSize,
		height: cellSize,
		bg:     color.RGBA{0, 0, 0, 0},
	}
}

func (c *Cell) Draw(screen *ebiten.Image) {
	// 四辺を描画
	vector.StrokeLine(screen, float32(c.x), float32(c.y), float32(c.x+c.width), float32(c.y), 1, color.White, false)
	vector.StrokeLine(screen, float32(c.x), float32(c.y), float32(c.x), float32(c.y+c.height), 1, color.White, false)
	vector.StrokeLine(screen, float32(c.x+c.width), float32(c.y), float32(c.x+c.width), float32(c.y+c.height), 1, color.White, false)
	vector.StrokeLine(screen, float32(c.x), float32(c.y+c.height), float32(c.x+c.width), float32(c.y+c.height), 1, color.White, false)
	// 背景を描画
	vector.DrawFilledRect(screen, float32(c.x), float32(c.y), float32(c.width), float32(c.height), c.bg, false)
}

func (c *Cell) Update() {
	x, y := ebiten.CursorPosition()
	if c.isIn(x, y) {
		c.bg = color.RGBA{128, 128, 128, 255}
	} else {
		c.bg = color.RGBA{0, 0, 0, 0}
	}
}

func (c *Cell) isIn(x, y int) bool {
	if c.x < x && x < c.x+c.width && c.y < y && y < c.y+c.height {
		return true
	}
	return false
}

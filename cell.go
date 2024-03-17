package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const cellSize = 10

type Bit int

const (
	Bit0 Bit = iota
	Bit1
)

var (
	bit0Color    = color.RGBA{0, 0, 0, 0}
	bit1Color    = color.RGBA{255, 255, 255, 255}
	hoveredColor = color.RGBA{128, 128, 128, 255}
)

type Cell struct {
	x        int
	y        int
	width    int
	height   int
	value    Bit
	hovered  bool
	canClick bool
}

func newCell(y, x int) *Cell {
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
	// 背景を描画
	if c.value == Bit1 {
		vector.DrawFilledRect(screen, float32(c.x), float32(c.y), float32(c.width), float32(c.height), bit1Color, false)
	} else {
		if c.hovered {
			vector.DrawFilledRect(screen, float32(c.x), float32(c.y), float32(c.width), float32(c.height), hoveredColor, false)
		} else {
			vector.DrawFilledRect(screen, float32(c.x), float32(c.y), float32(c.width), float32(c.height), bit0Color, false)
		}

	}
}

func (c *Cell) Update() {
	x, y := ebiten.CursorPosition()

	// onHover
	if c.isIn(x, y) {
		c.hovered = true
	} else {
		c.hovered = false
	}

	// onClick
	if c.canClick {
		if c.isIn(x, y) {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				if c.value == Bit0 {
					c.value = Bit1
				} else {
					c.value = Bit0
				}
			}
		}
	}
}

func (c *Cell) isIn(x, y int) bool {
	if c.x < x && x < c.x+c.width && c.y < y && y < c.y+c.height {
		return true
	}
	return false
}

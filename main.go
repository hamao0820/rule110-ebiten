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
	startButton *Button
	stopButton  *Button
	stepButton  *Button
	resetButton *Button
	cells       []*Cell
}

func newGame() *Game {
	startButton := newButton(10, 10, 60, 30, "Start")
	stopButton := newButton(10, 80, 60, 30, "Stop")
	stopButton.canClick = false
	stepButton := newButton(10, 150, 60, 30, "Step")
	restButton := newButton(10, 220, 60, 30, "Reset")
	cells := make([]*Cell, 0)
	for i := 0; ; i++ {
		y := 1 + 3*cellSize + i*cellSize
		if y >= ScreenHeight {
			break
		}
		for j := 0; ; j++ {
			x := 1 + j*cellSize
			if x >= ScreenWidth {
				break
			}
			cell := newCell(y, x)
			if i == 0 {
				cell.canClick = true
			}
			cells = append(cells, cell)
		}
	}

	return &Game{
		startButton: startButton,
		stopButton:  stopButton,
		stepButton:  stepButton,
		resetButton: restButton,
		cells:       cells,
	}
}

func (g *Game) Update() error {
	for _, cell := range g.cells {
		cell.Update()
	}

	// マウスカーソルの形状を設定
	ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	for _, cell := range g.cells {
		if cell.hovered {
			ebiten.SetCursorShape(ebiten.CursorShapeNotAllowed)
		}
		if cell.canClick && cell.hovered {
			ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
			break
		}
	}
	if g.startButton.hovered {
		if g.startButton.canClick {
			ebiten.SetCursorShape(ebiten.CursorShapePointer)
		} else {
			ebiten.SetCursorShape(ebiten.CursorShapeNotAllowed)
		}
	}
	if g.stopButton.hovered {
		if g.stopButton.canClick {
			ebiten.SetCursorShape(ebiten.CursorShapePointer)
		} else {
			ebiten.SetCursorShape(ebiten.CursorShapeNotAllowed)
		}
	}
	if g.stepButton.hovered {
		if g.stepButton.canClick {
			ebiten.SetCursorShape(ebiten.CursorShapePointer)
		} else {
			ebiten.SetCursorShape(ebiten.CursorShapeNotAllowed)
		}
	}
	if g.resetButton.hovered {
		if g.resetButton.canClick {
			ebiten.SetCursorShape(ebiten.CursorShapePointer)
		} else {
			ebiten.SetCursorShape(ebiten.CursorShapeNotAllowed)
		}
	}

	g.startButton.Update()
	g.stopButton.Update()
	g.stepButton.Update()
	g.resetButton.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.startButton.Draw(screen)
	g.stopButton.Draw(screen)
	g.stepButton.Draw(screen)
	g.resetButton.Draw(screen)
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
	ebiten.SetWindowTitle("Rule 110")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

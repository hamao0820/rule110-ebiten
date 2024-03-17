package main

import (
	"fmt"
	"log"
	"rule110/automaton"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
)

type Mode int

const (
	ModeInitializing Mode = iota
	ModeRunning
	ModePosed
	ModeEnded
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
	mode        Mode
	row, col    int
	automaton   *automaton.Automaton
	ticks       int
}

func newGame() *Game {
	startButton := newButton(10, 10, 60, 30, "Start")
	stopButton := newButton(10, 80, 60, 30, "Stop")
	stopButton.canClick = false
	stepButton := newButton(10, 150, 60, 30, "Step")
	restButton := newButton(10, 220, 60, 30, "Reset")
	var row, col int
	cells := make([]*Cell, 0)
	for i := 0; ; i++ {
		y := 51 + i*cellSize
		if y >= ScreenHeight {
			break
		}
		row++
		col_ := 0
		for j := 0; ; j++ {
			x := 1 + j*cellSize
			if x >= ScreenWidth {
				break
			}
			col_++
			cell := newCell(y, x)
			if i == 0 {
				cell.canClick = true
			}
			cells = append(cells, cell)
		}
		if i == 0 {
			col = col_
		}
	}

	return &Game{
		startButton: startButton,
		stopButton:  stopButton,
		stepButton:  stepButton,
		resetButton: restButton,
		cells:       cells,
		mode:        ModeInitializing,
		row:         row,
		col:         col,
	}
}

func (g *Game) initAutomaton() {
	initial := make([]uint, g.col)
	for i, cell := range g.cells[:g.col] {
		if cell.value == Bit1 {
			initial[i] = 1
		}
	}
	g.automaton = automaton.NewAutomaton(initial)
}

func (g *Game) updateAutomaton() {
	n := g.automaton.Update()
	for i := 0; i < g.col; i++ {
		if n*g.col+i >= len(g.cells) {
			g.mode = ModeEnded
			break
		}
		g.cells[n*g.col+i].value = Bit(g.automaton.State[i])
	}
}

func (g *Game) Update() error {
	for _, cell := range g.cells {
		cell.Update()
	}

	g.startButton.Update()
	g.stopButton.Update()
	g.stepButton.Update()
	g.resetButton.Update()

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
	switch g.mode {
	case ModeInitializing:
		if g.startButton.clicked {
			g.mode = ModeRunning
			g.startButton.canClick = false
			g.stopButton.canClick = true
			g.stepButton.canClick = false
			for i := range g.cells {
				g.cells[i].canClick = false
			}
			g.initAutomaton()
		}
		if g.stepButton.clicked {
			g.initAutomaton()
			g.updateAutomaton()
			g.mode = ModePosed
		}
	case ModeRunning:
		if g.stopButton.clicked {
			g.mode = ModePosed
			g.startButton.canClick = true
			g.stopButton.canClick = false
			g.stepButton.canClick = true
		}
		g.ticks++
		if g.ticks%15 == 0 {
			g.updateAutomaton()
		}
	case ModePosed:
		if g.startButton.clicked {
			g.mode = ModeRunning
			g.startButton.canClick = false
			g.stopButton.canClick = true
			g.stepButton.canClick = false
		}
		if g.stepButton.clicked {
			g.updateAutomaton()
		}
	}

	if g.resetButton.clicked {
		g.mode = ModeInitializing
		g.startButton.canClick = true
		g.stopButton.canClick = false
		g.stepButton.canClick = true
		cells := make([]*Cell, 0)
		for i := 0; ; i++ {
			y := 51 + i*cellSize
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
		g.cells = cells
	}

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
	bds, _ := font.BoundString(mPlusNormalFont, "Rule 110")
	messageWidth := (bds.Max.X - bds.Min.X).Ceil()
	ebitenutil.DebugPrintAt(screen, "Rule 110", ScreenWidth/2-messageWidth/2, 15)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS()), ScreenWidth-100, 8)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()), ScreenWidth-100, 24)
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

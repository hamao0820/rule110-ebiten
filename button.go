package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var mPlusNormalFont font.Face

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		panic(err)
	}
	const dpi = 72
	mPlusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}
}

type Button struct {
	x       int
	y       int
	width   int
	height  int
	message string

	hovered  bool
	canClick bool
	clicked  bool
}

func newButton(y, x, width, height int, message string) *Button {
	return &Button{
		x:        x,
		y:        y,
		width:    width,
		height:   height,
		message:  message,
		canClick: true,
	}
}

func (b *Button) Update() {
	x, y := ebiten.CursorPosition()

	// onHover
	if b.isIn(x, y) {
		b.hovered = true
	} else {
		b.hovered = false
	}

	// onClick
	b.clicked = false
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if b.isIn(x, y) {
			b.clicked = true
		}
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	// 四辺を描画
	vector.StrokeLine(screen, float32(b.x), float32(b.y), float32(b.x+b.width), float32(b.y), 1, color.White, false)
	vector.StrokeLine(screen, float32(b.x), float32(b.y), float32(b.x), float32(b.y+b.height), 1, color.White, false)
	vector.StrokeLine(screen, float32(b.x+b.width), float32(b.y), float32(b.x+b.width), float32(b.y+b.height), 1, color.White, false)
	vector.StrokeLine(screen, float32(b.x), float32(b.y+b.height), float32(b.x+b.width), float32(b.y+b.height), 1, color.White, false)
	// 背景を描画
	if b.hovered {
		vector.DrawFilledRect(screen, float32(b.x), float32(b.y), float32(b.width), float32(b.height), hoveredColor, false)
	}

	// メッセージを描画
	bds, _ := font.BoundString(mPlusNormalFont, b.message)
	stringWidth := (bds.Max.X - bds.Min.X).Ceil()
	stringHeight := (bds.Max.Y - bds.Min.Y).Ceil()
	text.Draw(screen, b.message, mPlusNormalFont, b.x+(b.width-stringWidth)/2, b.y+(b.height-stringHeight)/2+stringHeight, color.White)
}

func (b *Button) isIn(x, y int) bool {
	if b.x <= x && x <= b.x+b.width && b.y <= y && y <= b.y+b.height {
		return true
	}
	return false
}

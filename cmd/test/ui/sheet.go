package ui

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sheet struct {
	cs  int // rectangular cells
	img *ebiten.Image
}

func NewSheet(img *ebiten.Image, cs int) *Sheet {
	return &Sheet{
		img: img,
		cs:  cs,
	}
}

func (s *Sheet) Get(x, y int) *ebiten.Image {
	rect := image.Rect(x*s.cs, y*s.cs, (x+1)*s.cs, (y+1)*s.cs)
	return s.img.SubImage(rect).(*ebiten.Image)
}

func (s *Sheet) Colorize(col color.Color) *Sheet {
	img := ebiten.NewImage(s.img.Bounds().Dx(), s.img.Bounds().Dy())
	opts := &ebiten.DrawImageOptions{}
	opts.ColorScale.ScaleWithColor(col)
	img.DrawImage(s.img, opts)

	return &Sheet{
		img: img,
		cs:  s.cs,
	}
}

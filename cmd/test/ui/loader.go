package ui

import (
	"bytes"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func loadFrontFace(path string, size float64) (*text.Face, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	src, err := text.NewGoTextFaceSource(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	var face text.Face
	face = &text.GoTextFace{
		Source: src,
		Size:   size,
	}

	return &face, nil
}

func loadImage(path string) (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile(path)
	return img, err
}

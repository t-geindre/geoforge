package main

import (
	"bytes"
	"image/color"
	"os"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Theme struct {
	// Containers images
	MainImage     *image.NineSlice // Window background
	SidebarImage  *image.NineSlice // Status bar, menu bar, toolbar color
	SidebarHeight int              // Height of status bar, menu bar, toolbar
	SurfaceImage  *image.NineSlice // Panels, cards, sheets color

	// Widgets background colors
	BgIdleColor     color.Color
	BgHoverColor    color.Color
	BgPressedColor  color.Color
	BgDisabledColor color.Color

	// Widgets foreground colors
	FgIdle     color.Color
	FgHover    color.Color
	FgPressed  color.Color
	FgDisabled color.Color

	// Images
	BtnImage *widget.ButtonImage

	WidgetsFont  *text.Face
	WidgetsInset *widget.Insets
	SurfaceInset *widget.Insets
}

func NewDefaultTheme() *Theme {
	return &Theme{
		MainImage:     solidImage(color.RGBA{25, 25, 35, 255}),
		SidebarImage:  solidImage(color.RGBA{30, 30, 40, 255}),
		SidebarHeight: 28,
		SurfaceImage:  solidImage(color.RGBA{40, 40, 50, 255}),
		SurfaceInset:  &widget.Insets{Left: 10, Right: 10, Top: 6, Bottom: 6},

		BgIdleColor:     color.RGBA{30, 30, 40, 255},
		BgHoverColor:    color.RGBA{50, 50, 60, 255},
		BgPressedColor:  color.RGBA{20, 20, 30, 255},
		BgDisabledColor: color.RGBA{10, 10, 10, 255},

		FgIdle:     color.RGBA{220, 220, 220, 255},
		FgHover:    color.RGBA{240, 240, 240, 255},
		FgPressed:  color.RGBA{200, 200, 200, 255},
		FgDisabled: color.RGBA{100, 100, 100, 255},

		BtnImage: &widget.ButtonImage{
			Idle:     solidImage(color.RGBA{60, 60, 70, 255}),
			Hover:    solidImage(color.RGBA{90, 90, 110, 255}),
			Pressed:  solidImage(color.RGBA{40, 40, 50, 255}),
			Disabled: solidImage(color.RGBA{20, 20, 20, 255}),
		},
		WidgetsFont:  mustLoadTextFace("assets/fonts/Roboto-Regular.ttf", 16),
		WidgetsInset: &widget.Insets{Left: 8, Right: 8, Top: 4, Bottom: 4},
	}
}

func solidImage(col color.Color) *image.NineSlice {
	img := ebiten.NewImage(1, 1)
	img.Fill(col)
	return image.NewNineSlice(img, [3]int{0, 1, 0}, [3]int{0, 1, 0})
}

func mustLoadTextFace(path string, size float64) *text.Face {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	src, err := text.NewGoTextFaceSource(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	var face text.Face
	face = &text.GoTextFace{
		Source: src,
		Size:   size,
	}

	return &face
}

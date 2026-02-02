package ui

import (
	"bytes"
	"image/color"
	"os"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/themes"
	"github.com/ebitenui/ebitenui/widget"
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

	BgIdleImage     *image.NineSlice
	BgHoverImage    *image.NineSlice
	BgPressedImage  *image.NineSlice
	BgDisabledImage *image.NineSlice

	// Widgets foreground colors
	FgIdleColor     color.Color
	FgHoverColor    color.Color
	FgPressedColor  color.Color
	FgDisabledColor color.Color

	FgIdleImage     *image.NineSlice
	FgHoverImage    *image.NineSlice
	FgPressedImage  *image.NineSlice
	FgDisabledImage *image.NineSlice

	// Images
	BtnImage *widget.ButtonImage

	WidgetsFont                  *text.Face
	WidgetsInset                 *widget.Insets
	SurfaceInset                 *widget.Insets
	SurfaceSpaceX, SurfaceSpaceY int
}

func NewDefaultTheme() *Theme {
	return &Theme{
		MainImage:     image.NewNineSliceColor(color.RGBA{25, 25, 35, 255}),
		SidebarImage:  image.NewNineSliceColor(color.RGBA{30, 30, 40, 255}),
		SidebarHeight: 28,
		SurfaceImage: image.NewBorderedNineSliceColor(
			color.RGBA{50, 50, 60, 255},
			color.RGBA{80, 80, 90, 255},
			1,
		),
		SurfaceInset: &widget.Insets{Left: 4, Right: 4, Top: 4, Bottom: 4},

		BgIdleColor:     color.RGBA{30, 30, 40, 255},
		BgHoverColor:    color.RGBA{50, 50, 60, 255},
		BgPressedColor:  color.RGBA{20, 20, 30, 255},
		BgDisabledColor: color.RGBA{10, 10, 10, 255},

		FgIdleColor:     color.RGBA{220, 220, 220, 255},
		FgHoverColor:    color.RGBA{240, 240, 240, 255},
		FgPressedColor:  color.RGBA{200, 200, 200, 255},
		FgDisabledColor: color.RGBA{100, 100, 100, 255},

		BtnImage: &widget.ButtonImage{
			Idle:     image.NewNineSliceColor(color.RGBA{60, 60, 70, 255}),
			Hover:    image.NewNineSliceColor(color.RGBA{90, 90, 110, 255}),
			Pressed:  image.NewNineSliceColor(color.RGBA{40, 40, 50, 255}),
			Disabled: image.NewNineSliceColor(color.RGBA{20, 20, 20, 255}),
		},
		WidgetsFont:   mustLoadTextFace("assets/fonts/Roboto-Regular.ttf", 14),
		WidgetsInset:  &widget.Insets{Left: 4, Right: 4, Top: 4, Bottom: 4},
		SurfaceSpaceX: 4,
		SurfaceSpaceY: 4,

		FgIdleImage:     image.NewNineSliceColor(color.RGBA{220, 220, 220, 255}),
		FgHoverImage:    image.NewNineSliceColor(color.RGBA{240, 240, 240, 255}),
		FgPressedImage:  image.NewNineSliceColor(color.RGBA{200, 200, 200, 255}),
		FgDisabledImage: image.NewNineSliceColor(color.RGBA{100, 100, 100, 255}),

		BgIdleImage:     image.NewNineSliceColor(color.RGBA{30, 30, 40, 255}),
		BgHoverImage:    image.NewNineSliceColor(color.RGBA{50, 50, 60, 255}),
		BgPressedImage:  image.NewNineSliceColor(color.RGBA{20, 20, 30, 255}),
		BgDisabledImage: image.NewNineSliceColor(color.RGBA{10, 10, 10, 255}),
	}
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

func GetDefaultTheme() *widget.Theme {
	return themes.GetBasicLightTheme()
}

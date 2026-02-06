package ui

import (
	"bytes"
	img "image"
	"image/color"
	"os"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/utilities/constantutil"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

type Theme struct {
	*widget.Theme
	PanelTheme       *PanelTheme
	ConnectionsTheme *ConnectionsTheme
}

func NewTheme() *Theme {
	const borderSize = 1
	face := mustLoadTextFace("assets/fonts/Roboto-Regular.ttf", 14)

	return &Theme{
		PanelTheme: &PanelTheme{
			ForegroundImage: image.NewBorderedNineSliceColor(
				color.RGBA{R: 0x2b, G: 0x2d, B: 0x30, A: 0xff},
				color.RGBA{R: 0x21, G: 0x21, B: 0x21, A: 0xff},
				1),
			BackgroundImage: image.NewNineSliceColor(
				color.RGBA{R: 0x21, G: 0x21, B: 0x21, A: 0xff},
			),
			Padding: &widget.Insets{Left: 10, Right: 10, Top: 10, Bottom: 10},
			Spacing: 10,
		},
		Theme: &widget.Theme{
			DefaultFace:      face,
			DefaultTextColor: color.White,
			ButtonTheme: &widget.ButtonParams{
				TextColor: &widget.ButtonTextColor{
					Idle:    color.White,
					Hover:   color.White,
					Pressed: color.White,
				},
				TextFace: face,
				Image: &widget.ButtonImage{
					Idle:    image.NewBorderedNineSliceColor(color.NRGBA{51, 51, 51, 255}, color.NRGBA{81, 81, 81, 255}, borderSize),
					Hover:   image.NewBorderedNineSliceColor(color.NRGBA{77, 77, 77, 255}, color.NRGBA{51, 51, 51, 255}, borderSize),
					Pressed: image.NewBorderedNineSliceColor(color.NRGBA{119, 119, 119, 255}, color.NRGBA{77, 77, 77, 255}, borderSize),
				},
				TextPadding: &widget.Insets{Left: 15, Right: 15, Top: 5, Bottom: 5},
				TextPosition: &widget.TextPositioning{
					VTextPosition: widget.TextPositionCenter,
					HTextPosition: widget.TextPositionCenter,
				},
			},
			PanelTheme: &widget.PanelParams{
				BackgroundImage: image.NewNineSliceColor(colornames.Black),
			},
			LabelTheme: &widget.LabelParams{
				Face: face,
				Color: &widget.LabelColor{
					Idle:     color.White,
					Disabled: color.NRGBA{122, 122, 122, 255},
				},
			},
			TextTheme: &widget.TextParams{
				Face:  face,
				Color: color.White,
				Position: &widget.TextPositioning{
					VTextPosition: widget.TextPositionCenter,
				},
			},
			TabbookTheme: &widget.TabBookParams{
				TabButton: &widget.ButtonParams{
					TextColor: &widget.ButtonTextColor{
						Idle:    color.White,
						Hover:   color.White,
						Pressed: color.White,
					},
					TextFace: face,
					Image: &widget.ButtonImage{
						Idle:    image.NewNineSliceColor(color.NRGBA{51, 51, 51, 255}),
						Hover:   image.NewNineSliceColor(color.NRGBA{77, 77, 77, 255}),
						Pressed: image.NewNineSliceColor(color.NRGBA{119, 119, 119, 255}),
					},
					TextPadding: widget.NewInsetsSimple(5),
					MinSize:     &img.Point{98, 40},
				},
				TabSpacing: constantutil.ConstantToPointer(1),
			},
			TabTheme: &widget.TabParams{
				BackgroundImage: image.NewNineSliceColor(color.NRGBA{32, 32, 32, 255}),
			},
			TextInputTheme: &widget.TextInputParams{
				Face: face,
				Image: &widget.TextInputImage{
					Idle:     image.NewBorderedNineSliceColor(color.NRGBA{77, 77, 77, 255}, color.NRGBA{177, 177, 177, 255}, borderSize),
					Disabled: image.NewBorderedNineSliceColor(color.NRGBA{47, 47, 47, 255}, color.NRGBA{177, 177, 177, 255}, borderSize),
				},
				Color: &widget.TextInputColor{
					Idle:          color.White,
					Caret:         color.White,
					Disabled:      color.NRGBA{127, 122, 126, 255},
					DisabledCaret: color.NRGBA{127, 122, 126, 255},
				},
				Padding: widget.NewInsetsSimple(5),
			},
			SliderTheme: &widget.SliderParams{
				TrackPadding:    widget.NewInsetsSimple(0),
				FixedHandleSize: constantutil.ConstantToPointer(6),
				TrackOffset:     constantutil.ConstantToPointer(0),

				PageSizeFunc: func() int {
					return 1
				},
				TrackImage: &widget.SliderTrackImage{
					Idle:     image.NewBorderedNineSliceColor(color.NRGBA{77, 77, 77, 255}, color.NRGBA{177, 177, 177, 255}, borderSize),
					Disabled: image.NewBorderedNineSliceColor(color.NRGBA{47, 47, 47, 255}, color.NRGBA{177, 177, 177, 255}, borderSize),
				},
				HandleImage: &widget.ButtonImage{
					Idle:         image.NewBorderedNineSliceColor(color.White, color.NRGBA{177, 177, 177, 255}, 1),
					Hover:        image.NewBorderedNineSliceColor(color.NRGBA{235, 235, 235, 255}, color.NRGBA{177, 177, 177, 255}, borderSize),
					Pressed:      image.NewBorderedNineSliceColor(color.NRGBA{210, 210, 210, 255}, color.NRGBA{177, 177, 177, 255}, borderSize),
					PressedHover: image.NewBorderedNineSliceColor(color.NRGBA{210, 210, 210, 255}, color.NRGBA{177, 177, 177, 255}, borderSize),
				},
			},
			ListComboButtonTheme: &widget.ListComboButtonParams{
				List: &widget.ListParams{
					EntryFace:                   face,
					EntryTextPadding:            widget.NewInsetsSimple(5),
					EntryTextHorizontalPosition: constantutil.ConstantToPointer(widget.TextPositionStart),
					EntryTextVerticalPosition:   constantutil.ConstantToPointer(widget.TextPositionCenter),
					EntryColor: &widget.ListEntryColor{
						Unselected:         color.White,
						Selected:           color.White,
						DisabledUnselected: color.NRGBA{127, 122, 126, 255},
						DisabledSelected:   color.NRGBA{127, 122, 126, 255},

						SelectedBackground:        color.NRGBA{40, 40, 40, 255},
						SelectedFocusedBackground: color.NRGBA{50, 50, 50, 255},

						SelectingBackground:        color.NRGBA{99, 99, 99, 255},
						FocusedBackground:          color.NRGBA{99, 99, 99, 255},
						SelectingFocusedBackground: color.NRGBA{99, 99, 99, 255},
						DisabledSelectedBackground: color.NRGBA{99, 99, 99, 255},
					},
					ScrollContainerPadding: widget.NewInsetsSimple(4),
					ScrollContainerImage: &widget.ScrollContainerImage{
						Idle:     image.NewNineSliceColor(color.NRGBA{77, 77, 77, 255}),
						Disabled: image.NewNineSliceColor(color.NRGBA{47, 47, 47, 255}),
						Mask:     image.NewNineSliceColor(color.NRGBA{77, 77, 77, 255}),
					},
					Slider: &widget.SliderParams{
						TrackImage: &widget.SliderTrackImage{
							Idle:     image.NewNineSliceColor(color.NRGBA{77, 77, 77, 255}),
							Disabled: image.NewNineSliceColor(color.NRGBA{47, 47, 47, 255}),
						},
						HandleImage: &widget.ButtonImage{
							Idle:    image.NewNineSliceColor(color.NRGBA{110, 110, 110, 255}),
							Hover:   image.NewNineSliceColor(color.NRGBA{120, 120, 120, 255}),
							Pressed: image.NewBorderedNineSliceColor(color.NRGBA{120, 120, 120, 255}, color.NRGBA{110, 110, 110, 255}, borderSize),
						},
						TrackPadding: &widget.Insets{Top: 4, Left: 4, Right: 4, Bottom: 4},
					},
				},
				Button: &widget.ButtonParams{
					TextColor: &widget.ButtonTextColor{
						Idle:    color.White,
						Hover:   color.White,
						Pressed: color.White,
					},
					TextFace: face,
					Image: &widget.ButtonImage{
						Idle:    getComboListButtonImage(color.NRGBA{51, 51, 51, 255}, color.NRGBA{81, 81, 81, 255}, color.NRGBA{220, 220, 220, 255}),
						Hover:   getComboListButtonImage(color.NRGBA{77, 77, 77, 255}, color.NRGBA{51, 51, 51, 255}, color.NRGBA{220, 220, 220, 255}),
						Pressed: getComboListButtonImage(color.NRGBA{119, 119, 119, 255}, color.NRGBA{77, 77, 77, 255}, color.NRGBA{220, 220, 220, 255}),
					},
					TextPadding: &widget.Insets{Left: 15, Right: 25, Top: 5, Bottom: 5},
					TextPosition: &widget.TextPositioning{
						VTextPosition: widget.TextPositionCenter,
						HTextPosition: widget.TextPositionCenter,
					},
				},
				MaxContentHeight: constantutil.ConstantToPointer(200),
			},
			CheckboxTheme: &widget.CheckboxParams{
				Label: &widget.LabelParams{
					Face: face,
					Color: &widget.LabelColor{
						Idle:     color.White,
						Disabled: color.NRGBA{122, 122, 122, 255},
					},
				},
				Image: getCheckboxImage(),
			},
		},
		ConnectionsTheme: &ConnectionsTheme{
			CableColor:       color.RGBA{R: 0x2F, G: 0x7C, B: 0xF6, A: 100},
			CableActiveColor: color.RGBA{R: 0x4F, G: 0xD1, B: 0xFF, A: 200},
			CableWidth:       4,
			KnobColor:        color.RGBA{R: 0xFF, G: 0x9F, B: 0x1C, A: 220},
			KnobActiveColor:  color.RGBA{R: 0xFF, G: 0xFF, B: 0x8C, A: 255},
		},
	}
}

type PanelTheme struct {
	BackgroundImage *image.NineSlice
	ForegroundImage *image.NineSlice
	Padding         *widget.Insets
	Spacing         int
}

type ConnectionsTheme struct {
	CableColor       color.Color
	CableActiveColor color.Color
	CableWidth       float32
	KnobColor        color.Color
	KnobActiveColor  color.Color
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

func getCheckboxImage() *widget.CheckboxImage {
	const size = 16
	const borderSize = 1
	const crossPadding = 4
	const crossWidth = 2

	border := image.NewBorderedNineSliceColor(color.NRGBA{119, 119, 119, 255}, color.White, borderSize)
	idle := ebiten.NewImage(size, size)
	border.Draw(idle, size, size, nil)
	idle9s := image.NewFixedNineSlice(idle)

	// Create checked image
	checked := ebiten.NewImage(size, size)
	border.Draw(checked, size, size, nil)
	vector.StrokeLine(checked, crossPadding, crossPadding, size-crossPadding, size-crossPadding, crossWidth, color.White, true)
	vector.StrokeLine(checked, size-crossPadding, crossPadding, crossPadding, size-crossPadding, crossWidth, color.White, true)

	checked9s := image.NewFixedNineSlice(checked)

	// Create greyed image
	greyed := ebiten.NewImage(size, size)
	border.Draw(greyed, size, size, nil)
	vector.StrokeLine(greyed, 5, 16, 27, 16, 3, color.White, true)
	greyed9s := image.NewFixedNineSlice(greyed)

	return &widget.CheckboxImage{
		Unchecked:         idle9s,
		Checked:           checked9s,
		Greyed:            greyed9s,
		UncheckedHovered:  idle9s,
		CheckedHovered:    checked9s,
		GreyedHovered:     greyed9s,
		UncheckedDisabled: idle9s,
		CheckedDisabled:   checked9s,
		GreyedDisabled:    greyed9s,
	}
}

func getComboListButtonImage(bgCol, borderCol, arrowCol color.Color) *image.NineSlice {
	const (
		border = 1
		w, h   = 44, 24

		leftCap   = 4
		rightCap  = 30
		topCap    = 4
		bottomCap = 4
	)

	i := ebiten.NewImage(w, h)
	i.Fill(bgCol)

	vector.StrokeRect(i, 0, 0, float32(w-border), float32(h-border), float32(border), borderCol, false)

	cx := float32(w - rightCap/2)
	cy := float32(h / 2)

	size := float32(4)

	var p vector.Path
	p.MoveTo(cx-size, cy-size)
	p.LineTo(cx, cy+size)
	p.LineTo(cx+size, cy-size)
	p.Close()

	dpOpt := vector.DrawPathOptions{AntiAlias: true}
	dpOpt.ColorScale.ScaleWithColor(arrowCol)

	flOpt := vector.FillOptions{
		FillRule: vector.FillRuleEvenOdd,
	}

	vector.FillPath(i, &p, &flOpt, &dpOpt)

	ws := [3]int{
		leftCap,
		w - leftCap - rightCap,
		rightCap,
	}
	hs := [3]int{
		topCap,
		h - topCap - bottomCap,
		bottomCap,
	}

	return image.NewNineSlice(i, ws, hs)
}

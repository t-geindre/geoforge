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
	PanelTheme *PanelTheme
}

func NewTheme() *Theme {
	const borderSize = 1
	face := mustLoadTextFace("assets/fonts/Roboto-Regular.ttf", 14)
	//face := mustLoadTextFace("assets/fonts/Gobold Light.otf", 14)
	//face := mustLoadTextFace("assets/fonts/AlegreyaSansSC-Regular.ttf", 16)

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
				TextPadding: &widget.Insets{Left: 30, Right: 30, Top: 5, Bottom: 5},
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
			TextAreaTheme: &widget.TextAreaParams{
				Face:                   face,
				StripBBCode:            constantutil.ConstantToPointer(true),
				ControlWidgetSpacing:   constantutil.ConstantToPointer(2),
				TextPadding:            &widget.Insets{Right: 18},
				ForegroundColor:        color.White,
				ScrollContainerPadding: widget.NewInsetsSimple(4),
				ScrollContainerImage: &widget.ScrollContainerImage{
					Idle:     image.NewBorderedNineSliceColor(color.NRGBA{77, 77, 77, 255}, color.NRGBA{177, 177, 177, 255}, 1),
					Disabled: image.NewBorderedNineSliceColor(color.NRGBA{47, 47, 47, 255}, color.NRGBA{177, 177, 177, 255}, 1),
					Mask:     image.NewBorderedNineSliceColor(color.NRGBA{77, 77, 77, 255}, color.NRGBA{177, 177, 177, 255}, 1),
				},
				Slider: &widget.SliderParams{
					TrackImage: &widget.SliderTrackImage{
						Idle:     image.NewBorderedNineSliceColor(color.NRGBA{77, 77, 77, 255}, color.NRGBA{177, 177, 177, 255}, 1),
						Disabled: image.NewBorderedNineSliceColor(color.NRGBA{47, 47, 47, 255}, color.NRGBA{177, 177, 177, 255}, 1),
					},
					HandleImage: &widget.ButtonImage{
						Idle:    image.NewBorderedNineSliceColor(color.NRGBA{77, 77, 77, 255}, color.NRGBA{51, 51, 51, 255}, 2),
						Hover:   image.NewBorderedNineSliceColor(color.NRGBA{99, 99, 99, 255}, color.NRGBA{77, 77, 77, 255}, 2),
						Pressed: image.NewBorderedNineSliceColor(color.NRGBA{99, 99, 99, 255}, color.NRGBA{77, 77, 77, 255}, 2),
					},
				},
			},
			ProgressBarTheme: &widget.ProgressBarParams{
				TrackPadding: widget.NewInsetsSimple(2),
				TrackImage: &widget.ProgressBarImage{
					Idle:     image.NewBorderedNineSliceColor(color.NRGBA{77, 77, 77, 255}, color.NRGBA{177, 177, 177, 255}, 1),
					Disabled: image.NewBorderedNineSliceColor(color.NRGBA{47, 47, 47, 255}, color.NRGBA{177, 177, 177, 255}, 1),
					Hover:    image.NewBorderedNineSliceColor(color.NRGBA{77, 77, 77, 255}, color.NRGBA{177, 177, 177, 255}, 1),
				},
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
			ListTheme: &widget.ListParams{
				EntryFace:                   face,
				EntryTextPadding:            widget.NewInsetsSimple(5),
				EntryTextHorizontalPosition: constantutil.ConstantToPointer(widget.TextPositionStart),
				EntryTextVerticalPosition:   constantutil.ConstantToPointer(widget.TextPositionCenter),
				MinSize:                     &img.Point{150, 0},
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
					Idle:     image.NewBorderedNineSliceColor(color.NRGBA{77, 77, 77, 255}, color.NRGBA{177, 177, 177, 255}, 1),
					Disabled: image.NewBorderedNineSliceColor(color.NRGBA{47, 47, 47, 255}, color.NRGBA{177, 177, 177, 255}, 1),
					Mask:     image.NewBorderedNineSliceColor(color.NRGBA{77, 77, 77, 255}, color.NRGBA{177, 177, 177, 255}, 1),
				},
				Slider: &widget.SliderParams{
					TrackImage: &widget.SliderTrackImage{
						Idle:     image.NewBorderedNineSliceColor(color.NRGBA{77, 77, 77, 255}, color.NRGBA{177, 177, 177, 255}, 1),
						Disabled: image.NewBorderedNineSliceColor(color.NRGBA{47, 47, 47, 255}, color.NRGBA{177, 177, 177, 255}, 1),
					},
					HandleImage: &widget.ButtonImage{
						Idle:    image.NewBorderedNineSliceColor(color.NRGBA{77, 77, 77, 255}, color.NRGBA{51, 51, 51, 255}, 2),
						Hover:   image.NewBorderedNineSliceColor(color.NRGBA{99, 99, 99, 255}, color.NRGBA{77, 77, 77, 255}, 2),
						Pressed: image.NewBorderedNineSliceColor(color.NRGBA{99, 99, 99, 255}, color.NRGBA{77, 77, 77, 255}, 2),
					},
				},
			},
			ListComboButtonTheme: &widget.ListComboButtonParams{
				List: &widget.ListParams{
					EntryFace:                   face,
					EntryTextPadding:            widget.NewInsetsSimple(5),
					EntryTextHorizontalPosition: constantutil.ConstantToPointer(widget.TextPositionStart),
					EntryTextVerticalPosition:   constantutil.ConstantToPointer(widget.TextPositionCenter),
					MinSize:                     &img.Point{200, 0},
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
						Idle:     image.NewBorderedNineSliceColor(color.NRGBA{77, 77, 77, 255}, color.NRGBA{177, 177, 177, 255}, borderSize),
						Disabled: image.NewBorderedNineSliceColor(color.NRGBA{47, 47, 47, 255}, color.NRGBA{177, 177, 177, 255}, borderSize),
						Mask:     image.NewBorderedNineSliceColor(color.NRGBA{77, 77, 77, 255}, color.NRGBA{177, 177, 177, 255}, borderSize),
					},
					Slider: &widget.SliderParams{
						TrackImage: &widget.SliderTrackImage{
							Idle:     image.NewBorderedNineSliceColor(color.NRGBA{77, 77, 77, 255}, color.NRGBA{177, 177, 177, 255}, borderSize),
							Disabled: image.NewBorderedNineSliceColor(color.NRGBA{47, 47, 47, 255}, color.NRGBA{177, 177, 177, 255}, borderSize),
						},
						HandleImage: &widget.ButtonImage{
							Idle:    image.NewBorderedNineSliceColor(color.NRGBA{77, 77, 77, 255}, color.NRGBA{51, 51, 51, 255}, borderSize),
							Hover:   image.NewBorderedNineSliceColor(color.NRGBA{99, 99, 99, 255}, color.NRGBA{77, 77, 77, 255}, borderSize),
							Pressed: image.NewBorderedNineSliceColor(color.NRGBA{99, 99, 99, 255}, color.NRGBA{77, 77, 77, 255}, borderSize),
						},
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
						Idle:    image.NewBorderedNineSliceColor(color.NRGBA{51, 51, 51, 255}, color.NRGBA{81, 81, 81, 255}, borderSize),
						Hover:   image.NewBorderedNineSliceColor(color.NRGBA{77, 77, 77, 255}, color.NRGBA{51, 51, 51, 255}, borderSize),
						Pressed: image.NewBorderedNineSliceColor(color.NRGBA{119, 119, 119, 255}, color.NRGBA{77, 77, 77, 255}, borderSize),
					},
					TextPadding: &widget.Insets{
						Left:   30,
						Right:  30,
						Top:    5,
						Bottom: 5,
					},
					TextPosition: &widget.TextPositioning{
						VTextPosition: widget.TextPositionCenter,
						HTextPosition: widget.TextPositionCenter,
					},
					MinSize: &img.Point{200, 0},
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
				Image: getDarkCheckbox(),
			},
		},
	}
}

type PanelTheme struct {
	BackgroundImage *image.NineSlice
	ForegroundImage *image.NineSlice
	Padding         *widget.Insets
	Spacing         int
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

func getDarkCheckbox() *widget.CheckboxImage {
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

package theme

import (
	img "image"
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/utilities/constantutil"
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/colornames"
)

func NewDefaultTheme() (*Theme, error) {
	const borderSize = 1

	face, err := loadFrontFace("assets/fonts/Roboto-Regular.ttf", 14)
	if err != nil {
		return nil, err
	}

	menuFace, err := loadFrontFace("assets/fonts/Roboto-Regular.ttf", 18) // TODO: should cache fonts
	if err != nil {
		return nil, err
	}

	sheet, err := loadImage("assets/icons/icons.png")
	if err != nil {
		return nil, err
	}

	blue := color.RGBA{R: 0x2F, G: 0x7C, B: 0xF6, A: 100}
	orange := color.RGBA{R: 0xFF, G: 0x9F, B: 0x1C, A: 220}

	icons := NewSheet(sheet, 24)
	iconsOrange := icons.Colorize(orange)
	iconsBlue := icons.Colorize(blue)
	iconsBlueSmall := iconsBlue.Scale(.6)

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
			CableColor:       blue,
			CableActiveColor: color.RGBA{R: 0x4F, G: 0xD1, B: 0xFF, A: 200},
			CableWidth:       4,
			KnobColor:        orange,
			KnobActiveColor:  color.RGBA{R: 0xFF, G: 0xFF, B: 0x8C, A: 255},
		},
		IconsTheme: &IconsTheme{
			Add:        &widget.GraphicImage{Idle: icons.Get(0, 0), Hover: iconsBlue.Get(0, 0)},
			Camera:     &widget.GraphicImage{Idle: iconsOrange.Get(1, 0)},
			Rendering:  &widget.GraphicImage{Idle: iconsBlueSmall.Get(1, 0)},
			Center:     &widget.GraphicImage{Idle: icons.Get(2, 0), Hover: iconsBlue.Get(2, 0)},
			Delete:     &widget.GraphicImage{Idle: icons.Get(3, 0), Hover: iconsBlue.Get(3, 0)},
			File:       &widget.GraphicImage{Idle: iconsOrange.Get(4, 0)},
			Noise:      &widget.GraphicImage{Idle: iconsOrange.Get(6, 0)},
			NoiseSmall: &widget.GraphicImage{Idle: iconsBlueSmall.Get(6, 0)},
			Open:       &widget.GraphicImage{Idle: icons.Get(7, 0), Hover: iconsBlue.Get(7, 0)},
			Save:       &widget.GraphicImage{Idle: icons.Get(0, 1), Hover: iconsBlue.Get(0, 1)},
			Zoom:       &widget.GraphicImage{Idle: icons.Get(1, 1), Hover: iconsBlue.Get(1, 1)},
		},
		MainMenuTheme: &MainMenuTheme{
			ButtonImage: &widget.ButtonImage{
				Idle:  image.NewNineSliceColor(color.NRGBA{51, 51, 51, 255}),
				Hover: image.NewNineSliceColor(color.NRGBA{77, 77, 77, 255}),
			},
			ButtonPadding: &widget.Insets{Left: 10, Right: 10, Top: 5, Bottom: 5},
			IconSpacing:   5,
			Font:          menuFace,
			TextColor:     colornames.White,
		},
	}, nil
}

package main

import (
	"geoforge/cmd/test/ui"
	"geoforge/game"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// THEME AND LAYOUT
	theme := ui.NewTheme()
	layout := ui.NewLayout(theme)

	// MAIN MENU
	menu := ui.NewMenu(theme)
	layout.AddChild(menu)

	// SPLIT PAN
	split := ui.NewSplit(theme)
	layout.AddChild(split)

	// DESKTOP (LEFT)
	desktop := ui.NewDesktop(1000, 1000)
	split.AddChild(desktop)

	for j := 0; j < 3; j++ {
		drag := ui.NewDraggable(j*200, 50, theme)
		drag.AddChild(getGridForm(theme))
		drag.SetTitle("Draggable " + string(rune('A'+j)))
		desktop.AddDraggable(drag)
	}

	// INSPECTOR (RIGHT)
	split.AddChild(widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(colornames.White)),
	))

	// STATUS
	status := ui.NewStatus(theme)
	layout.AddChild(status)

	// UI
	ui := &ebitenui.UI{Container: layout.Container}
	ui.PrimaryTheme = theme.Theme

	if err := ebiten.RunGame(game.NewGame(ui)); err != nil {
		panic(err)
	}
}

func getGridForm(theme *ui.Theme) *widget.Container {
	grid := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(2),
				widget.GridLayoutOpts.Stretch([]bool{false, true}, []bool{false}),
				widget.GridLayoutOpts.Padding(theme.PanelTheme.Padding),
				widget.GridLayoutOpts.Spacing(theme.PanelTheme.Spacing, theme.PanelTheme.Spacing),
			),
		),
	)

	items := []any{
		"Fast noise",
		"Maths",
		"Sine",
		"Pow",
	}
	grid.AddChild(
		widget.NewText(widget.TextOpts.TextLabel("List")),
		widget.NewListComboButton(
			widget.ListComboButtonOpts.Entries(items),
			widget.ListComboButtonOpts.EntryLabelFunc(func(v any) string {
				return v.(string)
			}, func(v any) string {
				return v.(string)
			}),
		),
	)

	grid.AddChild(
		widget.NewText(widget.TextOpts.TextLabel("Button")),
		widget.NewButton(widget.ButtonOpts.TextLabel("Button")),
	)
	grid.AddChild(
		widget.NewText(widget.TextOpts.TextLabel("Checkbox")),
		widget.NewCheckbox(),
	)
	grid.AddChild(
		widget.NewText(widget.TextOpts.TextLabel("Slider")),
		widget.NewSlider(
			widget.SliderOpts.MinMax(0, 100),
			widget.SliderOpts.InitialCurrent(50),
		),
	)
	grid.AddChild(
		widget.NewText(
			widget.TextOpts.TextLabel("Text Input"),
		),
		widget.NewTextInput(),
	)

	return grid
}

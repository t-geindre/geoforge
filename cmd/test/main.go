package main

import (
	"geoforge/cmd/test/ui"
	"geoforge/game"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	theme := ui.NewTheme()

	menu := ui.NewMenu(theme)
	desktop := ui.NewDesktop(1000, 1000)
	status := ui.NewStatus(theme)

	layout := ui.NewLayout(theme)
	layout.AddChild(menu, desktop, status)

	for j := 0; j < 3; j++ {
		drag := ui.NewDraggable(j*150, 50, theme)
		dbody := widget.NewContainer(
			widget.ContainerOpts.Layout(
				widget.NewGridLayout(
					widget.GridLayoutOpts.Columns(2),
					widget.GridLayoutOpts.Stretch([]bool{false, true}, []bool{false}),
					widget.GridLayoutOpts.Padding(theme.PanelTheme.Padding),
					widget.GridLayoutOpts.Spacing(theme.PanelTheme.Spacing, theme.PanelTheme.Spacing),
				),
			),
		)
		drag.AddChild(dbody)
		drag.SetTitle("Draggable " + string(rune('A'+j)))

		items := []any{
			"Fast noise liter",
			"Maths",
			"Sine",
			"Pow",
		}
		dbody.AddChild(
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

		dbody.AddChild(
			widget.NewText(widget.TextOpts.TextLabel("Button")),
			widget.NewButton(widget.ButtonOpts.TextLabel("Button")),
		)
		dbody.AddChild(
			widget.NewText(widget.TextOpts.TextLabel("Checkbox")),
			widget.NewCheckbox(),
		)
		dbody.AddChild(
			widget.NewText(widget.TextOpts.TextLabel("Slider")),
			widget.NewSlider(
				widget.SliderOpts.MinMax(0, 100),
				widget.SliderOpts.InitialCurrent(50),
			),
		)
		dbody.AddChild(
			widget.NewText(
				widget.TextOpts.TextLabel("Text Input"),
			),
			widget.NewTextInput(),
		)

		desktop.AddDraggable(drag)
	}

	ui := &ebitenui.UI{Container: layout.Container}
	ui.PrimaryTheme = theme.Theme

	updater := game.NewUpdateFunc(desktop.UpdateDragging)

	if err := ebiten.RunGame(game.NewGame(updater, ui)); err != nil {
		panic(err)
	}
}

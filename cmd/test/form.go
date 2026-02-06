package main

import (
	"geoforge/cmd/test/ui"

	"github.com/ebitenui/ebitenui/widget"
)

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

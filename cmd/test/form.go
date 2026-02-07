package main

import (
	"geoforge/cmd/test/ui/theme"
	"geoforge/cmd/test/ui/widgets"

	"github.com/ebitenui/ebitenui/input"
	"github.com/ebitenui/ebitenui/widget"
)

func getGridForm(theme *theme.Theme, connections *widgets.Connections) *widget.Container {
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
		widget.NewButton(widget.ButtonOpts.TextLabel("Button"),
			widget.ButtonOpts.WidgetOpts(
				widget.WidgetOpts.CursorHovered(input.CURSOR_POINTER),
			),
		),
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

	grid.AddChild(
		widget.NewText(
			widget.TextOpts.TextLabel("Input"),
		),
		connections.NewConnector(widgets.ConDirectionInput),
	)

	grid.AddChild(
		widget.NewText(
			widget.TextOpts.TextLabel("Output"),
		),
		connections.NewConnector(widgets.ConDirectionOutput),
	)

	return grid
}

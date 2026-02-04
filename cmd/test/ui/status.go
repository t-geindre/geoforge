package ui

import "github.com/ebitenui/ebitenui/widget"

func NewStatus(theme *Theme) *widget.Container {
	c := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(theme.PanelTheme.ForegroundImage),
		widget.ContainerOpts.WidgetOpts(),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Padding(theme.PanelTheme.Padding),
				widget.GridLayoutOpts.Columns(2),
				widget.GridLayoutOpts.Stretch([]bool{true, false}, []bool{true}),
			),
		),
	)

	c.AddChild(widget.NewText(
		widget.TextOpts.TextLabel("Left"),
	))

	c.AddChild(widget.NewText(
		widget.TextOpts.TextLabel("Right"),
	))

	return c
}

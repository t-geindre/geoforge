package ui2

import "github.com/ebitenui/ebitenui/widget"

func NewStatus(theme *widget.Theme) *widget.Container {
	c := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(theme.TabTheme.BackgroundImage),
		widget.ContainerOpts.WidgetOpts(),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
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

package ui2

import "github.com/ebitenui/ebitenui/widget"

func NewMenu(theme *widget.Theme) *widget.Container {
	c := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(theme.TabTheme.BackgroundImage),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			),
		),
	)

	c.AddChild(widget.NewText(
		widget.TextOpts.TextLabel("Menu 1"),
	))

	c.AddChild(widget.NewText(
		widget.TextOpts.TextLabel("Menu 2"),
	))

	c.AddChild(widget.NewText(
		widget.TextOpts.TextLabel("Menu 3"),
	))

	return c
}

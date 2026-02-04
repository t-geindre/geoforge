package ui

import "github.com/ebitenui/ebitenui/widget"

func NewMenu(theme *Theme) *widget.Container {
	c := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(theme.PanelTheme.ForegroundImage),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Padding(theme.PanelTheme.Padding),
				widget.RowLayoutOpts.Spacing(theme.PanelTheme.Spacing),
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

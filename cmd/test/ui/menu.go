package ui

import (
	"github.com/ebitenui/ebitenui/widget"
)

func NewMenu(theme *Theme) *widget.Container {
	c := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(theme.PanelTheme.ForegroundImage),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(3),
				widget.GridLayoutOpts.Stretch([]bool{false, true, false}, []bool{true}),
				widget.GridLayoutOpts.Padding(theme.PanelTheme.Padding),
			),
		),
	)

	left := widget.NewContainer(widget.ContainerOpts.Layout(widget.NewRowLayout(
		widget.RowLayoutOpts.Spacing(theme.PanelTheme.Spacing),
	)))
	center := widget.NewContainer(widget.ContainerOpts.Layout(widget.NewAnchorLayout()))
	right := widget.NewContainer(widget.ContainerOpts.Layout(widget.NewAnchorLayout()))

	left.AddChild(widget.NewGraphic(
		widget.GraphicOpts.Image(theme.IconsTheme.Logo),
	))
	left.AddChild(widget.NewButton(
		widget.ButtonOpts.TextAndImage(
			"Open", nil, &widget.GraphicImage{Idle: theme.IconsTheme.Open}, nil,
		),
	))
	left.AddChild(widget.NewButton(
		widget.ButtonOpts.TextAndImage(
			"New", nil, &widget.GraphicImage{Idle: theme.IconsTheme.Delete}, nil,
		),
	))
	left.AddChild(widget.NewButton(
		widget.ButtonOpts.TextAndImage(
			"Save", nil, &widget.GraphicImage{Idle: theme.IconsTheme.Save}, nil,
		),
	))

	c.AddChild(left)
	c.AddChild(center)
	c.AddChild(right)

	return c
}

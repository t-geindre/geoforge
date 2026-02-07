package ui

import (
	"geoforge/cmd/test/ui/theme"

	"github.com/ebitenui/ebitenui/widget"
)

type Layout struct {
	*widget.Container
}

func NewLayout(theme *theme.Theme) *Layout {
	return &Layout{
		Container: widget.NewContainer(
			widget.ContainerOpts.BackgroundImage(theme.PanelTheme.BackgroundImage),
			widget.ContainerOpts.Layout(
				widget.NewGridLayout(
					widget.GridLayoutOpts.Columns(1),
					widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, true, false}),
				),
			),
		),
	}
}

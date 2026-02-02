package ui2

import "github.com/ebitenui/ebitenui/widget"

type Layout struct {
	*widget.Container
}

func NewLayout() *Layout {
	return &Layout{
		Container: widget.NewContainer(
			widget.ContainerOpts.Layout(
				widget.NewGridLayout(
					widget.GridLayoutOpts.Columns(1),
					widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, true, false}),
				),
			),
		),
	}
}

package layout

import "github.com/ebitenui/ebitenui/widget"

func NewRoot() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewStackedLayout()),
	)
}

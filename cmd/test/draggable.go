package main

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type Draggable struct {
	*widget.Container
	header  *widget.Container
	content *widget.Container
	wx, wy  int
}

func NewDraggable(wx, wy int, col color.Color) *Draggable {
	container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(col)),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(1),
				widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false}),
			),
		),
	)
	header := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(0, 20),
		),
	)
	container.AddChild(header)
	return &Draggable{
		wx:        wx,
		wy:        wy,
		Container: container,
		header:    header,
	}
}

func (d *Draggable) GetDragContainer() *widget.Container {
	return d.header
}

func (d *Draggable) Size() (int, int) {
	r := d.GetWidget().Rect
	w, h := r.Dx(), r.Dy()
	if w <= 0 || h <= 0 {
		return 100, 100 // Fallback size
	}
	return w, h
}

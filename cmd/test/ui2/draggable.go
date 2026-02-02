package ui2

import (
	"github.com/ebitenui/ebitenui/widget"
)

type Draggable struct {
	*widget.Container
	header  *widget.Container
	content *widget.Container
	title   *widget.Text
	wx, wy  int
	dirty   bool
	w, h    int
}

func NewDraggable(wx, wy int, theme *widget.Theme) *Draggable {
	container := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			Padding: &widget.Insets{Left: wx, Top: wy},
		})),
		widget.ContainerOpts.BackgroundImage(theme.TabTheme.BackgroundImage),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(1),
			),
		),
	)
	header := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(theme.TabTheme.BackgroundImage),
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(),
		),
	)
	title := widget.NewText(
		widget.TextOpts.TextLabel("Draggable"),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
				},
			),
		),
	)
	header.AddChild(title)
	container.AddChild(header)

	return &Draggable{
		wx:        wx,
		wy:        wy,
		Container: container,
		header:    header,
		title:     title,
		dirty:     true,
	}
}

func (d *Draggable) GetDragContainer() *widget.Container {
	return d.header
}

func (d *Draggable) SetTitle(title string) {
	d.title.Label = title
	d.MarkDirty()
}

func (d *Draggable) MarkDirty() {
	d.dirty = true
}

package widgets

import (
	"geoforge/cmd/test/ui/theme"

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

func NewDraggable(wx, wy int, theme *theme.Theme, icon *widget.GraphicImage) *Draggable {
	container := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			Padding: &widget.Insets{Left: wx, Top: wy},
		})),
		widget.ContainerOpts.BackgroundImage(theme.PanelTheme.ForegroundImage),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(1),
			),
		),
	)
	header := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(theme.PanelTheme.ForegroundImage),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Padding(theme.PanelTheme.Padding),
				widget.RowLayoutOpts.Spacing(theme.PanelTheme.Spacing),
			),
		),
	)
	titleIcon := widget.NewGraphic(
		widget.GraphicOpts.Image(icon.Idle),
	)
	title := widget.NewText(
		widget.TextOpts.TextLabel("Draggable"),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
				},
			),
		),
	)
	header.AddChild(titleIcon, title)
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

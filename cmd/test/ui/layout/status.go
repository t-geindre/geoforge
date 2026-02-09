package layout

import (
	"fmt"
	"geoforge/camera"
	"geoforge/cmd/test/ui/theme"
	"geoforge/render"
	"geoforge/world"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewStatus(th *theme.Theme, rdr *render.Renderer, wrld *world.World, cam camera.Camera) *widget.Container {
	c := newStatusContainer(th)

	// LEFT
	left := newStatusBox(th)
	c.AddChild(left)

	left.AddChild(newStatusIcon(th.IconsTheme.Small[theme.IconStats]))
	left.AddChild(newUpdateText(func() string {
		return fmt.Sprintf("TPS: %02.0f", ebiten.ActualTPS())
	}, nil))

	left.AddChild(newStatusIcon(th.IconsTheme.Small[theme.IconStats]))
	left.AddChild(newUpdateText(func() string {
		return fmt.Sprintf("FPS: %02.0f", ebiten.ActualFPS())
	}, nil))

	// RIGHT
	right := newStatusBox(th)
	c.AddChild(right)

	// Camera position
	right.AddChild(newStatusIcon(th.IconsTheme.Small[theme.IconCamera]))
	right.AddChild(newUpdateText(func() string {
		cx, cy := cam.Position()
		return fmt.Sprintf("%.0fx%.0f", cx, cy)
	}, c))

	// Camera zoom
	right.AddChild(newStatusIcon(th.IconsTheme.Small[theme.IconZoom]))
	right.AddChild(newUpdateText(func() string {
		cz := cam.Zoom() * 100
		return fmt.Sprintf("%.0f%%", cz)
	}, c))

	// Drawn chunks
	right.AddChild(newStatusIcon(th.IconsTheme.Small[theme.IconChunk]))
	right.AddChild(newUpdateText(func() string {
		return fmt.Sprintf("%d / %d", rdr.DrawnChunks(), len(wrld.Chunks()))
	}, c))

	return c
}

func newStatusContainer(theme *theme.Theme) *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(theme.PanelTheme.ForegroundImage),
		widget.ContainerOpts.WidgetOpts(),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Padding(theme.PanelTheme.Padding),
				widget.GridLayoutOpts.Columns(2),
				widget.GridLayoutOpts.Stretch([]bool{true, false}, []bool{true}),
			),
		),
	)
}

func newStatusBox(theme *theme.Theme) *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Spacing(theme.PanelTheme.Spacing),
		)),
	)
}

func newUpdateText(upd func() string, parent *widget.Container) *widget.Text {
	return widget.NewText(
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.OnUpdate(func(w widget.HasWidget) {
				wdg := w.(*widget.Text)
				str := upd()
				if str != wdg.Label {
					wdg.Label = str
					if parent != nil {
						parent.RequestRelayout()
					}
				}
			}),
		),
	)
}

func newStatusIcon(i *widget.GraphicImage) *widget.Graphic {
	return widget.NewGraphic(
		widget.GraphicOpts.Image(i.Idle),
	)
}

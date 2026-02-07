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

func NewStatus(theme *theme.Theme, rdr *render.Renderer, wrld *world.World, cam camera.Camera) *widget.Container {
	c := widget.NewContainer(
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

	var right *widget.Text
	right = widget.NewText(
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.OnUpdate(func(w widget.HasWidget) {
				cx, cy := cam.Position()
				cz := cam.Zoom() * 100
				str := fmt.Sprintf(
					"Cam: %.0f X %.0f - %.0f%%  |  Chunks: %d / %d",
					cx, cy, cz,
					rdr.DrawnChunks(), len(wrld.Chunks()),
				)
				if str != right.Label {
					right.Label = str
					c.RequestRelayout()
				}
			}),
		),
	)

	var left *widget.Text
	left = widget.NewText(
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.OnUpdate(func(w widget.HasWidget) {
				left.Label = fmt.Sprintf("TPS: %02.0f  |  FPS: %02.0f", ebiten.ActualTPS(), ebiten.ActualFPS())
			}),
		),
	)

	c.AddChild(left)
	c.AddChild(right)

	return c
}

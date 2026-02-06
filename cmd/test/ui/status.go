package ui

import (
	"fmt"
	"geoforge/camera"
	"geoforge/render"
	"geoforge/world"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewStatus(theme *Theme, rdr *render.Renderer, wrld *world.World, cam camera.Camera) *widget.Container {
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

	var fps *widget.Text
	fps = widget.NewText(
		widget.TextOpts.TextLabel("FPS: 00"),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.OnUpdate(func(w widget.HasWidget) {
				cx, cy := cam.Position()
				cz := cam.Zoom() * 100
				str := fmt.Sprintf(
					"Cam: %.0fx%.0f  %0.f%%    Chunks: %d / %d    TPS: %2.0f    FPS: %2.0f",
					cx, cy, cz,
					rdr.DrawnChunks(), len(wrld.Chunks()),
					ebiten.ActualTPS(), ebiten.ActualFPS())
				if str != fps.Label {
					fps.Label = str
					c.RequestRelayout()
				}
			}),
		))

	c.AddChild(widget.NewText(
		widget.TextOpts.TextLabel("Left"),
	))

	c.AddChild(fps)

	return c
}

package ui

import (
	"geoforge/camera"
	"geoforge/render"
	"image"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type Preview struct {
	cam      camera.Camera
	renderer *render.Renderer
	widget   *widget.Widget
}

func NewPreview(cam camera.Camera, rdr *render.Renderer) *Preview {
	return &Preview{
		cam:      cam,
		renderer: rdr,
		widget: widget.NewWidget(
			widget.WidgetOpts.CursorEnterHandler(func(args *widget.WidgetCursorEnterEventArgs) {
				cam.Lock(false)
			}),
			widget.WidgetOpts.CursorExitHandler(func(args *widget.WidgetCursorExitEventArgs) {
				cam.Lock(true)
			}),
		),
	}
}

func (p *Preview) Render(screen *ebiten.Image) {
	p.renderer.Draw(screen)
}

func (p *Preview) GetWidget() *widget.Widget {
	return p.widget
}

func (p *Preview) PreferredSize() (int, int) {
	return 0, 0
}

func (p *Preview) SetLocation(rect image.Rectangle) {
	p.cam.ScreenMoveTo(rect.Min.X, rect.Min.Y)
	p.cam.SetViewport(rect.Dx(), rect.Dy())
	p.widget.Rect = rect
}

func (p *Preview) Validate() {
}

func (p *Preview) Update(updObj *widget.UpdateObject) {
	p.widget.Update(updObj)
}

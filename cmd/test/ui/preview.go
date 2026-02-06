package ui

import (
	"geoforge/camera"
	"geoforge/render"
	"time"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type Preview struct {
	cam        camera.Camera
	renderer   *render.Renderer
	img        *ebiten.Image
	container  *widget.Container
	lastUpdate time.Time
}

func NewPreview(cam camera.Camera, rdr *render.Renderer) (*widget.Container, func(*ebiten.Image)) {
	var p *Preview

	cam.Lock(true)

	p = &Preview{
		cam:      cam,
		renderer: rdr,
		container: widget.NewContainer(
			widget.ContainerOpts.WidgetOpts(),
			widget.ContainerOpts.WidgetOpts(
				widget.WidgetOpts.OnUpdate(func(w widget.HasWidget) {
					p.Update()
				}),
				widget.WidgetOpts.CursorEnterHandler(func(args *widget.WidgetCursorEnterEventArgs) {
					cam.Lock(false)
				}),
				widget.WidgetOpts.CursorExitHandler(func(args *widget.WidgetCursorExitEventArgs) {
					cam.Lock(true)
				}),
			),
		),
	}

	p.Update()

	return p.container, p.Draw
}

func (p *Preview) Update() {
	//if time.Since(p.lastUpdate) < 100*time.Millisecond {
	//	return
	//}
	p.lastUpdate = time.Now()

	rect := p.container.GetWidget().Rect
	ww, wh := rect.Dx(), rect.Dy()
	wx, wy := rect.Min.X, rect.Min.Y

	if ww <= 0 || wh <= 0 {
		return
	}

	p.cam.ScreenMoveTo(wx, wy)
	p.cam.SetViewport(ww, wh)
}

func (p *Preview) Draw(screen *ebiten.Image) {
	p.renderer.Draw(screen)
}

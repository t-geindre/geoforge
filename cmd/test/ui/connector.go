package ui

import (
	theme2 "geoforge/cmd/test/ui/theme"
	"image"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type ConDirection int

const (
	ConDirectionInput ConDirection = iota
	ConDirectionOutput
)

type Connector struct {
	widget    *widget.Widget
	knob      *ebiten.Image
	knobHl    *ebiten.Image
	center    image.Point
	dir       ConDirection
	dest      *Connector
	highlight bool
	theme     *theme2.Theme
}

func NewConnector(theme *theme2.Theme, dir ConDirection) *Connector {
	var c *Connector
	c = &Connector{
		widget: widget.NewWidget(
			widget.WidgetOpts.CursorEnterHandler(func(args *widget.WidgetCursorEnterEventArgs) {
				c.highlight = true
			}),
			widget.WidgetOpts.CursorExitHandler(func(args *widget.WidgetCursorExitEventArgs) {
				c.highlight = false
			}),
		),
		dir:   dir,
		theme: theme,
	}
	return c
}

func (p *Connector) Render(screen *ebiten.Image) {
	if p.knob == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.widget.Rect.Min.X), float64(p.widget.Rect.Min.Y))

	if p.highlight {
		screen.DrawImage(p.knobHl, op)
	} else {
		screen.DrawImage(p.knob, op)
	}
}

func (p *Connector) GetWidget() *widget.Widget {
	return p.widget
}

func (p *Connector) PreferredSize() (int, int) {
	return 0, 0
}

func (p *Connector) SetLocation(rect image.Rectangle) {
	if p.widget.Rect.Eq(rect) {
		return
	}

	p.widget.Rect = rect

	w, h := rect.Dx(), rect.Dy()
	if w <= 0 || h <= 0 {
		p.knob = nil
		p.knobHl = nil
	}

	p.center = image.Point{X: rect.Min.X + rect.Dx()/2, Y: rect.Min.Y + rect.Dy()/2}

	p.knob = ebiten.NewImage(rect.Dx(), rect.Dy())
	col := p.theme.ConnectionsTheme.KnobColor
	vector.StrokeCircle(p.knob, float32(w)/2, float32(h)/2, float32(h)/2-2, 2, col, true)
	vector.FillCircle(p.knob, float32(w)/2, float32(h)/2, float32(h)/4-2, col, true)

	colHl := p.theme.ConnectionsTheme.KnobActiveColor
	p.knobHl = ebiten.NewImage(rect.Dx(), rect.Dy())
	vector.StrokeCircle(p.knobHl, float32(w)/2, float32(h)/2, float32(h)/2-2, 2, colHl, true)
	vector.FillCircle(p.knobHl, float32(w)/2, float32(h)/2, float32(h)/4-2, colHl, true)
}

func (p *Connector) Validate() {
}

func (p *Connector) Update(updObj *widget.UpdateObject) {
	p.widget.Update(updObj)
}

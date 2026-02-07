package widgets

import (
	theme2 "geoforge/cmd/test/ui/theme"
	"image"
	"image/color"
	"math"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type cable struct {
	from, to *Connector
}

type Connections struct {
	container    widget.HasWidget
	hover        *Connector
	dragging     *Connector
	handlerClear map[*Connector]func()
	connectors   []*Connector
	cables       []*cable
	theme        *theme2.Theme
}

func NewConnections(t *theme2.Theme, c widget.HasWidget) *Connections {
	return &Connections{
		container:    c,
		handlerClear: make(map[*Connector]func()),
		cables:       []*cable{},
		connectors:   []*Connector{},
		theme:        t,
	}
}

func (c *Connections) NewConnector(dir ConDirection) *Connector {
	conn := NewConnector(c.theme, dir)

	// Drag state
	dStartClear := conn.GetWidget().MouseButtonPressedEvent.AddHandler(func(e any) {
		if e.(*widget.WidgetMouseButtonPressedEventArgs).Button != ebiten.MouseButtonLeft {
			return
		}
		c.DragStarts(conn)
	})
	dEndClear := conn.GetWidget().MouseButtonReleasedEvent.AddHandler(func(e any) {
		if e.(*widget.WidgetMouseButtonReleasedEventArgs).Button != ebiten.MouseButtonLeft {
			return
		}
		c.DragEnds(conn)
	})

	c.handlerClear[conn] = func() {
		dStartClear()
		dEndClear()
	}

	c.connectors = append(c.connectors, conn)

	return conn
}

func (c *Connections) RemoveConnector(conn *Connector) {
	// Remove its cables
	for i := len(c.cables) - 1; i >= 0; i-- {
		cl := c.cables[i]
		if cl.from == conn || cl.to == conn {
			c.cables = append(c.cables[:i], c.cables[i+1:]...)
		}
	}

	// Remove from connectors list
	for i, con := range c.connectors {
		if con == conn {
			c.connectors = append(c.connectors[:i], c.connectors[i+1:]...)
			break
		}
	}

	// Clear event handlers
	if clr, ok := c.handlerClear[conn]; ok {
		clr()
		delete(c.handlerClear, conn)
	}

}

func (c *Connections) Update() {
	if c.dragging == nil {
		return
	}

	mouseX, mouseY := ebiten.CursorPosition()
	mousePt := image.Point{X: mouseX, Y: mouseY}

	c.hover = nil
	for _, conn := range c.connectors {
		if conn == c.dragging {
			continue
		}
		conn.highlight = false
		connRect := conn.GetWidget().Rect
		if mousePt.In(connRect) && c.isValidTarget(c.dragging, conn) {
			conn.highlight = true
			c.hover = conn
			break
		}
	}
}

func (c *Connections) DragStarts(conn *Connector) {
	if conn.dest == nil {
		c.dragging = conn
		return
	}

	c.dragging = conn.dest

	// Remove existing cable
	for i, cl := range c.cables {
		if cl.from == conn || cl.to == conn {
			c.cables = append(c.cables[:i], c.cables[i+1:]...)
			break
		}
	}

	// Clear destination links
	if conn.dest != nil {
		conn.dest.dest = nil
		conn.dest = nil
	}
}

func (c *Connections) DragEnds(conn *Connector) {
	if c.dragging == nil {
		return
	}

	if c.isValidTarget(c.dragging, c.hover) {
		c.cables = append(c.cables, &cable{
			from: c.dragging,
			to:   c.hover,
		})
		c.dragging.dest = c.hover
		c.hover.dest = c.dragging
		c.hover.highlight = false
	}

	c.dragging = nil
}

func (c *Connections) Draw(screen *ebiten.Image) {
	dst := screen.SubImage(c.container.GetWidget().Rect).(*ebiten.Image)
	for _, cl := range c.cables {
		c.DrawCable(dst, cl.from.center, cl.to.center, c.theme.ConnectionsTheme.CableColor)
	}

	if c.dragging != nil {
		if c.isValidTarget(c.dragging, c.hover) {
			c.DrawCable(dst, c.dragging.center, c.hover.center, c.theme.ConnectionsTheme.CableActiveColor)
			return
		}

		mouseX, mouseY := ebiten.CursorPosition()
		c.DrawCable(dst, c.dragging.center, image.Point{X: mouseX, Y: mouseY}, c.theme.ConnectionsTheme.CableActiveColor)
	}
}

func (c *Connections) DrawCable(screen *ebiten.Image, from, to image.Point, col color.Color) {
	x0, y0 := float32(from.X), float32(from.Y)
	x1, y1 := float32(to.X), float32(to.Y)

	dx, dy := x1-x0, y1-y0
	dist := float32(math.Sqrt(float64(dx*dx + dy*dy)))

	midx, midy := (x0+x1)*0.5, (y0+y1)*0.5

	sag := dist * 0.32 // Sag

	cx, cy := midx, midy+sag

	path := &vector.Path{}
	path.MoveTo(x0, y0)
	path.QuadTo(cx, cy, x1, y1)

	strokeOpt := &vector.StrokeOptions{
		Width:    c.theme.ConnectionsTheme.CableWidth,
		LineCap:  vector.LineCapRound,
		LineJoin: vector.LineJoinRound,
	}
	drawOpt := &vector.DrawPathOptions{
		AntiAlias: true,
	}
	drawOpt.ColorScale.ScaleWithColor(col)

	vector.StrokePath(screen, path, strokeOpt, drawOpt)
}

func (c *Connections) isValidTarget(from, to *Connector) bool {
	if from != nil && to != nil && from.dir != to.dir && to.dest == nil {
		return true
	}

	return false
}

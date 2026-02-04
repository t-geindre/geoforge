package ui

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

const WorldMargin = 200 // pixels

type Desktop struct {
	*widget.ScrollContainer
	content *widget.Container

	ww, wh int

	dragging     *Draggable
	pendingFront *Draggable

	draggables    []*Draggable // last = top
	clearHandlers map[*Draggable]func()
	grabX, grabY  int

	panning    bool
	panStartMX int
	panStartMY int
	panStartL  float64
	panStartT  float64
}

func NewDesktop(ww, wh int) *Desktop {
	d := &Desktop{
		clearHandlers: make(map[*Draggable]func()),
		ww:            ww,
		wh:            wh,
	}

	d.content = widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(ww, wh),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MouseButtonPressedHandler(d.PanStarts),
			widget.WidgetOpts.MouseButtonReleasedHandler(d.PanEnds),
		),
	)
	d.ScrollContainer = widget.NewScrollContainer(
		widget.ScrollContainerOpts.Content(d.content),
		widget.ScrollContainerOpts.Image(&widget.ScrollContainerImage{
			Idle: image.NewNineSliceColor(color.NRGBA{0, 0, 0, 0}),
			Mask: image.NewNineSliceColor(color.NRGBA{255, 255, 255, 255}),
		}),
	)
	return d
}

func (d *Desktop) AddDraggable(drag *Draggable) {
	pressClear := drag.GetDragContainer().GetWidget().MouseButtonPressedEvent.AddHandler(func(aargs any) {
		if args, ok := aargs.(*widget.WidgetMouseButtonPressedEventArgs); ok {
			d.DragStarts(args, drag)
		}
	})
	releaseClear := drag.GetDragContainer().GetWidget().MouseButtonReleasedEvent.AddHandler(func(aargs any) {
		if args, ok := aargs.(*widget.WidgetMouseButtonReleasedEventArgs); ok {
			d.DragEnds(args, drag)
		}
	})
	focusClear := drag.GetWidget().MouseButtonPressedEvent.AddHandler(func(_ any) {
		if d.TopMostAtCursor() != drag {
			return
		}

		if d.dragging == nil && !d.panning {
			d.pendingFront = drag
		}
	})

	d.clearHandlers[drag] = func() {
		pressClear()
		releaseClear()
		focusClear()
	}

	d.draggables = append(d.draggables, drag)
	d.content.AddChild(drag)
}

func (d *Desktop) RemoveDraggable(drag *Draggable) {
	for i, dr := range d.draggables {
		if dr == drag {
			d.draggables = append(d.draggables[:i], d.draggables[i+1:]...)
			break
		}
	}
	d.content.RemoveChild(drag)
	if c, ok := d.clearHandlers[drag]; ok {
		c()
		delete(d.clearHandlers, drag)
	}
}

func (d *Desktop) BringToFront() {
	if d.pendingFront == nil {
		// none
		return
	}
	drag := d.pendingFront
	d.pendingFront = nil

	if len(d.draggables) < 2 || drag == d.draggables[len(d.draggables)-1] {
		// already front
		return
	}

	intersect := false
	rect := drag.GetWidget().Rect
	for _, dr := range d.draggables {
		if dr == drag {
			continue
		}
		r := dr.GetWidget().Rect
		if rect.Min.X < r.Max.X && rect.Max.X > r.Min.X &&
			rect.Min.Y < r.Max.Y && rect.Max.Y > r.Min.Y {
			intersect = true
			break
		}
	}

	if !intersect {
		// no intersection, no need to reorder
		return
	}

	for i, dr := range d.draggables {
		if dr == drag {
			d.draggables = append(d.draggables[:i], d.draggables[i+1:]...)
			break
		}
	}
	d.draggables = append(d.draggables, drag)

	d.content.RemoveChild(drag)
	d.content.AddChild(drag)
}

func (d *Desktop) DragStarts(args *widget.WidgetMouseButtonPressedEventArgs, drag *Draggable) {
	if args.Button != ebiten.MouseButtonLeft || d.dragging != nil || d.TopMostAtCursor() != drag {
		return
	}

	d.dragging = drag
	d.grabX, d.grabY = args.OffsetX, args.OffsetY
}

func (d *Desktop) DragEnds(args *widget.WidgetMouseButtonReleasedEventArgs, drag *Draggable) {
	if args.Button != ebiten.MouseButtonLeft || d.dragging != drag {
		return
	}
	d.pendingFront = drag
	d.dragging = nil
	d.RefitWorld()
}

func (d *Desktop) PanStarts(args *widget.WidgetMouseButtonPressedEventArgs) {
	if args.Button != ebiten.MouseButtonLeft || d.panning || d.TopMostAtCursor() != nil || d.dragging != nil {
		return
	}

	d.panning = true
	mx, my := ebiten.CursorPosition()
	d.panStartMX, d.panStartMY = mx, my
	d.panStartL, d.panStartT = d.ScrollLeft, d.ScrollTop
}

func (d *Desktop) PanEnds(args *widget.WidgetMouseButtonReleasedEventArgs) {
	if args.Button != ebiten.MouseButtonLeft || !d.panning {
		return
	}
	d.panning = false
}

func (d *Desktop) contentOrigin() (ox, oy int) {
	cr := d.ScrollContainer.ContentRect()
	return cr.Min.X, cr.Min.Y
}

func (d *Desktop) UpdateDragging() {
	d.BringToFront()

	if d.panning {
		mx, my := ebiten.CursorPosition()
		dx := float64(mx - d.panStartMX)
		dy := float64(my - d.panStartMY)

		vr := d.GetWidget().Rect
		vw, vh := vr.Dx(), vr.Dy()

		maxLpx := float64(d.ww - vw)
		maxTpx := float64(d.wh - vh)
		if maxLpx < 0 {
			maxLpx = 0
		}
		if maxTpx < 0 {
			maxTpx = 0
		}

		startLpx := d.panStartL * maxLpx
		startTpx := d.panStartT * maxTpx

		newLpx := startLpx - dx
		newTpx := startTpx - dy

		if newLpx < 0 {
			newLpx = 0
		} else if newLpx > maxLpx {
			newLpx = maxLpx
		}
		if newTpx < 0 {
			newTpx = 0
		} else if newTpx > maxTpx {
			newTpx = maxTpx
		}

		if maxLpx > 0 {
			d.ScrollLeft = newLpx / maxLpx
		} else {
			d.ScrollLeft = 0
		}
		if maxTpx > 0 {
			d.ScrollTop = newTpx / maxTpx
		} else {
			d.ScrollTop = 0
		}
	}

	ox, oy := d.contentOrigin()

	if d.dragging != nil {
		mx, my := ebiten.CursorPosition()

		wx := (mx - ox) - d.grabX
		wy := (my - oy) - d.grabY

		if wx < 0 {
			wx = 0
		}
		if wy < 0 {
			wy = 0
		}

		d.dragging.wx, d.dragging.wy = wx, wy

		ld := d.dragging.GetWidget().LayoutData.(widget.AnchorLayoutData)
		if ld.Padding == nil {
			ld.Padding = &widget.Insets{}
		}
		ld.Padding.Left = wx
		ld.Padding.Top = wy
		d.dragging.GetWidget().LayoutData = ld

		d.content.RequestRelayout()
	}
}

func (d *Desktop) TopMostAtCursor() *Draggable {
	mx, my := ebiten.CursorPosition()

	for i := len(d.draggables) - 1; i >= 0; i-- {
		dr := d.draggables[i]
		r := dr.GetWidget().Rect
		if mx >= r.Min.X && mx < r.Max.X && my >= r.Min.Y && my < r.Max.Y {
			return dr
		}
	}
	return nil
}

func (d *Desktop) RefitWorld() {
	vr := d.GetWidget().Rect
	vw, vh := vr.Dx(), vr.Dy()

	needW, needH := vw, vh

	for _, dr := range d.draggables {
		w, h := dr.GetWidget().Rect.Dx(), dr.GetWidget().Rect.Dy()
		if w == 0 || h == 0 {
			pw, ph := dr.PreferredSize()
			w, h = pw, ph
		}

		x := dr.wx
		y := dr.wy
		r := x + w
		b := y + h

		if r > needW {
			needW = r
		}
		if b > needH {
			needH = b
		}
	}

	needW += WorldMargin
	needH += WorldMargin

	d.content.GetWidget().MinWidth = needW
	d.content.GetWidget().MinHeight = needH

	d.ww = needW
	d.wh = needH

	d.content.RequestRelayout()
}

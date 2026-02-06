package ui

import (
	img "image"

	"github.com/ebitenui/ebitenui/input"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewSplit(theme *Theme) *widget.Container {
	layout := NewSplitLayout()
	container := widget.NewContainer(
		widget.ContainerOpts.Layout(layout),
	)
	container.AddChild(NewSplitBar(theme, layout, container))

	return container
}

type SplitBar struct {
	*widget.Button
	dragging  bool
	grabDX    int
	layout    *SplitLayout
	container *widget.Container
}

func NewSplitBar(t *Theme, l *SplitLayout, c *widget.Container) *widget.Button {
	b := &SplitBar{layout: l, container: c}

	b.Button = widget.NewButton(
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    t.PanelTheme.ForegroundImage,
			Hover:   t.PanelTheme.BackgroundImage,
			Pressed: t.PanelTheme.ForegroundImage,
		}),
		widget.ButtonOpts.PressedHandler(b.DragStart),
		widget.ButtonOpts.ReleasedHandler(func(*widget.ButtonReleasedEventArgs) { b.dragging = false }),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.OnUpdate(b.Update),
			widget.WidgetOpts.MinSize(l.BarWidth, l.BarWidth),
			widget.WidgetOpts.CursorHovered(input.CURSOR_EWRESIZE),
			widget.WidgetOpts.CursorPressed(input.CURSOR_EWRESIZE),
		),
	)

	return b.Button
}

func (b *SplitBar) Update(_ widget.HasWidget) {
	if b.dragging {
		x, _ := ebiten.CursorPosition()
		b.layout.SplitX = x - b.grabDX
		b.container.RequestRelayout()
		input.SetCursorShape(input.CURSOR_EWRESIZE)
	}
}

func (b *SplitBar) DragStart(args *widget.ButtonPressedEventArgs) {
	b.dragging = true
	x, _ := ebiten.CursorPosition()
	b.grabDX = x - b.layout.SplitX
}

func (b *SplitBar) DragEnd(args *widget.WidgetMouseButtonReleasedEventArgs) {
	b.dragging = false
}

type SplitLayout struct {
	SplitX      int
	BarWidth    int
	MinLeft     int
	MinRight    int
	lastBoundsW int
}

func NewSplitLayout() *SplitLayout {
	return &SplitLayout{
		SplitX:   200,
		BarWidth: 10,
		MinLeft:  50,
		MinRight: 50,
	}
}

func (l *SplitLayout) PreferredSize(widgets []widget.PreferredSizeLocateableWidget) (int, int) {
	ln := len(widgets)

	w, h := widgets[0].PreferredSize()
	if ln == 1 {
		return w, h
	}

	for i := 1; i < 2; i++ {
		ww, wh := widgets[i].PreferredSize()
		w += ww
		if wh > h {
			h = wh
		}
	}

	return w, h
}

func (l *SplitLayout) Layout(children []widget.PreferredSizeLocateableWidget, bounds img.Rectangle) {
	if l.lastBoundsW == 0 {
		// First, no scaling
		l.lastBoundsW = bounds.Dx()
	} else if bounds.Dx() != l.lastBoundsW {
		// Scale the split position based on the change in width
		scale := float64(bounds.Dx()) / float64(l.lastBoundsW)
		l.SplitX = int(float64(l.SplitX) * scale)
		l.lastBoundsW = bounds.Dx()
	}
	w := bounds.Dx()
	h := bounds.Dy()

	s := l.SplitX
	if s < l.MinLeft {
		s = l.MinLeft
	}
	if s > w-l.MinRight-l.BarWidth {
		s = w - l.MinRight - l.BarWidth
	}
	l.SplitX = s

	bar := children[0]
	barRect := img.Rect(bounds.Min.X+s, bounds.Min.Y, bounds.Min.X+s+l.BarWidth, bounds.Min.Y+h)
	bar.(interface{ SetLocation(img.Rectangle) }).SetLocation(barRect)

	ln := len(children)

	if ln == 1 {
		return
	}

	left := children[1]
	leftRect := img.Rect(bounds.Min.X, bounds.Min.Y, bounds.Min.X+s, bounds.Min.Y+h)
	left.(interface{ SetLocation(img.Rectangle) }).SetLocation(leftRect)

	if ln == 2 {
		return
	}

	right := children[2]
	rightRect := img.Rect(bounds.Min.X+s+l.BarWidth, bounds.Min.Y, bounds.Max.X, bounds.Min.Y+h)
	right.(interface{ SetLocation(img.Rectangle) }).SetLocation(rightRect)
}

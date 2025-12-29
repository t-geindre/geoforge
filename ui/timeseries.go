package ui

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

type GraphStyle struct {
	Padding float32

	LineWidth float32
	LineColor color.Color

	MidlineWidth float32
	MidlineColor color.Color

	BorderWidth float32
	BorderColor color.Color

	BackgroundColor color.Color
	Background      bool

	Height int
}

func DefaultGraphStyle() GraphStyle {
	return GraphStyle{
		Padding: 0,

		LineWidth: 1,
		LineColor: colornames.Lime,

		MidlineWidth: 1,
		MidlineColor: colornames.Blue,

		BorderWidth: 1,
		BorderColor: colornames.Aliceblue,

		BackgroundColor: colornames.Black,
		Background:      true,

		Height: 100,
	}
}

type ScaleMode int

const (
	ScaleAuto ScaleMode = iota
	ScaleFixed
)

type TimeSeries struct {
	buf         []float32
	head        int
	count       int
	frame       int
	sampleEvery int

	scaleMode ScaleMode
	fixedMin  float32
	fixedMax  float32
	clampMin  float32
	clampMax  float32

	valFunc func() float32
	lastVal float32

	style GraphStyle
	label string
}

func NewTimeSeries(label string, size int, sampleEvery int, style GraphStyle, val func() float32) *TimeSeries {
	if size < 2 {
		size = 2
	}
	if sampleEvery < 1 {
		sampleEvery = 1
	}
	if val == nil {
		val = func() float32 { return 0 }
	}
	return &TimeSeries{
		buf:         make([]float32, size),
		sampleEvery: sampleEvery,
		scaleMode:   ScaleAuto,
		clampMin:    float32(-1e9),
		clampMax:    float32(+1e9),
		style:       style,
		valFunc:     val,
		label:       label,
	}
}

func (ts *TimeSeries) SetFixedScale(min, max float32) {
	ts.scaleMode = ScaleFixed
	ts.fixedMin, ts.fixedMax = min, max
}

func (ts *TimeSeries) SetAutoScale() {
	ts.scaleMode = ScaleAuto
}

func (ts *TimeSeries) SetClamp(min, max float32) {
	ts.clampMin, ts.clampMax = min, max
}

func (ts *TimeSeries) Push(v float32) {
	ts.lastVal = v

	if v < ts.clampMin {
		v = ts.clampMin
	}
	if v > ts.clampMax {
		v = ts.clampMax
	}

	ts.buf[ts.head] = v
	ts.head = (ts.head + 1) % len(ts.buf)
	if ts.count < len(ts.buf) {
		ts.count++
	}
}

func (ts *TimeSeries) Sample(sample func() float32) {
	ts.frame++
	if ts.frame%ts.sampleEvery != 0 {
		return
	}
	ts.Push(sample())
}

func (ts *TimeSeries) Len() int { return ts.count }

func (ts *TimeSeries) At(i int) float32 {
	start := (ts.head - ts.count + len(ts.buf)) % len(ts.buf)
	idx := (start + i) % len(ts.buf)
	return ts.buf[idx]
}

func (ts *TimeSeries) MinMax() (min, max float32) {
	if ts.count == 0 {
		return 0, 1
	}
	min, max = ts.At(0), ts.At(0)
	for i := 1; i < ts.count; i++ {
		v := ts.At(i)
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func (ts *TimeSeries) Draw(dst *ebiten.Image, bounds image.Rectangle) {
	// background
	if ts.style.Background {
		vector.FillRect(dst,
			float32(bounds.Min.X), float32(bounds.Min.Y),
			float32(bounds.Dx()), float32(bounds.Dy()),
			ts.style.BackgroundColor, false,
		)
	}

	// border
	if ts.style.BorderWidth > 0 {
		vector.StrokeRect(dst,
			float32(bounds.Min.X), float32(bounds.Min.Y),
			float32(bounds.Dx()), float32(bounds.Dy()),
			ts.style.BorderWidth, ts.style.BorderColor, false,
		)
	}

	if ts.Len() < 2 {
		return
	}

	pad := ts.style.Padding
	x0 := float32(bounds.Min.X) + pad
	y0 := float32(bounds.Min.Y) + pad
	w := float32(bounds.Dx()) - 2*pad
	h := float32(bounds.Dy()) - 2*pad
	if w <= 1 || h <= 1 {
		return
	}

	// scale
	var yMin, yMax float32
	switch ts.scaleMode {
	case ScaleFixed:
		yMin, yMax = ts.fixedMin, ts.fixedMax
	default:
		yMin, yMax = ts.MinMax()
		// avoid flatline division
		if yMax-yMin < 1e-6 {
			yMax = yMin + 1
		}
		// add a tiny margin to breathe
		margin := float32(0.05) * (yMax - yMin)
		yMin -= margin
		yMax += margin
	}

	// optional midline
	if ts.style.MidlineWidth > 0 {
		mid := y0 + h*0.5
		vector.StrokeLine(dst, x0, mid, x0+w, mid, ts.style.MidlineWidth, ts.style.MidlineColor, false)
	}

	mapX := func(i int) float32 {
		t := float32(i) / float32(ts.Len()-1)
		return x0 + t*w
	}
	mapY := func(v float32) float32 {
		t := (v - yMin) / (yMax - yMin)
		// screen Y grows downward
		t = float32(math.Max(0, math.Min(1, float64(t))))
		return y0 + (1-t)*h
	}

	// build path
	var p vector.Path
	p.MoveTo(mapX(0), mapY(ts.At(0)))
	for i := 1; i < ts.Len(); i++ {
		p.LineTo(mapX(i), mapY(ts.At(i)))
	}

	// stroke
	op := &vector.StrokeOptions{}
	op.Width = ts.style.LineWidth

	draw := &vector.DrawPathOptions{}
	draw.ColorScale.ScaleWithColor(ts.style.LineColor)
	draw.AntiAlias = true

	vector.StrokePath(dst, &p, op, draw)
}

func (ts *TimeSeries) Style() GraphStyle {
	return ts.style
}

func (ts *TimeSeries) Update() {
	ts.Sample(ts.valFunc)
}

func (ts *TimeSeries) Label() string {
	return ts.label
}

func (ts *TimeSeries) Value() float32 {
	return ts.lastVal
}

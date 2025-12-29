package ui

import (
	"fmt"
	"image"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
)

type metrics struct {
	metrics []Metric
}

func NewMetrics(ms ...Metric) Component {
	return &metrics{
		metrics: ms,
	}
}

func (d *metrics) UiUpdate(ctx *debugui.Context) {
	ctx.Header(
		"Metrics",
		true,
		func() {
			ctx.Loop(len(d.metrics), func(i int) {
				m := d.metrics[i]
				if ts, ok := m.(*TimeSeries); ok {
					ts.Update()
					ctx.Text(fmt.Sprintf("%s: %.0f", m.Label(), m.Value()))

					ctx.SetGridLayout(nil, []int{ts.Style().Height})
					ctx.GridCell(func(bounds image.Rectangle) {
						ctx.DrawOnlyWidget(func(dst *ebiten.Image) {
							ts.Draw(dst, bounds)
						})
					})

					ctx.SetGridLayout(nil, nil)
					return
				}

				ctx.Text(fmt.Sprintf("%s: %.0f", m.Label(), m.Value()))
			})
		},
	)
}

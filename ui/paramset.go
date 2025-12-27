package ui

import (
	"geoforge/preset"

	"github.com/ebitengine/debugui"
)

func ParamSetUI(ctx *debugui.Context, ps preset.ParamSet) {
	ctx.Loop(len(ps), func(i int) {
		p := ps[i]
		ctx.Text(p.Label())

		switch tp := p.(type) {
		case preset.Variable[int]:
			v := tp.Val()
			ctx.Slider(&v, tp.Min(), tp.Max(), tp.Step()).On(func() {
				tp.SetVal(v)
			})

		case preset.Variable[float32]:
			v := float64(tp.Val())
			ctx.SliderF(&v, float64(tp.Min()), float64(tp.Max()), float64(tp.Step()), tp.Digits()).On(func() {
				tp.SetVal(float32(v))
			})
		}
	})
}

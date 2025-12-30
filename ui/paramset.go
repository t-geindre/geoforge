package ui

import (
	"geoforge/preset"

	"github.com/ebitengine/debugui"
)

type paramSet struct {
	ps    preset.ParamSet
	label string
}

func NewParamSet(label string, ps preset.ParamSet) Component {
	return &paramSet{
		ps:    ps,
		label: label,
	}
}

func (p *paramSet) UiUpdate(ctx *debugui.Context) {
	ctx.Header(p.label, false, func() {
		p.handle(ctx, p.ps.All())
	})
}

func (p *paramSet) handle(ctx *debugui.Context, pms []preset.ParamGeneric) {
	ctx.Loop(len(pms), func(i int) {
		pm := pms[i]

		ctx.SetGridLayout([]int{100, -1}, nil)

		switch tp := pm.(type) {
		case preset.Variable[int]:
			ctx.Text(pm.Label())
			v := tp.Val()
			ctx.Slider(&v, tp.Min(), tp.Max(), tp.Step()).On(func() {
				tp.SetVal(v)
			})

		case preset.Variable[float32]:
			ctx.Text(pm.Label())
			v := float64(tp.Val())
			ctx.SliderF(&v, float64(tp.Min()), float64(tp.Max()), float64(tp.Step()), tp.Digits()).On(func() {
				tp.SetVal(float32(v))
			})

		case preset.Param[string]:
			ctx.Text(pm.Label())
			v := tp.Val()
			ctx.TextField(&v).On(func() {
				tp.SetVal(v)
			})

		case preset.Param[bool]:
			ctx.Text(pm.Label())
			v := tp.Val()
			ctx.Checkbox(&v, "").On(func() {
				tp.SetVal(v)
			})

		case preset.ChoiceGeneric:
			ctx.Text(pm.Label())
			v := tp.ValIndex()
			ctx.Dropdown(&v, tp.OptionsLabels()).On(func() {
				tp.SetValByIndex(v)
			})

		case preset.ParamSet:
			ctx.TreeNode(tp.Label(), func() {
				p.handle(ctx, tp.All())
			})
		}
	})
}

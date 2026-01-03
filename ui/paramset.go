package ui

import (
	"fmt"
	"geoforge/preset"
	"image"
	"image/color"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
	ctx.Header(p.label, true, func() {
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
		case preset.Param[color.RGBA]:
			col := tp.Val()
			ur, ug, ub, _ := col.R, col.G, col.B, col.A
			r, g, b := int(ur), int(ug), int(ub)

			ctx.TreeNode(pm.Label(), func() {
				ctx.SetGridLayout([]int{-3, -1}, []int{54})
				ctx.GridCell(func(bounds image.Rectangle) {
					ctx.SetGridLayout([]int{-1, -3}, nil)
					ctx.Text("Red:")
					ctx.Slider(&r, 0, 255, 1).On(func() {
						col.R = uint8(r)
						tp.SetVal(col)
					})
					ctx.Text("Green:")
					ctx.Slider(&g, 0, 255, 1).On(func() {
						col.G = uint8(g)
						tp.SetVal(col)
					})
					ctx.Text("Blue:")
					ctx.Slider(&b, 0, 255, 1).On(func() {
						col.B = uint8(b)
						tp.SetVal(col)
					})
				})
				ctx.GridCell(func(bounds image.Rectangle) {
					ctx.DrawOnlyWidget(func(screen *ebiten.Image) {
						scale := ctx.Scale()
						vector.FillRect(
							screen,
							float32(bounds.Min.X*scale),
							float32(bounds.Min.Y*scale),
							float32(bounds.Dx()*scale),
							float32(bounds.Dy()*scale),
							col,
							false)

						txtCol := color.White
						if (float64(col.R)*0.299 + float64(col.G)*0.587 + float64(col.B)*0.114) > 186.0 {
							txtCol = color.Black
						}

						txt := fmt.Sprintf("#%02X%02X%02X", r, g, b)

						op := &text.DrawOptions{}
						op.GeoM.Translate(float64((bounds.Min.X+bounds.Max.X)/2), float64((bounds.Min.Y+bounds.Max.Y)/2))
						op.GeoM.Scale(float64(scale), float64(scale))
						op.PrimaryAlign = text.AlignCenter
						op.SecondaryAlign = text.AlignCenter
						op.ColorScale.ScaleWithColor(txtCol)
						debugui.DrawText(screen, txt, op)
					})
				})
			})

		case preset.ParamSet:
			ctx.TreeNode(tp.Label(), func() {
				p.handle(ctx, tp.All())
			})
		}
	})
}

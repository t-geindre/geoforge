package render

import (
	_ "embed"
	"geoforge/preset"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed colorscale.kage
var colorscaleShdRaw []byte

type ColorScale struct {
	sh *ebiten.Shader
	ps preset.ParamSet

	from    color.RGBA
	fromF32 [3]float32

	to    color.RGBA
	toF32 [3]float32

	trans float32
}

func NewGrayScale() *ColorScale {

	shd, err := ebiten.NewShader(colorscaleShdRaw)
	if err != nil {
		panic(err)
	}

	g := &ColorScale{
		sh: shd,
		ps: preset.NewAnonymousParamSet(),
	}

	g.from = color.RGBA{A: 255}
	g.ps.Append(preset.NewParam(0, "Low", g.from, func(p preset.Param[color.RGBA]) {
		g.from = p.Val()
		g.fromF32 = [3]float32{
			float32(g.from.R) / 255.0,
			float32(g.from.G) / 255.0,
			float32(g.from.B) / 255.0,
		}
	}))

	g.to = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	g.ps.Append(preset.NewParam(0, "High", g.to, func(p preset.Param[color.RGBA]) {
		g.to = p.Val()
		g.toF32 = [3]float32{
			float32(g.to.R) / 255.0,
			float32(g.to.G) / 255.0,
			float32(g.to.B) / 255.0,
		}
	}))

	g.ps.Append(preset.NewVariable(0, "Transition", float32(0.5), 0.0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		g.trans = p.Val()
	}))

	return g
}

func (g *ColorScale) DrawChunk(dst *ebiten.Image, w, h int, op *ebiten.DrawRectShaderOptions) {
	op.Uniforms["ColorFrom"] = g.fromF32
	op.Uniforms["ColorTo"] = g.toF32
	op.Uniforms["Transition"] = g.trans
	dst.DrawRectShader(w, h, g.sh, op)
}

func (g *ColorScale) Params() preset.ParamSet {
	return g.ps
}

func (g *ColorScale) Name() string {
	return "Color scale"
}

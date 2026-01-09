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

	edge   float32
	smooth float32
}

func NewColorScale() *ColorScale {

	shd, err := ebiten.NewShader(colorscaleShdRaw)
	if err != nil {
		panic(err)
	}

	g := &ColorScale{
		sh: shd,
		ps: preset.NewAnonymousParamSet(),
	}

	g.from = color.RGBA{A: 255}
	g.ps.Append(preset.NewParam(0, "Color low", g.from, func(p preset.Param[color.RGBA]) {
		g.from = p.Val()
		g.fromF32 = [3]float32{
			float32(g.from.R) / 255.0,
			float32(g.from.G) / 255.0,
			float32(g.from.B) / 255.0,
		}
	}))

	g.to = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	g.ps.Append(preset.NewParam(0, "Color high", g.to, func(p preset.Param[color.RGBA]) {
		g.to = p.Val()
		g.toF32 = [3]float32{
			float32(g.to.R) / 255.0,
			float32(g.to.G) / 255.0,
			float32(g.to.B) / 255.0,
		}
	}))

	g.ps.Append(preset.NewVariable(0, "Edge", float32(0.5), 0.0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		g.edge = p.Val()
	}))

	g.ps.Append(preset.NewVariable(0, "Smoothness", float32(0.1), 0.0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		g.smooth = p.Val()
	}))

	return g
}

func (g *ColorScale) DrawChunk(dst *ebiten.Image, w, h int, op *ebiten.DrawRectShaderOptions) {
	op.Uniforms["ColorFrom"] = g.fromF32
	op.Uniforms["ColorTo"] = g.toF32
	op.Uniforms["Edge"] = g.edge
	op.Uniforms["Smooth"] = g.smooth
	dst.DrawRectShader(w, h, g.sh, op)
}

func (g *ColorScale) Params() preset.ParamSet {
	return g.ps
}

func (g *ColorScale) Name() string {
	return "Color scale"
}

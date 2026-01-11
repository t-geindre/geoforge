package render

import (
	_ "embed"
	"geoforge/preset"
	"image/color"
)

//go:embed colorscale.kage
var colorScaleShdRaw []byte

func NewColorScale() ChunkRenderer {
	r := newChunkRenderer("Color scale", colorScaleShdRaw)

	r.ps.Append(preset.NewParam(0, "Color low", color.RGBA{A: 255}, func(p preset.Param[color.RGBA]) {
		r.uniforms["ColorFrom"] = rgbaToF32(p.Val())
	}))

	r.ps.Append(preset.NewParam(0, "Color high", color.RGBA{R: 255, G: 255, B: 255, A: 255}, func(p preset.Param[color.RGBA]) {
		r.uniforms["ColorTo"] = rgbaToF32(p.Val())
	}))

	r.ps.Append(preset.NewVariable(0, "Edge", float32(0.5), 0.0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		r.uniforms["Edge"] = p.Val()
	}))

	r.ps.Append(preset.NewVariable(0, "Smoothness", float32(0.1), 0.0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		r.uniforms["Smooth"] = p.Val()
	}))
	return r
}

package render

import (
	_ "embed"
	"geoforge/preset"
	"geoforge/world"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed terrain.kage
var terrainShdRaw []byte

func NewTerrain() ChunkRenderer {
	r := newChunkRenderer("Terrain", terrainShdRaw)

	// Update, light follows mouse
	r.update = func() {
		x, y := ebiten.CursorPosition()
		r.uniforms["LightPos"] = [2]float32{
			float32(x),
			float32(y),
		}
	}

	// Lighting
	lights := preset.NewParamSet(0, "Lighting")

	lights.Append(preset.NewVariable(1, "Ambient", 0.35, 0.0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		r.uniforms["Ambient"] = p.Val()
	}))

	lights.Append(preset.NewVariable(1, "Normal Strength", 70, 0.0, 250.0, 1, 0, func(p preset.Param[float32]) {
		r.uniforms["NormalStrength"] = p.Val()
	}))
	lights.Append(preset.NewVariable(1, "Normal Epsilon", 1, 1, world.ChunkApron, 1, 0, func(p preset.Param[float32]) {
		r.uniforms["NormalEps"] = p.Val()
	}))

	r.ps.Append(lights)

	// Levels
	levels := preset.NewParamSet(0, "Levels")

	levels.Append(preset.NewVariable(1, "Sea", 0.5, -0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		r.uniforms["SeaLevel"] = p.Val()
	}))

	levels.Append(preset.NewVariable(1, "Beach", 0.06, -0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		r.uniforms["BeachLevel"] = p.Val()
	}))

	levels.Append(preset.NewVariable(1, "Plain", 0.35, -0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		r.uniforms["PlainLevel"] = p.Val()
	}))

	levels.Append(preset.NewVariable(1, "Hill", 0.60, -0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		r.uniforms["HillLevel"] = p.Val()
	}))

	levels.Append(preset.NewVariable(1, "Mountain", 0.85, -0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		r.uniforms["MountainLevel"] = p.Val()
	}))

	r.ps.Append(levels)

	// Colors
	colors := preset.NewParamSet(0, "Colors")

	colors.Append(preset.NewParam(0, "Watter shallow", f32ToRgba([3]float32{0.1, 0.4, 0.7}), func(p preset.Param[color.RGBA]) {
		r.uniforms["WaterShallowColor"] = rgbaToF32(p.Val())
	}))

	colors.Append(preset.NewParam(0, "Watter deep", f32ToRgba([3]float32{0.02, 0.08, 0.25}), func(p preset.Param[color.RGBA]) {
		r.uniforms["WaterDeepColor"] = rgbaToF32(p.Val())
	}))

	colors.Append(preset.NewParam(0, "Beach", f32ToRgba([3]float32{0.85, 0.80, 0.60}), func(p preset.Param[color.RGBA]) {
		r.uniforms["BeachColor"] = rgbaToF32(p.Val())
	}))

	colors.Append(preset.NewParam(0, "Plain", f32ToRgba([3]float32{0.15, 0.55, 0.20}), func(p preset.Param[color.RGBA]) {
		r.uniforms["PlainColor"] = rgbaToF32(p.Val())
	}))

	colors.Append(preset.NewParam(0, "Hill", f32ToRgba([3]float32{0.35, 0.45, 0.25}), func(p preset.Param[color.RGBA]) {
		r.uniforms["HillColor"] = rgbaToF32(p.Val())
	}))

	colors.Append(preset.NewParam(0, "Mountain", f32ToRgba([3]float32{0.55, 0.55, 0.55}), func(p preset.Param[color.RGBA]) {
		r.uniforms["MountainColor"] = rgbaToF32(p.Val())
	}))

	colors.Append(preset.NewParam(0, "Snow", f32ToRgba([3]float32{0.95, 0.95, 0.95}), func(p preset.Param[color.RGBA]) {
		r.uniforms["SnowColor"] = rgbaToF32(p.Val())
	}))

	r.ps.Append(colors)

	return r
}

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

	// Spot position update
	msTrack := true
	mx, my := ebiten.Monitor().Size()
	fmx, fmy := float32(mx), float32(my)
	sx, sy := fmx/2, fmy/2

	r.update = func() {
		if !msTrack {
			r.uniforms["LightPos"] = [2]float32{sx, sy}
			return
		}

		x, y := ebiten.CursorPosition()
		r.uniforms["LightPos"] = [2]float32{
			float32(x),
			float32(y),
		}
	}
	spotX := preset.NewVariable(200, "X", sx, -200, fmx+200, 1, 0, func(p preset.Param[float32]) {
		sx = p.Val()
	})
	spotY := preset.NewVariable(200, "Y", sy, -200, fmy+200, 1, 0, func(p preset.Param[float32]) {
		sy = p.Val()
	})

	// Lighting
	lights := preset.NewParamSet(0, "Lighting")

	ambient := preset.NewParamSet(0, "Ambient")
	ambient.Append(preset.NewVariable(1, "Intensity", 0.35, 0.0, 1.0, 0.01, 2, func(p preset.Param[float32]) { // Todo add color
		r.uniforms["Ambient"] = p.Val()
	}))
	ambient.Append(preset.NewParam(0, "Color", f32ToRgba([3]float32{1.0, 1.0, 1.0}), func(p preset.Param[color.RGBA]) {
		r.uniforms["AmbientColor"] = rgbaToF32(p.Val())
	}))
	lights.Append(ambient)

	spot := preset.NewParamSet(0, "Spot")
	spot.Append(preset.NewVariable(1, "Intensity", 1, 0.0, 2.0, 0.01, 2, func(p preset.Param[float32]) { // Todo add radius and color
		r.uniforms["LightIntensity"] = p.Val()
	}))
	spot.Append(preset.NewParam(0, "Track mouse", true, func(p preset.Param[bool]) {
		msTrack = p.Val()
		if !msTrack {
			spot.Append(spotX)
			spot.Append(spotY)
			return
		}
		spot.Remove(spotX)
		spot.Remove(spotY)
	}))
	lights.Append(spot)

	normals := preset.NewParamSet(0, "Normals")
	normals.Append(preset.NewVariable(1, "Strength", 70, 0.0, 500.0, 1, 0, func(p preset.Param[float32]) {
		r.uniforms["NormalStrength"] = p.Val()
	}))
	normals.Append(preset.NewVariable(1, "Epsilon", 1, 1, world.ChunkApron, 1, 0, func(p preset.Param[float32]) {
		r.uniforms["NormalEps"] = p.Val()
	}))
	lights.Append(normals)

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

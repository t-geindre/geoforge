package noise

import (
	"geoforge/preset"
)

type Noise interface {
	Fill(dst []float32, size int, x0, y0 float32)
	Params() preset.ParamSet
}

type fillFunc func(dst []float32, size int, x0, y0 float32)

type noise struct {
	fn     fillFunc
	params preset.ParamSet
}

func newNoise(ps preset.ParamSet, fn fillFunc) Noise {
	return &noise{
		fn:     fn,
		params: ps,
	}
}

func (n *noise) Fill(dst []float32, size int, x0, y0 float32) {
	n.fn(dst, size, x0, y0)
}

func (n *noise) Params() preset.ParamSet {
	return n.params
}

/*
type Func func(x, y float32) float32

func NewWorldNoise() (preset.ParamSet, Func) {
	var ps preset.ParamSet

	var noise = fastnoise.New[float32]()
	noise.NoiseType(fastnoise.OpenSimplex2S)

	ps.Append(preset.NewVariable(ParamSeed, ParamSeedLabel, 42, 0, 1_000_000_000, 1, 0, func(v int) {
		noise.Seed = v
	}))

	clampMin := float32(-1)
	clampMax := float32(1)

	return ps, func(x, y float32) float32 {
		// Simple placeholder noise function
		land := noise.GetNoise2D(x, y)
		return clamp(land, clampMin, clampMax)
	}
}
*/

package noise

import (
	"geoforge/preset"
)

type Noise interface {
	Get(x, y float32) float32
	Params() preset.ParamSet
}

type noise struct {
	fn     func(x, y float32) float32
	params preset.ParamSet
}

func newNoise(ps preset.ParamSet, fn func(x, y float32) float32) Noise {
	return &noise{
		fn:     fn,
		params: ps,
	}
}

func (n *noise) Get(x, y float32) float32 {
	return n.fn(x, y)
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

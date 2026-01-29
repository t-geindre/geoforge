package noise

import (
	"geoforge/preset"

	fastnoise "github.com/Auburn/FastNoiseLite/Go"
)

type warp struct {
	*fastnoise.State[float32]
	ps      preset.ParamSet
	enabled bool
}

func newWarp() *warp {
	w := &warp{
		State: fastnoise.New[float32](),
		ps:    preset.NewParamSet(0, "Warp"),
	}

	//Seed
	w.ps.Append(preset.NewVariable(ParamSeed, "Seed", 1337, SeedMin, SeeMax, 1, 0, func(p preset.Param[int]) {
		w.State.Seed = p.Val()
	}))

	// Warp Type
	const DomainWarpNone fastnoise.DomainWarpType = -1
	w.ps.Append(preset.NewChoice(ParamWarpType, "Type", DomainWarpNone, []preset.Option[fastnoise.DomainWarpType]{
		preset.NewOption(DomainWarpNone, "None"),
		preset.NewOption(fastnoise.DomainWarpOpenSimplex2, "OpenSimplex2"),
		preset.NewOption(fastnoise.DomainWarpOpenSimplex2Reduced, "OpenSimplex2Reduced"),
		preset.NewOption(fastnoise.DomainWarpBasicGrid, "BasicGrid"),
	}, func(p preset.Param[fastnoise.DomainWarpType]) {
		if p.Val() == DomainWarpNone {
			w.enabled = false
			return
		}

		w.enabled = true
		w.State.DomainWarpType = p.Val()
	}))

	// Amp/Freq
	w.ps.Append(preset.NewVariable(ParamWarpAmp, "Amplitude", 0.0, 0.0, 1000.0, 1.0, 2, func(p preset.Param[float32]) {
		w.State.DomainWarpAmp = p.Val()
	}))

	w.ps.Append(preset.NewVariable(ParamWarpFreq, "Frequency", 0.0001, 0.0001, 0.1, 0.0001, 4, func(p preset.Param[float32]) {
		w.State.Frequency = p.Val()
	}))

	// Fractal
	w.ps.Append(preset.NewChoice(ParamFractType, "Fractal type", fastnoise.FractalNone, []preset.Option[fastnoise.FractalType]{
		preset.NewOption(fastnoise.FractalNone, "None"),
		preset.NewOption(fastnoise.FractalDomainWarpProgressive, "Warp progressive"),
		preset.NewOption(fastnoise.FractalDomainWarpIndependent, "Warp independent"),
	}, func(p preset.Param[fastnoise.FractalType]) {
		w.State.FractalType(p.Val())
	}))
	w.ps.Append(preset.NewVariable(ParamFractOctaves, "Octaves", 1, 1, 10, 1, 0, func(p preset.Param[int]) {
		w.State.Octaves = p.Val()
	}))
	w.ps.Append(preset.NewVariable(ParamFractLacunarity, "Lacunarity", 1, 1.0, 4.0, 0.1, 2, func(p preset.Param[float32]) {
		w.State.Lacunarity = p.Val()
	}))
	w.ps.Append(preset.NewVariable(ParamFractGain, "Gain", 0, 0.0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		w.State.Gain = p.Val()
	}))

	return w
}

func (w *warp) Warp(x, y float32) (float32, float32) {
	if !w.enabled {
		return x, y
	}

	return w.State.DomainWarp2D(x, y)
}

func (w *warp) Params() preset.ParamSet {
	return w.ps
}

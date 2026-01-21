package noise

import (
	"geoforge/preset"

	fastnoise "github.com/Auburn/FastNoiseLite/Go"
)

type warp struct {
	src  Noise
	warp *fastnoise.State[float32]
	ps   preset.ParamSet
}

func NewWarp() Noise {
	w := &warp{
		warp: fastnoise.New[float32](),
		ps:   preset.NewAnonymousParamSet(),
	}

	w.ps.Append(preset.NewChoice(ParamWarpType, "Type", w.warp.DomainWarpType, []preset.Option[fastnoise.DomainWarpType]{
		preset.NewOption(fastnoise.DomainWarpOpenSimplex2, "OpenSimplex2"),
		preset.NewOption(fastnoise.DomainWarpOpenSimplex2Reduced, "OpenSimplex2Reduced"),
		preset.NewOption(fastnoise.DomainWarpBasicGrid, "BasicGrid"),
	}, func(p preset.Param[fastnoise.DomainWarpType]) {
		w.warp.DomainWarpType = p.Val()
	}))

	w.ps.Append(preset.NewParam(0, "Source", w.src, func(p preset.Param[Noise]) {
		w.src = p.Val()
	}))

	w.ps.Append(preset.NewVariable(ParamWarpAmp, "Amplitude", 0.0, 0.0, 1000.0, 1.0, 2, func(p preset.Param[float32]) {
		w.warp.DomainWarpAmp = p.Val()
	}))

	w.ps.Append(preset.NewVariable(ParamWarpFreq, "Frequency", 0.0001, 0.0001, 0.1, 0.0001, 4, func(p preset.Param[float32]) {
		w.warp.Frequency = p.Val()
	}))

	return w
}

func (w *warp) At(x, y float32) float32 {
	if w.src == nil {
		return 0
	}

	x, y = w.warp.DomainWarp2D(x, y)
	return w.src.At(x, y)
}

func (w *warp) Params() preset.ParamSet {
	return w.ps
}

func (w *warp) Name() string {
	return "Warp"
}

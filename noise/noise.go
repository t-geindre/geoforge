package noise

import (
	"geoforge/preset"

	fastnoise "github.com/Auburn/FastNoiseLite/Go"
)

const DomainWarpNone fastnoise.DomainWarpType = -1

type Noise interface {
	Fill(dst []float32, size int, x0, y0 float32)
	Params() preset.ParamSet
}

type noise struct {
	fsn    *fastnoise.State[float32]
	fsw    *fastnoise.State[float32]
	doWarp bool
	ps     preset.ParamSet
}

func NewNoise() Noise {
	n := &noise{
		fsn: fastnoise.New[float32](),
		fsw: fastnoise.New[float32](),
	}

	n.buildParams()

	return n
}

func (n *noise) Fill(dst []float32, size int, x0, y0 float32) {
	idx := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			px := x0 + float32(x)
			py := y0 + float32(y)
			if n.doWarp {
				px, py = n.fsw.DomainWarp2D(px, py)
			}
			dst[idx] = n.fsn.GetNoise2D(px, py)
			idx++
		}
	}
}

func (n *noise) Params() preset.ParamSet {
	return n.ps
}

func (n *noise) buildParams() {
	n.ps = preset.NewAnonymousParamSet()

	// Basic parameters
	n.ps.Append(preset.NewVariable(1, "Scale", 0.0005, 0.0001, 0.01, 0.0001, 4, func(p preset.Param[float32]) {
		n.fsn.Frequency = p.Val()
	}))

	n.ps.Append(preset.NewChoice(0, "Type", 0, []preset.Option[fastnoise.NoiseType]{
		preset.NewOption(fastnoise.OpenSimplex2, "OpenSimplex2"),
		preset.NewOption(fastnoise.OpenSimplex2S, "OpenSimplex2S"),
		preset.NewOption(fastnoise.Cellular, "Cellular"),
		preset.NewOption(fastnoise.Perlin, "Perlin"),
		preset.NewOption(fastnoise.ValueCubic, "ValueCubic"),
		preset.NewOption(fastnoise.Value, "Value"),
	}, func(p preset.Param[fastnoise.NoiseType]) {
		n.fsn.NoiseType(p.Val())
	}))

	// Fractal parameters

	fract := preset.NewParamSet(0, "Fractal")
	fract.Append(preset.NewChoice(0, "Fractal type", 0, []preset.Option[fastnoise.FractalType]{
		preset.NewOption(fastnoise.FractalNone, "None"),
		preset.NewOption(fastnoise.FractalFBm, "FBm"),
		preset.NewOption(fastnoise.FractalRidged, "Ridged"),
		preset.NewOption(fastnoise.FractalPingPong, "PingPong"),
	}, func(p preset.Param[fastnoise.FractalType]) {
		n.fsn.FractalType(p.Val())
	}))

	fract.Append(preset.NewVariable(2, "Octaves", 1, 1, 10, 1, 0, func(p preset.Param[int]) {
		n.fsn.Octaves = p.Val()
	}))
	fract.Append(preset.NewVariable(2, "Lacunarity", 1, 1.0, 4.0, 0.1, 2, func(p preset.Param[float32]) {
		n.fsn.Lacunarity = p.Val()
	}))
	fract.Append(preset.NewVariable(2, "Gain", 0, 0.0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		n.fsn.Gain = p.Val()
	}))

	fract.Append(preset.NewVariable(2, "Strength", 0.0, 0.0, 2.0, 0.01, 2, func(p preset.Param[float32]) {
		n.fsn.WeightedStrength = p.Val()
	}))

	n.ps.Append(fract)

	// Domain warp parameters

	warp := preset.NewParamSet(0, "Domain Warp")
	warp.Append(preset.NewChoice(0, "Type", -1, []preset.Option[fastnoise.DomainWarpType]{
		preset.NewOption(DomainWarpNone, "None"),
		preset.NewOption(fastnoise.DomainWarpOpenSimplex2, "OpenSimplex2"),
		preset.NewOption(fastnoise.DomainWarpOpenSimplex2Reduced, "OpenSimplex2Reduced"),
		preset.NewOption(fastnoise.DomainWarpBasicGrid, "BasicGrid"),
	}, func(p preset.Param[fastnoise.DomainWarpType]) {
		v := p.Val()

		if v == DomainWarpNone {
			n.doWarp = false
			return
		}

		n.doWarp = true
		n.fsw.DomainWarpType = v
	}))

	warp.Append(preset.NewVariable(2, "Amplitude", 0.0, 0.0, 100.0, 1.0, 2, func(p preset.Param[float32]) {
		n.fsw.DomainWarpAmp = p.Val()
	}))

	warp.Append(preset.NewVariable(2, "Frequency", 0.0001, 0.0001, 0.1, 0.0001, 4, func(p preset.Param[float32]) {
		n.fsw.Frequency = p.Val()
	}))

	n.ps.Append(warp)
}

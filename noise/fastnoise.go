package noise

import (
	"geoforge/preset"

	fastnoise "github.com/Auburn/FastNoiseLite/Go"
)

type fastNoise struct {
	fsn    *fastnoise.State[float32]
	fsw    *fastnoise.State[float32]
	doWarp bool
	ps     preset.ParamSet
}

func NewFastNoise() Noise {
	n := &fastNoise{
		fsn: fastnoise.New[float32](),
		fsw: fastnoise.New[float32](),
	}

	n.buildParams()

	return n
}

func (n *fastNoise) At(x, y float32) float32 {
	px, py := x, y
	if n.doWarp {
		px, py = n.fsw.DomainWarp2D(px, py)
	}
	return n.fsn.GetNoise2D(px, py)
}

func (n *fastNoise) Params() preset.ParamSet {
	return n.ps
}

func (n *fastNoise) buildParams() {
	n.ps = preset.NewAnonymousParamSet()

	// Seed
	n.ps.Append(preset.NewVariable(ParamSeed, "Seed", 1337, SeedMin, SeeMax, 1, 0, func(p preset.Param[int]) {
		n.fsn.Seed = p.Val()
	}))

	// Cellular specific parameters
	cellPs := preset.NewParamSet(0, "Cellular")
	cellPs.Append(preset.NewChoice(0, "Dist func", n.fsn.CellularDistanceFunc, []preset.Option[fastnoise.CellularDistanceFunc]{
		preset.NewOption(fastnoise.CellularDistanceEuclidean, "Euclidean"),
		preset.NewOption(fastnoise.CellularDistanceEuclideanSq, "EuclideanSq"),
		preset.NewOption(fastnoise.CellularDistanceManhattan, "Manhattan"),
		preset.NewOption(fastnoise.CellularDistanceHybrid, "Hybrid"),
	}, func(p preset.Param[fastnoise.CellularDistanceFunc]) {
		n.fsn.CellularDistanceFunc = p.Val()
	}))
	cellPs.Append(preset.NewChoice(0, "Return type", n.fsn.CellularReturnType, []preset.Option[fastnoise.CellularReturnType]{
		preset.NewOption(fastnoise.CellularReturnCellValue, "CellValue"),
		preset.NewOption(fastnoise.CellularReturnDistance, "Distance"),
		preset.NewOption(fastnoise.CellularReturnDistance2, "Distance2"),
		preset.NewOption(fastnoise.CellularReturnDistance2Add, "Distance2Add"),
		preset.NewOption(fastnoise.CellularReturnDistance2Sub, "Distance2Sub"),
		preset.NewOption(fastnoise.CellularReturnDistance2Mul, "Distance2Mul"),
	}, func(p preset.Param[fastnoise.CellularReturnType]) {
		n.fsn.CellularReturnType = p.Val()
	}))
	cellPs.Append(preset.NewVariable(0, "Jitter", n.fsn.CellularJitterMod, 0.0, 2.0, 0.01, 2, func(p preset.Param[float32]) {
		n.fsn.CellularJitterMod = p.Val()
	}))

	// Basic parameters
	n.ps.Append(preset.NewChoice(ParamSubType, "Type", 0, []preset.Option[fastnoise.NoiseType]{
		preset.NewOption(fastnoise.OpenSimplex2, "OpenSimplex2"),
		preset.NewOption(fastnoise.OpenSimplex2S, "OpenSimplex2S"),
		preset.NewOption(fastnoise.Cellular, "Cellular"),
		preset.NewOption(fastnoise.Perlin, "Perlin"),
		preset.NewOption(fastnoise.ValueCubic, "ValueCubic"),
		preset.NewOption(fastnoise.Value, "Value"),
	}, func(p preset.Param[fastnoise.NoiseType]) {
		n.fsn.NoiseType(p.Val())
		if p.Val() == fastnoise.Cellular {
			n.ps.Add(ParamScale, cellPs)
		} else {
			n.ps.Remove(cellPs)
		}
	}))

	n.ps.Append(preset.NewVariable(ParamScale, "Scale", 0.0005, 0.0001, 0.02, 0.0001, 4, func(p preset.Param[float32]) {
		n.fsn.Frequency = p.Val()
	}))

	// Fractal parameters
	pingPongStrength := preset.NewVariable(ParamFractPPStrength, "PP Strength", 2.0, 0.0, 10.0, 0.1, 2, func(p preset.Param[float32]) {
		n.fsn.PingPongStrength = p.Val()
	})

	fract := preset.NewParamSet(ParamSetFract, "Fractal")
	fract.Append(preset.NewChoice(ParamFractType, "Fractal type", 0, []preset.Option[fastnoise.FractalType]{
		preset.NewOption(fastnoise.FractalNone, "None"),
		preset.NewOption(fastnoise.FractalFBm, "FBm"),
		preset.NewOption(fastnoise.FractalRidged, "Ridged"),
		preset.NewOption(fastnoise.FractalPingPong, "PingPong"),
	}, func(p preset.Param[fastnoise.FractalType]) {
		n.fsn.FractalType(p.Val())
		if p.Val() == fastnoise.FractalPingPong {
			fract.Append(pingPongStrength)
		} else {
			fract.Remove(pingPongStrength)
		}
	}))

	fract.Append(preset.NewVariable(ParamFractOctaves, "Octaves", 1, 1, 10, 1, 0, func(p preset.Param[int]) {
		n.fsn.Octaves = p.Val()
	}))
	fract.Append(preset.NewVariable(ParamFractLacunarity, "Lacunarity", 1, 1.0, 4.0, 0.1, 2, func(p preset.Param[float32]) {
		n.fsn.Lacunarity = p.Val()
	}))
	fract.Append(preset.NewVariable(ParamFractGain, "Gain", 0, 0.0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		n.fsn.Gain = p.Val()
	}))

	fract.Append(preset.NewVariable(ParamFractWeighting, "Weighting", 0.0, 0.0, 2.0, 0.01, 2, func(p preset.Param[float32]) {
		n.fsn.WeightedStrength = p.Val()
	}))

	n.ps.Append(fract)

	// Domain warp parameters
	warp := preset.NewParamSet(ParamSetWarp, "Domain Warp")
	warp.Append(preset.NewChoice(ParamWarpType, "Type", -1, []preset.Option[fastnoise.DomainWarpType]{
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

	warp.Append(preset.NewVariable(ParamWarpAmp, "Amplitude", 0.0, 0.0, 100.0, 1.0, 2, func(p preset.Param[float32]) {
		n.fsw.DomainWarpAmp = p.Val()
	}))

	warp.Append(preset.NewVariable(ParamWarpFreq, "Frequency", 0.0001, 0.0001, 0.1, 0.0001, 4, func(p preset.Param[float32]) {
		n.fsw.Frequency = p.Val()
	}))

	n.ps.Append(warp)
}

func (n *fastNoise) Name() string {
	return "Fast noise lite"
}

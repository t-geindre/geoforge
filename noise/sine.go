package noise

import (
	"geoforge/preset"
	"math"
)

func NewSine() Noise {
	var freq, gain float64
	var octaves int
	var lacunarity, persistence float64

	var warpFreq, warpAmount float64
	var warpOctaves int

	ps := preset.NewParamSet()

	ps = append(ps, preset.NewVariable(
		ParamFrequency, ParamFrequencyLabel,
		.0002, .0002, .1, .0001, 4,
		func(v float32) { freq = float64(v) },
	))

	ps = append(ps, preset.NewVariable(
		ParamGain, ParamGainLabel,
		1, -1, 1, .01, 4,
		func(v float32) { gain = float64(v) },
	))

	ps = append(ps, preset.NewVariable(
		ParamOctaves, "Octaves",
		1, 1, 12, 1, 0,
		func(v float32) { octaves = int(v) },
	))

	ps = append(ps, preset.NewVariable(
		ParamLacunarity, ParamLacunarityLabel,
		0, 1, 4, .01, 2,
		func(v float32) { lacunarity = float64(v) },
	))

	ps = append(ps, preset.NewVariable(
		ParamPersistence, ParamPersistenceLabel,
		0, 0, 1, .01, 2,
		func(v float32) { persistence = float64(v) },
	))

	ps = append(ps, preset.NewVariable(
		ParamWarpFreq, ParamWarpFreqLabel,
		0, .00001, .1, .00001, 5,
		func(v float32) { warpFreq = float64(v) },
	))

	ps = append(ps, preset.NewVariable(
		ParamWarpAmount, ParamWarpAmountLabel,
		0, 0, 1000, 1, 0,
		func(v float32) { warpAmount = float64(v) },
	))

	ps = append(ps, preset.NewVariable(
		ParamWarpOctaves, ParamWarpOctavesLabel,
		1, 1, 8, 1, 0,
		func(v float32) { warpOctaves = int(v) },
	))

	// Helpers inline (pas d'alloc, pas de struct)
	fbmSine := func(x, y, f0 float64, oct int, lac, pers float64) float64 {
		f := f0
		a := 1.0
		sum := 0.0
		norm := 0.0
		if oct < 1 {
			oct = 1
		}
		for k := 0; k < oct; k++ {
			sum += (math.Sin(x*f) * math.Sin(y*f)) * a
			norm += a
			f *= lac
			a *= pers
		}
		if norm != 0 {
			sum /= norm
		}
		return sum
	}

	return newNoise(ps, func(dst []float32, size int, x0, y0 float32) {
		xf0, yf0 := float64(x0), float64(y0)

		n := size * size
		if len(dst) < n {
			panic("dst too small")
		}

		o := octaves
		if o < 1 {
			o = 1
		}
		wo := warpOctaves
		if wo < 1 {
			wo = 1
		}

		for i := 0; i < n; i++ {
			x := xf0 + float64(i%size)
			y := yf0 + float64(i/size)

			// --- Domain warp: deux champs (wx, wy) décorrélés via offsets
			wx := fbmSine(x+12.34, y-56.78, warpFreq, wo, lacunarity, persistence)
			wy := fbmSine(x-90.12, y+34.56, warpFreq, wo, lacunarity, persistence)

			xw := x + wx*warpAmount
			yw := y + wy*warpAmount

			// --- Base FBM sur coords warppées
			v := fbmSine(xw, yw, freq, o, lacunarity, persistence)

			dst[i] = float32(v * gain)
		}
	})
}

package noise

import (
	"geoforge/preset"
	"math"
)

func hash2i(x, y int32, seed uint32) uint32 {
	h := uint32(x)*0x8da6b343 ^ uint32(y)*0xd8163841 ^ seed*0xcb1ab31f
	h ^= h >> 16
	h *= 0x7feb352d
	h ^= h >> 15
	h *= 0x846ca68b
	h ^= h >> 16
	return h
}

func hashFloat(x, y int32, seed uint32) float32 {
	// 24 bits → float32 [-1,1]
	v := hash2i(x, y, seed) >> 8
	return float32(v)*(1.0/8388607.5) - 1.0
}

func NewWhite() Noise {
	var freq float32
	var gain float32
	var seed uint32

	ps := preset.NewParamSet()

	ps = append(ps, preset.NewVariable(
		ParamSeed, ParamSeedLabel,
		0, 0, 100000, 1, 0,
		func(v float32) {
			seed = uint32(v)
		},
	))

	ps = append(ps, preset.NewVariable(
		ParamFrequency, ParamFrequencyLabel,
		0.05, 0.001, 1.0, 0.001, 4,
		func(v float32) {
			freq = v
		},
	))

	ps = append(ps, preset.NewVariable(
		ParamGain, ParamGainLabel,
		1, -1, 1, 0.01, 3,
		func(v float32) {
			gain = v
		},
	))

	return newNoise(ps, func(dst []float32, size int, x0, y0 float32) {
		i := 0
		for y := 0; y < size; y++ {
			wy := y0 + float32(y)
			cy := int32(math.Floor(float64(wy * freq)))

			for x := 0; x < size; x++ {
				wx := x0 + float32(x)
				cx := int32(math.Floor(float64(wx * freq)))

				dst[i] = hashFloat(cx, cy, seed) * gain
				i++
			}
		}
	})
}

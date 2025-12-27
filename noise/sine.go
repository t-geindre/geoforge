package noise

import (
	"geoforge/preset"
	"math"
)

func NewSine() Noise {
	var freq, gain float64

	ps := preset.NewParamSet()
	ps = append(ps, preset.NewVariable(
		ParamFrequency, ParamFrequencyLabel,
		.0002, .0002, .1, .0001, 4,
		func(v float32) {
			freq = float64(v)
		}))
	ps = append(ps, preset.NewVariable(
		ParamGain, ParamGainLabel,
		1, -1, 1, .01, 4,
		func(v float32) {
			gain = float64(v)
		}))

	return newNoise(ps, func(x, y float32) float32 {
		return float32(
			math.Sin(float64(x)*freq)*math.Sin(float64(y)*freq),
		) * float32(gain)
	})
}

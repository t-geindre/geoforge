package noise

import "geoforge/preset"

func NewFixed() Noise {
	var val float32

	ps := preset.NewParamSet()
	ps = append(ps, preset.NewVariable(
		ParamFixedValue, ParamFixedValueLabel,
		0, -1, 1, .001, 3,
		func(v float32) {
			val = v
		}))

	return newNoise(ps, func(_, _ float32) float32 {
		return val
	})
}

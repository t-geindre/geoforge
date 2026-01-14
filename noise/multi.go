package noise

import "geoforge/preset"

type MultiNoise struct {
	noises  []Noise
	current int
	ps      preset.ParamSet
}

func NewMultiNoise(noises ...Noise) Noise {
	m := MultiNoise{
		noises:  noises,
		current: -1,
	}

	m.ps = preset.NewAnonymousParamSet()
	m.ps.Prepend(preset.NewChoice(0, "Noise type", 0, func() []preset.Option[int] {
		opts := make([]preset.Option[int], len(noises))
		for i, n := range noises {
			opts[i] = preset.NewOption(i, n.Name())
		}
		return opts
	}(), func(p preset.Param[int]) {
		if m.current == -1 {
			// None, append noise params
			m.ps.Append(m.noises[p.Val()].Params())
		} else {
			// Switch, remove old noise params, append new noise params
			m.ps.Replace(
				m.noises[m.current].Params(),
				m.noises[p.Val()].Params(),
			)
		}
		m.current = p.Val()
	}))

	return &m
}

func (m MultiNoise) At(x, y float32) float32 {
	return m.noises[m.current].At(x, y)
}

func (m MultiNoise) Params() preset.ParamSet {
	return m.ps
}

func (m MultiNoise) Name() string {
	return "Multi Noise"
}

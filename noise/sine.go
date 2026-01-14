package noise

import (
	"geoforge/preset"
	"math"
)

type sine struct {
	freq float32
	ps   preset.ParamSet
}

func NewSine() Noise {
	s := &sine{}

	s.ps = preset.NewAnonymousParamSet()
	s.ps.Append(preset.NewVariable(0, "Scale", 0.1, 0.001, .5, 0.001, 3, func(p preset.Param[float32]) {
		s.freq = p.Val()
	}))

	return s
}

func (s *sine) At(x, y float32) float32 {
	return float32(math.Sin(float64(x*s.freq)) * math.Sin(float64(y*s.freq)))
}

func (s *sine) Params() preset.ParamSet {
	return s.ps
}

func (s *sine) Name() string {
	return "Sine"
}

package noise

import (
	"geoforge/preset"
	"math"
)

type sine struct {
	freq float32
	ps   preset.ParamSet
	warp *warp
}

func NewSine() Noise {
	s := &sine{
		warp: newWarp(),
	}

	s.ps = preset.NewAnonymousParamSet()
	s.ps.Append(preset.NewVariable(0, "Scale", 0.1, 0.00001, .2, 0.00001, 5, func(p preset.Param[float32]) {
		s.freq = p.Val()
	}))
	s.ps.Append(s.warp.Params())

	return s
}

func (s *sine) At(x, y float32) float32 {
	x, y = s.warp.Warp(x, y)
	return float32(math.Sin(float64(x*s.freq)) * math.Sin(float64(y*s.freq)))
}

func (s *sine) Params() preset.ParamSet {
	return s.ps
}

func (s *sine) Name() string {
	return "Sine"
}

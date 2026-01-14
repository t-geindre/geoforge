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
	s.ps.Append(preset.NewVariable(0, "Scale", 0.5, 0.001, 1.0, 0.001, 3, func(p preset.Param[float32]) {
		s.freq = p.Val()
	}))

	return s
}

func (n *sine) Fill(dst []float32, size int, x0, y0 float32) {
	idx := 0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			px := x0 + float32(x)
			py := y0 + float32(y)
			dst[idx] = float32((math.Sin(float64(px*n.freq)) + math.Sin(float64(py*n.freq))) / 2.0)
			idx++
		}
	}
}

func (n *sine) Params() preset.ParamSet {
	return n.ps
}

func (n *sine) Name() string {
	return "Sine"
}

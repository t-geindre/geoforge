package noise

import (
	"geoforge/preset"
	"math"
)

type pow struct {
	src Noise
	exp float32
	ps  preset.ParamSet
}

func NewPow() Noise {
	p := &pow{
		exp: 1.0,
	}

	p.ps = preset.NewAnonymousParamSet()
	p.ps.Append(preset.NewParam(0, "Src", p.src, func(pr preset.Param[Noise]) {
		p.src = pr.Val()
	}))
	p.ps.Append(preset.NewVariable(0, "Exp", p.exp, 1, 20.0, .01, 2, func(pr preset.Param[float32]) {
		p.exp = pr.Val()
	}))

	return p
}

func (p *pow) At(x, y float32) float32 {
	if p.src == nil {
		return 0
	}

	v := (float64(p.src.At(x, y)) + 1) / 2 // Normalize to [0,1]
	v = math.Pow(v, float64(p.exp))

	return float32(v)*2 - 1 // Denormalize back to [-1,1]
}

func (p *pow) Params() preset.ParamSet {
	return p.ps
}

func (p *pow) Name() string {
	return "Pow"
}

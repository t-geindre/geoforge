package noise

import (
	"geoforge/preset"
)

type clamp struct {
	min, max float32
	src      Noise
	ps       preset.ParamSet
}

func NewClamp() Noise {
	c := &clamp{
		min: -1,
		max: 1,
	}
	c.ps = preset.NewAnonymousParamSet()
	c.ps.Append(preset.NewVariable(0, "Min", c.min, -1, 1, 0.001, 3, func(p preset.Param[float32]) {
		c.min = p.Val()
	}))
	c.ps.Append(preset.NewVariable(1, "Max", c.max, -1, 1, 0.001, 3, func(p preset.Param[float32]) {
		c.max = p.Val()
	}))
	c.ps.Append(preset.NewParam(0, "Source", c.src, func(p preset.Param[Noise]) {
		c.src = p.Val()
	}))

	return c
}

func (c *clamp) At(x, y float32) float32 {
	if c.src == nil {
		return 0
	}

	v := c.src.At(x, y)
	if v < c.min {
		return c.min
	}
	if v > c.max {
		return c.max
	}
	return v
}

func (c *clamp) Params() preset.ParamSet {
	return c.ps
}

func (c *clamp) Name() string {
	return "Clamp"
}

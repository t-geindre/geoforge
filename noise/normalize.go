package noise

import "geoforge/preset"

type normalize struct {
	src      Noise
	min, max float32
	ps       preset.ParamSet
}

func NewNormalize() Noise {
	n := &normalize{
		min: -1,
		max: 1,
		ps:  preset.NewAnonymousParamSet(),
	}

	n.ps.Append(preset.NewParam(0, "Src", n.src, func(pr preset.Param[Noise]) {
		n.src = pr.Val()
	}))

	n.ps.Append(preset.NewVariable(0, "Min", n.min, -1.0, 1.0, .01, 2, func(pr preset.Param[float32]) {
		n.min = pr.Val()
	}))

	n.ps.Append(preset.NewVariable(0, "Max", n.max, -1.0, 1.0, .01, 2, func(pr preset.Param[float32]) {
		n.max = pr.Val()
	}))

	return n
}

func (n *normalize) At(x, y float32) float32 {
	if n.src == nil {
		return 0
	}

	// Expect [-1,1] input, map to [0,1], then scale to [min,max]
	return (n.src.At(x, y)+1)/2*(n.max-n.min) + n.min

}

func (n *normalize) Params() preset.ParamSet {
	return n.ps
}

func (n *normalize) Name() string {
	return "Normalize"
}

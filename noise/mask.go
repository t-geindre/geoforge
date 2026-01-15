package noise

import "geoforge/preset"

type mask struct {
	low, high Noise
	mask      Noise
	edge      float32
	smooth    float32

	ps preset.ParamSet
}

func NewMask() Noise {
	m := &mask{}

	m.ps = preset.NewAnonymousParamSet()
	m.ps.Append(preset.NewParam(0, "Mask", nil, func(p preset.Param[Noise]) {
		m.mask = p.Val()
	}))
	m.ps.Append(preset.NewParam(0, "Low", nil, func(p preset.Param[Noise]) {
		m.low = p.Val()
	}))
	m.ps.Append(preset.NewParam(0, "High", nil, func(p preset.Param[Noise]) {
		m.high = p.Val()
	}))
	m.ps.Append(preset.NewVariable(0, "Edge", 0.0, -1.0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		m.edge = p.Val()
	}))
	m.ps.Append(preset.NewVariable(0, "Smoothness", 0.1, 0.0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		m.smooth = p.Val()
	}))

	return m
}

func (m *mask) At(x, y float32) float32 {
	if m.mask == nil || m.low == nil || m.high == nil {
		return 0
	}

	mv := m.mask.At(x, y)
	if mv < m.edge-m.smooth {
		return m.low.At(x, y)
	}

	if mv > m.edge+m.smooth {
		return m.high.At(x, y)
	}

	t := (mv - (m.edge - m.smooth)) / (2 * m.smooth)
	va := m.low.At(x, y)
	vb := m.high.At(x, y)
	return va*(1-t) + vb*t
}

func (m *mask) Params() preset.ParamSet {
	return m.ps
}

func (m *mask) Name() string {
	return "Mask"
}

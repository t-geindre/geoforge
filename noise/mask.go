package noise

import "geoforge/preset"

type mask struct {
	low, high     Noise
	clow, chigh   float32 // const
	wchigh, wclow bool    // with const

	mask   Noise
	edge   float32
	smooth float32

	ps preset.ParamSet
}

func NewMask() Noise {
	m := &mask{}

	m.ps = preset.NewAnonymousParamSet()
	m.ps.Append(preset.NewParam(0, "Mask", nil, func(p preset.Param[Noise]) {
		m.mask = p.Val()
	}))

	// Low param: either constant or noise
	pnl := preset.NewParam(798, "Low", nil, func(p preset.Param[Noise]) {
		m.low = p.Val()
	})
	pcl := preset.NewVariable(465, "Low value", 0.0, -1.0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		m.clow = p.Val()
	})
	m.ps.Append(preset.NewParam(0, "Low const", m.wclow, func(p preset.Param[bool]) {
		m.wclow = p.Val()
		if m.wclow {
			m.ps.ReplaceById(pnl.Id(), pcl)
		} else {
			m.ps.ReplaceById(pcl.Id(), pnl)
		}
	}))
	m.ps.Append(pnl)

	// High param: either constant or noise
	pnh := preset.NewParam(799, "High", nil, func(p preset.Param[Noise]) {
		m.high = p.Val()
	})
	pch := preset.NewVariable(466, "High value", 1.0, -1.0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		m.chigh = p.Val()
	})
	m.ps.Append(preset.NewParam(0, "High const", m.wchigh, func(p preset.Param[bool]) {
		m.wchigh = p.Val()
		if m.wchigh {
			m.ps.ReplaceById(pnh.Id(), pch)
		} else {
			m.ps.ReplaceById(pch.Id(), pnh)
		}
	}))
	m.ps.Append(pnh)

	// Transition
	m.ps.Append(preset.NewVariable(0, "Edge", 0.0, -1.0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		m.edge = p.Val()
	}))
	m.ps.Append(preset.NewVariable(0, "Smoothness", 0.1, 0.0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		m.smooth = p.Val()
	}))

	return m
}

func (m *mask) At(x, y float32) float32 {
	if m.mask == nil {
		return 0
	}

	mv := m.mask.At(x, y)
	if mv < m.edge-m.smooth {
		return m.getLow(x, y)
	}

	if mv > m.edge+m.smooth {
		return m.getHigh(x, y)
	}

	t := (mv - (m.edge - m.smooth)) / (2 * m.smooth)
	va := m.getLow(x, y)
	vb := m.getHigh(x, y)
	return va*(1-t) + vb*t
}

func (m *mask) Params() preset.ParamSet {
	return m.ps
}

func (m *mask) Name() string {
	return "Mask"
}

func (m *mask) getLow(x, y float32) float32 {
	if m.wclow {
		return m.clow
	}
	if m.low == nil {
		return 0
	}
	return m.low.At(x, y)
}

func (m *mask) getHigh(x, y float32) float32 {
	if m.wchigh {
		return m.chigh
	}
	if m.high == nil {
		return 0
	}
	return m.high.At(x, y)
}

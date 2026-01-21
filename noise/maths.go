package noise

import (
	"geoforge/preset"
)

const (
	mathsOpAdd = iota
	mathsOpSub
	mathsOpMul
	mathsOpDiv
	mathsOpMax
	mathsOpMin
)

const (
	mathsPmNoiseA = iota
	mathsPmNoiseB
	mathsPmWithConst
	mathsPmConst
	mathsPmOp
)

type Maths struct {
	nsa, nsb Noise
	c        float32
	wc       bool // With const
	ps       preset.ParamSet
	op       int
}

func NewMaths() Noise {
	m := &Maths{
		ps: preset.NewAnonymousParamSet(),
	}

	// Operations
	po := preset.NewChoice(mathsPmOp, "Operation", m.op, []preset.Option[int]{
		preset.NewOption[int](mathsOpAdd, "A + B"),
		preset.NewOption[int](mathsOpSub, "A - B"),
		preset.NewOption[int](mathsOpMul, "A x B"),
		preset.NewOption[int](mathsOpDiv, "A / B"),
		preset.NewOption[int](mathsOpMax, "MAX(A, B)"),
		preset.NewOption[int](mathsOpMin, "MIN(A, B)"),
	}, func(p preset.Param[int]) {
		m.op = p.Val()
	})

	// Noises sources
	pna := preset.NewParam(mathsPmNoiseA, "A", m.nsa, func(p preset.Param[Noise]) {
		m.nsa = p.Val()
	})
	pnb := preset.NewParam(mathsPmNoiseB, "B", m.nsb, func(p preset.Param[Noise]) {
		m.nsb = p.Val()
	})

	// Constant
	pc := preset.NewVariable(mathsPmConst, "B", m.c, -20, 20, .001, 3, func(p preset.Param[float32]) {
		m.c = p.Val()
	})

	pwc := preset.NewParam(mathsPmWithConst, "Use constant", m.wc, func(p preset.Param[bool]) {
		m.wc = p.Val()
		if m.wc {
			m.ps.ReplaceById(pnb.Id(), pc)
		} else {
			m.ps.ReplaceById(pc.Id(), pnb)
		}
	})

	// Append
	m.ps.Append(
		pwc, po, pna, pnb,
	)

	return m
}

func (m *Maths) At(x, y float32) float32 {
	if m.nsa == nil {
		return 0
	}
	a := m.nsa.At(x, y)

	b := m.c
	if !m.wc {
		if m.nsb == nil {
			return 0
		}
		b = m.nsb.At(x, y)
	}

	v := float32(0)
	switch m.op {
	case mathsOpAdd:
		v = a + b
	case mathsOpSub:
		v = a - b
	case mathsOpMul:
		v = a * b
	case mathsOpDiv:
		v = a / b
	case mathsOpMax:
		if a > b {
			v = a
		} else {
			v = b
		}
	case mathsOpMin:
		if a > b {
			v = b
		} else {
			v = a
		}
	}

	return clamp11(v)
}

func (m *Maths) Params() preset.ParamSet {
	return m.ps
}

func (m *Maths) Name() string {
	return "Maths"
}

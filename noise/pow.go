package noise

import "math"

type Pow struct {
	inner Noise
	exp   float64
}

func NewPow(inner Noise, exp float64) Noise {
	return &Pow{inner, exp}
}

func (n *Pow) Eval(x, y float64) float64 {
	v := n.inner.Eval(x, y)
	if v < 0 {
		v = 0
	}
	return math.Pow(v, n.exp)
}

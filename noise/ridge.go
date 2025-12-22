package noise

import "math"

type Ridge struct{ inner Noise }

func NewRidge(inner Noise) Noise {
	return &Ridge{inner}
}

func (n *Ridge) Eval(x, y float64) float64 {
	v := n.inner.Eval(x, y)
	// ridge = 1 - |2v-1|
	r := 1 - math.Abs(2*v-1)
	if r < 0 {
		return 0
	}
	return r
}

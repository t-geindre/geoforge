package noise

type Clamp struct {
	inner    Noise
	min, max float64
}

func NewClamp(inner Noise, min, max float64) Noise {
	return &Clamp{inner, min, max}
}

func (n *Clamp) Eval(x, y float64) float64 {
	v := n.inner.Eval(x, y)
	if v < n.min {
		return n.min
	}
	if v > n.max {
		return n.max
	}
	return v
}

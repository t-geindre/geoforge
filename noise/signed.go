package noise

type Signed struct{ inner Noise }

func NewSigned(inner Noise) Noise {
	return &Signed{inner}
}

func (n *Signed) Eval(x, y float64) float64 {
	return n.inner.Eval(x, y)*2 - 1
}

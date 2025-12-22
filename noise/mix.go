package noise

type Mix struct{ a, b, mask Noise }

func NewMix(a, b, mask Noise) Noise {
	return &Mix{a, b, mask}
}

func (n *Mix) Eval(x, y float64) float64 {
	m := n.mask.Eval(x, y)
	if m < 0 {
		m = 0
	} else if m > 1 {
		m = 1
	}
	av := n.a.Eval(x, y)
	bv := n.b.Eval(x, y)
	return av + m*(bv-av)
}

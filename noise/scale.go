package noise

type Scale struct {
	inner  Noise
	factor float64
}

func NewScale(inner Noise, factor float64) Noise {
	return &Scale{
		inner:  inner,
		factor: factor,
	}
}

func (s *Scale) Eval(x, y float64) float64 {
	return s.inner.Eval(x*s.factor, y*s.factor)
}

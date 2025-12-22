package noise

type Smoothstep struct {
	inner        Noise
	edge0, edge1 float64
}

func NewSmoothstep(inner Noise, edge0, edge1 float64) Noise {
	return &Smoothstep{
		inner: inner,
		edge0: edge0,
		edge1: edge1,
	}
}

func (s *Smoothstep) Eval(x, y float64) float64 {
	v := s.inner.Eval(x, y)
	if v <= s.edge0 {
		return 0
	}
	if v >= s.edge1 {
		return 1
	}
	// Scale v to [0, 1]
	t := (v - s.edge0) / (s.edge1 - s.edge0)
	// Apply smoothstep formula
	return t * t * (3 - 2*t)
}

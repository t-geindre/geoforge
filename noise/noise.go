package noise

type Noise interface {
	Eval(x, y float64) float64
}

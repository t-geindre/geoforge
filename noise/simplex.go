package noise

import (
	"github.com/ojrac/opensimplex-go"
)

type OpenSimplex struct {
	os opensimplex.Noise
}

func NewOpenSimplex(seed int64) Noise {
	return &OpenSimplex{
		os: opensimplex.NewNormalized(seed),
	}
}

func (n *OpenSimplex) Eval(x, y float64) float64 {
	return n.os.Eval2(x, y)
}

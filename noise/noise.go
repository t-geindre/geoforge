package noise

import (
	"geoforge/preset"

	fastnoise "github.com/Auburn/FastNoiseLite/Go"
)

const DomainWarpNone fastnoise.DomainWarpType = -1

type Noise interface {
	At(x, y float32) float32
	Params() preset.ParamSet
	Name() string
}

type atFunc func(x, y float32) float32

type noise struct {
	ps preset.ParamSet
	at atFunc
	n  string
}

func newNoise(name string, init func(ps preset.ParamSet) atFunc) Noise {
	n := &noise{
		n: name,
	}

	n.ps = preset.NewAnonymousParamSet()
	n.at = init(n.ps)

	return n
}

func (n *noise) At(x, y float32) float32 {
	return n.at(x, y)
}

func (n *noise) Params() preset.ParamSet {
	return n.ps
}

func (n *noise) Name() string {
	return n.n
}

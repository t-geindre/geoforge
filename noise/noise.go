package noise

import (
	"geoforge/preset"

	fastnoise "github.com/Auburn/FastNoiseLite/Go"
)

const DomainWarpNone fastnoise.DomainWarpType = -1

type Noise interface {
	Fill(dst []float32, size int, x0, y0 float32)
	Params() preset.ParamSet
}

type fillFunc func(dst []float32, size int, x0, y0 float32)

type noise struct {
	fill   fillFunc
	ps     preset.ParamSet
	fsw    *fastnoise.State[float32]
	doWarp bool
}

type fastNoise struct {
	fsn    *fastnoise.State[float32]
	fsw    *fastnoise.State[float32]
	doWarp bool
	ps     preset.ParamSet
}

package noise

import (
	"geoforge/preset"

	fastnoise "github.com/Auburn/FastNoiseLite/Go"
)

const DomainWarpNone fastnoise.DomainWarpType = -1

type Noise interface {
	Fill(dst []float32, size int, x0, y0 float32)
	Params() preset.ParamSet
	Name() string
}

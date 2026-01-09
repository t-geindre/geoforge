package render

import (
	"geoforge/preset"

	"github.com/hajimehoshi/ebiten/v2"
)

type ChunkRenderer interface {
	DrawChunk(dst *ebiten.Image, w, h int, op *ebiten.DrawRectShaderOptions)
	Params() preset.ParamSet
	Name() string
}

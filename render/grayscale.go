package render

import (
	_ "embed"
	"geoforge/preset"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed grayscale.kage
var grayscaleShdRaw []byte

type GrayScale struct {
	sh *ebiten.Shader
	ps preset.ParamSet
}

func NewGrayScale() *GrayScale {

	shd, err := ebiten.NewShader(grayscaleShdRaw)
	if err != nil {
		panic(err)
	}

	return &GrayScale{
		sh: shd,
		ps: preset.NewAnonymousParamSet(),
	}
}

func (g *GrayScale) DrawChunk(dst *ebiten.Image, w, h int, op *ebiten.DrawRectShaderOptions) {
	dst.DrawRectShader(w, h, g.sh, op)
}

func (g *GrayScale) Params() preset.ParamSet {
	return g.ps
}

func (g *GrayScale) Name() string {
	return "Gray scale"
}

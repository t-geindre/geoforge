package render

import (
	"geoforge/preset"

	"github.com/hajimehoshi/ebiten/v2"
)

type ChunkRenderer interface {
	DrawChunk(dst *ebiten.Image, w, h int, op *ebiten.DrawRectShaderOptions)
	Update()
	Params() preset.ParamSet
	Name() string
}

type chunkRenderer struct {
	sh       *ebiten.Shader
	ps       preset.ParamSet
	uniforms map[string]interface{}
	name     string
	update   func()
}

func newChunkRenderer(name string, shader []byte) *chunkRenderer {
	shd, err := ebiten.NewShader(shader)
	if err != nil {
		panic(err)
	}

	return &chunkRenderer{
		sh:       shd,
		ps:       preset.NewAnonymousParamSet(),
		uniforms: make(map[string]interface{}),
		name:     name,
		update:   func() {},
	}
}

func (g *chunkRenderer) DrawChunk(dst *ebiten.Image, w, h int, op *ebiten.DrawRectShaderOptions) {
	for k, v := range g.uniforms {
		op.Uniforms[k] = v
	}
	dst.DrawRectShader(w, h, g.sh, op)
}

func (g *chunkRenderer) Params() preset.ParamSet {
	return g.ps
}

func (g *chunkRenderer) Name() string {
	return g.name
}

func (g *chunkRenderer) Update() {
	g.update()
}

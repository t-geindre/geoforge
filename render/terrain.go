package render

import (
	_ "embed"
	"geoforge/preset"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed terrain.kage
var terrainShdRaw []byte

type Terrain struct {
	sh *ebiten.Shader
	ps preset.ParamSet

	ambientLight float32
	seaLevel     float32
}

func NewTerrain() *Terrain {
	t := &Terrain{}

	shd, err := ebiten.NewShader(terrainShdRaw)
	if err != nil {
		panic(err)
	}

	ps := preset.NewAnonymousParamSet()
	ps.Append(preset.NewVariable(1, "Ambient Light", 0.35, 0.0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		t.ambientLight = p.Val()
	}))
	ps.Append(preset.NewVariable(1, "Sea Level", 0.5, -0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		t.seaLevel = p.Val()
	}))

	t.sh = shd
	t.ps = ps

	return t
}

func (t *Terrain) DrawChunk(dst *ebiten.Image, w, h int, op *ebiten.DrawRectShaderOptions) {
	x, y := ebiten.CursorPosition()
	op.Uniforms["LightPos"] = [2]float32{
		float32(x),
		float32(y),
	}
	op.Uniforms["Ambient"] = t.ambientLight
	op.Uniforms["SeaLevel"] = t.seaLevel
	dst.DrawRectShader(w, h, t.sh, op)
}

func (t *Terrain) Params() preset.ParamSet {
	return t.ps
}

func (t *Terrain) Name() string {
	return "Terrain"
}

package render

import (
	_ "embed"
	"geoforge/preset"
	"geoforge/world"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed terrain.kage
var terrainShdRaw []byte

type Terrain struct {
	sh *ebiten.Shader
	ps preset.ParamSet

	ambientLight   float32
	seaLevel       float32
	normalStrength float32
	normalEps      float32
}

func NewTerrain() *Terrain {
	t := &Terrain{}

	shd, err := ebiten.NewShader(terrainShdRaw)
	if err != nil {
		panic(err)
	}

	lights := preset.NewParamSet(0, "Lighting")
	normals := preset.NewParamSet(0, "Normals")
	levels := preset.NewParamSet(0, "Levels")

	ps := preset.NewAnonymousParamSet()
	ps.Append(lights)
	ps.Append(normals)
	ps.Append(levels)

	lights.Append(preset.NewVariable(1, "Ambient", 0.35, 0.0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
		t.ambientLight = p.Val()
	}))

	normals.Append(preset.NewVariable(1, "Strength", 70, 0.0, 250.0, 1, 0, func(p preset.Param[float32]) {
		t.normalStrength = p.Val()
	}))
	normals.Append(preset.NewVariable(1, "Epsilon", 1, 1, world.ChunkApron, 1, 0, func(p preset.Param[float32]) {
		t.normalEps = p.Val()
	}))
	levels.Append(preset.NewVariable(1, "Sea", 0.5, -0, 1.0, 0.01, 2, func(p preset.Param[float32]) {
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
	op.Uniforms["NormalStrength"] = t.normalStrength
	op.Uniforms["NormalEps"] = t.normalEps
	dst.DrawRectShader(w, h, t.sh, op)
}

func (t *Terrain) Params() preset.ParamSet {
	return t.ps
}

func (t *Terrain) Name() string {
	return "Terrain"
}

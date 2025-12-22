package render

import (
	_ "embed"
	"geoforge/cam"
	"geoforge/geo"
	"geoforge/world"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

//go:embed shader.kage
var Shader []byte

type Renderer struct {
	drawn  int
	shader *ebiten.Shader
}

func NewRenderer() *Renderer {
	shd, err := ebiten.NewShader(Shader)
	if err != nil {
		panic(err) // todo
	}

	return &Renderer{
		shader: shd,
	}
}

func (r *Renderer) Draw(w *world.World, cam cam.Camera, dst *ebiten.Image) {
	r.drawn = 0
	z := cam.Zoom()
	if z <= 0 {
		return
	}

	csWorld := w.ChunkSize()
	csScreen := csWorld * z
	worldRect := cam.WorldRect()

	for _, c := range w.Chunks() {
		wx := float64(c.Id().X) * csWorld
		wy := float64(c.Id().Y) * csWorld

		cRect := geo.NewRect(wx, wy, wx+csWorld, wy+csWorld)
		if !worldRect.Intersects(cRect) {
			continue
		}

		sx, sy := cam.WorldToScreen(wx, wy)
		if c.Is(world.ChunkStateReady) {
			hm := c.GetLayer(world.LayerHeightMap)
			if hm != nil {
				op := &ebiten.DrawRectShaderOptions{}
				op.Images = [4]*ebiten.Image{hm}
				op.GeoM.Scale(z, z)
				op.GeoM.Translate(sx, sy)
				dst.DrawRectShader(int(csWorld), int(csWorld), r.shader, op)

				r.drawn++
				continue
			}
		}

		fill := debugChunkColor(c.Id(), 0x80)
		vector.FillRect(dst, float32(sx), float32(sy), float32(csScreen), float32(csScreen), fill, false)

		r.drawn++
	}
}

func (r *Renderer) DrawnChunks() int {
	return r.drawn
}

func debugChunkColor(id world.ChunkId, alpha uint8) color.RGBA {
	// Simple stable hash
	x := uint64(id.X)
	y := uint64(id.Y)
	h := x*0x9e3779b97f4a7c15 ^ y*0xbf58476d1ce4e5b9

	return color.RGBA{
		R: uint8(h),
		G: uint8(h >> 8),
		B: uint8(h >> 16),
		A: alpha,
	}
}

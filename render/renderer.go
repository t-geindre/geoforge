package render

import (
	_ "embed"
	"geoforge/cam"
	"geoforge/geo"
	"geoforge/preset"
	"geoforge/world"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Renderer struct {
	drawn     int
	ps        preset.ParamSet
	renderers []ChunkRenderer
	current   int
}

func NewRenderer() *Renderer {
	r := &Renderer{
		renderers: []ChunkRenderer{
			NewColorScale(),
			NewTerrain(),
		},
	}

	r.buildParams()

	return r
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
		hm := c.GetLayer(world.LayerHeightMap)

		if hm != nil && c.Is(world.ChunkStateReady) {
			bds := hm.Bounds()
			op := &ebiten.DrawRectShaderOptions{}
			op.Images = [4]*ebiten.Image{hm}
			originX := float32(sx - w.Apron()*z)
			originY := float32(sy - w.Apron()*z)

			op.Uniforms = map[string]any{
				"Apron":     float32(w.Apron()),
				"ChunkSize": float32(w.ChunkSize()),
				"Zoom":      float32(z),
				"Origin":    []float32{originX, originY},
			}
			op.GeoM.Scale(z, z)
			op.GeoM.Translate(
				sx-w.Apron()*z,
				sy-w.Apron()*z,
			)
			r.renderers[r.current].DrawChunk(dst, bds.Dx(), bds.Dy(), op)

			r.drawn++
			continue
		}

		fill := debugChunkColor(c.Id(), 0x80)
		vector.StrokeRect(dst, float32(sx), float32(sy), float32(csScreen-1), float32(csScreen-1), 1, fill, false)
		vector.StrokeLine(dst, float32(sx), float32(sy), float32(sx+csScreen), float32(sy+csScreen), 1, fill, false)
		vector.StrokeLine(dst, float32(sx+csScreen), float32(sy), float32(sx), float32(sy+csScreen), 1, fill, false)
		ebitenutil.DebugPrintAt(dst, c.Id().String(), int(sx+2), int(sy+2))

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

func (r *Renderer) Params() preset.ParamSet {
	return r.ps
}

func (r *Renderer) buildParams() {
	ops := make([]preset.Option[int], len(r.renderers))
	for i, rd := range r.renderers {
		ops[i] = preset.NewOption(i, rd.Name())
	}

	if r.ps == nil {
		r.ps = preset.NewAnonymousParamSet()
	}
	r.ps.Clear()

	r.ps.Append(preset.NewChoice(0, "Renderer", r.current, ops, func(p preset.Param[int]) {
		if r.current == p.Val() {
			return
		}

		r.current = p.Val()
		r.buildParams()
	}))

	r.ps.Append(r.renderers[r.current].Params().All()...)
}

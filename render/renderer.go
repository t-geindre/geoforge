package render

import (
	_ "embed"
	"geoforge/camera"
	"geoforge/geo"
	"geoforge/preset"
	"geoforge/world"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer struct {
	drawn     int
	ps        preset.ParamSet
	renderers []ChunkRenderer
	current   int

	world   *world.World
	worldSt uint8

	cam   camera.Camera
	camSt uint8

	hm *ebiten.Image
}

func NewRenderer(w *world.World, c camera.Camera) *Renderer {
	r := &Renderer{
		world:   w,
		cam:     c,
		camSt:   c.RegisterChangeId(),
		worldSt: w.RegisterChangeId(),
		renderers: []ChunkRenderer{
			NewColorScale(),
			NewTerrain(),
		},
	}

	r.buildParams()

	return r
}

func (r *Renderer) Update() {
	r.renderers[r.current].Update()
}

func (r *Renderer) Draw(dst *ebiten.Image) {
	if r.cam.HasChanged(r.camSt) || r.world.HasChanged(r.worldSt) {
		r.drawHeightMap()
	}

	ww, wh := r.cam.GetViewport()
	op := &ebiten.DrawRectShaderOptions{
		Images:   [4]*ebiten.Image{r.hm},
		Uniforms: map[string]any{},
	}

	r.renderers[r.current].Draw(dst, ww, wh, op)
}
func (r *Renderer) drawHeightMap() {
	ww, wh := r.cam.GetViewport()
	if r.hm == nil || r.hm.Bounds().Dx() != ww || r.hm.Bounds().Dy() != wh {
		r.hm = ebiten.NewImage(ww, wh)
	}

	r.drawn = 0
	z := r.cam.Zoom()
	if z <= 0 {
		return
	}

	worldRect := r.cam.WorldRect()

	for _, c := range r.world.Chunks() {
		wx := float64(c.Id().X) * world.ChunkSize
		wy := float64(c.Id().Y) * world.ChunkSize

		cRect := geo.NewRect(wx, wy, wx+world.ChunkSize, wy+world.ChunkSize)
		if !worldRect.Intersects(cRect) {
			continue
		}

		sx, sy := r.cam.WorldToScreen(wx, wy)
		hm := c.GetHeightMap()

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(z, z)
		op.GeoM.Translate(sx, sy)

		r.hm.DrawImage(hm, op)
		r.drawn++
	}
}

func (r *Renderer) DrawnChunks() int {
	return r.drawn
}

func debugChunkColor(id world.ChunkId, alpha uint8) color.RGBA {
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

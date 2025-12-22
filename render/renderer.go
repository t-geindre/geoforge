package render

import (
	"awesomeProject/cam"
	"awesomeProject/world"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer struct{}

func NewRenderer() *Renderer { return &Renderer{} }

func (r *Renderer) Draw(w *world.World, cam *cam.Camera, dst *ebiten.Image) {
	z := cam.Zoom()
	if z <= 0 {
		return
	}

	csWorld := w.ChunkSize()
	csScreen := csWorld * z

	for _, c := range w.Chunks() {
		wx := float64(c.Id().X) * csWorld
		wy := float64(c.Id().Y) * csWorld

		sx, sy := cam.WorldToScreen(wx, wy)

		vector.StrokeRect(dst, float32(sx), float32(sy), float32(csScreen), float32(csScreen), 0, colornames.Blue, false)
	}
}

func debugChunkColor(id world.ChunkId, alpha uint8) color.RGBA {
	// hash simple stable
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

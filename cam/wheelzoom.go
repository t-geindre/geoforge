package cam

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const minZoom = 0.3
const maxZoom = 10.0

type WheelZoom struct {
	Camera
}

func NewWheelZoom(inner Camera) Camera {
	return &WheelZoom{
		Camera: inner,
	}
}

func (c *WheelZoom) Update() {
	_, zw := ebiten.Wheel()

	if zw != 0 {
		old := c.Zoom()

		factor := 1.1
		if zw < 0 {
			factor = 1.0 / factor
		}

		// clamp via cible, puis retransforme en facteur réel
		target := old * factor
		if target < minZoom {
			target = minZoom
		} else if target > maxZoom {
			target = maxZoom
		}
		factor = target / old

		x, y := ebiten.CursorPosition()
		c.ZoomAt(factor, float64(x), float64(y))
	}

	c.Camera.Update()
}

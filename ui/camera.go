package ui

import (
	"fmt"

	"github.com/ebitengine/debugui"
)

type Cam interface {
	MoveTo(x, y float64)
	Position() (x, y float64)
	SetZoom(zoom float64)
	Zoom() float64
}

type Camera struct {
	cam Cam
}

func NewCamera(cam Cam) *Camera {
	return &Camera{cam: cam}
}

func (c *Camera) UiUpdate(ctx *debugui.Context) {
	ctx.Header("Camera", true, func() {
		ctx.SetGridLayout([]int{-1, 70}, nil)
		ctx.Text(fmt.Sprintf("Zoom: %.0f%%", c.cam.Zoom()*100))
		ctx.Button("Reset").On(func() {
			c.cam.SetZoom(1)
		})
		x, y := c.cam.Position()
		ctx.Text(fmt.Sprintf("Position: (%.0f, %.0f)", x, y))
		ctx.Button("Center").On(func() {
			c.cam.MoveTo(0, 0)
		})
	})
}

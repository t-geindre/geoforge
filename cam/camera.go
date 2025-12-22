package cam

import "awesomeProject/geo"

type Camera struct {
	x, y float64 // World coordinates
	zoom float64 // 1 = 100%, >1 = zoom in, <1 = zoom out
	w, h int     // Viewport size, pixels
}

func NewCamera() *Camera {
	return &Camera{
		x:    0,
		y:    0,
		zoom: 1,
		w:    800,
		h:    600,
	}
}

func (c *Camera) SetViewport(w, h int) {
	c.w = w
	c.h = h
}

func (c *Camera) WorldRect() geo.Rect {
	halfW := float64(c.w) * 0.5 / c.zoom
	halfH := float64(c.h) * 0.5 / c.zoom

	return geo.NewRect(c.x-halfW, c.y-halfH, c.x+halfW, c.y+halfH)
}

func (c *Camera) WorldToScreen(wx, wy float64) (sx, sy float64) {
	sx = (wx-c.x)*c.zoom + float64(c.w)/2
	sy = (wy-c.y)*c.zoom + float64(c.h)/2
	return
}

func (c *Camera) ScreenToWorld(sx, sy float64) (wx, wy float64) {
	wx = (sx-float64(c.w)/2)/c.zoom + c.x
	wy = (sy-float64(c.h)/2)/c.zoom + c.y
	return
}

func (c *Camera) Move(dx, dy float64) {
	c.x += dx
	c.y += dy
}

func (c *Camera) ZoomAt(factor float64, screenX, screenY float64) {
	if factor <= 0 {
		return
	}

	wx, wy := c.ScreenToWorld(screenX, screenY)
	c.zoom *= factor
	wx2, wy2 := c.ScreenToWorld(screenX, screenY)

	// Center the camera to keep the point under the cursor fixed
	c.x += wx - wx2
	c.y += wy - wy2
}

func (c *Camera) Zoom() float64 {
	return c.zoom
}

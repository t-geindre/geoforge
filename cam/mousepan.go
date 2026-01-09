package cam

import "github.com/hajimehoshi/ebiten/v2"

type MousePan struct {
	Camera
	isDragging   bool
	lastX, lastY int
}

func NewMousePan(inner Camera) Camera {
	return &MousePan{
		Camera: inner,
	}
}

func (mp *MousePan) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if !mp.isDragging {
			mp.isDragging = true
			mp.lastX = x
			mp.lastY = y
		} else {
			dx := x - mp.lastX
			dy := y - mp.lastY
			mp.Move(float64(-dx)/mp.Zoom(), float64(-dy)/mp.Zoom())
			mp.lastX = x
			mp.lastY = y
		}
	}
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mp.isDragging = false
	}
	mp.Camera.Update()
}

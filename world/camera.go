package world

import "awesomeProject/geo"

type Camera interface {
	WorldRect() geo.Rect
}

package geo

type Rect struct {
	MinX, MinY float64
	MaxX, MaxY float64
}

func NewRect(minX, minY, maxX, maxY float64) Rect {
	return Rect{
		MinX: minX,
		MinY: minY,
		MaxX: maxX,
		MaxY: maxY,
	}
}

func (r Rect) Expand(margin float64) Rect {
	return Rect{
		MinX: r.MinX - margin,
		MinY: r.MinY - margin,
		MaxX: r.MaxX + margin,
		MaxY: r.MaxY + margin,
	}
}

func (r Rect) Intersects(other Rect) bool {
	return r.MinX < other.MaxX && r.MaxX > other.MinX &&
		r.MinY < other.MaxY && r.MaxY > other.MinY
}

// SnapOut expands the rect to the grid defined by `size`.
// Result bounds are multiples of `size` and fully contain the original rect.
func (r Rect) SnapOut(size float64) Rect {
	if size <= 0 {
		return r
	}

	return Rect{
		MinX: floorTo(r.MinX, size),
		MinY: floorTo(r.MinY, size),
		MaxX: ceilTo(r.MaxX, size),
		MaxY: ceilTo(r.MaxY, size),
	}
}

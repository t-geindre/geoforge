package geo

import "math"

func floorTo(x, size float64) float64 { return math.Floor(x/size) * size }
func ceilTo(x, size float64) float64  { return math.Ceil(x/size) * size }

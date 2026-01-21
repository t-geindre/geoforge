package noise

func clamp11(v float32) float32 {
	if v < -1 {
		v = -1
	}
	if v > 1 {
		v = 1
	}
	return v
}

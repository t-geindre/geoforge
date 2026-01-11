package render

import "image/color"

func rgbaToF32(c color.RGBA) [3]float32 {
	return [3]float32{
		float32(c.R) / 255.0,
		float32(c.G) / 255.0,
		float32(c.B) / 255.0,
	}
}

func f32ToRgba(f [3]float32) color.RGBA {
	return color.RGBA{
		R: uint8(f[0] * 255.0),
		G: uint8(f[1] * 255.0),
		B: uint8(f[2] * 255.0),
		A: 255,
	}
}

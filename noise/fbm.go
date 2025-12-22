package noise

// Fbm implements fractal brownian motion noise.
type Fbm struct {
	inner       Noise
	octaves     int
	persistence float64
	lacunarity  float64
}

// NewFbm creates a new Fbm noise generator.
func NewFbm(inner Noise, octaves int, persistence, lacunarity float64) Noise {
	return &Fbm{
		inner:       inner,
		octaves:     octaves,
		persistence: persistence,
		lacunarity:  lacunarity,
	}
}

// Value returns the Fbm noise value at (x, y).
func (f *Fbm) Eval(x, y float64) float64 {
	var total float64
	var frequency float64 = 1.0
	var amplitude float64 = 1.0
	var maxAmplitude float64 = 0.0

	for i := 0; i < f.octaves; i++ {
		noiseValue := f.inner.Eval(x*frequency, y*frequency)
		total += noiseValue * amplitude

		maxAmplitude += amplitude
		amplitude *= f.persistence
		frequency *= f.lacunarity
	}

	if maxAmplitude == 0 {
		return 0
	}
	return total / maxAmplitude
}

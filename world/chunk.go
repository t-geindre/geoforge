package world

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ChunkStateDirty = iota
	ChunkStateQueued
	ChunkStateReady
)

const (
	LayerHeightMap = iota
)

const (
	// ChunkSize size of the chunk without apron
	ChunkSize = 256
	// ChunkSurface total surface including apron (X * Y)
	ChunkSurface = ChunkSize * ChunkSize
)

type Chunk struct {
	id    ChunkId
	state int
	gen   uint64
	hm    *ebiten.Image // heightmap, R = height
}

func NewChunk(id ChunkId) *Chunk {
	return &Chunk{
		id:    id,
		state: ChunkStateDirty,
		hm:    ebiten.NewImage(ChunkSize, ChunkSize),
	}
}

func (c *Chunk) Id() ChunkId {
	return c.id
}

// WritePixels writes the given grayscale pixels to the heightmap if the generation matches
// Converts grayscale (1 byte per pixel) to RGBA (4 bytes per pixel) for ebiten.Image
// rgbaBuf must be at least 4*ChunkSurface bytes and will be reused to avoid allocations
func (c *Chunk) WritePixels(gen uint64, pixels []byte, rgbaBuf []byte) bool {
	if c.gen != gen {
		return false
	}

	// Convert grayscale to RGBA format
	for i, gray := range pixels {
		idx := i * 4
		rgbaBuf[idx] = gray   // R
		rgbaBuf[idx+1] = gray // G
		rgbaBuf[idx+2] = gray // B
		rgbaBuf[idx+3] = 255  // A
	}

	c.hm.WritePixels(rgbaBuf)

	return true
}

func (c *Chunk) GetHeightMap() *ebiten.Image {
	return c.hm
}

func (c *Chunk) SetState(state int) {
	c.state = state
}

func (c *Chunk) Is(state int) bool {
	return c.state == state
}

func (c *Chunk) GetGeneration() uint64 {
	return c.gen
}

func (c *Chunk) BumpGeneration() {
	c.gen++
}

func (c *Chunk) ResetGeneration() {
	c.gen = 0
}

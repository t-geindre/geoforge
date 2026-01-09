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
	// ChunkApron size of the apron around the chunk for seamless generation
	ChunkApron = 10
	// ChunkDimSize one dimension size including apron (X or Y)
	ChunkDimSize = ChunkSize + 2*ChunkApron
	// ChunkSurface total surface including apron (X * Y)
	ChunkSurface = ChunkDimSize * ChunkDimSize
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
		hm:    ebiten.NewImage(ChunkDimSize, ChunkDimSize),
	}
}

func (c *Chunk) Id() ChunkId {
	return c.id
}

// WritePixels writes the given pixels to the heightmap if the generation matches
func (c *Chunk) WritePixels(gen uint64, pixels []byte) bool {
	if c.gen != gen {
		return false
	}

	c.hm.WritePixels(pixels)

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

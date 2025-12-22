package world

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Chunk struct {
	id       ChunkId
	dirty    bool
	layers   map[int]*ebiten.Image
	lastUsed uint64
}

func NewChunk(id ChunkId) *Chunk {
	return &Chunk{
		id:     id,
		dirty:  true,
		layers: make(map[int]*ebiten.Image),
	}
}

func (c *Chunk) Id() ChunkId {
	return c.id
}

func (c *Chunk) SetLayer(l int, img *ebiten.Image) {
	c.layers[l] = img
}

func (c *Chunk) GetLayer(l int) *ebiten.Image {
	return c.layers[l]
}

func (c *Chunk) MarkDirty() {
	c.dirty = true
}

func (c *Chunk) ClearDirty() {
	c.dirty = false
}

func (c *Chunk) IsDirty() bool {
	return c.dirty
}

func (c *Chunk) SetLastUsed(frame uint64) {
	c.lastUsed = frame
}

func (c *Chunk) LastUsed() uint64 {
	return c.lastUsed
}

package world

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ChunkStateDirty = iota
	ChunkStateQueued
	ChunkStateGenerating
	ChunkStateReady
	ChunkStateEvicted
)

const (
	LayerHeightMap = iota
)

type Chunk struct {
	id       ChunkId
	state    int
	layers   map[int]*ebiten.Image
	lastUsed uint64
	mutex    sync.Mutex
}

func NewChunk(id ChunkId) *Chunk {
	return &Chunk{
		id:     id,
		state:  ChunkStateDirty,
		layers: make(map[int]*ebiten.Image),
	}
}

func (c *Chunk) Id() ChunkId {
	return c.id
}

func (c *Chunk) SetLayer(l int, img *ebiten.Image) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.layers[l] = img
}

func (c *Chunk) GetLayer(l int) *ebiten.Image {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.layers[l]
}

// SetState sets the state of the chunk in a thread-safe manner.
func (c *Chunk) SetState(state int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.state = state
}

// Is checks if the chunk is in the given state in a thread-safe manner.
func (c *Chunk) Is(state int) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.state == state
}

func (c *Chunk) SetLastUsed(frame uint64) {
	c.lastUsed = frame
}

func (c *Chunk) LastUsed() uint64 {
	return c.lastUsed
}

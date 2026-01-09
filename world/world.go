package world

import (
	"geoforge/geo"
	"geoforge/noise"
	"runtime"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	margin float64
	chunks map[ChunkId]*Chunk
	queue  chan *Chunk

	noise   noise.Noise
	noiseMu sync.RWMutex
}

func NewWorld(margin int) *World {
	ws := runtime.NumCPU()
	w := &World{
		margin: float64(margin) * ChunkSize,
		chunks: make(map[ChunkId]*Chunk),
		queue:  make(chan *Chunk, ws*2),
	}

	for i := 0; i < ws; i++ {
		go w.worker(w.queue)
	}

	return w
}

func (w *World) Update(rect geo.Rect) {
	rect = rect.Expand(w.margin).SnapOut(ChunkSize)

	for y := rect.MinY; y < rect.MaxY; y += ChunkSize {
		for x := rect.MinX; x < rect.MaxX; x += ChunkSize {
			id := NewChunkId(int(x/ChunkSize), int(y/ChunkSize))
			w.ensure(id)
		}
	}

	w.evict(rect)
	w.populate()
}

func (w *World) Chunks() map[ChunkId]*Chunk {
	return w.chunks
}

func (w *World) MarkDirty() {
	for _, c := range w.chunks {
		c.SetState(ChunkStateDirty)
	}
}

func (w *World) populate() {
	if w.Noise() == nil {
		return
	}

	for _, c := range w.chunks {
		if c.Is(ChunkStateDirty) {
			select {
			case w.queue <- c:
				c.SetState(ChunkStateQueued)
			default:
				// queue is full, give up for now
				return
			}
		}
	}
}

func (w *World) ensure(id ChunkId) {
	c, exists := w.chunks[id]
	if !exists {
		c = NewChunk(id)
		w.chunks[id] = c
	}
}

func (w *World) evict(rect geo.Rect) {
	for id, c := range w.chunks {
		cx := float64(id.X) * ChunkSize
		cy := float64(id.Y) * ChunkSize

		cRect := geo.NewRect(cx, cy, cx+ChunkSize, cy+ChunkSize)
		if !rect.Intersects(cRect) {
			c.SetState(ChunkStateEvicted)
			delete(w.chunks, id)
		}
	}
}

func (w *World) worker(queue chan *Chunk) {
	hm := make([]float32, ChunkSurface) // heightmap raw
	hmp := make([]byte, 4*ChunkSurface) // heightmap RGBA

	for c := range queue {
		if !c.Is(ChunkStateQueued) {
			continue
		}

		baseX := c.Id().X*ChunkSize - ChunkApron
		baseY := c.Id().Y*ChunkSize - ChunkApron

		ns := w.Noise()
		if ns == nil {
			continue
		}

		c.SetState(ChunkStateGenerating)
		ns.Fill(hm, ChunkDimSize, float32(baseX), float32(baseY))

		for i := range hm {
			n := hm[i]
			n = (n + 1) / 2 // normalize to 0..1
			v := byte(n * 255)
			hmp[i*4+0] = v
			hmp[i*4+1] = v
			hmp[i*4+2] = v
			hmp[i*4+3] = 255
		}

		hmImg := c.GetLayer(LayerHeightMap)
		if hmImg == nil || hmImg.Bounds().Dx() != ChunkDimSize || hmImg.Bounds().Dy() != ChunkDimSize {
			hmImg = ebiten.NewImage(ChunkDimSize, ChunkDimSize)
			c.SetLayer(LayerHeightMap, hmImg)
		}
		hmImg.WritePixels(hmp)

		c.SetState(ChunkStateReady)
	}
}

func (w *World) SetNoise(n noise.Noise) {
	w.noiseMu.Lock()
	w.noise = n
	w.noiseMu.Unlock()

	w.MarkDirty()
}

func (w *World) Noise() noise.Noise {
	w.noiseMu.RLock()
	defer w.noiseMu.RUnlock()

	return w.noise
}

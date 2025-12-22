package world

import (
	"geoforge/geo"
	"geoforge/noise"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	chunkSize float64
	margin    float64
	chunks    map[ChunkId]*Chunk
	maxChunks int
	frame     uint64
	noise     noise.Noise
	queue     chan *Chunk
}

func NewWorld(chunkSize float64, margin, maxChunks int, noise noise.Noise) *World {
	w := &World{
		chunkSize: chunkSize,
		margin:    float64(margin) * chunkSize,
		maxChunks: maxChunks,
		chunks:    make(map[ChunkId]*Chunk),
		noise:     noise,
		queue:     make(chan *Chunk, maxChunks),
	}

	ws := runtime.NumCPU()
	if ws > 4 {
		ws = ws / 2
	} else if ws < 1 {
		ws = 1
	}

	for i := 0; i < ws; i++ {
		go w.worker(w.queue)
	}

	return w
}

func (w *World) Update(rect geo.Rect) {
	w.frame++

	rect = rect.Expand(w.margin).SnapOut(w.chunkSize)

	for y := rect.MinY; y < rect.MaxY; y += w.chunkSize {
		for x := rect.MinX; x < rect.MaxX; x += w.chunkSize {
			id := NewChunkId(int64(x/w.chunkSize), int64(y/w.chunkSize))
			w.ensure(id)
		}
	}

	w.evict()
	w.populate()
}

func (w *World) Chunks() map[ChunkId]*Chunk {
	return w.chunks
}

func (w *World) ChunkSize() float64 {
	return w.chunkSize
}

func (w *World) populate() {
	for _, c := range w.chunks {
		if c.Is(ChunkStateDirty) {
			select {
			case w.queue <- c:
				c.SetState(ChunkStateQueued)
			default:
				// queue is full
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

	c.SetLastUsed(w.frame)
}

func (w *World) evict() {
	for len(w.chunks) > w.maxChunks {
		var oldestID ChunkId
		var oldest *Chunk

		for id, ch := range w.chunks {
			if oldest == nil || ch.LastUsed() < oldest.LastUsed() {
				oldest = ch
				oldestID = id
			}
		}

		if oldest == nil {
			return
		}
		oldest.SetState(ChunkStateEvicted)
		delete(w.chunks, oldestID)
	}
}

func (w *World) worker(queue chan *Chunk) {
	for c := range queue {
		if !c.Is(ChunkStateQueued) {
			continue
		}

		c.SetState(ChunkStateGenerating)
		hm := make([]byte, int(4*w.chunkSize*w.chunkSize))
		for y := 0; y < int(w.chunkSize); y++ {
			row := y * int(w.chunkSize) * 4
			for x := 0; x < int(w.chunkSize); x++ {
				idx := row + x*4
				wx := float64(c.Id().X)*w.chunkSize + float64(x)
				wy := float64(c.Id().Y)*w.chunkSize + float64(y)
				v := byte(w.noise.Eval(wx*.5, wy*.5) * 255)
				hm[idx], hm[idx+1], hm[idx+2] = v, v, v
				hm[idx+3] = 255
			}
		}

		hmImg := c.GetLayer(LayerHeightMap)
		if hmImg == nil {
			hmImg = ebiten.NewImage(int(w.chunkSize), int(w.chunkSize))
			c.SetLayer(LayerHeightMap, hmImg)
		}
		hmImg.WritePixels(hm)

		c.SetState(ChunkStateReady)
	}
}

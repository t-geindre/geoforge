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
	apron     float64
}

func NewWorld(chunkSize, apron float64, margin, maxChunks int, noise noise.Noise) *World {
	ws := runtime.NumCPU()
	if ws > 4 {
		ws = ws / 2
	} else if ws < 1 {
		ws = 1
	}

	w := &World{
		chunkSize: chunkSize,
		margin:    float64(margin) * chunkSize,
		maxChunks: maxChunks,
		chunks:    make(map[ChunkId]*Chunk),
		noise:     noise,
		queue:     make(chan *Chunk, ws*2),
		apron:     apron,
	}

	for i := 0; i < ws; i++ {
		go w.worker(w.queue)
	}

	return w
}

func (w *World) Update(rect geo.Rect) {
	if w.noise.Params().HasChanged() {
		w.MarkDirty()
	}

	w.frame++

	rect = rect.Expand(w.margin).SnapOut(w.chunkSize)

	for y := rect.MinY; y < rect.MaxY; y += w.chunkSize {
		for x := rect.MinX; x < rect.MaxX; x += w.chunkSize {
			id := NewChunkId(int(x/w.chunkSize), int(y/w.chunkSize))
			w.ensure(id)
		}
	}

	w.evict(rect)
	w.populate()
}

func (w *World) Chunks() map[ChunkId]*Chunk {
	return w.chunks
}

func (w *World) ChunkSize() float64 {
	return w.chunkSize
}

func (w *World) Apron() float64 {
	return w.apron
}

func (w *World) MarkDirty() {
	for _, c := range w.chunks {
		c.SetState(ChunkStateDirty)
	}
}

func (w *World) populate() {
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

	c.SetLastUsed(w.frame)
}

func (w *World) evict(rect geo.Rect) {
	for id, c := range w.chunks {
		cx := float64(id.X) * w.chunkSize
		cy := float64(id.Y) * w.chunkSize

		cRect := geo.NewRect(cx, cy, cx+w.chunkSize, cy+w.chunkSize)
		if !rect.Intersects(cRect) {
			c.SetState(ChunkStateEvicted)
			delete(w.chunks, id)
		}
	}
}

func (w *World) worker(queue chan *Chunk) {
	N := int(w.chunkSize)
	A := int(w.apron)
	W := N + 2*A // texture width/height with apron

	hm := make([]float32, W*W) // heightmap raw
	hmp := make([]byte, 4*W*W) // heightmap RGBA

	for c := range queue {
		if !c.Is(ChunkStateQueued) {
			continue
		}

		c.SetState(ChunkStateGenerating)

		baseX := c.Id().X*N - A
		baseY := c.Id().Y*N - A
		w.noise.Fill(hm, W, float32(baseX), float32(baseY))

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
		if hmImg == nil || hmImg.Bounds().Dx() != W || hmImg.Bounds().Dy() != W {
			hmImg = ebiten.NewImage(W, W)
			c.SetLayer(LayerHeightMap, hmImg)
		}
		hmImg.WritePixels(hmp)

		c.SetState(ChunkStateReady)
	}
}

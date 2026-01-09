package world

import (
	"geoforge/geo"
	"geoforge/noise"
	"runtime"
	"sync"
)

type query struct {
	id  ChunkId
	gen uint64 // chunk generation, avoid stale
}

type result struct {
	id  ChunkId
	gen uint64 // chunk generation, avoid stale
	hm  []byte // heightmap, RGBA, R = height
}

type World struct {
	margin float64
	chunks map[ChunkId]*Chunk

	query   chan query
	results chan result
	hmPool  sync.Pool

	noise   noise.Noise
	noiseMu sync.RWMutex
}

func NewWorld(margin int) *World {
	ws := runtime.NumCPU()

	w := &World{
		margin:  float64(margin) * ChunkSize,
		chunks:  make(map[ChunkId]*Chunk),
		query:   make(chan query, ws*2),
		results: make(chan result, ws*2),
		hmPool: sync.Pool{
			New: func() any {
				return make([]byte, 4*ChunkSurface)
			},
		},
	}

	for i := 0; i < ws; i++ {
		go w.worker()
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
	w.generateHeightMaps()
	w.storeHeightMaps()
}

func (w *World) Chunks() map[ChunkId]*Chunk {
	return w.chunks
}

func (w *World) MarkDirty() {
	for _, c := range w.chunks {
		c.BumpGeneration()
		c.SetState(ChunkStateDirty)
	}
}

func (w *World) generateHeightMaps() {
	if w.Noise() == nil {
		return
	}

	for _, c := range w.chunks {
		if c.Is(ChunkStateDirty) {
			select {
			case w.query <- query{
				id:  c.Id(),
				gen: c.GetGeneration(),
			}:
				c.SetState(ChunkStateQueued)
			default:
				// query is full, give up for now
				return
			}
		}
	}
}

func (w *World) storeHeightMaps() {
	for {
		select {
		case res := <-w.results:
			c, exists := w.chunks[res.id]
			if exists {
				if c.WritePixels(res.gen, res.hm) {
					c.SetState(ChunkStateReady)
				}
			}
			w.hmPool.Put(res.hm)
		default:
			return
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
	for id, _ := range w.chunks {
		cx := float64(id.X) * ChunkSize
		cy := float64(id.Y) * ChunkSize

		cRect := geo.NewRect(cx, cy, cx+ChunkSize, cy+ChunkSize)
		if !rect.Intersects(cRect) {
			delete(w.chunks, id)
		}
	}
}

func (w *World) worker() {
	hm := make([]float32, ChunkSurface)

	for q := range w.query {
		baseX := q.id.X*ChunkSize - ChunkApron
		baseY := q.id.Y*ChunkSize - ChunkApron

		ns := w.Noise()
		if ns == nil {
			continue
		}

		ns.Fill(hm, ChunkDimSize, float32(baseX), float32(baseY))

		hmp := w.hmPool.Get().([]byte)
		for i := range hm {
			n := hm[i]
			n = (n + 1) / 2 // normalize to 0..1
			v := byte(n * 255)
			hmp[i*4+0] = v
			hmp[i*4+1] = v
			hmp[i*4+2] = v
			hmp[i*4+3] = 255
		}

		w.results <- result{
			id:  q.id,
			gen: q.gen,
			hm:  hmp,
		}
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

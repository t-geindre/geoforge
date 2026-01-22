package world

import (
	"geoforge/camera"
	"geoforge/game"
	"geoforge/geo"
	"geoforge/noise"
	"runtime"
	"sync"
)

type query struct {
	id    ChunkId
	gen   uint64 // chunk generation, avoid stale
	noise noise.Noise
}

type result struct {
	id  ChunkId
	gen uint64 // chunk generation, avoid stale
	hm  []byte // heightmap, grayscale, 1 byte per pixel
}

type World struct {
	margin float64
	chunks map[ChunkId]*Chunk

	query    chan query
	results  chan result
	hmPool   sync.Pool
	rgbaPool sync.Pool

	noise noise.Noise

	cam   camera.Camera
	camSt uint8

	game.StateChanged
}

func NewWorld(margin int, cam camera.Camera) *World {
	ws := runtime.NumCPU()

	w := &World{
		margin:       float64(margin) * ChunkSize,
		chunks:       make(map[ChunkId]*Chunk),
		query:        make(chan query, ws*2),
		results:      make(chan result, ws*2),
		cam:          cam,
		camSt:        cam.RegisterChangeId(),
		StateChanged: game.NewStateChanged(),
		hmPool: sync.Pool{
			New: func() any {
				return make([]byte, ChunkSurface)
			},
		},
		rgbaPool: sync.Pool{
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

func (w *World) Update() {
	if w.cam.HasChanged(w.camSt) {
		rect := w.cam.WorldRect().Expand(w.margin).SnapOut(ChunkSize)

		for y := rect.MinY; y < rect.MaxY; y += ChunkSize {
			for x := rect.MinX; x < rect.MaxX; x += ChunkSize {
				id := NewChunkId(int(x/ChunkSize), int(y/ChunkSize))
				w.ensure(id)
			}
		}

		w.evict(rect)
	}

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
	w.SetChanged()
}

func (w *World) generateHeightMaps() {
	if w.noise == nil {
		return
	}

	for _, c := range w.chunks {
		if c.Is(ChunkStateDirty) {
			select {
			case w.query <- query{
				id:    c.Id(),
				gen:   c.GetGeneration(),
				noise: w.noise,
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
				rgbaBuf := w.rgbaPool.Get().([]byte)
				if c.WritePixels(res.gen, res.hm, rgbaBuf) {
					c.SetState(ChunkStateReady)
					w.SetChanged()
				}
				w.rgbaPool.Put(rgbaBuf)
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
			w.SetChanged()
		}
	}
}

func (w *World) worker() {
	const normalizeFactor = 127.5 // (v + 1) * 127.5 = (v + 1) / 2 * 255
	const offset = 127.5

	for q := range w.query {
		baseX := q.id.X * ChunkSize
		baseY := q.id.Y * ChunkSize

		hmp := w.hmPool.Get().([]byte)

		idx := 0
		for y := 0; y < ChunkSize; y++ {
			for x := 0; x < ChunkSize; x++ {
				v := q.noise.At(float32(baseX+x), float32(baseY+y))
				hmp[idx] = byte(v*normalizeFactor + offset) // 0..255, grayscale
				idx++
			}
		}

		w.results <- result{
			id:  q.id,
			gen: q.gen,
			hm:  hmp,
		}
	}
}

func (w *World) SetNoise(n noise.Noise) {
	w.noise = n
	w.MarkDirty()
}

func (w *World) Close() {
	close(w.query)
}

package world

type World struct {
	chunkSize float64
	margin    float64
	chunks    map[ChunkId]*Chunk
	maxChunks int
	frame     uint64
}

func NewWorld(chunkSize float64, margin, maxChunks int) *World {
	return &World{
		chunkSize: chunkSize,
		margin:    float64(margin) * chunkSize,
		maxChunks: maxChunks,
		chunks:    make(map[ChunkId]*Chunk),
	}
}

func (w *World) Update(cam Camera) {
	w.frame++

	r := cam.WorldRect().
		Expand(w.margin).
		SnapOut(w.chunkSize)

	for y := r.MinY; y < r.MaxY; y += w.chunkSize {
		for x := r.MinX; x < r.MaxX; x += w.chunkSize {
			id := NewChunkId(int64(x/w.chunkSize), int64(y/w.chunkSize))
			w.ensure(id)
		}
	}

	w.evict()
}

func (w *World) Chunks() map[ChunkId]*Chunk {
	return w.chunks
}

func (w *World) ChunkSize() float64 {
	return w.chunkSize
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
		delete(w.chunks, oldestID)
	}
}

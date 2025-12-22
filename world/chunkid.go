package world

type ChunkId struct {
	X, Y int64
}

func NewChunkId(x, y int64) ChunkId {
	return ChunkId{X: x, Y: y}
}

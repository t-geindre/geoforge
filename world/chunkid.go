package world

import "fmt"

type ChunkId struct {
	X, Y int64
}

func NewChunkId(x, y int64) ChunkId {
	return ChunkId{X: x, Y: y}
}

func (id ChunkId) String() string {
	return fmt.Sprintf("(%d, %d)", id.X, id.Y)
}

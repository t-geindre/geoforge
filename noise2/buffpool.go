package noise2

import "sync"

type BuffPool struct {
	pool sync.Pool
}

func NewBuffPool() *BuffPool {
	return &BuffPool{
		pool: sync.Pool{
			New: func() any { return make([]float32, N) },
		},
	}
}

func (p *BuffPool) Get() []float32  { return p.pool.Get().([]float32) }
func (p *BuffPool) Put(b []float32) { p.pool.Put(b) }

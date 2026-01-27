package noise2

import (
	"fmt"
	"math"
)

type OpKind uint8

const (
	OpSine OpKind = iota
	OpClamp11
	OpAdd
	OpAddConst
)

type Op struct {
	Kind       OpKind
	Src0, Src1 int
	Dst        int
	F0         float32 // param (freq, etc.)
}

type Node interface {
	Op() OpKind
	Inputs() []Node
}

type SineNode struct{ Freq float32 }

func (n *SineNode) Op() OpKind     { return OpSine }
func (n *SineNode) Inputs() []Node { return nil }

type AddNode struct{ A, B Node }

func (n *AddNode) Op() OpKind     { return OpAdd }
func (n *AddNode) Inputs() []Node { return []Node{n.A, n.B} }

type AddConstNode struct {
	A Node
	C float32
}

func (n *AddConstNode) Op() OpKind     { return OpAddConst }
func (n *AddConstNode) Inputs() []Node { return []Node{n.A} }

type Clamp11Node struct{ A Node }

func (n *Clamp11Node) Op() OpKind     { return OpClamp11 }
func (n *Clamp11Node) Inputs() []Node { return []Node{n.A} }

const (
	ChunkSize = 256
	N         = ChunkSize * ChunkSize
)

type Grid struct {
	Size   int
	X0, Y0 float32
	Step   float32
}

func opSine(out []float32, g Grid, freq float32) {
	size := g.Size
	k := 0
	y := g.Y0
	for j := 0; j < size; j++ {
		x := g.X0
		sy := float32(math.Sin(float64(y * freq)))
		for i := 0; i < size; i++ {
			sx := float32(math.Sin(float64(x * freq)))
			out[k] = sx * sy
			k++
			x += g.Step
		}
		y += g.Step
	}
}

func opClamp11(out, in []float32) {
	for i := 0; i < N; i++ {
		if in[i] < -1 {
			out[i] = -1
		} else if in[i] > 1 {
			out[i] = 1
		} else {
			out[i] = in[i]
		}
	}
}

func opAdd(out, a, b []float32) {
	for i := 0; i < N; i++ {
		out[i] = a[i] + b[i]
	}
}

func opAddConst(out, in []float32, c float32) {
	for i := 0; i < N; i++ {
		out[i] = in[i] + c
	}
}

type Program struct {
	Ops      []Op
	NumBuffs int
}

func Compile(root Node) Program {
	var ops []Op
	buffOf := map[Node]int{} // node -> buf id
	nextBuff := 1

	alloc := func() int {
		b := nextBuff
		nextBuff++
		return b
	}

	var emit func(n Node) int
	emit = func(n Node) int {
		if n == nil {
			panic("nil node")
		}
		if b, ok := buffOf[n]; ok {
			return b // shared node (branches) -> compute once
		}

		ins := n.Inputs()
		src0, src1 := -1, -1
		if len(ins) >= 1 {
			src0 = emit(ins[0])
		}
		if len(ins) >= 2 {
			src1 = emit(ins[1])
		}

		dst := 0
		if n != root {
			dst = alloc() // todo optimize: in-place for some ops when possible
		}

		buffOf[n] = dst

		op := Op{Kind: n.Op(), Src0: src0, Src1: src1, Dst: dst}

		// Attach parameters for parametric nodes
		switch t := n.(type) {
		case *SineNode:
			op.F0 = t.Freq
		case *AddConstNode:
			op.F0 = t.C
		}

		ops = append(ops, op)
		return dst
	}

	out := emit(root)
	if out != 0 {
		panic("internal: root must compile to buff0")
	}

	return Program{
		Ops:      ops,
		NumBuffs: nextBuff,
	}
}

func UseCounts(p Program) []int {
	uc := make([]int, p.NumBuffs)
	for _, op := range p.Ops {
		switch op.Kind {
		case OpSine:
			// no inputs

		case OpClamp11, OpAddConst:
			if op.Src0 >= 0 {
				uc[op.Src0]++
			}

		case OpAdd:
			if op.Src0 >= 0 {
				uc[op.Src0]++
			}
			if op.Src1 >= 0 {
				uc[op.Src1]++
			}

		default:
			panic("unknown op kind")
		}
	}
	return uc
}

func RunInto(g Grid, p Program, out0 []float32, pool *BuffPool) {
	if len(out0) != N {
		panic("out0 must be N")
	}

	uc := UseCounts(p) // buff use counts

	// Allocate buffers
	buffs := make([][]float32, p.NumBuffs)
	buffs[0] = out0

	ensure := func(id int) {
		if id <= 0 {
			return // buf0 is provided; ignore -1
		}
		if buffs[id] == nil {
			b := pool.Get()
			if len(b) != N {
				panic("pool returned wrong buffer size")
			}
			buffs[id] = b
		}
	}

	release := func(id int) {
		if id <= 0 {
			return // never release buf0; ignore -1
		}
		uc[id]--
		if uc[id] == 0 {
			if buffs[id] == nil {
				panic("internal: releasing nil buffer")
			}
			pool.Put(buffs[id])
			buffs[id] = nil
		}
	}

	// Run ops
	for _, op := range p.Ops {
		ensure(op.Dst)

		switch op.Kind {
		case OpSine:
			opSine(buffs[op.Dst], g, op.F0)

		case OpClamp11:
			opClamp11(buffs[op.Dst], buffs[op.Src0])
			release(op.Src0)

		case OpAdd:
			opAdd(buffs[op.Dst], buffs[op.Src0], buffs[op.Src1])
			release(op.Src0)
			release(op.Src1)

		case OpAddConst:
			opAddConst(buffs[op.Dst], buffs[op.Src0], op.F0)
			release(op.Src0)

		default:
			panic("unknown op kind")
		}
	}

	// sanity check / debug / optional
	for i := 1; i < len(buffs); i++ {
		if buffs[i] != nil {
			panic("leak temp buffer")
		}
	}
}

func Demo() { // for testing
	stats := func(a []float32) (min, max float32) {
		min, max = a[0], a[0]
		for i := 1; i < len(a); i++ {
			v := a[i]
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
		return
	}

	var graph Node
	graph = &Clamp11Node{
		A: &AddNode{
			A: &AddConstNode{
				A: &SineNode{Freq: 0.1},
				C: 0.5,
			},
			B: &SineNode{Freq: 0.2},
		},
	}

	//graph = &SineNode{Freq: 10}

	buffPool := NewBuffPool()
	prog := Compile(graph)
	out := make([]float32, N)
	RunInto(Grid{Size: ChunkSize, X0: 0, Y0: 0, Step: 1}, prog, out, buffPool)

	fmt.Printf("ops=%d buffs=%d\n", len(prog.Ops), prog.NumBuffs)
	opsLbl := map[OpKind]string{
		OpSine:     "Sine",
		OpClamp11:  "Clamp11",
		OpAdd:      "Add",
		OpAddConst: "AddConst",
	}
	for i, op := range prog.Ops {
		fmt.Printf("%02d kind=%s src0=%d src1=%d dst=%d f0=%.3f\n",
			i, opsLbl[op.Kind], op.Src0, op.Src1, op.Dst, op.F0)
	}

	center := 128*ChunkSize + 128
	mn, mx := stats(out)
	fmt.Printf("out: min=%.10f max=%.10f out[0]=%.10f out[center]=%.10f out[last]=%.10f\n",
		mn, mx, out[0], out[center], out[N-1])
}

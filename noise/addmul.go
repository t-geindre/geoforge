package noise

type Add struct{ a, b Noise }

func NewAdd(a, b Noise) Noise            { return &Add{a, b} }
func (n *Add) Eval(x, y float64) float64 { return n.a.Eval(x, y) + n.b.Eval(x, y) }

type Mul struct{ a, b Noise }

func NewMul(a, b Noise) Noise            { return &Mul{a, b} }
func (n *Mul) Eval(x, y float64) float64 { return n.a.Eval(x, y) * n.b.Eval(x, y) }

type AddConst struct {
	inner Noise
	k     float64
}

func NewAddConst(inner Noise, k float64) Noise { return &AddConst{inner, k} }
func (n *AddConst) Eval(x, y float64) float64  { return n.inner.Eval(x, y) + n.k }

type MulConst struct {
	inner Noise
	k     float64
}

func NewMulConst(inner Noise, k float64) Noise { return &MulConst{inner, k} }
func (n *MulConst) Eval(x, y float64) float64  { return n.inner.Eval(x, y) * n.k }

package noise

type Warp struct {
	inner  Noise
	warpX  Noise // idéalement [-1..1]
	warpY  Noise // idéalement [-1..1]
	amount float64
}

func NewWarp(inner, warpX, warpY Noise, amount float64) Noise {
	return &Warp{inner: inner, warpX: warpX, warpY: warpY, amount: amount}
}

func (n *Warp) Eval(x, y float64) float64 {
	dx := n.warpX.Eval(x, y) * n.amount
	dy := n.warpY.Eval(x, y) * n.amount
	return n.inner.Eval(x+dx, y+dy)
}

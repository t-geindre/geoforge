package preset

type Numeric interface {
	~int | ~float32
}

type Variable[T Numeric] interface {
	Param[T]
	Min() T
	Max() T
	Step() T
	Digits() int
}

type variable[T Numeric] struct {
	Param[T]
	min, max, step T
	digits         int
}

func NewVariable[T Numeric](id ParamId, label string, val, min, max, step T, digits int, onChange func(T)) Variable[T] {
	return &variable[T]{
		Param:  NewParam[T](id, label, val, onChange),
		min:    min,
		max:    max,
		step:   step,
		digits: digits,
	}
}

func (v *variable[T]) Min() T {
	return v.min
}

func (v *variable[T]) Max() T {
	return v.max
}

func (v *variable[T]) Step() T {
	return v.step
}

func (v *variable[T]) Digits() int {
	return v.digits
}

package preset

type Option[T Numeric] struct {
	val   T
	label string
}

func NewOption[T Numeric](val T, label string) Option[T] {
	return Option[T]{val: val, label: label}
}

func (o Option[T]) Val() T {
	return o.val
}

func (o Option[T]) Label() string {
	return o.label
}

type Choice[T Numeric] interface {
	Param[T]
	Options() []Option[T]
}

type choice[T Numeric] struct {
	Param[T]
	options []Option[T]
}

func NewChoice[T Numeric](id ParamId, label string, val T, opts []Option[T], onChange func(T)) Choice[T] {
	c := &choice[T]{
		Param:   NewParam[T](id, label, val, onChange),
		options: opts,
	}
	return c
}

func (c *choice[T]) Options() []Option[T] {
	return c.options
}

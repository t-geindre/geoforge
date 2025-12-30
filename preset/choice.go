package preset

type ChoiceGeneric interface {
	ParamGeneric
	OptionsLabels() []string
	SetValByIndex(idx int)
	ValIndex() int
}

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

func (c *choice[T]) OptionsLabels() []string {
	labels := make([]string, len(c.options))
	for i, opt := range c.options {
		labels[i] = opt.Label()
	}
	return labels
}

func (c *choice[T]) SetValByIndex(idx int) {
	if idx < 0 || idx >= len(c.options) {
		return
	}
	c.SetVal(c.options[idx].Val())
}

func (c *choice[T]) ValIndex() int {
	val := c.Val()
	for i, opt := range c.options {
		if opt.Val() == val {
			return i
		}
	}
	return -1
}

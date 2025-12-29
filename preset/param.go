package preset

type ParamId int

type ParamGeneric interface {
	Id() ParamId
	Label() string
	HasChanged() bool
}

type Param[T comparable] interface {
	ParamGeneric
	Val() T
	SetVal(v T)
}

type param[T comparable] struct {
	id         ParamId
	label      string
	val        T
	hasChanged bool
	onChange   func(T)
}

func NewParam[T comparable](id ParamId, label string, val T, onChange func(T)) Param[T] {
	if onChange == nil {
		onChange = func(T) {}
	}

	onChange(val)

	return &param[T]{
		id:       id,
		label:    label,
		val:      val,
		onChange: onChange,
	}
}

func (p *param[T]) Id() ParamId {
	return p.id
}

func (p *param[T]) Label() string {
	return p.label
}

func (p *param[T]) Val() T {
	return p.val
}

func (p *param[T]) SetVal(v T) {
	if p.val == v {
		return
	}

	p.hasChanged = true
	p.val = v
	p.onChange(p.Val())
}

// HasChanged since last call
func (p *param[T]) HasChanged() bool {
	defer func() { p.hasChanged = false }()
	return p.hasChanged
}

package ui

type Metric interface {
	Label() string
	Value() float32
}

type metric struct {
	label string
	value func() float32
}

func NewMetric(label string, value func() float32) Metric {
	if value == nil {
		value = func() float32 { return 0 }
	}
	return &metric{
		label: label,
		value: value,
	}
}

func (m *metric) Label() string {
	return m.label
}

func (m *metric) Value() float32 {
	return m.value()
}

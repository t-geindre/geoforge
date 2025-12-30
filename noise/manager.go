package noise

import "geoforge/preset"

type Receiver interface {
	SetNoise(n Noise)
	MarkDirty()
}

type Manager struct {
	noises   []Noise
	receiver Receiver
	params   preset.ParamSet
}

func NewNoiseManager(r Receiver) *Manager {
	m := &Manager{
		noises:   []Noise{},
		receiver: r,
		params:   preset.NewAnonymousParamSet(),
	}

	m.AddNoise(NewNoise())
	m.AddNoise(NewNoise())

	return m
}

func (m *Manager) AddNoise(n Noise) {
	// Label
	n.Params().SetLabel("Unnamed")
	n.Params().Prepend(preset.NewParam(0, "Name", n.Params().Label(), func(s string) {
		n.Params().SetLabel(s)
	}))

	// Rendering
	preset.NewParam(1000, "Render", false, func(b bool) {
		if b {
			others := m.params.QueryParamById(1000)
			for _, o := range others {
				if o != n.Params() {
					o.(preset.Param[bool]).SetVal(false)
				}
			}
			m.receiver.SetNoise(n)
		}
	})
	n.Params().Prepend()

	// Store
	m.noises = append(m.noises, n)
	m.receiver.SetNoise(n)
	m.params.Append(n.Params())
}

func (m *Manager) Params() preset.ParamSet {
	return m.params
}

func (m *Manager) Update() {
	if m.params.HasChanged() {
		m.receiver.MarkDirty()
	}
}

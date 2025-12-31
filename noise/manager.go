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
	m.AddNoise(NewNoise())
	m.AddNoise(NewNoise())

	return m
}

func (m *Manager) AddNoise(n Noise) {
	// Rendering
	isRendered := false
	if len(m.noises) == 0 {
		m.receiver.SetNoise(n)
		isRendered = true
	}

	n.Params().Prepend(preset.NewParam(1000, "Render", isRendered, func(p preset.Param[bool]) {
		if p.Val() {
			others := m.params.QueryParamById(1000)
			for _, o := range others {
				if o != p {
					o.(preset.Param[bool]).SetVal(false)
				}
			}
			m.receiver.SetNoise(n)
			return
		}

		// Prevent disabling all noises
		others := m.params.QueryParamById(1000)
		anyRendered := false
		for _, o := range others {
			if o != p && o.(preset.Param[bool]).Val() {
				anyRendered = true
				break
			}
		}
		if !anyRendered {
			p.SetVal(true)
		}
	}))

	// Label
	n.Params().SetLabel("Unnamed")
	n.Params().Prepend(preset.NewParam(0, "Name", n.Params().Label(), func(p preset.Param[string]) {
		n.Params().SetLabel(p.Val())
	}))

	// Store
	m.noises = append(m.noises, n)
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

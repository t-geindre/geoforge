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
	}

	m.params = preset.NewAnonymousParamSet()
	m.params.Append(preset.NewAction(0, "Add Noise", func() {
		m.AddNoise(NewNoise())
	}))

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

	// Add remove action
	n.Params().Append(preset.NewAction(1, "Remove Noise", func() {
		m.RemoveNoise(n)
	}))

	// Store
	m.noises = append(m.noises, n)
	m.params.Append(n.Params())
}

func (m *Manager) RemoveNoise(n Noise) {
	// Remove from list
	newNoises := []Noise{}
	for _, noise := range m.noises {
		if noise != n {
			newNoises = append(newNoises, noise)
		}
	}
	m.noises = newNoises

	// Remove params
	m.params.Remove(n.Params())

	// If rendered, disable
	rendered := n.Params().QueryParamById(1000)[0].(preset.Param[bool]).Val()
	if rendered {
		m.receiver.SetNoise(nil)
	}
}

func (m *Manager) Params() preset.ParamSet {
	return m.params
}

func (m *Manager) Update() {
	if m.params.HasChanged() {
		m.receiver.MarkDirty()
	}
}

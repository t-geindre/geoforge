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
	m.params.Append(preset.NewAction(ParamActionAdd, "Add Noise", func(preset.Action) {
		m.AddNoise(NewMultiNoise(
			NewFastNoise(),
			NewSine(),
		))
	}))

	return m
}

func (m *Manager) AddNoise(n Noise) {
	first := true
	n.Params().Prepend(preset.NewParam(ParamIsRendered, "Render", len(m.noises) == 0, func(p preset.Param[bool]) {
		if p.Val() {
			for _, op := range m.params.QueryParamById(ParamIsRendered) {
				if op != p {
					op.(preset.Param[bool]).SetVal(false)
				}
			}
			m.receiver.SetNoise(n)
			return
		}
		// Avoid disabling the noise when first added
		if first {
			first = false
			return
		}

		m.receiver.SetNoise(nil)
	}))

	// Label
	n.Params().SetLabel("Unnamed")
	n.Params().Prepend(preset.NewParam(ParamName, "Name", n.Params().Label(), func(p preset.Param[string]) {
		if p.Val() == "" {
			p.SetVal("Unnamed")
			return
		}

		n.Params().SetLabel(p.Val())
	}))

	// Add remove action
	n.Params().Append(preset.NewAction(ParamActionRemove, "Remove Noise", func(preset.Action) {
		m.RemoveNoise(n)
	}))

	// Store
	m.noises = append(m.noises, n)
	m.params.Append(n.Params())
}

func (m *Manager) RemoveNoise(n Noise) {
	// Remove from list
	var newNoises []Noise
	for _, ns := range m.noises {
		if ns != n {
			newNoises = append(newNoises, ns)
		}
	}
	m.noises = newNoises

	// Remove params
	m.params.Remove(n.Params())

	// If rendered, disable
	if n.Params().QueryParamById(ParamIsRendered)[0].(preset.Param[bool]).Val() {
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

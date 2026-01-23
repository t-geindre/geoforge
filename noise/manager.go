package noise

import (
	"fmt"
	"geoforge/preset"
)

type Receiver interface {
	SetNoise(n Noise)
	MarkDirty()
}

const defaultName = "Unnamed"
const noiseNone = -1

type Manager struct {
	noises   []Noise
	receiver Receiver
	params   preset.ParamSet
}

func NewNoiseManager(r Receiver) *Manager {
	m := &Manager{
		noises:   make([]Noise, 0),
		receiver: r,
	}

	m.params = preset.NewAnonymousParamSet()
	m.params.Append(preset.NewAction(ParamActionAdd, "Add Noise", func(preset.Action) {
		m.AddNoise(NewMultiNoise(
			NewFastNoise(),
			NewSine(),
			NewMask(),
			NewPow(),
			NewClamp(),
			NewMaths(),
			NewWarp(),
			NewNormalize(),
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
	n.Params().SetLabel(defaultName)
	n.Params().Prepend(preset.NewParam(ParamName, "Name", n.Params().Label(), func(p preset.Param[string]) {
		v := p.Val()
		if v == "" {
			v = defaultName
		}

		n.Params().SetLabel(v)
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
		m.injectNoiseChoices(m.params, nil)
		m.receiver.MarkDirty()
	}
}

func (m *Manager) injectNoiseChoices(p preset.ParamSet, owner Noise) {
	for _, pm := range p.All() {
		switch tpm := pm.(type) {
		case preset.Choice[Noise]:
			tpm.SetOptions(m.noiseOptions(owner))
		case preset.Param[Noise]:
			p.Replace(tpm, m.noiseChoice(tpm, owner))
		case preset.ParamSet:
			if owner == nil {
				for _, ns := range m.noises {
					if ns.Params() == p {
						owner = ns
						break
					}
				}
			}
			m.injectNoiseChoices(tpm, owner)
		}
	}
}

func (m *Manager) noiseChoice(pm preset.Param[Noise], owner Noise) preset.Choice[Noise] {
	return preset.NewChoice(pm.Id(), pm.Label(), nil, m.noiseOptions(owner), func(p preset.Param[Noise]) {
		pm.SetVal(p.Val())
	})
}

func (m *Manager) noiseOptions(owner Noise) []preset.Option[Noise] {
	opts := make([]preset.Option[Noise], 0, len(m.noises)+1)
	opts = append(opts, preset.NewOption[Noise](nil, "None"))
	for _, ns := range m.noises {
		if ns == owner {
			continue
		}
		ps := ns.Params()
		nsType := ps.QueryParamById(ParamType)[0].(preset.ChoiceGeneric)
		label := fmt.Sprintf("%s (%s)", ps.Label(), nsType.OptionsLabels()[nsType.ValIndex()])

		opts = append(opts, preset.NewOption[Noise](ns, label))
	}
	return opts
}

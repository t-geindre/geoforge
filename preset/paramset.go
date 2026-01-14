package preset

type ParamSet interface {
	ParamGeneric
	Append(p ...ParamGeneric)
	Prepend(p ...ParamGeneric)
	Add(after ParamId, params ...ParamGeneric)
	Remove(p ...ParamGeneric)
	Replace(old, new ParamGeneric)
	All() []ParamGeneric
	SetLabel(label string)
	QueryParamById(id ParamId) []ParamGeneric
	Clear()
	IsAnonymous() bool
}

type paramSet struct {
	set     []ParamGeneric
	id      ParamId
	label   string
	changed bool
}

func NewAnonymousParamSet() ParamSet {
	return &paramSet{}
}

func NewParamSet(id ParamId, label string) ParamSet {
	return &paramSet{
		id:    id,
		label: label,
	}
}

func (p *paramSet) Id() ParamId {
	return p.id
}

func (p *paramSet) Label() string {
	return p.label
}

func (p *paramSet) Prepend(params ...ParamGeneric) {
	p.changed = true
	for _, pr := range params {
		p.set = append([]ParamGeneric{pr}, p.set...)
	}

}

func (p *paramSet) All() []ParamGeneric {
	return p.set
}

// HasChanged since last call
func (p *paramSet) HasChanged() bool {
	changed := p.changed
	p.changed = false

	for _, pm := range p.set {
		if pm.HasChanged() {
			// range over all to reset their states
			changed = true
		}
	}
	return changed
}

func (p *paramSet) SetLabel(label string) {
	p.label = label
	p.changed = true
}

func (p *paramSet) QueryParamById(id ParamId) []ParamGeneric {
	var results []ParamGeneric

	for _, pm := range p.set {
		if pm.Id() == id {
			results = append(results, pm)
		}
		if subSet, ok := pm.(ParamSet); ok {
			subResults := subSet.QueryParamById(id)
			results = append(results, subResults...)
		}
	}

	return results
}

func (p *paramSet) Clear() {
	p.changed = true
	p.set = []ParamGeneric{}
}

func (p *paramSet) Remove(params ...ParamGeneric) {
	for _, toRemove := range params {
		for i, pm := range p.set {
			if pm == toRemove {
				// Remove the parameter by slicing
				p.set = append(p.set[:i], p.set[i+1:]...)
				p.changed = true
				break
			}
		}
	}
}

func (p *paramSet) Append(params ...ParamGeneric) {
	p.changed = true
	for _, pr := range params {
		p.set = append(p.set, pr)
	}
}

func (p *paramSet) Add(after ParamId, params ...ParamGeneric) {
	p.changed = true
	index := -1
	for i, pm := range p.set {
		if pm.Id() == after {
			index = i
			break
		}
	}
	if index != -1 {
		// Insert after the found index
		p.set = append(p.set[:index+1], append(params, p.set[index+1:]...)...)
	} else {
		// If not found, append to the end
		p.Append(params...)
	}
}

func (p *paramSet) IsAnonymous() bool {
	return p.label == ""
}

func (p *paramSet) Replace(old, new ParamGeneric) {
	for i, pm := range p.set {
		if pm == old {
			p.changed = true
			p.set[i] = new
			return
		}
	}
}

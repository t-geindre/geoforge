package preset

type ParamSet interface {
	ParamGeneric
	Append(p ...ParamGeneric)
	Prepend(p ...ParamGeneric)
	Remove(p ...ParamGeneric)
	All() []ParamGeneric
	SetLabel(label string)
	QueryParamById(id ParamId) []ParamGeneric
	Clear()
}

type paramSet struct {
	set   []ParamGeneric
	id    ParamId
	label string
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

func (p *paramSet) Append(params ...ParamGeneric) {
	for _, pr := range params {
		p.set = append(p.set, pr)
	}
}

func (p *paramSet) Prepend(params ...ParamGeneric) {
	for _, pr := range params {
		p.set = append([]ParamGeneric{pr}, p.set...)
	}
}

func (p *paramSet) All() []ParamGeneric {
	return p.set
}

// HasChanged since last call
func (p *paramSet) HasChanged() bool {
	changed := false
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
	p.set = []ParamGeneric{}
}

func (p *paramSet) Remove(params ...ParamGeneric) {
	for _, toRemove := range params {
		for i, pm := range p.set {
			if pm == toRemove {
				// Remove the parameter by slicing
				p.set = append(p.set[:i], p.set[i+1:]...)
				break
			}
		}
	}
}

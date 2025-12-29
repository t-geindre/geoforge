package preset

type ParamSet interface {
	ParamGeneric
	Append(p ParamGeneric)
	Prepend(p ParamGeneric)
	All() []ParamGeneric
	SetLabel(label string)
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

func (p *paramSet) Append(param ParamGeneric) {
	p.set = append(p.set, param)
}

func (p *paramSet) Prepend(param ParamGeneric) {
	p.set = append([]ParamGeneric{param}, p.set...)
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

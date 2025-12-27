package preset

type ParamSet []ParamGeneric

func NewParamSet() ParamSet {
	return ParamSet{}
}

// HasChanged since last call
func (ps ParamSet) HasChanged() bool {
	changed := false
	for _, p := range ps {
		if p.HasChanged() {
			changed = true
		}
	}
	return changed
}

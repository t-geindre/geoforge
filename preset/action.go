package preset

type Action interface {
	ParamGeneric
	Execute()
}

type action struct {
	id      ParamId
	label   string
	execute func()
}

func NewAction(id ParamId, label string, execute func()) Action {
	return &action{
		id:      id,
		label:   label,
		execute: execute,
	}
}

func (a *action) Execute() {
	a.execute()
}

func (a *action) Id() ParamId {
	return a.id
}

func (a *action) Label() string {
	return a.label
}

func (a *action) HasChanged() bool {
	return false // Actions never change, only execute
}

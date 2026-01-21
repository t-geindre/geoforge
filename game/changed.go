package game

type StateChanged interface {
	SetChanged()
	HasChanged(id uint8) bool
	RegisterChangeId() uint8
	ClearChangeId(id uint8)
}

type stateChanged struct {
	changed map[uint8]bool
	idx     uint8
}

func NewStateChanged() StateChanged {
	return &stateChanged{
		changed: make(map[uint8]bool),
	}
}

func (s *stateChanged) SetChanged() {
	for k := range s.changed {
		s.changed[k] = true
	}
}

func (s *stateChanged) HasChanged(id uint8) bool {
	changed, exists := s.changed[id]
	if !exists {
		changed = true
	}

	s.changed[id] = false

	return changed
}

func (s *stateChanged) RegisterChangeId() uint8 {
	s.idx++
	return s.idx
}

func (s *stateChanged) ClearChangeId(id uint8) {
	delete(s.changed, id)
}

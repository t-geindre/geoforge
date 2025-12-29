package game

type Updater interface {
	Update()
}

type updateFunc struct {
	f func()
}

func NewUpdateFunc(f func()) Updater {
	return &updateFunc{f: f}
}

func (u *updateFunc) Update() {
	u.f()
}

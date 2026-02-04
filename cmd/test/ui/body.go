package ui

import "github.com/ebitenui/ebitenui/widget"

type Body struct {
	*widget.Container
}

func NewBody() *Body {
	return &Body{
		Container: widget.NewContainer(),
	}
}

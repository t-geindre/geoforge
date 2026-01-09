package ui

import (
	"image"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
)

type Component interface {
	UiUpdate(ctx *debugui.Context)
}

type Window struct {
	dui        debugui.DebugUI
	components []Component
	label      string
}

func NewWindow(label string, components ...Component) *Window {
	return &Window{
		dui:        debugui.DebugUI{},
		components: components,
		label:      label,
	}
}

func (g *Window) Draw(screen *ebiten.Image) {
	g.dui.Draw(screen)
}

func (g *Window) Update() bool {
	ics, err := g.dui.Update(func(ctx *debugui.Context) error {
		ctx.Window(g.label, image.Rect(0, 0, 400, 700), func(layout debugui.ContainerLayout) {
			ctx.Loop(len(g.components), func(i int) {
				g.components[i].UiUpdate(ctx)
			})
		})
		return nil
	})

	if err != nil {
		panic(err)
	}

	return ics > 0
}

func (g *Window) Append(c Component) {
	g.components = append(g.components, c)
}

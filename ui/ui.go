package ui

import (
	"image"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
)

type Component interface {
	UiUpdate(ctx *debugui.Context)
}

type Ui struct {
	dui        debugui.DebugUI
	components []Component
	label      string
}

func NewUi(label string, components ...Component) *Ui {
	return &Ui{
		dui:        debugui.DebugUI{},
		components: components,
		label:      label,
	}
}

func (g *Ui) Draw(screen *ebiten.Image) {
	g.dui.Draw(screen)
}

func (g *Ui) Update() bool {
	ics, err := g.dui.Update(func(ctx *debugui.Context) error {
		ctx.Window(g.label, image.Rect(0, 0, 400, 700), func(layout debugui.ContainerLayout) {
			for _, comp := range g.components {
				comp.UiUpdate(ctx)
			}
		})
		return nil
	})

	if err != nil {
		panic(err)
	}

	return ics > 0
}

func (g *Ui) Append(c Component) {
	g.components = append(g.components, c)
}

/*
type Ui struct {
	dui debugui.DebugUI
}

func NewUi() *Ui {
	return &Ui{
		dui: debugui.DebugUI{},
	}
}

// UiUpdate the metrics UI, returns true if capturing events
func (ui *Ui) UiUpdate() bool {
	ics, _ := ui.dui.UiUpdate(func(ctx *debugui.Context) error {
		ctx.Window("Settings", image.Rect(0, 0, 320, 600), func(layout debugui.ContainerLayout) {
			ctx.SetGridLayout(nil, []int{28, -1})
			ctx.Panel(func(layout debugui.ContainerLayout) {
				ctx.Text(fmt.Sprintf(
					"FPS: %.0f, TPS: %.0f, Chunks: %d/%d",
					ebiten.ActualFPS(),
					ebiten.ActualTPS(),
					g.renderer.DrawnChunks(),
					len(g.world.Chunks()),
				))
			})
			ctx.Header("Camera", false, func() {
				x, y := g.cam.Position()
				ctx.Text(fmt.Sprintf("Position: %.0f, %.0f", x, y))
				ctx.Text(fmt.Sprintf("Zoom: %.0f%%", g.cam.Zoom()*100))
				ctx.Button("Reset").On(func() {
					g.cam.Reset()
				})
			})
			ctx.Header("Noise", true, func() {
				ParamSetUI(ctx, g.world.Noise().Params())
			})
			ctx.Header("Presets", false, func() {
				ctx.SetGridLayout([]int{-4, -1, -1, -1}, nil)
				ctx.Text("Default")
				ctx.Button("Load")
				ctx.Button("Save")
				ctx.Button("Delete")

				t := ""
				ctx.SetGridLayout([]int{-4, -1}, nil)
				ctx.TextField(&t)
				ctx.Button("New")
			})
		})
		return nil
	})

	return ics > 0
}
*/

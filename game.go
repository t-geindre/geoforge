package main

import (
	"fmt"
	"geoforge/cam"
	"geoforge/noise"
	"geoforge/render"
	"geoforge/ui"
	"geoforge/world"
	"image"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	cam      cam.Camera
	world    *world.World
	renderer *render.Renderer
	ui       debugui.DebugUI
}

func NewGame() *Game {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.MaximizeWindow()

	return &Game{
		cam:      cam.NewWheelZoom(cam.NewMousePan(cam.NewCamera())),
		world:    world.NewWorld(256, 10, 1, 4096, noise.NewFixed()),
		renderer: render.NewRenderer(),
	}
}

func (g *Game) Update() error {
	ics, err := g.ui.Update(func(ctx *debugui.Context) error {
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
				ui.ParamSetUI(ctx, g.world.Noise().Params())
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

	if err != nil {
		return err
	}

	g.world.Update(g.cam.WorldRect())

	if ics == 0 {
		g.cam.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.Draw(g.world, g.cam, screen)
	g.ui.Draw(screen)
}

func (g *Game) Layout(x, y int) (int, int) {
	g.cam.SetViewport(x, y)
	return x, y
}

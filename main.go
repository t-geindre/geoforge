package main

import (
	"geoforge/cam"
	"geoforge/game"
	"geoforge/noise"
	"geoforge/render"
	"geoforge/ui"
	"geoforge/world"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	rdr := render.NewRenderer()
	wld := world.NewWorld(2)

	nmg := noise.NewNoiseManager(wld)

	camera := cam.NewWheelZoom(cam.NewMousePan(cam.NewCamera()))

	// Metrics
	mStyle := ui.DefaultGraphStyle()

	mFps := ui.NewTimeSeries("FPS", 200, 5, mStyle, func() float32 {
		return float32(ebiten.ActualFPS())
	})
	mFps.SetFixedScale(0, 120)

	mTps := ui.NewTimeSeries("TPS", 200, 5, mStyle, func() float32 {
		return float32(ebiten.ActualTPS())
	})
	mTps.SetFixedScale(0, 120)

	mChunksDrawn := ui.NewMetric("Chunks drawn", func() float32 {
		return float32(rdr.DrawnChunks())
	})
	mChunks := ui.NewMetric("Chunks total", func() float32 {
		return float32(len(wld.Chunks()))
	})

	// UI
	gui := ui.NewWindow(
		"Geoforge",
		ui.NewMetrics(mFps, mTps, mChunksDrawn, mChunks),
		ui.NewCamera(camera),
		ui.NewParamSet("Noise", nmg.Params()),
		ui.NewParamSet("Renderer", rdr.Params()),
	)

	err := ebiten.RunGame(game.NewGame(
		game.NewUpdateFunc(func() {
			wld.Update(camera.WorldRect())
			if !gui.Update() {
				// Window not capturing events
				camera.Update()
			}
		}),
		game.NewDrawFunc(func(screen *ebiten.Image) {
			rdr.Draw(wld, camera, screen)
		}),
		game.NewLayoutFunc(func(x, y int) (int, int) {
			camera.SetViewport(x, y)
			return x, y
		}),
		gui,
		nmg,
	))

	if err != nil {
		panic(err)
	}
}

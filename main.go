package main

import (
	"geoforge/cam"
	"geoforge/game"
	"geoforge/noise"
	"geoforge/preset"
	"geoforge/render"
	"geoforge/ui"
	"geoforge/world"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.MaximizeWindow()

	nse := noise.NewSine()
	rdr := render.NewRenderer()
	wld := world.NewWorld(256, 10, 1, 4096, nse)

	nsePs := nse.Params()
	nsePs.Prepend(preset.NewParam(0, "Name", "Land mask", func(v string) {
		nsePs.SetLabel(v)
	}))

	ps := preset.NewAnonymousParamSet()
	ps.Append(nsePs)

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
	gui := ui.NewUi(
		"Geoforge",
		ui.NewMetrics(mFps, mTps, mChunksDrawn, mChunks),
		ui.NewCamera(camera),
		ui.NewParamSet("Noise", ps),
	)

	err := ebiten.RunGame(game.NewGame(
		game.NewUpdateFunc(func() {
			wld.Update(camera.WorldRect())
			if !gui.Update() {
				// Ui not capturing events
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
	))

	if err != nil {
		panic(err)
	}
}

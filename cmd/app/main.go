package main

import (
	"geoforge/camera"
	"geoforge/game"
	"geoforge/noise"
	"geoforge/render"
	"geoforge/ui"
	"geoforge/world"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.MaximizeWindow()

	// Core components
	cam := camera.NewWheelZoom(camera.NewMousePan(camera.NewCamera()))
	wld := world.NewWorld(2, cam)
	defer wld.Close()
	rdr := render.NewRenderer(wld, cam)
	nmg := noise.NewNoiseManager(wld)

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
		ui.NewCamera(cam),
		ui.NewParamSet("Noise", nmg.Params()),
		ui.NewParamSet("Renderer", rdr.Params()),
	)

	err := ebiten.RunGame(game.NewGame(
		game.NewUpdateFunc(func() {
			cam.Lock(gui.Update())
		}),
		game.NewLayoutFunc(func(x, y int) (int, int) {
			cam.SetViewport(x, y)
			return x, y
		}),
		rdr, gui, nmg, wld, cam,
	))

	if err != nil {
		panic(err)
	}
}

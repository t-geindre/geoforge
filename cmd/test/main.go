package main

import (
	"geoforge/camera"
	"geoforge/cmd/test/ui"
	"geoforge/game"
	"geoforge/noise"
	"geoforge/render"
	"geoforge/world"
	"math"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	//debug()
	//return

	// THEME AND LAYOUT
	theme := ui.NewTheme()
	layout := ui.NewLayout(theme)

	// NOISE SETUP
	cam, wrld, rdr := noiseSetup()

	// MAIN MENU
	menu := ui.NewMenu(theme)
	layout.AddChild(menu)

	// SPLIT PAN
	split := ui.NewSplit(theme)
	layout.AddChild(split)

	// DESKTOP (LEFT)
	desktop := ui.NewDesktop(1000, 1000)
	split.AddChild(desktop)

	for j := 0; j < 3; j++ {
		drag := ui.NewDraggable(j*200, 50, theme)
		drag.AddChild(getGridForm(theme))
		drag.SetTitle("Draggable " + string(rune('A'+j)))
		desktop.AddDraggable(drag)
	}

	// NOISE PREVIEW (RIGHT)
	previewWidget, previewDraw := ui.NewPreview(cam, rdr)
	split.AddChild(previewWidget)

	// STATUS
	status := ui.NewStatus(theme, rdr, wrld, cam)
	layout.AddChild(status)

	// UI
	ui := &ebitenui.UI{Container: layout.Container}
	ui.PrimaryTheme = theme.Theme

	// UPDATES
	updates := game.NewUpdateFunc(func() {
		cam.Update()
		wrld.Update()
		rdr.Update()
		ui.Update()
	})

	// DRAWS
	draws := game.NewDrawFunc(func(screen *ebiten.Image) {
		ui.Draw(screen)
		previewDraw(screen)
	})

	if err := ebiten.RunGame(game.NewGame(updates, draws)); err != nil {
		panic(err)
	}
}

func debug() {
	cam := camera.NewWheelZoom(camera.NewMousePan(camera.NewCamera()))
	cam.SetViewport(200, 200)
	cam.ScreenMoveTo(200, 200)

	wrld := world.NewWorld(1, cam)

	nse := noise.NewFastNoise()
	wrld.SetNoise(nse)

	rdr := render.NewRenderer(wrld, cam)
	draw := game.NewDrawFunc(func(screen *ebiten.Image) {
		rdr.Draw(screen)
	})
	rdrUpdate := game.NewUpdateFunc(func() {
		rdr.Update()
	})

	var time float64
	move := game.NewUpdateFunc(func() {
		time += 0.01
		cam.ScreenMoveTo(int(math.Sin(time)*100+200), int(math.Cos(time)*100+100))
	})

	ebiten.RunGame(game.NewGame(cam, wrld, draw, rdrUpdate, move))

}

package main

import (
	"geoforge/cmd/test/ui"
	"geoforge/game"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("GeoForge - Noise Editor")
	ebiten.MaximizeWindow()

	// THEME AND LAYOUT
	theme, err := ui.NewTheme()
	if err != nil {
		panic(err)
	}
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

	// CONNECTORS
	conn := ui.NewConnections(theme, desktop)

	// DRAGGABLES
	for j := 0; j < 4; j++ {
		icon := theme.IconsTheme.NoiseSmall
		title := "Noise " + string(rune('A'+j))
		if j == 0 {
			icon = theme.IconsTheme.Rendering
			title = "Renderer"
		}
		drag := ui.NewDraggable(j*200, 50, theme, icon)
		drag.AddChild(getGridForm(theme, conn))
		drag.SetTitle(title)
		desktop.AddDraggable(drag)
	}

	// NOISE PREVIEW (RIGHT)
	split.AddChild(ui.NewPreview(cam, rdr))

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
		conn.Update()
	})

	// DRAWS
	draws := game.NewDrawFunc(func(screen *ebiten.Image) {
		ui.Draw(screen)
		conn.Draw(screen)
	})

	if err := ebiten.RunGame(game.NewGame(updates, draws)); err != nil {
		panic(err)
	}
}

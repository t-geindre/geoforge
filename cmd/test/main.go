package main

import (
	"geoforge/cmd/test/ui/layout"
	"geoforge/cmd/test/ui/theme"
	"geoforge/cmd/test/ui/widgets"
	"geoforge/game"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("GeoForge - Noise Editor")
	ebiten.MaximizeWindow()

	// THEME AND LAYOUT
	theme, err := theme.NewDefaultTheme()
	if err != nil {
		panic(err)
	}
	lyt := layout.NewLayout(theme)

	// NOISE SETUP
	cam, wrld, rdr := noiseSetup()

	// MAIN MENU
	menu := layout.NewMenu(theme)
	lyt.AddChild(menu)

	// SPLIT PAN
	split := widgets.NewSplit(theme)
	lyt.AddChild(split)

	// DESKTOP (LEFT)
	desktop := widgets.NewDesktop(1000, 1000)
	split.AddChild(desktop)

	// CONNECTORS
	conn := widgets.NewConnections(theme, desktop)

	// DRAGGABLES
	for j := 0; j < 4; j++ {
		icon := theme.IconsTheme.NoiseSmall
		title := "Noise " + string(rune('A'+j))
		if j == 0 {
			icon = theme.IconsTheme.Rendering
			title = "Renderer"
		}
		drag := widgets.NewDraggable(j*200, 50, theme, icon)
		drag.AddChild(getGridForm(theme, conn))
		drag.SetTitle(title)
		desktop.AddDraggable(drag)
	}

	// NOISE PREVIEW (RIGHT)
	split.AddChild(widgets.NewPreview(cam, rdr))

	// STATUS
	status := layout.NewStatus(theme, rdr, wrld, cam)
	lyt.AddChild(status)

	// UI
	ui := &ebitenui.UI{Container: lyt.Container}
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

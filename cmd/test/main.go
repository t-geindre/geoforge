package main

import (
	"geoforge/cmd/test/ui/layout"
	"geoforge/cmd/test/ui/theme"
	"geoforge/cmd/test/ui/widgets"
	"geoforge/game"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	root := layout.NewRoot()
	grid := layout.NewGrid(theme)
	root.AddChild(grid)

	// NOISE SETUP
	cam, wrld, rdr := noiseSetup()

	// MAIN MENU
	menu := layout.NewMenu(theme)
	grid.AddChild(menu)

	menu.AddIcon(layout.MenuPosLeft, theme.IconsTheme.Noise)
	menu.AddButton(layout.MenuPosLeft, "Add", theme.IconsTheme.Add, nil)
	menu.AddButton(layout.MenuPosLeft, "Clear", theme.IconsTheme.Delete, nil)

	menu.AddIcon(layout.MenuPosCenter, theme.IconsTheme.File)
	menu.AddButton(layout.MenuPosCenter, "Open", theme.IconsTheme.Open, nil)
	menu.AddButton(layout.MenuPosCenter, "Save", theme.IconsTheme.Save, nil)

	// SPLIT PAN
	split := widgets.NewSplit(theme)
	grid.AddChild(split)

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
			icon = theme.IconsTheme.CameraSmall
			title = "Renderer"
		}
		drag := widgets.NewDraggable(j*200, 50, theme, icon)
		drag.AddChild(getGridForm(theme, conn))
		drag.SetTitle(title)
		desktop.AddDraggable(drag)
	}

	// NOISE PREVIEW (RIGHT)
	preview := widgets.NewPreview(cam, rdr)
	split.AddChild(preview)

	// NOISE PREVIEW (TOGGLE FS)
	toggleFS := func(fs bool) {
		if fs {
			root.RemoveChild(grid)
			root.AddChild(preview)
			conn.SetVisible(false)
			ebiten.SetFullscreen(true)
		} else {
			root.RemoveChild(preview)
			root.AddChild(grid)
			conn.SetVisible(true)
			ebiten.SetFullscreen(false)
		}
	}

	fsKeysListener := func() {
		if ebiten.IsFullscreen() {
			if inpututil.IsKeyJustPressed(ebiten.KeyEscape) || inpututil.IsKeyJustPressed(ebiten.KeyF11) {
				toggleFS(false)
			}
		} else {
			if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
				toggleFS(true)
			}
		}
	}

	// CAMERA MENU
	menu.AddIcon(layout.MenuPosRight, theme.IconsTheme.Camera)
	menu.AddButton(layout.MenuPosRight, "Center", theme.IconsTheme.Center, func() {
		cam.ByPassLock(func() { cam.MoveTo(0, 0) })
	})
	menu.AddButton(layout.MenuPosRight, "Reset", theme.IconsTheme.Zoom, func() {
		cam.ByPassLock(func() { cam.SetZoom(1) })
	})
	menu.AddButton(layout.MenuPosRight, "Full screen", theme.IconsTheme.Fullscreen, func() {
		toggleFS(true)
	})

	// STATUS
	status := layout.NewStatus(theme, rdr, wrld, cam)
	grid.AddChild(status)

	// UI
	ui := &ebitenui.UI{Container: root}
	ui.PrimaryTheme = theme.Theme

	// UPDATES
	updates := game.NewUpdateFunc(func() {
		cam.Update()
		wrld.Update()
		rdr.Update()
		ui.Update()
		conn.Update()
		fsKeysListener()
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

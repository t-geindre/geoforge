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
	th, err := theme.NewDefaultTheme()
	if err != nil {
		panic(err)
	}
	root := layout.NewRoot()
	grid := layout.NewGrid(th)
	root.AddChild(grid)

	// NOISE SETUP
	cam, wrld, rdr := noiseSetup()

	// MAIN MENU
	menu := layout.NewMenu(th)
	grid.AddChild(menu)

	menu.AddIcon(layout.MenuPosLeft, th.IconsTheme.NormalAccent[theme.IconNoise])
	menu.AddButton(layout.MenuPosLeft, "Add", th.IconsTheme.Normal[theme.IconAdd], nil)
	menu.AddButton(layout.MenuPosLeft, "Clear", th.IconsTheme.Normal[theme.IconDelete], nil)

	menu.AddIcon(layout.MenuPosCenter, th.IconsTheme.NormalAccent[theme.IconFile])
	menu.AddButton(layout.MenuPosCenter, "Open", th.IconsTheme.Normal[theme.IconOpen], nil)
	menu.AddButton(layout.MenuPosCenter, "Save", th.IconsTheme.Normal[theme.IconSave], nil)

	// SPLIT PAN
	split := widgets.NewSplit(th)
	grid.AddChild(split)

	// DESKTOP (LEFT)
	desktop := widgets.NewDesktop(1000, 1000)
	split.AddChild(desktop)

	// CONNECTORS
	conn := widgets.NewConnections(th, desktop)

	// DRAGGABLES
	for j := 0; j < 4; j++ {
		icon := th.IconsTheme.Small[theme.IconNoise]
		title := "Noise " + string(rune('A'+j))
		if j == 0 {
			icon = th.IconsTheme.Small[theme.IconCamera]
			title = "Renderer"
		}
		drag := widgets.NewDraggable(j*200, 50, th, icon)
		drag.AddChild(getGridForm(th, conn))
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
	menu.AddIcon(layout.MenuPosRight, th.IconsTheme.NormalAccent[theme.IconCamera])
	menu.AddButton(layout.MenuPosRight, "Center", th.IconsTheme.Normal[theme.IconCenter], func() {
		cam.ByPassLock(func() { cam.MoveTo(0, 0) })
	})
	menu.AddButton(layout.MenuPosRight, "Reset", th.IconsTheme.Normal[theme.IconZoom], func() {
		cam.ByPassLock(func() { cam.SetZoom(1) })
	})
	menu.AddButton(layout.MenuPosRight, "Full screen", th.IconsTheme.Normal[theme.IconFullscreen], func() {
		toggleFS(true)
	})

	// STATUS
	status := layout.NewStatus(th, rdr, wrld, cam)
	grid.AddChild(status)

	// UI
	ui := &ebitenui.UI{Container: root}
	ui.PrimaryTheme = th.Theme

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

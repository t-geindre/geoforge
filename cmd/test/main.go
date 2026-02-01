package main

import (
	"fmt"
	"geoforge/game"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	theme := NewDefaultTheme()
	builder := NewBuilder(theme)

	root := builder.NewMainContainer()

	menuBar := builder.NewMenuBar()
	menuBar.AddChild(builder.NewButton("File"))
	menuBar.AddChild(builder.NewButton("View"))

	statusBar := builder.NewStatusBar()
	statusBar.AddChild(builder.NewText("Ready."))

	fps := builder.NewText("FPS 60")
	statusBar.AddChild(fps)
	fpdUpdater := game.NewUpdateFunc(func() {
		fps.Label = fmt.Sprintf("FPS %0.f", ebiten.ActualFPS())
	})

	root.AddChild(menuBar)

	desktop := NewDesktop(20000, 20000)

	dragA := NewDraggable(100, 100, colornames.Red)
	dragB := NewDraggable(10, 20, colornames.Green)
	dragC := NewDraggable(20, 30, colornames.Blue)

	desktop.AddDraggable(dragA)
	desktop.AddDraggable(dragB)
	desktop.AddDraggable(dragC)

	updater := game.NewUpdateFunc(desktop.UpdateDragging)

	dragABody := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(1),
				widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false}),
				widget.GridLayoutOpts.Padding(&widget.Insets{Top: 5, Bottom: 5, Left: 5, Right: 5}),
				widget.GridLayoutOpts.Spacing(5, 5),
			),
		),
	)
	dragABody.AddChild(builder.NewButton("Hello"))
	dragABody.AddChild(builder.NewButton("Hello"))
	dragABody.AddChild(builder.NewButton("Hello"))
	dragABody.AddChild(builder.NewButton("Hello"))
	dragABody.AddChild(builder.NewButton("Hello"))
	dragA.AddChild(dragABody)

	dragB.AddChild(builder.NewButton("World"))
	dragB.AddChild(builder.NewButton("World"))
	dragB.AddChild(builder.NewButton("World"))
	dragB.AddChild(builder.NewButton("World"))
	dragB.AddChild(builder.NewButton("World"))
	dragB.AddChild(builder.NewButton("World"))
	dragC.AddChild(builder.NewText("!"))

	root.AddChild(desktop)
	root.AddChild(statusBar)

	ui := &ebitenui.UI{Container: root}

	if err := ebiten.RunGame(game.NewGame(ui, fpdUpdater, updater)); err != nil {
		panic(err)
	}
}

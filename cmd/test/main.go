package main

import (
	"geoforge/cmd/test/ui2"
	"geoforge/game"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/themes"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	theme := themes.GetBasicDarkTheme()

	menu := ui2.NewMenu(theme)
	desktop := ui2.NewDesktop(1000, 1000)
	status := ui2.NewStatus(theme)

	layout := ui2.NewLayout()
	layout.AddChild(menu, desktop, status)

	drag := ui2.NewDraggable(50, 50, theme)
	dbody := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(2),
				widget.GridLayoutOpts.Stretch([]bool{false, true}, []bool{false}),
				widget.GridLayoutOpts.Padding(&widget.Insets{Top: 10, Bottom: 10, Left: 10, Right: 10}),
				widget.GridLayoutOpts.Spacing(15, 15),
			),
		),
	)
	drag.AddChild(dbody)
	dbody.AddChild(
		widget.NewText(widget.TextOpts.TextLabel("Button")),
		widget.NewButton(widget.ButtonOpts.TextLabel("Button")),
	)
	dbody.AddChild(
		widget.NewText(widget.TextOpts.TextLabel("Checkbox")),
		widget.NewCheckbox(),
	)
	dbody.AddChild(
		widget.NewText(widget.TextOpts.TextLabel("Slider")),
		widget.NewSlider(
			widget.SliderOpts.MinMax(0, 100),
			widget.SliderOpts.InitialCurrent(50),
		),
	)
	dbody.AddChild(
		widget.NewText(widget.TextOpts.TextLabel("Text Input")),
		widget.NewTextInput(),
	)

	var items []any
	for i := 1; i <= 15; i++ {
		items = append(items, "Entry "+string(rune('A'+i-1)))
	}
	dbody.AddChild(
		widget.NewText(widget.TextOpts.TextLabel("List")),
		widget.NewListComboButton(
			widget.ListComboButtonOpts.Entries(items),
			widget.ListComboButtonOpts.EntryLabelFunc(func(v any) string {
				return v.(string)
			}, func(v any) string {
				return v.(string)
			}),
		),
	)

	desktop.AddDraggable(drag)

	ui := &ebitenui.UI{Container: layout.Container}
	ui.PrimaryTheme = theme

	updater := game.NewUpdateFunc(desktop.UpdateDragging)

	if err := ebiten.RunGame(game.NewGame(ui, updater)); err != nil {
		panic(err)
	}
}

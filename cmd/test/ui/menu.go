package ui

import (
	"geoforge/cmd/test/ui/theme"

	"github.com/ebitenui/ebitenui/input"
	"github.com/ebitenui/ebitenui/widget"
)

func NewMenu(theme *theme.Theme) *widget.Container {
	c := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(theme.PanelTheme.ForegroundImage),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(5),
				widget.GridLayoutOpts.Stretch([]bool{false, true, false, true, false}, []bool{true}),
				widget.GridLayoutOpts.Padding(theme.PanelTheme.Padding),
			),
		),
	)

	left := newMenuBox(theme)
	center := newMenuBox(theme)
	right := newMenuBox(theme)

	left.AddChild(newMenuIcon(theme.IconsTheme.Noise))
	left.AddChild(newMenuEntry(theme, "Add", theme.IconsTheme.Add))
	left.AddChild(newMenuEntry(theme, "Clear", theme.IconsTheme.Delete))

	center.AddChild(newMenuIcon(theme.IconsTheme.File))
	center.AddChild(newMenuEntry(theme, "Open", theme.IconsTheme.Open))
	center.AddChild(newMenuEntry(theme, "Save", theme.IconsTheme.Save))

	right.AddChild(newMenuIcon(theme.IconsTheme.Camera))
	right.AddChild(newMenuEntry(theme, "Center", theme.IconsTheme.Center))
	right.AddChild(newMenuEntry(theme, "Reset", theme.IconsTheme.Zoom))

	c.AddChild(left, newMenuBox(theme), center, newMenuBox(theme), right)

	return c
}

func newMenuEntry(theme *theme.Theme, label string, icon *widget.GraphicImage) *widget.Container {
	iconWidget := widget.NewGraphic(
		widget.GraphicOpts.Image(icon.Idle),
		widget.GraphicOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)

	txtWidget := widget.NewText(
		widget.TextOpts.Text(label, theme.MainMenuTheme.Font, theme.DefaultTextColor),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)

	var c *widget.Container
	c = widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(theme.MainMenuTheme.ButtonImage.Idle),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Spacing(theme.MainMenuTheme.IconSpacing),
			widget.RowLayoutOpts.Padding(theme.MainMenuTheme.ButtonPadding),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.CursorEnterHandler(func(args *widget.WidgetCursorEnterEventArgs) {
				iconWidget.Image = icon.Hover
				c.SetBackgroundImage(theme.MainMenuTheme.ButtonImage.Hover)
			}),
			widget.WidgetOpts.CursorExitHandler(func(args *widget.WidgetCursorExitEventArgs) {
				iconWidget.Image = icon.Idle
				c.SetBackgroundImage(theme.MainMenuTheme.ButtonImage.Idle)
			}),
			widget.WidgetOpts.CursorHovered(input.CURSOR_POINTER),
		),
	)

	c.AddChild(iconWidget)
	c.AddChild(txtWidget)

	return c
}

func newMenuIcon(icon *widget.GraphicImage) *widget.Graphic {
	return widget.NewGraphic(
		widget.GraphicOpts.Image(icon.Idle),
		widget.GraphicOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)
}

func newMenuBox(theme *theme.Theme) *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Spacing(theme.PanelTheme.Spacing),
		)),
	)
}

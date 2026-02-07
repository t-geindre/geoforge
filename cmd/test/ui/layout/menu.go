package layout

import (
	"geoforge/cmd/test/ui/theme"

	"github.com/ebitenui/ebitenui/input"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type MenuPos int

const (
	MenuPosLeft MenuPos = iota
	MenuPosCenter
	MenuPosRight
)

type Menu struct {
	*widget.Container
	boxes map[MenuPos]*widget.Container
	theme *theme.Theme
}

func NewMenu(theme *theme.Theme) *Menu {
	m := &Menu{theme: theme}

	m.Container = widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(theme.PanelTheme.ForegroundImage),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(5),
				widget.GridLayoutOpts.Stretch([]bool{false, true, false, true, false}, []bool{true}),
				widget.GridLayoutOpts.Padding(theme.PanelTheme.Padding),
			),
		),
	)

	m.boxes = make(map[MenuPos]*widget.Container)
	m.boxes[MenuPosLeft] = m.newBox()
	m.boxes[MenuPosCenter] = m.newBox()
	m.boxes[MenuPosRight] = m.newBox()

	m.Container.AddChild(
		m.boxes[MenuPosLeft],
		m.newBox(),
		m.boxes[MenuPosCenter],
		m.newBox(),
		m.boxes[MenuPosRight],
	)

	return m
}

func (m *Menu) AddButton(pos MenuPos, label string, icon *widget.GraphicImage, on func()) {
	box, ok := m.boxes[pos]
	if !ok {
		return
	}

	if on == nil {
		on = func() {}
	}

	iconWidget := widget.NewGraphic(
		widget.GraphicOpts.Image(icon.Idle),
		widget.GraphicOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)

	txtWidget := widget.NewText(
		widget.TextOpts.Text(label, m.theme.MainMenuTheme.Font, m.theme.DefaultTextColor),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)

	var c *widget.Container
	c = widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(m.theme.MainMenuTheme.ButtonImage.Idle),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Spacing(m.theme.MainMenuTheme.IconSpacing),
			widget.RowLayoutOpts.Padding(m.theme.MainMenuTheme.ButtonPadding),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.CursorEnterHandler(func(args *widget.WidgetCursorEnterEventArgs) {
				iconWidget.Image = icon.Hover
				c.SetBackgroundImage(m.theme.MainMenuTheme.ButtonImage.Hover)
			}),
			widget.WidgetOpts.CursorExitHandler(func(args *widget.WidgetCursorExitEventArgs) {
				iconWidget.Image = icon.Idle
				c.SetBackgroundImage(m.theme.MainMenuTheme.ButtonImage.Idle)
			}),
			widget.WidgetOpts.CursorHovered(input.CURSOR_POINTER),
			widget.WidgetOpts.MouseButtonClickedHandler(func(args *widget.WidgetMouseButtonClickedEventArgs) {
				if args.Button == ebiten.MouseButtonLeft {
					on()
				}
			}),
		),
	)

	c.AddChild(iconWidget)
	c.AddChild(txtWidget)

	box.AddChild(c)
}

func (m *Menu) AddIcon(pos MenuPos, icon *widget.GraphicImage) {
	box, ok := m.boxes[pos]
	if !ok {
		return
	}

	box.AddChild(widget.NewGraphic(
		widget.GraphicOpts.Image(icon.Idle),
		widget.GraphicOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	))
}

func (m *Menu) newBox() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Spacing(m.theme.PanelTheme.Spacing),
		)),
	)
}

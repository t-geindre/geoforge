package layout

import (
	"geoforge/cmd/test/ui/theme"

	"github.com/ebitenui/ebitenui/widget"
)

func NewGrid(theme *theme.Theme) *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(&widget.StackedLayoutData{}),
		),
		widget.ContainerOpts.BackgroundImage(theme.PanelTheme.BackgroundImage),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(1),
				widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, true, false}),
			),
		),
	)
}

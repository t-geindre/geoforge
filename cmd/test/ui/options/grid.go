package options

import (
	"geoforge/cmd/test/ui"

	"github.com/ebitenui/ebitenui/widget"
)

type GridOptions struct {
	theme *ui.Theme
	opts  []widget.GridLayoutOpt
	*ContainerOptions
}

func NewGridOptions(t *ui.Theme) *GridOptions {
	return &GridOptions{
		theme:            t,
		ContainerOptions: NewContainerOptions(t),
	}
}

func (g *GridOptions) WithColumns(cols int) {
	g.opts = append(g.opts, widget.GridLayoutOpts.Columns(cols))
}

func (g *GridOptions) WithSpacing() *GridOptions {
	g.opts = append(
		g.opts,
		widget.GridLayoutOpts.Spacing(g.theme.SurfaceSpaceX, g.theme.SurfaceSpaceY),
		widget.GridLayoutOpts.Padding(g.theme.SurfaceInset),
	)

	return g
}

func (g *GridOptions) WithStretch(c []bool, r []bool) *GridOptions {
	g.opts = append(
		g.opts,
		widget.GridLayoutOpts.Stretch(c, r),
	)

	return g
}

func (g *GridOptions) GetOptions() []widget.ContainerOpt {
	return append(g.ContainerOptions.GetOptions(),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(g.opts...),
		),
	)
}

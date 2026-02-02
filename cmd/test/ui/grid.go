package ui

import "github.com/ebitenui/ebitenui/widget"

type Grid struct {
	*Builder
	*Container
	container *widget.Container
	opts      []widget.GridLayoutOpt
	cols      int
}

func newGrid(b *Builder) *Grid {
	return &Grid{
		Builder:   b,
		Container: newContainer(b),
		opts:      []widget.GridLayoutOpt{},
	}
}

func (g *Grid) WithColumns(cols int) *Grid {
	g.ensureContainerBuilding()

	if cols < 1 {
		panic("columns must be at least 1")
	}

	g.cols = cols
	g.opts = append(g.opts, widget.GridLayoutOpts.Columns(cols))

	return g
}

func (g *Grid) WithSpacing() *Grid {
	g.ensureContainerBuilding()

	g.opts = append(
		g.opts,
		widget.GridLayoutOpts.Spacing(g.theme.SurfaceSpaceX, g.theme.SurfaceSpaceY),
		widget.GridLayoutOpts.Padding(g.theme.SurfaceInset),
	)

	return g
}

func (g *Grid) WithStretch(c []bool, r []bool) *Grid {
	g.ensureContainerBuilding()

	g.opts = append(
		g.opts,
		widget.GridLayoutOpts.Stretch(c, r),
	)

	return g
}

func (g *Grid) AddRow(children ...widget.PreferredSizeLocateableWidget) *Grid {
	g.buildContainer()

	if len(children) != g.cols {
		panic("number of children does not match number of columns")
	}

	g.container.AddChild(children...)

	return g
}

func (g *Grid) EndGrid() *widget.Container {
	g.buildContainer()

	return g.container
}

func (g *Grid) buildContainer() {
	if g.container != nil {
		return
	}

	g.container = widget.NewContainer(
		append([]widget.ContainerOpt{
			widget.ContainerOpts.Layout(
				widget.NewGridLayout(g.opts...),
			),
		})...,
	)
}

func (g *Grid) ensureContainerBuilding() {
	if g.container != nil {
		panic("grid is already created")
	}
}

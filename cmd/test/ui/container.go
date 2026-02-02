package ui

import "github.com/ebitenui/ebitenui/widget"

type Container struct {
	*Builder
	container *widget.Container
	opts      []widget.ContainerOpt
}

func newContainer(b *Builder) *Container {
	return &Container{
		Builder: b,
		opts:    []widget.ContainerOpt{},
	}
}

func (c *Container) WithBackground() *Container {
	c.ensureContainerBuilding()

	c.opts = append(
		c.opts,
		widget.ContainerOpts.BackgroundImage(c.theme.SurfaceImage),
	)

	return c
}

func (c *Container) ensureContainerBuilding() {
	if c.container != nil {
		panic("container already built; cannot modify options")
	}
}

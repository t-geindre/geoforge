package options

import (
	"geoforge/cmd/test/ui"

	"github.com/ebitenui/ebitenui/widget"
)

type ContainerOptions struct {
	*WidgetOptions
	theme *ui.Theme
	opts  []widget.ContainerOpt
}

func NewContainerOptions(t *ui.Theme) *ContainerOptions {
	return &ContainerOptions{
		WidgetOptions: NewWidgetOptions(t),
		theme:         t,
	}
}

func (c *ContainerOptions) WithBackground() {
	c.opts = append(
		c.opts,
		widget.ContainerOpts.BackgroundImage(c.theme.SurfaceImage),
	)
}

func (c *ContainerOptions) GetOptions() []widget.ContainerOpt {
	return append(
		c.opts,
		widget.ContainerOpts.WidgetOpts(c.WidgetOptions.GetOptions()...),
	)
}

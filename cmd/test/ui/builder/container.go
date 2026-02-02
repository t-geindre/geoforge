package builder

import (
	"geoforge/cmd/test/ui"
	"geoforge/cmd/test/ui/options"

	"github.com/ebitenui/ebitenui/widget"
)

type ContainerBuilder struct {
	*options.ContainerOptions
}

func NewContainerBuilder(b *ui.Builder) *ContainerBuilder {
	return &ContainerBuilder{
		ContainerOptions: options.NewContainerOptions(b.GetTheme()),
	}
}

func (cb *ContainerBuilder) GetWidget() *widget.Container {
	return widget.NewContainer(cb.GetOptions()...)
}

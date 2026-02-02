package options

import (
	"geoforge/cmd/test/ui"

	"github.com/ebitenui/ebitenui/widget"
)

type WidgetOptions struct {
	opts  []widget.WidgetOpt
	theme *ui.Theme
}

func NewWidgetOptions(t *ui.Theme) *WidgetOptions {
	return &WidgetOptions{
		theme: t,
	}
}

func (w *WidgetOptions) WithMinSize(width, height int) {
	w.opts = append(
		w.opts,
		widget.WidgetOpts.MinSize(width, height),
	)
}

func (w *WidgetOptions) GetOptions() []widget.WidgetOpt {
	return w.opts
}

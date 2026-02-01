package main

import (
	"github.com/ebitenui/ebitenui/widget"
)

type ContainerId uint8

type Builder struct {
	theme *Theme
}

func NewBuilder(theme *Theme) *Builder {
	return &Builder{
		theme: theme,
	}
}

func (b *Builder) NewButton(label string, opts ...widget.ButtonOpt) *widget.Button {
	return widget.NewButton(
		append([]widget.ButtonOpt{
			widget.ButtonOpts.Text(label, b.theme.WidgetsFont, &widget.ButtonTextColor{
				Idle:     b.theme.FgIdle,
				Hover:    b.theme.FgHover,
				Pressed:  b.theme.FgPressed,
				Disabled: b.theme.FgDisabled,
			}),
			widget.ButtonOpts.Image(b.theme.BtnImage),
			widget.ButtonOpts.TextPadding(b.theme.WidgetsInset),
		}, opts...)...,
	)
}

func (b *Builder) NewText(label string, opts ...widget.TextOpt) *widget.Text {
	return widget.NewText(
		append([]widget.TextOpt{
			widget.TextOpts.Text(label, b.theme.WidgetsFont, b.theme.FgIdle),
			widget.TextOpts.Padding(b.theme.WidgetsInset),
		}, opts...)...,
	)
}

func (b *Builder) NewStatusBar(opts ...widget.ContainerOpt) *widget.Container {
	return widget.NewContainer(
		append([]widget.ContainerOpt{
			widget.ContainerOpts.BackgroundImage(b.theme.SidebarImage),
			widget.ContainerOpts.WidgetOpts(
				widget.WidgetOpts.MinSize(0, b.theme.SidebarHeight),
			),
			widget.ContainerOpts.Layout(
				widget.NewGridLayout(
					widget.GridLayoutOpts.Columns(2),
					widget.GridLayoutOpts.Stretch([]bool{true, false}, []bool{true}),
					widget.GridLayoutOpts.Padding(b.theme.SurfaceInset),
				),
			),
		}, opts...)...,
	)
}

func (b *Builder) NewMenuBar(opts ...widget.ContainerOpt) *widget.Container {
	return widget.NewContainer(
		append([]widget.ContainerOpt{
			widget.ContainerOpts.BackgroundImage(b.theme.SidebarImage),
			widget.ContainerOpts.WidgetOpts(
				widget.WidgetOpts.MinSize(0, b.theme.SidebarHeight),
			),
			widget.ContainerOpts.Layout(
				widget.NewRowLayout(
					widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
					widget.RowLayoutOpts.Spacing(8),
					widget.RowLayoutOpts.Padding(b.theme.SurfaceInset),
				),
			),
		}, opts...)...,
	)
}

func (b *Builder) NewMainContainer(opts ...widget.ContainerOpt) *widget.Container {
	return widget.NewContainer(
		append([]widget.ContainerOpt{
			widget.ContainerOpts.Layout(
				widget.NewGridLayout(
					widget.GridLayoutOpts.Columns(1),
					widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, true, false}),
				),
			),
		}, opts...)...,
	)
}

func (b *Builder) NewContainer(opts ...widget.ContainerOpt) *widget.Container {
	return widget.NewContainer(
		append([]widget.ContainerOpt{
			widget.ContainerOpts.BackgroundImage(b.theme.SurfaceImage),
		}, opts...)...,
	)
}

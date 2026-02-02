package main

import (
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/colornames"
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

func (b *Builder) NewGrid() *Grid {
	return newGrid(b)
}

func (b *Builder) NewButton(label string, opts ...widget.ButtonOpt) *widget.Button {
	return widget.NewButton(
		append([]widget.ButtonOpt{
			widget.ButtonOpts.Text(label, b.theme.WidgetsFont, &widget.ButtonTextColor{
				Idle:     b.theme.FgIdleColor,
				Hover:    b.theme.FgHoverColor,
				Pressed:  b.theme.FgPressedColor,
				Disabled: b.theme.FgDisabledColor,
			}),
			widget.ButtonOpts.Image(b.theme.BtnImage),
			widget.ButtonOpts.TextPadding(b.theme.WidgetsInset),
		}, opts...)...,
	)
}

func (b *Builder) NewText(label string, opts ...widget.TextOpt) *widget.Text {
	return widget.NewText(
		append([]widget.TextOpt{
			widget.TextOpts.Text(label, b.theme.WidgetsFont, b.theme.FgIdleColor),
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

func (b *Builder) NewInvisibleContainer(opts ...widget.ContainerOpt) *widget.Container {
	return widget.NewContainer(
		append([]widget.ContainerOpt{}, opts...)...,
	)
}

func (b *Builder) NewGridContainer(c int, opts ...widget.ContainerOpt) *widget.Container {
	return widget.NewContainer(
		append([]widget.ContainerOpt{
			widget.ContainerOpts.Layout(
				widget.NewGridLayout(
					widget.GridLayoutOpts.Columns(c),
					widget.GridLayoutOpts.Stretch([]bool{false, true}, []bool{false}),
					widget.GridLayoutOpts.Padding(b.theme.SurfaceInset),
					widget.GridLayoutOpts.Spacing(b.theme.SurfaceSpaceX, b.theme.SurfaceSpaceY),
				),
			),
		}, opts...)...,
	)
}

func (b *Builder) NewDesktop(ww, wh int) *Desktop {
	return NewDesktop(ww, wh)
}

func (b *Builder) NewDraggable(wx, wy int) *Draggable {
	return NewDraggable(wx, wy, b)
}

func (b *Builder) NewSlider(opts ...widget.SliderOpt) *widget.Slider {
	return widget.NewSlider(
		append([]widget.SliderOpt{
			widget.SliderOpts.MinMax(0, 10),
			widget.SliderOpts.InitialCurrent(5),
			widget.SliderOpts.Images(
				&widget.SliderTrackImage{
					Idle: b.theme.BgIdleImage,
				},
				&widget.ButtonImage{
					Idle:    b.theme.FgIdleImage,
					Hover:   b.theme.FgHoverImage,
					Pressed: b.theme.FgPressedImage,
				},
			),
			widget.SliderOpts.FixedHandleSize(6),
			widget.SliderOpts.TrackOffset(0),
			widget.SliderOpts.PageSizeFunc(func() int {
				return 1
			}),
		}, opts...)...,
	)
}

func (b *Builder) NewCheckbox(opts ...widget.CheckboxOpt) *widget.Checkbox {
	return widget.NewCheckbox(
		append([]widget.CheckboxOpt{
			widget.CheckboxOpts.Image(
				&widget.CheckboxImage{
					Unchecked:         b.theme.BgIdleImage,
					UncheckedDisabled: b.theme.BgDisabledImage,
					Checked:           b.theme.FgIdleImage,
					CheckedHovered:    b.theme.FgHoverImage,
				},
			),
			widget.CheckboxOpts.WidgetOpts(
				widget.WidgetOpts.MinSize(16, 16),
			),
		}, opts...)...,
	)
}

func (b *Builder) NewCheckboxWithLabel(on, off string) *widget.Container {
	var text *widget.Text
	checkbox := b.NewCheckbox(
		widget.CheckboxOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
		widget.CheckboxOpts.StateChangedHandler(func(args *widget.CheckboxChangedEventArgs) {
			if args.State == widget.WidgetChecked {
				text.Label = on
			} else {
				text.Label = off
			}
		}),
	)
	text = b.NewText(
		off,
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
			widget.WidgetOpts.MouseButtonPressedHandler(func(args *widget.WidgetMouseButtonPressedEventArgs) {
				checkbox.Click()
			}),
		),
	)

	container := b.NewInvisibleContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			),
		),
	)
	container.AddChild(checkbox)
	container.AddChild(text)

	return container
}

func (b *Builder) NewTextInput(opts ...widget.TextInputOpt) *widget.TextInput {
	return widget.NewTextInput(
		append([]widget.TextInputOpt{
			widget.TextInputOpts.Image(&widget.TextInputImage{
				Idle: b.theme.BgIdleImage,
			}),
			widget.TextInputOpts.Face(b.theme.WidgetsFont),
			widget.TextInputOpts.Color(&widget.TextInputColor{
				Idle:  b.theme.FgIdleColor,
				Caret: b.theme.FgHoverColor,
			}),
			widget.TextInputOpts.Padding(b.theme.WidgetsInset),
		}, opts...)...,
	)
}

func (b *Builder) NewSeparator(opts ...widget.ContainerOpt) *widget.Container {
	return widget.NewContainer(
		append([]widget.ContainerOpt{
			widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(colornames.White)),
			widget.ContainerOpts.WidgetOpts(
				widget.WidgetOpts.MinSize(0, 1),
			),
		}, opts...)...,
	)
}

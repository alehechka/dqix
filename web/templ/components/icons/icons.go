package icons

import "github.com/a-h/templ"

type IconConfig struct {
	Stroke      string
	Fill        string
	StrokeWidth string
	ViewBox     string
	Height      string
	Width       string
	Xmlns       string
	Class       templ.CSSClass
}

var Default IconConfig = IconConfig{
	Stroke:      "currentColor",
	Fill:        "currentColor",
	StrokeWidth: "0",
	ViewBox:     "0 0 16 16",
	Height:      "1em",
	Width:       "1em",
	Xmlns:       "http://www.w3.org/2000/svg",
	Class:       templ.Class(""),
}

type IconOption interface {
	apply(*IconConfig)
}

func newConfig(options ...IconOption) IconConfig {
	config := Default

	for _, option := range options {
		option.apply(&config)
	}

	return config
}

type IconFunc func(config IconConfig) templ.Component

func icon(options []IconOption, defaults ...IconOption) templ.Component {
	config := newConfig(append(defaults, options...)...)

	return svg(config)
}

type withSizeOption struct {
	size string
}

func WithSize(size string) IconOption {
	return &withSizeOption{size: size}
}

func (o withSizeOption) apply(config *IconConfig) {
	config.Height = o.size
	config.Width = o.size
}

type withStrokeOption struct {
	stroke string
}

func WithStroke(stroke string) IconOption {
	return &withStrokeOption{stroke: stroke}
}

func (o withStrokeOption) apply(config *IconConfig) {
	config.Stroke = o.stroke
}

type withFillOption struct {
	fill string
}

func WithFill(fill string) IconOption {
	return &withFillOption{fill: fill}
}

func (o withFillOption) apply(config *IconConfig) {
	config.Fill = o.fill
}

type withStrokeWidthOption struct {
	strokeWidth string
}

func WithStrokeWidth(strokeWidth string) IconOption {
	return &withStrokeWidthOption{strokeWidth: strokeWidth}
}

func (o withStrokeWidthOption) apply(config *IconConfig) {
	config.StrokeWidth = o.strokeWidth
}

type withViewBoxOption struct {
	viewBox string
}

func WithViewBox(viewBox string) IconOption {
	return &withViewBoxOption{viewBox: viewBox}
}

func (o withViewBoxOption) apply(config *IconConfig) {
	config.ViewBox = o.viewBox
}

type withXmlnsOption struct {
	xmlns string
}

func WithXmlns(xmlns string) IconOption {
	return &withXmlnsOption{xmlns: xmlns}
}

func (o withXmlnsOption) apply(config *IconConfig) {
	config.Xmlns = o.xmlns
}

type withClassOption struct {
	class templ.CSSClass
}

func WithClass(class string) IconOption {
	return &withClassOption{class: templ.Class(class)}
}

func WithSafeClass(class string) IconOption {
	return &withClassOption{class: templ.SafeClass(class)}
}

func (o withClassOption) apply(config *IconConfig) {
	config.Class = o.class
}

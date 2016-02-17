package impl

import (
	cairo_impl "github.com/gotk3/gotk3/cairo/impl"
	"github.com/gotk3/gotk3/pango"
)

import "github.com/gotk3/gotk3/cairo"

type RealPango struct{}

var Real = &RealPango{}

func (*RealPango) CairoCreateContext(cr cairo.Context) pango.Context {
	return CairoCreateContext(cairo_impl.CastToContext(cr))
}

func (*RealPango) CairoCreateLayout(cr cairo.Context) pango.Layout {
	return CairoCreateLayout(cairo_impl.CastToContext(cr))
}

func (*RealPango) CairoErrorUnderlinePath(cr cairo.Context, x float64, y float64, width float64, height float64) {
	CairoErrorUnderlinePath(cairo_impl.CastToContext(cr), x, y, width, height)
}

func (*RealPango) CairoGlyphStringPath(cr cairo.Context, font pango.Font, glyphs pango.GlyphString) {
	CairoGlyphStringPath(cairo_impl.CastToContext(cr), toFont(font), toGlyphString(glyphs))
}

func (*RealPango) CairoLayoutLinePath(cr cairo.Context, line pango.LayoutLine) {
	CairoLayoutLinePath(cairo_impl.CastToContext(cr), toLayoutLine(line))
}

func (*RealPango) CairoLayoutPath(cr cairo.Context, layout pango.Layout) {
	CairoLayoutPath(cairo_impl.CastToContext(cr), toLayout(layout))
}

func (*RealPango) CairoShowGlyphItem(cr cairo.Context, text string, glyph_item pango.GlyphItem) {
	CairoShowGlyphItem(cairo_impl.CastToContext(cr), text, toGlyphItem(glyph_item))
}

func (*RealPango) CairoShowGlyphString(cr cairo.Context, font pango.Font, glyphs pango.GlyphString) {
	CairoShowGlyphString(cairo_impl.CastToContext(cr), toFont(font), toGlyphString(glyphs))
}

func (*RealPango) CairoShowLayout(cr cairo.Context, layout pango.Layout) {
	CairoShowLayout(cairo_impl.CastToContext(cr), toLayout(layout))
}

func (*RealPango) CairoShowLayoutLine(cr cairo.Context, line pango.LayoutLine) {
	CairoShowLayoutLine(cairo_impl.CastToContext(cr), toLayoutLine(line))
}

func (*RealPango) CairoUpdateLayout(cr cairo.Context, v pango.Layout) {
	CairoUpdateLayout(cairo_impl.CastToContext(cr), toLayout(v))
}

func (*RealPango) ContextNew() pango.Context {
	return ContextNew()
}

func (*RealPango) FontDescriptionFromString(str string) pango.FontDescription {
	return FontDescriptionFromString(str)
}

func (*RealPango) FontDescriptionNew() pango.FontDescription {
	return FontDescriptionNew()
}

func (*RealPango) GravityToRotation(gravity pango.Gravity) float64 {
	return GravityToRotation(gravity)
}

func (*RealPango) LayoutNew(context pango.Context) pango.Layout {
	return LayoutNew(toContext(context))
}

func (*RealPango) RectangleNew(x int, y int, width int, height int) pango.Rectangle {
	return RectangleNew(x, y, width, height)
}

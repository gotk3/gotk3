package impl

import (
	cairo_impl "github.com/gotk3/gotk3/cairo/impl"
	"github.com/gotk3/gotk3/pango"
)

import "github.com/gotk3/gotk3/cairo"

type RealPango struct{}

var Real = &RealPango{}

func (*RealPango) CairoCreateContext(cr cairo.Context) pango.Context {
	return CairoCreateContext(cr.(*cairo_impl.Context))
}

func (*RealPango) CairoCreateLayout(cr cairo.Context) pango.Layout {
	return CairoCreateLayout(cr.(*cairo_impl.Context))
}

func (*RealPango) CairoErrorUnderlinePath(cr cairo.Context, x float64, y float64, width float64, height float64) {
	CairoErrorUnderlinePath(cr.(*cairo_impl.Context), x, y, width, height)
}

func (*RealPango) CairoGlyphStringPath(cr cairo.Context, font pango.Font, glyphs pango.GlyphString) {
	CairoGlyphStringPath(cr.(*cairo_impl.Context), font.(*Font), glyphs.(*GlyphString))
}

func (*RealPango) CairoLayoutLinePath(cr cairo.Context, line pango.LayoutLine) {
	CairoLayoutLinePath(cr.(*cairo_impl.Context), line.(*LayoutLine))
}

func (*RealPango) CairoLayoutPath(cr cairo.Context, layout pango.Layout) {
	CairoLayoutPath(cr.(*cairo_impl.Context), layout.(*Layout))
}

func (*RealPango) CairoShowGlyphItem(cr cairo.Context, text string, glyph_item pango.GlyphItem) {
	CairoShowGlyphItem(cr.(*cairo_impl.Context), text, glyph_item.(*GlyphItem))
}

func (*RealPango) CairoShowGlyphString(cr cairo.Context, font pango.Font, glyphs pango.GlyphString) {
	CairoShowGlyphString(cr.(*cairo_impl.Context), font.(*Font), glyphs.(*GlyphString))
}

func (*RealPango) CairoShowLayout(cr cairo.Context, layout pango.Layout) {
	CairoShowLayout(cr.(*cairo_impl.Context), layout.(*Layout))
}

func (*RealPango) CairoShowLayoutLine(cr cairo.Context, line pango.LayoutLine) {
	CairoShowLayoutLine(cr.(*cairo_impl.Context), line.(*LayoutLine))
}

func (*RealPango) CairoUpdateLayout(cr cairo.Context, v pango.Layout) {
	CairoUpdateLayout(cr.(*cairo_impl.Context), v.(*Layout))
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
	return LayoutNew(context.(*Context))
}

func (*RealPango) RectangleNew(x int, y int, width int, height int) pango.Rectangle {
	return RectangleNew(x, y, width, height)
}

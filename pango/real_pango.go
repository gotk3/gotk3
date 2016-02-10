package pango

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/pango/iface"
)
import cairo_iface "github.com/gotk3/gotk3/cairo/iface"

type RealPango struct{}

var Real = &RealPango{}

func (*RealPango) CairoCreateContext(cr cairo_iface.Context) iface.Context {
	return CairoCreateContext(cr.(*cairo.Context))
}

func (*RealPango) CairoCreateLayout(cr cairo_iface.Context) iface.Layout {
	return CairoCreateLayout(cr.(*cairo.Context))
}

func (*RealPango) CairoErrorUnderlinePath(cr cairo_iface.Context, x float64, y float64, width float64, height float64) {
	CairoErrorUnderlinePath(cr.(*cairo.Context), x, y, width, height)
}

func (*RealPango) CairoGlyphStringPath(cr cairo_iface.Context, font iface.Font, glyphs iface.GlyphString) {
	CairoGlyphStringPath(cr.(*cairo.Context), font.(*Font), glyphs.(*GlyphString))
}

func (*RealPango) CairoLayoutLinePath(cr cairo_iface.Context, line iface.LayoutLine) {
	CairoLayoutLinePath(cr.(*cairo.Context), line.(*LayoutLine))
}

func (*RealPango) CairoLayoutPath(cr cairo_iface.Context, layout iface.Layout) {
	CairoLayoutPath(cr.(*cairo.Context), layout.(*Layout))
}

func (*RealPango) CairoShowGlyphItem(cr cairo_iface.Context, text string, glyph_item iface.GlyphItem) {
	CairoShowGlyphItem(cr.(*cairo.Context), text, glyph_item.(*GlyphItem))
}

func (*RealPango) CairoShowGlyphString(cr cairo_iface.Context, font iface.Font, glyphs iface.GlyphString) {
	CairoShowGlyphString(cr.(*cairo.Context), font.(*Font), glyphs.(*GlyphString))
}

func (*RealPango) CairoShowLayout(cr cairo_iface.Context, layout iface.Layout) {
	CairoShowLayout(cr.(*cairo.Context), layout.(*Layout))
}

func (*RealPango) CairoShowLayoutLine(cr cairo_iface.Context, line iface.LayoutLine) {
	CairoShowLayoutLine(cr.(*cairo.Context), line.(*LayoutLine))
}

func (*RealPango) CairoUpdateLayout(cr cairo_iface.Context, v iface.Layout) {
	CairoUpdateLayout(cr.(*cairo.Context), v.(*Layout))
}

func (*RealPango) ContextNew() iface.Context {
	return ContextNew()
}

func (*RealPango) FontDescriptionFromString(str string) iface.FontDescription {
	return FontDescriptionFromString(str)
}

func (*RealPango) FontDescriptionNew() iface.FontDescription {
	return FontDescriptionNew()
}

func (*RealPango) GravityToRotation(gravity iface.Gravity) float64 {
	return GravityToRotation(gravity)
}

func (*RealPango) LayoutNew(context iface.Context) iface.Layout {
	return LayoutNew(context.(*Context))
}

func (*RealPango) RectangleNew(x int, y int, width int, height int) iface.Rectangle {
	return RectangleNew(x, y, width, height)
}

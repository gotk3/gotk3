package pango

import "github.com/gotk3/gotk3/cairo"

type RealPango interface {
	CairoCreateContext(cairo.Context) Context
	CairoCreateLayout(cairo.Context) Layout
	CairoErrorUnderlinePath(cairo.Context, float64, float64, float64, float64)
	CairoGlyphStringPath(cairo.Context, Font, GlyphString)
	CairoLayoutLinePath(cairo.Context, LayoutLine)
	CairoLayoutPath(cairo.Context, Layout)
	CairoShowGlyphItem(cairo.Context, string, GlyphItem)
	CairoShowGlyphString(cairo.Context, Font, GlyphString)
	CairoShowLayout(cairo.Context, Layout)
	CairoShowLayoutLine(cairo.Context, LayoutLine)
	CairoUpdateLayout(cairo.Context, Layout)
	ContextNew() Context
	FontDescriptionFromString(string) FontDescription
	FontDescriptionNew() FontDescription
	GravityToRotation(Gravity) float64
	LayoutNew(Context) Layout
	RectangleNew(int, int, int, int) Rectangle
} // end of RealPango

func AssertRealPango(_ RealPango) {}

package iface

import cairo_iface "github.com/gotk3/gotk3/cairo/iface"

type RealPango interface {
    CairoCreateContext(cairo_iface.Context) Context
    CairoCreateLayout(cairo_iface.Context) Layout
    CairoErrorUnderlinePath(cairo_iface.Context, float64, float64, float64, float64)
    CairoGlyphStringPath(cairo_iface.Context, Font, GlyphString)
    CairoLayoutLinePath(cairo_iface.Context, LayoutLine)
    CairoLayoutPath(cairo_iface.Context, Layout)
    CairoShowGlyphItem(cairo_iface.Context, string, GlyphItem)
    CairoShowGlyphString(cairo_iface.Context, Font, GlyphString)
    CairoShowLayout(cairo_iface.Context, Layout)
    CairoShowLayoutLine(cairo_iface.Context, LayoutLine)
    CairoUpdateLayout(cairo_iface.Context, Layout)
    ContextNew() Context
    FontDescriptionFromString(string) FontDescription
    FontDescriptionNew() FontDescription
    GravityToRotation(Gravity) float64
    LayoutNew(Context) Layout
    RectangleNew(int, int, int, int) Rectangle
} // end of RealPango

func AssertRealPango(_ RealPango) {}

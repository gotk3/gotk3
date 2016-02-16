package impl

import "github.com/gotk3/gotk3/cairo"

func init() {
	cairo.AssertCairo(&RealCairo{})
	cairo.AssertContext(&Context{})
	cairo.AssertSurface(&surface{})
}

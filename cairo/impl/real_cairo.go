package impl

import "github.com/gotk3/gotk3/cairo"

type RealCairo struct{}

var Real = &RealCairo{}

func (*RealCairo) Create(target cairo.Surface) cairo.Context {
	return Create(target.(*Surface))
}

func (*RealCairo) NewSurface(s uintptr, needsRef bool) cairo.Surface {
	return NewSurface(s, needsRef)
}

func (*RealCairo) NewSurfaceFromPNG(fileName string) (cairo.Surface, error) {
	return NewSurfaceFromPNG(fileName)
}

func (*RealCairo) StatusToString(status cairo.Status) string {
	return StatusToString(status)
}

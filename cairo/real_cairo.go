package cairo

import "github.com/gotk3/gotk3/cairo/iface"

type RealCairo struct{}

var Real = &RealCairo{}

func (*RealCairo) Create(target iface.Surface) iface.Context {
	return Create(target.(*Surface))
}

func (*RealCairo) NewSurface(s uintptr, needsRef bool) iface.Surface {
	return NewSurface(s, needsRef)
}

func (*RealCairo) NewSurfaceFromPNG(fileName string) (iface.Surface, error) {
	return NewSurfaceFromPNG(fileName)
}

func (*RealCairo) StatusToString(status iface.Status) string {
	return StatusToString(status)
}

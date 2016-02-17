package impl

import "github.com/gotk3/gotk3/cairo"

func toSurface(s cairo.Surface) *Surface {
	if s == nil {
		return nil
	}
	return s.(*Surface)
}

func CastToContext(s cairo.Context) *Context {
	if s == nil {
		return nil
	}
	return s.(*Context)
}

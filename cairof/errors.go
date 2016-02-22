package cairof

import "github.com/gotk3/gotk3/cairo"

type ErrorStatus cairo.Status

func (e ErrorStatus) Error() string {
	return StatusToString(cairo.Status(e))
}

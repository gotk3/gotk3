package cairo

import "github.com/gotk3/gotk3/cairo/iface"

type ErrorStatus iface.Status

func (e ErrorStatus) Error() string {
	return StatusToString(iface.Status(e))
}

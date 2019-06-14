package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
// #include "fixed.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_fixed_get_type()), marshalFixed},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkFixed"] = wrapFixed
}

/*
 * GtkFixed
 */

// Fixed is a representation of GTK's GtkFixed.
type Fixed struct {
	Container
}

func (v *Fixed) native() *C.GtkFixed {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFixed(p)
}

func marshalFixed(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFixed(obj), nil
}

func wrapFixed(obj *glib.Object) *Fixed {
	return &Fixed{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// FixedNew is a wrapper around gtk_fixed_new().
func FixedNew() (*Fixed, error) {
	c := C.gtk_fixed_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFixed(obj), nil
}

// Put is a wrapper around gtk_fixed_put().
func (v *Fixed) Put(w IWidget, x, y int) {
	C.gtk_fixed_put(v.native(), w.toWidget(), C.gint(x), C.gint(y))
}

// Move is a wrapper around gtk_fixed_move().
func (v *Fixed) Move(w IWidget, x, y int) {
	C.gtk_fixed_move(v.native(), w.toWidget(), C.gint(x), C.gint(y))
}

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_info_bar_get_type()), marshalInfoBar},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkInfoBar"] = wrapInfoBar
}

type InfoBar struct {
	Box
}

func (v *InfoBar) native() *C.GtkInfoBar {
	if v == nil || v.GObject == nil {
		return nil
	}

	p := unsafe.Pointer(v.GObject)
	return C.toGtkInfoBar(p)
}

func marshalInfoBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapInfoBar(wrapObject(unsafe.Pointer(c))), nil
}

func wrapInfoBar(obj *glib.Object) *InfoBar {
	return &InfoBar{Box{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

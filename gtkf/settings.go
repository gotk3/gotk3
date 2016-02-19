package gtkf

// #include <gtk/gtk.h>
// #include "settings.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	glib_impl "github.com/gotk3/gotk3/glibf"
)

func init() {
	tm := []glib_impl.TypeMarshaler{
		{glib.Type(C.gtk_settings_get_type()), marshalSettings},
	}

	glib_impl.RegisterGValueMarshalers(tm)

	WrapMap["GtkSettings"] = wrapSettings
}

//GtkSettings
type settings struct {
	*glib_impl.Object
}

func (v *settings) native() *C.GtkSettings {
	if v == nil || v.GObject == nil {
		return nil
	}

	p := unsafe.Pointer(v.GObject)
	return C.toGtkSettings(p)
}

func marshalSettings(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapSettings(wrapObject(unsafe.Pointer(c))), nil
}

func wrapSettings(obj *glib_impl.Object) *settings {
	return &settings{obj}
}

//Get the global non window specific settings
func SettingsGetDefault() (*settings, error) {
	c := C.gtk_settings_get_default()
	if c == nil {
		return nil, nilPtrErr
	}

	return wrapSettings(wrapObject(unsafe.Pointer(c))), nil
}

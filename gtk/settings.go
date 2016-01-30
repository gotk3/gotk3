package gtk

// #include <gtk/gtk.h>
// #include "settings.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_settings_get_type()), marshalSettings},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkSettings"] = SettingsWrap
}

//GtkSettings
type Settings struct {
	*glib.Object
}

func (v *Settings) native() *C.GtkSettings {
	if v == nil || v.GObject == nil {
		return nil
	}

	p := unsafe.Pointer(v.GObject)
	return C.toGtkSettings(p)
}

func marshalSettings(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return SettingsWrap(wrapObject(unsafe.Pointer(c))), nil
}

func SettingsWrap(obj *glib.Object) *Settings {
	return &Settings{obj}
}

//Get the global non window specific settings
func SettingsGetDefault() (*Settings, error) {
	c := C.gtk_settings_get_default()
	if c == nil {
		return nil, nilPtrErr
	}

	return SettingsWrap(wrapObject(unsafe.Pointer(c))), nil
}

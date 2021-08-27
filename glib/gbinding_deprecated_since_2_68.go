// +build glib_2_40 glib_2_42 glib_2_44 glib_2_46 glib_2_48 glib_2_50 glib_2_52 glib_2_54 glib_2_56 glib_2_58 glib_2_60 glib_2_62 glib_2_64 glib_2_66 gtk_deprecated

package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

// GetSource is a wrapper around g_binding_get_source().
func (v *Binding) GetSource() *Object {
	obj := C.g_binding_get_source(v.native())
	if obj == nil {
		return nil
	}
	return wrapObject(unsafe.Pointer(obj))
}

// GetTarget is a wrapper around g_binding_get_target().
func (v *Binding) GetTarget() *Object {
	obj := C.g_binding_get_target(v.native())
	if obj == nil {
		return nil
	}
	return wrapObject(unsafe.Pointer(obj))
}

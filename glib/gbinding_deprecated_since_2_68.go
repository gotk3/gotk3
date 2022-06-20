// +build glib_deprecated

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

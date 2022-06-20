// +build !glib_deprecated

package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

// DupSource is a wrapper around g_binding_dup_source().
func (v *Binding) DupSource() *Object {
	obj := C.g_binding_dup_source(v.native())
	if obj == nil {
		return nil
	}
	return wrapObject(unsafe.Pointer(obj))
}

// DupTarget is a wrapper around g_binding_dup_target().
func (v *Binding) DupTarget() *Object {
	obj := C.g_binding_dup_target(v.native())
	if obj == nil {
		return nil
	}
	return wrapObject(unsafe.Pointer(obj))
}

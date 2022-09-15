//go:build !glib_deprecated && !glib_2_40 && !glib_2_42 && !glib_2_44 && !glib_2_46 && !glib_2_48 && !glib_2_50 && !glib_2_52 && !glib_2_54 && !glib_2_56 && !glib_2_58 && !glib_2_60 && !glib_2_62 && !glib_2_64 && !glib_2_66
// +build !glib_deprecated,!glib_2_40,!glib_2_42,!glib_2_44,!glib_2_46,!glib_2_48,!glib_2_50,!glib_2_52,!glib_2_54,!glib_2_56,!glib_2_58,!glib_2_60,!glib_2_62,!glib_2_64,!glib_2_66

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

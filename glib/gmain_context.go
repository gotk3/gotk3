package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"

type MainContext C.GMainContext

// native returns a pointer to the underlying GMainContext.
func (v *MainContext) native() *C.GMainContext {
	if v == nil {
		return nil
	}
	return (*C.GMainContext)(v)
}

// MainContextDefault is a wrapper around g_main_context_default().
func MainContextDefault() *MainContext {
	c := C.g_main_context_default()
	if c == nil {
		return nil
	}
	return (*MainContext)(c)
}

// Iteration is a wrapper around g_main_context_iteration()
func (v *MainContext) Iteration(mayBlock bool) bool {
	return gobool(C.g_main_context_iteration(v.native(), gbool(mayBlock)))
}

// Pending is a wrapper around g_main_context_pending()
func (v *MainContext) Pending() bool {
	return gobool(C.g_main_context_pending(v.native()))
}

// MainDepth is a wrapper around g_main_depth().
func MainDepth() int {
	return int(C.g_main_depth())
}

// FindSourceById is a wrapper around g_main_context_find_source_by_id()
func (v *MainContext) FindSourceById(hdlSrc SourceHandle) *Source {
	c := C.g_main_context_find_source_by_id(v.native(), C.guint(hdlSrc))
	if c == nil {
		return nil
	}
	return (*Source)(c)
}

// Acquire is a wrapper around g_main_context_acquire().
func (v *MainContext) Acquire() bool {
	return gobool(C.g_main_context_acquire(v.native()))
}

// Release is a wrapper around g_main_context_release().
func (v *MainContext) Release() {
	C.g_main_context_release(v.native())
}

// IsOwner is a wrapper around g_main_context_is_owner().
func (v *MainContext) IsOwner() bool {
	return gobool(C.g_main_context_is_owner(v.native()))
}

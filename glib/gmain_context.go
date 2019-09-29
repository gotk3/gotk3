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

// SourceRemove is a wrapper around g_main_context_pending()
func (v *MainContext) Pending(src SourceHandle) bool {
	return gobool(C.g_main_context_pending(v.native()))
}

package glib

import (
	"unsafe"
)

// #include <glib.h>
// #include <stdlib.h>
import "C"

// QuarkFromString is a wrapper around g_quark_from_string().
func QuarkFromString(str string) Quark {
	cstr := (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(cstr))

	return Quark(C.g_quark_from_string(cstr))
}

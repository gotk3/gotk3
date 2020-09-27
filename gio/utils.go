package gio

// #include <glib.h>
// #include "gresource.go.h"
import "C"

// same implementation as package glib
func toGoStringArray(c **C.char) []string {
	var strs []string
	originalc := c
	defer C.char_g_strfreev(originalc)

	for *c != nil {
		strs = append(strs, C.GoString((*C.char)(*c)))
		c = C.next_charptr(c)
	}

	return strs
}

func goString(cstr *C.gchar) string {
	return C.GoString((*C.char)(cstr))
}

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

func gobool(b C.gboolean) bool {
	return b != C.FALSE
}

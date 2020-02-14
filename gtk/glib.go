package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func nativeGPermission(permission *glib.Permission) *C.GPermission {
	// Note: would return C type prefixed with glib package.
	// Go issue: here https://github.com/golang/go/issues/13467.
	var perm *C.GPermission
	if permission != nil {
		perm = (*C.GPermission)(unsafe.Pointer(permission.Native()))
	}
	return perm
}

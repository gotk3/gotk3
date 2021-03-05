package glib

// #cgo pkg-config: gio-2.0
// #include <gio/gio.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/internal/callback"
)

//export goAsyncReadyCallbacks
func goAsyncReadyCallbacks(sourceObject *C.GObject, res *C.GAsyncResult, userData C.gpointer) {
	var source *Object
	if sourceObject != nil {
		source = wrapObject(unsafe.Pointer(sourceObject))
	}

	fn := callback.Get(uintptr(userData)).(AsyncReadyCallback)
	fn(source, wrapAsyncResult(wrapObject(unsafe.Pointer(res))))
}

//export goCompareDataFuncs
func goCompareDataFuncs(a, b C.gconstpointer, userData C.gpointer) C.gint {
	fn := callback.Get(uintptr(userData)).(CompareDataFunc)
	return C.gint(fn(uintptr(a), uintptr(b)))
}

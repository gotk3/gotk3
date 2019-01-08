// +build !gtk_3_6,!gtk_3_8

package gtk

// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

//export goListBoxFilterFuncs
func goListBoxFilterFuncs(row *C.GtkListBoxRow, userData C.gpointer) C.gboolean {
	id := int(uintptr(userData))

	listBoxFilterFuncRegistry.Lock()
	r := listBoxFilterFuncRegistry.m[id]
	// TODO: figure out a way to determine when we can clean up
	//delete(printSettingsCallbackRegistry.m, id)
	listBoxFilterFuncRegistry.Unlock()

	return gbool(r.fn(wrapListBoxRow(glib.Take(unsafe.Pointer(row))), r.userData))
}

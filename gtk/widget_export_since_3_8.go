// +build !gtk_3_6

package gtk

// #include <gdk/gdk.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
)

//export goTickCallbacks
func goTickCallbacks(widget *C.GtkWidget, frameClock *C.GdkFrameClock, userData C.gpointer) C.gboolean {
	id := int(uintptr(userData))

	tickCallbackRegistry.RLock()
	r := tickCallbackRegistry.m[id]
	tickCallbackRegistry.RUnlock()

	return gbool(r.fn(
		wrapWidget(glib.Take(unsafe.Pointer(widget))),
		gdk.WrapFrameClock(unsafe.Pointer(frameClock)),
		r.userData,
	))
}

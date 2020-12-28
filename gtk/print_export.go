package gtk

// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/internal/callback"
)

//export goPageSetupDone
func goPageSetupDone(setup *C.GtkPageSetup, data C.gpointer) {
	// This callback is only used once, so we can clean up immediately
	fn := callback.GetAndDelete(uintptr(data)).(PageSetupDoneCallback)
	fn(wrapPageSetup(glib.Take(unsafe.Pointer(setup))))
}

//export goPrintSettings
func goPrintSettings(key *C.gchar, value *C.gchar, userData C.gpointer) {
	fn := callback.Get(uintptr(userData)).(PrintSettingsCallback)
	fn(C.GoString((*C.char)(key)), C.GoString((*C.char)(value)))
}

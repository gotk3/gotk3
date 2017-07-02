package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

//export goBuilderConnect
func goBuilderConnect(builder *C.GtkBuilder,
	object *C.GObject,
	signal_name *C.gchar,
	handler_name *C.gchar,
	connect_object *C.GObject,
	flags C.GConnectFlags,
	user_data C.gpointer) {

	builderSignals.Lock()
	signals, ok := builderSignals.m[builder]
	builderSignals.Unlock()

	if !ok {
		panic("no signal mapping defined for this GtkBuilder")
	}

	h := C.GoString((*C.char)(handler_name))
	s := C.GoString((*C.char)(signal_name))

	handler, ok := signals[h]
	if !ok {
		return
	}

	if object == nil {
		panic("unexpected nil object from builder")
	}

	//TODO: figure out a better way to get a glib.Object from a *C.GObject
	gobj := glib.Object{glib.ToGObject(unsafe.Pointer(object))}
	gobj.Connect(s, handler)
}

//export goPageSetupDone
func goPageSetupDone(setup *C.GtkPageSetup,
	data C.gpointer) {

	id := int(uintptr(data))

	pageSetupDoneCallbackRegistry.Lock()
	r := pageSetupDoneCallbackRegistry.m[id]
	delete(pageSetupDoneCallbackRegistry.m, id)
	pageSetupDoneCallbackRegistry.Unlock()

	obj := wrapObject(unsafe.Pointer(setup))
	r.fn(wrapPageSetup(obj), r.data)

}

//export goPrintSettings
func goPrintSettings(key *C.gchar,
	value *C.gchar,
	userData C.gpointer) {

	id := int(uintptr(userData))

	printSettingsCallbackRegistry.Lock()
	r := printSettingsCallbackRegistry.m[id]
	// TODO: figure out a way to determine when we can clean up
	//delete(printSettingsCallbackRegistry.m, id)
	printSettingsCallbackRegistry.Unlock()

	r.fn(C.GoString((*C.char)(key)), C.GoString((*C.char)(value)), r.userData)

}

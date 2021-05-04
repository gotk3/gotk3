package glib

// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/gotk3/gotk3/internal/closure"
)

/*
 * Events
 */

// SignalHandle is the ID of a signal handler.
type SignalHandle uint

// Connect is a wrapper around g_signal_connect_closure(). f must be a function
// with at least one parameter matching the type it is connected to.
//
// It is optional to list the rest of the required types from Gtk, as values
// that don't fit into the function parameter will simply be ignored; however,
// extraneous types will trigger a runtime panic. Arguments for f must be a
// matching Go equivalent type for the C callback, or an interface type which
// the value may be packed in. If the type is not suitable, a runtime panic will
// occur when the signal is emitted.
//
// Circular References
//
// To prevent circular references, prefer declaring Connect functions like so:
//
//    obj.Connect(func(obj *ObjType) { obj.Do() })
//
// Instead of directly referencing the object from outside like so:
//
//    obj.Connect(func() { obj.Do() })
//
// When using Connect, beware of referencing variables outside the closure that
// may cause a circular reference that prevents both Go from garbage collecting
// the callback and GTK from successfully unreferencing its values.
//
// Below is an example piece of code that is considered "leaky":
//
//    type ChatBox struct {
//        gtk.TextView
//        Loader *gdk.PixbufLoader
//
//        State State
//    }
//
//    func (box *ChatBox) Method() {
//        box.Loader.Connect("size-allocate", func(loader *gdk.PixbufLoader) {
//            // Here, we're dereferencing box to get the state, which might
//            // keep box alive along with the PixbufLoader, causing a circular
//            // reference.
//            loader.SetSize(box.State.Width, box.State.Height)
//        })
//    }
//
// There are many solutions to fix the above piece of code. For example,
// box.Loader could be discarded manually immediately after it's done by setting
// it to nil, or the signal handle could be disconnected manually, or box could
// be set to nil after its first call in the callback.
func (v *Object) Connect(detailedSignal string, f interface{}) SignalHandle {
	return v.connectClosure(false, detailedSignal, f)
}

// ConnectAfter is a wrapper around g_signal_connect_closure(). The difference
// between Connect and ConnectAfter is that the latter will be invoked after the
// default handler, not before. For more information, refer to Connect.
func (v *Object) ConnectAfter(detailedSignal string, f interface{}) SignalHandle {
	return v.connectClosure(true, detailedSignal, f)
}

// ClosureCheckReceiver, if true, will make GLib check for every single
// closure's first argument to ensure that it is correct, otherwise it will
// panic with a message warning about the possible circular references. The
// receiver in this case is most often the first argument of the callback.
//
// This constant can be changed by using go.mod's replace directive for
// debugging purposes.
const ClosureCheckReceiver = false

func (v *Object) connectClosure(after bool, detailedSignal string, f interface{}) SignalHandle {
	fs := closure.NewFuncStack(f, 2)

	if ClosureCheckReceiver {
		// This is a bit slow, but we could be careful.
		objValue, err := v.goValue()
		if err == nil {
			fsType := fs.Func.Type()
			if fsType.NumIn() < 1 {
				fs.Panicf("callback should have the object receiver to avoid circular references")
			}
			objType := reflect.TypeOf(objValue)
			if first := fsType.In(0); !objType.ConvertibleTo(first) {
				fs.Panicf("receiver not convertible to expected type %s, got %s", objType, first)
			}
		}

		// Allow the type check to fail if we can't get a value marshaler. This
		// rarely happens, but it might, and we want to at least allow working
		// around it.
	}

	cstr := C.CString(detailedSignal)
	defer C.free(unsafe.Pointer(cstr))

	gclosure := ClosureNewFunc(fs)
	c := C.g_signal_connect_closure(C.gpointer(v.native()), (*C.gchar)(cstr), gclosure, gbool(after))

	// TODO: There's a slight race condition here, where
	// g_signal_connect_closure may trigger signal callbacks before the signal
	// is registered. It is therefore ideal to have another intermediate ID to
	// pass into the connect function. This is not a big issue though, since
	// there isn't really any guarantee that signals should arrive until after
	// the Connect functions return successfully.
	closure.RegisterSignal(uint(c), unsafe.Pointer(gclosure))

	return SignalHandle(c)
}

// ClosureNew creates a new GClosure and adds its callback function to the
// internal registry. It's exported for visibility to other gotk3 packages and
// should not be used in a regular application.
func ClosureNew(f interface{}) *C.GClosure {
	return ClosureNewFunc(closure.NewFuncStack(f, 2))
}

// ClosureNewFunc creates a new GClosure and adds its callback function to the
// internal registry. It's exported for visibility to other gotk3 packages; it
// cannot be used in application code, as package closure is part of the
// internals.
func ClosureNewFunc(funcStack closure.FuncStack) *C.GClosure {
	gclosure := C._g_closure_new()
	closure.Assign(unsafe.Pointer(gclosure), funcStack)

	return gclosure
}

// removeClosure removes a closure from the internal closures map. This is
// needed to prevent a leak where Go code can access the closure context
// (along with rf and userdata) even after an object has been destroyed and
// the GClosure is invalidated and will never run.
//
//export removeClosure
func removeClosure(_ C.gpointer, gclosure *C.GClosure) {
	closure.Delete(unsafe.Pointer(gclosure))
}

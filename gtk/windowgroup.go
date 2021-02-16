// Same copyright and license as the rest of the files in this project

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
)

/*
 * GtkWindowGroup
 */

type WindowGroup struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkWindowGroup.
func (v *WindowGroup) native() *C.GtkWindowGroup {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkWindowGroup(p)
}

func marshalWindowGroup(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWindowGroup(obj), nil
}

func wrapWindowGroup(obj *glib.Object) *WindowGroup {
	if obj == nil {
		return nil
	}

	return &WindowGroup{obj}
}

// WindowGroupNew is a wrapper around gtk_window_group_new().
func WindowGroupNew() (*WindowGroup, error) {
	c := C.gtk_window_group_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWindowGroup(glib.Take(unsafe.Pointer(c))), nil
}

// AddWindow is a wrapper around gtk_window_group_add_window().
func (v *WindowGroup) AddWindow(window IWindow) {
	var pw *C.GtkWindow = nil
	if window != nil {
		pw = window.toWindow()
	}
	C.gtk_window_group_add_window(v.native(), pw)
}

// RemoveWindow is a wrapper around gtk_window_group_remove_window().
func (v *WindowGroup) RemoveWindow(window IWindow) {
	var pw *C.GtkWindow = nil
	if window != nil {
		pw = window.toWindow()
	}
	C.gtk_window_group_remove_window(v.native(), pw)
}

// ListWindows is a wrapper around gtk_window_group_list_windows().
// Returned list is wrapped to return *gtk.Window elements.
// TODO: Use IWindow and wrap to correct type
func (v *WindowGroup) ListWindows() *glib.List {
	clist := C.gtk_window_group_list_windows(v.native())
	if clist == nil {
		return nil
	}
	glist := glib.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		return wrapWindow(glib.Take(ptr))
	})
	runtime.SetFinalizer(glist, func(l *glib.List) {
		l.Free()
	})
	return glist
}

// GetCurrentGrab is a wrapper around gtk_window_group_get_current_grab().
func (v *WindowGroup) GetCurrentGrab() (IWidget, error) {
	c := C.gtk_window_group_get_current_grab(v.native())
	if c == nil {
		return nil, nil
	}
	return castWidget(c)
}

// GetCurrentDeviceGrab is a wrapper around gtk_window_group_get_current_device_grab().
func (v *WindowGroup) GetCurrentDeviceGrab(device *gdk.Device) (IWidget, error) {
	c := C.gtk_window_group_get_current_device_grab(v.native(), C.toGdkDevice(unsafe.Pointer(device.Native())))
	if c == nil {
		return nil, nil
	}
	return castWidget(c)
}

// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18,!gtk_3_20

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
// #include "shortcutswindow_since_3_22.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_shortcuts_window_get_type()), marshalShortcutsWindow},
		{glib.Type(C.gtk_shortcuts_section_get_type()), marshalShortcutsSection},
		{glib.Type(C.gtk_shortcuts_group_get_type()), marshalShortcutsGroup},
		{glib.Type(C.gtk_shortcuts_shortcut_get_type()), marshalShortcutsShortcut},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkShortcutsWindow"] = wrapShortcutsWindow
	WrapMap["GtkShortcutsSection"] = wrapShortcutsSection
	WrapMap["GtkShortcutsGroup"] = wrapShortcutsGroup
	WrapMap["GtkShortcutsShortcut"] = wrapShortcutsShortcut
}

/*
 * GtkShortcutsWindow
 */

// ShortcutsWindow is a representation of GTK's GtkShortcutsWindow.
type ShortcutsWindow struct {
	Window
}

func (v *ShortcutsWindow) native() *C.GtkShortcutsWindow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkShortcutsWindow(p)
}

func marshalShortcutsWindow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapShortcutsWindow(obj), nil
}

func wrapShortcutsWindow(obj *glib.Object) *ShortcutsWindow {
	return &ShortcutsWindow{Window{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}}
}

/*
 * GtkShortcutsSection
 */

// ShortcutsWindow is a representation of GTK's GtkShortcutsSection.
type ShortcutsSection struct {
	Box
}

// native returns a pointer to the underlying GtkShortcutsSection.
func (v *ShortcutsSection) native() *C.GtkShortcutsSection {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkShortcutsSection(p)
}

func marshalShortcutsSection(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapShortcutsSection(obj), nil
}

func wrapShortcutsSection(obj *glib.Object) *ShortcutsSection {
	return &ShortcutsSection{Box{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

/*
 * GtkShortcutsSection
 */

// ShortcutsWindow is a representation of GTK's GtkShortcutsGroup.
type ShortcutsGroup struct {
	Box
}

// native returns a pointer to the underlying GtkShortcutsGroup.
func (v *ShortcutsGroup) native() *C.GtkShortcutsGroup {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkShortcutsGroup(p)
}

func marshalShortcutsGroup(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapShortcutsGroup(obj), nil
}

func wrapShortcutsGroup(obj *glib.Object) *ShortcutsGroup {
	return &ShortcutsGroup{Box{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

/*
 * GtkShortcutsShortcut
 */

// ShortcutsWindow is a representation of GTK's GtkShortcutsShortcut.
type ShortcutsShortcut struct {
	Box
}

// native returns a pointer to the underlying GtkShortcutsShortcut.
func (v *ShortcutsShortcut) native() *C.GtkShortcutsShortcut {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkShortcutsShortcut(p)
}

func marshalShortcutsShortcut(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapShortcutsShortcut(obj), nil
}

func wrapShortcutsShortcut(obj *glib.Object) *ShortcutsShortcut {
	return &ShortcutsShortcut{Box{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

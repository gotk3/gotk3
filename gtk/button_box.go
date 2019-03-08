// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"github.com/gotk3/gotk3/glib"
	"unsafe"
)

type ButtonBoxStyle int

const (
	BUTTONBOX_SPREAD ButtonBoxStyle = C.GTK_BUTTONBOX_SPREAD
	BUTTONBOX_EDGE   ButtonBoxStyle = C.GTK_BUTTONBOX_EDGE
	BUTTONBOX_START  ButtonBoxStyle = C.GTK_BUTTONBOX_START
	BUTTONBOX_END    ButtonBoxStyle = C.GTK_BUTTONBOX_END
	BUTTONBOX_CENTER ButtonBoxStyle = C.GTK_BUTTONBOX_CENTER
)

/*
 * GtkButtonBox
 */

// ButtonBox is a representation of GTK's GtkButtonBox.
type ButtonBox struct {
	Box
}

// native returns a pointer to the underlying GtkButtonBox.
func (v *ButtonBox) native() *C.GtkButtonBox {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkButtonBox(p)
}

func marshalButtonBox(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapButtonBox(obj), nil
}

func wrapButtonBox(obj *glib.Object) *ButtonBox {
	return &ButtonBox{Box{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// ButtonBoxNew is a wrapper around gtk_button_box_new().
func ButtonBoxNew(o Orientation) (*ButtonBox, error) {
	c := C.gtk_button_box_new(C.GtkOrientation(o))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapButtonBox(obj), nil
}

// GetLayout() is a wrapper around gtk_button_box_get_layout().
func (v *ButtonBox) GetLayout() ButtonBoxStyle {
	c := C.gtk_button_box_get_layout(v.native())
	return ButtonBoxStyle(c)
}

// GetChildSecondary() is a wrapper around gtk_button_box_get_child_secondary().
func (v *ButtonBox) GetChildSecondary(child IWidget) bool {
	c := C.gtk_button_box_get_child_secondary(v.native(), child.toWidget())
	return gobool(c)
}

// GetChildNonHomogeneous() is a wrapper around gtk_button_box_get_child_non_homogeneous().
func (v *ButtonBox) GetChildNonHomogeneous(child IWidget) bool {
	c := C.gtk_button_box_get_child_non_homogeneous(v.native(), child.toWidget())
	return gobool(c)
}

// SetLayout() is a wrapper around gtk_button_box_set_layout().
func (v *ButtonBox) SetLayout(style ButtonBoxStyle) {
	C.gtk_button_box_set_layout(v.native(), C.GtkButtonBoxStyle(style))
}

// SetChildSecondary() is a wrapper around gtk_button_box_set_child_secondary().
func (v *ButtonBox) SetChildSecondary(child IWidget, isSecondary bool) {
	C.gtk_button_box_set_child_secondary(v.native(), child.toWidget(), gbool(isSecondary))
}

// SetChildNonHomogeneous() is a wrapper around gtk_button_box_set_child_non_homogeneous().
func (v *ButtonBox) SetChildNonHomogeneous(child IWidget, nonHomogeneous bool) {
	C.gtk_button_box_set_child_non_homogeneous(v.native(), child.toWidget(), gbool(nonHomogeneous))
}

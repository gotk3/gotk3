// Same copyright and license as the rest of the files in this project

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {

	WrapMap["GtkTextMark"] = wrapTextMark
}

/*
 * GtkTextMark
 */

// TextMark is a representation of GTK's GtkTextMark.
// A position in the buffer preserved across buffer modifications
type TextMark struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkTextMark.
func (v *TextMark) native() *C.GtkTextMark {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTextMark(p)
}

func marshalTextMark(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTextMark(obj), nil
}

func wrapTextMark(obj *glib.Object) *TextMark {
	return &TextMark{obj}
}

// TextMarkNew is a wrapper around gtk_text_mark_new().
func TextMarkNew(name string, leftGravity bool) (*TextMark, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_text_mark_new((*C.gchar)(cstr), gbool(leftGravity))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapTextMark(glib.Take(unsafe.Pointer(c))), nil
}

// SetVisible is a wrapper around gtk_text_mark_set_visible().
func (v *TextMark) SetVisible(setting bool) {
	C.gtk_text_mark_set_visible(v.native(), gbool(setting))
}

// GetVisible is a wrapper around gtk_text_mark_get_visible().
func (v *TextMark) GetVisible() bool {
	return gobool(C.gtk_text_mark_get_visible(v.native()))
}

// GetDeleted is a wrapper around gtk_text_mark_get_deleted().
func (v *TextMark) GetDeleted() bool {
	return gobool(C.gtk_text_mark_get_deleted(v.native()))
}

// GetName is a wrapper around gtk_text_mark_get_name().
func (v *TextMark) GetName() string {
	return goString(C.gtk_text_mark_get_name(v.native()))
}

// GetBuffer is a wrapper around gtk_text_mark_get_buffer().
func (v *TextMark) GetBuffer() (*TextBuffer, error) {
	c := C.gtk_text_mark_get_buffer(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapTextBuffer(glib.Take(unsafe.Pointer(c))), nil
}

// GetLeftGravity is a wrapper around gtk_text_mark_get_left_gravity().
func (v *TextMark) GetLeftGravity() bool {
	return gobool(C.gtk_text_mark_get_left_gravity(v.native()))
}

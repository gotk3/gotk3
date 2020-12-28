// Same copyright and license as the rest of the files in this project

package gtk

// #include <gtk/gtk.h>
// #include "text_child_anchor.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Objects/Interfaces
		{glib.Type(C.gtk_text_child_anchor_get_type()), marshalTextChildAnchor},
	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * GtkTextChildAnchor
 */

// TextChildAnchor is a representation of GTK's GtkTextChildAnchor
type TextChildAnchor struct {
	glib.InitiallyUnowned
}

// native returns a pointer to the underlying GtkTextChildAnchor.
func (v *TextChildAnchor) native() *C.GtkTextChildAnchor {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTextChildAnchor(p)
}

func marshalTextChildAnchor(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapTextChildAnchor(obj), nil
}

func wrapTextChildAnchor(obj *glib.Object) *TextChildAnchor {
	return &TextChildAnchor{glib.InitiallyUnowned{obj}}
}

// TextChildAnchorNew is a wrapper around gtk_text_child_anchor_new ()
func TextChildAnchorNew() (*TextChildAnchor, error) {
	c := C.gtk_text_child_anchor_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapTextChildAnchor(glib.Take(unsafe.Pointer(c))), nil
}

// GetWidgets is a wrapper around gtk_text_child_anchor_get_widgets ().
func (v *TextChildAnchor) GetWidgets() *glib.List {
	clist := C.gtk_text_child_anchor_get_widgets(v.native())
	if clist == nil {
		return nil
	}

	glist := glib.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		return wrapWidget(glib.Take(ptr))
	})

	return glist
}

// GetDeleted is a wrapper around gtk_text_child_anchor_get_deleted().
func (v *TextChildAnchor) GetDeleted() bool {
	return gobool(C.gtk_text_child_anchor_get_deleted(v.native()))
}

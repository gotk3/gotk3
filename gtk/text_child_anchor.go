// Same copyright and license as the rest of the files in this project

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

/*
 * GtkTextChildAnchor
 */

// TextChildAnchor is a representation of GTK's GtkTextChildAnchor
type TextChildAnchor C.GtkTextChildAnchor

// native returns a pointer to the underlying GtkTextChildAnchor.
func (v *TextChildAnchor) native() *C.GtkTextChildAnchor {
	return (*C.GtkTextChildAnchor)(v)
}

// TextChildAnchorNew is a wrapper around gtk_text_child_anchor_new ()
func TextChildAnchorNew() *TextChildAnchor {
	ret := C.gtk_text_child_anchor_new()
	return (*TextChildAnchor)(ret)
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

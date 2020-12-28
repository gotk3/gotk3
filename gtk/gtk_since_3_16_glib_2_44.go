// Same copyright and license as the rest of the files in this project
// The code in this file is only for GTK+ version 3.16+, as well as Glib version 2.44+

// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!glib_2_40,!glib_2_42

package gtk

// #include <gtk/gtk.h>
// #include "gtk_since_3_16.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/internal/callback"
	"github.com/gotk3/gotk3/glib"
)

// BindModel is a wrapper around gtk_list_box_bind_model().
func (v *ListBox) BindModel(listModel *glib.ListModel, createWidgetFunc ListBoxCreateWidgetFunc) {
	C._gtk_list_box_bind_model(
		v.native(),
		C.toGListModel(unsafe.Pointer(listModel.Native())),
		C.gpointer(callback.Assign(createWidgetFunc)),
	)
}

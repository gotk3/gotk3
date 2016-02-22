// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,gtk_3_18

// See: https://developer.gnome.org/gtk3/3.18/api-index-3-18.html

package gtkf

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	glib_impl "github.com/gotk3/gotk3/glibf"
	"github.com/gotk3/gotk3/gtk"
)

//void
//gtk_popover_set_default_widget (GtkPopover *popover, GtkWidget *widget);
func (p *Popover) SetDefaultWidget(widget gtk.Widget) {
	C.gtk_popover_set_default_widget(p.native(), castToWidget(widget))
}

//GtkWidget *
//gtk_popover_get_default_widget (GtkPopover *popover);
func (p *Popover) GetDefaultWidget() gtk.Widget {
	w := C.gtk_popover_get_default_widget(p.native())
	if w == nil {
		return nil
	}
	return &Widget{glib_impl.InitiallyUnowned{wrapObject(unsafe.Pointer(w))}}
}

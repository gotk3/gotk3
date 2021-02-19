// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"

	"github.com/gotk3/gotk3/glib"
)

/*
 * GtkMenu
 */

// Menu is a representation of GTK's GtkMenu.
type Menu struct {
	MenuShell
}

// IMenu is an interface type implemented by all structs embedding
// a Menu.  It is meant to be used as an argument type for wrapper
// functions that wrap around a C GTK function taking a
// GtkMenu.
type IMenu interface {
	toMenu() *C.GtkMenu
	toWidget() *C.GtkWidget
}

// native() returns a pointer to the underlying GtkMenu.
func (v *Menu) native() *C.GtkMenu {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenu(p)
}

func (v *Menu) toMenu() *C.GtkMenu {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalMenu(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenu(obj), nil
}

func wrapMenu(obj *glib.Object) *Menu {
	if obj == nil {
		return nil
	}

	return &Menu{MenuShell{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// MenuNew() is a wrapper around gtk_menu_new().
func MenuNew() (*Menu, error) {
	c := C.gtk_menu_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapMenu(glib.Take(unsafe.Pointer(c))), nil
}

// GtkMenuNewFromModel is a wrapper around gtk_menu_new_from_model().
func GtkMenuNewFromModel(model *glib.MenuModel) (*Menu, error) {
	c := C.gtk_menu_new_from_model(C.toGMenuModel(unsafe.Pointer(model.Native())))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapMenu(glib.Take(unsafe.Pointer(c))), nil
}

// SetScreen is a wrapper around gtk_menu_set_screen().
func (v *Menu) SetScreen(screen *gdk.Screen) {
	C.gtk_menu_set_screen(v.native(), (*C.GdkScreen)(unsafe.Pointer(screen.Native())))
}

// Attach is a wrapper around gtk_menu_attach().
func (v *Menu) Attach(child IWidget, l, r, t, b uint) {
	C.gtk_menu_attach(
		v.native(),
		child.toWidget(),
		C.guint(l),
		C.guint(r),
		C.guint(t),
		C.guint(b))
}

// SetMonitor() is a wrapper around gtk_menu_set_monitor().
func (v *Menu) SetMonitor(monitor_num int) {
	C.gtk_menu_set_monitor(v.native(), C.gint(monitor_num))
}

// GetMonitor() is a wrapper around gtk_menu_get_monitor().
func (v *Menu) GetMonitor() int {
	return int(C.gtk_menu_get_monitor(v.native()))
}

// ReorderChild() is a wrapper around gtk_menu_reorder_child().
func (v *Menu) ReorderChild(child IWidget, position int) {
	C.gtk_menu_reorder_child(v.native(), child.toWidget(), C.gint(position))
}

// SetReserveToggleSize() is a wrapper around gtk_menu_set_reserve_toggle_size().
func (v *Menu) SetReserveToggleSize(reserve bool) {
	C.gtk_menu_set_reserve_toggle_size(v.native(), gbool(reserve))
}

// GetReserveToggleSize() is a wrapper around gtk_menu_get_reserve_toggle_size().
func (v *Menu) GetReserveToggleSize() bool {
	return gobool(C.gtk_menu_get_reserve_toggle_size(v.native()))
}

// Popdown() is a wrapper around gtk_menu_popdown().
func (v *Menu) Popdown() {
	C.gtk_menu_popdown(v.native())
}

// TODO
/*
gtk_menu_reposition () require 'GtkMenuPositionFunc' (according to its position function.)
*/

// GetActive() is a wrapper around gtk_menu_get_active().
func (v *Menu) GetActive() (*Menu, error) {
	c := C.gtk_menu_get_active(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapMenu(glib.Take(unsafe.Pointer(c))), nil
}

// SetActive() is a wrapper around gtk_menu_set_active().
func (v *Menu) SetActive(index uint) {
	C.gtk_menu_set_active(v.native(), C.guint(index))
}

// TODO
/*
void
gtk_menu_attach_to_widget (GtkMenu *menu,
                           GtkWidget *attach_widget,
                           GtkMenuDetachFunc detacher);

void
gtk_menu_detach (GtkMenu *menu);
*/

// GetAttachWidget() is a wrapper around gtk_menu_get_attach_widget().
func (v *Menu) GetAttachWidget() (IWidget, error) {
	c := C.gtk_menu_get_attach_widget(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return castWidget(c)
}

// TODO
/*
GList *
gtk_menu_get_for_attach_widget (GtkWidget *widget);
*/

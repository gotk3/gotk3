// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

package impl

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	glib_impl "github.com/gotk3/gotk3/glib/impl"
	"github.com/gotk3/gotk3/gtk"
)

/*
 * GtkMenuShell
 */

// MenuShell is a representation of GTK's GtkMenuShell.
type menuShell struct {
	container
}

// native returns a pointer to the underlying GtkMenuShell.
func (v *menuShell) native() *C.GtkMenuShell {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenuShell(p)
}

func marshalMenuShell(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapMenuShell(obj), nil
}

func wrapMenuShell(obj *glib_impl.Object) *menuShell {
	return &menuShell{container{widget{glib_impl.InitiallyUnowned{obj}}}}
}

// Append is a wrapper around gtk_menu_shell_append().
func (v *menuShell) Append(child gtk.MenuItem) {
	C.gtk_menu_shell_append(v.native(), child.(IMenuItem).toWidget())
}

// Prepend is a wrapper around gtk_menu_shell_prepend().
func (v *menuShell) Prepend(child gtk.MenuItem) {
	C.gtk_menu_shell_prepend(v.native(), child.(IMenuItem).toWidget())
}

// Insert is a wrapper around gtk_menu_shell_insert().
func (v *menuShell) Insert(child gtk.MenuItem, position int) {
	C.gtk_menu_shell_insert(v.native(), child.(IMenuItem).toWidget(), C.gint(position))
}

// Deactivate is a wrapper around gtk_menu_shell_deactivate().
func (v *menuShell) Deactivate() {
	C.gtk_menu_shell_deactivate(v.native())
}

// SelectItem is a wrapper around gtk_menu_shell_select_item().
func (v *menuShell) SelectItem(child gtk.MenuItem) {
	C.gtk_menu_shell_select_item(v.native(), child.(IMenuItem).toWidget())
}

// SelectFirst is a wrapper around gtk_menu_shell_select_first().
func (v *menuShell) SelectFirst(searchSensitive bool) {
	C.gtk_menu_shell_select_first(v.native(), gbool(searchSensitive))
}

// Deselect is a wrapper around gtk_menu_shell_deselect().
func (v *menuShell) Deselect() {
	C.gtk_menu_shell_deselect(v.native())
}

// ActivateItem is a wrapper around gtk_menu_shell_activate_item().
func (v *menuShell) ActivateItem(child gtk.MenuItem, forceDeactivate bool) {
	C.gtk_menu_shell_activate_item(v.native(), child.(IMenuItem).toWidget(), gbool(forceDeactivate))
}

// Cancel is a wrapper around gtk_menu_shell_cancel().
func (v *menuShell) Cancel() {
	C.gtk_menu_shell_cancel(v.native())
}

// SetTakeFocus is a wrapper around gtk_menu_shell_set_take_focus().
func (v *menuShell) SetTakeFocus(takeFocus bool) {
	C.gtk_menu_shell_set_take_focus(v.native(), gbool(takeFocus))
}

// gboolean 	gtk_menu_shell_get_take_focus ()
// GtkWidget * 	gtk_menu_shell_get_selected_item ()
// GtkWidget * 	gtk_menu_shell_get_parent_shell ()
// void 	gtk_menu_shell_bind_model ()

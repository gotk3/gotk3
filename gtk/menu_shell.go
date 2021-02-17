// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_menu_direction_type_get_type()), marshalMenuDirectionType},
	}

	glib.RegisterGValueMarshalers(tm)
}

// MenuDirectionType is a representation of GTK's GtkMenuDirectionType.
type MenuDirectionType int

const (
	MENU_DIR_PARENT MenuDirectionType = C.GTK_MENU_DIR_PARENT
	MENU_DIR_CHILD  MenuDirectionType = C.GTK_MENU_DIR_CHILD
	MENU_DIR_NEXT   MenuDirectionType = C.GTK_MENU_DIR_NEXT
	MENU_DIR_PREV   MenuDirectionType = C.GTK_MENU_DIR_PREV
)

func marshalMenuDirectionType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return MenuDirectionType(c), nil
}

/*
 * GtkMenuShell
 */

// MenuShell is a representation of GTK's GtkMenuShell.
type MenuShell struct {
	Container
}

// native returns a pointer to the underlying GtkMenuShell.
func (v *MenuShell) native() *C.GtkMenuShell {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenuShell(p)
}

func marshalMenuShell(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenuShell(obj), nil
}

func wrapMenuShell(obj *glib.Object) *MenuShell {
	if obj == nil {
		return nil
	}

	return &MenuShell{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// Append is a wrapper around gtk_menu_shell_append().
func (v *MenuShell) Append(child IMenuItem) {
	C.gtk_menu_shell_append(v.native(), child.toWidget())
}

// Prepend is a wrapper around gtk_menu_shell_prepend().
func (v *MenuShell) Prepend(child IMenuItem) {
	C.gtk_menu_shell_prepend(v.native(), child.toWidget())
}

// Insert is a wrapper around gtk_menu_shell_insert().
func (v *MenuShell) Insert(child IMenuItem, position int) {
	C.gtk_menu_shell_insert(v.native(), child.toWidget(), C.gint(position))
}

// Deactivate is a wrapper around gtk_menu_shell_deactivate().
func (v *MenuShell) Deactivate() {
	C.gtk_menu_shell_deactivate(v.native())
}

// SelectItem is a wrapper around gtk_menu_shell_select_item().
func (v *MenuShell) SelectItem(child IMenuItem) {
	C.gtk_menu_shell_select_item(v.native(), child.toWidget())
}

// SelectFirst is a wrapper around gtk_menu_shell_select_first().
func (v *MenuShell) SelectFirst(searchSensitive bool) {
	C.gtk_menu_shell_select_first(v.native(), gbool(searchSensitive))
}

// Deselect is a wrapper around gtk_menu_shell_deselect().
func (v *MenuShell) Deselect() {
	C.gtk_menu_shell_deselect(v.native())
}

// ActivateItem is a wrapper around gtk_menu_shell_activate_item().
func (v *MenuShell) ActivateItem(child IMenuItem, forceDeactivate bool) {
	C.gtk_menu_shell_activate_item(v.native(), child.toWidget(), gbool(forceDeactivate))
}

// Cancel is a wrapper around gtk_menu_shell_cancel().
func (v *MenuShell) Cancel() {
	C.gtk_menu_shell_cancel(v.native())
}

// SetTakeFocus is a wrapper around gtk_menu_shell_set_take_focus().
func (v *MenuShell) SetTakeFocus(takeFocus bool) {
	C.gtk_menu_shell_set_take_focus(v.native(), gbool(takeFocus))
}

// GetTakeFocus is a wrapper around gtk_menu_shell_get_take_focus().
func (v *MenuShell) GetTakeFocus() bool {
	return gobool(C.gtk_menu_shell_get_take_focus(v.native()))
}

// GetSelectedItem is a wrapper around gtk_menu_shell_get_selected_item().
func (v *MenuShell) GetSelectedItem() (IMenuItem, error) {
	c := C.gtk_menu_shell_get_selected_item(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenuItem(obj), nil
}

// GetParentShell is a wrapper around gtk_menu_shell_get_parent_shell().
func (v *MenuShell) GetParentShell() (*MenuShell, error) {
	c := C.gtk_menu_shell_get_parent_shell(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenuShell(obj), nil
}

// BindModel is a wrapper around gtk_menu_shell_bind_model().
func (v *MenuShell) BindModel(model *glib.MenuModel,
	action_namespace string, with_separators bool) {

	cstr := C.CString(action_namespace)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_menu_shell_bind_model(
		v.native(),
		(*C.GMenuModel)(unsafe.Pointer(model.Native())),
		cstr,
		gbool(with_separators))
}

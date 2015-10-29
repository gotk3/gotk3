// Same copyright and license as the rest of the files in this project
// This file contains style related functions and structures
package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

/*
 * GtkApplication
 */

// Application is a representation of GTK's GtkApplication.
type Application struct {
	app *C.GtkApplication
}

// native returns a pointer to the underlying GtkApplication.
func (v *Application) native() *C.GtkApplication {
	if v == nil || v.app == nil {
		return nil
	}
	p := unsafe.Pointer(v.app)
	return C.toGtkApplication(p)
}

func wrapApplication(obj *C.GtkApplication) *Application {
	return &Application{obj}
}

// ApplicationNew is a wrapper around gtk_application_new().
func ApplicationNew(appId string, flags glib.ApplicationFlags) (*Application, error) {
	cstr := (*C.gchar)(C.CString(appId))
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_application_new(cstr, C.GApplicationFlags(flags))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapApplication(c), nil
}

// AddWindow is a wrapper around gtk_application_add_window().
func (v *Application) AddWindow(w *Window) {
	C.gtk_application_add_window(v.native(), w.native())
}

// RemoveWindow is a wrapper around gtk_application_remove_window().
func (v *Application) RemoveWindow(w *Window) {
	C.gtk_application_remove_window(v.native(), w.native())
}

// GetWindowByID is a wrapper around gtk_application_get_window_by_id().
func (v *Application) GetWindowByID(id uint) *Window {
	c := C.gtk_application_get_window_by_id(v.native(), C.guint(id))
	if c == nil {
		return nil
	}
	return wrapWindow(wrapObject(unsafe.Pointer(c)))
}

// GetActiveWindow is a wrapper around gtk_application_get_active_window().
func (v *Application) GetActiveWindow() *Window {
	c := C.gtk_application_get_active_window(v.native())
	if c == nil {
		return nil
	}
	return wrapWindow(wrapObject(unsafe.Pointer(c)))
}

// Uninhibit is a wrapper around gtk_application_uninhibit().
func (v *Application) Uninhibit(cookie uint) {
	C.gtk_application_uninhibit(v.native(), C.guint(cookie))
}

// GMenuModel * 	gtk_application_get_app_menu ()
// void 	gtk_application_set_app_menu ()
// GMenuModel * 	gtk_application_get_menubar ()
// void 	gtk_application_set_menubar ()
// GMenu * 	gtk_application_get_menu_by_id ()
// void 	gtk_application_add_accelerator ()
// void 	gtk_application_remove_accelerator ()
// gboolean 	gtk_application_is_inhibited ()
// guint 	gtk_application_inhibit ()
// GList * 	gtk_application_get_windows ()
// gchar ** 	gtk_application_list_action_descriptions ()
// gchar ** 	gtk_application_get_accels_for_action ()
// void 	gtk_application_set_accels_for_action ()
// gchar ** 	gtk_application_get_actions_for_accel ()

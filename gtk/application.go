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

// ApplicationInhibitFlags is a representation of GTK's GtkApplicationInhibitFlags.
type ApplicationInhibitFlags int

const (
	APPLICATION_INHIBIT_LOGOUT  ApplicationInhibitFlags = C.GTK_APPLICATION_INHIBIT_LOGOUT
	APPLICATION_INHIBIT_SWITCH  ApplicationInhibitFlags = C.GTK_APPLICATION_INHIBIT_SWITCH
	APPLICATION_INHIBIT_SUSPEND ApplicationInhibitFlags = C.GTK_APPLICATION_INHIBIT_SUSPEND
	APPLICATION_INHIBIT_IDLE    ApplicationInhibitFlags = C.GTK_APPLICATION_INHIBIT_IDLE
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

// GetAppMenu is a wrapper around gtk_application_get_app_menu().
func (v *Application) GetAppMenu() *glib.MenuModel {
	c := C.gtk_application_get_app_menu(v.native())
	if c == nil {
		return nil
	}
	return &glib.MenuModel{wrapObject(unsafe.Pointer(c))}
}

// SetAppMenu is a wrapper around gtk_application_set_app_menu().
func (v *Application) SetAppMenu(m *glib.MenuModel) {
	mptr := (*C.GMenuModel)(unsafe.Pointer(m.Native()))
	C.gtk_application_set_app_menu(v.native(), mptr)
}

// GetMenubar is a wrapper around gtk_application_get_menubar().
func (v *Application) GetMenubar() *glib.MenuModel {
	c := C.gtk_application_get_menubar(v.native())
	if c == nil {
		return nil
	}
	return &glib.MenuModel{wrapObject(unsafe.Pointer(c))}
}

// SetMenubar is a wrapper around gtk_application_set_menubar().
func (v *Application) SetMenubar(m *glib.MenuModel) {
	mptr := (*C.GMenuModel)(unsafe.Pointer(m.Native()))
	C.gtk_application_set_menubar(v.native(), mptr)
}

// GetMenuByID is a wrapper around gtk_application_get_menu_by_id().
func (v *Application) GetMenuByID(id string) *glib.Menu {
	cstr1 := (*C.gchar)(C.CString(id))
	defer C.free(unsafe.Pointer(cstr1))

	c := C.gtk_application_get_menu_by_id(v.native(), cstr1)
	if c == nil {
		return nil
	}
	return &glib.Menu{glib.MenuModel{wrapObject(unsafe.Pointer(c))}}
}

// ListActionDescriptions is a wrapper around gtk_application_list_action_descriptions().
func (v *Application) ListActionDescriptions() []string {
	var descs []string
	c := C.gtk_application_list_action_descriptions(v.native())
	originalc := c
	defer C.g_strfreev(originalc)

	for *c != nil {
		descs = append(descs, C.GoString((*C.char)(*c)))
		c = C.next_gcharptr(c)
	}

	return descs
}

// GetAccelsForAction is a wrapper around gtk_application_get_accels_for_action().
func (v *Application) GetAccelsForAction(act string) []string {
	cstr1 := (*C.gchar)(C.CString(act))
	defer C.free(unsafe.Pointer(cstr1))

	var descs []string
	c := C.gtk_application_get_accels_for_action(v.native(), cstr1)
	originalc := c
	defer C.g_strfreev(originalc)

	for *c != nil {
		descs = append(descs, C.GoString((*C.char)(*c)))
		c = C.next_gcharptr(c)
	}

	return descs
}

// SetAccelsForAction is a wrapper around gtk_application_set_accels_for_action().
func (v *Application) SetAccelsForAction(act string, accels []string) {
	cstr1 := (*C.gchar)(C.CString(act))
	defer C.free(unsafe.Pointer(cstr1))

	caccels := C.make_strings(C.int(len(accels) + 1))
	defer C.destroy_strings(caccels)

	for i, accel := range accels {
		cstr := C.CString(accel)
		defer C.free(unsafe.Pointer(cstr))
		C.set_string(caccels, C.int(i), (*C.gchar)(cstr))
	}

	C.set_string(caccels, C.int(len(accels)), nil)

	C.gtk_application_set_accels_for_action(v.native(), cstr1, caccels)
}

// IsInhibited is a wrapper around gtk_application_is_inhibited().
func (v *Application) IsInhibited(flags ApplicationInhibitFlags) bool {
	return gobool(C.gtk_application_is_inhibited(v.native(), C.GtkApplicationInhibitFlags(flags)))
}

// Inhibit is a wrapper around gtk_application_inhibit().
func (v *Application) Inhibited(w *Window, flags ApplicationInhibitFlags, reason string) uint {
	cstr1 := (*C.gchar)(C.CString(reason))
	defer C.free(unsafe.Pointer(cstr1))

	return uint(C.gtk_application_inhibit(v.native(), w.native(), C.GtkApplicationInhibitFlags(flags), cstr1))
}

// void 	gtk_application_add_accelerator () // deprecated and uses a gvariant paramater
// void 	gtk_application_remove_accelerator () // deprecated and uses a gvariant paramater
// GList * 	gtk_application_get_windows () // glist

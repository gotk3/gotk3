// Same copyright and license as the rest of the files in this project
// This file contains style related functions and structures

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	glib_iface "github.com/gotk3/gotk3/glib/iface"
	"github.com/gotk3/gotk3/gtk/iface"
)

func init() {
	iface.APPLICATION_INHIBIT_LOGOUT = C.GTK_APPLICATION_INHIBIT_LOGOUT
	iface.APPLICATION_INHIBIT_SWITCH = C.GTK_APPLICATION_INHIBIT_SWITCH
	iface.APPLICATION_INHIBIT_SUSPEND = C.GTK_APPLICATION_INHIBIT_SUSPEND
	iface.APPLICATION_INHIBIT_IDLE = C.GTK_APPLICATION_INHIBIT_IDLE
}

/*
 * GtkApplication
 */

// Application is a representation of GTK's GtkApplication.
type Application struct {
	glib.Application
}

// native returns a pointer to the underlying GtkApplication.
func (v *Application) native() *C.GtkApplication {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkApplication(unsafe.Pointer(v.GObject))
}

func wrapApplication(obj *glib.Object) *Application {
	return &Application{glib.Application{obj}}
}

// ApplicationNew is a wrapper around gtk_application_new().
func ApplicationNew(appId string, flags glib_iface.ApplicationFlags) (*Application, error) {
	cstr := (*C.gchar)(C.CString(appId))
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_application_new(cstr, C.GApplicationFlags(flags))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapApplication(wrapObject(unsafe.Pointer(c))), nil
}

// AddWindow is a wrapper around gtk_application_add_window().
func (v *Application) AddWindow(w iface.Window) {
	C.gtk_application_add_window(v.native(), w.(*Window).native())
}

// RemoveWindow is a wrapper around gtk_application_remove_window().
func (v *Application) RemoveWindow(w iface.Window) {
	C.gtk_application_remove_window(v.native(), w.(*Window).native())
}

// GetWindowByID is a wrapper around gtk_application_get_window_by_id().
func (v *Application) GetWindowByID(id uint) iface.Window {
	c := C.gtk_application_get_window_by_id(v.native(), C.guint(id))
	if c == nil {
		return nil
	}
	return wrapWindow(wrapObject(unsafe.Pointer(c)))
}

// GetActiveWindow is a wrapper around gtk_application_get_active_window().
func (v *Application) GetActiveWindow() iface.Window {
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
func (v *Application) GetAppMenu() glib_iface.MenuModel {
	c := C.gtk_application_get_app_menu(v.native())
	if c == nil {
		return nil
	}
	return &glib.MenuModel{wrapObject(unsafe.Pointer(c))}
}

// SetAppMenu is a wrapper around gtk_application_set_app_menu().
func (v *Application) SetAppMenu(m glib_iface.MenuModel) {
	mptr := (*C.GMenuModel)(unsafe.Pointer(m.(*glib.MenuModel).Native()))
	C.gtk_application_set_app_menu(v.native(), mptr)
}

// GetMenubar is a wrapper around gtk_application_get_menubar().
func (v *Application) GetMenubar() glib_iface.MenuModel {
	c := C.gtk_application_get_menubar(v.native())
	if c == nil {
		return nil
	}
	return &glib.MenuModel{wrapObject(unsafe.Pointer(c))}
}

// SetMenubar is a wrapper around gtk_application_set_menubar().
func (v *Application) SetMenubar(m glib_iface.MenuModel) {
	mptr := (*C.GMenuModel)(unsafe.Pointer(m.(*glib.MenuModel).Native()))
	C.gtk_application_set_menubar(v.native(), mptr)
}

// IsInhibited is a wrapper around gtk_application_is_inhibited().
func (v *Application) IsInhibited(flags iface.ApplicationInhibitFlags) bool {
	return gobool(C.gtk_application_is_inhibited(v.native(), C.GtkApplicationInhibitFlags(flags)))
}

// Inhibited is a wrapper around gtk_application_inhibit().
func (v *Application) Inhibited(w iface.Window, flags iface.ApplicationInhibitFlags, reason string) uint {
	cstr1 := (*C.gchar)(C.CString(reason))
	defer C.free(unsafe.Pointer(cstr1))

	return uint(C.gtk_application_inhibit(v.native(), w.(*Window).native(), C.GtkApplicationInhibitFlags(flags), cstr1))
}

// void 	gtk_application_add_accelerator () // deprecated and uses a gvariant paramater
// void 	gtk_application_remove_accelerator () // deprecated and uses a gvariant paramater

// GetWindows is a wrapper around gtk_application_get_windows().
// Returned list is wrapped to return *gtk.Window elements.
func (v *Application) GetWindows() glib_iface.List {
	glist := C.gtk_application_get_windows(v.native())
	list := glib.WrapList(uintptr(unsafe.Pointer(glist)))
	list.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		return wrapWindow(wrapObject(ptr))
	})
	runtime.SetFinalizer(list, func(l *glib.List) {
		l.Free()
	})
	return list
}

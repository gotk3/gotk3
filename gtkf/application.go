// Same copyright and license as the rest of the files in this project
// This file contains style related functions and structures

package gtkf

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	glib_impl "github.com/gotk3/gotk3/glibf"
	"github.com/gotk3/gotk3/gtk"
)

func init() {
	gtk.APPLICATION_INHIBIT_LOGOUT = C.GTK_APPLICATION_INHIBIT_LOGOUT
	gtk.APPLICATION_INHIBIT_SWITCH = C.GTK_APPLICATION_INHIBIT_SWITCH
	gtk.APPLICATION_INHIBIT_SUSPEND = C.GTK_APPLICATION_INHIBIT_SUSPEND
	gtk.APPLICATION_INHIBIT_IDLE = C.GTK_APPLICATION_INHIBIT_IDLE
}

/*
 * GtkApplication
 */

// Application is a representation of GTK's GtkApplication.
type application struct {
	glib_impl.Application
}

// native returns a pointer to the underlying GtkApplication.
func (v *application) native() *C.GtkApplication {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkApplication(unsafe.Pointer(v.GObject))
}

func wrapApplication(obj *glib_impl.Object) *application {
	return &application{glib_impl.Application{obj}}
}

// ApplicationNew is a wrapper around gtk_application_new().
func ApplicationNew(appId string, flags glib.ApplicationFlags) (*application, error) {
	cstr := (*C.gchar)(C.CString(appId))
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_application_new(cstr, C.GApplicationFlags(flags))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapApplication(wrapObject(unsafe.Pointer(c))), nil
}

// AddWindow is a wrapper around gtk_application_add_window().
func (v *application) AddWindow(w gtk.Window) {
	C.gtk_application_add_window(v.native(), asWindowImpl(w).native())
}

// RemoveWindow is a wrapper around gtk_application_remove_window().
func (v *application) RemoveWindow(w gtk.Window) {
	C.gtk_application_remove_window(v.native(), asWindowImpl(w).native())
}

// GetWindowByID is a wrapper around gtk_application_get_window_by_id().
func (v *application) GetWindowByID(id uint) gtk.Window {
	c := C.gtk_application_get_window_by_id(v.native(), C.guint(id))
	if c == nil {
		return nil
	}
	return wrapWindow(wrapObject(unsafe.Pointer(c)))
}

// GetActiveWindow is a wrapper around gtk_application_get_active_window().
func (v *application) GetActiveWindow() gtk.Window {
	c := C.gtk_application_get_active_window(v.native())
	if c == nil {
		return nil
	}
	return wrapWindow(wrapObject(unsafe.Pointer(c)))
}

// Uninhibit is a wrapper around gtk_application_uninhibit().
func (v *application) Uninhibit(cookie uint) {
	C.gtk_application_uninhibit(v.native(), C.guint(cookie))
}

// GetAppMenu is a wrapper around gtk_application_get_app_menu().
func (v *application) GetAppMenu() glib.MenuModel {
	c := C.gtk_application_get_app_menu(v.native())
	if c == nil {
		return nil
	}
	return &glib_impl.MenuModel{wrapObject(unsafe.Pointer(c))}
}

// SetAppMenu is a wrapper around gtk_application_set_app_menu().
func (v *application) SetAppMenu(m glib.MenuModel) {
	mptr := (*C.GMenuModel)(unsafe.Pointer(glib_impl.CastToMenuModel(m).Native()))
	C.gtk_application_set_app_menu(v.native(), mptr)
}

// GetMenubar is a wrapper around gtk_application_get_menubar().
func (v *application) GetMenubar() glib.MenuModel {
	c := C.gtk_application_get_menubar(v.native())
	if c == nil {
		return nil
	}
	return &glib_impl.MenuModel{wrapObject(unsafe.Pointer(c))}
}

// SetMenubar is a wrapper around gtk_application_set_menubar().
func (v *application) SetMenubar(m glib.MenuModel) {
	mptr := (*C.GMenuModel)(unsafe.Pointer(glib_impl.CastToMenuModel(m).Native()))
	C.gtk_application_set_menubar(v.native(), mptr)
}

// IsInhibited is a wrapper around gtk_application_is_inhibited().
func (v *application) IsInhibited(flags gtk.ApplicationInhibitFlags) bool {
	return gobool(C.gtk_application_is_inhibited(v.native(), C.GtkApplicationInhibitFlags(flags)))
}

// Inhibited is a wrapper around gtk_application_inhibit().
func (v *application) Inhibited(w gtk.Window, flags gtk.ApplicationInhibitFlags, reason string) uint {
	cstr1 := (*C.gchar)(C.CString(reason))
	defer C.free(unsafe.Pointer(cstr1))

	return uint(C.gtk_application_inhibit(v.native(), asWindowImpl(w).native(), C.GtkApplicationInhibitFlags(flags), cstr1))
}

// void 	gtk_application_add_accelerator () // deprecated and uses a gvariant paramater
// void 	gtk_application_remove_accelerator () // deprecated and uses a gvariant paramater

// GetWindows is a wrapper around gtk_application_get_windows().
// Returned list is wrapped to return *gtk.Window elements.
func (v *application) GetWindows() glib.List {
	glist := C.gtk_application_get_windows(v.native())
	list := glib_impl.WrapList(uintptr(unsafe.Pointer(glist)))
	list.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		return wrapWindow(wrapObject(ptr))
	})
	runtime.SetFinalizer(list, func(l *glib_impl.List) {
		l.Free()
	})
	return list
}

// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

package gtkf

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"errors"
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	gdk_impl "github.com/gotk3/gotk3/gdkf"
	glib_impl "github.com/gotk3/gotk3/glibf"
	"github.com/gotk3/gotk3/gtk"
)

/*
 * GtkWindow
 */

// Window is a representation of GTK's GtkWindow.
type window struct {
	bin
}

// IWindow is an interface type implemented by all structs embedding a
// Window.  It is meant to be used as an argument type for wrapper
// functions that wrap around a C GTK function taking a GtkWindow.
type IWindow interface {
	asWindowImpl() *window
	toWindow() *C.GtkWindow
}

// native returns a pointer to the underlying GtkWindow.
func (v *window) native() *C.GtkWindow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkWindow(p)
}

func asWindowImpl(v gtk.Window) *window {
	if v == nil {
		return nil
	}
	return v.(IWindow).asWindowImpl()
}

func (v *window) asWindowImpl() *window {
	if v == nil {
		return nil
	}
	return v
}

func (v *window) toWindow() *C.GtkWindow {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalWindow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapWindow(obj), nil
}

func wrapWindow(obj *glib_impl.Object) *window {
	return &window{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}
}

// WindowNew is a wrapper around gtk_window_new().
func WindowNew(t gtk.WindowType) (*window, error) {
	c := C.gtk_window_new(C.GtkWindowType(t))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWindow(wrapObject(unsafe.Pointer(c))), nil
}

// SetTitle is a wrapper around gtk_window_set_title().
func (v *window) SetTitle(title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_window_set_title(v.native(), (*C.gchar)(cstr))
}

// SetResizable is a wrapper around gtk_window_set_resizable().
func (v *window) SetResizable(resizable bool) {
	C.gtk_window_set_resizable(v.native(), gbool(resizable))
}

// GetResizable is a wrapper around gtk_window_get_resizable().
func (v *window) GetResizable() bool {
	c := C.gtk_window_get_resizable(v.native())
	return gobool(c)
}

// ActivateFocus is a wrapper around gtk_window_activate_focus().
func (v *window) ActivateFocus() bool {
	c := C.gtk_window_activate_focus(v.native())
	return gobool(c)
}

// ActivateDefault is a wrapper around gtk_window_activate_default().
func (v *window) ActivateDefault() bool {
	c := C.gtk_window_activate_default(v.native())
	return gobool(c)
}

// SetModal is a wrapper around gtk_window_set_modal().
func (v *window) SetModal(modal bool) {
	C.gtk_window_set_modal(v.native(), gbool(modal))
}

// SetDefaultSize is a wrapper around gtk_window_set_default_size().
func (v *window) SetDefaultSize(width, height int) {
	C.gtk_window_set_default_size(v.native(), C.gint(width), C.gint(height))
}

// SetDefaultGeometry is a wrapper around gtk_window_set_default_geometry().
func (v *window) SetDefaultGeometry(width, height int) {
	C.gtk_window_set_default_geometry(v.native(), C.gint(width),
		C.gint(height))
}

// GetScreen is a wrapper around gtk_window_get_screen().
func (v *window) GetScreen() (gdk.Screen, error) {
	c := C.gtk_window_get_screen(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	s := &gdk_impl.Screen{wrapObject(unsafe.Pointer(c))}
	return s, nil
}

// SetIcon is a wrapper around gtk_window_set_icon().
func (v *window) SetIcon(icon gdk.Pixbuf) {
	iconPtr := (*C.GdkPixbuf)(unsafe.Pointer(gdk_impl.CastToPixbuf(icon).Native()))
	C.gtk_window_set_icon(v.native(), iconPtr)
}

// WindowSetDefaultIcon is a wrapper around gtk_window_set_default_icon().
func WindowSetDefaultIcon(icon gdk.Pixbuf) {
	iconPtr := (*C.GdkPixbuf)(unsafe.Pointer(gdk_impl.CastToPixbuf(icon).Native()))
	C.gtk_window_set_default_icon(iconPtr)
}

// TODO(jrick) GdkGeometry GdkWindowHints.
/*
func (v *window) SetGeometryHints() {
}
*/

// TODO(jrick) GdkGravity.
/*
func (v *window) SetGravity() {
}
*/

// TODO(jrick) GdkGravity.
/*
func (v *window) GetGravity() {
}
*/

// SetPosition is a wrapper around gtk_window_set_position().
func (v *window) SetPosition(position gtk.WindowPosition) {
	C.gtk_window_set_position(v.native(), C.GtkWindowPosition(position))
}

// SetTransientFor is a wrapper around gtk_window_set_transient_for().
func (v *window) SetTransientFor(parent gtk.Window) {
	var pw *C.GtkWindow = nil
	if parent != nil {
		pw = parent.(IWindow).toWindow()
	}
	C.gtk_window_set_transient_for(v.native(), pw)
}

// SetDestroyWithParent is a wrapper around
// gtk_window_set_destroy_with_parent().
func (v *window) SetDestroyWithParent(setting bool) {
	C.gtk_window_set_destroy_with_parent(v.native(), gbool(setting))
}

// SetHideTitlebarWhenMaximized is a wrapper around
// gtk_window_set_hide_titlebar_when_maximized().
func (v *window) SetHideTitlebarWhenMaximized(setting bool) {
	C.gtk_window_set_hide_titlebar_when_maximized(v.native(),
		gbool(setting))
}

// IsActive is a wrapper around gtk_window_is_active().
func (v *window) IsActive() bool {
	c := C.gtk_window_is_active(v.native())
	return gobool(c)
}

// HasToplevelFocus is a wrapper around gtk_window_has_toplevel_focus().
func (v *window) HasToplevelFocus() bool {
	c := C.gtk_window_has_toplevel_focus(v.native())
	return gobool(c)
}

// GetFocus is a wrapper around gtk_window_get_focus().
func (v *window) GetFocus() (gtk.Widget, error) {
	c := C.gtk_window_get_focus(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWidget(wrapObject(unsafe.Pointer(c))), nil
}

// SetFocus is a wrapper around gtk_window_set_focus().
func (v *window) SetFocus(w gtk.Widget) {
	C.gtk_window_set_focus(v.native(), asWidgetImpl(w).native())
}

// GetDefaultWidget is a wrapper arround gtk_window_get_default_widget().
func (v *window) GetDefaultWidget() gtk.Widget {
	c := C.gtk_window_get_default_widget(v.native())
	if c == nil {
		return nil
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapWidget(obj)
}

// SetDefault is a wrapper arround gtk_window_set_default().
func (v *window) SetDefault(widget gtk.Widget) {
	C.gtk_window_set_default(v.native(), castToWidget(widget))
}

// Present is a wrapper around gtk_window_present().
func (v *window) Present() {
	C.gtk_window_present(v.native())
}

// PresentWithTime is a wrapper around gtk_window_present_with_time().
func (v *window) PresentWithTime(ts uint32) {
	C.gtk_window_present_with_time(v.native(), C.guint32(ts))
}

// Iconify is a wrapper around gtk_window_iconify().
func (v *window) Iconify() {
	C.gtk_window_iconify(v.native())
}

// Deiconify is a wrapper around gtk_window_deiconify().
func (v *window) Deiconify() {
	C.gtk_window_deiconify(v.native())
}

// Stick is a wrapper around gtk_window_stick().
func (v *window) Stick() {
	C.gtk_window_stick(v.native())
}

// Unstick is a wrapper around gtk_window_unstick().
func (v *window) Unstick() {
	C.gtk_window_unstick(v.native())
}

// Maximize is a wrapper around gtk_window_maximize().
func (v *window) Maximize() {
	C.gtk_window_maximize(v.native())
}

// Unmaximize is a wrapper around gtk_window_unmaximize().
func (v *window) Unmaximize() {
	C.gtk_window_unmaximize(v.native())
}

// Fullscreen is a wrapper around gtk_window_fullscreen().
func (v *window) Fullscreen() {
	C.gtk_window_fullscreen(v.native())
}

// Unfullscreen is a wrapper around gtk_window_unfullscreen().
func (v *window) Unfullscreen() {
	C.gtk_window_unfullscreen(v.native())
}

// SetKeepAbove is a wrapper around gtk_window_set_keep_above().
func (v *window) SetKeepAbove(setting bool) {
	C.gtk_window_set_keep_above(v.native(), gbool(setting))
}

// SetKeepBelow is a wrapper around gtk_window_set_keep_below().
func (v *window) SetKeepBelow(setting bool) {
	C.gtk_window_set_keep_below(v.native(), gbool(setting))
}

// SetDecorated is a wrapper around gtk_window_set_decorated().
func (v *window) SetDecorated(setting bool) {
	C.gtk_window_set_decorated(v.native(), gbool(setting))
}

// SetDeletable is a wrapper around gtk_window_set_deletable().
func (v *window) SetDeletable(setting bool) {
	C.gtk_window_set_deletable(v.native(), gbool(setting))
}

// SetSkipTaskbarHint is a wrapper around gtk_window_set_skip_taskbar_hint().
func (v *window) SetSkipTaskbarHint(setting bool) {
	C.gtk_window_set_skip_taskbar_hint(v.native(), gbool(setting))
}

// SetSkipPagerHint is a wrapper around gtk_window_set_skip_pager_hint().
func (v *window) SetSkipPagerHint(setting bool) {
	C.gtk_window_set_skip_pager_hint(v.native(), gbool(setting))
}

// SetUrgencyHint is a wrapper around gtk_window_set_urgency_hint().
func (v *window) SetUrgencyHint(setting bool) {
	C.gtk_window_set_urgency_hint(v.native(), gbool(setting))
}

// SetAcceptFocus is a wrapper around gtk_window_set_accept_focus().
func (v *window) SetAcceptFocus(setting bool) {
	C.gtk_window_set_accept_focus(v.native(), gbool(setting))
}

// SetFocusOnMap is a wrapper around gtk_window_set_focus_on_map().
func (v *window) SetFocusOnMap(setting bool) {
	C.gtk_window_set_focus_on_map(v.native(), gbool(setting))
}

// SetStartupID is a wrapper around gtk_window_set_startup_id().
func (v *window) SetStartupID(sid string) {
	cstr := (*C.gchar)(C.CString(sid))
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_window_set_startup_id(v.native(), cstr)
}

// SetRole is a wrapper around gtk_window_set_role().
func (v *window) SetRole(s string) {
	cstr := (*C.gchar)(C.CString(s))
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_window_set_role(v.native(), cstr)
}

// SetWMClass is a wrapper around gtk_window_set_wmclass().
func (v *window) SetWMClass(name, class string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cClass := C.CString(class)
	defer C.free(unsafe.Pointer(cClass))
	C.gtk_window_set_wmclass(v.native(), (*C.gchar)(cName), (*C.gchar)(cClass))
}

// GetDecorated is a wrapper around gtk_window_get_decorated().
func (v *window) GetDecorated() bool {
	c := C.gtk_window_get_decorated(v.native())
	return gobool(c)
}

// GetDeletable is a wrapper around gtk_window_get_deletable().
func (v *window) GetDeletable() bool {
	c := C.gtk_window_get_deletable(v.native())
	return gobool(c)
}

// WindowGetDefaultIconName is a wrapper around gtk_window_get_default_icon_name().
func WindowGetDefaultIconName() (string, error) {
	return stringReturn(C.gtk_window_get_default_icon_name())
}

// GetDefaultSize is a wrapper around gtk_window_get_default_size().
func (v *window) GetDefaultSize() (width, height int) {
	var w, h C.gint
	C.gtk_window_get_default_size(v.native(), &w, &h)
	return int(w), int(h)
}

// GetDestroyWithParent is a wrapper around
// gtk_window_get_destroy_with_parent().
func (v *window) GetDestroyWithParent() bool {
	c := C.gtk_window_get_destroy_with_parent(v.native())
	return gobool(c)
}

// GetHideTitlebarWhenMaximized is a wrapper around
// gtk_window_get_hide_titlebar_when_maximized().
func (v *window) GetHideTitlebarWhenMaximized() bool {
	c := C.gtk_window_get_hide_titlebar_when_maximized(v.native())
	return gobool(c)
}

// GetIcon is a wrapper around gtk_window_get_icon().
func (v *window) GetIcon() (gdk.Pixbuf, error) {
	c := C.gtk_window_get_icon(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	p := &gdk_impl.Pixbuf{wrapObject(unsafe.Pointer(c))}
	return p, nil
}

// GetIconName is a wrapper around gtk_window_get_icon_name().
func (v *window) GetIconName() (string, error) {
	return stringReturn(C.gtk_window_get_icon_name(v.native()))
}

// GetModal is a wrapper around gtk_window_get_modal().
func (v *window) GetModal() bool {
	c := C.gtk_window_get_modal(v.native())
	return gobool(c)
}

// GetPosition is a wrapper around gtk_window_get_position().
func (v *window) GetPosition() (root_x, root_y int) {
	var x, y C.gint
	C.gtk_window_get_position(v.native(), &x, &y)
	return int(x), int(y)
}

func stringReturn(c *C.gchar) (string, error) {
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// GetRole is a wrapper around gtk_window_get_role().
func (v *window) GetRole() (string, error) {
	return stringReturn(C.gtk_window_get_role(v.native()))
}

// GetSize is a wrapper around gtk_window_get_size().
func (v *window) GetSize() (width, height int) {
	var w, h C.gint
	C.gtk_window_get_size(v.native(), &w, &h)
	return int(w), int(h)
}

// GetTitle is a wrapper around gtk_window_get_title().
func (v *window) GetTitle() (string, error) {
	return stringReturn(C.gtk_window_get_title(v.native()))
}

// GetTransientFor is a wrapper around gtk_window_get_transient_for().
func (v *window) GetTransientFor() (gtk.Window, error) {
	c := C.gtk_window_get_transient_for(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWindow(wrapObject(unsafe.Pointer(c))), nil
}

// GetAttachedTo is a wrapper around gtk_window_get_attached_to().
func (v *window) GetAttachedTo() (gtk.Widget, error) {
	c := C.gtk_window_get_attached_to(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWidget(wrapObject(unsafe.Pointer(c))), nil
}

// GetSkipTaskbarHint is a wrapper around gtk_window_get_skip_taskbar_hint().
func (v *window) GetSkipTaskbarHint() bool {
	c := C.gtk_window_get_skip_taskbar_hint(v.native())
	return gobool(c)
}

// GetSkipPagerHint is a wrapper around gtk_window_get_skip_pager_hint().
func (v *window) GetSkipPagerHint() bool {
	c := C.gtk_window_get_skip_taskbar_hint(v.native())
	return gobool(c)
}

// GetUrgencyHint is a wrapper around gtk_window_get_urgency_hint().
func (v *window) GetUrgencyHint() bool {
	c := C.gtk_window_get_urgency_hint(v.native())
	return gobool(c)
}

// GetAcceptFocus is a wrapper around gtk_window_get_accept_focus().
func (v *window) GetAcceptFocus() bool {
	c := C.gtk_window_get_accept_focus(v.native())
	return gobool(c)
}

// GetFocusOnMap is a wrapper around gtk_window_get_focus_on_map().
func (v *window) GetFocusOnMap() bool {
	c := C.gtk_window_get_focus_on_map(v.native())
	return gobool(c)
}

// HasGroup is a wrapper around gtk_window_has_group().
func (v *window) HasGroup() bool {
	c := C.gtk_window_has_group(v.native())
	return gobool(c)
}

// Move is a wrapper around gtk_window_move().
func (v *window) Move(x, y int) {
	C.gtk_window_move(v.native(), C.gint(x), C.gint(y))
}

// Resize is a wrapper around gtk_window_resize().
func (v *window) Resize(width, height int) {
	C.gtk_window_resize(v.native(), C.gint(width), C.gint(height))
}

// ResizeToGeometry is a wrapper around gtk_window_resize_to_geometry().
func (v *window) ResizeToGeometry(width, height int) {
	C.gtk_window_resize_to_geometry(v.native(), C.gint(width), C.gint(height))
}

// WindowSetDefaultIconFromFile is a wrapper around gtk_window_set_default_icon_from_file().
func WindowSetDefaultIconFromFile(file string) error {
	cstr := C.CString(file)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	res := C.gtk_window_set_default_icon_from_file((*C.gchar)(cstr), &err)
	if res == 0 {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

// WindowSetDefaultIconName is a wrapper around gtk_window_set_default_icon_name().
func WindowSetDefaultIconName(s string) {
	cstr := (*C.gchar)(C.CString(s))
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_window_set_default_icon_name(cstr)
}

// SetIconFromFile is a wrapper around gtk_window_set_icon_from_file().
func (v *window) SetIconFromFile(file string) error {
	cstr := C.CString(file)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	res := C.gtk_window_set_icon_from_file(v.native(), (*C.gchar)(cstr), &err)
	if res == 0 {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

// SetIconName is a wrapper around gtk_window_set_icon_name().
func (v *window) SetIconName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_window_set_icon_name(v.native(), (*C.gchar)(cstr))
}

// SetAutoStartupNotification is a wrapper around
// gtk_window_set_auto_startup_notification().
// This doesn't seem write.  Might need to rethink?
/*
func (v *window) SetAutoStartupNotification(setting bool) {
	C.gtk_window_set_auto_startup_notification(gbool(setting))
}
*/

// GetMnemonicsVisible is a wrapper around
// gtk_window_get_mnemonics_visible().
func (v *window) GetMnemonicsVisible() bool {
	c := C.gtk_window_get_mnemonics_visible(v.native())
	return gobool(c)
}

// SetMnemonicsVisible is a wrapper around
// gtk_window_get_mnemonics_visible().
func (v *window) SetMnemonicsVisible(setting bool) {
	C.gtk_window_set_mnemonics_visible(v.native(), gbool(setting))
}

// GetFocusVisible is a wrapper around gtk_window_get_focus_visible().
func (v *window) GetFocusVisible() bool {
	c := C.gtk_window_get_focus_visible(v.native())
	return gobool(c)
}

// SetFocusVisible is a wrapper around gtk_window_set_focus_visible().
func (v *window) SetFocusVisible(setting bool) {
	C.gtk_window_set_focus_visible(v.native(), gbool(setting))
}

// GetApplication is a wrapper around gtk_window_get_application().
func (v *window) GetApplication() (gtk.Application, error) {
	c := C.gtk_window_get_application(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return wrapApplication(wrapObject(unsafe.Pointer(c))), nil
}

// SetApplication is a wrapper around gtk_window_set_application().
func (v *window) SetApplication(a gtk.Application) {
	C.gtk_window_set_application(v.native(), castToApplication(a).native())
}

// TODO gtk_window_activate_key().
// TODO gtk_window_add_mnemonic().
// TODO gtk_window_begin_move_drag().
// TODO gtk_window_begin_resize_drag().
// TODO gtk_window_get_default_icon_list().
// TODO gtk_window_get_group().
// TODO gtk_window_get_icon_list().
// TODO gtk_window_get_mnemonic_modifier().
// TODO gtk_window_get_type_hint().
// TODO gtk_window_get_window_type().
// TODO gtk_window_list_toplevels().
// TODO gtk_window_mnemonic_activate().
// TODO gtk_window_parse_geometry().
// TODO gtk_window_propogate_key_event().
// TODO gtk_window_remove_mnemonic().
// TODO gtk_window_set_attached_to().
// TODO gtk_window_set_default_icon_list().
// TODO gtk_window_set_icon_list().
// TODO gtk_window_set_mnemonic_modifier().
// TODO gtk_window_set_screen().
// TODO gtk_window_set_type_hint().
// TODO gtk_window_get_resize_grip_area().

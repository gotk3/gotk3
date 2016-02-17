// Same copyright and license as the rest of the files in this project
// This file contains style related functions and structures

package impl

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	gdk_impl "github.com/gotk3/gotk3/gdk/impl"
	glib_impl "github.com/gotk3/gotk3/glib/impl"
	"github.com/gotk3/gotk3/gtk"
)

func init() {
	gtk.STYLE_PROVIDER_PRIORITY_FALLBACK = C.GTK_STYLE_PROVIDER_PRIORITY_FALLBACK
	gtk.STYLE_PROVIDER_PRIORITY_THEME = C.GTK_STYLE_PROVIDER_PRIORITY_THEME
	gtk.STYLE_PROVIDER_PRIORITY_SETTINGS = C.GTK_STYLE_PROVIDER_PRIORITY_SETTINGS
	gtk.STYLE_PROVIDER_PRIORITY_APPLICATION = C.GTK_STYLE_PROVIDER_PRIORITY_APPLICATION
	gtk.STYLE_PROVIDER_PRIORITY_USER = C.GTK_STYLE_PROVIDER_PRIORITY_USER
}

/*
 * GtkStyleContext
 */

// StyleContext is a representation of GTK's GtkStyleContext.
type styleContext struct {
	*glib_impl.Object
}

// native returns a pointer to the underlying GtkStyleContext.
func (v *styleContext) native() *C.GtkStyleContext {
	if v == nil || v.Object == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkStyleContext(p)
}

func wrapStyleContext(obj *glib_impl.Object) *styleContext {
	return &styleContext{obj}
}

func (v *styleContext) AddClass(class_name string) {
	cstr := C.CString(class_name)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_style_context_add_class(v.native(), (*C.gchar)(cstr))
}

func (v *styleContext) RemoveClass(class_name string) {
	cstr := C.CString(class_name)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_style_context_remove_class(v.native(), (*C.gchar)(cstr))
}

func fromNativeStyleContext(c *C.GtkStyleContext) (*styleContext, error) {
	if c == nil {
		return nil, nilPtrErr
	}

	obj := wrapObject(unsafe.Pointer(c))
	return wrapStyleContext(obj), nil
}

// GetStyleContext is a wrapper around gtk_widget_get_style_context().
func (v *widget) GetStyleContext() (gtk.StyleContext, error) {
	return fromNativeStyleContext(C.gtk_widget_get_style_context(v.native()))
}

// GetParent is a wrapper around gtk_style_context_get_parent().
func (v *styleContext) GetParent() (gtk.StyleContext, error) {
	return fromNativeStyleContext(C.gtk_style_context_get_parent(v.native()))
}

// GetProperty is a wrapper around gtk_style_context_get_property().
func (v *styleContext) GetProperty2(property string, state gtk.StateFlags) (interface{}, error) {
	cstr := (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(cstr))

	var gval C.GValue
	C.gtk_style_context_get_property(v.native(), cstr, C.GtkStateFlags(state), &gval)
	val := glib_impl.ValueFromNative(unsafe.Pointer(&gval))
	return val.GoValue()
}

// GetStyleProperty is a wrapper around gtk_style_context_get_style_property().
func (v *styleContext) GetStyleProperty(property string) (interface{}, error) {
	cstr := (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(cstr))

	var gval C.GValue
	C.gtk_style_context_get_style_property(v.native(), cstr, &gval)
	val := glib_impl.ValueFromNative(unsafe.Pointer(&gval))
	return val.GoValue()
}

// GetScreen is a wrapper around gtk_style_context_get_screen().
func (v *styleContext) GetScreen() (gdk.Screen, error) {
	c := C.gtk_style_context_get_screen(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	d := &gdk_impl.Screen{wrapObject(unsafe.Pointer(c))}
	return d, nil
}

// GetState is a wrapper around gtk_style_context_get_state().
func (v *styleContext) GetState() gtk.StateFlags {
	return gtk.StateFlags(C.gtk_style_context_get_state(v.native()))
}

// GetColor is a wrapper around gtk_style_context_get_color().
func (v *styleContext) GetColor(state gtk.StateFlags) gdk.RGBA {
	gdkColor := gdk_impl.NewRGBA()
	C.gtk_style_context_get_color(v.native(), C.GtkStateFlags(state), (*C.GdkRGBA)(unsafe.Pointer(gdkColor.Native())))
	return gdkColor
}

// LookupColor is a wrapper around gtk_style_context_lookup_color().
func (v *styleContext) LookupColor(colorName string) (gdk.RGBA, bool) {
	cstr := (*C.gchar)(C.CString(colorName))
	defer C.free(unsafe.Pointer(cstr))
	gdkColor := gdk_impl.NewRGBA()
	ret := C.gtk_style_context_lookup_color(v.native(), cstr, (*C.GdkRGBA)(unsafe.Pointer(gdkColor.Native())))
	return gdkColor, gobool(ret)
}

// StyleContextResetWidgets is a wrapper around gtk_style_context_reset_widgets().
func StyleContextResetWidgets(v *gdk_impl.Screen) {
	C.gtk_style_context_reset_widgets((*C.GdkScreen)(unsafe.Pointer(v.Native())))
}

// Restore is a wrapper around gtk_style_context_restore().
func (v *styleContext) Restore() {
	C.gtk_style_context_restore(v.native())
}

// Save is a wrapper around gtk_style_context_save().
func (v *styleContext) Save() {
	C.gtk_style_context_save(v.native())
}

// SetParent is a wrapper around gtk_style_context_set_parent().
func (v *styleContext) SetParent(p gtk.StyleContext) {
	C.gtk_style_context_set_parent(v.native(), castToStyleContext(p).native())
}

// HasClass is a wrapper around gtk_style_context_has_class().
func (v *styleContext) HasClass(className string) bool {
	cstr := C.CString(className)
	defer C.free(unsafe.Pointer(cstr))

	return gobool(C.gtk_style_context_has_class(v.native(), (*C.gchar)(cstr)))
}

// SetScreen is a wrapper around gtk_style_context_set_screen().
func (v *styleContext) SetScreen(s gdk.Screen) {
	C.gtk_style_context_set_screen(v.native(), (*C.GdkScreen)(unsafe.Pointer(gdk_impl.CastToScreen(s).Native())))
}

// SetState is a wrapper around gtk_style_context_set_state().
func (v *styleContext) SetState(state gtk.StateFlags) {
	C.gtk_style_context_set_state(v.native(), C.GtkStateFlags(state))
}

type IStyleProvider interface {
	toStyleProvider() *C.GtkStyleProvider
}

// AddProvider is a wrapper around gtk_style_context_add_provider().
func (v *styleContext) AddProvider(provider gtk.StyleProvider, prio uint) {
	C.gtk_style_context_add_provider(v.native(), provider.(IStyleProvider).toStyleProvider(), C.guint(prio))
}

// AddProviderForScreen is a wrapper around gtk_style_context_add_provider_for_screen().
func AddProviderForScreen(s *gdk_impl.Screen, provider gtk.StyleProvider, prio uint) {
	C.gtk_style_context_add_provider_for_screen((*C.GdkScreen)(unsafe.Pointer(s.Native())), provider.(IStyleProvider).toStyleProvider(), C.guint(prio))
}

// RemoveProvider is a wrapper around gtk_style_context_remove_provider().
func (v *styleContext) RemoveProvider(provider gtk.StyleProvider) {
	C.gtk_style_context_remove_provider(v.native(), provider.(IStyleProvider).toStyleProvider())
}

// RemoveProviderForScreen is a wrapper around gtk_style_context_remove_provider_for_screen().
func RemoveProviderForScreen(s *gdk_impl.Screen, provider gtk.StyleProvider) {
	C.gtk_style_context_remove_provider_for_screen((*C.GdkScreen)(unsafe.Pointer(s.Native())), provider.(IStyleProvider).toStyleProvider())
}

// GtkStyleContext * 	gtk_style_context_new ()
// void 	gtk_style_context_get ()
// GtkTextDirection 	gtk_style_context_get_direction ()
// GtkJunctionSides 	gtk_style_context_get_junction_sides ()
// const GtkWidgetPath * 	gtk_style_context_get_path ()
// GdkFrameClock * 	gtk_style_context_get_frame_clock ()
// void 	gtk_style_context_get_style ()
// void 	gtk_style_context_get_style_valist ()
// void 	gtk_style_context_get_valist ()
// GtkCssSection * 	gtk_style_context_get_section ()
// void 	gtk_style_context_get_background_color ()
// void 	gtk_style_context_get_border_color ()
// void 	gtk_style_context_get_border ()
// void 	gtk_style_context_get_padding ()
// void 	gtk_style_context_get_margin ()
// const PangoFontDescription * 	gtk_style_context_get_font ()
// void 	gtk_style_context_invalidate ()
// gboolean 	gtk_style_context_state_is_running ()
// GtkIconSet * 	gtk_style_context_lookup_icon_set ()
// void 	gtk_style_context_cancel_animations ()
// void 	gtk_style_context_scroll_animations ()
// void 	gtk_style_context_notify_state_change ()
// void 	gtk_style_context_pop_animatable_region ()
// void 	gtk_style_context_push_animatable_region ()
// void 	gtk_style_context_set_background ()
// void 	gtk_style_context_set_direction ()
// void 	gtk_style_context_set_junction_sides ()
// void 	gtk_style_context_set_path ()
// void 	gtk_style_context_add_region ()
// void 	gtk_style_context_remove_region ()
// gboolean 	gtk_style_context_has_region ()
// GList * 	gtk_style_context_list_regions ()
// void 	gtk_style_context_set_frame_clock ()
// void 	gtk_style_context_set_scale ()
// gint 	gtk_style_context_get_scale ()
// GList * 	gtk_style_context_list_classes ()
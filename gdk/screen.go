package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

/*
 * GdkScreen
 */

// Screen is a representation of GDK's GdkScreen.
type Screen struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkScreen.
func (v *Screen) native() *C.GdkScreen {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkScreen(p)
}

// Native returns a pointer to the underlying GdkScreen.
func (v *Screen) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalScreen(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &Screen{obj}, nil
}

func toScreen(s *C.GdkScreen) (*Screen, error) {
	if s == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(s))}
	return &Screen{obj}, nil
}

// GetRGBAVisual is a wrapper around gdk_screen_get_rgba_visual().
func (v *Screen) GetRGBAVisual() (*Visual, error) {
	c := C.gdk_screen_get_rgba_visual(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	visual := &Visual{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return visual, nil
}

// GetSystemVisual is a wrapper around gdk_screen_get_system_visual().
func (v *Screen) GetSystemVisual() (*Visual, error) {
	c := C.gdk_screen_get_system_visual(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	visual := &Visual{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return visual, nil
}

// GetWidth is a wrapper around gdk_screen_get_width().
func (v *Screen) GetWidth() int {
	c := C.gdk_screen_get_width(v.native())
	return int(c)
}

// GetHeight is a wrapper around gdk_screen_get_height().
func (v *Screen) GetHeight() int {
	c := C.gdk_screen_get_height(v.native())
	return int(c)
}

// ScreenGetDefault is a wrapper aroud gdk_screen_get_default().
func ScreenGetDefault() (*Screen, error) {
	return toScreen(C.gdk_screen_get_default())
}

// IsComposited is a wrapper around gdk_screen_is_composited().
func (v *Screen) IsComposited() bool {
	return gobool(C.gdk_screen_is_composited(v.native()))
}

// GetRootWindow is a wrapper around gdk_screen_get_root_window().
func (v *Screen) GetRootWindow() (*Window, error) {
	return toWindow(C.gdk_screen_get_root_window(v.native()))
}

// GetDisplay is a wrapper around gdk_screen_get_display().
func (v *Screen) GetDisplay() (*Display, error) {
	return toDisplay(C.gdk_screen_get_display(v.native()))
}

// GetNumber is a wrapper around gdk_screen_get_number().
func (v *Screen) GetNumber() int {
	return int(C.gdk_screen_get_number(v.native()))
}

// GetWidthMM is a wrapper around gdk_screen_get_width_mm().
func (v *Screen) GetWidthMM() int {
	return int(C.gdk_screen_get_width_mm(v.native()))
}

// GetHeightMM is a wrapper around gdk_screen_get_height_mm().
func (v *Screen) GetHeightMM() int {
	return int(C.gdk_screen_get_height_mm(v.native()))
}

func toString(c *C.gchar) (string, error) {
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// MakeDisplayName is a wrapper around gdk_screen_make_display_name().
func (v *Screen) MakeDisplayName() (string, error) {
	return toString(C.gdk_screen_make_display_name(v.native()))
}

// GetNMonitors is a wrapper around gdk_screen_get_n_monitors().
func (v *Screen) GetNMonitors() int {
	return int(C.gdk_screen_get_n_monitors(v.native()))
}

// GetPrimaryMonitor is a wrapper around gdk_screen_get_primary_monitor().
func (v *Screen) GetPrimaryMonitor() int {
	return int(C.gdk_screen_get_primary_monitor(v.native()))
}

// GetMonitorAtPoint is a wrapper around gdk_screen_get_monitor_at_point().
func (v *Screen) GetMonitorAtPoint(x, y int) int {
	return int(C.gdk_screen_get_monitor_at_point(v.native(), C.gint(x), C.gint(y)))
}

// GetMonitorAtWindow is a wrapper around gdk_screen_get_monitor_at_window().
func (v *Screen) GetMonitorAtWindow(w *Window) int {
	return int(C.gdk_screen_get_monitor_at_window(v.native(), w.native()))
}

// GetMonitorHeightMM is a wrapper around gdk_screen_get_monitor_height_mm().
func (v *Screen) GetMonitorHeightMM(m int) int {
	return int(C.gdk_screen_get_monitor_height_mm(v.native(), C.gint(m)))
}

// GetMonitorWidthMM is a wrapper around gdk_screen_get_monitor_width_mm().
func (v *Screen) GetMonitorWidthMM(m int) int {
	return int(C.gdk_screen_get_monitor_width_mm(v.native(), C.gint(m)))
}

// GetMonitorPlugName is a wrapper around gdk_screen_get_monitor_plug_name().
func (v *Screen) GetMonitorPlugName(m int) (string, error) {
	return toString(C.gdk_screen_get_monitor_plug_name(v.native(), C.gint(m)))
}

// GetMonitorScaleFactor is a wrapper around gdk_screen_get_monitor_scale_factor().
func (v *Screen) GetMonitorScaleFactor(m int) int {
	return int(C.gdk_screen_get_monitor_scale_factor(v.native(), C.gint(m)))
}

// GetResolution is a wrapper around gdk_screen_get_resolution().
func (v *Screen) GetResolution() float64 {
	return float64(C.gdk_screen_get_resolution(v.native()))
}

// SetResolution is a wrapper around gdk_screen_set_resolution().
func (v *Screen) SetResolution(r float64) {
	C.gdk_screen_set_resolution(v.native(), C.gdouble(r))
}

// GetActiveWindow is a wrapper around gdk_screen_get_active_window().
func (v *Screen) GetActiveWindow() (*Window, error) {
	return toWindow(C.gdk_screen_get_active_window(v.native()))
}

// void 	gdk_screen_set_font_options ()
// gboolean 	gdk_screen_get_setting ()
// const cairo_font_options_t * 	gdk_screen_get_font_options ()
// GList * 	gdk_screen_get_window_stack ()
// GList * 	gdk_screen_list_visuals ()
// GList * 	gdk_screen_get_toplevel_windows ()
// void 	gdk_screen_get_monitor_geometry ()
// void 	gdk_screen_get_monitor_workarea ()

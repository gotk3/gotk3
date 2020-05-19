//+build gtk_3_6 gtk_3_8 gtk_3_10 gtk_3_12 gtk_3_14 gtk_3_16 gtk_3_18 gtk_3_20 gtk_deprecated

package gdk

// #include <gdk/gdk.h>
import "C"

/*
 * Constants
 */

// TODO:
// GdkByteOrder

/*
 * GdkScreen
 */

// GetActiveWindow is a wrapper around gdk_screen_get_active_window().
func (v *Screen) GetActiveWindow() (*Window, error) {

	c := C.gdk_screen_get_active_window(v.native())

	// transfer full -> i.e. don't Ref(), but ensure Unref() via finalizer
	return toWindowWithFinalizer(c)
}

// GetHeight is a wrapper around gdk_screen_get_height().
func (v *Screen) GetHeight() int {
	c := C.gdk_screen_get_height(v.native())
	return int(c)
}

// GetHeightMM is a wrapper around gdk_screen_get_height_mm().
func (v *Screen) GetHeightMM() int {
	return int(C.gdk_screen_get_height_mm(v.native()))
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

// GetMonitorPlugName is a wrapper around gdk_screen_get_monitor_plug_name().
func (v *Screen) GetMonitorPlugName(m int) (string, error) {
	return toString(C.gdk_screen_get_monitor_plug_name(v.native(), C.gint(m)))
}

// GetMonitorScaleFactor is a wrapper around gdk_screen_get_monitor_scale_factor().
func (v *Screen) GetMonitorScaleFactor(m int) int {
	return int(C.gdk_screen_get_monitor_scale_factor(v.native(), C.gint(m)))
}

// GetMonitorWidthMM is a wrapper around gdk_screen_get_monitor_width_mm().
func (v *Screen) GetMonitorWidthMM(m int) int {
	return int(C.gdk_screen_get_monitor_width_mm(v.native(), C.gint(m)))
}

// GetNMonitors is a wrapper around gdk_screen_get_n_monitors().
func (v *Screen) GetNMonitors() int {
	return int(C.gdk_screen_get_n_monitors(v.native()))
}

// GetNumber is a wrapper around gdk_screen_get_number().
func (v *Screen) GetNumber() int {
	return int(C.gdk_screen_get_number(v.native()))
}

// GetPrimaryMonitor is a wrapper around gdk_screen_get_primary_monitor().
func (v *Screen) GetPrimaryMonitor() int {
	return int(C.gdk_screen_get_primary_monitor(v.native()))
}

// GetWidth is a wrapper around gdk_screen_get_width().
func (v *Screen) GetWidth() int {
	c := C.gdk_screen_get_width(v.native())
	return int(c)
}

// GetWidthMM is a wrapper around gdk_screen_get_width_mm().
func (v *Screen) GetWidthMM() int {
	return int(C.gdk_screen_get_width_mm(v.native()))
}

// MakeDisplayName is a wrapper around gdk_screen_make_display_name().
func (v *Screen) MakeDisplayName() (string, error) {
	return toString(C.gdk_screen_make_display_name(v.native()))
}

/*
 * GdkVisuals
 */

// TODO:
// gdk_query_depths().
// gdk_query_visual_types().
// gdk_list_visuals().
// gdk_visual_get_bits_per_rgb().
// gdk_visual_get_byte_order().
// gdk_visual_get_colormap_size().
// gdk_visual_get_best_depth().
// gdk_visual_get_best_type().
// gdk_visual_get_system().
// gdk_visual_get_best().
// gdk_visual_get_best_with_depth().
// gdk_visual_get_best_with_type().
// gdk_visual_get_best_with_both().

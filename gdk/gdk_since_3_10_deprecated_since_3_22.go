//+build gtk_3_10 gtk_3_12 gtk_3_14 gtk_3_16 gtk_3_18 gtk_3_20 gtk_deprecated

package gdk

// #include <gdk/gdk.h>
import "C"


// GetMonitorScaleFactor is a wrapper around gdk_screen_get_monitor_scale_factor().
func (v *Screen) GetMonitorScaleFactor(m int) int {
	return int(C.gdk_screen_get_monitor_scale_factor(v.native(), C.gint(m)))
}


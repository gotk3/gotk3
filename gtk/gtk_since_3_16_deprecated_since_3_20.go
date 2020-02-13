// Same copyright and license as the rest of the files in this project.
// This file is normally only compiled for GTK 3.16 and 3.18.
// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,gtk_deprecated gtk_3_16 gtk_3_18

package gtk

// #include <gtk/gtk.h>
import "C"

/*
 * GtkPopover
 */

// SetTransitionsEnabled is a wrapper gtk_popover_set_transitions_enabled().
func (v *Popover) SetTransitionsEnabled(transitionsEnabled bool) {
	C.gtk_popover_set_transitions_enabled(v.native(), gbool(transitionsEnabled))
}

// GetTransitionsEnabled is a wrapper gtk_popover_get_transitions_enabled().
func (v *Popover) GetTransitionsEnabled() bool {
	return gobool(C.gtk_popover_get_transitions_enabled(v.native()))
}

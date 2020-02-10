// Same copyright and license as the rest of the files in this project

// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

// SetHelpOverlay is a wrapper around gtk_application_window_set_help_overlay().
func (v *ApplicationWindow) SetHelpOverlay(helpOverlay *ShortcutsWindow) {
	C.gtk_application_window_set_help_overlay(v.native(), helpOverlay.native())
}

// GetHelpOverlay is a wrapper around gtk_application_window_get_help_overlay().
func (v *ApplicationWindow) GetHelpOverlay() *ShortcutsWindow {
	c := C.gtk_application_window_get_help_overlay(v.native())
	if c == nil {
		return nil
	}
	return wrapShortcutsWindow(glib.Take(unsafe.Pointer(c)))
}

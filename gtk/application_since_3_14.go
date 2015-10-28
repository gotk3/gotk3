// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12

// See: https://developer.gnome.org/gtk3/3.14/api-index-3-14.html

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"

// PrefersAppMenu is a wrapper around gtk_application_prefers_app_menu().
func (v *Application) PrefersAppMenu() bool {
	return gobool(C.gtk_application_prefers_app_menu(v.native()))
}

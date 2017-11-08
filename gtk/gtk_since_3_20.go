// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18,gtk_3_20

// See: https://developer.gnome.org/gtk3/3.20/api-index-3-20.html

package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "gtk_since_3_20.go.h"
import "C"

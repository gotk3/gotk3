// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"

/*
 * Constants
 */

const (
	LEVEL_BAR_OFFSET_FULL string = C.GTK_LEVEL_BAR_OFFSET_FULL
)

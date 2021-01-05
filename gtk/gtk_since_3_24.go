// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18,!gtk_3_20,!gtk_3_22

package gtk

// #include <gtk/gtk.h>
import "C"

/*
 * GtkInputPurpose
 */

const (
	INPUT_PURPOSE_TERMINAL InputPurpose = C.GTK_INPUT_PURPOSE_TERMINAL
)

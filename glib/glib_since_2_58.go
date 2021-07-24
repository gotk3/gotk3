// Same copyright and license as the rest of the files in this project

// +build !glib_2_40,!glib_2_42,!glib_2_44,!glib_2_46,!glib_2_48,!glib_2_50,!glib_2_52,!glib_2_54,!glib_2_56

package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"

const (
	FORMAT_SIZE_BITS FormatSizeFlags = C.G_FORMAT_SIZE_BITS
)

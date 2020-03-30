// Same copyright and license as the rest of the files in this project

// +build !gdk_pixbuf_2_2,!gdk_pixbuf_2_4,!gdk_pixbuf_2_6,!gdk_pixbuf_2_8,!gdk_pixbuf_2_12,!gdk_pixbuf_2_14

package gdk

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
// #include "pixbuf.go.h"
import "C"

/*
 * PixbufFormat
 */

// TODO:
// gdk_pixbuf_format_copy().

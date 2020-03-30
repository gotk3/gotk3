// Same copyright and license as the rest of the files in this project

// +build !gdk_pixbuf_2_2,!gdk_pixbuf_2_4,!gdk_pixbuf_2_6,!gdk_pixbuf_2_8,!gdk_pixbuf_2_12,!gdk_pixbuf_2_14,!gdk_pixbuf_2_22,!gdk_pixbuf_2_24,!gdk_pixbuf_2_26,!gdk_pixbuf_2_28,!gdk_pixbuf_2_30,!gdk_pixbuf_2_32

package gdk

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
// #include "pixbuf.go.h"
import "C"

// File saving

// TODO:
// gdk_pixbuf_save_to_streamv().
// gdk_pixbuf_save_to_streamv_async().

// The GdkPixbuf Structure

// TODO:
// gdk_pixbuf_remove_option().
// gdk_pixbuf_copy_options().

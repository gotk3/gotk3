// Same copyright and license as the rest of the files in this project

// +build !gdk_pixbuf_2_2,!gdk_pixbuf_2_4,!gdk_pixbuf_2_6,!gdk_pixbuf_2_8,!gdk_pixbuf_2_12,!gdk_pixbuf_2_14,!gdk_pixbuf_2_22

package gdk

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
// #include "pixbuf.go.h"
import "C"

// File Loading

// TODO:
// gdk_pixbuf_new_from_stream_finish().
// gdk_pixbuf_new_from_stream_at_scale_async().

// File saving

// TODO:
// gdk_pixbuf_save_to_stream_async().
// gdk_pixbuf_save_to_stream_finish().

// Same copyright and license as the rest of the files in this project

//+build gdk_pixbuf_2_2 gdk_pixbuf_2_4 gdk_pixbuf_2_6 gdk_pixbuf_2_8 gdk_pixbuf_2_12 gdk_pixbuf_2_14 gdk_pixbuf_2_24 gdk_pixbuf_2_26 gdk_pixbuf_2_28 gdk_pixbuf_2_30 gdk_pixbuf_deprecated

package gdk

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
// #include "pixbuf.go.h"
import "C"

// Image Data in Memory

// TODO:
// gdk_pixbuf_new_from_inline().

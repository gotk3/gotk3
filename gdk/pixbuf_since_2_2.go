// Same copyright and license as the rest of the files in this project

package gdk

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
// #include "pixbuf.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

// The GdkPixbuf Structure

// TODO:
// gdk_pixbuf_set_option().

/*
 * GdkPixbufLoader
 */

// SetSize is a wrapper around gdk_pixbuf_loader_set_size().
func (v *PixbufLoader) SetSize(width, height int) {
	C.gdk_pixbuf_loader_set_size(v.native(), C.int(width), C.int(height))
}

/*
 * PixbufFormat
 */

// PixbufGetFormats is a wrapper around gdk_pixbuf_get_formats().
func PixbufGetFormats() []*PixbufFormat {
	l := (*C.struct__GSList)(C.gdk_pixbuf_get_formats())
	formats := glib.WrapSList(uintptr(unsafe.Pointer(l)))
	if formats == nil {
		return nil // no error. A nil list is considered to be empty.
	}

	// "The structures themselves are owned by GdkPixbuf". Free the list only.
	defer formats.Free()

	ret := make([]*PixbufFormat, 0, formats.Length())
	formats.Foreach(func(item interface{}) {
		ret = append(
			ret,
			&PixbufFormat{
				(*C.GdkPixbufFormat)(item.(unsafe.Pointer))})
	})

	return ret
}

// GetName is a wrapper around gdk_pixbuf_format_get_name().
func (f *PixbufFormat) GetName() (string, error) {
	c := C.gdk_pixbuf_format_get_name(f.native())
	return C.GoString((*C.char)(c)), nil
}

// GetDescription is a wrapper around gdk_pixbuf_format_get_description().
func (f *PixbufFormat) GetDescription() (string, error) {
	c := C.gdk_pixbuf_format_get_description(f.native())
	return C.GoString((*C.char)(c)), nil
}

// GetMimeTypes is a wrapper around gdk_pixbuf_format_get_mime_types().
func (f *PixbufFormat) GetMimeTypes() []string {
	var types []string
	c := C.gdk_pixbuf_format_get_mime_types(f.native())
	if c == nil {
		return nil
	}
	for *c != nil {
		types = append(types, C.GoString((*C.char)(*c)))
		c = C.next_gcharptr(c)
	}
	return types
}

// GetExtensions is a wrapper around gdk_pixbuf_format_get_extensions().
func (f *PixbufFormat) GetExtensions() []string {
	var extensions []string
	c := C.gdk_pixbuf_format_get_extensions(f.native())
	if c == nil {
		return nil
	}
	for *c != nil {
		extensions = append(extensions, C.GoString((*C.char)(*c)))
		c = C.next_gcharptr(c)
	}
	return extensions
}

// GetLicense is a wrapper around gdk_pixbuf_format_get_license().
func (f *PixbufFormat) GetLicense() (string, error) {
	c := C.gdk_pixbuf_format_get_license(f.native())
	return C.GoString((*C.char)(c)), nil
}

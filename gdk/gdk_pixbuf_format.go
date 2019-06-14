package gdk

// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

type PixbufFormat struct {
	format *C.GdkPixbufFormat
}

// native returns a pointer to the underlying GdkPixbuf.
func (v *PixbufFormat) native() *C.GdkPixbufFormat {
	if v == nil {
		return nil
	}

	return v.format
}

// Native returns a pointer to the underlying GdkPixbuf.
func (v *PixbufFormat) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (f *PixbufFormat) GetName() (string, error) {
	c := C.gdk_pixbuf_format_get_name(f.native())
	return C.GoString((*C.char)(c)), nil
}

func (f *PixbufFormat) GetDescription() (string, error) {
	c := C.gdk_pixbuf_format_get_description(f.native())
	return C.GoString((*C.char)(c)), nil
}

func (f *PixbufFormat) GetLicense() (string, error) {
	c := C.gdk_pixbuf_format_get_license(f.native())
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

func PixbufGetFormats() []*PixbufFormat {
	l := (*C.struct__GSList)(C.gdk_pixbuf_get_formats())
	formats := glib.WrapSList(uintptr(unsafe.Pointer(l)))
	if formats == nil {
		return nil // no error. A nil list is considered to be empty.
	}

	// "The structures themselves are owned by GdkPixbuf". Free the list only.
	defer formats.Free()

	ret := make([]*PixbufFormat, 0, formats.Length())
	formats.Foreach(func(ptr unsafe.Pointer) {
		ret = append(ret, &PixbufFormat{(*C.GdkPixbufFormat)(ptr)})
	})

	return ret
}

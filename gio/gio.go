package gio

// #cgo pkg-config: gio-2.0 glib-2.0 gobject-2.0
// #include <gio/gio.h>
// #include <stdlib.h>
// #include "gio.go.h"
import "C"
import (
	"errors"
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {

	tm := []glib.TypeMarshaler{
		{glib.Type(C.g_icon_get_type()), marshalIcon},
		{glib.Type(C.g_file_get_type()), marshalFile},
		{glib.Type(C.g_file_icon_get_type()), marshalFileIcon},
	}

	glib.RegisterGValueMarshalers(tm)
}

/*
 * Unexported vars
 */

var nilPtrErr = errors.New("cgo returned unexpected nil pointer")

/*
 * GIcon
 */

// Icon is a representation of GIO's GIcon.
// Interface for icons
type Icon struct {
	*glib.Object
}

// native returns a pointer to the underlying GIcon.
func (v *Icon) native() *C.GIcon {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGIcon(p)
}

// NativePrivate: to be used inside Gotk3 only.
func (v *Icon) NativePrivate() *C.GIcon {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGIcon(p)
}

// Native returns a pointer to the underlying GIcon.
func (v *Icon) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalIcon(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapIcon(obj), nil
}

func wrapIcon(obj *glib.Object) *Icon {
	return &Icon{obj}
}

// TODO I dont know how to handle it ...
/*
guint
g_icon_hash (gconstpointer icon);
*/

// Equal is a wrapper around g_icon_equal().
func (v *Icon) Equal(icon *Icon) bool {
	return gobool(C.g_icon_equal(v.native(), icon.native()))
}

// ToString is a wrapper around g_icon_to_string().
func (v *Icon) ToString() string {
	return goString(C.g_icon_to_string(v.native()))
}

// IconNewForString is a wrapper around g_icon_new_for_string().
func IconNewForString(str string) (*Icon, error) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError
	c := C.g_icon_new_for_string((*C.char)(cstr), &err)
	if c == nil {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	i := &Icon{obj}

	runtime.SetFinalizer(i, func(_ interface{}) { obj.Unref() })
	return i, nil
}

// TODO Requiere GVariant
/*
GVariant * 	g_icon_serialize ()
GIcon * 	g_icon_deserialize ()
*/

/*
 * GFileIcon
 */

// FileIcon is a representation of GIO's GFileIcon.
type FileIcon struct {
	*glib.Object
}

// native returns a pointer to the underlying GFileIcon.
func (v *FileIcon) native() *C.GFileIcon {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGFileIcon(p)
}

// NativePrivate: to be used inside Gotk3 only.
func (v *FileIcon) NativePrivate() *C.GFileIcon {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGFileIcon(p)
}

// Native returns a pointer to the underlying GFileIcon.
func (v *FileIcon) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalFileIcon(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFileIcon(obj), nil
}

func wrapFileIcon(obj *glib.Object) *FileIcon {
	return &FileIcon{obj}
}

// FileIconNew is a wrapper around g_file_icon_new().
func FileIconNew(file *File) (*Icon, error) {

	c := C.g_file_icon_new(file.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapIcon(glib.Take(unsafe.Pointer(c))), nil
}

/*
 * GFile
 */

// File is a representation of GIO's GFile.
type File struct {
	*glib.Object
}

// native returns a pointer to the underlying GFile.
func (v *File) native() *C.GFile {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGFile(p)
}

// NativePrivate: to be used inside Gotk3 only.
func (v *File) NativePrivate() *C.GFile {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGFile(p)
}

// Native returns a pointer to the underlying GFile.
func (v *File) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalFile(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFile(obj), nil
}

func wrapFile(obj *glib.Object) *File {
	return &File{obj}
}

// FileNewForPath is a wrapper around g_file_new_for_path().
func FileNewForPath(path string) (*File, error) {
	cstr := (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_file_new_for_path(cstr)
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapFile(glib.Take(unsafe.Pointer(c))), nil
}

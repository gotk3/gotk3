package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

func init() {

	tm := []TypeMarshaler{
		{Type(C.g_file_get_type()), marshalFile},
		{Type(C.g_file_icon_get_type()), marshalFileIcon},
	}

	RegisterGValueMarshalers(tm)
}

/*
 * GIcon
 */

// Icon is a representation of GIO's GIcon.
// Interface for icons
type Icon struct {
	*Object
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
	obj := Take(unsafe.Pointer(c))
	return wrapIcon(obj), nil
}

func wrapIcon(obj *Object) *Icon {
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
	var s string
	if c := C.g_icon_to_string(v.native()); c != nil {
		s = goString(c)
		defer C.g_free((C.gpointer)(c))
	}

	return s
}

// IconNewForString is a wrapper around g_icon_new_for_string().
func IconNewForString(str string) (*Icon, error) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError
	c := C.g_icon_new_for_string((*C.gchar)(cstr), &err)
	if c == nil {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}

	obj := &Object{ToGObject(unsafe.Pointer(c))}
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
	*Object
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
	obj := Take(unsafe.Pointer(c))
	return wrapFileIcon(obj), nil
}

func wrapFileIcon(obj *Object) *FileIcon {
	return &FileIcon{obj}
}

// FileIconNewN is a wrapper around g_file_icon_new().
// This version respect Gtk3 documentation.
func FileIconNewN(file *File) (*Icon, error) {

	c := C.g_file_icon_new(file.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapIcon(Take(unsafe.Pointer(c))), nil
}

// FileIconNew is a wrapper around g_file_icon_new().
// To not break previous implementation of GFileIcon ...
func FileIconNew(path string) *Icon {
	file := FileNew(path)

	c := C.g_file_icon_new(file.native())
	if c == nil {
		return nil
	}
	return wrapIcon(Take(unsafe.Pointer(c)))
}

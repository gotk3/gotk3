package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

type File struct {
  *Object
}

// Native() returns a pointer to the underlying GFile.
func (v *File ) native() *C.GFile  {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGFile(unsafe.Pointer(v.GObject))
}

// FileNew is a wrapper around g_file_new_for_path().
func FileNew(title string) *File {
	cstr1 := (*C.char)(C.CString(title))
	defer C.free(unsafe.Pointer(cstr1))

	c := C.g_file_new_for_path(cstr1)
	if c == nil {
		return nil
	}
	return wrapFile(wrapObject(unsafe.Pointer(c)))
}

func wrapFile(obj *Object) *File {
	return &File{obj}
}

package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

/*
 * GFile
 */

// File is a representation of GIO's GFile.
type File struct {
  *Object
}

// native returns a pointer to the underlying GFile.
func (v *File ) native() *C.GFile  {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGFile(unsafe.Pointer(v.GObject))
}

// Native returns a pointer to the underlying GFile.
func (v *File ) Native() uintptr  {
	return uintptr(unsafe.Pointer(v.native()))
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

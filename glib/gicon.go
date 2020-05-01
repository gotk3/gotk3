package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

type FileIcon struct {
  *Object
}

// native() returns a pointer to the underlying GFileIcon.
func (v *FileIcon) native() *C.GFileIcon {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGFileIcon(unsafe.Pointer(v.GObject))
}

// FileIconNew is a wrapper around g_file_icon_new().
func FileIconNew(path string) *FileIcon {
  file := FileNew(path)

	c := C.g_file_icon_new(file.native())
	if c == nil {
		return nil
	}
	return wrapFileIcon(wrapObject(unsafe.Pointer(c)))
}

func wrapFileIcon(obj *Object) *FileIcon {
	return &FileIcon{obj}
}

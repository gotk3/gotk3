package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
// #include "gfile.go.h"
import "C"
import (
	"errors"
	"unsafe"
)

func init() {

	tm := []TypeMarshaler{
		{Type(C.g_file_get_type()), marshalFile},
		{Type(C.g_file_input_stream_get_type()), marshalFileInputStream},
		{Type(C.g_file_output_stream_get_type()), marshalFileOutputStream},
	}

	RegisterGValueMarshalers(tm)
}

func goString(cstr *C.gchar) string {
	return C.GoString((*C.char)(cstr))
}

/*
 * GFile
 */

// File is a representation of GIO's GFile.
type File struct {
	*Object
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
	obj := Take(unsafe.Pointer(c))
	return wrapFile(obj), nil
}

func wrapFile(obj *Object) *File {
	return &File{obj}
}

// FileNew is a wrapper around g_file_new_for_path().
// To avoid breaking previous implementation of GFile ...
func FileNew(path string) *File {
	f, e := FileNewForPath(path)
	if e != nil {
		return nil
	}
	return f
}

// FileNewForPath is a wrapper around g_file_new_for_path().
func FileNewForPath(path string) (*File, error) {
	cstr := (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_file_new_for_path(cstr)
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapFile(Take(unsafe.Pointer(c))), nil
}

// TODO g_file_*** and more
/*
void 	(*GFileProgressCallback) ()
gboolean 	(*GFileReadMoreCallback) ()
void 	(*GFileMeasureProgressCallback) ()
GFile * 	g_file_new_for_uri ()
GFile * 	g_file_new_for_commandline_arg ()
GFile * 	g_file_new_for_commandline_arg_and_cwd ()
GFile * 	g_file_new_tmp ()
GFile * 	g_file_parse_name ()
GFile * 	g_file_new_build_filename ()
GFile * 	g_file_dup ()
guint 	g_file_hash ()
gboolean 	g_file_equal ()
char * 	g_file_get_basename ()
*/

/*
char *
g_file_get_path (GFile *file);
*/
// GetPath is a wrapper around g_file_get_path().
func (v *File) GetPath() string {
	var s string
	if c := C.g_file_get_path(v.native()); c != nil {
		s = C.GoString(c)
		defer C.g_free((C.gpointer)(c))
	}

	return s
}

/*
const char * 	g_file_peek_path ()
char * 	g_file_get_uri ()
char * 	g_file_get_parse_name ()
GFile * 	g_file_get_parent ()
gboolean 	g_file_has_parent ()
GFile * 	g_file_get_child ()
GFile * 	g_file_get_child_for_display_name ()
gboolean 	g_file_has_prefix ()
char * 	g_file_get_relative_path ()
GFile * 	g_file_resolve_relative_path ()
gboolean 	g_file_is_native ()
gboolean 	g_file_has_uri_scheme ()
char * 	g_file_get_uri_scheme ()
*/

/*
GFileInputStream *
g_file_read (GFile *file,
             GCancellable *cancellable,
             GError **error);
*/
// Read is a wrapper around g_file_read().
// Object.Unref() must be used after use
func (v *File) Read(cancellable *Cancellable) (*FileInputStream, error) {
	var gerr *C.GError
	c := C.g_file_read(
		v.native(),
		cancellable.native(),
		&gerr)
	if c == nil {
		defer C.g_error_free(gerr)
		return nil, errors.New(goString(gerr.message))
	}
	return wrapFileInputStream(Take(unsafe.Pointer(c))), nil
}

/*
void 	g_file_read_async ()
GFileInputStream * 	g_file_read_finish ()
GFileOutputStream * 	g_file_append_to ()
GFileOutputStream * 	g_file_create ()
GFileOutputStream * 	g_file_replace ()
void 	g_file_append_to_async ()
GFileOutputStream * 	g_file_append_to_finish ()
void 	g_file_create_async ()
GFileOutputStream * 	g_file_create_finish ()
void 	g_file_replace_async ()
GFileOutputStream * 	g_file_replace_finish ()
GFileInfo * 	g_file_query_info ()
void 	g_file_query_info_async ()
GFileInfo * 	g_file_query_info_finish ()
gboolean 	g_file_query_exists ()
GFileType 	g_file_query_file_type ()
GFileInfo * 	g_file_query_filesystem_info ()
void 	g_file_query_filesystem_info_async ()
GFileInfo * 	g_file_query_filesystem_info_finish ()
GAppInfo * 	g_file_query_default_handler ()
void 	g_file_query_default_handler_async ()
GAppInfo * 	g_file_query_default_handler_finish ()
gboolean 	g_file_measure_disk_usage ()
void 	g_file_measure_disk_usage_async ()
gboolean 	g_file_measure_disk_usage_finish ()
GMount * 	g_file_find_enclosing_mount ()
void 	g_file_find_enclosing_mount_async ()
GMount * 	g_file_find_enclosing_mount_finish ()
GFileEnumerator * 	g_file_enumerate_children ()
void 	g_file_enumerate_children_async ()
GFileEnumerator * 	g_file_enumerate_children_finish ()
GFile * 	g_file_set_display_name ()
void 	g_file_set_display_name_async ()
GFile * 	g_file_set_display_name_finish ()
gboolean 	g_file_delete ()
void 	g_file_delete_async ()
gboolean 	g_file_delete_finish ()
gboolean 	g_file_trash ()
void 	g_file_trash_async ()
gboolean 	g_file_trash_finish ()
gboolean 	g_file_copy ()
void 	g_file_copy_async ()
gboolean 	g_file_copy_finish ()
gboolean 	g_file_move ()
gboolean 	g_file_make_directory ()
void 	g_file_make_directory_async ()
gboolean 	g_file_make_directory_finish ()
gboolean 	g_file_make_directory_with_parents ()
gboolean 	g_file_make_symbolic_link ()
GFileAttributeInfoList * 	g_file_query_settable_attributes ()
GFileAttributeInfoList * 	g_file_query_writable_namespaces ()
gboolean 	g_file_set_attribute ()
gboolean 	g_file_set_attributes_from_info ()
void 	g_file_set_attributes_async ()
gboolean 	g_file_set_attributes_finish ()
gboolean 	g_file_set_attribute_string ()
gboolean 	g_file_set_attribute_byte_string ()
gboolean 	g_file_set_attribute_uint32 ()
gboolean 	g_file_set_attribute_int32 ()
gboolean 	g_file_set_attribute_uint64 ()
gboolean 	g_file_set_attribute_int64 ()
void 	g_file_mount_mountable ()
GFile * 	g_file_mount_mountable_finish ()
void 	g_file_unmount_mountable ()
gboolean 	g_file_unmount_mountable_finish ()
void 	g_file_unmount_mountable_with_operation ()
gboolean 	g_file_unmount_mountable_with_operation_finish ()
void 	g_file_eject_mountable ()
gboolean 	g_file_eject_mountable_finish ()
void 	g_file_eject_mountable_with_operation ()
gboolean 	g_file_eject_mountable_with_operation_finish ()
void 	g_file_start_mountable ()
gboolean 	g_file_start_mountable_finish ()
void 	g_file_stop_mountable ()
gboolean 	g_file_stop_mountable_finish ()
void 	g_file_poll_mountable ()
gboolean 	g_file_poll_mountable_finish ()
void 	g_file_mount_enclosing_volume ()
gboolean 	g_file_mount_enclosing_volume_finish ()
GFileMonitor * 	g_file_monitor_directory ()
GFileMonitor * 	g_file_monitor_file ()
GFileMonitor * 	g_file_monitor ()
GBytes * 	g_file_load_bytes ()
void 	g_file_load_bytes_async ()
GBytes * 	g_file_load_bytes_finish ()
gboolean 	g_file_load_contents ()
void 	g_file_load_contents_async ()
gboolean 	g_file_load_contents_finish ()
void 	g_file_load_partial_contents_async ()
gboolean 	g_file_load_partial_contents_finish ()
gboolean 	g_file_replace_contents ()
void 	g_file_replace_contents_async ()
void 	g_file_replace_contents_bytes_async ()
gboolean 	g_file_replace_contents_finish ()
gboolean 	g_file_copy_attributes ()
GFileIOStream * 	g_file_create_readwrite ()
void 	g_file_create_readwrite_async ()
GFileIOStream * 	g_file_create_readwrite_finish ()
GFileIOStream * 	g_file_open_readwrite ()
void 	g_file_open_readwrite_async ()
GFileIOStream * 	g_file_open_readwrite_finish ()
GFileIOStream * 	g_file_replace_readwrite ()
void 	g_file_replace_readwrite_async ()
GFileIOStream * 	g_file_replace_readwrite_finish ()
gboolean 	g_file_supports_thread_contexts ()
*/

/*
 * GFileInputStream
 */

// FileInputStream is a representation of GIO's GFileInputStream.
type FileInputStream struct {
	*InputStream
}

// native returns a pointer to the underlying GFileInputStream.
func (v *FileInputStream) native() *C.GFileInputStream {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGFileInputStream(p)
}

// NativePrivate: to be used inside Gotk3 only.
func (v *FileInputStream) NativePrivate() *C.GFileInputStream {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGFileInputStream(p)
}

// Native returns a pointer to the underlying GFileInputStream.
func (v *FileInputStream) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalFileInputStream(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	return wrapFileInputStream(obj), nil
}

func wrapFileInputStream(obj *Object) *FileInputStream {
	return &FileInputStream{wrapInputStream(obj)}
}

// TODO g_file_input_stream_query_info and more
/*
GFileInfo * 	g_file_input_stream_query_info ()
void 	g_file_input_stream_query_info_async ()
GFileInfo * 	g_file_input_stream_query_info_finish ()
*/

/*
 * GFileOutputStream
 */

// FileOutputStream is a representation of GIO's GFileOutputStream.
type FileOutputStream struct {
	*OutputStream
}

// native returns a pointer to the underlying GFileOutputStream.
func (v *FileOutputStream) native() *C.GFileOutputStream {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGFileOutputStream(p)
}

// NativePrivate: to be used inside Gotk3 only.
func (v *FileOutputStream) NativePrivate() *C.GFileOutputStream {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGFileOutputStream(p)
}

// Native returns a pointer to the underlying GFileOutputStream.
func (v *FileOutputStream) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalFileOutputStream(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	return wrapFileOutputStream(obj), nil
}

func wrapFileOutputStream(obj *Object) *FileOutputStream {
	return &FileOutputStream{wrapOutputStream(obj)}
}

// TODO g_file_output_stream_query_info and more
/*
GFileInfo * 	g_file_output_stream_query_info ()
void 	g_file_output_stream_query_info_async ()
GFileInfo * 	g_file_output_stream_query_info_finish ()
char * 	g_file_output_stream_get_etag ()
*/

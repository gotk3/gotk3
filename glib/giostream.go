package glib

// #cgo pkg-config: gio-2.0 glib-2.0 gobject-2.0
// #include <gio/gio.h>
// #include <stdlib.h>
// #include "giostream.go.h"
import "C"
import (
	"bytes"
	"errors"
	"unsafe"
)

func init() {

	tm := []TypeMarshaler{
		{Type(C.g_io_stream_get_type()), marshalIOStream},
		{Type(C.g_output_stream_get_type()), marshalOutputStream},
		{Type(C.g_input_stream_get_type()), marshalInputStream},
	}

	RegisterGValueMarshalers(tm)
}

// OutputStreamSpliceFlags is a representation of GTK's GOutputStreamSpliceFlags.
type OutputStreamSpliceFlags int

const (
	OUTPUT_STREAM_SPLICE_NONE         OutputStreamSpliceFlags = C.G_OUTPUT_STREAM_SPLICE_NONE
	OUTPUT_STREAM_SPLICE_CLOSE_SOURCE                         = C.G_OUTPUT_STREAM_SPLICE_CLOSE_SOURCE
	OUTPUT_STREAM_SPLICE_CLOSE_TARGET                         = C.G_OUTPUT_STREAM_SPLICE_CLOSE_TARGET
)

/*
 * GIOStream
 */

// IOStream is a representation of GIO's GIOStream.
// Base class for implementing read/write streams
type IOStream struct {
	*Object
}

// native returns a pointer to the underlying GIOStream.
func (v *IOStream) native() *C.GIOStream {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGIOStream(p)
}

// NativePrivate: to be used inside Gotk3 only.
func (v *IOStream) NativePrivate() *C.GIOStream {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGIOStream(p)
}

// Native returns a pointer to the underlying GIOStream.
func (v *IOStream) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalIOStream(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	return wrapIOStream(obj), nil
}

func wrapIOStream(obj *Object) *IOStream {
	return &IOStream{obj}
}

/*
GInputStream * 	g_io_stream_get_input_stream ()
GOutputStream * 	g_io_stream_get_output_stream ()
void 	g_io_stream_splice_async ()
gboolean 	g_io_stream_splice_finish ()
*/

// Close is a wrapper around g_io_stream_close().
func (v *IOStream) Close(cancellable *Cancellable) (bool, error) {
	var gerr *C.GError
	ok := gobool(C.g_io_stream_close(
		v.native(),
		cancellable.native(),
		&gerr))
	if !ok {
		defer C.g_error_free(gerr)
		return false, errors.New(goString(gerr.message))
	}
	return ok, nil
}

/*
void 	g_io_stream_close_async ()
gboolean 	g_io_stream_close_finish ()
gboolean 	g_io_stream_is_closed ()
gboolean 	g_io_stream_has_pending ()
gboolean 	g_io_stream_set_pending ()
void 	g_io_stream_clear_pending ()
*/

/*
 * GInputStream
 */

// InputStream is a representation of GIO's GInputStream.
// Base class for implementing streaming input
type InputStream struct {
	*Object
}

// native returns a pointer to the underlying GInputStream.
func (v *InputStream) native() *C.GInputStream {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGInputStream(p)
}

// NativePrivate: to be used inside Gotk3 only.
func (v *InputStream) NativePrivate() *C.GInputStream {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGInputStream(p)
}

// Native returns a pointer to the underlying GInputStream.
func (v *InputStream) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalInputStream(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	return wrapInputStream(obj), nil
}

func wrapInputStream(obj *Object) *InputStream {
	return &InputStream{obj}
}

// Read is a wrapper around g_input_stream_read().
func (v *InputStream) Read(length uint, cancellable *Cancellable) (*bytes.Buffer, int, error) {
	var gerr *C.GError
	var buffer = bytes.NewBuffer(make([]byte, length))

	c := C.g_input_stream_read(
		v.native(),
		unsafe.Pointer(&buffer.Bytes()[0]),
		C.gsize(length),
		cancellable.native(),
		&gerr)
	if c == -1 {
		defer C.g_error_free(gerr)
		return nil, -1, errors.New(goString(gerr.message))
	}
	return buffer, int(c), nil
}

// TODO find a way to get size to be read without asking for ...
/*
gboolean
g_input_stream_read_all (GInputStream *stream,
                         void *buffer,
                         gsize count,
                         gsize *bytes_read,
                         GCancellable *cancellable,
                         GError **error);
*/

/*
void 	g_input_stream_read_all_async ()
gboolean 	g_input_stream_read_all_finish ()
gssize 	g_input_stream_skip ()
*/

// Close is a wrapper around g_input_stream_close().
func (v *InputStream) Close(cancellable *Cancellable) (bool, error) {
	var gerr *C.GError
	ok := gobool(C.g_input_stream_close(
		v.native(),
		cancellable.native(),
		&gerr))
	if !ok {
		defer C.g_error_free(gerr)
		return false, errors.New(goString(gerr.message))
	}
	return ok, nil
}

// TODO g_input_stream***
/*
void 	g_input_stream_read_async ()
gssize 	g_input_stream_read_finish ()
void 	g_input_stream_skip_async ()
gssize 	g_input_stream_skip_finish ()
void 	g_input_stream_close_async ()
gboolean 	g_input_stream_close_finish ()
*/

// IsClosed is a wrapper around g_input_stream_is_closed().
func (v *InputStream) IsClosed() bool {
	return gobool(C.g_input_stream_is_closed(v.native()))
}

// HasPending is a wrapper around g_input_stream_has_pending().
func (v *InputStream) HasPending() bool {
	return gobool(C.g_input_stream_has_pending(v.native()))
}

// SetPending is a wrapper around g_input_stream_set_pending().
func (v *InputStream) SetPending() (bool, error) {
	var gerr *C.GError
	ok := gobool(C.g_input_stream_set_pending(
		v.native(),
		&gerr))
	if !ok {
		defer C.g_error_free(gerr)
		return false, errors.New(goString(gerr.message))
	}
	return ok, nil
}

// ClearPending is a wrapper around g_input_stream_clear_pending().
func (v *InputStream) ClearPending() {
	C.g_input_stream_clear_pending(v.native())
}

/* Useless functions due to Go language specification and actual
   implementation of (*InputStream).Read that do same thing.

GBytes * 	g_input_stream_read_bytes ()
void 	g_input_stream_read_bytes_async ()
GBytes * 	g_input_stream_read_bytes_finish ()
*/

/*
 * GOutputStream
 */

// OutputStream is a representation of GIO's GOutputStream.
// Base class for implementing streaming output
type OutputStream struct {
	*Object
}

// native returns a pointer to the underlying GOutputStream.
func (v *OutputStream) native() *C.GOutputStream {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGOutputStream(p)
}

// NativePrivate: to be used inside Gotk3 only.
func (v *OutputStream) NativePrivate() *C.GOutputStream {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGOutputStream(p)
}

// Native returns a pointer to the underlying GOutputStream.
func (v *OutputStream) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalOutputStream(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := Take(unsafe.Pointer(c))
	return wrapOutputStream(obj), nil
}

func wrapOutputStream(obj *Object) *OutputStream {
	return &OutputStream{obj}
}

/*
gssize
g_output_stream_write (GOutputStream *stream,
                       const void *buffer,
                       gsize count,
                       GCancellable *cancellable,
                       GError **error);
*/

// Write is a wrapper around g_output_stream_write().
// buffer := bytes.NewBuffer(make([]byte, length))
func (v *OutputStream) Write(buffer *bytes.Buffer, cancellable *Cancellable) (int, error) {
	var gerr *C.GError
	length := buffer.Len()

	c := C.g_output_stream_write(
		v.native(),
		unsafe.Pointer(&buffer.Bytes()[0]),
		C.gsize(length),
		cancellable.native(),
		&gerr)
	if c == -1 {
		defer C.g_error_free(gerr)
		return -1, errors.New(goString(gerr.message))
	}
	return int(c), nil
}

// Write is a wrapper around g_output_stream_write().
// func (v *OutputStream) Write(buffer *[]byte, cancellable *Cancellable) (int, error) {
// 	// cdata := C.CString(data)
// 	// defer C.free(unsafe.Pointer(cdata))
// 	var gerr *C.GError
// 	c := C.g_output_stream_write(
// 		v.native(),
// 		unsafe.Pointer(buffer),
// 		C.gsize(len(*buffer)),
// 		cancellable.native(),
// 		&gerr)
// 	if c == -1 {
// 		defer C.g_error_free(gerr)
// 		return 0, errors.New(goString(gerr.message))
// 	}
// 	return int(c), nil
// }

/*
gboolean 	g_output_stream_write_all ()
*/

// TODO outputStream asynch functions
/*
void 	g_output_stream_write_all_async ()
gboolean 	g_output_stream_write_all_finish ()
gboolean 	g_output_stream_writev ()
gboolean 	g_output_stream_writev_all ()
void 	g_output_stream_writev_async ()
gboolean 	g_output_stream_writev_finish ()
void 	g_output_stream_writev_all_async ()
gboolean 	g_output_stream_writev_all_finish ()
*/
/*
gssize
g_output_stream_splice (GOutputStream *stream,
                        GInputStream *source,
                        GOutputStreamSpliceFlags flags,
                        GCancellable *cancellable,
                        GError **error);
*/

// Flush is a wrapper around g_output_stream_flush().
func (v *OutputStream) Flush(cancellable *Cancellable) (bool, error) {
	var gerr *C.GError
	ok := gobool(C.g_output_stream_flush(
		v.native(),
		cancellable.native(),
		&gerr))
	if !ok {
		defer C.g_error_free(gerr)
		return false, errors.New(goString(gerr.message))
	}
	return ok, nil
}

// Close is a wrapper around g_output_stream_close().
func (v *OutputStream) Close(cancellable *Cancellable) (bool, error) {
	var gerr *C.GError
	ok := gobool(C.g_output_stream_close(
		v.native(),
		cancellable.native(),
		&gerr))
	if !ok {
		defer C.g_error_free(gerr)
		return false, errors.New(goString(gerr.message))
	}
	return ok, nil
}

// TODO outputStream asynch functions
/*
void 	g_output_stream_write_async ()
gssize 	g_output_stream_write_finish ()
void 	g_output_stream_splice_async ()
gssize 	g_output_stream_splice_finish ()
void 	g_output_stream_flush_async ()
gboolean 	g_output_stream_flush_finish ()
void 	g_output_stream_close_async ()
gboolean 	g_output_stream_close_finish ()
*/

// IsClosing is a wrapper around g_output_stream_is_closing().
func (v *OutputStream) IsClosing() bool {
	return gobool(C.g_output_stream_is_closing(v.native()))
}

// IsClosed is a wrapper around g_output_stream_is_closed().
func (v *OutputStream) IsClosed() bool {
	return gobool(C.g_output_stream_is_closed(v.native()))
}

// HasPending is a wrapper around g_output_stream_has_pending().
func (v *OutputStream) HasPending() bool {
	return gobool(C.g_output_stream_has_pending(v.native()))
}

// SetPending is a wrapper around g_output_stream_set_pending().
func (v *OutputStream) SetPending() (bool, error) {
	var gerr *C.GError
	ok := gobool(C.g_output_stream_set_pending(
		v.native(),
		&gerr))
	if !ok {
		defer C.g_error_free(gerr)
		return false, errors.New(goString(gerr.message))
	}
	return ok, nil
}

// ClearPending is a wrapper around g_output_stream_clear_pending().
func (v *OutputStream) ClearPending() {
	C.g_output_stream_clear_pending(v.native())
}

/*
gssize 	g_output_stream_write_bytes ()
void 	g_output_stream_write_bytes_async ()
gssize 	g_output_stream_write_bytes_finish ()
gboolean 	g_output_stream_printf ()
gboolean 	g_output_stream_vprintf ()
*/

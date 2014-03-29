// Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
//
// This file originated from: http://opensource.conformal.com/
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// Go bindings for GDK 3.  Supports version 3.6 and later.
package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import (
	"errors"
	"github.com/conformal/gotk3/glib"
	"reflect"
	"runtime"
	"unsafe"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.gdk_colorspace_get_type()), marshalColorspace},
		{glib.Type(C.gdk_pixbuf_alpha_mode_get_type()), marshalPixbufAlphaMode},

		// Objects/Interfaces
		{glib.Type(C.gdk_device_get_type()), marshalDevice},
		{glib.Type(C.gdk_device_manager_get_type()), marshalDeviceManager},
		{glib.Type(C.gdk_display_get_type()), marshalDisplay},
		{glib.Type(C.gdk_pixbuf_get_type()), marshalPixbuf},
		{glib.Type(C.gdk_screen_get_type()), marshalScreen},
		{glib.Type(C.gdk_window_get_type()), marshalWindow},

		// Boxed
		{glib.Type(C.gdk_event_get_type()), marshalEvent},
	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * Type conversions
 */

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}
func gobool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

/*
 * Unexported vars
 */

var nilPtrErr = errors.New("cgo returned unexpected nil pointer")

/*
 * Constants
 */

// Colorspace is a representation of GDK's GdkColorspace.
type Colorspace int

const (
	COLORSPACE_RGB Colorspace = C.GDK_COLORSPACE_RGB
)

func marshalColorspace(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Colorspace(c), nil
}

// PixbufAlphaMode is a representation of GDK's GdkPixbufAlphaMode.
type PixbufAlphaMode int

const (
	GDK_PIXBUF_ALPHA_BILEVEL PixbufAlphaMode = C.GDK_PIXBUF_ALPHA_BILEVEL
	GDK_PIXBUF_ALPHA_FULL    PixbufAlphaMode = C.GDK_PIXBUF_ALPHA_FULL
)

func marshalPixbufAlphaMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PixbufAlphaMode(c), nil
}

// Selections
const (
	SELECTION_PRIMARY       Atom = 1
	SELECTION_SECONDARY     Atom = 2
	SELECTION_CLIPBOARD     Atom = 69
	TARGET_BITMAP           Atom = 5
	TARGET_COLORMAP         Atom = 7
	TARGET_DRAWABLE         Atom = 17
	TARGET_PIXMAP           Atom = 20
	TARGET_STRING           Atom = 31
	SELECTION_TYPE_ATOM     Atom = 4
	SELECTION_TYPE_BITMAP   Atom = 5
	SELECTION_TYPE_COLORMAP Atom = 7
	SELECTION_TYPE_DRAWABLE Atom = 17
	SELECTION_TYPE_INTEGER  Atom = 19
	SELECTION_TYPE_PIXMAP   Atom = 20
	SELECTION_TYPE_WINDOW   Atom = 33
	SELECTION_TYPE_STRING   Atom = 31
)

/*
 * GdkAtom
 */

// Atom is a representation of GDK's GdkAtom.
type Atom uintptr

// native returns the underlying GdkAtom.
func (v Atom) native() C.GdkAtom {
	return C.toGdkAtom(unsafe.Pointer(uintptr(v)))
}

/*
 * GdkDevice
 */

// Device is a representation of GDK's GdkDevice.
type Device struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkDevice.
func (v *Device) native() *C.GdkDevice {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkDevice(p)
}

// Native returns a pointer to the underlying GdkDevice.
func (v *Device) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalDevice(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &Device{obj}, nil
}

/*
 * GdkDeviceManager
 */

// DeviceManager is a representation of GDK's GdkDeviceManager.
type DeviceManager struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkDeviceManager.
func (v *DeviceManager) native() *C.GdkDeviceManager {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkDeviceManager(p)
}

// Native returns a pointer to the underlying GdkDeviceManager.
func (v *DeviceManager) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalDeviceManager(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &DeviceManager{obj}, nil
}

/*
 * GdkDisplay
 */

// Display is a representation of GDK's GdkDisplay.
type Display struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkDisplay.
func (v *Display) native() *C.GdkDisplay {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkDisplay(p)
}

// Native returns a pointer to the underlying GdkDisplay.
func (v *Display) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalDisplay(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &Display{obj}, nil
}

// DisplayOpen() is a wrapper around gdk_display_open().
func DisplayOpen(displayName string) (*Display, error) {
	cstr := C.CString(displayName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_display_open((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	d := &Display{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return d, nil
}

// DisplayGetDefault() is a wrapper around gdk_display_get_default().
func DisplayGetDefault() (*Display, error) {
	c := C.gdk_display_get_default()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	d := &Display{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return d, nil
}

// GetName() is a wrapper around gdk_display_get_name().
func (v *Display) GetName() (string, error) {
	c := C.gdk_display_get_name(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// GetScreen() is a wrapper around gdk_display_get_screen().
func (v *Display) GetScreen(screenNum int) (*Screen, error) {
	c := C.gdk_display_get_screen(v.native(), C.gint(screenNum))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := &Screen{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

// GetDefaultScreen() is a wrapper around gdk_display_get_default_screen().
func (v *Display) GetDefaultScreen() (*Screen, error) {
	c := C.gdk_display_get_default_screen(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := &Screen{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

// GetDeviceManager() is a wrapper around gdk_display_get_device_manager().
func (v *Display) GetDeviceManager() (*DeviceManager, error) {
	c := C.gdk_display_get_device_manager(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	d := &DeviceManager{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return d, nil
}

// DeviceIsGrabbed() is a wrapper around gdk_display_device_is_grabbed().
func (v *Display) DeviceIsGrabbed(device *Device) bool {
	c := C.gdk_display_device_is_grabbed(v.native(), device.native())
	return gobool(c)
}

// Beep() is a wrapper around gdk_display_beep().
func (v *Display) Beep() {
	C.gdk_display_beep(v.native())
}

// Sync() is a wrapper around gdk_display_sync().
func (v *Display) Sync() {
	C.gdk_display_sync(v.native())
}

// Flush() is a wrapper around gdk_display_flush().
func (v *Display) Flush() {
	C.gdk_display_flush(v.native())
}

// Close() is a wrapper around gdk_display_close().
func (v *Display) Close() {
	C.gdk_display_close(v.native())
}

// IsClosed() is a wrapper around gdk_display_is_closed().
func (v *Display) IsClosed() bool {
	c := C.gdk_display_is_closed(v.native())
	return gobool(c)
}

// GetEvent() is a wrapper around gdk_display_get_event().
func (v *Display) GetEvent() (*Event, error) {
	c := C.gdk_display_get_event(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	e := &Event{c}
	runtime.SetFinalizer(e, (*Event).free)
	return e, nil
}

// PeekEvent() is a wrapper around gdk_display_peek_event().
func (v *Display) PeekEvent() (*Event, error) {
	c := C.gdk_display_peek_event(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	e := &Event{c}
	runtime.SetFinalizer(e, (*Event).free)
	return e, nil
}

// PutEvent() is a wrapper around gdk_display_put_event().
func (v *Display) PutEvent(event *Event) {
	C.gdk_display_put_event(v.native(), event.native())
}

// HasPending() is a wrapper around gdk_display_has_pending().
func (v *Display) HasPending() bool {
	c := C.gdk_display_has_pending(v.native())
	return gobool(c)
}

// SetDoubleClickTime() is a wrapper around gdk_display_set_double_click_time().
func (v *Display) SetDoubleClickTime(msec uint) {
	C.gdk_display_set_double_click_time(v.native(), C.guint(msec))
}

// SetDoubleClickDistance() is a wrapper around gdk_display_set_double_click_distance().
func (v *Display) SetDoubleClickDistance(distance uint) {
	C.gdk_display_set_double_click_distance(v.native(), C.guint(distance))
}

// SupportsColorCursor() is a wrapper around gdk_display_supports_cursor_color().
func (v *Display) SupportsColorCursor() bool {
	c := C.gdk_display_supports_cursor_color(v.native())
	return gobool(c)
}

// SupportsCursorAlpha() is a wrapper around gdk_display_supports_cursor_alpha().
func (v *Display) SupportsCursorAlpha() bool {
	c := C.gdk_display_supports_cursor_alpha(v.native())
	return gobool(c)
}

// GetDefaultCursorSize() is a wrapper around gdk_display_get_default_cursor_size().
func (v *Display) GetDefaultCursorSize() uint {
	c := C.gdk_display_get_default_cursor_size(v.native())
	return uint(c)
}

// GetMaximalCursorSize() is a wrapper around gdk_display_get_maximal_cursor_size().
func (v *Display) GetMaximalCursorSize() (width, height uint) {
	var w, h C.guint
	C.gdk_display_get_maximal_cursor_size(v.native(), &w, &h)
	return uint(w), uint(h)
}

// GetDefaultGroup() is a wrapper around gdk_display_get_default_group().
func (v *Display) GetDefaultGroup() (*Window, error) {
	c := C.gdk_display_get_default_group(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := &Window{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// SupportsSelectionNotification() is a wrapper around
// gdk_display_supports_selection_notification().
func (v *Display) SupportsSelectionNotification() bool {
	c := C.gdk_display_supports_selection_notification(v.native())
	return gobool(c)
}

// RequestSelectionNotification() is a wrapper around
// gdk_display_request_selection_notification().
func (v *Display) RequestSelectionNotification(selection Atom) bool {
	c := C.gdk_display_request_selection_notification(v.native(),
		selection.native())
	return gobool(c)
}

// SupportsClipboardPersistence() is a wrapper around
// gdk_display_supports_clipboard_persistence().
func (v *Display) SupportsClipboardPersistence() bool {
	c := C.gdk_display_supports_clipboard_persistence(v.native())
	return gobool(c)
}

// TODO(jrick)
func (v *Display) StoreClipboard(clipboardWindow *Window, time uint32, targets ...Atom) {
}

// SupportsShapes() is a wrapper around gdk_display_supports_shapes().
func (v *Display) SupportsShapes() bool {
	c := C.gdk_display_supports_shapes(v.native())
	return gobool(c)
}

// SupportsInputShapes() is a wrapper around gdk_display_supports_input_shapes().
func (v *Display) SupportsInputShapes() bool {
	c := C.gdk_display_supports_input_shapes(v.native())
	return gobool(c)
}

// SupportsComposite() is a wrapper around gdk_display_supports_composite().
func (v *Display) SupportsComposite() bool {
	c := C.gdk_display_supports_composite(v.native())
	return gobool(c)
}

// TODO(jrick) glib.AppLaunchContext GdkAppLaunchContext
func (v *Display) GetAppLaunchContext() {
}

// NotifyStartupComplete() is a wrapper around gdk_display_notify_startup_complete().
func (v *Display) NotifyStartupComplete(startupID string) {
	cstr := C.CString(startupID)
	defer C.free(unsafe.Pointer(cstr))
	C.gdk_display_notify_startup_complete(v.native(), (*C.gchar)(cstr))
}

/*
 * GdkEvent
 */

// Event is a representation of GDK's GdkEvent.
type Event struct {
	GdkEvent *C.GdkEvent
}

// native returns a pointer to the underlying GdkEvent.
func (v *Event) native() *C.GdkEvent {
	if v == nil {
		return nil
	}
	return v.GdkEvent
}

// Native returns a pointer to the underlying GdkEvent.
func (v *Event) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalEvent(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &Event{(*C.GdkEvent)(unsafe.Pointer(c))}, nil
}

func (v *Event) free() {
	C.gdk_event_free(v.native())
}

/*
 * GdkPixbuf
 */

// Pixbuf is a representation of GDK's GdkPixbuf.
type Pixbuf struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkPixbuf.
func (v *Pixbuf) native() *C.GdkPixbuf {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkPixbuf(p)
}

// Native returns a pointer to the underlying GdkPixbuf.
func (v *Pixbuf) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalPixbuf(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &Pixbuf{obj}, nil
}

// GetColorspace is a wrapper around gdk_pixbuf_get_colorspace().
func (v *Pixbuf) GetColorspace() Colorspace {
	c := C.gdk_pixbuf_get_colorspace(v.native())
	return Colorspace(c)
}

// GetNChannels is a wrapper around gdk_pixbuf_get_n_channels().
func (v *Pixbuf) GetNChannels() int {
	c := C.gdk_pixbuf_get_n_channels(v.native())
	return int(c)
}

// GetHasAlpha is a wrapper around gdk_pixbuf_get_has_alpha().
func (v *Pixbuf) GetHasAlpha() bool {
	c := C.gdk_pixbuf_get_has_alpha(v.native())
	return gobool(c)
}

// GetBitsPerSample is a wrapper around gdk_pixbuf_get_bits_per_sample().
func (v *Pixbuf) GetBitsPerSample() int {
	c := C.gdk_pixbuf_get_bits_per_sample(v.native())
	return int(c)
}

// GetPixels is a wrapper around gdk_pixbuf_get_pixels_with_length().
// A Go slice is used to represent the underlying Pixbuf data array, one
// byte per channel.
func (v *Pixbuf) GetPixels() (channels []byte) {
	var length C.guint
	c := C.gdk_pixbuf_get_pixels_with_length(v.native(), &length)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&channels))
	sliceHeader.Data = uintptr(unsafe.Pointer(c))
	sliceHeader.Len = int(length)
	sliceHeader.Cap = int(length)
	// To make sure the slice doesn't outlive the Pixbuf, add a reference
	v.Ref()
	runtime.SetFinalizer(&channels, func(_ *[]byte) {
		v.Unref()
	})
	return
}

// GetWidth is a wrapper around gdk_pixbuf_get_width().
func (v *Pixbuf) GetWidth() int {
	c := C.gdk_pixbuf_get_width(v.native())
	return int(c)
}

// GetHeight is a wrapper around gdk_pixbuf_get_height().
func (v *Pixbuf) GetHeight() int {
	c := C.gdk_pixbuf_get_height(v.native())
	return int(c)
}

// GetRowstride is a wrapper around gdk_pixbuf_get_rowstride().
func (v *Pixbuf) GetRowstride() int {
	c := C.gdk_pixbuf_get_rowstride(v.native())
	return int(c)
}

// GetByteLength is a wrapper around gdk_pixbuf_get_byte_length().
func (v *Pixbuf) GetByteLength() int {
	c := C.gdk_pixbuf_get_byte_length(v.native())
	return int(c)
}

// GetOption is a wrapper around gdk_pixbuf_get_option().  ok is true if
// the key has an associated value.
func (v *Pixbuf) GetOption(key string) (value string, ok bool) {
	cstr := C.CString(key)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_pixbuf_get_option(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return "", false
	}
	return C.GoString((*C.char)(c)), true
}

// PixbufNew is a wrapper around gdk_pixbuf_new().
func PixbufNew(colorspace Colorspace, hasAlpha bool, bitsPerSample, width, height int) (*Pixbuf, error) {
	c := C.gdk_pixbuf_new(C.GdkColorspace(colorspace), gbool(hasAlpha),
		C.int(bitsPerSample), C.int(width), C.int(height))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

/*
 * GdkScreen
 */

// Screen is a representation of GDK's GdkScreen.
type Screen struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkScreen.
func (v *Screen) native() *C.GdkScreen {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkScreen(p)
}

// Native returns a pointer to the underlying GdkScreen.
func (v *Screen) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalScreen(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &Screen{obj}, nil
}

/*
 * GdkWindow
 */

// Window is a representation of GDK's GdkWindow.
type Window struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkWindow.
func (v *Window) native() *C.GdkWindow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkWindow(p)
}

// Native returns a pointer to the underlying GdkWindow.
func (v *Window) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalWindow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &Window{obj}, nil
}

/*
 * Copyright (c) 2013 Conformal Systems <info@conformal.com>
 *
 * This file originated from: http://opensource.conformal.com/
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

/*
Go bindings for GDK 3.  Supports version 3.8 and later.
*/
package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import (
	"errors"
	"github.com/conformal/gotk3/glib"
	"runtime"
	"unsafe"
)

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

// Selections
const (
	SELECTION_PRIMARY       = Atom(uintptr(1))
	SELECTION_SECONDARY     = Atom(uintptr(2))
	SELECTION_CLIPBOARD     = Atom(uintptr(69))
	TARGET_BITMAP           = Atom(uintptr(5))
	TARGET_COLORMAP         = Atom(uintptr(7))
	TARGET_DRAWABLE         = Atom(uintptr(17))
	TARGET_PIXMAP           = Atom(uintptr(20))
	TARGET_STRING           = Atom(uintptr(31))
	SELECTION_TYPE_ATOM     = Atom(uintptr(4))
	SELECTION_TYPE_BITMAP   = Atom(uintptr(5))
	SELECTION_TYPE_COLORMAP = Atom(uintptr(7))
	SELECTION_TYPE_DRAWABLE = Atom(uintptr(17))
	SELECTION_TYPE_INTEGER  = Atom(uintptr(19))
	SELECTION_TYPE_PIXMAP   = Atom(uintptr(20))
	SELECTION_TYPE_WINDOW   = Atom(uintptr(33))
	SELECTION_TYPE_STRING   = Atom(uintptr(31))
)

/*
 * GdkAtom
 */

type Atom uintptr

func (v Atom) Native() C.GdkAtom {
	return C.toGdkAtom(unsafe.Pointer(v))
}

/*
 * GdkDevice
 */

type Device struct {
	*glib.Object
}

func (v *Device) Native() *C.GdkDevice {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkDevice(p)
}

/*
 * GdkDeviceManager
 */

type DeviceManager struct {
	*glib.Object
}

func (v *DeviceManager) Native() *C.GdkDeviceManager {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkDeviceManager(p)
}

/*
 * GdkDisplay
 */

type Display struct {
	*glib.Object
}

func (v *Display) Native() *C.GdkDisplay {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkDisplay(p)
}

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

func (v *Display) GetName() (string, error) {
	c := C.gdk_display_get_name(v.Native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

func (v *Display) GetNScreens() int {
	c := C.gdk_display_get_n_screens(v.Native())
	return int(c)
}

func (v *Display) GetScreen(screenNum int) (*Screen, error) {
	c := C.gdk_display_get_screen(v.Native(), C.gint(screenNum))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := &Screen{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

func (v *Display) GetDefaultScreen() (*Screen, error) {
	c := C.gdk_display_get_default_screen(v.Native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := &Screen{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

func (v *Display) GetDeviceManager() (*DeviceManager, error) {
	c := C.gdk_display_get_device_manager(v.Native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	d := &DeviceManager{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return d, nil
}

func (v *Display) DeviceIsGrabbed(device *Device) bool {
	c := C.gdk_display_device_is_grabbed(v.Native(), device.Native())
	return gobool(c)
}

func (v *Display) Beep() {
	C.gdk_display_beep(v.Native())
}

func (v *Display) Sync() {
	C.gdk_display_sync(v.Native())
}

func (v *Display) Flush() {
	C.gdk_display_flush(v.Native())
}

func (v *Display) Close() {
	C.gdk_display_close(v.Native())
}

func (v *Display) IsClosed() bool {
	c := C.gdk_display_is_closed(v.Native())
	return gobool(c)
}

func (v *Display) GetEvent() (*Event, error) {
	c := C.gdk_display_get_event(v.Native())
	if c == nil {
		return nil, nilPtrErr
	}
	e := &Event{c}
	runtime.SetFinalizer(e, (*Event).free)
	return e, nil
}

func (v *Display) PeekEvent() (*Event, error) {
	c := C.gdk_display_peek_event(v.Native())
	if c == nil {
		return nil, nilPtrErr
	}
	e := &Event{c}
	runtime.SetFinalizer(e, (*Event).free)
	return e, nil
}

func (v *Display) PutEvent(event *Event) {
	C.gdk_display_put_event(v.Native(), event.Native())
}

func (v *Display) HasPending() bool {
	c := C.gdk_display_has_pending(v.Native())
	return gobool(c)
}

func (v *Display) SetDoubleClickTime(msec uint) {
	C.gdk_display_set_double_click_time(v.Native(), C.guint(msec))
}

func (v *Display) SetDoubleClickDistance(distance uint) {
	C.gdk_display_set_double_click_distance(v.Native(), C.guint(distance))
}

func (v *Display) SupportsColorCursor() bool {
	c := C.gdk_display_supports_cursor_color(v.Native())
	return gobool(c)
}

func (v *Display) SupportsCursorAlpha() bool {
	c := C.gdk_display_supports_cursor_alpha(v.Native())
	return gobool(c)
}

func (v *Display) GetDefaultCursorSize() uint {
	c := C.gdk_display_get_default_cursor_size(v.Native())
	return uint(c)
}

func (v *Display) GetMaximalCursorSize() (width, height uint) {
	var w, h C.guint
	C.gdk_display_get_maximal_cursor_size(v.Native(), &w, &h)
	return uint(w), uint(h)
}

func (v *Display) GetDefaultGroup() (*Window, error) {
	c := C.gdk_display_get_default_group(v.Native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := &Window{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

func (v *Display) SupportsSelectionNotification() bool {
	c := C.gdk_display_supports_selection_notification(v.Native())
	return gobool(c)
}

func (v *Display) RequestSelectionNotification(selection Atom) bool {
	c := C.gdk_display_request_selection_notification(v.Native(),
		selection.Native())
	return gobool(c)
}

func (v *Display) SupportsClipboardPersistence() bool {
	c := C.gdk_display_supports_clipboard_persistence(v.Native())
	return gobool(c)
}

// TODO(jrick)
func (v *Display) StoreClipboard(clipboardWindow *Window, time uint32, targets ...Atom) {
}

func (v *Display) SupportsShapes() bool {
	c := C.gdk_display_supports_shapes(v.Native())
	return gobool(c)
}

func (v *Display) SupportsInputShapes() bool {
	c := C.gdk_display_supports_input_shapes(v.Native())
	return gobool(c)
}

func (v *Display) SupportsComposite() bool {
	c := C.gdk_display_supports_composite(v.Native())
	return gobool(c)
}

// TODO(jrick) glib.AppLaunchContext GdkAppLaunchContext
func (v *Display) GetAppLaunchContext() {
}

func (v *Display) NotifyStartupComplete(startupID string) {
	cstr := C.CString(startupID)
	defer C.free(unsafe.Pointer(cstr))
	C.gdk_display_notify_startup_complete(v.Native(), (*C.gchar)(cstr))
}

/*
 * GdkEvent
 */

type Event struct {
	GdkEvent *C.GdkEvent
}

func (v *Event) Native() *C.GdkEvent {
	if v == nil {
		return nil
	}
	return v.GdkEvent
}

func (v *Event) free() {
	C.gdk_event_free(v.Native())
}

/*
 * GdkScreen
 */

type Screen struct {
	*glib.Object
}

func (v *Screen) Native() *C.GdkScreen {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkScreen(p)
}

/*
 * GdkWindow
 */

type Window struct {
	*glib.Object
}

func (v *Window) Native() *C.GdkWindow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkWindow(p)
}

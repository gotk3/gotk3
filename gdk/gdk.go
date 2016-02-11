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
	"reflect"
	"runtime"
	"strconv"
	"unsafe"

	"github.com/gotk3/gotk3/gdk/iface"
	"github.com/gotk3/gotk3/glib"
	glib_iface "github.com/gotk3/gotk3/glib/iface"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib_iface.Type(C.gdk_drag_action_get_type()), marshalDragAction},
		{glib_iface.Type(C.gdk_colorspace_get_type()), marshalColorspace},
		{glib_iface.Type(C.gdk_event_type_get_type()), marshalEventType},
		{glib_iface.Type(C.gdk_interp_type_get_type()), marshalInterpType},
		{glib_iface.Type(C.gdk_modifier_type_get_type()), marshalModifierType},
		{glib_iface.Type(C.gdk_pixbuf_alpha_mode_get_type()), marshalPixbufAlphaMode},
		{glib_iface.Type(C.gdk_event_mask_get_type()), marshalEventMask},

		// Objects/Interfaces
		{glib_iface.Type(C.gdk_device_get_type()), marshalDevice},
		{glib_iface.Type(C.gdk_cursor_get_type()), marshalCursor},
		{glib_iface.Type(C.gdk_device_manager_get_type()), marshalDeviceManager},
		{glib_iface.Type(C.gdk_display_get_type()), marshalDisplay},
		{glib_iface.Type(C.gdk_drag_context_get_type()), marshalDragContext},
		{glib_iface.Type(C.gdk_pixbuf_get_type()), marshalPixbuf},
		{glib_iface.Type(C.gdk_rgba_get_type()), marshalRGBA},
		{glib_iface.Type(C.gdk_screen_get_type()), marshalScreen},
		{glib_iface.Type(C.gdk_visual_get_type()), marshalVisual},
		{glib_iface.Type(C.gdk_window_get_type()), marshalWindow},

		// Boxed
		{glib_iface.Type(C.gdk_event_get_type()), marshalEvent},
	}
	glib.RegisterGValueMarshalers(tm)

	iface.ACTION_DEFAULT = C.GDK_ACTION_DEFAULT
	iface.ACTION_COPY = C.GDK_ACTION_COPY
	iface.ACTION_MOVE = C.GDK_ACTION_MOVE
	iface.ACTION_LINK = C.GDK_ACTION_LINK
	iface.ACTION_PRIVATE = C.GDK_ACTION_PRIVATE
	iface.ACTION_ASK = C.GDK_ACTION_ASK

	iface.COLORSPACE_RGB = C.GDK_COLORSPACE_RGB

	iface.INTERP_NEAREST = C.GDK_INTERP_NEAREST
	iface.INTERP_TILES = C.GDK_INTERP_TILES
	iface.INTERP_BILINEAR = C.GDK_INTERP_BILINEAR
	iface.INTERP_HYPER = C.GDK_INTERP_HYPER

	iface.PIXBUF_ROTATE_NONE = C.GDK_PIXBUF_ROTATE_NONE
	iface.PIXBUF_ROTATE_COUNTERCLOCKWISE = C.GDK_PIXBUF_ROTATE_COUNTERCLOCKWISE
	iface.PIXBUF_ROTATE_UPSIDEDOWN = C.GDK_PIXBUF_ROTATE_UPSIDEDOWN
	iface.PIXBUF_ROTATE_CLOCKWISE = C.GDK_PIXBUF_ROTATE_CLOCKWISE

	iface.GDK_SHIFT_MASK = C.GDK_SHIFT_MASK
	iface.GDK_LOCK_MASK = C.GDK_LOCK_MASK
	iface.GDK_CONTROL_MASK = C.GDK_CONTROL_MASK
	iface.GDK_MOD1_MASK = C.GDK_MOD1_MASK
	iface.GDK_MOD2_MASK = C.GDK_MOD2_MASK
	iface.GDK_MOD3_MASK = C.GDK_MOD3_MASK
	iface.GDK_MOD4_MASK = C.GDK_MOD4_MASK
	iface.GDK_MOD5_MASK = C.GDK_MOD5_MASK
	iface.GDK_BUTTON1_MASK = C.GDK_BUTTON1_MASK
	iface.GDK_BUTTON2_MASK = C.GDK_BUTTON2_MASK
	iface.GDK_BUTTON3_MASK = C.GDK_BUTTON3_MASK
	iface.GDK_BUTTON4_MASK = C.GDK_BUTTON4_MASK
	iface.GDK_BUTTON5_MASK = C.GDK_BUTTON5_MASK
	iface.GDK_SUPER_MASK = C.GDK_SUPER_MASK
	iface.GDK_HYPER_MASK = C.GDK_HYPER_MASK
	iface.GDK_META_MASK = C.GDK_META_MASK
	iface.GDK_RELEASE_MASK = C.GDK_RELEASE_MASK
	iface.GDK_MODIFIER_MASK = C.GDK_MODIFIER_MASK

	iface.GDK_PIXBUF_ALPHA_BILEVEL = C.GDK_PIXBUF_ALPHA_BILEVEL
	iface.GDK_PIXBUF_ALPHA_FULL = C.GDK_PIXBUF_ALPHA_FULL

	iface.SELECTION_PRIMARY = 1
	iface.SELECTION_SECONDARY = 2
	iface.SELECTION_CLIPBOARD = 69
	iface.TARGET_BITMAP = 5
	iface.TARGET_COLORMAP = 7
	iface.TARGET_DRAWABLE = 17
	iface.TARGET_PIXMAP = 20
	iface.TARGET_STRING = 31
	iface.SELECTION_TYPE_ATOM = 4
	iface.SELECTION_TYPE_BITMAP = 5
	iface.SELECTION_TYPE_COLORMAP = 7
	iface.SELECTION_TYPE_DRAWABLE = 17
	iface.SELECTION_TYPE_INTEGER = 19
	iface.SELECTION_TYPE_PIXMAP = 20
	iface.SELECTION_TYPE_WINDOW = 33
	iface.SELECTION_TYPE_STRING = 31

	iface.EXPOSURE_MASK = C.GDK_EXPOSURE_MASK
	iface.POINTER_MOTION_MASK = C.GDK_POINTER_MOTION_MASK
	iface.POINTER_MOTION_HINT_MASK = C.GDK_POINTER_MOTION_HINT_MASK
	iface.BUTTON_MOTION_MASK = C.GDK_BUTTON_MOTION_MASK
	iface.BUTTON1_MOTION_MASK = C.GDK_BUTTON1_MOTION_MASK
	iface.BUTTON2_MOTION_MASK = C.GDK_BUTTON2_MOTION_MASK
	iface.BUTTON3_MOTION_MASK = C.GDK_BUTTON3_MOTION_MASK
	iface.BUTTON_PRESS_MASK = C.GDK_BUTTON_PRESS_MASK
	iface.BUTTON_RELEASE_MASK = C.GDK_BUTTON_RELEASE_MASK
	iface.KEY_PRESS_MASK = C.GDK_KEY_PRESS_MASK
	iface.KEY_RELEASE_MASK = C.GDK_KEY_RELEASE_MASK
	iface.ENTER_NOTIFY_MASK = C.GDK_ENTER_NOTIFY_MASK
	iface.LEAVE_NOTIFY_MASK = C.GDK_LEAVE_NOTIFY_MASK
	iface.FOCUS_CHANGE_MASK = C.GDK_FOCUS_CHANGE_MASK
	iface.STRUCTURE_MASK = C.GDK_STRUCTURE_MASK
	iface.PROPERTY_CHANGE_MASK = C.GDK_PROPERTY_CHANGE_MASK
	iface.VISIBILITY_NOTIFY_MASK = C.GDK_VISIBILITY_NOTIFY_MASK
	iface.PROXIMITY_IN_MASK = C.GDK_PROXIMITY_IN_MASK
	iface.PROXIMITY_OUT_MASK = C.GDK_PROXIMITY_OUT_MASK
	iface.SUBSTRUCTURE_MASK = C.GDK_SUBSTRUCTURE_MASK
	iface.SCROLL_MASK = C.GDK_SCROLL_MASK
	iface.TOUCH_MASK = C.GDK_TOUCH_MASK
	iface.SMOOTH_SCROLL_MASK = C.GDK_SMOOTH_SCROLL_MASK
	iface.ALL_EVENTS_MASK = C.GDK_ALL_EVENTS_MASK

	iface.SCROLL_UP = C.GDK_SCROLL_UP
	iface.SCROLL_DOWN = C.GDK_SCROLL_DOWN
	iface.SCROLL_LEFT = C.GDK_SCROLL_LEFT
	iface.SCROLL_RIGHT = C.GDK_SCROLL_RIGHT
	iface.SCROLL_SMOOTH = C.GDK_SCROLL_SMOOTH

	iface.GRAB_SUCCESS = C.GDK_GRAB_SUCCESS
	iface.GRAB_ALREADY_GRABBED = C.GDK_GRAB_ALREADY_GRABBED
	iface.GRAB_INVALID_TIME = C.GDK_GRAB_INVALID_TIME
	iface.GRAB_FROZEN = C.GDK_GRAB_FROZEN
	// Only exists since 3.16
	// GRAB_FAILED GrabStatus = C.GDK_GRAB_FAILED
	iface.GRAB_FAILED = 5

	iface.OWNERSHIP_NONE = C.GDK_OWNERSHIP_NONE
	iface.OWNERSHIP_WINDOW = C.GDK_OWNERSHIP_WINDOW
	iface.OWNERSHIP_APPLICATION = C.GDK_OWNERSHIP_APPLICATION

	iface.DEVICE_TYPE_MASTER = C.GDK_DEVICE_TYPE_MASTER
	iface.DEVICE_TYPE_SLAVE = C.GDK_DEVICE_TYPE_SLAVE
	iface.DEVICE_TYPE_FLOATING = C.GDK_DEVICE_TYPE_FLOATING

	iface.EVENT_NOTHING = C.GDK_NOTHING
	iface.EVENT_DELETE = C.GDK_DELETE
	iface.EVENT_DESTROY = C.GDK_DESTROY
	iface.EVENT_EXPOSE = C.GDK_EXPOSE
	iface.EVENT_MOTION_NOTIFY = C.GDK_MOTION_NOTIFY
	iface.EVENT_BUTTON_PRESS = C.GDK_BUTTON_PRESS
	iface.EVENT_2BUTTON_PRESS = C.GDK_2BUTTON_PRESS
	iface.EVENT_DOUBLE_BUTTON_PRESS = C.GDK_DOUBLE_BUTTON_PRESS
	iface.EVENT_3BUTTON_PRESS = C.GDK_3BUTTON_PRESS
	iface.EVENT_TRIPLE_BUTTON_PRESS = C.GDK_TRIPLE_BUTTON_PRESS
	iface.EVENT_BUTTON_RELEASE = C.GDK_BUTTON_RELEASE
	iface.EVENT_KEY_PRESS = C.GDK_KEY_PRESS
	iface.EVENT_KEY_RELEASE = C.GDK_KEY_RELEASE
	iface.EVENT_LEAVE_NOTIFY = C.GDK_ENTER_NOTIFY
	iface.EVENT_FOCUS_CHANGE = C.GDK_FOCUS_CHANGE
	iface.EVENT_CONFIGURE = C.GDK_CONFIGURE
	iface.EVENT_MAP = C.GDK_MAP
	iface.EVENT_UNMAP = C.GDK_UNMAP
	iface.EVENT_PROPERTY_NOTIFY = C.GDK_PROPERTY_NOTIFY
	iface.EVENT_SELECTION_CLEAR = C.GDK_SELECTION_CLEAR
	iface.EVENT_SELECTION_REQUEST = C.GDK_SELECTION_REQUEST
	iface.EVENT_SELECTION_NOTIFY = C.GDK_SELECTION_NOTIFY
	iface.EVENT_PROXIMITY_IN = C.GDK_PROXIMITY_IN
	iface.EVENT_PROXIMITY_OUT = C.GDK_PROXIMITY_OUT
	iface.EVENT_DRAG_ENTER = C.GDK_DRAG_ENTER
	iface.EVENT_DRAG_LEAVE = C.GDK_DRAG_LEAVE
	iface.EVENT_DRAG_MOTION = C.GDK_DRAG_MOTION
	iface.EVENT_DRAG_STATUS = C.GDK_DRAG_STATUS
	iface.EVENT_DROP_START = C.GDK_DROP_START
	iface.EVENT_DROP_FINISHED = C.GDK_DROP_FINISHED
	iface.EVENT_CLIENT_EVENT = C.GDK_CLIENT_EVENT
	iface.EVENT_VISIBILITY_NOTIFY = C.GDK_VISIBILITY_NOTIFY
	iface.EVENT_SCROLL = C.GDK_SCROLL
	iface.EVENT_WINDOW_STATE = C.GDK_WINDOW_STATE
	iface.EVENT_SETTING = C.GDK_SETTING
	iface.EVENT_OWNER_CHANGE = C.GDK_OWNER_CHANGE
	iface.EVENT_GRAB_BROKEN = C.GDK_GRAB_BROKEN
	iface.EVENT_DAMAGE = C.GDK_DAMAGE
	iface.EVENT_TOUCH_BEGIN = C.GDK_TOUCH_BEGIN
	iface.EVENT_TOUCH_UPDATE = C.GDK_TOUCH_UPDATE
	iface.EVENT_TOUCH_END = C.GDK_TOUCH_END
	iface.EVENT_TOUCH_CANCEL = C.GDK_TOUCH_CANCEL
	iface.EVENT_LAST = C.GDK_EVENT_LAST
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

func marshalDragAction(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return iface.DragAction(c), nil
}

func marshalColorspace(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return iface.Colorspace(c), nil
}

func marshalInterpType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return iface.InterpType(c), nil
}

func marshalModifierType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return iface.ModifierType(c), nil
}

func marshalPixbufAlphaMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return iface.PixbufAlphaMode(c), nil
}

func marshalEventMask(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return iface.EventMask(c), nil
}

/*
 * GdkAtom
 */

// nativeAtom returns the underlying GdkAtom.
func nativeAtom(v iface.Atom) C.GdkAtom {
	return C.toGdkAtom(unsafe.Pointer(uintptr(v)))
}

// AtomName returns the name of the atom
func AtomName(v iface.Atom) string {
	c := C.gdk_atom_name(nativeAtom(v))
	defer C.g_free(C.gpointer(c))
	return C.GoString((*C.char)(c))
}

// GdkAtomIntern is a wrapper around gdk_atom_intern
func GdkAtomIntern(atomName string, onlyIfExists bool) iface.Atom {
	cstr := C.CString(atomName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_atom_intern((*C.gchar)(cstr), gbool(onlyIfExists))
	return iface.Atom(uintptr(unsafe.Pointer(c)))
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

// Grab() is a wrapper around gdk_device_grab().
func (v *Device) Grab(w iface.Window, ownership iface.GrabOwnership, owner_events bool, event_mask iface.EventMask, cursor iface.Cursor, time uint32) iface.GrabStatus {
	ret := C.gdk_device_grab(
		v.native(),
		w.(*Window).native(),
		C.GdkGrabOwnership(ownership),
		gbool(owner_events),
		C.GdkEventMask(event_mask),
		cursor.(*Cursor).native(),
		C.guint32(time),
	)
	return iface.GrabStatus(ret)
}

// Ungrab() is a wrapper around gdk_device_ungrab().
func (v *Device) Ungrab(time uint32) {
	C.gdk_device_ungrab(v.native(), C.guint32(time))
}

/*
 * GdkCursor
 */

// Cursor is a representation of GdkCursor.
type Cursor struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkCursor.
func (v *Cursor) native() *C.GdkCursor {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkCursor(p)
}

// Native returns a pointer to the underlying GdkCursor.
func (v *Cursor) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalCursor(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &Cursor{obj}, nil
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

// GetClientPointer() is a wrapper around gdk_device_manager_get_client_pointer().
func (v *DeviceManager) GetClientPointer() (iface.Device, error) {
	c := C.gdk_device_manager_get_client_pointer(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return &Device{obj}, nil
}

// GetDisplay() is a wrapper around gdk_device_manager_get_display().
func (v *DeviceManager) GetDisplay() (iface.Display, error) {
	c := C.gdk_device_manager_get_display(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return &Display{obj}, nil
}

// ListDevices() is a wrapper around gdk_device_manager_list_devices().
func (v *DeviceManager) ListDevices(tp iface.DeviceType) glib_iface.List {
	clist := C.gdk_device_manager_list_devices(v.native(), C.GdkDeviceType(tp))
	if clist == nil {
		return nil
	}
	glist := glib.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		return &Device{&glib.Object{glib.ToGObject(ptr)}}
	})
	runtime.SetFinalizer(glist, func(glist *glib.List) {
		glist.Free()
	})
	return glist
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

func toDisplay(s *C.GdkDisplay) (*Display, error) {
	if s == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(s))}
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
func (v *Display) GetScreen(screenNum int) (iface.Screen, error) {
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
func (v *Display) GetDefaultScreen() (iface.Screen, error) {
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
func (v *Display) GetDeviceManager() (iface.DeviceManager, error) {
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
func (v *Display) DeviceIsGrabbed(device iface.Device) bool {
	c := C.gdk_display_device_is_grabbed(v.native(), device.(*Device).native())
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
func (v *Display) GetEvent() (iface.Event, error) {
	c := C.gdk_display_get_event(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	e := &Event{c}
	runtime.SetFinalizer(e, (*Event).free)
	return e, nil
}

// PeekEvent() is a wrapper around gdk_display_peek_event().
func (v *Display) PeekEvent() (iface.Event, error) {
	c := C.gdk_display_peek_event(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	e := &Event{c}
	runtime.SetFinalizer(e, (*Event).free)
	return e, nil
}

// PutEvent() is a wrapper around gdk_display_put_event().
func (v *Display) PutEvent(event iface.Event) {
	C.gdk_display_put_event(v.native(), event.(*Event).native())
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
func (v *Display) GetDefaultGroup() (iface.Window, error) {
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
func (v *Display) RequestSelectionNotification(selection iface.Atom) bool {
	c := C.gdk_display_request_selection_notification(v.native(),
		nativeAtom(selection))
	return gobool(c)
}

// SupportsClipboardPersistence() is a wrapper around
// gdk_display_supports_clipboard_persistence().
func (v *Display) SupportsClipboardPersistence() bool {
	c := C.gdk_display_supports_clipboard_persistence(v.native())
	return gobool(c)
}

// TODO(jrick)
func (v *Display) StoreClipboard(clipboardWindow iface.Window, time uint32, targets ...iface.Atom) {
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

// TODO(jrick) glib.AppLaunchContext GdkAppLaunchContext
func (v *Display) GetAppLaunchContext() {
}

// NotifyStartupComplete() is a wrapper around gdk_display_notify_startup_complete().
func (v *Display) NotifyStartupComplete(startupID string) {
	cstr := C.CString(startupID)
	defer C.free(unsafe.Pointer(cstr))
	C.gdk_display_notify_startup_complete(v.native(), (*C.gchar)(cstr))
}

func marshalEventType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return iface.EventType(c), nil
}

/*
 * GDK Keyval
 */

// KeyvalFromName() is a wrapper around gdk_keyval_from_name().
func KeyvalFromName(keyvalName string) uint {
	str := (*C.gchar)(C.CString(keyvalName))
	defer C.free(unsafe.Pointer(str))
	return uint(C.gdk_keyval_from_name(str))
}

func KeyvalConvertCase(v uint) (lower, upper uint) {
	var l, u C.guint
	l = 0
	u = 0
	C.gdk_keyval_convert_case(C.guint(v), &l, &u)
	return uint(l), uint(u)
}

func KeyvalIsLower(v uint) bool {
	return gobool(C.gdk_keyval_is_lower(C.guint(v)))
}

func KeyvalIsUpper(v uint) bool {
	return gobool(C.gdk_keyval_is_upper(C.guint(v)))
}

func KeyvalToLower(v uint) uint {
	return uint(C.gdk_keyval_to_lower(C.guint(v)))
}

func KeyvalToUpper(v uint) uint {
	return uint(C.gdk_keyval_to_upper(C.guint(v)))
}

/*
 * GdkDragContext
 */

// DragContext is a representation of GDK's GdkDragContext.
type DragContext struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkDragContext.
func (v *DragContext) native() *C.GdkDragContext {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkDragContext(p)
}

// Native returns a pointer to the underlying GdkDragContext.
func (v *DragContext) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalDragContext(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &DragContext{obj}, nil
}

func (v *DragContext) ListTargets() glib_iface.List {
	c := C.gdk_drag_context_list_targets(v.native())
	return glib.WrapList(uintptr(unsafe.Pointer(c)))
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
 * GdkEventButton
 */

// EventButton is a representation of GDK's GdkEventButton.
type EventButton struct {
	*Event
}

// Native returns a pointer to the underlying GdkEventButton.
func (v *EventButton) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *EventButton) native() *C.GdkEventButton {
	return (*C.GdkEventButton)(unsafe.Pointer(v.Event.native()))
}

func EventButtonFrom(ev *Event) iface.EventButton {
	return &EventButton{ev}
}

func (v *EventButton) X() float64 {
	c := v.native().x
	return float64(c)
}

func (v *EventButton) Y() float64 {
	c := v.native().y
	return float64(c)
}

// XRoot returns the x coordinate of the pointer relative to the root of the screen.
func (v *EventButton) XRoot() float64 {
	c := v.native().x_root
	return float64(c)
}

// YRoot returns the y coordinate of the pointer relative to the root of the screen.
func (v *EventButton) YRoot() float64 {
	c := v.native().y_root
	return float64(c)
}

func (v *EventButton) Button() uint {
	c := v.native().button
	return uint(c)
}

func (v *EventButton) State() uint {
	c := v.native().state
	return uint(c)
}

// Time returns the time of the event in milliseconds.
func (v *EventButton) Time() uint32 {
	c := v.native().time
	return uint32(c)
}

func (v *EventButton) Type() iface.EventType {
	c := v.native()._type
	return iface.EventType(c)
}

func (v *EventButton) MotionVal() (float64, float64) {
	x := v.native().x
	y := v.native().y
	return float64(x), float64(y)
}

func (v *EventButton) MotionValRoot() (float64, float64) {
	x := v.native().x_root
	y := v.native().y_root
	return float64(x), float64(y)
}

func (v *EventButton) ButtonVal() uint {
	c := v.native().button
	return uint(c)
}

/*
 * GdkEventKey
 */

// EventKey is a representation of GDK's GdkEventKey.
type EventKey struct {
	*Event
}

func EventKeyNew() *EventKey {
	ee := (*C.GdkEvent)(unsafe.Pointer(&C.GdkEventKey{}))
	ev := Event{ee}
	return &EventKey{&ev}
}

// Native returns a pointer to the underlying GdkEventKey.
func (v *EventKey) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *EventKey) native() *C.GdkEventKey {
	return (*C.GdkEventKey)(unsafe.Pointer(v.Event.native()))
}

func (v *EventKey) KeyVal() uint {
	c := v.native().keyval
	return uint(c)
}

func (v *EventKey) Type() iface.EventType {
	c := v.native()._type
	return iface.EventType(c)
}

func (v *EventKey) State() uint {
	c := v.native().state
	return uint(c)
}

/*
 * GdkEventMotion
 */

type EventMotion struct {
	*Event
}

// Native returns a pointer to the underlying GdkEventMotion.
func (v *EventMotion) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *EventMotion) native() *C.GdkEventMotion {
	return (*C.GdkEventMotion)(unsafe.Pointer(v.Event.native()))
}

func (v *EventMotion) MotionVal() (float64, float64) {
	x := v.native().x
	y := v.native().y
	return float64(x), float64(y)
}

func (v *EventMotion) MotionValRoot() (float64, float64) {
	x := v.native().x_root
	y := v.native().y_root
	return float64(x), float64(y)
}

/*
 * GdkEventScroll
 */

// EventScroll is a representation of GDK's GdkEventScroll.
type EventScroll struct {
	*Event
}

// Native returns a pointer to the underlying GdkEventScroll.
func (v *EventScroll) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *EventScroll) native() *C.GdkEventScroll {
	return (*C.GdkEventScroll)(unsafe.Pointer(v.Event.native()))
}

func (v *EventScroll) DeltaX() float64 {
	return float64(v.native().delta_x)
}

func (v *EventScroll) DeltaY() float64 {
	return float64(v.native().delta_y)
}

func (v *EventScroll) X() float64 {
	return float64(v.native().x)
}

func (v *EventScroll) Y() float64 {
	return float64(v.native().y)
}

func (v *EventScroll) Type() iface.EventType {
	c := v.native()._type
	return iface.EventType(c)
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
func (v *Pixbuf) GetColorspace() iface.Colorspace {
	c := C.gdk_pixbuf_get_colorspace(v.native())
	return iface.Colorspace(c)
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
func PixbufNew(colorspace iface.Colorspace, hasAlpha bool, bitsPerSample, width, height int) (*Pixbuf, error) {
	c := C.gdk_pixbuf_new(C.GdkColorspace(colorspace), gbool(hasAlpha),
		C.int(bitsPerSample), C.int(width), C.int(height))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// PixbufCopy is a wrapper around gdk_pixbuf_copy().
func PixbufCopy(v *Pixbuf) (*Pixbuf, error) {
	c := C.gdk_pixbuf_copy(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// PixbufNewFromFile is a wrapper around gdk_pixbuf_new_from_file().
func PixbufNewFromFile(filename string) (*Pixbuf, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError
	res := C.gdk_pixbuf_new_from_file((*C.char)(cstr), &err)
	if res == nil {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(res))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// PixbufNewFromFileAtSize is a wrapper around gdk_pixbuf_new_from_file_at_size().
func PixbufNewFromFileAtSize(filename string, width, height int) (*Pixbuf, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	res := C.gdk_pixbuf_new_from_file_at_size(cstr, C.int(width), C.int(height), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}
	if res == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(res))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// PixbufNewFromFileAtScale is a wrapper around gdk_pixbuf_new_from_file_at_scale().
func PixbufNewFromFileAtScale(filename string, width, height int, preserveAspectRatio bool) (*Pixbuf, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	var err *C.GError = nil
	res := C.gdk_pixbuf_new_from_file_at_scale(cstr, C.int(width), C.int(height),
		gbool(preserveAspectRatio), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}
	if res == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(res))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// ScaleSimple is a wrapper around gdk_pixbuf_scale_simple().
func (v *Pixbuf) ScaleSimple(destWidth, destHeight int, interpType iface.InterpType) (iface.Pixbuf, error) {
	c := C.gdk_pixbuf_scale_simple(v.native(), C.int(destWidth),
		C.int(destHeight), C.GdkInterpType(interpType))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// RotateSimple is a wrapper around gdk_pixbuf_rotate_simple().
func (v *Pixbuf) RotateSimple(angle iface.PixbufRotation) (iface.Pixbuf, error) {
	c := C.gdk_pixbuf_rotate_simple(v.native(), C.GdkPixbufRotation(angle))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// ApplyEmbeddedOrientation is a wrapper around gdk_pixbuf_apply_embedded_orientation().
func (v *Pixbuf) ApplyEmbeddedOrientation() (iface.Pixbuf, error) {
	c := C.gdk_pixbuf_apply_embedded_orientation(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// Flip is a wrapper around gdk_pixbuf_flip().
func (v *Pixbuf) Flip(horizontal bool) (iface.Pixbuf, error) {
	c := C.gdk_pixbuf_flip(v.native(), gbool(horizontal))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// SaveJPEG is a wrapper around gdk_pixbuf_save().
// Quality is a number between 0...100
func (v *Pixbuf) SaveJPEG(path string, quality int) error {
	cpath := C.CString(path)
	cquality := C.CString(strconv.Itoa(quality))
	defer C.free(unsafe.Pointer(cpath))
	defer C.free(unsafe.Pointer(cquality))

	var err *C.GError
	c := C._gdk_pixbuf_save_jpeg(v.native(), cpath, &err, cquality)
	if !gobool(c) {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

// SavePNG is a wrapper around gdk_pixbuf_save().
// Compression is a number between 0...9
func (v *Pixbuf) SavePNG(path string, compression int) error {
	cpath := C.CString(path)
	ccompression := C.CString(strconv.Itoa(compression))
	defer C.free(unsafe.Pointer(cpath))
	defer C.free(unsafe.Pointer(ccompression))

	var err *C.GError
	c := C._gdk_pixbuf_save_png(v.native(), cpath, &err, ccompression)
	if !gobool(c) {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

// PixbufGetFileInfo is a wrapper around gdk_pixbuf_get_file_info().
// TODO: need to wrap the returned format to GdkPixbufFormat.
func PixbufGetFileInfo(filename string) (format interface{}, width, height int) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	var cw, ch C.gint
	format = C.gdk_pixbuf_get_file_info((*C.gchar)(cstr), &cw, &ch)
	// TODO: need to wrap the returned format to GdkPixbufFormat.
	return format, int(cw), int(ch)
}

/*
 * GdkPixbufLoader
 */

// PixbufLoader is a representation of GDK's GdkPixbufLoader.
// Users of PixbufLoader are expected to call Close() when they are finished.
type PixbufLoader struct {
	*glib.Object
}

// native() returns a pointer to the underlying GdkPixbufLoader.
func (v *PixbufLoader) native() *C.GdkPixbufLoader {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkPixbufLoader(p)
}

// PixbufLoaderNew() is a wrapper around gdk_pixbuf_loader_new().
func PixbufLoaderNew() (*PixbufLoader, error) {
	c := C.gdk_pixbuf_loader_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &PixbufLoader{obj}
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

// Write() is a wrapper around gdk_pixbuf_loader_write().  The
// function signature differs from the C equivalent to satisify the
// io.Writer interface.
func (v *PixbufLoader) Write(data []byte) (int, error) {
	// n is set to 0 on error, and set to len(data) otherwise.
	// This is a tiny hacky to satisfy io.Writer and io.WriteCloser,
	// which would allow access to all io and ioutil goodies,
	// and play along nice with go environment.

	if len(data) == 0 {
		return 0, nil
	}

	var err *C.GError
	c := C.gdk_pixbuf_loader_write(v.native(),
		(*C.guchar)(unsafe.Pointer(&data[0])), C.gsize(len(data)),
		&err)
	if !gobool(c) {
		defer C.g_error_free(err)
		return 0, errors.New(C.GoString((*C.char)(err.message)))
	}

	return len(data), nil
}

// Close is a wrapper around gdk_pixbuf_loader_close().  An error is
// returned instead of a bool like the native C function to support the
// io.Closer interface.
func (v *PixbufLoader) Close() error {
	var err *C.GError

	if ok := gobool(C.gdk_pixbuf_loader_close(v.native(), &err)); !ok {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

// SetSize is a wrapper around gdk_pixbuf_loader_set_size().
func (v *PixbufLoader) SetSize(width, height int) {
	C.gdk_pixbuf_loader_set_size(v.native(), C.int(width), C.int(height))
}

// GetPixbuf is a wrapper around gdk_pixbuf_loader_get_pixbuf().
func (v *PixbufLoader) GetPixbuf() (iface.Pixbuf, error) {
	c := C.gdk_pixbuf_loader_get_pixbuf(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return p, nil
}

type RGBA struct {
	rgba *C.GdkRGBA
}

func marshalRGBA(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c2 := (*C.GdkRGBA)(unsafe.Pointer(c))
	return wrapRGBA(c2), nil
}

func wrapRGBA(obj *C.GdkRGBA) *RGBA {
	return &RGBA{obj}
}

func NewRGBA(values ...float64) *RGBA {
	cval := C.GdkRGBA{}
	c := &RGBA{&cval}
	if len(values) > 0 {
		c.rgba.red = C.gdouble(values[0])
	}
	if len(values) > 1 {
		c.rgba.green = C.gdouble(values[1])
	}
	if len(values) > 2 {
		c.rgba.blue = C.gdouble(values[2])
	}
	if len(values) > 3 {
		c.rgba.alpha = C.gdouble(values[3])
	}
	return c
}

func (c *RGBA) Floats() []float64 {
	return []float64{float64(c.rgba.red), float64(c.rgba.green), float64(c.rgba.blue), float64(c.rgba.alpha)}
}

func (v *RGBA) Native() uintptr {
	return uintptr(unsafe.Pointer(v.rgba))
}

// Parse is a representation of gdk_rgba_parse().
func (v *RGBA) Parse(spec string) bool {
	cstr := (*C.gchar)(C.CString(spec))
	defer C.free(unsafe.Pointer(cstr))

	return gobool(C.gdk_rgba_parse(v.rgba, cstr))
}

// String is a representation of gdk_rgba_to_string().
func (v *RGBA) String() string {
	return C.GoString((*C.char)(C.gdk_rgba_to_string(v.rgba)))
}

// GdkRGBA * 	gdk_rgba_copy ()
// void 	gdk_rgba_free ()
// gboolean 	gdk_rgba_equal ()
// guint 	gdk_rgba_hash ()

// PixbufGetType is a wrapper around gdk_pixbuf_get_type().
func PixbufGetType() glib_iface.Type {
	return glib_iface.Type(C.gdk_pixbuf_get_type())
}

/*
 * GdkRectangle
 */

// Rectangle is a representation of GDK's GdkRectangle type.
type Rectangle struct {
	GdkRectangle C.GdkRectangle
}

// Native() returns a pointer to the underlying GdkRectangle.
func (r *Rectangle) native() *C.GdkRectangle {
	return &r.GdkRectangle
}

// GetX returns x field of the underlying GdkRectangle.
func (r *Rectangle) GetX() int {
	return int(r.native().x)
}

// GetY returns y field of the underlying GdkRectangle.
func (r *Rectangle) GetY() int {
	return int(r.native().y)
}

// GetWidth returns width field of the underlying GdkRectangle.
func (r *Rectangle) GetWidth() int {
	return int(r.native().width)
}

// GetHeight returns height field of the underlying GdkRectangle.
func (r *Rectangle) GetHeight() int {
	return int(r.native().height)
}

/*
 * GdkVisual
 */

// Visual is a representation of GDK's GdkVisual.
type Visual struct {
	*glib.Object
}

func (v *Visual) native() *C.GdkVisual {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkVisual(p)
}

func (v *Visual) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalVisual(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &Visual{obj}, nil
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

func toWindow(s *C.GdkWindow) (*Window, error) {
	if s == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(s))}
	return &Window{obj}, nil
}

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
package impl

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

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	glib_impl "github.com/gotk3/gotk3/glib/impl"
)

func init() {
	tm := []glib_impl.TypeMarshaler{
		// Enums
		{glib.Type(C.gdk_drag_action_get_type()), marshalDragAction},
		{glib.Type(C.gdk_colorspace_get_type()), marshalColorspace},
		{glib.Type(C.gdk_event_type_get_type()), marshalEventType},
		{glib.Type(C.gdk_interp_type_get_type()), marshalInterpType},
		{glib.Type(C.gdk_modifier_type_get_type()), marshalModifierType},
		{glib.Type(C.gdk_pixbuf_alpha_mode_get_type()), marshalPixbufAlphaMode},
		{glib.Type(C.gdk_event_mask_get_type()), marshalEventMask},

		// Objects/Interfaces
		{glib.Type(C.gdk_device_get_type()), marshalDevice},
		{glib.Type(C.gdk_cursor_get_type()), marshalCursor},
		{glib.Type(C.gdk_device_manager_get_type()), marshalDeviceManager},
		{glib.Type(C.gdk_display_get_type()), marshalDisplay},
		{glib.Type(C.gdk_drag_context_get_type()), marshalDragContext},
		{glib.Type(C.gdk_pixbuf_get_type()), marshalPixbuf},
		{glib.Type(C.gdk_rgba_get_type()), marshalRGBA},
		{glib.Type(C.gdk_screen_get_type()), marshalScreen},
		{glib.Type(C.gdk_visual_get_type()), marshalVisual},
		{glib.Type(C.gdk_window_get_type()), marshalWindow},

		// Boxed
		{glib.Type(C.gdk_event_get_type()), marshalEvent},
	}
	glib_impl.RegisterGValueMarshalers(tm)

	gdk.ACTION_DEFAULT = C.GDK_ACTION_DEFAULT
	gdk.ACTION_COPY = C.GDK_ACTION_COPY
	gdk.ACTION_MOVE = C.GDK_ACTION_MOVE
	gdk.ACTION_LINK = C.GDK_ACTION_LINK
	gdk.ACTION_PRIVATE = C.GDK_ACTION_PRIVATE
	gdk.ACTION_ASK = C.GDK_ACTION_ASK

	gdk.COLORSPACE_RGB = C.GDK_COLORSPACE_RGB

	gdk.INTERP_NEAREST = C.GDK_INTERP_NEAREST
	gdk.INTERP_TILES = C.GDK_INTERP_TILES
	gdk.INTERP_BILINEAR = C.GDK_INTERP_BILINEAR
	gdk.INTERP_HYPER = C.GDK_INTERP_HYPER

	gdk.PIXBUF_ROTATE_NONE = C.GDK_PIXBUF_ROTATE_NONE
	gdk.PIXBUF_ROTATE_COUNTERCLOCKWISE = C.GDK_PIXBUF_ROTATE_COUNTERCLOCKWISE
	gdk.PIXBUF_ROTATE_UPSIDEDOWN = C.GDK_PIXBUF_ROTATE_UPSIDEDOWN
	gdk.PIXBUF_ROTATE_CLOCKWISE = C.GDK_PIXBUF_ROTATE_CLOCKWISE

	gdk.GDK_SHIFT_MASK = C.GDK_SHIFT_MASK
	gdk.GDK_LOCK_MASK = C.GDK_LOCK_MASK
	gdk.GDK_CONTROL_MASK = C.GDK_CONTROL_MASK
	gdk.GDK_MOD1_MASK = C.GDK_MOD1_MASK
	gdk.GDK_MOD2_MASK = C.GDK_MOD2_MASK
	gdk.GDK_MOD3_MASK = C.GDK_MOD3_MASK
	gdk.GDK_MOD4_MASK = C.GDK_MOD4_MASK
	gdk.GDK_MOD5_MASK = C.GDK_MOD5_MASK
	gdk.GDK_BUTTON1_MASK = C.GDK_BUTTON1_MASK
	gdk.GDK_BUTTON2_MASK = C.GDK_BUTTON2_MASK
	gdk.GDK_BUTTON3_MASK = C.GDK_BUTTON3_MASK
	gdk.GDK_BUTTON4_MASK = C.GDK_BUTTON4_MASK
	gdk.GDK_BUTTON5_MASK = C.GDK_BUTTON5_MASK
	gdk.GDK_SUPER_MASK = C.GDK_SUPER_MASK
	gdk.GDK_HYPER_MASK = C.GDK_HYPER_MASK
	gdk.GDK_META_MASK = C.GDK_META_MASK
	gdk.GDK_RELEASE_MASK = C.GDK_RELEASE_MASK
	gdk.GDK_MODIFIER_MASK = C.GDK_MODIFIER_MASK

	gdk.GDK_PIXBUF_ALPHA_BILEVEL = C.GDK_PIXBUF_ALPHA_BILEVEL
	gdk.GDK_PIXBUF_ALPHA_FULL = C.GDK_PIXBUF_ALPHA_FULL

	gdk.SELECTION_PRIMARY = 1
	gdk.SELECTION_SECONDARY = 2
	gdk.SELECTION_CLIPBOARD = 69
	gdk.TARGET_BITMAP = 5
	gdk.TARGET_COLORMAP = 7
	gdk.TARGET_DRAWABLE = 17
	gdk.TARGET_PIXMAP = 20
	gdk.TARGET_STRING = 31
	gdk.SELECTION_TYPE_ATOM = 4
	gdk.SELECTION_TYPE_BITMAP = 5
	gdk.SELECTION_TYPE_COLORMAP = 7
	gdk.SELECTION_TYPE_DRAWABLE = 17
	gdk.SELECTION_TYPE_INTEGER = 19
	gdk.SELECTION_TYPE_PIXMAP = 20
	gdk.SELECTION_TYPE_WINDOW = 33
	gdk.SELECTION_TYPE_STRING = 31

	gdk.EXPOSURE_MASK = C.GDK_EXPOSURE_MASK
	gdk.POINTER_MOTION_MASK = C.GDK_POINTER_MOTION_MASK
	gdk.POINTER_MOTION_HINT_MASK = C.GDK_POINTER_MOTION_HINT_MASK
	gdk.BUTTON_MOTION_MASK = C.GDK_BUTTON_MOTION_MASK
	gdk.BUTTON1_MOTION_MASK = C.GDK_BUTTON1_MOTION_MASK
	gdk.BUTTON2_MOTION_MASK = C.GDK_BUTTON2_MOTION_MASK
	gdk.BUTTON3_MOTION_MASK = C.GDK_BUTTON3_MOTION_MASK
	gdk.BUTTON_PRESS_MASK = C.GDK_BUTTON_PRESS_MASK
	gdk.BUTTON_RELEASE_MASK = C.GDK_BUTTON_RELEASE_MASK
	gdk.KEY_PRESS_MASK = C.GDK_KEY_PRESS_MASK
	gdk.KEY_RELEASE_MASK = C.GDK_KEY_RELEASE_MASK
	gdk.ENTER_NOTIFY_MASK = C.GDK_ENTER_NOTIFY_MASK
	gdk.LEAVE_NOTIFY_MASK = C.GDK_LEAVE_NOTIFY_MASK
	gdk.FOCUS_CHANGE_MASK = C.GDK_FOCUS_CHANGE_MASK
	gdk.STRUCTURE_MASK = C.GDK_STRUCTURE_MASK
	gdk.PROPERTY_CHANGE_MASK = C.GDK_PROPERTY_CHANGE_MASK
	gdk.VISIBILITY_NOTIFY_MASK = C.GDK_VISIBILITY_NOTIFY_MASK
	gdk.PROXIMITY_IN_MASK = C.GDK_PROXIMITY_IN_MASK
	gdk.PROXIMITY_OUT_MASK = C.GDK_PROXIMITY_OUT_MASK
	gdk.SUBSTRUCTURE_MASK = C.GDK_SUBSTRUCTURE_MASK
	gdk.SCROLL_MASK = C.GDK_SCROLL_MASK
	gdk.TOUCH_MASK = C.GDK_TOUCH_MASK
	gdk.SMOOTH_SCROLL_MASK = C.GDK_SMOOTH_SCROLL_MASK
	gdk.ALL_EVENTS_MASK = C.GDK_ALL_EVENTS_MASK

	gdk.SCROLL_UP = C.GDK_SCROLL_UP
	gdk.SCROLL_DOWN = C.GDK_SCROLL_DOWN
	gdk.SCROLL_LEFT = C.GDK_SCROLL_LEFT
	gdk.SCROLL_RIGHT = C.GDK_SCROLL_RIGHT
	gdk.SCROLL_SMOOTH = C.GDK_SCROLL_SMOOTH

	gdk.GRAB_SUCCESS = C.GDK_GRAB_SUCCESS
	gdk.GRAB_ALREADY_GRABBED = C.GDK_GRAB_ALREADY_GRABBED
	gdk.GRAB_INVALID_TIME = C.GDK_GRAB_INVALID_TIME
	gdk.GRAB_FROZEN = C.GDK_GRAB_FROZEN
	// Only exists since 3.16
	// GRAB_FAILED GrabStatus = C.GDK_GRAB_FAILED
	gdk.GRAB_FAILED = 5

	gdk.OWNERSHIP_NONE = C.GDK_OWNERSHIP_NONE
	gdk.OWNERSHIP_WINDOW = C.GDK_OWNERSHIP_WINDOW
	gdk.OWNERSHIP_APPLICATION = C.GDK_OWNERSHIP_APPLICATION

	gdk.DEVICE_TYPE_MASTER = C.GDK_DEVICE_TYPE_MASTER
	gdk.DEVICE_TYPE_SLAVE = C.GDK_DEVICE_TYPE_SLAVE
	gdk.DEVICE_TYPE_FLOATING = C.GDK_DEVICE_TYPE_FLOATING

	gdk.EVENT_NOTHING = C.GDK_NOTHING
	gdk.EVENT_DELETE = C.GDK_DELETE
	gdk.EVENT_DESTROY = C.GDK_DESTROY
	gdk.EVENT_EXPOSE = C.GDK_EXPOSE
	gdk.EVENT_MOTION_NOTIFY = C.GDK_MOTION_NOTIFY
	gdk.EVENT_BUTTON_PRESS = C.GDK_BUTTON_PRESS
	gdk.EVENT_2BUTTON_PRESS = C.GDK_2BUTTON_PRESS
	gdk.EVENT_DOUBLE_BUTTON_PRESS = C.GDK_DOUBLE_BUTTON_PRESS
	gdk.EVENT_3BUTTON_PRESS = C.GDK_3BUTTON_PRESS
	gdk.EVENT_TRIPLE_BUTTON_PRESS = C.GDK_TRIPLE_BUTTON_PRESS
	gdk.EVENT_BUTTON_RELEASE = C.GDK_BUTTON_RELEASE
	gdk.EVENT_KEY_PRESS = C.GDK_KEY_PRESS
	gdk.EVENT_KEY_RELEASE = C.GDK_KEY_RELEASE
	gdk.EVENT_LEAVE_NOTIFY = C.GDK_ENTER_NOTIFY
	gdk.EVENT_FOCUS_CHANGE = C.GDK_FOCUS_CHANGE
	gdk.EVENT_CONFIGURE = C.GDK_CONFIGURE
	gdk.EVENT_MAP = C.GDK_MAP
	gdk.EVENT_UNMAP = C.GDK_UNMAP
	gdk.EVENT_PROPERTY_NOTIFY = C.GDK_PROPERTY_NOTIFY
	gdk.EVENT_SELECTION_CLEAR = C.GDK_SELECTION_CLEAR
	gdk.EVENT_SELECTION_REQUEST = C.GDK_SELECTION_REQUEST
	gdk.EVENT_SELECTION_NOTIFY = C.GDK_SELECTION_NOTIFY
	gdk.EVENT_PROXIMITY_IN = C.GDK_PROXIMITY_IN
	gdk.EVENT_PROXIMITY_OUT = C.GDK_PROXIMITY_OUT
	gdk.EVENT_DRAG_ENTER = C.GDK_DRAG_ENTER
	gdk.EVENT_DRAG_LEAVE = C.GDK_DRAG_LEAVE
	gdk.EVENT_DRAG_MOTION = C.GDK_DRAG_MOTION
	gdk.EVENT_DRAG_STATUS = C.GDK_DRAG_STATUS
	gdk.EVENT_DROP_START = C.GDK_DROP_START
	gdk.EVENT_DROP_FINISHED = C.GDK_DROP_FINISHED
	gdk.EVENT_CLIENT_EVENT = C.GDK_CLIENT_EVENT
	gdk.EVENT_VISIBILITY_NOTIFY = C.GDK_VISIBILITY_NOTIFY
	gdk.EVENT_SCROLL = C.GDK_SCROLL
	gdk.EVENT_WINDOW_STATE = C.GDK_WINDOW_STATE
	gdk.EVENT_SETTING = C.GDK_SETTING
	gdk.EVENT_OWNER_CHANGE = C.GDK_OWNER_CHANGE
	gdk.EVENT_GRAB_BROKEN = C.GDK_GRAB_BROKEN
	gdk.EVENT_DAMAGE = C.GDK_DAMAGE
	gdk.EVENT_TOUCH_BEGIN = C.GDK_TOUCH_BEGIN
	gdk.EVENT_TOUCH_UPDATE = C.GDK_TOUCH_UPDATE
	gdk.EVENT_TOUCH_END = C.GDK_TOUCH_END
	gdk.EVENT_TOUCH_CANCEL = C.GDK_TOUCH_CANCEL
	gdk.EVENT_LAST = C.GDK_EVENT_LAST
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
	return gdk.DragAction(c), nil
}

func marshalColorspace(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gdk.Colorspace(c), nil
}

func marshalInterpType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gdk.InterpType(c), nil
}

func marshalModifierType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gdk.ModifierType(c), nil
}

func marshalPixbufAlphaMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gdk.PixbufAlphaMode(c), nil
}

func marshalEventMask(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gdk.EventMask(c), nil
}

/*
 * GdkAtom
 */

// nativeAtom returns the underlying GdkAtom.
func nativeAtom(v gdk.Atom) C.GdkAtom {
	return C.toGdkAtom(unsafe.Pointer(uintptr(v)))
}

// AtomName returns the name of the atom
func AtomName(v gdk.Atom) string {
	c := C.gdk_atom_name(nativeAtom(v))
	defer C.g_free(C.gpointer(c))
	return C.GoString((*C.char)(c))
}

// GdkAtomIntern is a wrapper around gdk_atom_intern
func GdkAtomIntern(atomName string, onlyIfExists bool) gdk.Atom {
	cstr := C.CString(atomName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_atom_intern((*C.gchar)(cstr), gbool(onlyIfExists))
	return gdk.Atom(uintptr(unsafe.Pointer(c)))
}

/*
 * GdkDevice
 */

// Device is a representation of GDK's GdkDevice.
type Device struct {
	*glib_impl.Object
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
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	return &Device{obj}, nil
}

// Grab() is a wrapper around gdk_device_grab().
func (v *Device) Grab(w gdk.Window, ownership gdk.GrabOwnership, owner_events bool, event_mask gdk.EventMask, cursor gdk.Cursor, time uint32) gdk.GrabStatus {
	ret := C.gdk_device_grab(
		v.native(),
		CastToWindow(w).native(),
		C.GdkGrabOwnership(ownership),
		gbool(owner_events),
		C.GdkEventMask(event_mask),
		castToCursor(cursor).native(),
		C.guint32(time),
	)
	return gdk.GrabStatus(ret)
}

// Ungrab() is a wrapper around gdk_device_ungrab().
func (v *Device) Ungrab(time uint32) {
	C.gdk_device_ungrab(v.native(), C.guint32(time))
}

/*
 * GdkCursor
 */

// Cursor is a representation of GdkCursor.
type cursor struct {
	*glib_impl.Object
}

// native returns a pointer to the underlying GdkCursor.
func (v *cursor) native() *C.GdkCursor {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkCursor(p)
}

// Native returns a pointer to the underlying GdkCursor.
func (v *cursor) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalCursor(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	return &cursor{obj}, nil
}

/*
 * GdkDeviceManager
 */

// DeviceManager is a representation of GDK's GdkDeviceManager.
type deviceManager struct {
	*glib_impl.Object
}

// native returns a pointer to the underlying GdkDeviceManager.
func (v *deviceManager) native() *C.GdkDeviceManager {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkDeviceManager(p)
}

// Native returns a pointer to the underlying GdkDeviceManager.
func (v *deviceManager) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalDeviceManager(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	return &deviceManager{obj}, nil
}

// GetClientPointer() is a wrapper around gdk_device_manager_get_client_pointer().
func (v *deviceManager) GetClientPointer() (gdk.Device, error) {
	c := C.gdk_device_manager_get_client_pointer(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
	return &Device{obj}, nil
}

// GetDisplay() is a wrapper around gdk_device_manager_get_display().
func (v *deviceManager) GetDisplay() (gdk.Display, error) {
	c := C.gdk_device_manager_get_display(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
	return &Display{obj}, nil
}

// ListDevices() is a wrapper around gdk_device_manager_list_devices().
func (v *deviceManager) ListDevices(tp gdk.DeviceType) glib.List {
	clist := C.gdk_device_manager_list_devices(v.native(), C.GdkDeviceType(tp))
	if clist == nil {
		return nil
	}
	glist := glib_impl.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		return &Device{&glib_impl.Object{glib_impl.ToGObject(ptr)}}
	})
	runtime.SetFinalizer(glist, func(glist glib.List) {
		glist.Free()
	})
	return glist
}

/*
 * GdkDisplay
 */

// Display is a representation of GDK's GdkDisplay.
type Display struct {
	*glib_impl.Object
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
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	return &Display{obj}, nil
}

func toDisplay(s *C.GdkDisplay) (*Display, error) {
	if s == nil {
		return nil, nilPtrErr
	}
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(s))}
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
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	d := &Display{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
	return d, nil
}

// DisplayGetDefault() is a wrapper around gdk_display_get_default().
func DisplayGetDefault() (*Display, error) {
	c := C.gdk_display_get_default()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	d := &Display{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
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
func (v *Display) GetScreen(screenNum int) (gdk.Screen, error) {
	c := C.gdk_display_get_screen(v.native(), C.gint(screenNum))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	s := &Screen{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
	return s, nil
}

// GetDefaultScreen() is a wrapper around gdk_display_get_default_screen().
func (v *Display) GetDefaultScreen() (gdk.Screen, error) {
	c := C.gdk_display_get_default_screen(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	s := &Screen{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
	return s, nil
}

// GetDeviceManager() is a wrapper around gdk_display_get_device_manager().
func (v *Display) GetDeviceManager() (gdk.DeviceManager, error) {
	c := C.gdk_display_get_device_manager(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	d := &deviceManager{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
	return d, nil
}

// DeviceIsGrabbed() is a wrapper around gdk_display_device_is_grabbed().
func (v *Display) DeviceIsGrabbed(device gdk.Device) bool {
	c := C.gdk_display_device_is_grabbed(v.native(), CastToDevice(device).native())
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
func (v *Display) GetEvent() (gdk.Event, error) {
	c := C.gdk_display_get_event(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	e := &Event{c}
	runtime.SetFinalizer(e, (*Event).free)
	return e, nil
}

// PeekEvent() is a wrapper around gdk_display_peek_event().
func (v *Display) PeekEvent() (gdk.Event, error) {
	c := C.gdk_display_peek_event(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	e := &Event{c}
	runtime.SetFinalizer(e, (*Event).free)
	return e, nil
}

// PutEvent() is a wrapper around gdk_display_put_event().
func (v *Display) PutEvent(event gdk.Event) {
	C.gdk_display_put_event(v.native(), CastToEvent(event).native())
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
func (v *Display) GetDefaultGroup() (gdk.Window, error) {
	c := C.gdk_display_get_default_group(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	w := &Window{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
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
func (v *Display) RequestSelectionNotification(selection gdk.Atom) bool {
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
func (v *Display) StoreClipboard(clipboardWindow gdk.Window, time uint32, targets ...gdk.Atom) {
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
	return gdk.EventType(c), nil
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
type dragContext struct {
	*glib_impl.Object
}

// native returns a pointer to the underlying GdkDragContext.
func (v *dragContext) native() *C.GdkDragContext {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkDragContext(p)
}

// Native returns a pointer to the underlying GdkDragContext.
func (v *dragContext) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalDragContext(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	return &dragContext{obj}, nil
}

func (v *dragContext) ListTargets() glib.List {
	c := C.gdk_drag_context_list_targets(v.native())
	return glib_impl.WrapList(uintptr(unsafe.Pointer(c)))
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
type eventButton struct {
	*Event
}

// Native returns a pointer to the underlying GdkEventButton.
func (v *eventButton) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *eventButton) native() *C.GdkEventButton {
	return (*C.GdkEventButton)(unsafe.Pointer(v.Event.native()))
}

func EventButtonFrom(ev *Event) gdk.EventButton {
	return &eventButton{ev}
}

func (v *eventButton) X() float64 {
	c := v.native().x
	return float64(c)
}

func (v *eventButton) Y() float64 {
	c := v.native().y
	return float64(c)
}

// XRoot returns the x coordinate of the pointer relative to the root of the screen.
func (v *eventButton) XRoot() float64 {
	c := v.native().x_root
	return float64(c)
}

// YRoot returns the y coordinate of the pointer relative to the root of the screen.
func (v *eventButton) YRoot() float64 {
	c := v.native().y_root
	return float64(c)
}

func (v *eventButton) Button() uint {
	c := v.native().button
	return uint(c)
}

func (v *eventButton) State() uint {
	c := v.native().state
	return uint(c)
}

// Time returns the time of the event in milliseconds.
func (v *eventButton) Time() uint32 {
	c := v.native().time
	return uint32(c)
}

func (v *eventButton) Type() gdk.EventType {
	c := v.native()._type
	return gdk.EventType(c)
}

func (v *eventButton) MotionVal() (float64, float64) {
	x := v.native().x
	y := v.native().y
	return float64(x), float64(y)
}

func (v *eventButton) MotionValRoot() (float64, float64) {
	x := v.native().x_root
	y := v.native().y_root
	return float64(x), float64(y)
}

func (v *eventButton) ButtonVal() uint {
	c := v.native().button
	return uint(c)
}

/*
 * GdkEventKey
 */

// EventKey is a representation of GDK's GdkEventKey.
type eventKey struct {
	*Event
}

func EventKeyNew() *eventKey {
	ee := (*C.GdkEvent)(unsafe.Pointer(&C.GdkEventKey{}))
	ev := Event{ee}
	return &eventKey{&ev}
}

// Native returns a pointer to the underlying GdkEventKey.
func (v *eventKey) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *eventKey) native() *C.GdkEventKey {
	return (*C.GdkEventKey)(unsafe.Pointer(v.Event.native()))
}

func (v *eventKey) KeyVal() uint {
	c := v.native().keyval
	return uint(c)
}

func (v *eventKey) Type() gdk.EventType {
	c := v.native()._type
	return gdk.EventType(c)
}

func (v *eventKey) State() uint {
	c := v.native().state
	return uint(c)
}

/*
 * GdkEventMotion
 */

type eventMotion struct {
	*Event
}

// Native returns a pointer to the underlying GdkEventMotion.
func (v *eventMotion) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *eventMotion) native() *C.GdkEventMotion {
	return (*C.GdkEventMotion)(unsafe.Pointer(v.Event.native()))
}

func (v *eventMotion) MotionVal() (float64, float64) {
	x := v.native().x
	y := v.native().y
	return float64(x), float64(y)
}

func (v *eventMotion) MotionValRoot() (float64, float64) {
	x := v.native().x_root
	y := v.native().y_root
	return float64(x), float64(y)
}

/*
 * GdkEventScroll
 */

// EventScroll is a representation of GDK's GdkEventScroll.
type eventScroll struct {
	*Event
}

// Native returns a pointer to the underlying GdkEventScroll.
func (v *eventScroll) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *eventScroll) native() *C.GdkEventScroll {
	return (*C.GdkEventScroll)(unsafe.Pointer(v.Event.native()))
}

func (v *eventScroll) DeltaX() float64 {
	return float64(v.native().delta_x)
}

func (v *eventScroll) DeltaY() float64 {
	return float64(v.native().delta_y)
}

func (v *eventScroll) X() float64 {
	return float64(v.native().x)
}

func (v *eventScroll) Y() float64 {
	return float64(v.native().y)
}

func (v *eventScroll) Type() gdk.EventType {
	c := v.native()._type
	return gdk.EventType(c)
}

/*
 * GdkPixbuf
 */

// Pixbuf is a representation of GDK's GdkPixbuf.
type Pixbuf struct {
	*glib_impl.Object
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
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	return &Pixbuf{obj}, nil
}

// GetColorspace is a wrapper around gdk_pixbuf_get_colorspace().
func (v *Pixbuf) GetColorspace() gdk.Colorspace {
	c := C.gdk_pixbuf_get_colorspace(v.native())
	return gdk.Colorspace(c)
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
func PixbufNew(colorspace gdk.Colorspace, hasAlpha bool, bitsPerSample, width, height int) (*Pixbuf, error) {
	c := C.gdk_pixbuf_new(C.GdkColorspace(colorspace), gbool(hasAlpha),
		C.int(bitsPerSample), C.int(width), C.int(height))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
	return p, nil
}

// PixbufCopy is a wrapper around gdk_pixbuf_copy().
func PixbufCopy(v *Pixbuf) (*Pixbuf, error) {
	c := C.gdk_pixbuf_copy(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
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
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(res))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
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
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(res))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
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
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(res))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
	return p, nil
}

// ScaleSimple is a wrapper around gdk_pixbuf_scale_simple().
func (v *Pixbuf) ScaleSimple(destWidth, destHeight int, interpType gdk.InterpType) (gdk.Pixbuf, error) {
	c := C.gdk_pixbuf_scale_simple(v.native(), C.int(destWidth),
		C.int(destHeight), C.GdkInterpType(interpType))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
	return p, nil
}

// RotateSimple is a wrapper around gdk_pixbuf_rotate_simple().
func (v *Pixbuf) RotateSimple(angle gdk.PixbufRotation) (gdk.Pixbuf, error) {
	c := C.gdk_pixbuf_rotate_simple(v.native(), C.GdkPixbufRotation(angle))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
	return p, nil
}

// ApplyEmbeddedOrientation is a wrapper around gdk_pixbuf_apply_embedded_orientation().
func (v *Pixbuf) ApplyEmbeddedOrientation() (gdk.Pixbuf, error) {
	c := C.gdk_pixbuf_apply_embedded_orientation(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
	return p, nil
}

// Flip is a wrapper around gdk_pixbuf_flip().
func (v *Pixbuf) Flip(horizontal bool) (gdk.Pixbuf, error) {
	c := C.gdk_pixbuf_flip(v.native(), gbool(horizontal))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
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
type pixbufLoader struct {
	*glib_impl.Object
}

// native() returns a pointer to the underlying GdkPixbufLoader.
func (v *pixbufLoader) native() *C.GdkPixbufLoader {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkPixbufLoader(p)
}

// PixbufLoaderNew() is a wrapper around gdk_pixbuf_loader_new().
func PixbufLoaderNew() (*pixbufLoader, error) {
	c := C.gdk_pixbuf_loader_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	p := &pixbufLoader{obj}
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
	return p, nil
}

// Write() is a wrapper around gdk_pixbuf_loader_write().  The
// function signature differs from the C equivalent to satisify the
// io.Writer interface.
func (v *pixbufLoader) Write(data []byte) (int, error) {
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
func (v *pixbufLoader) Close() error {
	var err *C.GError

	if ok := gobool(C.gdk_pixbuf_loader_close(v.native(), &err)); !ok {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

// SetSize is a wrapper around gdk_pixbuf_loader_set_size().
func (v *pixbufLoader) SetSize(width, height int) {
	C.gdk_pixbuf_loader_set_size(v.native(), C.int(width), C.int(height))
}

// GetPixbuf is a wrapper around gdk_pixbuf_loader_get_pixbuf().
func (v *pixbufLoader) GetPixbuf() (gdk.Pixbuf, error) {
	c := C.gdk_pixbuf_loader_get_pixbuf(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib_impl.Object).Unref)
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
func PixbufGetType() glib.Type {
	return glib.Type(C.gdk_pixbuf_get_type())
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
	*glib_impl.Object
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
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	return &Visual{obj}, nil
}

/*
 * GdkWindow
 */

// Window is a representation of GDK's GdkWindow.
type Window struct {
	*glib_impl.Object
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
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(c))}
	return &Window{obj}, nil
}

func toWindow(s *C.GdkWindow) (*Window, error) {
	if s == nil {
		return nil, nilPtrErr
	}
	obj := &glib_impl.Object{glib_impl.ToGObject(unsafe.Pointer(s))}
	return &Window{obj}, nil
}

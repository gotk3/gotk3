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

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import (
	"errors"
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.gdk_drag_action_get_type()), marshalDragAction},
		{glib.Type(C.gdk_colorspace_get_type()), marshalColorspace},
		{glib.Type(C.gdk_event_type_get_type()), marshalEventType},
		{glib.Type(C.gdk_interp_type_get_type()), marshalInterpType},
		{glib.Type(C.gdk_modifier_type_get_type()), marshalModifierType},
		{glib.Type(C.gdk_event_mask_get_type()), marshalEventMask},
		{glib.Type(C.gdk_gravity_get_type()), marshalGravity},

		// Objects/Interfaces
		{glib.Type(C.gdk_device_get_type()), marshalDevice},
		{glib.Type(C.gdk_cursor_get_type()), marshalCursor},
		{glib.Type(C.gdk_device_manager_get_type()), marshalDeviceManager},
		{glib.Type(C.gdk_display_get_type()), marshalDisplay},
		{glib.Type(C.gdk_drag_context_get_type()), marshalDragContext},
		{glib.Type(C.gdk_rgba_get_type()), marshalRGBA},
		{glib.Type(C.gdk_screen_get_type()), marshalScreen},
		{glib.Type(C.gdk_visual_get_type()), marshalVisual},
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

// DragAction is a representation of GDK's GdkDragAction.
type DragAction int

const (
	ACTION_DEFAULT DragAction = C.GDK_ACTION_DEFAULT
	ACTION_COPY    DragAction = C.GDK_ACTION_COPY
	ACTION_MOVE    DragAction = C.GDK_ACTION_MOVE
	ACTION_LINK    DragAction = C.GDK_ACTION_LINK
	ACTION_PRIVATE DragAction = C.GDK_ACTION_PRIVATE
	ACTION_ASK     DragAction = C.GDK_ACTION_ASK
)

func marshalDragAction(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return DragAction(c), nil
}

// Colorspace is a representation of GDK's GdkColorspace.
type Colorspace int

const (
	COLORSPACE_RGB Colorspace = C.GDK_COLORSPACE_RGB
)

func marshalColorspace(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Colorspace(c), nil
}

// InterpType is a representation of GDK's GdkInterpType.
type InterpType int

const (
	INTERP_NEAREST  InterpType = C.GDK_INTERP_NEAREST
	INTERP_TILES    InterpType = C.GDK_INTERP_TILES
	INTERP_BILINEAR InterpType = C.GDK_INTERP_BILINEAR
	INTERP_HYPER    InterpType = C.GDK_INTERP_HYPER
)

func marshalInterpType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return InterpType(c), nil
}

// ModifierType is a representation of GDK's GdkModifierType.
type ModifierType uint

const (
	SHIFT_MASK    ModifierType = C.GDK_SHIFT_MASK
	LOCK_MASK                  = C.GDK_LOCK_MASK
	CONTROL_MASK               = C.GDK_CONTROL_MASK
	MOD1_MASK                  = C.GDK_MOD1_MASK
	MOD2_MASK                  = C.GDK_MOD2_MASK
	MOD3_MASK                  = C.GDK_MOD3_MASK
	MOD4_MASK                  = C.GDK_MOD4_MASK
	MOD5_MASK                  = C.GDK_MOD5_MASK
	BUTTON1_MASK               = C.GDK_BUTTON1_MASK
	BUTTON2_MASK               = C.GDK_BUTTON2_MASK
	BUTTON3_MASK               = C.GDK_BUTTON3_MASK
	BUTTON4_MASK               = C.GDK_BUTTON4_MASK
	BUTTON5_MASK               = C.GDK_BUTTON5_MASK
	SUPER_MASK                 = C.GDK_SUPER_MASK
	HYPER_MASK                 = C.GDK_HYPER_MASK
	META_MASK                  = C.GDK_META_MASK
	RELEASE_MASK               = C.GDK_RELEASE_MASK
	MODIFIER_MASK              = C.GDK_MODIFIER_MASK
)

func marshalModifierType(p uintptr) (interface{}, error) {
	c := C.g_value_get_flags((*C.GValue)(unsafe.Pointer(p)))
	return ModifierType(c), nil
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

// added by terrak
// EventMask is a representation of GDK's GdkEventMask.
type EventMask int

const (
	EXPOSURE_MASK            EventMask = C.GDK_EXPOSURE_MASK
	POINTER_MOTION_MASK      EventMask = C.GDK_POINTER_MOTION_MASK
	POINTER_MOTION_HINT_MASK EventMask = C.GDK_POINTER_MOTION_HINT_MASK
	BUTTON_MOTION_MASK       EventMask = C.GDK_BUTTON_MOTION_MASK
	BUTTON1_MOTION_MASK      EventMask = C.GDK_BUTTON1_MOTION_MASK
	BUTTON2_MOTION_MASK      EventMask = C.GDK_BUTTON2_MOTION_MASK
	BUTTON3_MOTION_MASK      EventMask = C.GDK_BUTTON3_MOTION_MASK
	BUTTON_PRESS_MASK        EventMask = C.GDK_BUTTON_PRESS_MASK
	BUTTON_RELEASE_MASK      EventMask = C.GDK_BUTTON_RELEASE_MASK
	KEY_PRESS_MASK           EventMask = C.GDK_KEY_PRESS_MASK
	KEY_RELEASE_MASK         EventMask = C.GDK_KEY_RELEASE_MASK
	ENTER_NOTIFY_MASK        EventMask = C.GDK_ENTER_NOTIFY_MASK
	LEAVE_NOTIFY_MASK        EventMask = C.GDK_LEAVE_NOTIFY_MASK
	FOCUS_CHANGE_MASK        EventMask = C.GDK_FOCUS_CHANGE_MASK
	STRUCTURE_MASK           EventMask = C.GDK_STRUCTURE_MASK
	PROPERTY_CHANGE_MASK     EventMask = C.GDK_PROPERTY_CHANGE_MASK
	VISIBILITY_NOTIFY_MASK   EventMask = C.GDK_VISIBILITY_NOTIFY_MASK
	PROXIMITY_IN_MASK        EventMask = C.GDK_PROXIMITY_IN_MASK
	PROXIMITY_OUT_MASK       EventMask = C.GDK_PROXIMITY_OUT_MASK
	SUBSTRUCTURE_MASK        EventMask = C.GDK_SUBSTRUCTURE_MASK
	SCROLL_MASK              EventMask = C.GDK_SCROLL_MASK
	TOUCH_MASK               EventMask = C.GDK_TOUCH_MASK
	SMOOTH_SCROLL_MASK       EventMask = C.GDK_SMOOTH_SCROLL_MASK
	ALL_EVENTS_MASK          EventMask = C.GDK_ALL_EVENTS_MASK
)

func marshalEventMask(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return EventMask(c), nil
}

// added by lazyshot
// ScrollDirection is a representation of GDK's GdkScrollDirection
type ScrollDirection int

const (
	SCROLL_UP     ScrollDirection = C.GDK_SCROLL_UP
	SCROLL_DOWN   ScrollDirection = C.GDK_SCROLL_DOWN
	SCROLL_LEFT   ScrollDirection = C.GDK_SCROLL_LEFT
	SCROLL_RIGHT  ScrollDirection = C.GDK_SCROLL_RIGHT
	SCROLL_SMOOTH ScrollDirection = C.GDK_SCROLL_SMOOTH
)

// WindowEdge is a representation of GDK's GdkWindowEdge
type WindowEdge int

const (
	WINDOW_EDGE_NORTH_WEST WindowEdge = C.GDK_WINDOW_EDGE_NORTH_WEST
	WINDOW_EDGE_NORTH      WindowEdge = C.GDK_WINDOW_EDGE_NORTH
	WINDOW_EDGE_NORTH_EAST WindowEdge = C.GDK_WINDOW_EDGE_NORTH_EAST
	WINDOW_EDGE_WEST       WindowEdge = C.GDK_WINDOW_EDGE_WEST
	WINDOW_EDGE_EAST       WindowEdge = C.GDK_WINDOW_EDGE_EAST
	WINDOW_EDGE_SOUTH_WEST WindowEdge = C.GDK_WINDOW_EDGE_SOUTH_WEST
	WINDOW_EDGE_SOUTH      WindowEdge = C.GDK_WINDOW_EDGE_SOUTH
	WINDOW_EDGE_SOUTH_EAST WindowEdge = C.GDK_WINDOW_EDGE_SOUTH_EAST
)

// WindowState is a representation of GDK's GdkWindowState
type WindowState int

const (
	WINDOW_STATE_WITHDRAWN  WindowState = C.GDK_WINDOW_STATE_WITHDRAWN
	WINDOW_STATE_ICONIFIED  WindowState = C.GDK_WINDOW_STATE_ICONIFIED
	WINDOW_STATE_MAXIMIZED  WindowState = C.GDK_WINDOW_STATE_MAXIMIZED
	WINDOW_STATE_STICKY     WindowState = C.GDK_WINDOW_STATE_STICKY
	WINDOW_STATE_FULLSCREEN WindowState = C.GDK_WINDOW_STATE_FULLSCREEN
	WINDOW_STATE_ABOVE      WindowState = C.GDK_WINDOW_STATE_ABOVE
	WINDOW_STATE_BELOW      WindowState = C.GDK_WINDOW_STATE_BELOW
	WINDOW_STATE_FOCUSED    WindowState = C.GDK_WINDOW_STATE_FOCUSED
	WINDOW_STATE_TILED      WindowState = C.GDK_WINDOW_STATE_TILED
)

// WindowTypeHint is a representation of GDK's GdkWindowTypeHint
type WindowTypeHint int

const (
	WINDOW_TYPE_HINT_NORMAL        WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_NORMAL
	WINDOW_TYPE_HINT_DIALOG        WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_DIALOG
	WINDOW_TYPE_HINT_MENU          WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_MENU
	WINDOW_TYPE_HINT_TOOLBAR       WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_TOOLBAR
	WINDOW_TYPE_HINT_SPLASHSCREEN  WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_SPLASHSCREEN
	WINDOW_TYPE_HINT_UTILITY       WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_UTILITY
	WINDOW_TYPE_HINT_DOCK          WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_DOCK
	WINDOW_TYPE_HINT_DESKTOP       WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_DESKTOP
	WINDOW_TYPE_HINT_DROPDOWN_MENU WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_DROPDOWN_MENU
	WINDOW_TYPE_HINT_POPUP_MENU    WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_POPUP_MENU
	WINDOW_TYPE_HINT_TOOLTIP       WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_TOOLTIP
	WINDOW_TYPE_HINT_NOTIFICATION  WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_NOTIFICATION
	WINDOW_TYPE_HINT_COMBO         WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_COMBO
	WINDOW_TYPE_HINT_DND           WindowTypeHint = C.GDK_WINDOW_TYPE_HINT_DND
)

// WindowHints is a representation of GDK's GdkWindowHints
type WindowHints int

const (
	HINT_POS         WindowHints = C.GDK_HINT_POS
	HINT_MIN_SIZE    WindowHints = C.GDK_HINT_MIN_SIZE
	HINT_MAX_SIZE    WindowHints = C.GDK_HINT_MAX_SIZE
	HINT_BASE_SIZE   WindowHints = C.GDK_HINT_BASE_SIZE
	HINT_ASPECT      WindowHints = C.GDK_HINT_ASPECT
	HINT_RESIZE_INC  WindowHints = C.GDK_HINT_RESIZE_INC
	HINT_WIN_GRAVITY WindowHints = C.GDK_HINT_WIN_GRAVITY
	HINT_USER_POS    WindowHints = C.GDK_HINT_USER_POS
	HINT_USER_SIZE   WindowHints = C.GDK_HINT_USER_SIZE
)

// CURRENT_TIME is a representation of GDK_CURRENT_TIME

const CURRENT_TIME = C.GDK_CURRENT_TIME

// GrabStatus is a representation of GdkGrabStatus

type GrabStatus int

const (
	GRAB_SUCCESS         GrabStatus = C.GDK_GRAB_SUCCESS
	GRAB_ALREADY_GRABBED GrabStatus = C.GDK_GRAB_ALREADY_GRABBED
	GRAB_INVALID_TIME    GrabStatus = C.GDK_GRAB_INVALID_TIME
	GRAB_NOT_VIEWABLE    GrabStatus = C.GDK_GRAB_NOT_VIEWABLE
	GRAB_FROZEN          GrabStatus = C.GDK_GRAB_FROZEN
)

// GrabOwnership is a representation of GdkGrabOwnership

type GrabOwnership int

const (
	OWNERSHIP_NONE        GrabOwnership = C.GDK_OWNERSHIP_NONE
	OWNERSHIP_WINDOW      GrabOwnership = C.GDK_OWNERSHIP_WINDOW
	OWNERSHIP_APPLICATION GrabOwnership = C.GDK_OWNERSHIP_APPLICATION
)

// TODO:
// GdkInputSource
// GdkInputMode
// GdkAxisUse
// GdkAxisFlags
// GdkDeviceToolType

// DeviceType is a representation of GdkDeviceType

type DeviceType int

const (
	DEVICE_TYPE_MASTER   DeviceType = C.GDK_DEVICE_TYPE_MASTER
	DEVICE_TYPE_SLAVE    DeviceType = C.GDK_DEVICE_TYPE_SLAVE
	DEVICE_TYPE_FLOATING DeviceType = C.GDK_DEVICE_TYPE_FLOATING
)

// TODO:
// GdkColorspace
// GdkVisualType
// GdkTimeCoord

// EventPropagation constants

const (
	GDK_EVENT_PROPAGATE bool = C.GDK_EVENT_PROPAGATE != 0
	GDK_EVENT_STOP      bool = C.GDK_EVENT_STOP != 0
)

// Button constants
type Button uint

const (
	BUTTON_PRIMARY   Button = C.GDK_BUTTON_PRIMARY
	BUTTON_MIDDLE    Button = C.GDK_BUTTON_MIDDLE
	BUTTON_SECONDARY Button = C.GDK_BUTTON_SECONDARY
)

// CrossingMode is a representation of GDK's GdkCrossingMode.

type CrossingMode int

const (
	CROSSING_NORMAL        CrossingMode = C.GDK_CROSSING_NORMAL
	CROSSING_GRAB          CrossingMode = C.GDK_CROSSING_GRAB
	CROSSING_UNGRAB        CrossingMode = C.GDK_CROSSING_UNGRAB
	CROSSING_GTK_GRAB      CrossingMode = C.GDK_CROSSING_GTK_GRAB
	CROSSING_GTK_UNGRAB    CrossingMode = C.GDK_CROSSING_GTK_UNGRAB
	CROSSING_STATE_CHANGED CrossingMode = C.GDK_CROSSING_STATE_CHANGED
	CROSSING_TOUCH_BEGIN   CrossingMode = C.GDK_CROSSING_TOUCH_BEGIN
	CROSSING_TOUCH_END     CrossingMode = C.GDK_CROSSING_TOUCH_END
	CROSSING_DEVICE_SWITCH CrossingMode = C.GDK_CROSSING_DEVICE_SWITCH
)

// NotifyType is a representation of GDK's GdkNotifyType.

type NotifyType int

const (
	NOTIFY_ANCESTOR          NotifyType = C.GDK_NOTIFY_ANCESTOR
	NOTIFY_VIRTUAL           NotifyType = C.GDK_NOTIFY_VIRTUAL
	NOTIFY_INFERIOR          NotifyType = C.GDK_NOTIFY_INFERIOR
	NOTIFY_NONLINEAR         NotifyType = C.GDK_NOTIFY_NONLINEAR
	NOTIFY_NONLINEAR_VIRTUAL NotifyType = C.GDK_NOTIFY_NONLINEAR_VIRTUAL
	NOTIFY_UNKNOWN           NotifyType = C.GDK_NOTIFY_UNKNOWN
)

// EventType is a representation of GDK's GdkEventType.
// Do not confuse these event types with the signals that GTK+ widgets emit
type EventType int

func marshalEventType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return EventType(c), nil
}

const (
	EVENT_NOTHING             EventType = C.GDK_NOTHING
	EVENT_DELETE              EventType = C.GDK_DELETE
	EVENT_DESTROY             EventType = C.GDK_DESTROY
	EVENT_EXPOSE              EventType = C.GDK_EXPOSE
	EVENT_MOTION_NOTIFY       EventType = C.GDK_MOTION_NOTIFY
	EVENT_BUTTON_PRESS        EventType = C.GDK_BUTTON_PRESS
	EVENT_2BUTTON_PRESS       EventType = C.GDK_2BUTTON_PRESS
	EVENT_DOUBLE_BUTTON_PRESS EventType = C.GDK_DOUBLE_BUTTON_PRESS
	EVENT_3BUTTON_PRESS       EventType = C.GDK_3BUTTON_PRESS
	EVENT_TRIPLE_BUTTON_PRESS EventType = C.GDK_TRIPLE_BUTTON_PRESS
	EVENT_BUTTON_RELEASE      EventType = C.GDK_BUTTON_RELEASE
	EVENT_KEY_PRESS           EventType = C.GDK_KEY_PRESS
	EVENT_KEY_RELEASE         EventType = C.GDK_KEY_RELEASE
	EVENT_ENTER_NOTIFY        EventType = C.GDK_ENTER_NOTIFY
	EVENT_LEAVE_NOTIFY        EventType = C.GDK_LEAVE_NOTIFY
	EVENT_FOCUS_CHANGE        EventType = C.GDK_FOCUS_CHANGE
	EVENT_CONFIGURE           EventType = C.GDK_CONFIGURE
	EVENT_MAP                 EventType = C.GDK_MAP
	EVENT_UNMAP               EventType = C.GDK_UNMAP
	EVENT_PROPERTY_NOTIFY     EventType = C.GDK_PROPERTY_NOTIFY
	EVENT_SELECTION_CLEAR     EventType = C.GDK_SELECTION_CLEAR
	EVENT_SELECTION_REQUEST   EventType = C.GDK_SELECTION_REQUEST
	EVENT_SELECTION_NOTIFY    EventType = C.GDK_SELECTION_NOTIFY
	EVENT_PROXIMITY_IN        EventType = C.GDK_PROXIMITY_IN
	EVENT_PROXIMITY_OUT       EventType = C.GDK_PROXIMITY_OUT
	EVENT_DRAG_ENTER          EventType = C.GDK_DRAG_ENTER
	EVENT_DRAG_LEAVE          EventType = C.GDK_DRAG_LEAVE
	EVENT_DRAG_MOTION         EventType = C.GDK_DRAG_MOTION
	EVENT_DRAG_STATUS         EventType = C.GDK_DRAG_STATUS
	EVENT_DROP_START          EventType = C.GDK_DROP_START
	EVENT_DROP_FINISHED       EventType = C.GDK_DROP_FINISHED
	EVENT_CLIENT_EVENT        EventType = C.GDK_CLIENT_EVENT
	EVENT_VISIBILITY_NOTIFY   EventType = C.GDK_VISIBILITY_NOTIFY
	EVENT_SCROLL              EventType = C.GDK_SCROLL
	EVENT_WINDOW_STATE        EventType = C.GDK_WINDOW_STATE
	EVENT_SETTING             EventType = C.GDK_SETTING
	EVENT_OWNER_CHANGE        EventType = C.GDK_OWNER_CHANGE
	EVENT_GRAB_BROKEN         EventType = C.GDK_GRAB_BROKEN
	EVENT_DAMAGE              EventType = C.GDK_DAMAGE
	EVENT_TOUCH_BEGIN         EventType = C.GDK_TOUCH_BEGIN
	EVENT_TOUCH_UPDATE        EventType = C.GDK_TOUCH_UPDATE
	EVENT_TOUCH_END           EventType = C.GDK_TOUCH_END
	EVENT_TOUCH_CANCEL        EventType = C.GDK_TOUCH_CANCEL
	EVENT_LAST                EventType = C.GDK_EVENT_LAST
)

/*
 * General
 */

// TODO:
// gdk_init().
// gdk_init_check().
// gdk_parse_args().
// gdk_get_display_arg_name().
// gdk_notify_startup_complete().
// gdk_notify_startup_complete_with_id().
// gdk_get_program_class().
// gdk_set_program_class().
// gdk_get_display(). deprecated since version 3.8
// gdk_flush(). deprecated
// gdk_screen_width(). deprecated since version 3.22
// gdk_screen_height(). deprecated since version 3.22
// gdk_screen_width_mm(). deprecated since version 3.22
// gdk_screen_height_mm(). deprecated since version 3.22
// gdk_set_double_click_time(). deprecated
// gdk_beep(). deprecated
// gdk_error_trap_push(). deprecated
// gdk_error_trap_pop(). deprecated
// gdk_error_trap_pop_ignored(). deprecated

// SetAllowedBackends is a wrapper around gdk_set_allowed_backends
func SetAllowedBackends(backends string) {
	cstr := C.CString(backends)
	defer C.free(unsafe.Pointer(cstr))
	C.gdk_set_allowed_backends((*C.gchar)(cstr))
}

/*
 * GdkAtom
 */

// Atom is a representation of GDK's GdkAtom.
type Atom uintptr

// native returns the underlying GdkAtom.
func (v Atom) native() C.GdkAtom {
	return C.toGdkAtom(unsafe.Pointer(uintptr(v)))
}

func (v Atom) Name() string {
	c := C.gdk_atom_name(v.native())
	defer C.g_free(C.gpointer(c))
	return C.GoString((*C.char)(c))
}

// GdkAtomIntern is a wrapper around gdk_atom_intern
func GdkAtomIntern(atomName string, onlyIfExists bool) Atom {
	cstr := C.CString(atomName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_atom_intern((*C.gchar)(cstr), gbool(onlyIfExists))
	return Atom(uintptr(unsafe.Pointer(c)))
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

func toDevice(d *C.GdkDevice) (*Device, error) {
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(d))}
	return &Device{obj}, nil
}

func (v *Device) GetPosition(screen **Screen, x, y *int) error {
	cs := (**C.GdkScreen)(unsafe.Pointer(uintptr(0)))
	if screen != nil {
		var cval *C.GdkScreen
		cs = &cval
	}

	cx := (*C.gint)(unsafe.Pointer(uintptr(0)))
	if x != nil {
		var cval C.gint
		cx = &cval
	}

	cy := (*C.gint)(unsafe.Pointer(uintptr(0)))
	if y != nil {
		var cval C.gint
		cy = &cval
	}
	C.gdk_device_get_position(v.native(), cs, cx, cy)

	if cs != (**C.GdkScreen)(unsafe.Pointer(uintptr(0))) {
		ms, err := toScreen(*cs)
		if err != nil {
			return err
		}
		*screen = ms
	}
	if cx != (*C.gint)(unsafe.Pointer(uintptr(0))) {
		*x = int(*cx)
	}
	if cy != (*C.gint)(unsafe.Pointer(uintptr(0))) {
		*y = int(*cy)
	}
	return nil
}

// TODO:
// gdk_device_get_name().
// gdk_device_get_source().
// gdk_device_set_mode().
// gdk_device_get_mode().
// gdk_device_set_key().
// gdk_device_get_key().
// gdk_device_set_axis_use().
// gdk_device_get_axis_use().
// gdk_device_get_associated_device().
// gdk_device_list_slave_devices().
// gdk_device_get_device_type().
// gdk_device_get_display().
// gdk_device_get_has_cursor().
// gdk_device_get_n_axes().
// gdk_device_get_n_keys().
// gdk_device_warp().
// gdk_device_get_state().
// gdk_device_get_window_at_position().
// gdk_device_get_window_at_position_double().
// gdk_device_get_history().
// gdk_device_free_history().
// gdk_device_get_axis().
// gdk_device_list_axes().
// gdk_device_get_axis_value().

/*
 * GdkCursor
 */

// Cursor is a representation of GdkCursor.
type Cursor struct {
	*glib.Object
}

// CursorNewFromName is a wrapper around gdk_cursor_new_from_name().
func CursorNewFromName(display *Display, name string) (*Cursor, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_cursor_new_from_name(display.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}

	return &Cursor{glib.Take(unsafe.Pointer(c))}, nil
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

// GetDisplay() is a wrapper around gdk_device_manager_get_display().
func (v *DeviceManager) GetDisplay() (*Display, error) {
	c := C.gdk_device_manager_get_display(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Display{glib.Take(unsafe.Pointer(c))}, nil
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

// DisplayOpen is a wrapper around gdk_display_open().
func DisplayOpen(displayName string) (*Display, error) {
	cstr := C.CString(displayName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gdk_display_open((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}

	return &Display{glib.Take(unsafe.Pointer(c))}, nil
}

// DisplayGetDefault is a wrapper around gdk_display_get_default().
func DisplayGetDefault() (*Display, error) {
	c := C.gdk_display_get_default()
	if c == nil {
		return nil, nilPtrErr
	}

	return &Display{glib.Take(unsafe.Pointer(c))}, nil
}

// GetName is a wrapper around gdk_display_get_name().
func (v *Display) GetName() (string, error) {
	c := C.gdk_display_get_name(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// GetDefaultScreen is a wrapper around gdk_display_get_default_screen().
func (v *Display) GetDefaultScreen() (*Screen, error) {
	c := C.gdk_display_get_default_screen(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Screen{glib.Take(unsafe.Pointer(c))}, nil
}

// DeviceIsGrabbed is a wrapper around gdk_display_device_is_grabbed().
func (v *Display) DeviceIsGrabbed(device *Device) bool {
	c := C.gdk_display_device_is_grabbed(v.native(), device.native())
	return gobool(c)
}

// Beep is a wrapper around gdk_display_beep().
func (v *Display) Beep() {
	C.gdk_display_beep(v.native())
}

// Sync is a wrapper around gdk_display_sync().
func (v *Display) Sync() {
	C.gdk_display_sync(v.native())
}

// Flush is a wrapper around gdk_display_flush().
func (v *Display) Flush() {
	C.gdk_display_flush(v.native())
}

// Close is a wrapper around gdk_display_close().
func (v *Display) Close() {
	C.gdk_display_close(v.native())
}

// IsClosed is a wrapper around gdk_display_is_closed().
func (v *Display) IsClosed() bool {
	c := C.gdk_display_is_closed(v.native())
	return gobool(c)
}

// GetEvent is a wrapper around gdk_display_get_event().
func (v *Display) GetEvent() (*Event, error) {
	c := C.gdk_display_get_event(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	//The finalizer is not on the glib.Object but on the event.
	e := &Event{c}
	runtime.SetFinalizer(e, (*Event).free)
	return e, nil
}

// PeekEvent is a wrapper around gdk_display_peek_event().
func (v *Display) PeekEvent() (*Event, error) {
	c := C.gdk_display_peek_event(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	//The finalizer is not on the glib.Object but on the event.
	e := &Event{c}
	runtime.SetFinalizer(e, (*Event).free)
	return e, nil
}

// PutEvent is a wrapper around gdk_display_put_event().
func (v *Display) PutEvent(event *Event) {
	C.gdk_display_put_event(v.native(), event.native())
}

// HasPending is a wrapper around gdk_display_has_pending().
func (v *Display) HasPending() bool {
	c := C.gdk_display_has_pending(v.native())
	return gobool(c)
}

// SetDoubleClickTime is a wrapper around gdk_display_set_double_click_time().
func (v *Display) SetDoubleClickTime(msec uint) {
	C.gdk_display_set_double_click_time(v.native(), C.guint(msec))
}

// SetDoubleClickDistance is a wrapper around gdk_display_set_double_click_distance().
func (v *Display) SetDoubleClickDistance(distance uint) {
	C.gdk_display_set_double_click_distance(v.native(), C.guint(distance))
}

// SupportsColorCursor is a wrapper around gdk_display_supports_cursor_color().
func (v *Display) SupportsColorCursor() bool {
	c := C.gdk_display_supports_cursor_color(v.native())
	return gobool(c)
}

// SupportsCursorAlpha is a wrapper around gdk_display_supports_cursor_alpha().
func (v *Display) SupportsCursorAlpha() bool {
	c := C.gdk_display_supports_cursor_alpha(v.native())
	return gobool(c)
}

// GetDefaultCursorSize is a wrapper around gdk_display_get_default_cursor_size().
func (v *Display) GetDefaultCursorSize() uint {
	c := C.gdk_display_get_default_cursor_size(v.native())
	return uint(c)
}

// GetMaximalCursorSize is a wrapper around gdk_display_get_maximal_cursor_size().
func (v *Display) GetMaximalCursorSize() (width, height uint) {
	var w, h C.guint
	C.gdk_display_get_maximal_cursor_size(v.native(), &w, &h)
	return uint(w), uint(h)
}

// GetDefaultGroup is a wrapper around gdk_display_get_default_group().
func (v *Display) GetDefaultGroup() (*Window, error) {
	c := C.gdk_display_get_default_group(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return &Window{glib.Take(unsafe.Pointer(c))}, nil
}

// SupportsSelectionNotification is a wrapper around gdk_display_supports_selection_notification().
func (v *Display) SupportsSelectionNotification() bool {
	c := C.gdk_display_supports_selection_notification(v.native())
	return gobool(c)
}

// RequestSelectionNotification is a wrapper around gdk_display_request_selection_notification().
func (v *Display) RequestSelectionNotification(selection Atom) bool {
	c := C.gdk_display_request_selection_notification(v.native(),
		selection.native())
	return gobool(c)
}

// SupportsClipboardPersistence is a wrapper around gdk_display_supports_clipboard_persistence().
func (v *Display) SupportsClipboardPersistence() bool {
	c := C.gdk_display_supports_clipboard_persistence(v.native())
	return gobool(c)
}

// TODO:
// gdk_display_store_clipboard().
// func (v *Display) StoreClipboard(clipboardWindow *Window, time uint32, targets ...Atom) {
// 	panic("Not implemented")
// }

// SupportsShapes is a wrapper around gdk_display_supports_shapes().
func (v *Display) SupportsShapes() bool {
	c := C.gdk_display_supports_shapes(v.native())
	return gobool(c)
}

// SupportsInputShapes is a wrapper around gdk_display_supports_input_shapes().
func (v *Display) SupportsInputShapes() bool {
	c := C.gdk_display_supports_input_shapes(v.native())
	return gobool(c)
}

// TODO:
// gdk_display_get_app_launch_context().
// func (v *Display) GetAppLaunchContext() {
// 	panic("Not implemented")
// }

// NotifyStartupComplete is a wrapper around gdk_display_notify_startup_complete().
func (v *Display) NotifyStartupComplete(startupID string) {
	cstr := C.CString(startupID)
	defer C.free(unsafe.Pointer(cstr))
	C.gdk_display_notify_startup_complete(v.native(), (*C.gchar)(cstr))
}

/*
 * GdkDisplayManager
 */

// TODO:
// gdk_display_manager_get().
// gdk_display_manager_get_default_display().
// gdk_display_manager_set_default_display().
// gdk_display_manager_list_displays().
// gdk_display_manager_open_display().

/*
 * GdkKeymap
 */

type Keymap struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkKeymap.
func (v *Keymap) native() *C.GdkKeymap {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkKeymap(p)
}

// Native returns a pointer to the underlying GdkKeymap.
func (v *Keymap) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalKeymap(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &Keymap{obj}, nil
}

func wrapKeymap(obj *glib.Object) *Keymap {
	return &Keymap{obj}
}

// GetKeymap is a wrapper around gdk_keymap_get_for_display().
func (v *Display) GetKeymap() (*Keymap, error) {
	c := C.gdk_keymap_get_for_display(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &Keymap{obj}, nil
}

// TranslateKeyboardState is a wrapper around gdk_keymap_translate_keyboard_state().
func (v *Keymap) TranslateKeyboardState(hardwareKeycode uint, state ModifierType, group int) (bool, *uint, *int, *int, *ModifierType) {

	var cKeyval C.guint
	var keyval *uint
	var cEffectiveGroup, cLevel C.gint
	var effectiveGroup, level *int
	var cConsumedModifiers C.GdkModifierType
	var consumedModifiers *ModifierType

	c := C.gdk_keymap_translate_keyboard_state(
		v.native(),
		C.guint(hardwareKeycode),
		C.GdkModifierType(state),
		C.gint(group),
		&cKeyval,
		&cEffectiveGroup,
		&cLevel,
		&cConsumedModifiers,
	)

	if &cKeyval == nil {
		keyval = nil
	} else {
		*keyval = uint(cKeyval)
	}
	if &cEffectiveGroup == nil {
		effectiveGroup = nil
	} else {
		*effectiveGroup = int(cEffectiveGroup)
	}
	if &cLevel == nil {
		level = nil
	} else {
		*level = int(cLevel)
	}
	if &cConsumedModifiers == nil {
		consumedModifiers = nil
	} else {
		*consumedModifiers = ModifierType(cConsumedModifiers)
	}

	return gobool(c), keyval, effectiveGroup, level, consumedModifiers
}

// HaveBidiLayouts is a wrapper around gdk_keymap_have_bidi_layouts().
func (v *Keymap) HaveBidiLayouts() bool {
	return gobool(C.gdk_keymap_have_bidi_layouts(v.native()))
}

// GetCapsLockState is a wrapper around gdk_keymap_get_caps_lock_state().
func (v *Keymap) GetCapsLockState() bool {
	return gobool(C.gdk_keymap_get_caps_lock_state(v.native()))
}

// GetNumLockState is a wrapper around gdk_keymap_get_num_lock_state().
func (v *Keymap) GetNumLockState() bool {
	return gobool(C.gdk_keymap_get_num_lock_state(v.native()))
}

// GetModifierState is a wrapper around gdk_keymap_get_modifier_state().
func (v *Keymap) GetModifierState() uint {
	return uint(C.gdk_keymap_get_modifier_state(v.native()))
}

// TODO:
// gdk_keymap_get_default(). deprecated since 3.22
// gdk_keymap_get_direction().
// gdk_keymap_add_virtual_modifiers().
// gdk_keymap_map_virtual_modifiers().
// gdk_keymap_get_modifier_mask().

/*
 * GdkKeymapKey
 */

// TODO:
// gdk_keymap_lookup_key().
// gdk_keymap_get_entries_for_keyval().
// gdk_keymap_get_entries_for_keycode().

/*
 * GDK Keyval
 */

// TODO:
// gdk_keyval_name().

// KeyvalFromName() is a wrapper around gdk_keyval_from_name().
func KeyvalFromName(keyvalName string) uint {
	str := (*C.gchar)(C.CString(keyvalName))
	defer C.free(unsafe.Pointer(str))
	return uint(C.gdk_keyval_from_name(str))
}

// KeyvalConvertCase is a wrapper around gdk_keyval_convert_case().
func KeyvalConvertCase(v uint) (lower, upper uint) {
	var l, u C.guint
	l = 0
	u = 0
	C.gdk_keyval_convert_case(C.guint(v), &l, &u)
	return uint(l), uint(u)
}

// KeyvalIsLower is a wrapper around gdk_keyval_is_lower().
func KeyvalIsLower(v uint) bool {
	return gobool(C.gdk_keyval_is_lower(C.guint(v)))
}

// KeyvalIsUpper is a wrapper around gdk_keyval_is_upper().
func KeyvalIsUpper(v uint) bool {
	return gobool(C.gdk_keyval_is_upper(C.guint(v)))
}

// KeyvalToLower is a wrapper around gdk_keyval_to_lower().
func KeyvalToLower(v uint) uint {
	return uint(C.gdk_keyval_to_lower(C.guint(v)))
}

// KeyvalToUpper is a wrapper around gdk_keyval_to_upper().
func KeyvalToUpper(v uint) uint {
	return uint(C.gdk_keyval_to_upper(C.guint(v)))
}

// KeyvalToUnicode is a wrapper around gdk_keyval_to_unicode().
func KeyvalToUnicode(v uint) rune {
	return rune(C.gdk_keyval_to_unicode(C.guint(v)))
}

// UnicodeToKeyval is a wrapper around gdk_unicode_to_keyval().
func UnicodeToKeyval(v rune) uint {
	return uint(C.gdk_unicode_to_keyval(C.guint32(v)))
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

func (v *DragContext) ListTargets() *glib.List {
	clist := C.gdk_drag_context_list_targets(v.native())
	if clist == nil {
		return nil
	}
	return glib.WrapList(uintptr(unsafe.Pointer(clist)))
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

func EventButtonNew() *EventButton {
	ee := (*C.GdkEvent)(unsafe.Pointer(&C.GdkEventButton{}))
	ev := Event{ee}
	return &EventButton{&ev}
}

// EventButtonNewFromEvent returns an EventButton from an Event.
//
// Using widget.Connect() for a key related signal such as
// "button-press-event" results in a *Event being passed as
// the callback's second argument. The argument is actually a
// *EventButton. EventButtonNewFromEvent provides a means of creating
// an EventKey from the Event.
func EventButtonNewFromEvent(event *Event) *EventButton {
	ee := (*C.GdkEvent)(unsafe.Pointer(event.native()))
	ev := Event{ee}
	return &EventButton{&ev}
}

// Native returns a pointer to the underlying GdkEventButton.
func (v *EventButton) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *EventButton) native() *C.GdkEventButton {
	return (*C.GdkEventButton)(unsafe.Pointer(v.Event.native()))
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

func (v *EventButton) Button() Button {
	c := v.native().button
	return Button(c)
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

func (v *EventButton) Type() EventType {
	c := v.native()._type
	return EventType(c)
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

// EventKeyNewFromEvent returns an EventKey from an Event.
//
// Using widget.Connect() for a key related signal such as
// "key-press-event" results in a *Event being passed as
// the callback's second argument. The argument is actually a
// *EventKey. EventKeyNewFromEvent provides a means of creating
// an EventKey from the Event.
func EventKeyNewFromEvent(event *Event) *EventKey {
	ee := (*C.GdkEvent)(unsafe.Pointer(event.native()))
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

func (v *EventKey) HardwareKeyCode() uint16 {
	c := v.native().hardware_keycode
	return uint16(c)
}

func (v *EventKey) Type() EventType {
	c := v.native()._type
	return EventType(c)
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

func EventMotionNew() *EventMotion {
	ee := (*C.GdkEvent)(unsafe.Pointer(&C.GdkEventMotion{}))
	ev := Event{ee}
	return &EventMotion{&ev}
}

// EventMotionNewFromEvent returns an EventMotion from an Event.
//
// Using widget.Connect() for a key related signal such as
// "button-press-event" results in a *Event being passed as
// the callback's second argument. The argument is actually a
// *EventMotion. EventMotionNewFromEvent provides a means of creating
// an EventKey from the Event.
func EventMotionNewFromEvent(event *Event) *EventMotion {
	ee := (*C.GdkEvent)(unsafe.Pointer(event.native()))
	ev := Event{ee}
	return &EventMotion{&ev}
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

// Time returns the time of the event in milliseconds.
func (v *EventMotion) Time() uint32 {
	c := v.native().time
	return uint32(c)
}

func (v *EventMotion) Type() EventType {
	c := v.native()._type
	return EventType(c)
}

// A bit-mask representing the state of the modifier keys (e.g. Control, Shift
// and Alt) and the pointer buttons. See gdk.ModifierType constants.
func (v *EventMotion) State() ModifierType {
	c := v.native().state
	return ModifierType(c)
}

/*
 * GdkEventCrossing
 */

// EventCrossing is a representation of GDK's GdkEventCrossing.
type EventCrossing struct {
	*Event
}

func EventCrossingNew() *EventCrossing {
	ee := (*C.GdkEvent)(unsafe.Pointer(&C.GdkEventCrossing{}))
	ev := Event{ee}
	return &EventCrossing{&ev}
}

// EventCrossingNewFromEvent returns an EventCrossing from an Event.
//
// Using widget.Connect() for a key related signal such as
// "enter-notify-event" results in a *Event being passed as
// the callback's second argument. The argument is actually a
// *EventCrossing. EventCrossingNewFromEvent provides a means of creating
// an EventCrossing from the Event.
func EventCrossingNewFromEvent(event *Event) *EventCrossing {
	ee := (*C.GdkEvent)(unsafe.Pointer(event.native()))
	ev := Event{ee}
	return &EventCrossing{&ev}
}

// Native returns a pointer to the underlying GdkEventCrossing.
func (v *EventCrossing) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *EventCrossing) native() *C.GdkEventCrossing {
	return (*C.GdkEventCrossing)(unsafe.Pointer(v.Event.native()))
}

func (v *EventCrossing) X() float64 {
	c := v.native().x
	return float64(c)
}

func (v *EventCrossing) Y() float64 {
	c := v.native().y
	return float64(c)
}

// XRoot returns the x coordinate of the pointer relative to the root of the screen.
func (v *EventCrossing) XRoot() float64 {
	c := v.native().x_root
	return float64(c)
}

// YRoot returns the y coordinate of the pointer relative to the root of the screen.
func (v *EventCrossing) YRoot() float64 {
	c := v.native().y_root
	return float64(c)
}

func (v *EventCrossing) State() uint {
	c := v.native().state
	return uint(c)
}

// Time returns the time of the event in milliseconds.
func (v *EventCrossing) Time() uint32 {
	c := v.native().time
	return uint32(c)
}

func (v *EventCrossing) Type() EventType {
	c := v.native()._type
	return EventType(c)
}

func (v *EventCrossing) MotionVal() (float64, float64) {
	x := v.native().x
	y := v.native().y
	return float64(x), float64(y)
}

func (v *EventCrossing) MotionValRoot() (float64, float64) {
	x := v.native().x_root
	y := v.native().y_root
	return float64(x), float64(y)
}

func (v *EventCrossing) Mode() CrossingMode {
	c := v.native().mode
	return CrossingMode(c)
}

func (v *EventCrossing) Detail() NotifyType {
	c := v.native().detail
	return NotifyType(c)
}

func (v *EventCrossing) Focus() bool {
	c := v.native().focus
	return gobool(c)
}

/*
 * GdkEventScroll
 */

// EventScroll is a representation of GDK's GdkEventScroll.
type EventScroll struct {
	*Event
}

func EventScrollNew() *EventScroll {
	ee := (*C.GdkEvent)(unsafe.Pointer(&C.GdkEventScroll{}))
	ev := Event{ee}
	return &EventScroll{&ev}
}

// EventScrollNewFromEvent returns an EventScroll from an Event.
//
// Using widget.Connect() for a key related signal such as
// "button-press-event" results in a *Event being passed as
// the callback's second argument. The argument is actually a
// *EventScroll. EventScrollNewFromEvent provides a means of creating
// an EventKey from the Event.
func EventScrollNewFromEvent(event *Event) *EventScroll {
	ee := (*C.GdkEvent)(unsafe.Pointer(event.native()))
	ev := Event{ee}
	return &EventScroll{&ev}
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

func (v *EventScroll) Type() EventType {
	c := v.native()._type
	return EventType(c)
}

func (v *EventScroll) Direction() ScrollDirection {
	c := v.native().direction
	return ScrollDirection(c)
}

// A bit-mask representing the state of the modifier keys (e.g. Control, Shift
// and Alt) and the pointer buttons. See gdk.ModifierType constants.
func (v *EventScroll) State() ModifierType {
	c := v.native().state
	return ModifierType(c)
}

/*
 * GdkEventWindowState
 */

// EventWindowState is a representation of GDK's GdkEventWindowState.
type EventWindowState struct {
	*Event
}

func EventWindowStateNew() *EventWindowState {
	ee := (*C.GdkEvent)(unsafe.Pointer(&C.GdkEventWindowState{}))
	ev := Event{ee}
	return &EventWindowState{&ev}
}

// EventWindowStateNewFromEvent returns an EventWindowState from an Event.
//
// Using widget.Connect() for the
// "window-state-event" signal results in a *Event being passed as
// the callback's second argument. The argument is actually a
// *EventWindowState. EventWindowStateNewFromEvent provides a means of creating
// an EventWindowState from the Event.
func EventWindowStateNewFromEvent(event *Event) *EventWindowState {
	ee := (*C.GdkEvent)(unsafe.Pointer(event.native()))
	ev := Event{ee}
	return &EventWindowState{&ev}
}

// Native returns a pointer to the underlying GdkEventWindowState.
func (v *EventWindowState) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *EventWindowState) native() *C.GdkEventWindowState {
	return (*C.GdkEventWindowState)(unsafe.Pointer(v.Event.native()))
}

func (v *EventWindowState) Type() EventType {
	c := v.native()._type
	return EventType(c)
}

func (v *EventWindowState) ChangedMask() WindowState {
	c := v.native().changed_mask
	return WindowState(c)
}

func (v *EventWindowState) NewWindowState() WindowState {
	c := v.native().new_window_state
	return WindowState(c)
}

/*
 * GdkEventConfigure
 */

// EventConfigure is a representation of GDK's GdkEventConfigure.
type EventConfigure struct {
	*Event
}

func EventConfigureNew() *EventConfigure {
	ee := (*C.GdkEvent)(unsafe.Pointer(&C.GdkEventConfigure{}))
	ev := Event{ee}
	return &EventConfigure{&ev}
}

// EventConfigureNewFromEvent returns an EventConfigure from an Event.
//
// Using widget.Connect() for the
// "configure-event" signal results in a *Event being passed as
// the callback's second argument. The argument is actually a
// *EventConfigure. EventConfigureNewFromEvent provides a means of creating
// an EventConfigure from the Event.
func EventConfigureNewFromEvent(event *Event) *EventConfigure {
	ee := (*C.GdkEvent)(unsafe.Pointer(event.native()))
	ev := Event{ee}
	return &EventConfigure{&ev}
}

// Native returns a pointer to the underlying GdkEventConfigure.
func (v *EventConfigure) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *EventConfigure) native() *C.GdkEventConfigure {
	return (*C.GdkEventConfigure)(unsafe.Pointer(v.Event.native()))
}

func (v *EventConfigure) Type() EventType {
	c := v.native()._type
	return EventType(c)
}

func (v *EventConfigure) X() int {
	c := v.native().x
	return int(c)
}

func (v *EventConfigure) Y() int {
	c := v.native().y
	return int(c)
}

func (v *EventConfigure) Width() int {
	c := v.native().width
	return int(c)
}

func (v *EventConfigure) Height() int {
	c := v.native().height
	return int(c)
}

/*
 * GdkGravity
 */

type Gravity int

const (
	GDK_GRAVITY_NORTH_WEST = C.GDK_GRAVITY_NORTH_WEST
	GDK_GRAVITY_NORTH      = C.GDK_GRAVITY_NORTH
	GDK_GRAVITY_NORTH_EAST = C.GDK_GRAVITY_NORTH_EAST
	GDK_GRAVITY_WEST       = C.GDK_GRAVITY_WEST
	GDK_GRAVITY_CENTER     = C.GDK_GRAVITY_CENTER
	GDK_GRAVITY_EAST       = C.GDK_GRAVITY_EAST
	GDK_GRAVITY_SOUTH_WEST = C.GDK_GRAVITY_SOUTH_WEST
	GDK_GRAVITY_SOUTH      = C.GDK_GRAVITY_SOUTH
	GDK_GRAVITY_SOUTH_EAST = C.GDK_GRAVITY_SOUTH_EAST
	GDK_GRAVITY_STATIC     = C.GDK_GRAVITY_STATIC
)

func marshalGravity(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Gravity(c), nil
}

/*
 * GdkRGBA
 */

type RGBA struct {
	rgba *C.GdkRGBA
}

func marshalRGBA(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRGBA(unsafe.Pointer(c)), nil
}

func WrapRGBA(p unsafe.Pointer) *RGBA {
	return wrapRGBA((*C.GdkRGBA)(p))
}

func wrapRGBA(cRgba *C.GdkRGBA) *RGBA {
	return &RGBA{cRgba}
}

func NewRGBA(values ...float64) *RGBA {

	cRgba := new(C.GdkRGBA)
	if len(values) > 0 {
		cRgba.red = C.gdouble(values[0])
	}
	if len(values) > 1 {
		cRgba.green = C.gdouble(values[1])
	}
	if len(values) > 2 {
		cRgba.blue = C.gdouble(values[2])
	}
	if len(values) > 3 {
		cRgba.alpha = C.gdouble(values[3])
	}
	return wrapRGBA(cRgba)
}

func (c *RGBA) Floats() []float64 {
	return []float64{float64(c.rgba.red), float64(c.rgba.green), float64(c.rgba.blue), float64(c.rgba.alpha)}
}

// SetColors sets all colors values in the RGBA.
func (c *RGBA) SetColors(r, g, b, a float64) {
	c.rgba.red = C.gdouble(r)
	c.rgba.green = C.gdouble(g)
	c.rgba.blue = C.gdouble(b)
	c.rgba.alpha = C.gdouble(a)
}

func (c *RGBA) Native() uintptr {
	return uintptr(unsafe.Pointer(c.rgba))
}

// Parse is a representation of gdk_rgba_parse().
func (c *RGBA) Parse(spec string) bool {
	cstr := (*C.gchar)(C.CString(spec))
	defer C.free(unsafe.Pointer(cstr))

	return gobool(C.gdk_rgba_parse(c.rgba, cstr))
}

// String is a representation of gdk_rgba_to_string().
func (c *RGBA) String() string {
	return C.GoString((*C.char)(C.gdk_rgba_to_string(c.rgba)))
}

// free is a representation of gdk_rgba_free().
func (c *RGBA) free() {
	C.gdk_rgba_free(c.rgba)
}

// Copy is a representation of gdk_rgba_copy().
func (c *RGBA) Copy() (*RGBA, error) {
	cRgba := C.gdk_rgba_copy(c.rgba)

	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapRGBA(cRgba)

	runtime.SetFinalizer(obj, (*RGBA).free)
	return obj, nil
}

// Equal is a representation of gdk_rgba_equal().
func (c *RGBA) Equal(rgba *RGBA) bool {
	return gobool(C.gdk_rgba_equal(
		C.gconstpointer(c.rgba),
		C.gconstpointer(rgba.rgba)))
}

// Hash is a representation of gdk_rgba_hash().
func (c *RGBA) Hash() uint {
	return uint(C.gdk_rgba_hash(C.gconstpointer(c.rgba)))
}

/*
 * GdkRectangle
 */

// Rectangle is a representation of GDK's GdkRectangle type.
type Rectangle struct {
	GdkRectangle C.GdkRectangle
}

func WrapRectangle(p uintptr) *Rectangle {
	return wrapRectangle((*C.GdkRectangle)(unsafe.Pointer(p)))
}

func wrapRectangle(obj *C.GdkRectangle) *Rectangle {
	if obj == nil {
		return nil
	}
	return &Rectangle{*obj}
}

// Native() returns a pointer to the underlying GdkRectangle.
func (r *Rectangle) native() *C.GdkRectangle {
	return &r.GdkRectangle
}

// GetX returns x field of the underlying GdkRectangle.
func (r *Rectangle) GetX() int {
	return int(r.native().x)
}

// SetX sets x field of the underlying GdkRectangle.
func (r *Rectangle) SetX(x int) {
	r.native().x = C.int(x)
}

// GetY returns y field of the underlying GdkRectangle.
func (r *Rectangle) GetY() int {
	return int(r.native().y)
}

// SetY sets y field of the underlying GdkRectangle.
func (r *Rectangle) SetY(y int) {
	r.native().y = C.int(y)
}

// GetWidth returns width field of the underlying GdkRectangle.
func (r *Rectangle) GetWidth() int {
	return int(r.native().width)
}

// SetWidth sets width field of the underlying GdkRectangle.
func (r *Rectangle) SetWidth(width int) {
	r.native().width = C.int(width)
}

// GetHeight returns height field of the underlying GdkRectangle.
func (r *Rectangle) GetHeight() int {
	return int(r.native().height)
}

// SetHeight sets height field of the underlying GdkRectangle.
func (r *Rectangle) SetHeight(height int) {
	r.native().height = C.int(height)
}

/*
 * GdkGeometry
 */

type Geometry struct {
	GdkGeometry C.GdkGeometry
}

func WrapGeometry(p uintptr) *Geometry {
	return wrapGeometry((*C.GdkGeometry)(unsafe.Pointer(p)))
}

func wrapGeometry(obj *C.GdkGeometry) *Geometry {
	if obj == nil {
		return nil
	}
	return &Geometry{*obj}
}

// native returns a pointer to the underlying GdkGeometry.
func (r *Geometry) native() *C.GdkGeometry {
	return &r.GdkGeometry
}

// GetMinWidth returns min_width field of the underlying GdkGeometry.
func (r *Geometry) GetMinWidth() int {
	return int(r.native().min_width)
}

// SetMinWidth sets min_width field of the underlying GdkGeometry.
func (r *Geometry) SetMinWidth(minWidth int) {
	r.native().min_width = C.gint(minWidth)
}

// GetMinHeight returns min_height field of the underlying GdkGeometry.
func (r *Geometry) GetMinHeight() int {
	return int(r.native().min_height)
}

// SetMinHeight sets min_height field of the underlying GdkGeometry.
func (r *Geometry) SetMinHeight(minHeight int) {
	r.native().min_height = C.gint(minHeight)
}

// GetMaxWidth returns max_width field of the underlying GdkGeometry.
func (r *Geometry) GetMaxWidth() int {
	return int(r.native().max_width)
}

// SetMaxWidth sets max_width field of the underlying GdkGeometry.
func (r *Geometry) SetMaxWidth(maxWidth int) {
	r.native().max_width = C.gint(maxWidth)
}

// GetMaxHeight returns max_height field of the underlying GdkGeometry.
func (r *Geometry) GetMaxHeight() int {
	return int(r.native().max_height)
}

// SetMaxHeight sets max_height field of the underlying GdkGeometry.
func (r *Geometry) SetMaxHeight(maxHeight int) {
	r.native().max_height = C.gint(maxHeight)
}

// GetBaseWidth returns base_width field of the underlying GdkGeometry.
func (r *Geometry) GetBaseWidth() int {
	return int(r.native().base_width)
}

// SetBaseWidth sets base_width field of the underlying GdkGeometry.
func (r *Geometry) SetBaseWidth(baseWidth int) {
	r.native().base_width = C.gint(baseWidth)
}

// GetBaseHeight returns base_height field of the underlying GdkGeometry.
func (r *Geometry) GetBaseHeight() int {
	return int(r.native().base_height)
}

// SetBaseHeight sets base_height field of the underlying GdkGeometry.
func (r *Geometry) SetBaseHeight(baseHeight int) {
	r.native().base_height = C.gint(baseHeight)
}

// GetWidthInc returns width_inc field of the underlying GdkGeometry.
func (r *Geometry) GetWidthInc() int {
	return int(r.native().width_inc)
}

// SetWidthInc sets width_inc field of the underlying GdkGeometry.
func (r *Geometry) SetWidthInc(widthInc int) {
	r.native().width_inc = C.gint(widthInc)
}

// GetHeightInc returns height_inc field of the underlying GdkGeometry.
func (r *Geometry) GetHeightInc() int {
	return int(r.native().height_inc)
}

// SetHeightInc sets height_inc field of the underlying GdkGeometry.
func (r *Geometry) SetHeightInc(heightInc int) {
	r.native().height_inc = C.gint(heightInc)
}

// GetMinAspect returns min_aspect field of the underlying GdkGeometry.
func (r *Geometry) GetMinAspect() float64 {
	return float64(r.native().min_aspect)
}

// SetMinAspect sets min_aspect field of the underlying GdkGeometry.
func (r *Geometry) SetMinAspect(minAspect float64) {
	r.native().min_aspect = C.gdouble(minAspect)
}

// GetMaxAspect returns max_aspect field of the underlying GdkGeometry.
func (r *Geometry) GetMaxAspect() float64 {
	return float64(r.native().max_aspect)
}

// SetMaxAspect sets max_aspect field of the underlying GdkGeometry.
func (r *Geometry) SetMaxAspect(maxAspect float64) {
	r.native().max_aspect = C.gdouble(maxAspect)
}

// GetWinGravity returns win_gravity field of the underlying GdkGeometry.
func (r *Geometry) GetWinGravity() Gravity {
	return Gravity(r.native().win_gravity)
}

// SetWinGravity sets win_gravity field of the underlying GdkGeometry.
func (r *Geometry) SetWinGravity(winGravity Gravity) {
	r.native().win_gravity = C.GdkGravity(winGravity)
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

// TODO:
// gdk_visual_get_blue_pixel_details().
// gdk_visual_get_depth().
// gdk_visual_get_green_pixel_details().
// gdk_visual_get_red_pixel_details().
// gdk_visual_get_visual_type().
// gdk_visual_get_screen().

/*
 * GdkWindow
 */

// Window is a representation of GDK's GdkWindow.
type Window struct {
	*glib.Object
}

// SetCursor is a wrapper around gdk_window_set_cursor().
func (v *Window) SetCursor(cursor *Cursor) {
	C.gdk_window_set_cursor(v.native(), cursor.native())
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

// WindowGetWidth is a wrapper around gdk_window_get_width()
func (v *Window) WindowGetWidth() (width int) {
	return int(C.gdk_window_get_width(v.native()))
}

// WindowGetHeight is a wrapper around gdk_window_get_height()
func (v *Window) WindowGetHeight() (height int) {
	return int(C.gdk_window_get_height(v.native()))
}

// CreateSimilarSurface is a wrapper around gdk_window_create_similar_surface().
func (v *Window) CreateSimilarSurface(content cairo.Content, w, h int) (*cairo.Surface, error) {
	surface := C.gdk_window_create_similar_surface(v.native(), C.cairo_content_t(content), C.gint(w), C.gint(h))

	status := cairo.Status(C.cairo_surface_status(surface))
	if status != cairo.STATUS_SUCCESS {
		return nil, cairo.ErrorStatus(status)
	}

	return cairo.NewSurface(uintptr(unsafe.Pointer(surface)), false), nil
}

//PixbufGetFromWindow is a wrapper around gdk_pixbuf_get_from_window()
func (v *Window) PixbufGetFromWindow(x, y, w, h int) (*Pixbuf, error) {
	c := C.gdk_pixbuf_get_from_window(v.native(), C.gint(x), C.gint(y), C.gint(w), C.gint(h))
	if c == nil {
		return nil, nilPtrErr
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	//obj.Ref()
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })
	return p, nil
}

// GetDevicePosition is a wrapper around gdk_window_get_device_position()
func (v *Window) GetDevicePosition(d *Device) (*Window, int, int, ModifierType) {
	var x C.gint
	var y C.gint
	var mt C.GdkModifierType
	underneathWindow := C.gdk_window_get_device_position(v.native(), d.native(), &x, &y, &mt)
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(underneathWindow))}
	rw := &Window{obj}
	runtime.SetFinalizer(rw, func(_ interface{}) { obj.Unref() })
	return rw, int(x), int(y), ModifierType(mt)
}

func PixbufGetFromSurface(surface *cairo.Surface, src_x, src_y, width, height int) (*Pixbuf, error) {
	c := C.gdk_pixbuf_get_from_surface((*C.cairo_surface_t)(unsafe.Pointer(surface.Native())), C.gint(src_x), C.gint(src_y), C.gint(width), C.gint(height))
	if c == nil {
		return nil, nilPtrErr
	}

	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	p := &Pixbuf{obj}
	//obj.Ref()
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })
	return p, nil
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

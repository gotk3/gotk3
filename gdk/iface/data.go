package iface

// DragAction is a representation of GDK's GdkDragAction.
type DragAction int

var (
	ACTION_DEFAULT DragAction
	ACTION_COPY    DragAction
	ACTION_MOVE    DragAction
	ACTION_LINK    DragAction
	ACTION_PRIVATE DragAction
	ACTION_ASK     DragAction
)

// Colorspace is a representation of GDK's GdkColorspace.
type Colorspace int

var (
	COLORSPACE_RGB Colorspace
)

// InterpType is a representation of GDK's GdkInterpType.
type InterpType int

var (
	INTERP_NEAREST  InterpType
	INTERP_TILES    InterpType
	INTERP_BILINEAR InterpType
	INTERP_HYPER    InterpType
)

// PixbufRotation is a representation of GDK's GdkPixbufRotation.
type PixbufRotation int

var (
	PIXBUF_ROTATE_NONE             PixbufRotation
	PIXBUF_ROTATE_COUNTERCLOCKWISE PixbufRotation
	PIXBUF_ROTATE_UPSIDEDOWN       PixbufRotation
	PIXBUF_ROTATE_CLOCKWISE        PixbufRotation
)

// ModifierType is a representation of GDK's GdkModifierType.
type ModifierType uint

var (
	GDK_SHIFT_MASK    ModifierType
	GDK_LOCK_MASK     ModifierType
	GDK_CONTROL_MASK  ModifierType
	GDK_MOD1_MASK     ModifierType
	GDK_MOD2_MASK     ModifierType
	GDK_MOD3_MASK     ModifierType
	GDK_MOD4_MASK     ModifierType
	GDK_MOD5_MASK     ModifierType
	GDK_BUTTON1_MASK  ModifierType
	GDK_BUTTON2_MASK  ModifierType
	GDK_BUTTON3_MASK  ModifierType
	GDK_BUTTON4_MASK  ModifierType
	GDK_BUTTON5_MASK  ModifierType
	GDK_SUPER_MASK    ModifierType
	GDK_HYPER_MASK    ModifierType
	GDK_META_MASK     ModifierType
	GDK_RELEASE_MASK  ModifierType
	GDK_MODIFIER_MASK ModifierType
)

// PixbufAlphaMode is a representation of GDK's GdkPixbufAlphaMode.
type PixbufAlphaMode int

var (
	GDK_PIXBUF_ALPHA_BILEVEL PixbufAlphaMode
	GDK_PIXBUF_ALPHA_FULL    PixbufAlphaMode
)

// Selections
var (
	SELECTION_PRIMARY       Atom
	SELECTION_SECONDARY     Atom
	SELECTION_CLIPBOARD     Atom
	TARGET_BITMAP           Atom
	TARGET_COLORMAP         Atom
	TARGET_DRAWABLE         Atom
	TARGET_PIXMAP           Atom
	TARGET_STRING           Atom
	SELECTION_TYPE_ATOM     Atom
	SELECTION_TYPE_BITMAP   Atom
	SELECTION_TYPE_COLORMAP Atom
	SELECTION_TYPE_DRAWABLE Atom
	SELECTION_TYPE_INTEGER  Atom
	SELECTION_TYPE_PIXMAP   Atom
	SELECTION_TYPE_WINDOW   Atom
	SELECTION_TYPE_STRING   Atom
)

// added by terrak
// EventMask is a representation of GDK's GdkEventMask.
type EventMask int

var (
	EXPOSURE_MASK            EventMask
	POINTER_MOTION_MASK      EventMask
	POINTER_MOTION_HINT_MASK EventMask
	BUTTON_MOTION_MASK       EventMask
	BUTTON1_MOTION_MASK      EventMask
	BUTTON2_MOTION_MASK      EventMask
	BUTTON3_MOTION_MASK      EventMask
	BUTTON_PRESS_MASK        EventMask
	BUTTON_RELEASE_MASK      EventMask
	KEY_PRESS_MASK           EventMask
	KEY_RELEASE_MASK         EventMask
	ENTER_NOTIFY_MASK        EventMask
	LEAVE_NOTIFY_MASK        EventMask
	FOCUS_CHANGE_MASK        EventMask
	STRUCTURE_MASK           EventMask
	PROPERTY_CHANGE_MASK     EventMask
	VISIBILITY_NOTIFY_MASK   EventMask
	PROXIMITY_IN_MASK        EventMask
	PROXIMITY_OUT_MASK       EventMask
	SUBSTRUCTURE_MASK        EventMask
	SCROLL_MASK              EventMask
	TOUCH_MASK               EventMask
	SMOOTH_SCROLL_MASK       EventMask
	ALL_EVENTS_MASK          EventMask
)

// ScrollDirection is a representation of GDK's GdkScrollDirection
type ScrollDirection int

var (
	SCROLL_UP     ScrollDirection
	SCROLL_DOWN   ScrollDirection
	SCROLL_LEFT   ScrollDirection
	SCROLL_RIGHT  ScrollDirection
	SCROLL_SMOOTH ScrollDirection
)

// CURRENT_TIME is a representation of GDK_CURRENT_TIME

var CURRENT_TIME int

// GrabStatus is a representation of GdkGrabStatus

type GrabStatus int

var (
	GRAB_SUCCESS         GrabStatus
	GRAB_ALREADY_GRABBED GrabStatus
	GRAB_INVALID_TIME    GrabStatus
	GRAB_FROZEN          GrabStatus
	// Only exists since 3.16
	// GRAB_FAILED GrabStatus = C.GDK_GRAB_FAILED
	GRAB_FAILED GrabStatus
)

// GrabOwnership is a representation of GdkGrabOwnership

type GrabOwnership int

var (
	OWNERSHIP_NONE        GrabOwnership
	OWNERSHIP_WINDOW      GrabOwnership
	OWNERSHIP_APPLICATION GrabOwnership
)

// DeviceType is a representation of GdkDeviceType

type DeviceType int

var (
	DEVICE_TYPE_MASTER   DeviceType
	DEVICE_TYPE_SLAVE    DeviceType
	DEVICE_TYPE_FLOATING DeviceType
)

// Atom is a representation of GDK's GdkAtom.
type Atom uintptr

// EventType is a representation of GDK's GdkEventType.
// Do not confuse these event types with the signals that GTK+ widgets emit
type EventType int

var (
	EVENT_NOTHING             EventType
	EVENT_DELETE              EventType
	EVENT_DESTROY             EventType
	EVENT_EXPOSE              EventType
	EVENT_MOTION_NOTIFY       EventType
	EVENT_BUTTON_PRESS        EventType
	EVENT_2BUTTON_PRESS       EventType
	EVENT_DOUBLE_BUTTON_PRESS EventType
	EVENT_3BUTTON_PRESS       EventType
	EVENT_TRIPLE_BUTTON_PRESS EventType
	EVENT_BUTTON_RELEASE      EventType
	EVENT_KEY_PRESS           EventType
	EVENT_KEY_RELEASE         EventType
	EVENT_LEAVE_NOTIFY        EventType
	EVENT_FOCUS_CHANGE        EventType
	EVENT_CONFIGURE           EventType
	EVENT_MAP                 EventType
	EVENT_UNMAP               EventType
	EVENT_PROPERTY_NOTIFY     EventType
	EVENT_SELECTION_CLEAR     EventType
	EVENT_SELECTION_REQUEST   EventType
	EVENT_SELECTION_NOTIFY    EventType
	EVENT_PROXIMITY_IN        EventType
	EVENT_PROXIMITY_OUT       EventType
	EVENT_DRAG_ENTER          EventType
	EVENT_DRAG_LEAVE          EventType
	EVENT_DRAG_MOTION         EventType
	EVENT_DRAG_STATUS         EventType
	EVENT_DROP_START          EventType
	EVENT_DROP_FINISHED       EventType
	EVENT_CLIENT_EVENT        EventType
	EVENT_VISIBILITY_NOTIFY   EventType
	EVENT_SCROLL              EventType
	EVENT_WINDOW_STATE        EventType
	EVENT_SETTING             EventType
	EVENT_OWNER_CHANGE        EventType
	EVENT_GRAB_BROKEN         EventType
	EVENT_DAMAGE              EventType
	EVENT_TOUCH_BEGIN         EventType
	EVENT_TOUCH_UPDATE        EventType
	EVENT_TOUCH_END           EventType
	EVENT_TOUCH_CANCEL        EventType
	EVENT_LAST                EventType
)

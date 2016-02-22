package gtk

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
)

type Widget interface {
	glib.InitiallyUnowned

	Activate() bool
	AddAccelerator(string, AccelGroup, uint, gdk.ModifierType, AccelFlags)
	AddEvents(int)
	CanActivateAccel(uint) bool
	Destroy()
	Event(gdk.Event) bool
	GetAllocatedHeight() int
	GetAllocatedWidth() int
	GetAllocation() Allocation
	GetAppPaintable() bool
	GetCanDefault() bool
	GetCanFocus() bool
	GetDeviceEnabled(gdk.Device) bool
	GetEvents() int
	GetHAlign() Align
	GetHExpand() bool
	GetHasWindow() bool
	GetMapped() bool
	GetMarginBottom() int
	GetMarginTop() int
	GetName() (string, error)
	GetNoShowAll() bool
	GetParent() (Widget, error)
	GetParentWindow() (gdk.Window, error)
	GetRealized() bool
	GetSensitive() bool
	GetSizeRequest() (int, int)
	GetStyleContext() (StyleContext, error)
	GetTooltipText() (string, error)
	GetToplevel() (Widget, error)
	GetVAlign() Align
	GetVExpand() bool
	GetVisible() bool
	GetWindow() (gdk.Window, error)
	GrabDefault()
	GrabFocus()
	HasDefault() bool
	HasFocus() bool
	HasGrab() bool
	HasVisibleFocus() bool
	Hide()
	HideOnDelete()
	InDestruction() bool
	IsDrawable() bool
	IsFocus() bool
	IsSensitive() bool
	IsToplevel() bool
	Map()
	QueueDraw()
	QueueDrawArea(int, int, int, int)
	RemoveAccelerator(AccelGroup, uint, gdk.ModifierType) bool
	ResetStyle()
	SetAccelPath2(string, AccelGroup)
	SetAllocation(Allocation)
	SetAppPaintable(bool)
	SetCanDefault(bool)
	SetCanFocus(bool)
	SetDeviceEnabled(gdk.Device, bool)
	SetEvents(int)
	SetHAlign(Align)
	SetHExpand(bool)
	SetHasWindow(bool)
	SetMapped(bool)
	SetMarginBottom(int)
	SetMarginTop(int)
	SetName(string)
	SetNoShowAll(bool)
	SetParent(Widget)
	SetParentWindow(gdk.Window)
	SetRealized(bool)
	SetSensitive(bool)
	SetSizeRequest(int, int)
	SetStateFlags(StateFlags, bool)
	SetTooltipText(string)
	SetVAlign(Align)
	SetVExpand(bool)
	SetVisible(bool)
	SetVisual(gdk.Visual)
	Show()
	ShowAll()
	ShowNow()
	SizeAllocate(Allocation)
	TranslateCoordinates(Widget, int, int) (int, int, error)
	Unmap()
	Unparent()
} // end of Widget

func AssertWidget(_ Widget) {}

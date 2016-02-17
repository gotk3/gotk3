package gdk

import "github.com/gotk3/gotk3/glib"

type Display interface {
	glib.Object

	Beep()
	Close()
	DeviceIsGrabbed(Device) bool
	Flush()
	GetAppLaunchContext()
	GetDefaultCursorSize() uint
	GetDefaultGroup() (Window, error)
	GetDefaultScreen() (Screen, error)
	GetDeviceManager() (DeviceManager, error)
	GetEvent() (Event, error)
	GetMaximalCursorSize() (uint, uint)
	GetName() (string, error)
	GetScreen(int) (Screen, error)
	HasPending() bool
	IsClosed() bool
	NotifyStartupComplete(string)
	PeekEvent() (Event, error)
	PutEvent(Event)
	RequestSelectionNotification(Atom) bool
	SetDoubleClickDistance(uint)
	SetDoubleClickTime(uint)
	StoreClipboard(Window, uint32, ...Atom)
	SupportsClipboardPersistence() bool
	SupportsColorCursor() bool
	SupportsCursorAlpha() bool
	SupportsInputShapes() bool
	SupportsSelectionNotification() bool
	SupportsShapes() bool
	Sync()
} // end of Display

func AssertDisplay(_ Display) {}

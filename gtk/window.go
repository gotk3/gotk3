package gtk

import "github.com/gotk3/gotk3/gdk"

type Window interface {
	Bin

	ActivateDefault() bool
	ActivateFocus() bool
	AddAccelGroup(AccelGroup)
	Deiconify()
	Fullscreen()
	GetAcceptFocus() bool
	GetApplication() (Application, error)
	GetAttachedTo() (Widget, error)
	GetDecorated() bool
	GetDefaultSize() (int, int)
	GetDefaultWidget() Widget
	GetDeletable() bool
	GetDestroyWithParent() bool
	GetFocus() (Widget, error)
	GetFocusOnMap() bool
	GetFocusVisible() bool
	GetHideTitlebarWhenMaximized() bool
	GetIcon() (gdk.Pixbuf, error)
	GetIconName() (string, error)
	GetMnemonicsVisible() bool
	GetModal() bool
	GetPosition() (int, int)
	GetResizable() bool
	GetRole() (string, error)
	GetScreen() (gdk.Screen, error)
	GetSize() (int, int)
	GetSkipPagerHint() bool
	GetSkipTaskbarHint() bool
	GetTitle() (string, error)
	GetTransientFor() (Window, error)
	GetUrgencyHint() bool
	HasGroup() bool
	HasToplevelFocus() bool
	Iconify()
	IsActive() bool
	Maximize()
	Move(int, int)
	Present()
	PresentWithTime(uint32)
	RemoveAccelGroup(AccelGroup)
	Resize(int, int)
	ResizeToGeometry(int, int)
	SetAcceptFocus(bool)
	SetApplication(Application)
	SetDecorated(bool)
	SetDefault(Widget)
	SetDefaultGeometry(int, int)
	SetDefaultSize(int, int)
	SetDeletable(bool)
	SetDestroyWithParent(bool)
	SetFocus(Widget)
	SetFocusOnMap(bool)
	SetFocusVisible(bool)
	SetHideTitlebarWhenMaximized(bool)
	SetIcon(gdk.Pixbuf)
	SetIconFromFile(string) error
	SetIconName(string)
	SetKeepAbove(bool)
	SetKeepBelow(bool)
	SetMnemonicsVisible(bool)
	SetModal(bool)
	SetPosition(WindowPosition)
	SetResizable(bool)
	SetRole(string)
	SetSkipPagerHint(bool)
	SetSkipTaskbarHint(bool)
	SetStartupID(string)
	SetTitle(string)
	SetTransientFor(Window)
	SetUrgencyHint(bool)
	SetWMClass(string, string)
	Stick()
	Unfullscreen()
	Unmaximize()
	Unstick()
} // end of Window

func AssertWindow(_ Window) {}

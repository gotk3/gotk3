package gtk

import "github.com/gotk3/gotk3/glib"

type Application interface {
	glib.Application

	AddWindow(Window)
	GetActiveWindow() Window
	GetAppMenu() glib.MenuModel
	GetMenubar() glib.MenuModel
	GetWindowByID(uint) Window
	GetWindows() glib.List
	Inhibited(Window, ApplicationInhibitFlags, string) uint
	IsInhibited(ApplicationInhibitFlags) bool
	RemoveWindow(Window)
	SetAppMenu(glib.MenuModel)
	SetMenubar(glib.MenuModel)
	Uninhibit(uint)
} // end of Application

func AssertApplication(_ Application) {}

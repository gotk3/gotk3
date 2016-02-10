package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type Application interface {
    glib_iface.Application

    AddWindow(Window)
    GetActiveWindow() Window
    GetAppMenu() glib_iface.MenuModel
    GetMenubar() glib_iface.MenuModel
    GetWindowByID(uint) Window
    GetWindows() glib_iface.List
    Inhibited(Window, ApplicationInhibitFlags, string) uint
    IsInhibited(ApplicationInhibitFlags) bool
    RemoveWindow(Window)
    SetAppMenu(glib_iface.MenuModel)
    SetMenubar(glib_iface.MenuModel)
    Uninhibit(uint)
} // end of Application

func AssertApplication(_ Application) {}

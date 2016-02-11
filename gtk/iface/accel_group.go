package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"
import gdk_iface "github.com/gotk3/gotk3/gdk/iface"

type AccelGroup interface {
	glib_iface.Object

	Activate(glib_iface.Quark, glib_iface.Object, uint, gdk_iface.ModifierType) bool

	Connect2(uint, gdk_iface.ModifierType, AccelFlags, interface{})
	ConnectByPath(string, interface{})
	Disconnect(interface{})
	DisconnectKey(uint, gdk_iface.ModifierType)
	GetModifierMask() gdk_iface.ModifierType
	IsLocked() bool
	Lock()
	Unlock()
} // end of AccelGroup

func AssertAccelGroup(_ AccelGroup) {}

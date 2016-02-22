package gtk

import "github.com/gotk3/gotk3/glib"

import "github.com/gotk3/gotk3/gdk"

type AccelGroup interface {
	glib.Object

	Activate(glib.Quark, glib.Object, uint, gdk.ModifierType) bool

	Connect2(uint, gdk.ModifierType, AccelFlags, interface{})
	ConnectByPath(string, interface{})
	Disconnect(interface{})
	DisconnectKey(uint, gdk.ModifierType)
	GetModifierMask() gdk.ModifierType
	IsLocked() bool
	Lock()
	Unlock()
} // end of AccelGroup

func AssertAccelGroup(_ AccelGroup) {}

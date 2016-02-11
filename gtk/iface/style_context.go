package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"
import gdk_iface "github.com/gotk3/gotk3/gdk/iface"

type StyleContext interface {
	glib_iface.Object

	AddClass(string)
	AddProvider(StyleProvider, uint)
	GetColor(StateFlags) gdk_iface.RGBA
	GetParent() (StyleContext, error)
	GetProperty2(string, StateFlags) (interface{}, error)
	GetScreen() (gdk_iface.Screen, error)
	GetState() StateFlags
	GetStyleProperty(string) (interface{}, error)
	HasClass(string) bool
	LookupColor(string) (gdk_iface.RGBA, bool)
	RemoveClass(string)
	RemoveProvider(StyleProvider)
	Restore()
	Save()
	SetParent(StyleContext)
	SetScreen(gdk_iface.Screen)
	SetState(StateFlags)
} // end of StyleContext

func AssertStyleContext(_ StyleContext) {}

package gtk

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
)

type StyleContext interface {
	glib.Object

	AddClass(string)
	AddProvider(StyleProvider, uint)
	GetColor(StateFlags) gdk.RGBA
	GetParent() (StyleContext, error)
	GetProperty2(string, StateFlags) (interface{}, error)
	GetScreen() (gdk.Screen, error)
	GetState() StateFlags
	GetStyleProperty(string) (interface{}, error)
	HasClass(string) bool
	LookupColor(string) (gdk.RGBA, bool)
	RemoveClass(string)
	RemoveProvider(StyleProvider)
	Restore()
	Save()
	SetParent(StyleContext)
	SetScreen(gdk.Screen)
	SetState(StateFlags)
} // end of StyleContext

func AssertStyleContext(_ StyleContext) {}

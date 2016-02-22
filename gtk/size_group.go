package gtk

import "github.com/gotk3/gotk3/glib"

type SizeGroup interface {
	glib.Object

	AddWidget(Widget)
	GetIgnoreHidden() bool
	GetMode() SizeGroupMode
	GetWidgets() glib.SList
	RemoveWidget(Widget)
	SetIgnoreHidden(bool)
	SetMode(SizeGroupMode)
}

func AssertSizeGroup(_ SizeGroup) {}

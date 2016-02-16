package gtk

import "github.com/gotk3/gotk3/glib"

type Scrollable interface {
	glib.Object

	GetHAdjustment() (Adjustment, error)
	GetVAdjustment() (Adjustment, error)
	SetHAdjustment(Adjustment)
	SetVAdjustment(Adjustment)
} // end of Scrollable

func AssertScrollable(_ Scrollable) {}

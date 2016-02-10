package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type Scrollable interface {
    glib_iface.Object

    GetHAdjustment() (Adjustment, error)
    GetVAdjustment() (Adjustment, error)
    SetHAdjustment(Adjustment)
    SetVAdjustment(Adjustment)
} // end of Scrollable

func AssertScrollable(_ Scrollable) {}

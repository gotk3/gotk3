package iface

import gdk_iface "github.com/gotk3/gotk3/gdk/iface"

type Allocation interface {
    gdk_iface.Rectangle
} // end of Allocation

func AssertAllocation(_ Allocation) {}

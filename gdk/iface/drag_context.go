package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type DragContext interface {
    glib_iface.Object

    ListTargets() glib_iface.List
} // end of DragContext

func AssertDragContext(_ DragContext) {}

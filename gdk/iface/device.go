package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type Device interface {
    glib_iface.Object

    Grab(Window, GrabOwnership, bool, EventMask, Cursor, uint32) GrabStatus
    Ungrab(uint32)
} // end of Device

func AssertDevice(_ Device) {}

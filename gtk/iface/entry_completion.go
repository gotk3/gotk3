package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type EntryCompletion interface {
    glib_iface.Object
} // end of EntryCompletion

func AssertEntryCompletion(_ EntryCompletion) {}

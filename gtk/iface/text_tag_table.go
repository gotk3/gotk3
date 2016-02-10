package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type TextTagTable interface {
    glib_iface.Object

    Add(TextTag)
    Lookup(string) (TextTag, error)
    Remove(TextTag)
} // end of TextTagTable

func AssertTextTagTable(_ TextTagTable) {}

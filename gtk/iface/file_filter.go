package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type FileFilter interface {
    glib_iface.Object

    AddPattern(string)
    AddPixbufFormats()
    SetName(string)
} // end of FileFilter

func AssertFileFilter(_ FileFilter) {}

package iface

import gdk_iface "github.com/gotk3/gotk3/gdk/iface"

type Image interface {
    Widget

    Clear()
    GetIconName() (string, IconSize)
    GetPixbuf() gdk_iface.Pixbuf
    GetPixelSize() int
    GetStorageType() ImageType
    SetFromFile(string)
    SetFromIconName(string, IconSize)
    SetFromPixbuf(gdk_iface.Pixbuf)
    SetFromResource(string)
    SetPixelSize(int)
} // end of Image

func AssertImage(_ Image) {}

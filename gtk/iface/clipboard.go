package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"
import gdk_iface "github.com/gotk3/gotk3/gdk/iface"

type Clipboard interface {
    glib_iface.Object

    SetImage(gdk_iface.Pixbuf)
    SetText(string)
    Store()
    WaitForContents(gdk_iface.Atom) (SelectionData, error)
    WaitForImage() (gdk_iface.Pixbuf, error)
    WaitForText() (string, error)
    WaitIsImageAvailable() bool
    WaitIsRichTextAvailable(TextBuffer) bool
    WaitIsTargetAvailable(gdk_iface.Atom) bool
    WaitIsTextAvailable() bool
    WaitIsUrisAvailable() bool
} // end of Clipboard

func AssertClipboard(_ Clipboard) {}

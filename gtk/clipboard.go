package gtk

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
)

type Clipboard interface {
	glib.Object

	SetImage(gdk.Pixbuf)
	SetText(string)
	Store()
	WaitForContents(gdk.Atom) (SelectionData, error)
	WaitForImage() (gdk.Pixbuf, error)
	WaitForText() (string, error)
	WaitIsImageAvailable() bool
	WaitIsRichTextAvailable(TextBuffer) bool
	WaitIsTargetAvailable(gdk.Atom) bool
	WaitIsTextAvailable() bool
	WaitIsUrisAvailable() bool
} // end of Clipboard

func AssertClipboard(_ Clipboard) {}

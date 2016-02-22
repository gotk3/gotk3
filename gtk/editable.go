package gtk

import "github.com/gotk3/gotk3/glib"

type Editable interface {
	glib.Object

	CopyClipboard()
	CutClipboard()
	DeleteSelection()
	DeleteText(int, int)
	GetChars(int, int) string
	GetEditable() bool
	GetPosition() int
	GetSelectionBounds() (int, int, bool)
	InsertText(string, int) int
	PasteClipboard()
	SelectRegion(int, int)
	SetEditable(bool)
	SetPosition(int)
} // end of Editable

func AssertEditable(_ Editable) {}

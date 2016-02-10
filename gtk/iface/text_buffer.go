package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type TextBuffer interface {
	glib_iface.Object

	ApplyTag(TextTag, TextIter, TextIter)
	ApplyTagByName(string, TextIter, TextIter)
	Delete(TextIter, TextIter)
	GetBounds() (TextIter, TextIter)
	GetCharCount() int
	GetEndIter() TextIter
	GetIterAtOffset(int) TextIter
	GetLineCount() int
	GetModified() bool
	GetStartIter() TextIter
	GetTagTable() (TextTagTable, error)
	GetText(TextIter, TextIter, bool) (string, error)
	Insert(TextIter, string)
	InsertAtCursor(string)
	RemoveTag(TextTag, TextIter, TextIter)
	SetModified(bool)
	SetText(string)
} // end of TextBuffer

func AssertTextBuffer(_ TextBuffer) {}

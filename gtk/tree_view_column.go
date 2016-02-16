package gtk

import "github.com/gotk3/gotk3/glib"

type TreeViewColumn interface {
	glib.InitiallyUnowned

	AddAttribute(CellRenderer, string, int)
	CellIsVisible() bool
	Clear()
	ClearAttributes(CellRenderer)
	Clicked()
	FocusCell(CellRenderer)
	GetClickable() bool
	GetExpand() bool
	GetFixedWidth() int
	GetMaxWidth() int
	GetMinWidth() int
	GetReorderable() bool
	GetResizable() bool
	GetSortColumnID() int
	GetSortIndicator() bool
	GetSpacing() int
	GetTitle() string
	GetVisible() bool
	GetWidth() int
	GetXOffset() int
	PackEnd(CellRenderer, bool)
	PackStart(CellRenderer, bool)
	QueueResize()
	SetClickable(bool)
	SetExpand(bool)
	SetFixedWidth(int)
	SetMaxWidth(int)
	SetMinWidth(int)
	SetReorderable(bool)
	SetResizable(bool)
	SetSortColumnID(int)
	SetSortIndicator(bool)
	SetSpacing(int)
	SetTitle(string)
	SetVisible(bool)
} // end of TreeViewColumn

func AssertTreeViewColumn(_ TreeViewColumn) {}

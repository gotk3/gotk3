package glib

type Menu interface {
	MenuModel

	Append(string, string)
	AppendItem(MenuItem)
	AppendSection(string, MenuModel)
	AppendSubmenu(string, MenuModel)
	Freeze()
	Insert(int, string, string)
	InsertItem(int, MenuItem)
	InsertSection(int, string, MenuModel)
	InsertSubmenu(int, string, MenuModel)
	Prepend(string, string)
	PrependItem(MenuItem)
	PrependSection(string, MenuModel)
	PrependSubmenu(string, MenuModel)
	Remove(int)
	RemoveAll()
} // end of Menu

func AssertMenu(_ Menu) {}

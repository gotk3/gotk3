package iface

type Menu interface {
	MenuShell

	GetAccelGroup() AccelGroup
	GetAccelPath() string
	Popdown()
	PopupAtMouseCursor(Menu, MenuItem, int, uint32)
	ReorderChild(Widget, int)
	SetAccelGroup(AccelGroup)
} // end of Menu

func AssertMenu(_ Menu) {}

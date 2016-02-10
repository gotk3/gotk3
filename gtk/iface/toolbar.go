package iface

type Toolbar interface {
	Container

	GetDropIndex(int, int) int
	GetIconSize() IconSize
	GetItemIndex(ToolItem) int
	GetNItems() int
	GetNthItem(int) ToolItem
	GetReliefStyle() ReliefStyle
	GetShowArrow() bool
	GetStyle() ToolbarStyle
	Insert(ToolItem, int)
	SetDropHighlightItem(ToolItem, int)
	SetIconSize(IconSize)
	SetShowArrow(bool)
	SetStyle(ToolbarStyle)
	UnsetIconSize()
	UnsetStyle()
} // end of Toolbar

func AssertToolbar(_ Toolbar) {}

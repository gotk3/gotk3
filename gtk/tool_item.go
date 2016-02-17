package gtk

type ToolItem interface {
	Bin

	GetExpand() bool
	GetHomogeneous() bool
	GetIconSize() IconSize
	GetIsImportant() bool
	GetOrientation() Orientation
	GetReliefStyle() ReliefStyle
	GetTextAlignment() float32
	GetTextOrientation() Orientation
	GetUseDragWindow() bool
	GetVisibleHorizontal() bool
	GetVisibleVertical() bool
	RebuildMenu()
	RetrieveProxyMenuItem() MenuItem
	SetExpand(bool)
	SetHomogeneous(bool)
	SetIsImportant(bool)
	SetProxyMenuItem(string, MenuItem)
	SetTooltipMarkup(string)
	SetUseDragWindow(bool)
	SetVisibleHorizontal(bool)
	SetVisibleVertical(bool)
	ToolbarReconfigured()
} // end of ToolItem

func AssertToolItem(_ ToolItem) {}

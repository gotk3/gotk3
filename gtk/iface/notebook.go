package iface

type Notebook interface {
	Container

	AppendPage(Widget, Widget) int
	AppendPageMenu(Widget, Widget, Widget) int
	GetActionWidget(PackType) (Widget, error)
	GetCurrentPage() int
	GetGroupName() (string, error)
	GetMenuLabel(Widget) (Widget, error)
	GetMenuLabelText(Widget) (string, error)
	GetNPages() int
	GetNthPage(int) (Widget, error)
	GetScrollable() bool
	GetShowBorder() bool
	GetShowTabs() bool
	GetTabDetachable(Widget) bool
	GetTabLabel(Widget) (Widget, error)
	GetTabLabelText(Widget) (string, error)
	GetTabPos() PositionType
	GetTabReorderable(Widget) bool
	InsertPage(Widget, Widget, int) int
	InsertPageMenu(Widget, Widget, Widget, int) int
	NextPage()
	PageNum(Widget) int
	PopupDisable()
	PopupEnable()
	PrependPage(Widget, Widget) int
	PrependPageMenu(Widget, Widget, Widget) int
	PrevPage()
	RemovePage(int)
	ReorderChild(Widget, int)
	SetActionWidget(Widget, PackType)
	SetCurrentPage(int)
	SetGroupName(string)
	SetMenuLabel(Widget, Widget)
	SetMenuLabelText(Widget, string)
	SetScrollable(bool)
	SetShowBorder(bool)
	SetShowTabs(bool)
	SetTabDetachable(Widget, bool)
	SetTabLabel(Widget, Widget)
	SetTabLabelText(Widget, string)
	SetTabPos(PositionType)
	SetTabReorderable(Widget, bool)
} // end of Notebook

func AssertNotebook(_ Notebook) {}

package iface

type Assistant interface {
	Window

	AddActionWidget(Widget)
	AppendPage(Widget) int
	Commit()
	GetCurrentPage() int
	GetNPages() int
	GetNthPage(int) (Widget, error)
	GetPageComplete(Widget) bool
	GetPageTitle(Widget) string
	GetPageType(Widget) AssistantPageType
	InsertPage(Widget, int) int
	NextPage()
	PrependPage(Widget) int
	PreviousPage()
	RemoveActionWidget(Widget)
	RemovePage(int)
	SetCurrentPage(int)
	SetPageComplete(Widget, bool)
	SetPageTitle(Widget, string)
	SetPageType(Widget, AssistantPageType)
	UpdateButtonsState()
} // end of Assistant

func AssertAssistant(_ Assistant) {}

package gtk

type SeparatorToolItem interface {
	ToolItem

	GetDraw() bool
	SetDraw(bool)
} // end of SeparatorToolItem

func AssertSeparatorToolItem(_ SeparatorToolItem) {}

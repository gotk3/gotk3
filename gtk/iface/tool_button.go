package iface

type ToolButton interface {
	ToolItem

	GetIconName() string
	GetIconWidget() Widget
	GetLabel() string
	GetLabelWidget() Widget
	GetuseUnderline() bool
	SetGetUnderline(bool)
	SetIconName(string)
	SetIconWidget(Widget)
	SetLabel(string)
	SetLabelWidget(Widget)
} // end of ToolButton

func AssertToolButton(_ ToolButton) {}

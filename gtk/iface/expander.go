package iface

type Expander interface {
	Bin

	GetExpanded() bool
	GetLabel() string
	SetExpanded(bool)
	SetLabel(string)
	SetLabelWidget(Widget)
} // end of Expander

func AssertExpander(_ Expander) {}

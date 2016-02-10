package iface

type MenuItem interface {
	Bin

	GetAccelPath() string
	SetLabel(string)
	SetSubmenu(Widget)
} // end of MenuItem

func AssertMenuItem(_ MenuItem) {}

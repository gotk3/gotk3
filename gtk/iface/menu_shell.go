package iface

type MenuShell interface {
	Container

	ActivateItem(MenuItem, bool)
	Append(MenuItem)
	Cancel()
	Deactivate()
	Deselect()
	Insert(MenuItem, int)
	Prepend(MenuItem)
	SelectFirst(bool)
	SelectItem(MenuItem)
	SetTakeFocus(bool)
} // end of MenuShell

func AssertMenuShell(_ MenuShell) {}

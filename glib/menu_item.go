package glib

type MenuItem interface {
	Object

	GetLink(string) MenuModel
	SetDetailedAction(string)
	SetLabel(string)
	SetLink(string, MenuModel)
	SetSection(MenuModel)
	SetSubmenu(MenuModel)
} // end of MenuItem

func AssertMenuItem(_ MenuItem) {}

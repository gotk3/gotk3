package impl

import "github.com/gotk3/gotk3/glib"

func CastToList(s glib.List) *List {
	if s == nil {
		return nil
	}
	return s.(*List)
}

func CastToObject(s glib.Object) *Object {
	if s == nil {
		return nil
	}
	return s.(*Object)
}

func CastToSList(s glib.SList) *SList {
	if s == nil {
		return nil
	}
	return s.(*SList)
}

func toNotification(s glib.Notification) *notification {
	if s == nil {
		return nil
	}
	return s.(*notification)
}

func toMenuItem(s glib.MenuItem) *menuItem {
	if s == nil {
		return nil
	}
	return s.(*menuItem)
}

func CastToMenuModel(s glib.MenuModel) *MenuModel {
	if s == nil {
		return nil
	}
	return s.(*MenuModel)
}
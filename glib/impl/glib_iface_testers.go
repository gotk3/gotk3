package impl

import "github.com/gotk3/gotk3/glib"

func init() {
	glib.AssertGlib(&RealGlib{})
	glib.AssertApplication(&Application{})
	glib.AssertInitiallyUnowned(&InitiallyUnowned{})
	glib.AssertList(&List{})
	glib.AssertMenu(&Menu{})
	glib.AssertMenuItem(&menuItem{})
	glib.AssertMenuModel(&MenuModel{})
	glib.AssertNotification(&notification{})
	glib.AssertObject(&Object{})
	glib.AssertSList(&SList{})
	glib.AssertSignal(&signal{})
	glib.AssertTypeMarshaler(&TypeMarshaler{})
	glib.AssertValue(&value{})
}

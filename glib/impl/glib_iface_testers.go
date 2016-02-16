package impl

import "github.com/gotk3/gotk3/glib"

func init() {
	glib.AssertGlib(&RealGlib{})
	glib.AssertApplication(&Application{})
	glib.AssertInitiallyUnowned(&InitiallyUnowned{})
	glib.AssertList(&List{})
	glib.AssertMenu(&Menu{})
	glib.AssertMenuItem(&MenuItem{})
	glib.AssertMenuModel(&MenuModel{})
	glib.AssertNotification(&Notification{})
	glib.AssertObject(&Object{})
	glib.AssertSList(&SList{})
	glib.AssertSignal(&Signal{})
	glib.AssertTypeMarshaler(&TypeMarshaler{})
	glib.AssertValue(&Value{})
}

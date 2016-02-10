package glib

import "github.com/gotk3/gotk3/glib/iface"

func init() {
  iface.AssertGlib(&RealGlib{})
  iface.AssertApplication(&Application{})
  iface.AssertInitiallyUnowned(&InitiallyUnowned{})
  iface.AssertList(&List{})
  iface.AssertMenu(&Menu{})
  iface.AssertMenuItem(&MenuItem{})
  iface.AssertMenuModel(&MenuModel{})
  iface.AssertNotification(&Notification{})
  iface.AssertObject(&Object{})
  iface.AssertSList(&SList{})
  iface.AssertSignal(&Signal{})
  iface.AssertTypeMarshaler(&TypeMarshaler{})
  iface.AssertValue(&Value{})
}

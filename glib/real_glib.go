package glib

import "github.com/gotk3/gotk3/glib/iface"
import "unsafe"

type RealGlib struct{}

var Real = &RealGlib{}

func (*RealGlib) ApplicationGetDefault() iface.Application {
	return ApplicationGetDefault()
}

func (*RealGlib) ApplicationIDIsValid(id string) bool {
	return ApplicationIDIsValid(id)
}

func (*RealGlib) ApplicationNew(appID string, flags iface.ApplicationFlags) iface.Application {
	return ApplicationNew(appID, flags)
}

func (*RealGlib) GValue(v interface{}) (iface.Value, error) {
	return GValue(v)
}

func (*RealGlib) GetApplicationName() string {
	return GetApplicationName()
}

func (*RealGlib) GetUserSpecialDir(directory iface.UserDirectory) (string, error) {
	return GetUserSpecialDir(directory)
}

func (*RealGlib) IdleAdd(f interface{}, args ...interface{}) (iface.SourceHandle, error) {
	return IdleAdd(f, args...)
}

func (*RealGlib) InitI18n(domain string, dir string) {
	InitI18n(domain, dir)
}

func (*RealGlib) MenuItemNew(label string, detailed_action string) iface.MenuItem {
	return MenuItemNew(label, detailed_action)
}

func (*RealGlib) MenuItemNewFromModel(model iface.MenuModel, index int) iface.MenuItem {
	return MenuItemNewFromModel(model.(*MenuModel), index)
}

func (*RealGlib) MenuItemNewSection(label string, section iface.MenuModel) iface.MenuItem {
	return MenuItemNewSection(label, section.(*MenuModel))
}

func (*RealGlib) MenuItemNewSubmenu(label string, submenu iface.MenuModel) iface.MenuItem {
	return MenuItemNewSubmenu(label, submenu.(*MenuModel))
}

func (*RealGlib) MenuNew() iface.Menu {
	return MenuNew()
}

func (*RealGlib) NotificationNew(title string) iface.Notification {
	return NotificationNew(title)
}

func (*RealGlib) RegisterGValueMarshalers(tm []iface.TypeMarshaler) {
	tp := []TypeMarshaler{}
	for _, tt := range tm {
		tp = append(tp, tt.(TypeMarshaler))
	}
	RegisterGValueMarshalers(tp)
}

func (*RealGlib) SetApplicationName(name string) {
	SetApplicationName(name)
}

func (*RealGlib) SignalNew(s string) (iface.Signal, error) {
	return SignalNew(s)
}

func (*RealGlib) TimeoutAdd(timeout uint, f interface{}, args ...interface{}) (iface.SourceHandle, error) {
	return TimeoutAdd(timeout, f, args...)
}

func (*RealGlib) TypeDepth(t iface.Type) uint {
	return TypeDepth(t)
}

func (*RealGlib) TypeName(t iface.Type) string {
	return TypeName(t)
}

func (*RealGlib) TypeParent(t iface.Type) iface.Type {
	return TypeParent(t)
}

func (*RealGlib) ValueAlloc() (iface.Value, error) {
	return ValueAlloc()
}

func (*RealGlib) ValueFromNative(l unsafe.Pointer) iface.Value {
	return ValueFromNative(l)
}

func (*RealGlib) ValueInit(t iface.Type) (iface.Value, error) {
	return ValueInit(t)
}

func (*RealGlib) WrapList(obj uintptr) iface.List {
	return WrapList(obj)
}

func (*RealGlib) WrapSList(obj uintptr) iface.SList {
	return WrapSList(obj)
}

package impl

import "github.com/gotk3/gotk3/glib"
import "unsafe"

type RealGlib struct{}

var Real = &RealGlib{}

func (*RealGlib) ApplicationGetDefault() glib.Application {
	return ApplicationGetDefault()
}

func (*RealGlib) ApplicationIDIsValid(id string) bool {
	return ApplicationIDIsValid(id)
}

func (*RealGlib) ApplicationNew(appID string, flags glib.ApplicationFlags) glib.Application {
	return ApplicationNew(appID, flags)
}

func (*RealGlib) GValue(v interface{}) (glib.Value, error) {
	return GValue(v)
}

func (*RealGlib) GetApplicationName() string {
	return GetApplicationName()
}

func (*RealGlib) GetUserSpecialDir(directory glib.UserDirectory) (string, error) {
	return GetUserSpecialDir(directory)
}

func (*RealGlib) IdleAdd(f interface{}, args ...interface{}) (glib.SourceHandle, error) {
	return IdleAdd(f, args...)
}

func (*RealGlib) InitI18n(domain string, dir string) {
	InitI18n(domain, dir)
}

func (*RealGlib) MenuItemNew(label string, detailed_action string) glib.MenuItem {
	return MenuItemNew(label, detailed_action)
}

func (*RealGlib) MenuItemNewFromModel(model glib.MenuModel, index int) glib.MenuItem {
	return MenuItemNewFromModel(CastToMenuModel(model), index)
}

func (*RealGlib) MenuItemNewSection(label string, section glib.MenuModel) glib.MenuItem {
	return MenuItemNewSection(label, CastToMenuModel(section))
}

func (*RealGlib) MenuItemNewSubmenu(label string, submenu glib.MenuModel) glib.MenuItem {
	return MenuItemNewSubmenu(label, CastToMenuModel(submenu))
}

func (*RealGlib) MenuNew() glib.Menu {
	return MenuNew()
}

func (*RealGlib) NotificationNew(title string) glib.Notification {
	return NotificationNew(title)
}

func (*RealGlib) RegisterGValueMarshalers(tm []glib.TypeMarshaler) {
	tp := []TypeMarshaler{}
	for _, tt := range tm {
		tp = append(tp, tt.(TypeMarshaler))
	}
	RegisterGValueMarshalers(tp)
}

func (*RealGlib) SetApplicationName(name string) {
	SetApplicationName(name)
}

func (*RealGlib) SignalNew(s string) (glib.Signal, error) {
	return SignalNew(s)
}

func (*RealGlib) TimeoutAdd(timeout uint, f interface{}, args ...interface{}) (glib.SourceHandle, error) {
	return TimeoutAdd(timeout, f, args...)
}

func (*RealGlib) TypeDepth(t glib.Type) uint {
	return TypeDepth(t)
}

func (*RealGlib) TypeName(t glib.Type) string {
	return TypeName(t)
}

func (*RealGlib) TypeParent(t glib.Type) glib.Type {
	return TypeParent(t)
}

func (*RealGlib) ValueAlloc() (glib.Value, error) {
	return ValueAlloc()
}

func (*RealGlib) ValueFromNative(l unsafe.Pointer) glib.Value {
	return ValueFromNative(l)
}

func (*RealGlib) ValueInit(t glib.Type) (glib.Value, error) {
	return ValueInit(t)
}

func (*RealGlib) WrapList(obj uintptr) glib.List {
	return WrapList(obj)
}

func (*RealGlib) WrapSList(obj uintptr) glib.SList {
	return WrapSList(obj)
}

package iface

import "unsafe"

type Glib interface {
    ApplicationGetDefault() Application
    ApplicationIDIsValid(string) bool
    ApplicationNew(string, ApplicationFlags) Application
    GValue(interface{}) (Value, error)
    GetApplicationName() string
    GetUserSpecialDir(UserDirectory) (string, error)
    IdleAdd(interface{}, ...interface{}) (SourceHandle, error)
    InitI18n(string, string)
    MenuItemNew(string, string) MenuItem
    MenuItemNewFromModel(MenuModel, int) MenuItem
    MenuItemNewSection(string, MenuModel) MenuItem
    MenuItemNewSubmenu(string, MenuModel) MenuItem
    MenuNew() Menu
    NotificationNew(string) Notification
    RegisterGValueMarshalers([]TypeMarshaler)
    SetApplicationName(string)
    SignalNew(string) (Signal, error)
    TimeoutAdd(uint, interface{}, ...interface{}) (SourceHandle, error)
    TypeDepth(Type) uint
    TypeName(Type) string
    TypeParent(Type) Type
    ValueAlloc() (Value, error)
    ValueFromNative(unsafe.Pointer) Value
    ValueInit(Type) (Value, error)
    WrapList(uintptr) List
    WrapSList(uintptr) SList
} // end of Glib

func AssertGlib(_ Glib) {}

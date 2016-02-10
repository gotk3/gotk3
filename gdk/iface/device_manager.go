package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type DeviceManager interface {
    glib_iface.Object

    GetClientPointer() (Device, error)
    GetDisplay() (Display, error)
    ListDevices(DeviceType) glib_iface.List
} // end of DeviceManager

func AssertDeviceManager(_ DeviceManager) {}

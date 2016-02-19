package gdk

import "github.com/gotk3/gotk3/glib"

type DeviceManager interface {
	glib.Object

	GetClientPointer() (Device, error)
	GetDisplay() (Display, error)
	ListDevices(DeviceType) glib.List
} // end of DeviceManager

func AssertDeviceManager(_ DeviceManager) {}

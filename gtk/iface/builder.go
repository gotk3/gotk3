package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type Builder interface {
	glib_iface.Object

	AddFromFile(string) error
	AddFromResource(string) error
	AddFromString(string) error
	ConnectSignals(map[string]interface{})
	GetObject(string) (glib_iface.Object, error)
} // end of Builder

func AssertBuilder(_ Builder) {}

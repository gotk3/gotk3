// +build !gtk_3_6,!gtk_3_8

package gtk

import "github.com/gotk3/gotk3/gtk/iface"

type RealGtkSince310 struct {
	RealGtk
}

var RealSince310 = &RealGtkSince310{}

func (*RealGtkSince310) RevealerNew() (iface.Revealer, error) {
	return RevealerNew()
}

func init() {
	iface.AssertGtkSince310(&RealGtkSince310{})
}

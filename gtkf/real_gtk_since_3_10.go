// +build !gtk_3_6,!gtk_3_8

package gtkf

import "github.com/gotk3/gotk3/gtk"

type RealGtkSince310 struct {
	RealGtk
}

var RealSince310 = &RealGtkSince310{}

func (*RealGtkSince310) RevealerNew() (gtk.Revealer, error) {
	return RevealerNew()
}

func init() {
	gtk.AssertGtkSince310(&RealGtkSince310{})
}

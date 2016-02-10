package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"
import cairo_iface "github.com/gotk3/gotk3/cairo/iface"

type Container interface {
	Widget

	Add(Widget)
	CheckResize()
	ChildNotify(Widget, string)
	ChildSetProperty(Widget, string, interface{}) error
	ChildType() glib_iface.Type
	GetBorderWidth() uint
	GetFocusChain() ([]Widget, bool)
	GetFocusChild() Widget
	GetFocusHAdjustment() Adjustment
	GetFocusVAdjustment() Adjustment
	PropagateDraw(Widget, cairo_iface.Context)
	Remove(Widget)
	SetBorderWidth(uint)
	SetFocusChain([]Widget)
	SetFocusChild(Widget)
	SetFocusHAdjustment(Adjustment)
	SetFocusVAdjustment(Adjustment)
} // end of Container

func AssertContainer(_ Container) {}

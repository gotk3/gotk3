package gtk

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/glib"
)

type Container interface {
	Widget

	Add(Widget)
	CheckResize()
	ChildNotify(Widget, string)
	ChildSetProperty(Widget, string, interface{}) error
	ChildType() glib.Type
	GetBorderWidth() uint
	GetFocusChain() ([]Widget, bool)
	GetFocusChild() Widget
	GetFocusHAdjustment() Adjustment
	GetFocusVAdjustment() Adjustment
	PropagateDraw(Widget, cairo.Context)
	Remove(Widget)
	SetBorderWidth(uint)
	SetFocusChain([]Widget)
	SetFocusChild(Widget)
	SetFocusHAdjustment(Adjustment)
	SetFocusVAdjustment(Adjustment)
} // end of Container

func AssertContainer(_ Container) {}

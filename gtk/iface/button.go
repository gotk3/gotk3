package iface

import gdk_iface "github.com/gotk3/gotk3/gdk/iface"

type Button interface {
	Bin

	Clicked()
	GetAlwaysShowImage() bool
	GetEventWindow() (gdk_iface.Window, error)
	GetFocusOnClick() bool
	GetImage() (Widget, error)
	GetImagePosition() PositionType
	GetLabel() (string, error)
	GetRelief() ReliefStyle
	GetUseUnderline() bool
	SetAlwaysShowImage(bool)
	SetFocusOnClick(bool)
	SetImage(Widget)
	SetImagePosition(PositionType)
	SetLabel(string)
	SetRelief(ReliefStyle)
	SetUseUnderline(bool)
} // end of Button

func AssertButton(_ Button) {}

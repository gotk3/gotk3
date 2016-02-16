package gtk

import "github.com/gotk3/gotk3/glib"

type RadioButton interface {
	CheckButton

	GetGroup() (glib.SList, error)
	JoinGroup(RadioButton)
	SetGroup(glib.SList)
} // end of RadioButton

func AssertRadioButton(_ RadioButton) {}

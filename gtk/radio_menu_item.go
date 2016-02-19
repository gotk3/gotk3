package gtk

import "github.com/gotk3/gotk3/glib"

type RadioMenuItem interface {
	CheckMenuItem

	GetGroup() (glib.SList, error)
	SetGroup(glib.SList)
} // end of RadioMenuItem

func AssertRadioMenuItem(_ RadioMenuItem) {}

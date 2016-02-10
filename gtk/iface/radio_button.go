package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type RadioButton interface {
    CheckButton

    GetGroup() (glib_iface.SList, error)
    JoinGroup(RadioButton)
    SetGroup(glib_iface.SList)
} // end of RadioButton

func AssertRadioButton(_ RadioButton) {}

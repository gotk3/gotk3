package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type Adjustment interface {
    glib_iface.InitiallyUnowned

    Configure(float64, float64, float64, float64, float64, float64)
    GetLower() float64
    GetMinimumIncrement() float64
    GetPageIncrement() float64
    GetPageSize() float64
    GetStepIncrement() float64
    GetUpper() float64
    GetValue() float64
    SetLower(float64)
    SetPageIncrement(float64)
    SetPageSize(float64)
    SetStepIncrement(float64)
    SetUpper(float64)
    SetValue(float64)
} // end of Adjustment

func AssertAdjustment(_ Adjustment) {}

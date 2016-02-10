package iface


type ScaleButton interface {
    Button

    GetAdjustment() Adjustment
    GetPopup() (Widget, error)
    GetValue() float64
    SetAdjustment(Adjustment)
    SetValue(float64)
} // end of ScaleButton

func AssertScaleButton(_ ScaleButton) {}

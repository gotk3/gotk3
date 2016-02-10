package iface


type LevelBar interface {
    Widget

    AddOffsetValue(string, float64)
    GetMaxValue() float64
    GetMinValue() float64
    GetMode() LevelBarMode
    GetOffsetValue(string) (float64, bool)
    GetValue() float64
    RemoveOffsetValue(string)
    SetMaxValue(float64)
    SetMinValue(float64)
    SetMode(LevelBarMode)
    SetValue(float64)
} // end of LevelBar

func AssertLevelBar(_ LevelBar) {}

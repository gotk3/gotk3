package gtk

type SpinButton interface {
	Entry

	Configure(Adjustment, float64, uint)
	GetAdjustment() Adjustment
	GetValue() float64
	GetValueAsInt() int
	SetIncrements(float64, float64)
	SetRange(float64, float64)
	SetValue(float64)
} // end of SpinButton

func AssertSpinButton(_ SpinButton) {}

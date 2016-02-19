package gtk

type Range interface {
	Widget

	GetValue() float64
	SetIncrements(float64, float64)
	SetRange(float64, float64)
	SetValue(float64)
} // end of Range

func AssertRange(_ Range) {}

package gtk

type ProgressBar interface {
	Widget

	GetFraction() float64
	GetShowText() bool
	SetFraction(float64)
	SetShowText(bool)
	SetText(string)
} // end of ProgressBar

func AssertProgressBar(_ ProgressBar) {}

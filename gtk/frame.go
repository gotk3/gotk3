package gtk

type Frame interface {
	Bin

	GetLabel() string
	GetLabelAlign() (float32, float32)
	GetLabelWidget() (Widget, error)
	GetShadowType() ShadowType
	SetLabel(string)
	SetLabelAlign(float32, float32)
	SetLabelWidget(Widget)
	SetShadowType(ShadowType)
} // end of Frame

func AssertFrame(_ Frame) {}

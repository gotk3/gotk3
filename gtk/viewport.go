package gtk

type Viewport interface {
	Bin

	GetHAdjustment() (Adjustment, error)
	GetVAdjustment() (Adjustment, error)
	SetHAdjustment(Adjustment)
	SetVAdjustment(Adjustment)
} // end of Viewport

func AssertViewport(_ Viewport) {}

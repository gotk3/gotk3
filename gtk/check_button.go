package gtk

type CheckButton interface {
	ToggleButton
} // end of CheckButton

func AssertCheckButton(_ CheckButton) {}

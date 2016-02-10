package iface


type CheckButton interface {
    ToggleButton
} // end of CheckButton

func AssertCheckButton(_ CheckButton) {}

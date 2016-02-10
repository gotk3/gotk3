package iface

type ColorButton interface {
	Button
} // end of ColorButton

func AssertColorButton(_ ColorButton) {}

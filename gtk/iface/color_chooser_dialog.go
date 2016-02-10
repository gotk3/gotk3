package iface

type ColorChooserDialog interface {
	Dialog
} // end of ColorChooserDialog

func AssertColorChooserDialog(_ ColorChooserDialog) {}

package iface

type FileChooserWidget interface {
	Box
} // end of FileChooserWidget

func AssertFileChooserWidget(_ FileChooserWidget) {}

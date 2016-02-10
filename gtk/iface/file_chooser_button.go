package iface

type FileChooserButton interface {
	Box
} // end of FileChooserButton

func AssertFileChooserButton(_ FileChooserButton) {}

package gtk

type FileChooserDialog interface {
	Dialog
} // end of FileChooserDialog

func AssertFileChooserDialog(_ FileChooserDialog) {}

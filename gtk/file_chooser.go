package gtk

import "github.com/gotk3/gotk3/glib"

type FileChooser interface {
	glib.Object

	AddFilter(FileFilter)
	AddShortcutFolder(string) bool
	GetCurrentFolder() (string, error)
	GetFilename() string
	GetPreviewFilename() string
	GetURI() string
	SetCurrentFolder(string) bool
	SetCurrentName(string)
	SetPreviewWidget(Widget)
	SetPreviewWidgetActive(bool)
} // end of FileChooser

func AssertFileChooser(_ FileChooser) {}

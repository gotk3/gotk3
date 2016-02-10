package iface

import glib_iface "github.com/gotk3/gotk3/glib/iface"

type FileChooser interface {
	glib_iface.Object

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

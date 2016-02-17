package gdk

import "github.com/gotk3/gotk3/glib"

type Gdk interface {
	DisplayGetDefault() (Display, error)
	DisplayOpen(string) (Display, error)
	EventButtonFrom(Event) EventButton
	EventKeyNew() EventKey
	GdkAtomIntern(string, bool) Atom
	KeyvalConvertCase(uint) (uint, uint)
	KeyvalFromName(string) uint
	KeyvalIsLower(uint) bool
	KeyvalIsUpper(uint) bool
	KeyvalToLower(uint) uint
	KeyvalToUpper(uint) uint
	NewRGBA(...float64) RGBA
	PixbufCopy(Pixbuf) (Pixbuf, error)
	PixbufGetFileInfo(string) (interface{}, int, int)
	PixbufGetType() glib.Type
	PixbufLoaderNew() (PixbufLoader, error)
	PixbufNew(Colorspace, bool, int, int, int) (Pixbuf, error)
	PixbufNewFromFile(string) (Pixbuf, error)
	PixbufNewFromFileAtScale(string, int, int, bool) (Pixbuf, error)
	PixbufNewFromFileAtSize(string, int, int) (Pixbuf, error)
	ScreenGetDefault() (Screen, error)
	WorkspaceControlSupported() bool
} // end of Gdk

func AssertGdk(_ Gdk) {}

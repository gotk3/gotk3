package impl

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
)

type RealGdk struct{}

var Real = &RealGdk{}

func (*RealGdk) DisplayGetDefault() (gdk.Display, error) {
	return DisplayGetDefault()
}

func (*RealGdk) DisplayOpen(displayName string) (gdk.Display, error) {
	return DisplayOpen(displayName)
}

func (*RealGdk) EventButtonFrom(ev gdk.Event) gdk.EventButton {
	return EventButtonFrom(ev.(*Event))
}

func (*RealGdk) EventKeyNew() gdk.EventKey {
	return EventKeyNew()
}

func (*RealGdk) GdkAtomIntern(atomName string, onlyIfExists bool) gdk.Atom {
	return GdkAtomIntern(atomName, onlyIfExists)
}

func (*RealGdk) KeyvalConvertCase(v uint) (uint, uint) {
	return KeyvalConvertCase(v)
}

func (*RealGdk) KeyvalFromName(keyvalName string) uint {
	return KeyvalFromName(keyvalName)
}

func (*RealGdk) KeyvalIsLower(v uint) bool {
	return KeyvalIsLower(v)
}

func (*RealGdk) KeyvalIsUpper(v uint) bool {
	return KeyvalIsUpper(v)
}

func (*RealGdk) KeyvalToLower(v uint) uint {
	return KeyvalToLower(v)
}

func (*RealGdk) KeyvalToUpper(v uint) uint {
	return KeyvalToUpper(v)
}

func (*RealGdk) NewRGBA(values ...float64) gdk.RGBA {
	return NewRGBA(values...)
}

func (*RealGdk) PixbufCopy(v gdk.Pixbuf) (gdk.Pixbuf, error) {
	return PixbufCopy(v.(*Pixbuf))
}

func (*RealGdk) PixbufGetFileInfo(filename string) (interface{}, int, int) {
	return PixbufGetFileInfo(filename)
}

func (*RealGdk) PixbufGetType() glib.Type {
	return PixbufGetType()
}

func (*RealGdk) PixbufLoaderNew() (gdk.PixbufLoader, error) {
	return PixbufLoaderNew()
}

func (*RealGdk) PixbufNew(colorspace gdk.Colorspace, hasAlpha bool, bitsPerSample int, width int, height int) (gdk.Pixbuf, error) {
	return PixbufNew(colorspace, hasAlpha, bitsPerSample, width, height)
}

func (*RealGdk) PixbufNewFromFile(filename string) (gdk.Pixbuf, error) {
	return PixbufNewFromFile(filename)
}

func (*RealGdk) PixbufNewFromFileAtScale(filename string, width int, height int, preserveAspectRatio bool) (gdk.Pixbuf, error) {
	return PixbufNewFromFileAtScale(filename, width, height, preserveAspectRatio)
}

func (*RealGdk) PixbufNewFromFileAtSize(filename string, width int, height int) (gdk.Pixbuf, error) {
	return PixbufNewFromFileAtSize(filename, width, height)
}

func (*RealGdk) ScreenGetDefault() (gdk.Screen, error) {
	return ScreenGetDefault()
}

func (*RealGdk) WorkspaceControlSupported() bool {
	return WorkspaceControlSupported()
}

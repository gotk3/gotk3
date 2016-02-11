package gdk

import "github.com/gotk3/gotk3/gdk/iface"
import glib_iface "github.com/gotk3/gotk3/glib/iface"

type RealGdk struct{}

var Real = &RealGdk{}

func (*RealGdk) DisplayGetDefault() (iface.Display, error) {
	return DisplayGetDefault()
}

func (*RealGdk) DisplayOpen(displayName string) (iface.Display, error) {
	return DisplayOpen(displayName)
}

func (*RealGdk) EventButtonFrom(ev iface.Event) iface.EventButton {
	return EventButtonFrom(ev.(*Event))
}

func (*RealGdk) EventKeyNew() iface.EventKey {
	return EventKeyNew()
}

func (*RealGdk) GdkAtomIntern(atomName string, onlyIfExists bool) iface.Atom {
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

func (*RealGdk) NewRGBA(values ...float64) iface.RGBA {
	return NewRGBA(values...)
}

func (*RealGdk) PixbufCopy(v iface.Pixbuf) (iface.Pixbuf, error) {
	return PixbufCopy(v.(*Pixbuf))
}

func (*RealGdk) PixbufGetFileInfo(filename string) (interface{}, int, int) {
	return PixbufGetFileInfo(filename)
}

func (*RealGdk) PixbufGetType() glib_iface.Type {
	return PixbufGetType()
}

func (*RealGdk) PixbufLoaderNew() (iface.PixbufLoader, error) {
	return PixbufLoaderNew()
}

func (*RealGdk) PixbufNew(colorspace iface.Colorspace, hasAlpha bool, bitsPerSample int, width int, height int) (iface.Pixbuf, error) {
	return PixbufNew(colorspace, hasAlpha, bitsPerSample, width, height)
}

func (*RealGdk) PixbufNewFromFile(filename string) (iface.Pixbuf, error) {
	return PixbufNewFromFile(filename)
}

func (*RealGdk) PixbufNewFromFileAtScale(filename string, width int, height int, preserveAspectRatio bool) (iface.Pixbuf, error) {
	return PixbufNewFromFileAtScale(filename, width, height, preserveAspectRatio)
}

func (*RealGdk) PixbufNewFromFileAtSize(filename string, width int, height int) (iface.Pixbuf, error) {
	return PixbufNewFromFileAtSize(filename, width, height)
}

func (*RealGdk) ScreenGetDefault() (iface.Screen, error) {
	return ScreenGetDefault()
}

func (*RealGdk) WorkspaceControlSupported() bool {
	return WorkspaceControlSupported()
}

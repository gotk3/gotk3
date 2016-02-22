package gdkf

import "github.com/gotk3/gotk3/gdk"

func CastToDisplay(s gdk.Display) *Display {
	if s == nil {
		return nil
	}
	return s.(*Display)
}

func CastToEvent(s gdk.Event) *Event {
	if s == nil {
		return nil
	}
	return s.(*Event)
}

func CastToPixbuf(s gdk.Pixbuf) *Pixbuf {
	if s == nil {
		return nil
	}
	return s.(*Pixbuf)
}

func CastToRGBA(s gdk.RGBA) *RGBA {
	if s == nil {
		return nil
	}
	return s.(*RGBA)
}

func CastToScreen(s gdk.Screen) *Screen {
	if s == nil {
		return nil
	}
	return s.(*Screen)
}

func CastToWindow(s gdk.Window) *Window {
	if s == nil {
		return nil
	}
	return s.(*Window)
}

func CastToDevice(s gdk.Device) *Device {
	if s == nil {
		return nil
	}
	return s.(*Device)
}

func castToCursor(s gdk.Cursor) *cursor {
	if s == nil {
		return nil
	}
	return s.(*cursor)
}

func CastToVisual(s gdk.Visual) *Visual {
	if s == nil {
		return nil
	}
	return s.(*Visual)
}

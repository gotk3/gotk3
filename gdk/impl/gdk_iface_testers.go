package impl

import "github.com/gotk3/gotk3/gdk"

func init() {
	gdk.AssertGdk(&RealGdk{})
	gdk.AssertCursor(&Cursor{})
	gdk.AssertDevice(&Device{})
	gdk.AssertDeviceManager(&DeviceManager{})
	gdk.AssertDisplay(&Display{})
	gdk.AssertDragContext(&DragContext{})
	gdk.AssertEvent(&Event{})
	gdk.AssertEventButton(&EventButton{})
	gdk.AssertEventKey(&EventKey{})
	gdk.AssertEventMotion(&EventMotion{})
	gdk.AssertEventScroll(&EventScroll{})
	gdk.AssertPixbuf(&Pixbuf{})
	gdk.AssertPixbufLoader(&PixbufLoader{})
	gdk.AssertRGBA(&RGBA{})
	gdk.AssertRectangle(&Rectangle{})
	gdk.AssertScreen(&Screen{})
	gdk.AssertVisual(&Visual{})
	gdk.AssertWindow(&Window{})
}

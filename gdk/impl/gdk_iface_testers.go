package impl

import "github.com/gotk3/gotk3/gdk"

func init() {
	gdk.AssertGdk(&RealGdk{})
	gdk.AssertCursor(&cursor{})
	gdk.AssertDevice(&Device{})
	gdk.AssertDeviceManager(&deviceManager{})
	gdk.AssertDisplay(&Display{})
	gdk.AssertDragContext(&dragContext{})
	gdk.AssertEvent(&Event{})
	gdk.AssertEventButton(&eventButton{})
	gdk.AssertEventKey(&eventKey{})
	gdk.AssertEventMotion(&eventMotion{})
	gdk.AssertEventScroll(&eventScroll{})
	gdk.AssertPixbuf(&Pixbuf{})
	gdk.AssertPixbufLoader(&pixbufLoader{})
	gdk.AssertRGBA(&RGBA{})
	gdk.AssertRectangle(&Rectangle{})
	gdk.AssertScreen(&Screen{})
	gdk.AssertVisual(&Visual{})
	gdk.AssertWindow(&Window{})
}

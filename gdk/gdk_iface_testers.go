package gdk

import "github.com/gotk3/gotk3/gdk/iface"

func init() {
  iface.AssertGdk(&RealGdk{})
  iface.AssertCursor(&Cursor{})
  iface.AssertDevice(&Device{})
  iface.AssertDeviceManager(&DeviceManager{})
  iface.AssertDisplay(&Display{})
  iface.AssertDragContext(&DragContext{})
  iface.AssertEvent(&Event{})
  iface.AssertEventButton(&EventButton{})
  iface.AssertEventKey(&EventKey{})
  iface.AssertEventMotion(&EventMotion{})
  iface.AssertEventScroll(&EventScroll{})
  iface.AssertPixbuf(&Pixbuf{})
  iface.AssertPixbufLoader(&PixbufLoader{})
  iface.AssertRGBA(&RGBA{})
  iface.AssertRectangle(&Rectangle{})
  iface.AssertScreen(&Screen{})
  iface.AssertVisual(&Visual{})
  iface.AssertWindow(&Window{})
}

package cairo

import "github.com/gotk3/gotk3/cairo/iface"

func init() {
  iface.AssertCairo(&RealCairo{})
  iface.AssertContext(&Context{})
  iface.AssertSurface(&Surface{})
}

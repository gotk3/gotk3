package gtk

import "github.com/gotk3/gotk3/gdk"

type Allocation interface {
	gdk.Rectangle
} // end of Allocation

func AssertAllocation(_ Allocation) {}

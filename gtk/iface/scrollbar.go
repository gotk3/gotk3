package iface


type Scrollbar interface {
    Range
} // end of Scrollbar

func AssertScrollbar(_ Scrollbar) {}

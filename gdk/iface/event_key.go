package iface


type EventKey interface {
    Event

    KeyVal() uint
    State() uint
    Type() EventType
} // end of EventKey

func AssertEventKey(_ EventKey) {}

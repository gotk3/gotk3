package iface


type Signal interface {
    String() string
} // end of Signal

func AssertSignal(_ Signal) {}

package iface


type Scale interface {
    Range
} // end of Scale

func AssertScale(_ Scale) {}

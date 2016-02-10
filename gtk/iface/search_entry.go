package iface


type SearchEntry interface {
    Entry
} // end of SearchEntry

func AssertSearchEntry(_ SearchEntry) {}

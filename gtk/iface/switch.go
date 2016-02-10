package iface


type Switch interface {
    Widget

    GetActive() bool
    SetActive(bool)
} // end of Switch

func AssertSwitch(_ Switch) {}

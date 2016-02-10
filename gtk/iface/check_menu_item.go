package iface


type CheckMenuItem interface {
    MenuItem

    GetActive() bool
    GetDrawAsRadio() bool
    GetInconsistent() bool
    SetActive(bool)
    SetDrawAsRadio(bool)
    SetInconsistent(bool)
    Toggled()
} // end of CheckMenuItem

func AssertCheckMenuItem(_ CheckMenuItem) {}

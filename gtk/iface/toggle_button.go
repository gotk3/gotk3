package iface


type ToggleButton interface {
    Button

    GetActive() bool
    SetActive(bool)
} // end of ToggleButton

func AssertToggleButton(_ ToggleButton) {}

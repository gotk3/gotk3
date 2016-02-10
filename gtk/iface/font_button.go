package iface


type FontButton interface {
    Button

    GetFontName() string
    SetFontName(string) bool
} // end of FontButton

func AssertFontButton(_ FontButton) {}

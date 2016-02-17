package pango

type Color interface {
	Copy(Color) Color
	Free()
	Get() (uint16, uint16, uint16)
	Parse(string) bool
	Set(uint16, uint16, uint16)
	ToString() string
} // end of Color

func AssertColor(_ Color) {}

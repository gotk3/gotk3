package gdk

type RGBA interface {
	Floats() []float64
	Parse(string) bool
	String() string
} // end of RGBA

func AssertRGBA(_ RGBA) {}

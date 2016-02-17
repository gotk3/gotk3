package gtk

type ApplicationWindow interface {
	Window

	GetID() uint
	GetShowMenubar() bool
	SetShowMenubar(bool)
} // end of ApplicationWindow

func AssertApplicationWindow(_ ApplicationWindow) {}

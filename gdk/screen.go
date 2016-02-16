package gdk

import "github.com/gotk3/gotk3/glib"

type Screen interface {
	glib.Object

	GetActiveWindow() (Window, error)
	GetCurrentDesktop() uint32
	GetDisplay() (Display, error)
	GetHeight() int
	GetHeightMM() int
	GetMonitorAtPoint(int, int) int
	GetMonitorAtWindow(Window) int
	GetMonitorHeightMM(int) int
	GetMonitorPlugName(int) (string, error)
	GetMonitorScaleFactor(int) int
	GetMonitorWidthMM(int) int
	GetNMonitors() int
	GetNumber() int
	GetNumberOfDesktops() uint32
	GetPrimaryMonitor() int
	GetRGBAVisual() (Visual, error)
	GetResolution() float64
	GetRootWindow() (Window, error)
	GetScreenNumber() int
	GetSystemVisual() (Visual, error)
	GetWidth() int
	GetWidthMM() int
	IsComposited() bool
	MakeDisplayName() (string, error)
	SetResolution(float64)
} // end of Screen

func AssertScreen(_ Screen) {}

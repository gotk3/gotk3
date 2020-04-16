// Same copyright and license as the rest of the files in this project

// Package gtk_test is a separate package for unit tests.
// Doing this actually utilizes the go build cache
// for package gtk when changing unit test code.
package gtk_test

import (
	"testing"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func TestWindowNew(t *testing.T) {
	cut, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		t.Error("unexpected error:", err.Error())
	}

	actual := cut.GetWindowType()
	if gtk.WINDOW_TOPLEVEL != actual {
		t.Errorf("Expected WindowType '%d'; Got '%d'", gtk.WINDOW_TOPLEVEL, actual)
	}
}

func createTestWindow(t *testing.T) *gtk.Window {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		t.Error("unexpected error:", err.Error())
	}

	return win
}

func TestWindowGetSetTitle(t *testing.T) {
	win := createTestWindow(t)

	expected := "Unit Test Window"
	win.SetTitle(expected)

	actual, err := win.GetTitle()
	if err != nil {
		t.Error("unexpected error:", err.Error())
	}
	if expected != actual {
		t.Errorf("Expected '%s'; Got '%s'", expected, actual)
	}
}

func TestWindowGetSetIconName(t *testing.T) {
	win := createTestWindow(t)

	expected := "unit-icon-symbolic"
	win.SetIconName(expected)

	actual, err := win.GetIconName()
	if err != nil {
		t.Error("unexpected error:", err.Error())
	}
	if expected != actual {
		t.Errorf("Expected '%s'; Got '%s'", expected, actual)
	}
}

func TestWindowGetSetDefaultIconName(t *testing.T) {
	expected := "unit-icon-symbolic"
	gtk.WindowSetDefaultIconName(expected)

	actual, err := gtk.WindowGetDefaultIconName()
	if err != nil {
		t.Error("unexpected error:", err.Error())
	}
	if expected != actual {
		t.Errorf("Expected '%s'; Got '%s'", expected, actual)
	}
}

func TestWindowGetSetRole(t *testing.T) {
	win := createTestWindow(t)

	expected := "Unit Test Role"
	win.SetRole(expected)

	actual, err := win.GetRole()
	if err != nil {
		t.Error("unexpected error:", err.Error())
	}
	if expected != actual {
		t.Errorf("Expected '%s'; Got '%s'", expected, actual)
	}
}

func TestWindowGetSetTransientFor(t *testing.T) {
	win := createTestWindow(t)

	expected := createTestWindow(t)
	win.SetTransientFor(expected)

	actual, err := win.GetTransientFor()
	if err != nil {
		t.Error("unexpected error:", err.Error())
	}
	if expected.Native() != actual.Native() {
		t.Errorf("Expected '0x%x'; Got '0x%x'", expected.Native(), actual.Native())
	}
}

func TestWindowGetSetAttachedTo(t *testing.T) {
	win := createTestWindow(t)

	expected := createTestWindow(t)
	win.SetAttachedTo(expected)

	a, err := win.GetAttachedTo()
	if err != nil {
		t.Error("unexpected error:", err.Error())
	}
	actual := a.ToWidget()
	if expected.Native() != actual.Native() {
		t.Errorf("Expected '0x%x'; Got '0x%x'", expected.Native(), actual.Native())
	}
}

func TestWindowGetSetDefaultSize(t *testing.T) {
	win := createTestWindow(t)

	expectedW, expectedH := 123, 345
	win.SetDefaultSize(expectedW, expectedH)

	actualW, actualH := win.GetDefaultSize()

	if expectedW != actualW || expectedH != actualH {
		t.Errorf("Expected %dx%d; Got %dx%d", expectedW, expectedH, actualW, actualH)
	}
}

func TestWindowGetSetGravity(t *testing.T) {
	win := createTestWindow(t)

	var expected gdk.Gravity
	expected = gdk.GDK_GRAVITY_EAST
	win.SetGravity(expected)

	actual := win.GetGravity()
	if expected != actual {
		t.Errorf("Expected '%d'; Got '%d'", expected, actual)
	}
}

func TestWindowGetSetDefaultWidget(t *testing.T) {
	win := createTestWindow(t)

	// Create test button, SetCanDefault is required for SetDefault to work
	expected, err := gtk.ButtonNew()
	if err != nil {
		t.Error("unexpected error:", err.Error())
	}
	expected.SetCanDefault(true)

	win.SetDefault(expected)

	a, err := win.GetDefaultWidget()
	if err != nil {
		t.Error("unexpected error:", err.Error())
	}
	actual := a.ToWidget()
	if expected.Native() != actual.Native() {
		t.Errorf("Expected '0x%x'; Got '0x%x'", expected.Native(), actual.Native())
	}
}

func TestWindowGetSetDestroyWithParent(t *testing.T) {
	win := createTestWindow(t)

	testCases := []struct {
		desc  string
		value bool
	}{
		{
			desc:  "true",
			value: true,
		},
		{
			desc:  "true",
			value: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			win.SetDestroyWithParent(tC.value)

			actual := win.GetDestroyWithParent()
			if tC.value != actual {
				t.Errorf("Expected '%t'; Got '%t'", tC.value, actual)
			}
		})
	}
}

func TestWindowGetSetHideTitlebarWhenMaximized(t *testing.T) {
	win := createTestWindow(t)

	testCases := []struct {
		desc  string
		value bool
	}{
		{
			desc:  "true",
			value: true,
		},
		{
			desc:  "true",
			value: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			win.SetHideTitlebarWhenMaximized(tC.value)

			actual := win.GetHideTitlebarWhenMaximized()
			if tC.value != actual {
				t.Errorf("Expected '%t'; Got '%t'", tC.value, actual)
			}
		})
	}
}

func TestWindowGetSetResizable(t *testing.T) {
	win := createTestWindow(t)

	testCases := []struct {
		desc  string
		value bool
	}{
		{
			desc:  "true",
			value: true,
		},
		{
			desc:  "true",
			value: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			win.SetResizable(tC.value)

			actual := win.GetResizable()
			if tC.value != actual {
				t.Errorf("Expected '%t'; Got '%t'", tC.value, actual)
			}
		})
	}
}

func TestWindowGetSetModal(t *testing.T) {
	win := createTestWindow(t)

	testCases := []struct {
		desc  string
		value bool
	}{
		{
			desc:  "true",
			value: true,
		},
		{
			desc:  "true",
			value: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			win.SetModal(tC.value)

			actual := win.GetModal()
			if tC.value != actual {
				t.Errorf("Expected '%t'; Got '%t'", tC.value, actual)
			}
		})
	}
}

func TestWindowGetSetDecorated(t *testing.T) {
	win := createTestWindow(t)

	testCases := []struct {
		desc  string
		value bool
	}{
		{
			desc:  "true",
			value: true,
		},
		{
			desc:  "true",
			value: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			win.SetDecorated(tC.value)

			actual := win.GetDecorated()
			if tC.value != actual {
				t.Errorf("Expected '%t'; Got '%t'", tC.value, actual)
			}
		})
	}
}

func TestWindowGetSetDeletable(t *testing.T) {
	win := createTestWindow(t)

	testCases := []struct {
		desc  string
		value bool
	}{
		{
			desc:  "true",
			value: true,
		},
		{
			desc:  "true",
			value: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			win.SetDeletable(tC.value)

			actual := win.GetDeletable()
			if tC.value != actual {
				t.Errorf("Expected '%t'; Got '%t'", tC.value, actual)
			}
		})
	}
}

func TestWindowGetSetSkipTaskbarHint(t *testing.T) {
	win := createTestWindow(t)

	testCases := []struct {
		desc  string
		value bool
	}{
		{
			desc:  "true",
			value: true,
		},
		{
			desc:  "true",
			value: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			win.SetSkipTaskbarHint(tC.value)

			actual := win.GetSkipTaskbarHint()
			if tC.value != actual {
				t.Errorf("Expected '%t'; Got '%t'", tC.value, actual)
			}
		})
	}
}

func TestWindowGetSetSkipPagerHint(t *testing.T) {
	win := createTestWindow(t)

	testCases := []struct {
		desc  string
		value bool
	}{
		{
			desc:  "true",
			value: true,
		},
		{
			desc:  "true",
			value: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			win.SetSkipPagerHint(tC.value)

			actual := win.GetSkipPagerHint()
			if tC.value != actual {
				t.Errorf("Expected '%t'; Got '%t'", tC.value, actual)
			}
		})
	}
}

func TestWindowGetSetUrgencyHint(t *testing.T) {
	win := createTestWindow(t)

	testCases := []struct {
		desc  string
		value bool
	}{
		{
			desc:  "true",
			value: true,
		},
		{
			desc:  "true",
			value: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			win.SetUrgencyHint(tC.value)

			actual := win.GetUrgencyHint()
			if tC.value != actual {
				t.Errorf("Expected '%t'; Got '%t'", tC.value, actual)
			}
		})
	}
}

func TestWindowGetSetAcceptFocus(t *testing.T) {
	win := createTestWindow(t)

	testCases := []struct {
		desc  string
		value bool
	}{
		{
			desc:  "true",
			value: true,
		},
		{
			desc:  "true",
			value: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			win.SetAcceptFocus(tC.value)

			actual := win.GetAcceptFocus()
			if tC.value != actual {
				t.Errorf("Expected '%t'; Got '%t'", tC.value, actual)
			}
		})
	}
}

func TestWindowGetSetFocusOnMap(t *testing.T) {
	win := createTestWindow(t)

	testCases := []struct {
		desc  string
		value bool
	}{
		{
			desc:  "true",
			value: true,
		},
		{
			desc:  "true",
			value: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			win.SetFocusOnMap(tC.value)

			actual := win.GetFocusOnMap()
			if tC.value != actual {
				t.Errorf("Expected '%t'; Got '%t'", tC.value, actual)
			}
		})
	}
}

func TestWindowGetSetMnemonicsVisible(t *testing.T) {
	win := createTestWindow(t)

	testCases := []struct {
		desc  string
		value bool
	}{
		{
			desc:  "true",
			value: true,
		},
		{
			desc:  "true",
			value: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			win.SetMnemonicsVisible(tC.value)

			actual := win.GetMnemonicsVisible()
			if tC.value != actual {
				t.Errorf("Expected '%t'; Got '%t'", tC.value, actual)
			}
		})
	}
}

func TestWindowGetSetFocusVisible(t *testing.T) {
	win := createTestWindow(t)

	testCases := []struct {
		desc  string
		value bool
	}{
		{
			desc:  "true",
			value: true,
		},
		{
			desc:  "true",
			value: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			win.SetFocusVisible(tC.value)

			actual := win.GetFocusVisible()
			if tC.value != actual {
				t.Errorf("Expected '%t'; Got '%t'", tC.value, actual)
			}
		})
	}
}

func TestWindowGetSetTypeHint(t *testing.T) {
	win := createTestWindow(t)

	var expected gdk.WindowTypeHint
	expected = gdk.WINDOW_TYPE_HINT_UTILITY
	win.SetTypeHint(expected)

	actual := win.GetTypeHint()
	if expected != actual {
		t.Errorf("Expected '%d'; Got '%d'", expected, actual)
	}
}

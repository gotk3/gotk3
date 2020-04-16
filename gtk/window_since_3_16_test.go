// Same copyright and license as the rest of the files in this project

// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14

// Package gtk_test is a separate package for unit tests.
// Doing this actually utilizes the go build cache
// for package gtk when changing unit test code.
package gtk_test

import (
	"testing"

	"github.com/gotk3/gotk3/gtk"
)

func TestWindowGetSetTitlebar(t *testing.T) {
	win := createTestWindow(t)

	// Create test button, SetCanDefault is required for SetDefault to work
	expected, err := gtk.ButtonNew()
	if err != nil {
		t.Error("unexpected error:", err.Error())
	}
	win.SetTitlebar(expected)

	a, err := win.GetTitlebar()
	if err != nil {
		t.Error("unexpected cast failure:", err.Error())
	}
	actual := a.ToWidget()
	if expected.Native() != actual.Native() {
		t.Errorf("Expected '0x%x'; Got '0x%x'", expected.Native(), actual.Native())
	}
}

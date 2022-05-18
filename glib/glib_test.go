package glib_test

import (
	"testing"

	"github.com/gotk3/gotk3/glib"
)

// TestConnectSignal tests that specific callback connected to the signal.
func TestConnectSignal(t *testing.T) {
	ctx := glib.MainContextDefault()
	mainLoop := glib.MainLoopNew(ctx, true)

	// Create any GObject that has defined properties.
	obj, _ := glib.CancellableNew()

	// Connect to a "notify::" signal to listen on property changes.
	obj.Connect("cancelled", func() {
		mainLoop.Quit()
	})

	glib.IdleAdd(func() bool {
		obj.Cancel()
		return false
	})

	mainLoop.Run()
}

// TestTypeNames tests both glib.TypeFromName and glib.Type.Name
func TestTypeNames(t *testing.T) {
	tp := glib.TypeFromName("GObject")
	name := tp.Name()

	if name != "GObject" {
		t.Error("Expected GObject, got", name)
	}
}

func TestTypeIsA(t *testing.T) {
	tp := glib.TypeFromName("GCancellable")
	tpParent := glib.TypeFromName("GObject")

	isA := tp.IsA(tpParent)

	if !isA {
		t.Error("Expected true, GCancellable is a GObject")
	}
}

func TestTypeNextBase(t *testing.T) {
	// http://manual.freeshell.org/glibmm-2.4/reference/html/classGlib_1_1ObjectBase.html
	tpLeaf := glib.TypeFromName("GFileInputStream")
	tpParent := glib.TypeFromName("GObject")

	tpNextBase := glib.TypeNextBase(tpLeaf, tpParent)
	name := tpNextBase.Name()

	if name != "GInputStream" {
		t.Error("Expected GInputStream, got", name)
	}
}

func TestValueString_NonEmpty(t *testing.T) {
	expected := "test"
	value, err := glib.GValue(expected)
	if err != nil {
		t.Error("acquiring gvalue failed:", err.Error())
		return
	}

	actual, err := value.GetString()
	if err != nil {
		t.Error(err.Error())
		return
	}

	if actual != expected {
		t.Errorf("Expected %q, got %q", expected, actual)
	}
}

func TestValueString_Empty(t *testing.T) {
	expected := ""
	value, err := glib.GValue(expected)
	if err != nil {
		t.Error("acquiring gvalue failed:", err.Error())
		return
	}

	actual, err := value.GetString()
	if err != nil {
		t.Error(err.Error())
		return
	}

	if actual != expected {
		t.Errorf("Expected %q, got %q", expected, actual)
	}
}

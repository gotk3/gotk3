package glib_test

import (
	"runtime"
	"testing"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func init() {
	gtk.Init(nil)
}

// TestConnectNotifySignal ensures that property notification signals (those
// whose name begins with "notify::") are queried by the name "notify" (with the
// "::" and the property name omitted). This is because the signal is "notify"
// and the characters after the "::" are not recognized by the signal system.
//
// See
// https://developer.gnome.org/gobject/stable/gobject-The-Base-Object-Type.html#GObject-notify
// for background, and
// https://developer.gnome.org/gobject/stable/gobject-Signals.html#g-signal-new
// for the specification of valid signal names.
func TestConnectNotifySignal(t *testing.T) {
	runtime.LockOSThread()

	// Create any GObject that has defined properties.
	spacing := 0
	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, spacing)

	// Connect to a "notify::" signal to listen on property changes.
	box.Connect("notify::spacing", func() {
		gtk.MainQuit()
	})

	glib.IdleAdd(func(s string) bool {
		t.Log(s)
		spacing++
		box.SetSpacing(spacing)
		return true
	}, "IdleAdd executed")

	gtk.Main()
}

/*At this moment Visionect specific*/
func TestTimeoutAdd(t *testing.T) {
	runtime.LockOSThread()

	glib.TimeoutAdd(2500, func(s string) bool {
		t.Log(s)
		gtk.MainQuit()
		return false
	}, "TimeoutAdd executed")

	gtk.Main()
}

// TestTypeNames tests both glib.TypeFromName and glib.Type.Name
func TestTypeNames(t *testing.T) {
	tp := glib.TypeFromName("GtkWindow")
	name := tp.Name()

	if name != "GtkWindow" {
		t.Error("Expected GtkWindow, got", name)
	}
}

func TestTypeIsA(t *testing.T) {
	tp := glib.TypeFromName("GtkApplicationWindow")
	tpParent := glib.TypeFromName("GtkWindow")

	isA := tp.IsA(tpParent)

	if !isA {
		t.Error("Expected true, GtkApplicationWindow is a GtkWindow")
	}
}

func TestTypeNextBase(t *testing.T) {
	tpLeaf := glib.TypeFromName("GtkWindow")
	tpParent := glib.TypeFromName("GtkContainer")

	tpNextBase := glib.TypeNextBase(tpLeaf, tpParent)
	name := tpNextBase.Name()

	if name != "GtkBin" {
		t.Error("Expected GtkBin, got", name)
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


package impl_test

import (
	"runtime"
	"testing"

	glib_impl "github.com/gotk3/gotk3/glib/impl"
	"github.com/gotk3/gotk3/gtk"
	gtk_impl "github.com/gotk3/gotk3/gtk/impl"
)

func init() {
	gtk_impl.Init(nil)
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
	box, _ := gtk_impl.BoxNew(gtk.ORIENTATION_HORIZONTAL, spacing)

	// Connect to a "notify::" signal to listen on property changes.
	box.Connect("notify::spacing", func() {
		gtk_impl.MainQuit()
	})

	glib_impl.IdleAdd(func(s string) bool {
		t.Log(s)
		spacing++
		box.SetSpacing(spacing)
		return true
	}, "IdleAdd executed")

	gtk_impl.Main()
}

/*At this moment Visionect specific*/
func TestTimeoutAdd(t *testing.T) {
	runtime.LockOSThread()

	glib_impl.TimeoutAdd(2500, func(s string) bool {
		t.Log(s)
		gtk_impl.MainQuit()
		return false
	}, "TimeoutAdd executed")

	gtk_impl.Main()
}

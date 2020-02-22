package gdk

// #include <gdk/gdk.h>
import "C"

// TestRenderSync retrieves a pixel from window to force the windowing system to carry out any pending rendering commands.
// This function is intended to be used to synchronize with rendering pipelines, to benchmark windowing system rendering operations.
// This is a wrapper around gdk_test_render_sync().
func TestRenderSync(window *Window) {
	C.gdk_test_render_sync(window.native())
}

// TestSimulateButton simulates a single mouse button event (press or release) at the given coordinates relative to the window.
// Hint: a single click of a button requires this method to be called twice, once for pressed and once for released.
// In most cases, gtk.TestWidgetClick() should be used.
//
// button: Mouse button number, starts with 0
// modifiers: Keyboard modifiers for the button event
// buttonPressRelease: either GDK_BUTTON_PRESS or GDK_BUTTON_RELEASE
//
// This is a wrapper around gdk_test_simulate_button().
func TestSimulateButton(window *Window, x, y int, button Button, modifiers ModifierType, buttonPressRelease EventType) bool {
	return gobool(C.gdk_test_simulate_button(window.native(), C.gint(x), C.gint(y), C.guint(button), C.GdkModifierType(modifiers), C.GdkEventType(buttonPressRelease)))
}

// TestSimulateButton simulates a keyboard event (press or release) at the given coordinates relative to the window.
// If the coordinates (-1, -1) are used, the window origin is used instead.
// Hint: a single key press requires this method to be called twice, once for pressed and once for released.
// In most cases, gtk.TestWidgetSendKey() should be used.
//
// keyval: A GDK keyboard value (See KeyvalFromName(), UnicodeToKeyval(), ...)
// modifiers: Keyboard modifiers for the key event
// buttonPressRelease: either GDK_BUTTON_PRESS or GDK_BUTTON_RELEASE
//
// This is a wrapper around gdk_test_simulate_key().
func TestSimulateKey(window *Window, x, y int, keyval uint, modifiers ModifierType, buttonPressRelease EventType) bool {
	return gobool(C.gdk_test_simulate_key(window.native(), C.gint(x), C.gint(y), C.guint(keyval), C.GdkModifierType(modifiers), C.GdkEventType(buttonPressRelease)))
}

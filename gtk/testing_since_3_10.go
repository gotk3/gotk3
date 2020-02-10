// +build !gtk_3_6,!gtk_3_8

package gtk

// #include <gtk/gtk.h>
import "C"

// TestWidgetWaitForDraw is a wrapper around gtk_test_widget_wait_for_draw().
// Enters the main loop and waits for widget to be “drawn”. In this context that means it waits for the frame clock of widget to have run a full styling, layout and drawing cycle.
// This function is intended to be used for syncing with actions that depend on widget relayouting or on interaction with the display server.
func TestWidgetWaitForDraw(widget IWidget) {
	C.gtk_test_widget_wait_for_draw(widget.toWidget())
}

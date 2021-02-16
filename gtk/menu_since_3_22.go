// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18,!gtk_3_20

package gtk

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
)

// PopupAtRect is a wrapper around gtk_menu_popup_at_rect().
func (v *Menu) PopupAtRect(rect_window *gdk.Window,
	rect *gdk.Rectangle, rect_anchor, menu_anchor gdk.Gravity,
	trigger_event *gdk.Event) {

	C.gtk_menu_popup_at_rect(
		v.native(),
		(*C.GdkWindow)(unsafe.Pointer(rect_window.Native())),
		(*C.GdkRectangle)(unsafe.Pointer(&rect.GdkRectangle)),
		C.GdkGravity(rect_anchor),
		C.GdkGravity(menu_anchor),
		(*C.GdkEvent)(unsafe.Pointer(trigger_event.Native())))
}

// PopupAtWidget() is a wrapper for gtk_menu_popup_at_widget()
func (v *Menu) PopupAtWidget(widget IWidget, widgetAnchor gdk.Gravity, menuAnchor gdk.Gravity, triggerEvent *gdk.Event) {
	e := (*C.GdkEvent)(unsafe.Pointer(triggerEvent.Native()))
	C.gtk_menu_popup_at_widget(v.native(), widget.toWidget(), C.GdkGravity(widgetAnchor), C.GdkGravity(menuAnchor), e)
}

// PopupAtPointer() is a wrapper for gtk_menu_popup_at_pointer(), on older versions it uses PopupAtMouseCursor
func (v *Menu) PopupAtPointer(triggerEvent *gdk.Event) {
	e := (*C.GdkEvent)(unsafe.Pointer(triggerEvent.Native()))
	C.gtk_menu_popup_at_pointer(v.native(), e)
}

// PlaceOnMonitor() is a wrapper around gtk_menu_place_on_monitor().
func (v *Menu) PlaceOnMonitor(monitor *gdk.Monitor) {
	C.gtk_menu_place_on_monitor(
		v.native(),
		(*C.GdkMonitor)(unsafe.Pointer(monitor.Native())))
}

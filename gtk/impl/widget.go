// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

package impl

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"errors"
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	gdk_impl "github.com/gotk3/gotk3/gdk/impl"
	glib_impl "github.com/gotk3/gotk3/glib/impl"
	"github.com/gotk3/gotk3/gtk"
)

/*
 * GtkWidget
 */

// Widget is a representation of GTK's GtkWidget.
type widget struct {
	glib_impl.InitiallyUnowned
}

// IWidget is an interface type implemented by all structs
// embedding a Widget.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkWidget.
type IWidget interface {
	toWidget() *C.GtkWidget
	Set(string, interface{}) error
}

type IWidgetable interface {
	asWidgetImpl() *widget
	toWidget() *C.GtkWidget
}

func nullableWidget(v interface{}) *C.GtkWidget {
	if v == nil {
		return nil
	}

	return v.(IWidgetable).toWidget()
}

// native returns a pointer to the underlying GtkWidget.
func (v *widget) native() *C.GtkWidget {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkWidget(p)
}

func asWidgetImpl(v gtk.Widget) *widget {
	if v == nil {
		return nil
	}
	return v.(IWidgetable).asWidgetImpl()
}

func (v *widget) asWidgetImpl() *widget {
	if v == nil {
		return nil
	}
	return v
}

func (v *widget) toWidget() *C.GtkWidget {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalWidget(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

func wrapWidget(obj *glib_impl.Object) *widget {
	return &widget{glib_impl.InitiallyUnowned{obj}}
}

// Destroy is a wrapper around gtk_widget_destroy().
func (v *widget) Destroy() {
	C.gtk_widget_destroy(v.native())
}

func (v *widget) HideOnDelete() {
	C._gtk_widget_hide_on_delete(v.native())
}

/* TODO
func (v *widget) DragDestSet(flags DestDefaults, targets []TargetEntry, actions gdk_impl.DragAction) {
	C.gtk_drag_dest_set(v.native(), C.GtkDestDefaults(flags), (*C.GtkTargetEntry)(&targets[0]),
		C.gint(len(targets)), C.GdkDragAction(actions))
}
*/

// ResetStyle is a wrapper around gtk_widget_reset_style().
func (v *widget) ResetStyle() {
	C.gtk_widget_reset_style(v.native())
}

// InDestruction is a wrapper around gtk_widget_in_destruction().
func (v *widget) InDestruction() bool {
	return gobool(C.gtk_widget_in_destruction(v.native()))
}

// TODO(jrick) this may require some rethinking
/*
func (v *widget) Destroyed(widgetPointer **widget) {
}
*/

// Unparent is a wrapper around gtk_widget_unparent().
func (v *widget) Unparent() {
	C.gtk_widget_unparent(v.native())
}

// Show is a wrapper around gtk_widget_show().
func (v *widget) Show() {
	C.gtk_widget_show(v.native())
}

// Hide is a wrapper around gtk_widget_hide().
func (v *widget) Hide() {
	C.gtk_widget_hide(v.native())
}

// GetCanFocus is a wrapper around gtk_widget_get_can_focus().
func (v *widget) GetCanFocus() bool {
	c := C.gtk_widget_get_can_focus(v.native())
	return gobool(c)
}

// SetCanFocus is a wrapper around gtk_widget_set_can_focus().
func (v *widget) SetCanFocus(canFocus bool) {
	C.gtk_widget_set_can_focus(v.native(), gbool(canFocus))
}

// GetCanDefault is a wrapper around gtk_widget_get_can_default().
func (v *widget) GetCanDefault() bool {
	c := C.gtk_widget_get_can_default(v.native())
	return gobool(c)
}

// SetCanDefault is a wrapper around gtk_widget_set_can_default().
func (v *widget) SetCanDefault(canDefault bool) {
	C.gtk_widget_set_can_default(v.native(), gbool(canDefault))
}

// GetMapped is a wrapper around gtk_widget_get_mapped().
func (v *widget) GetMapped() bool {
	c := C.gtk_widget_get_mapped(v.native())
	return gobool(c)
}

// SetMapped is a wrapper around gtk_widget_set_mapped().
func (v *widget) SetMapped(mapped bool) {
	C.gtk_widget_set_can_focus(v.native(), gbool(mapped))
}

// GetRealized is a wrapper around gtk_widget_get_realized().
func (v *widget) GetRealized() bool {
	c := C.gtk_widget_get_realized(v.native())
	return gobool(c)
}

// SetRealized is a wrapper around gtk_widget_set_realized().
func (v *widget) SetRealized(realized bool) {
	C.gtk_widget_set_realized(v.native(), gbool(realized))
}

// GetHasWindow is a wrapper around gtk_widget_get_has_window().
func (v *widget) GetHasWindow() bool {
	c := C.gtk_widget_get_has_window(v.native())
	return gobool(c)
}

// SetHasWindow is a wrapper around gtk_widget_set_has_window().
func (v *widget) SetHasWindow(hasWindow bool) {
	C.gtk_widget_set_has_window(v.native(), gbool(hasWindow))
}

// ShowNow is a wrapper around gtk_widget_show_now().
func (v *widget) ShowNow() {
	C.gtk_widget_show_now(v.native())
}

// ShowAll is a wrapper around gtk_widget_show_all().
func (v *widget) ShowAll() {
	C.gtk_widget_show_all(v.native())
}

// SetNoShowAll is a wrapper around gtk_widget_set_no_show_all().
func (v *widget) SetNoShowAll(noShowAll bool) {
	C.gtk_widget_set_no_show_all(v.native(), gbool(noShowAll))
}

// GetNoShowAll is a wrapper around gtk_widget_get_no_show_all().
func (v *widget) GetNoShowAll() bool {
	c := C.gtk_widget_get_no_show_all(v.native())
	return gobool(c)
}

// Map is a wrapper around gtk_widget_map().
func (v *widget) Map() {
	C.gtk_widget_map(v.native())
}

// Unmap is a wrapper around gtk_widget_unmap().
func (v *widget) Unmap() {
	C.gtk_widget_unmap(v.native())
}

// QueueDrawArea is a wrapper aroung gtk_widget_queue_draw_area().
func (v *widget) QueueDrawArea(x, y, w, h int) {
	C.gtk_widget_queue_draw_area(v.native(), C.gint(x), C.gint(y), C.gint(w), C.gint(h))
}

//void gtk_widget_realize(GtkWidget *widget);
//void gtk_widget_unrealize(GtkWidget *widget);
//void gtk_widget_draw(GtkWidget *widget, cairo_t *cr);
//void gtk_widget_queue_resize(GtkWidget *widget);
//void gtk_widget_queue_resize_no_redraw(GtkWidget *widget);
//GdkFrameClock *gtk_widget_get_frame_clock(GtkWidget *widget);
//guint gtk_widget_add_tick_callback (GtkWidget *widget,
//                                    GtkTickCallback callback,
//                                    gpointer user_data,
//                                    GDestroyNotify notify);
//void gtk_widget_remove_tick_callback(GtkWidget *widget, guint id);

// TODO(jrick) GtkAllocation
/*
func (v *widget) SizeAllocate() {
}
*/

// Allocation is a representation of GTK's GtkAllocation type.
type allocation struct {
	gdk_impl.Rectangle
}

// Native returns a pointer to the underlying GtkAllocation.
func (v *allocation) native() *C.GtkAllocation {
	return (*C.GtkAllocation)(unsafe.Pointer(&v.GdkRectangle))
}

// GetAllocatedWidth() is a wrapper around gtk_widget_get_allocated_width().
func (v *widget) GetAllocatedWidth() int {
	return int(C.gtk_widget_get_allocated_width(v.native()))
}

// GetAllocatedHeight() is a wrapper around gtk_widget_get_allocated_height().
func (v *widget) GetAllocatedHeight() int {
	return int(C.gtk_widget_get_allocated_height(v.native()))
}

// Event() is a wrapper around gtk_widget_event().
func (v *widget) Event(event gdk.Event) bool {
	c := C.gtk_widget_event(v.native(),
		(*C.GdkEvent)(unsafe.Pointer(gdk_impl.CastToEvent(event).Native())))
	return gobool(c)
}

// Activate() is a wrapper around gtk_widget_activate().
func (v *widget) Activate() bool {
	return gobool(C.gtk_widget_activate(v.native()))
}

// TODO(jrick) GdkRectangle
/*
func (v *widget) Intersect() {
}
*/

// IsFocus() is a wrapper around gtk_widget_is_focus().
func (v *widget) IsFocus() bool {
	return gobool(C.gtk_widget_is_focus(v.native()))
}

// GrabFocus() is a wrapper around gtk_widget_grab_focus().
func (v *widget) GrabFocus() {
	C.gtk_widget_grab_focus(v.native())
}

// GrabDefault() is a wrapper around gtk_widget_grab_default().
func (v *widget) GrabDefault() {
	C.gtk_widget_grab_default(v.native())
}

// SetName() is a wrapper around gtk_widget_set_name().
func (v *widget) SetName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_widget_set_name(v.native(), (*C.gchar)(cstr))
}

// GetName() is a wrapper around gtk_widget_get_name().  A non-nil
// error is returned in the case that gtk_widget_get_name returns NULL to
// differentiate between NULL and an empty string.
func (v *widget) GetName() (string, error) {
	c := C.gtk_widget_get_name(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// GetSensitive is a wrapper around gtk_widget_get_sensitive().
func (v *widget) GetSensitive() bool {
	c := C.gtk_widget_get_sensitive(v.native())
	return gobool(c)
}

// IsSensitive is a wrapper around gtk_widget_is_sensitive().
func (v *widget) IsSensitive() bool {
	c := C.gtk_widget_is_sensitive(v.native())
	return gobool(c)
}

// SetSensitive is a wrapper around gtk_widget_set_sensitive().
func (v *widget) SetSensitive(sensitive bool) {
	C.gtk_widget_set_sensitive(v.native(), gbool(sensitive))
}

// GetVisible is a wrapper around gtk_widget_get_visible().
func (v *widget) GetVisible() bool {
	c := C.gtk_widget_get_visible(v.native())
	return gobool(c)
}

// SetVisible is a wrapper around gtk_widget_set_visible().
func (v *widget) SetVisible(visible bool) {
	C.gtk_widget_set_visible(v.native(), gbool(visible))
}

// SetParent is a wrapper around gtk_widget_set_parent().
func (v *widget) SetParent(parent gtk.Widget) {
	C.gtk_widget_set_parent(v.native(), parent.(IWidget).toWidget())
}

// GetParent is a wrapper around gtk_widget_get_parent().
func (v *widget) GetParent() (gtk.Widget, error) {
	c := C.gtk_widget_get_parent(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWidget(wrapObject(unsafe.Pointer(c))), nil
}

// SetSizeRequest is a wrapper around gtk_widget_set_size_request().
func (v *widget) SetSizeRequest(width, height int) {
	C.gtk_widget_set_size_request(v.native(), C.gint(width), C.gint(height))
}

// GetSizeRequest is a wrapper around gtk_widget_get_size_request().
func (v *widget) GetSizeRequest() (width, height int) {
	var w, h C.gint
	C.gtk_widget_get_size_request(v.native(), &w, &h)
	return int(w), int(h)
}

// SetParentWindow is a wrapper around gtk_widget_set_parent_window().
func (v *widget) SetParentWindow(parentWindow gdk.Window) {
	C.gtk_widget_set_parent_window(v.native(),
		(*C.GdkWindow)(unsafe.Pointer(gdk_impl.CastToWindow(parentWindow).Native())))
}

// GetParentWindow is a wrapper around gtk_widget_get_parent_window().
func (v *widget) GetParentWindow() (gdk.Window, error) {
	c := C.gtk_widget_get_parent_window(v.native())
	if v == nil {
		return nil, nilPtrErr
	}

	w := &gdk_impl.Window{wrapObject(unsafe.Pointer(c))}
	return w, nil
}

// SetEvents is a wrapper around gtk_widget_set_events().
func (v *widget) SetEvents(events int) {
	C.gtk_widget_set_events(v.native(), C.gint(events))
}

// GetEvents is a wrapper around gtk_widget_get_events().
func (v *widget) GetEvents() int {
	return int(C.gtk_widget_get_events(v.native()))
}

// AddEvents is a wrapper around gtk_widget_add_events().
func (v *widget) AddEvents(events int) {
	C.gtk_widget_add_events(v.native(), C.gint(events))
}

// HasDefault is a wrapper around gtk_widget_has_default().
func (v *widget) HasDefault() bool {
	c := C.gtk_widget_has_default(v.native())
	return gobool(c)
}

// HasFocus is a wrapper around gtk_widget_has_focus().
func (v *widget) HasFocus() bool {
	c := C.gtk_widget_has_focus(v.native())
	return gobool(c)
}

// HasVisibleFocus is a wrapper around gtk_widget_has_visible_focus().
func (v *widget) HasVisibleFocus() bool {
	c := C.gtk_widget_has_visible_focus(v.native())
	return gobool(c)
}

// HasGrab is a wrapper around gtk_widget_has_grab().
func (v *widget) HasGrab() bool {
	c := C.gtk_widget_has_grab(v.native())
	return gobool(c)
}

// IsDrawable is a wrapper around gtk_widget_is_drawable().
func (v *widget) IsDrawable() bool {
	c := C.gtk_widget_is_drawable(v.native())
	return gobool(c)
}

// IsToplevel is a wrapper around gtk_widget_is_toplevel().
func (v *widget) IsToplevel() bool {
	c := C.gtk_widget_is_toplevel(v.native())
	return gobool(c)
}

// TODO(jrick) GdkEventMask
/*
func (v *widget) SetDeviceEvents() {
}
*/

// TODO(jrick) GdkEventMask
/*
func (v *widget) GetDeviceEvents() {
}
*/

// TODO(jrick) GdkEventMask
/*
func (v *widget) AddDeviceEvents() {
}
*/

// SetDeviceEnabled is a wrapper around gtk_widget_set_device_enabled().
func (v *widget) SetDeviceEnabled(device gdk.Device, enabled bool) {
	C.gtk_widget_set_device_enabled(v.native(),
		(*C.GdkDevice)(unsafe.Pointer(gdk_impl.CastToDevice(device).Native())), gbool(enabled))
}

// GetDeviceEnabled is a wrapper around gtk_widget_get_device_enabled().
func (v *widget) GetDeviceEnabled(device gdk.Device) bool {
	c := C.gtk_widget_get_device_enabled(v.native(),
		(*C.GdkDevice)(unsafe.Pointer(gdk_impl.CastToDevice(device).Native())))
	return gobool(c)
}

// GetToplevel is a wrapper around gtk_widget_get_toplevel().
func (v *widget) GetToplevel() (gtk.Widget, error) {
	c := C.gtk_widget_get_toplevel(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWidget(wrapObject(unsafe.Pointer(c))), nil
}

// GetTooltipText is a wrapper around gtk_widget_get_tooltip_text().
// A non-nil error is returned in the case that
// gtk_widget_get_tooltip_text returns NULL to differentiate between NULL
// and an empty string.
func (v *widget) GetTooltipText() (string, error) {
	c := C.gtk_widget_get_tooltip_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// SetTooltipText is a wrapper around gtk_widget_set_tooltip_text().
func (v *widget) SetTooltipText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_widget_set_tooltip_text(v.native(), (*C.gchar)(cstr))
}

// GetHAlign is a wrapper around gtk_widget_get_halign().
func (v *widget) GetHAlign() gtk.Align {
	c := C.gtk_widget_get_halign(v.native())
	return gtk.Align(c)
}

// SetHAlign is a wrapper around gtk_widget_set_halign().
func (v *widget) SetHAlign(align gtk.Align) {
	C.gtk_widget_set_halign(v.native(), C.GtkAlign(align))
}

// GetVAlign is a wrapper around gtk_widget_get_valign().
func (v *widget) GetVAlign() gtk.Align {
	c := C.gtk_widget_get_valign(v.native())
	return gtk.Align(c)
}

// SetVAlign is a wrapper around gtk_widget_set_valign().
func (v *widget) SetVAlign(align gtk.Align) {
	C.gtk_widget_set_valign(v.native(), C.GtkAlign(align))
}

// GetMarginTop is a wrapper around gtk_widget_get_margin_top().
func (v *widget) GetMarginTop() int {
	c := C.gtk_widget_get_margin_top(v.native())
	return int(c)
}

// SetMarginTop is a wrapper around gtk_widget_set_margin_top().
func (v *widget) SetMarginTop(margin int) {
	C.gtk_widget_set_margin_top(v.native(), C.gint(margin))
}

// GetMarginBottom is a wrapper around gtk_widget_get_margin_bottom().
func (v *widget) GetMarginBottom() int {
	c := C.gtk_widget_get_margin_bottom(v.native())
	return int(c)
}

// SetMarginBottom is a wrapper around gtk_widget_set_margin_bottom().
func (v *widget) SetMarginBottom(margin int) {
	C.gtk_widget_set_margin_bottom(v.native(), C.gint(margin))
}

// GetHExpand is a wrapper around gtk_widget_get_hexpand().
func (v *widget) GetHExpand() bool {
	c := C.gtk_widget_get_hexpand(v.native())
	return gobool(c)
}

// SetHExpand is a wrapper around gtk_widget_set_hexpand().
func (v *widget) SetHExpand(expand bool) {
	C.gtk_widget_set_hexpand(v.native(), gbool(expand))
}

// GetVExpand is a wrapper around gtk_widget_get_vexpand().
func (v *widget) GetVExpand() bool {
	c := C.gtk_widget_get_vexpand(v.native())
	return gobool(c)
}

// SetVExpand is a wrapper around gtk_widget_set_vexpand().
func (v *widget) SetVExpand(expand bool) {
	C.gtk_widget_set_vexpand(v.native(), gbool(expand))
}

// TranslateCoordinates is a wrapper around gtk_widget_translate_coordinates().
func (v *widget) TranslateCoordinates(dest gtk.Widget, srcX, srcY int) (destX, destY int, e error) {
	cdest := nullableWidget(dest)

	var cdestX, cdestY C.gint
	c := C.gtk_widget_translate_coordinates(v.native(), cdest, C.gint(srcX), C.gint(srcY), &cdestX, &cdestY)
	if !gobool(c) {
		return 0, 0, errors.New("translate coordinates failed")
	}
	return int(cdestX), int(cdestY), nil
}

// SetVisual is a wrapper around gtk_widget_set_visual().
func (v *widget) SetVisual(visual gdk.Visual) {
	C.gtk_widget_set_visual(v.native(),
		(*C.GdkVisual)(unsafe.Pointer(gdk_impl.CastToVisual(visual).Native())))
}

// SetAppPaintable is a wrapper around gtk_widget_set_app_paintable().
func (v *widget) SetAppPaintable(paintable bool) {
	C.gtk_widget_set_app_paintable(v.native(), gbool(paintable))
}

// GetAppPaintable is a wrapper around gtk_widget_get_app_paintable().
func (v *widget) GetAppPaintable() bool {
	c := C.gtk_widget_get_app_paintable(v.native())
	return gobool(c)
}

// QueueDraw is a wrapper around gtk_widget_queue_draw().
func (v *widget) QueueDraw() {
	C.gtk_widget_queue_draw(v.native())
}

// GetAllocation is a wrapper around gtk_widget_get_allocation().
func (v *widget) GetAllocation() gtk.Allocation {
	var a allocation
	C.gtk_widget_get_allocation(v.native(), a.native())
	return &a
}

// SetAllocation is a wrapper around gtk_widget_set_allocation().
func (v *widget) SetAllocation(allocation gtk.Allocation) {
	C.gtk_widget_set_allocation(v.native(), castToAllocation(allocation).native())
}

// SizeAllocate is a wrapper around gtk_widget_size_allocate().
func (v *widget) SizeAllocate(allocation gtk.Allocation) {
	C.gtk_widget_size_allocate(v.native(), castToAllocation(allocation).native())
}

// SetStateFlags is a wrapper around gtk_widget_set_state_flags().
func (v *widget) SetStateFlags(stateFlags gtk.StateFlags, clear bool) {
	C.gtk_widget_set_state_flags(v.native(), C.GtkStateFlags(stateFlags), gbool(clear))
}

// GetWindow is a wrapper around gtk_widget_get_window().
func (v *widget) GetWindow() (gdk.Window, error) {
	c := C.gtk_widget_get_window(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	w := &gdk_impl.Window{wrapObject(unsafe.Pointer(c))}
	return w, nil
}

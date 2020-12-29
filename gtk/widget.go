// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"errors"
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
)

/*
 * GtkWidget
 */

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.gtk_size_request_mode_get_type()), marshalSizeRequestMode},

		// Boxed
		{glib.Type(C.gtk_requisition_get_type()), marshalRequisition},
	}
	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkRequisition"] = wrapRequisition
}

// SizeRequestMode is a representation of GTK's GtkSizeRequestMode.
type SizeRequestMode int

const (
	SIZE_REQUEST_HEIGHT_FOR_WIDTH SizeRequestMode = C.GTK_SIZE_REQUEST_HEIGHT_FOR_WIDTH
	SIZE_REQUEST_WIDTH_FOR_HEIGHT SizeRequestMode = C.GTK_SIZE_REQUEST_WIDTH_FOR_HEIGHT
	SIZE_REQUEST_CONSTANT_SIZE    SizeRequestMode = C.GTK_SIZE_REQUEST_CONSTANT_SIZE
)

func marshalSizeRequestMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SizeRequestMode(c), nil
}

// Widget is a representation of GTK's GtkWidget.
type Widget struct {
	glib.InitiallyUnowned
}

// IWidget is an interface type implemented by all structs
// embedding a Widget.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkWidget.
type IWidget interface {
	toWidget() *C.GtkWidget
	ToWidget() *Widget
	Set(string, interface{}) error
}

type IWidgetable interface {
	toWidget() *C.GtkWidget
}

func nullableWidget(v IWidgetable) *C.GtkWidget {
	if v == nil {
		return nil
	}

	return v.toWidget()
}

// native returns a pointer to the underlying GtkWidget.
func (v *Widget) native() *C.GtkWidget {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkWidget(p)
}

func (v *Widget) toWidget() *C.GtkWidget {
	if v == nil {
		return nil
	}
	return v.native()
}

// ToWidget is a helper getter, e.g.: it returns *gtk.Label as a *gtk.Widget.
// In other cases, where you have a gtk.IWidget, use the type assertion.
func (v *Widget) ToWidget() *Widget {
	return v
}

// Cast changes the widget to an object of interface type IWidget.
// This is only useful if you don't already have an object of type IWidget at hand (see example below).
// This func is similar to gtk.Builder.GetObject():
// The returned value needs to be type-asserted, before it can be used.
//
// Example:
//   // you know that the parent is an object of type *gtk.ApplicationWindow,
//   // or you want to check just in case
//   parentWindow, _ := myWindow.GetTransientFor()
//   intermediate, _ := parentWindow.Cast()
//   appWindow, typeAssertSuccessful := intermediate.(*gtk.ApplicationWindow)
func (v *Widget) Cast() (IWidget, error) {
	return castWidget(v.native())
}

func marshalWidget(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj), nil
}

func wrapWidget(obj *glib.Object) *Widget {
	return &Widget{glib.InitiallyUnowned{obj}}
}

// TODO:
// GtkCallback().
// gtk_widget_new().

// Destroy is a wrapper around gtk_widget_destroy().
func (v *Widget) Destroy() {
	C.gtk_widget_destroy(v.native())
}

// HideOnDelete is a wrapper around gtk_widget_hide_on_delete().
// Calling this func adds gtk_widget_hide_on_delete to the widget's "delete-event" signal.
func (v *Widget) HideOnDelete() {
	C._gtk_widget_hide_on_delete(v.native())
}

// TODO:
// gtk_widget_set_direction().
// gtk_widget_get_direction().
// gtk_widget_set_default_direction().
// gtk_widget_get_default_direction().
// gtk_widget_input_shape_combine_region().
// gtk_widget_create_pango_context().
// gtk_widget_create_pango_context().
// gtk_widget_get_pango_context().
// gtk_widget_create_pango_layout().

// QueueDrawArea is a wrapper aroung gtk_widget_queue_draw_area().
func (v *Widget) QueueDrawArea(x, y, w, h int) {
	C.gtk_widget_queue_draw_area(v.native(), C.gint(x), C.gint(y), C.gint(w), C.gint(h))
}

// QueueDrawRegion is a wrapper aroung gtk_widget_queue_draw_region().
func (v *Widget) QueueDrawRegion(region *cairo.Region) {
	C.gtk_widget_queue_draw_region(v.native(), (*C.cairo_region_t)(unsafe.Pointer(region.Native())))
}

// TODO:
// gtk_widget_set_redraw_on_allocate().
// gtk_widget_mnemonic_activate().
// gtk_widget_class_install_style_property().
// gtk_widget_class_install_style_property_parser().
// gtk_widget_class_find_style_property().
// gtk_widget_class_list_style_properties().
// gtk_widget_send_focus_change().
// gtk_widget_style_get().
// gtk_widget_style_get_property().
// gtk_widget_style_get_valist().
// gtk_widget_class_set_accessible_type().
// gtk_widget_get_accessible().
// gtk_widget_child_focus().
// gtk_widget_child_notify().
// gtk_widget_get_child_visible().
// gtk_widget_get_settings().
// gtk_widget_get_clipboard().

// GetDisplay is a wrapper around gtk_widget_get_display().
func (v *Widget) GetDisplay() (*gdk.Display, error) {
	c := C.gtk_widget_get_display(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	s := &gdk.Display{glib.Take(unsafe.Pointer(c))}
	return s, nil
}

// GetScreen is a wrapper around gtk_widget_get_screen().
func (v *Widget) GetScreen() (*gdk.Screen, error) {
	c := C.gtk_widget_get_screen(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	s := &gdk.Screen{glib.Take(unsafe.Pointer(c))}
	return s, nil
}

// TODO:
// gtk_widget_has_screen().
// gtk_widget_get_size_request().
// gtk_widget_set_child_visible().
// gtk_widget_list_mnemonic_labels().
// gtk_widget_add_mnemonic_label().
// gtk_widget_remove_mnemonic_label().
// gtk_widget_error_bell().
// gtk_widget_keynav_failed().
// gtk_widget_get_tooltip_window().
// gtk_widget_get_has_tooltip().
// gtk_widget_set_has_tooltip().
// gtk_widget_trigger_tooltip_query().
// gtk_cairo_should_draw_window().
// gtk_cairo_transform_to_window().

// DragDestSet is a wrapper around gtk_drag_dest_set().
func (v *Widget) DragDestSet(flags DestDefaults, targets []TargetEntry, actions gdk.DragAction) {
	C.gtk_drag_dest_set(v.native(), C.GtkDestDefaults(flags), (*C.GtkTargetEntry)(&targets[0]), C.gint(len(targets)), C.GdkDragAction(actions))
}

// DragSourceSet is a wrapper around gtk_drag_source_set().
func (v *Widget) DragSourceSet(startButtonMask gdk.ModifierType, targets []TargetEntry, actions gdk.DragAction) {
	C.gtk_drag_source_set(v.native(), C.GdkModifierType(startButtonMask), (*C.GtkTargetEntry)(&targets[0]), C.gint(len(targets)), C.GdkDragAction(actions))
}

// ResetStyle is a wrapper around gtk_widget_reset_style().
func (v *Widget) ResetStyle() {
	C.gtk_widget_reset_style(v.native())
}

// InDestruction is a wrapper around gtk_widget_in_destruction().
func (v *Widget) InDestruction() bool {
	return gobool(C.gtk_widget_in_destruction(v.native()))
}

// TODO(jrick) this may require some rethinking
/*
// Destroyed is a wrapper around gtk_widget_destroyed().
func (v *Widget) Destroyed(widgetPointer **Widget) {
}
*/

// Unparent is a wrapper around gtk_widget_unparent().
func (v *Widget) Unparent() {
	C.gtk_widget_unparent(v.native())
}

// Show is a wrapper around gtk_widget_show().
func (v *Widget) Show() {
	C.gtk_widget_show(v.native())
}

// Hide is a wrapper around gtk_widget_hide().
func (v *Widget) Hide() {
	C.gtk_widget_hide(v.native())
}

// GetCanFocus is a wrapper around gtk_widget_get_can_focus().
func (v *Widget) GetCanFocus() bool {
	c := C.gtk_widget_get_can_focus(v.native())
	return gobool(c)
}

// SetCanFocus is a wrapper around gtk_widget_set_can_focus().
func (v *Widget) SetCanFocus(canFocus bool) {
	C.gtk_widget_set_can_focus(v.native(), gbool(canFocus))
}

// GetCanDefault is a wrapper around gtk_widget_get_can_default().
func (v *Widget) GetCanDefault() bool {
	c := C.gtk_widget_get_can_default(v.native())
	return gobool(c)
}

// SetCanDefault is a wrapper around gtk_widget_set_can_default().
func (v *Widget) SetCanDefault(canDefault bool) {
	C.gtk_widget_set_can_default(v.native(), gbool(canDefault))
}

// SetMapped is a wrapper around gtk_widget_set_mapped().
func (v *Widget) SetMapped(mapped bool) {
	C.gtk_widget_set_mapped(v.native(), gbool(mapped))
}

// GetMapped is a wrapper around gtk_widget_get_mapped().
func (v *Widget) GetMapped() bool {
	c := C.gtk_widget_get_mapped(v.native())
	return gobool(c)
}

// TODO:
// gtk_widget_device_is_shadowed().
// gtk_widget_get_modifier_mask().

// InsertActionGroup is a wrapper around gtk_widget_insert_action_group().
func (v *Widget) InsertActionGroup(name string, group glib.IActionGroup) {
	C.gtk_widget_insert_action_group(v.native(), (*C.gchar)(C.CString(name)), C.toGActionGroup(unsafe.Pointer(group.Native())))
}

// TODO:
// gtk_widget_get_path().

// GetPreferredHeight is a wrapper around gtk_widget_get_preferred_height().
func (v *Widget) GetPreferredHeight() (int, int) {
	var minimum, natural C.gint
	C.gtk_widget_get_preferred_height(v.native(), &minimum, &natural)
	return int(minimum), int(natural)
}

// GetPreferredWidth is a wrapper around gtk_widget_get_preferred_width().
func (v *Widget) GetPreferredWidth() (int, int) {
	var minimum, natural C.gint
	C.gtk_widget_get_preferred_width(v.native(), &minimum, &natural)
	return int(minimum), int(natural)
}

// GetPreferredHeightForWidth is a wrapper around gtk_widget_get_preferred_height_for_width().
func (v *Widget) GetPreferredHeightForWidth(width int) (int, int) {

	var minimum, natural C.gint

	C.gtk_widget_get_preferred_height_for_width(
		v.native(),
		C.gint(width),
		&minimum,
		&natural)
	return int(minimum), int(natural)
}

// GetPreferredWidthForHeight is a wrapper around gtk_widget_get_preferred_width_for_height().
func (v *Widget) GetPreferredWidthForHeight(height int) (int, int) {

	var minimum, natural C.gint

	C.gtk_widget_get_preferred_width_for_height(
		v.native(),
		C.gint(height),
		&minimum,
		&natural)
	return int(minimum), int(natural)
}

// GetRequestMode is a wrapper around gtk_widget_get_request_mode().
func (v *Widget) GetRequestMode() SizeRequestMode {
	return SizeRequestMode(C.gtk_widget_get_request_mode(v.native()))
}

// GetPreferredSize is a wrapper around gtk_widget_get_preferred_size().
func (v *Widget) GetPreferredSize() (*Requisition, *Requisition) {

	minimum_size := new(C.GtkRequisition)
	natural_size := new(C.GtkRequisition)

	C.gtk_widget_get_preferred_size(v.native(), minimum_size, natural_size)

	minR, err := requisitionFromNative(minimum_size)
	if err != nil {
		minR = nil
	}
	natR, err := requisitionFromNative(natural_size)
	if err != nil {
		natR = nil
	}

	return minR, natR
}

// TODO:
/*
gint
gtk_distribute_natural_allocation (gint extra_space,
                                   guint n_requested_sizes,
                                   GtkRequestedSize *sizes);
*/

// GetHAlign is a wrapper around gtk_widget_get_halign().
func (v *Widget) GetHAlign() Align {
	c := C.gtk_widget_get_halign(v.native())
	return Align(c)
}

// SetHAlign is a wrapper around gtk_widget_set_halign().
func (v *Widget) SetHAlign(align Align) {
	C.gtk_widget_set_halign(v.native(), C.GtkAlign(align))
}

// GetVAlign is a wrapper around gtk_widget_get_valign().
func (v *Widget) GetVAlign() Align {
	c := C.gtk_widget_get_valign(v.native())
	return Align(c)
}

// SetVAlign is a wrapper around gtk_widget_set_valign().
func (v *Widget) SetVAlign(align Align) {
	C.gtk_widget_set_valign(v.native(), C.GtkAlign(align))
}

// GetMarginTop is a wrapper around gtk_widget_get_margin_top().
func (v *Widget) GetMarginTop() int {
	c := C.gtk_widget_get_margin_top(v.native())
	return int(c)
}

// SetMarginTop is a wrapper around gtk_widget_set_margin_top().
func (v *Widget) SetMarginTop(margin int) {
	C.gtk_widget_set_margin_top(v.native(), C.gint(margin))
}

// GetMarginBottom is a wrapper around gtk_widget_get_margin_bottom().
func (v *Widget) GetMarginBottom() int {
	c := C.gtk_widget_get_margin_bottom(v.native())
	return int(c)
}

// SetMarginBottom is a wrapper around gtk_widget_set_margin_bottom().
func (v *Widget) SetMarginBottom(margin int) {
	C.gtk_widget_set_margin_bottom(v.native(), C.gint(margin))
}

// GetHExpand is a wrapper around gtk_widget_get_hexpand().
func (v *Widget) GetHExpand() bool {
	c := C.gtk_widget_get_hexpand(v.native())
	return gobool(c)
}

// SetHExpand is a wrapper around gtk_widget_set_hexpand().
func (v *Widget) SetHExpand(expand bool) {
	C.gtk_widget_set_hexpand(v.native(), gbool(expand))
}

// TODO:
// gtk_widget_get_hexpand_set().
// gtk_widget_set_hexpand_set().

// GetVExpand is a wrapper around gtk_widget_get_vexpand().
func (v *Widget) GetVExpand() bool {
	c := C.gtk_widget_get_vexpand(v.native())
	return gobool(c)
}

// SetVExpand is a wrapper around gtk_widget_set_vexpand().
func (v *Widget) SetVExpand(expand bool) {
	C.gtk_widget_set_vexpand(v.native(), gbool(expand))
}

// TODO:
// gtk_widget_get_vexpand_set().
// gtk_widget_set_vexpand_set().
// gtk_widget_queue_compute_expand().
// gtk_widget_compute_expand().

// GetRealized is a wrapper around gtk_widget_get_realized().
func (v *Widget) GetRealized() bool {
	c := C.gtk_widget_get_realized(v.native())
	return gobool(c)
}

// SetRealized is a wrapper around gtk_widget_set_realized().
func (v *Widget) SetRealized(realized bool) {
	C.gtk_widget_set_realized(v.native(), gbool(realized))
}

// GetHasWindow is a wrapper around gtk_widget_get_has_window().
func (v *Widget) GetHasWindow() bool {
	c := C.gtk_widget_get_has_window(v.native())
	return gobool(c)
}

// SetHasWindow is a wrapper around gtk_widget_set_has_window().
func (v *Widget) SetHasWindow(hasWindow bool) {
	C.gtk_widget_set_has_window(v.native(), gbool(hasWindow))
}

// ShowNow is a wrapper around gtk_widget_show_now().
func (v *Widget) ShowNow() {
	C.gtk_widget_show_now(v.native())
}

// ShowAll is a wrapper around gtk_widget_show_all().
func (v *Widget) ShowAll() {
	C.gtk_widget_show_all(v.native())
}

// SetNoShowAll is a wrapper around gtk_widget_set_no_show_all().
func (v *Widget) SetNoShowAll(noShowAll bool) {
	C.gtk_widget_set_no_show_all(v.native(), gbool(noShowAll))
}

// GetNoShowAll is a wrapper around gtk_widget_get_no_show_all().
func (v *Widget) GetNoShowAll() bool {
	c := C.gtk_widget_get_no_show_all(v.native())
	return gobool(c)
}

// Map is a wrapper around gtk_widget_map().
func (v *Widget) Map() {
	C.gtk_widget_map(v.native())
}

// Unmap is a wrapper around gtk_widget_unmap().
func (v *Widget) Unmap() {
	C.gtk_widget_unmap(v.native())
}

// TODO:
//void gtk_widget_realize(GtkWidget *widget);
//void gtk_widget_unrealize(GtkWidget *widget);
//void gtk_widget_draw(GtkWidget *widget, cairo_t *cr);
//void gtk_widget_queue_resize(GtkWidget *widget);
//void gtk_widget_queue_resize_no_redraw(GtkWidget *widget);
// gtk_widget_queue_allocate().

// Event() is a wrapper around gtk_widget_event().
func (v *Widget) Event(event *gdk.Event) bool {
	c := C.gtk_widget_event(v.native(),
		(*C.GdkEvent)(unsafe.Pointer(event.Native())))
	return gobool(c)
}

// Activate() is a wrapper around gtk_widget_activate().
func (v *Widget) Activate() bool {
	return gobool(C.gtk_widget_activate(v.native()))
}

// Intersect is a wrapper around gtk_widget_intersect().
func (v *Widget) Intersect(area gdk.Rectangle) (*gdk.Rectangle, bool) {
	var cRect *C.GdkRectangle
	hadIntersection := C.gtk_widget_intersect(v.native(), nativeGdkRectangle(area), cRect)
	intersection := gdk.WrapRectangle(uintptr(unsafe.Pointer(cRect)))
	return intersection, gobool(hadIntersection)
}

// IsFocus() is a wrapper around gtk_widget_is_focus().
func (v *Widget) IsFocus() bool {
	return gobool(C.gtk_widget_is_focus(v.native()))
}

// GrabFocus() is a wrapper around gtk_widget_grab_focus().
func (v *Widget) GrabFocus() {
	C.gtk_widget_grab_focus(v.native())
}

// GrabDefault() is a wrapper around gtk_widget_grab_default().
func (v *Widget) GrabDefault() {
	C.gtk_widget_grab_default(v.native())
}

// SetName() is a wrapper around gtk_widget_set_name().
func (v *Widget) SetName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_widget_set_name(v.native(), (*C.gchar)(cstr))
}

// GetName() is a wrapper around gtk_widget_get_name().  A non-nil
// error is returned in the case that gtk_widget_get_name returns NULL to
// differentiate between NULL and an empty string.
func (v *Widget) GetName() (string, error) {
	c := C.gtk_widget_get_name(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// GetSensitive is a wrapper around gtk_widget_get_sensitive().
func (v *Widget) GetSensitive() bool {
	c := C.gtk_widget_get_sensitive(v.native())
	return gobool(c)
}

// IsSensitive is a wrapper around gtk_widget_is_sensitive().
func (v *Widget) IsSensitive() bool {
	c := C.gtk_widget_is_sensitive(v.native())
	return gobool(c)
}

// SetSensitive is a wrapper around gtk_widget_set_sensitive().
func (v *Widget) SetSensitive(sensitive bool) {
	C.gtk_widget_set_sensitive(v.native(), gbool(sensitive))
}

// GetVisible is a wrapper around gtk_widget_get_visible().
func (v *Widget) GetVisible() bool {
	c := C.gtk_widget_get_visible(v.native())
	return gobool(c)
}

// SetVisible is a wrapper around gtk_widget_set_visible().
func (v *Widget) SetVisible(visible bool) {
	C.gtk_widget_set_visible(v.native(), gbool(visible))
}

// SetParent is a wrapper around gtk_widget_set_parent().
func (v *Widget) SetParent(parent IWidget) {
	C.gtk_widget_set_parent(v.native(), parent.toWidget())
}

// GetParent is a wrapper around gtk_widget_get_parent().
func (v *Widget) GetParent() (IWidget, error) {
	c := C.gtk_widget_get_parent(v.native())
	if c == nil {
		return nil, nil
	}
	return castWidget(c)
}

// SetSizeRequest is a wrapper around gtk_widget_set_size_request().
func (v *Widget) SetSizeRequest(width, height int) {
	C.gtk_widget_set_size_request(v.native(), C.gint(width), C.gint(height))
}

// GetSizeRequest is a wrapper around gtk_widget_get_size_request().
func (v *Widget) GetSizeRequest() (width, height int) {
	var w, h C.gint
	C.gtk_widget_get_size_request(v.native(), &w, &h)
	return int(w), int(h)
}

// SetParentWindow is a wrapper around gtk_widget_set_parent_window().
func (v *Widget) SetParentWindow(parentWindow *gdk.Window) {
	C.gtk_widget_set_parent_window(v.native(),
		(*C.GdkWindow)(unsafe.Pointer(parentWindow.Native())))
}

// GetParentWindow is a wrapper around gtk_widget_get_parent_window().
func (v *Widget) GetParentWindow() (*gdk.Window, error) {
	c := C.gtk_widget_get_parent_window(v.native())
	if v == nil {
		return nil, nilPtrErr
	}

	w := &gdk.Window{glib.Take(unsafe.Pointer(c))}
	return w, nil
}

// SetEvents is a wrapper around gtk_widget_set_events().
func (v *Widget) SetEvents(events int) {
	C.gtk_widget_set_events(v.native(), C.gint(events))
}

// GetEvents is a wrapper around gtk_widget_get_events().
func (v *Widget) GetEvents() int {
	return int(C.gtk_widget_get_events(v.native()))
}

// AddEvents is a wrapper around gtk_widget_add_events().
func (v *Widget) AddEvents(events int) {
	C.gtk_widget_add_events(v.native(), C.gint(events))
}

// TODO:
/*
// gtk_widget_set_device_events().
func (v *Widget) SetDeviceEvents() {
}
*/

/*
// gtk_widget_get_device_events().
func (v *Widget) GetDeviceEvents() {
}
*/

/*
// gtk_widget_add_device_events().
func (v *Widget) AddDeviceEvents() {
}
*/

// FreezeChildNotify is a wrapper around gtk_widget_freeze_child_notify().
func (v *Widget) FreezeChildNotify() {
	C.gtk_widget_freeze_child_notify(v.native())
}

// ThawChildNotify is a wrapper around gtk_widget_thaw_child_notify().
func (v *Widget) ThawChildNotify() {
	C.gtk_widget_thaw_child_notify(v.native())
}

// HasDefault is a wrapper around gtk_widget_has_default().
func (v *Widget) HasDefault() bool {
	c := C.gtk_widget_has_default(v.native())
	return gobool(c)
}

// HasFocus is a wrapper around gtk_widget_has_focus().
func (v *Widget) HasFocus() bool {
	c := C.gtk_widget_has_focus(v.native())
	return gobool(c)
}

// HasVisibleFocus is a wrapper around gtk_widget_has_visible_focus().
func (v *Widget) HasVisibleFocus() bool {
	c := C.gtk_widget_has_visible_focus(v.native())
	return gobool(c)
}

// HasGrab is a wrapper around gtk_widget_has_grab().
func (v *Widget) HasGrab() bool {
	c := C.gtk_widget_has_grab(v.native())
	return gobool(c)
}

// IsDrawable is a wrapper around gtk_widget_is_drawable().
func (v *Widget) IsDrawable() bool {
	c := C.gtk_widget_is_drawable(v.native())
	return gobool(c)
}

// IsToplevel is a wrapper around gtk_widget_is_toplevel().
func (v *Widget) IsToplevel() bool {
	c := C.gtk_widget_is_toplevel(v.native())
	return gobool(c)
}

// FIXME:
// gtk_widget_set_window().
// gtk_widget_set_receives_default().
// gtk_widget_get_receives_default().
// gtk_widget_set_support_multidevice().
// gtk_widget_get_support_multidevice().

// SetDeviceEnabled is a wrapper around gtk_widget_set_device_enabled().
func (v *Widget) SetDeviceEnabled(device *gdk.Device, enabled bool) {
	C.gtk_widget_set_device_enabled(v.native(),
		(*C.GdkDevice)(unsafe.Pointer(device.Native())), gbool(enabled))
}

// GetDeviceEnabled is a wrapper around gtk_widget_get_device_enabled().
func (v *Widget) GetDeviceEnabled(device *gdk.Device) bool {
	c := C.gtk_widget_get_device_enabled(v.native(),
		(*C.GdkDevice)(unsafe.Pointer(device.Native())))
	return gobool(c)
}

// GetToplevel is a wrapper around gtk_widget_get_toplevel().
func (v *Widget) GetToplevel() (IWidget, error) {
	c := C.gtk_widget_get_toplevel(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return castWidget(c)
}

// TODO:
// gtk_widget_get_ancestor().
// gtk_widget_get_visual().
// gtk_widget_is_ancestor().

// GetTooltipMarkup is a wrapper around gtk_widget_get_tooltip_markup().
// A non-nil error is returned in the case that gtk_widget_get_tooltip_markup
// returns NULL to differentiate between NULL and an empty string.
func (v *Widget) GetTooltipMarkup() (string, error) {
	c := C.gtk_widget_get_tooltip_markup(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// SetTooltipMarkup is a wrapper around gtk_widget_set_tooltip_markup().
func (v *Widget) SetTooltipMarkup(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_widget_set_tooltip_markup(v.native(), (*C.gchar)(cstr))
}

// GetTooltipText is a wrapper around gtk_widget_get_tooltip_text().
// A non-nil error is returned in the case that
// gtk_widget_get_tooltip_text returns NULL to differentiate between NULL
// and an empty string.
func (v *Widget) GetTooltipText() (string, error) {
	c := C.gtk_widget_get_tooltip_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// SetTooltipText is a wrapper around gtk_widget_set_tooltip_text().
func (v *Widget) SetTooltipText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_widget_set_tooltip_text(v.native(), (*C.gchar)(cstr))
}

// TranslateCoordinates is a wrapper around gtk_widget_translate_coordinates().
func (v *Widget) TranslateCoordinates(dest IWidget, srcX, srcY int) (destX, destY int, e error) {
	cdest := nullableWidget(dest)

	var cdestX, cdestY C.gint
	c := C.gtk_widget_translate_coordinates(v.native(), cdest, C.gint(srcX), C.gint(srcY), &cdestX, &cdestY)
	if !gobool(c) {
		return 0, 0, errors.New("translate coordinates failed")
	}
	return int(cdestX), int(cdestY), nil
}

// SetVisual is a wrapper around gtk_widget_set_visual().
func (v *Widget) SetVisual(visual *gdk.Visual) {
	C.gtk_widget_set_visual(v.native(),
		(*C.GdkVisual)(unsafe.Pointer(visual.Native())))
}

// SetAppPaintable is a wrapper around gtk_widget_set_app_paintable().
func (v *Widget) SetAppPaintable(paintable bool) {
	C.gtk_widget_set_app_paintable(v.native(), gbool(paintable))
}

// GetAppPaintable is a wrapper around gtk_widget_get_app_paintable().
func (v *Widget) GetAppPaintable() bool {
	c := C.gtk_widget_get_app_paintable(v.native())
	return gobool(c)
}

// QueueDraw is a wrapper around gtk_widget_queue_draw().
func (v *Widget) QueueDraw() {
	C.gtk_widget_queue_draw(v.native())
}

// GetAllocation is a wrapper around gtk_widget_get_allocation().
func (v *Widget) GetAllocation() *Allocation {
	var a Allocation
	C.gtk_widget_get_allocation(v.native(), a.native())
	return &a
}

// SetAllocation is a wrapper around gtk_widget_set_allocation().
func (v *Widget) SetAllocation(allocation *Allocation) {
	C.gtk_widget_set_allocation(v.native(), allocation.native())
}

// SizeAllocate is a wrapper around gtk_widget_size_allocate().
func (v *Widget) SizeAllocate(allocation *Allocation) {
	C.gtk_widget_size_allocate(v.native(), allocation.native())
}

// TODO:
// gtk_widget_size_allocate_with_baseline().

// SetStateFlags is a wrapper around gtk_widget_set_state_flags().
func (v *Widget) SetStateFlags(stateFlags StateFlags, clear bool) {
	C.gtk_widget_set_state_flags(v.native(), C.GtkStateFlags(stateFlags), gbool(clear))
}

// TODO:
// gtk_widget_unset_state_flags().
// gtk_widget_get_state_flags().

// GetWindow is a wrapper around gtk_widget_get_window().
func (v *Widget) GetWindow() (*gdk.Window, error) {
	c := C.gtk_widget_get_window(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	w := &gdk.Window{glib.Take(unsafe.Pointer(c))}
	return w, nil
}

/*
 * GtkRequisition
 */

// Requisition is a representation of GTK's GtkRequisition
type Requisition struct {
	requisition *C.GtkRequisition
	Width,
	Height int
}

func (v *Requisition) native() *C.GtkRequisition {
	if v == nil {
		return nil
	}
	v.requisition.width = C.int(v.Width)
	v.requisition.height = C.int(v.Height)
	return v.requisition
}

// Native returns a pointer to the underlying GtkRequisition.
func (v *Requisition) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalRequisition(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	requisition := (*C.GtkRequisition)(unsafe.Pointer(c))
	return wrapRequisition(requisition), nil
}

func wrapRequisition(requisition *C.GtkRequisition) *Requisition {
	if requisition == nil {
		return nil
	}
	return &Requisition{requisition, int(requisition.width), int(requisition.height)}
}

// requisitionFromNative that handle finalizer.
func requisitionFromNative(requisitionNative *C.GtkRequisition) (*Requisition, error) {
	requisition := wrapRequisition(requisitionNative)
	if requisition == nil {
		return nil, nilPtrErr
	}
	runtime.SetFinalizer(requisition, (*Requisition).free)
	return requisition, nil
}

// RequisitionNew is a wrapper around gtk_requisition_new().
func RequisitionNew() (*Requisition, error) {
	c := C.gtk_requisition_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return requisitionFromNative(c)
}

// free is a wrapper around gtk_requisition_free().
func (v *Requisition) free() {
	C.gtk_requisition_free(v.native())
}

// Copy is a wrapper around gtk_requisition_copy().
func (v *Requisition) Copy() (*Requisition, error) {
	c := C.gtk_requisition_copy(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return requisitionFromNative(c)
}

/*
 * GtkAllocation
 */

// Allocation is a representation of GTK's GtkAllocation type.
type Allocation struct {
	gdk.Rectangle
}

// Native returns a pointer to the underlying GtkAllocation.
func (v *Allocation) native() *C.GtkAllocation {
	return (*C.GtkAllocation)(unsafe.Pointer(&v.GdkRectangle))
}

// GetAllocatedWidth() is a wrapper around gtk_widget_get_allocated_width().
func (v *Widget) GetAllocatedWidth() int {
	return int(C.gtk_widget_get_allocated_width(v.native()))
}

// GetAllocatedHeight() is a wrapper around gtk_widget_get_allocated_height().
func (v *Widget) GetAllocatedHeight() int {
	return int(C.gtk_widget_get_allocated_height(v.native()))
}

// TODO:
// gtk_widget_get_allocated_baseline().
// gtk_widget_get_allocated_size().

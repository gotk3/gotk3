// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16

// See: https://developer.gnome.org/gtk3/3.18/api-index-3-18.html

// For gtk_overlay_reorder_overlay():
// See: https://git.gnome.org/browse/gtk+/tree/gtk/gtkoverlay.h?h=gtk-3-18

package gtk

// #include <gtk/gtk.h>
import "C"

/*
 * GtkStack
 */

// TODO:
// gtk_stack_get_interpolate_size().
// gtk_stack_set_interpolate_size().

/*
 * GtkRadioMenuItem
 */

// JoinGroup is a wrapper around gtk_radio_menu_item_join_group().
func (v *RadioMenuItem) JoinGroup(group_source *RadioMenuItem) {
	C.gtk_radio_menu_item_join_group(v.native(), group_source.native())
}

/*
 * GtkOverlay
 */

// ReorderOverlay() is a wrapper around gtk_overlay_reorder_overlay().
func (v *Overlay) ReorderOverlay(child IWidget, position int) {
	C.gtk_overlay_reorder_overlay(v.native(), child.toWidget(), C.int(position))
}

// GetOverlayPassThrough() is a wrapper around gtk_overlay_get_overlay_pass_through().
func (v *Overlay) GetOverlayPassThrough(widget IWidget) bool {
	c := C.gtk_overlay_get_overlay_pass_through(v.native(), widget.toWidget())
	return gobool(c)
}

// SetOverlayPassThrough() is a wrapper around gtk_overlay_set_overlay_pass_through().
func (v *Overlay) SetOverlayPassThrough(widget IWidget, passThrough bool) {
	C.gtk_overlay_set_overlay_pass_through(v.native(), widget.toWidget(), gbool(passThrough))
}

/*
 * GtkPlacesSidebar
 */

// TODO:
// gtk_places_sidebar_set_show_recent().
// gtk_places_sidebar_get_show_recent().
// gtk_places_sidebar_get_show_trash().
// gtk_places_sidebar_set_show_trash().
// gtk_places_sidebar_get_show_other_locations().
// gtk_places_sidebar_set_show_other_locations().
// gtk_places_sidebar_set_drop_targets_visible().

/*
 * GtkPopover
 */

// SetDefaultWidget is a wrapper around gtk_popover_set_default_widget().
func (p *Popover) SetDefaultWidget(widget IWidget) {
	C.gtk_popover_set_default_widget(p.native(), widget.toWidget())
}

// GetDefaultWidget is a wrapper around gtk_popover_get_default_widget().
func (p *Popover) GetDefaultWidget() (IWidget, error) {
	w := C.gtk_popover_get_default_widget(p.native())
	if w == nil {
		return nil, nil
	}
	return castWidget(w)
}

/*
 * GtkTextView
 */

// SetTopMargin is a wrapper around gtk_text_view_set_top_margin().
func (v *TextView) SetTopMargin(topMargin int) {
	C.gtk_text_view_set_top_margin(v.native(), C.gint(topMargin))
}

// GetTopMargin is a wrapper around gtk_text_view_get_top_margin().
func (v *TextView) GetTopMargin() int {
	return int(C.gtk_text_view_get_top_margin(v.native()))
}

// SetBottomMargin is a wrapper around gtk_text_view_set_bottom_margin().
func (v *TextView) SetBottomMargin(bottomMargin int) {
	C.gtk_text_view_set_bottom_margin(v.native(), C.gint(bottomMargin))
}

// GetBottomMargin is a wrapper around gtk_text_view_get_bottom_margin().
func (v *TextView) GetBottomMargin() int {
	return int(C.gtk_text_view_get_bottom_margin(v.native()))
}

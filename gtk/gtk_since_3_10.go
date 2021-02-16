// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

// +build !gtk_3_6,!gtk_3_8
// not use this: go build -tags gtk_3_8'. Otherwise, if no build tags are used, GTK 3.10

package gtk

// #include <stdlib.h>
// #include <gtk/gtk.h>
// #include "gtk_since_3_10.go.h"
import "C"
import (
	"errors"
	"sync"
	"unsafe"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/pango"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.gtk_revealer_transition_type_get_type()), marshalRevealerTransitionType},
		{glib.Type(C.gtk_stack_transition_type_get_type()), marshalStackTransitionType},

		// Objects/Interfaces
		{glib.Type(C.gtk_header_bar_get_type()), marshalHeaderBar},
		{glib.Type(C.gtk_list_box_get_type()), marshalListBox},
		{glib.Type(C.gtk_list_box_row_get_type()), marshalListBoxRow},
		{glib.Type(C.gtk_revealer_get_type()), marshalRevealer},
		{glib.Type(C.gtk_search_bar_get_type()), marshalSearchBar},
		{glib.Type(C.gtk_stack_get_type()), marshalStack},
		{glib.Type(C.gtk_stack_switcher_get_type()), marshalStackSwitcher},
	}
	glib.RegisterGValueMarshalers(tm)

	//Contribute to casting
	for k, v := range map[string]WrapFn{
		"GtkHeaderBar":  wrapHeaderBar,
		"GtkListBox":    wrapListBox,
		"GtkListBoxRow": wrapListBoxRow,
		"GtkRevealer":   wrapRevealer,
		"GtkSearchBar":  wrapSearchBar,
		"GtkStack":      wrapStack,
	} {
		WrapMap[k] = v
	}
}

/*
 * Constants
 */

const (
	ALIGN_BASELINE Align = C.GTK_ALIGN_BASELINE
)

// ImageType
const (
	IMAGE_SURFACE ImageType = C.GTK_IMAGE_SURFACE
)

// RevealerTransitionType is a representation of GTK's GtkRevealerTransitionType.
type RevealerTransitionType int

const (
	REVEALER_TRANSITION_TYPE_NONE        RevealerTransitionType = C.GTK_REVEALER_TRANSITION_TYPE_NONE
	REVEALER_TRANSITION_TYPE_CROSSFADE   RevealerTransitionType = C.GTK_REVEALER_TRANSITION_TYPE_CROSSFADE
	REVEALER_TRANSITION_TYPE_SLIDE_RIGHT RevealerTransitionType = C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_RIGHT
	REVEALER_TRANSITION_TYPE_SLIDE_LEFT  RevealerTransitionType = C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_LEFT
	REVEALER_TRANSITION_TYPE_SLIDE_UP    RevealerTransitionType = C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_UP
	REVEALER_TRANSITION_TYPE_SLIDE_DOWN  RevealerTransitionType = C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_DOWN
)

func marshalRevealerTransitionType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return RevealerTransitionType(c), nil
}

// StackTransitionType is a representation of GTK's GtkStackTransitionType.
type StackTransitionType int

const (
	STACK_TRANSITION_TYPE_NONE             StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_NONE
	STACK_TRANSITION_TYPE_CROSSFADE        StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_CROSSFADE
	STACK_TRANSITION_TYPE_SLIDE_RIGHT      StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_SLIDE_RIGHT
	STACK_TRANSITION_TYPE_SLIDE_LEFT       StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT
	STACK_TRANSITION_TYPE_SLIDE_UP         StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_SLIDE_UP
	STACK_TRANSITION_TYPE_SLIDE_DOWN       StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_SLIDE_DOWN
	STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT
	STACK_TRANSITION_TYPE_SLIDE_UP_DOWN    StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_SLIDE_UP_DOWN
)

func marshalStackTransitionType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return StackTransitionType(c), nil
}

/*
 * GtkWidget
 */

// TODO:
// gtk_widget_get_composite_name().

/*
 * GtkImage
 */

// ImageNewFromSurface is a wrapper around gtk_image_new_from_surface().
func ImageNewFromSurface(surface *cairo.Surface) (*Image, error) {
	c := C.gtk_image_new_from_surface((*C.cairo_surface_t)(surface.GetCSurface()))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapImage(obj), nil
}

// SetFromSurface is a wrapper around gtk_image_set_from_surface().
func (v *Image) SetFromSurface(surface *cairo.Surface) {
	csurface := (*C.cairo_surface_t)(surface.GetCSurface())
	C.gtk_image_set_from_surface(v.native(), csurface)
}

/*
 * GtkIconTheme
 */

// HasIcon is a wrapper around gtk_icon_theme_load_icon_for_scale().
func (v *IconTheme) LoadIconForScale(iconName string, size, scale int, flags IconLookupFlags) (*gdk.Pixbuf, error) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError = nil
	c := C.gtk_icon_theme_load_icon_for_scale(v.Theme, (*C.gchar)(cstr), C.gint(size), C.gint(scale), C.GtkIconLookupFlags(flags), &err)
	if c == nil {
		defer C.g_error_free(err)
		return nil, errors.New(goString(err.message))
	}
	return &gdk.Pixbuf{glib.Take(unsafe.Pointer(c))}, nil
}

/*
 * GtkEntry
 */

// GetTabs is a wrapper around gtk_entry_get_tabs().
func (v *Entry) GetTabs() (*pango.TabArray, error) {
	c := C.gtk_entry_get_tabs(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return pango.WrapTabArray(uintptr(unsafe.Pointer(c))), nil
}

// SetTabs is a wrapper around gtk_entry_set_tabs().
func (v *Entry) SetTabs(tabs *pango.TabArray) {
	C.gtk_entry_set_tabs(v.native(), (*C.PangoTabArray)(unsafe.Pointer(tabs.Native())))
}

/*
 * GtkButton
 */

// ButtonNewFromIconName is a wrapper around gtk_button_new_from_icon_name().
func ButtonNewFromIconName(iconName string, size IconSize) (*Button, error) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_button_new_from_icon_name((*C.gchar)(cstr),
		C.GtkIconSize(size))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapButton(glib.Take(unsafe.Pointer(c))), nil
}

/*
 * GtkBox
 */

// TODO:
// gtk_box_get_baseline_position().
// gtk_box_set_baseline_position().

/*
 * GtkGrid
 */

// RemoveRow is a wrapper around gtk_grid_remove_row().
func (v *Grid) RemoveRow(position int) {
	C.gtk_grid_remove_row(v.native(), C.gint(position))
}

// RemoveColumn is a wrapper around gtk_grid_remove_column().
func (v *Grid) RemoveColumn(position int) {
	C.gtk_grid_remove_column(v.native(), C.gint(position))
}

// TODO:
// gtk_grid_get_baseline_row().
// gtk_grid_set_baseline_row().
// gtk_grid_get_row_baseline_position().
// gtk_grid_set_row_baseline_position().

/*
 * GtkHeaderBar
 */

// HeaderBar is a representation of GtkHeaderBar
type HeaderBar struct {
	Container
}

// native returns a pointer to the underlying GtkHeaderBar.
func (v *HeaderBar) native() *C.GtkHeaderBar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkHeaderBar(p)
}

func marshalHeaderBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapHeaderBar(obj), nil
}

func wrapHeaderBar(obj *glib.Object) *HeaderBar {
	if obj == nil {
		return nil
	}

	return &HeaderBar{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// HeaderBarNew is a wrapper around gtk_header_bar_new().
func HeaderBarNew() (*HeaderBar, error) {
	c := C.gtk_header_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapHeaderBar(glib.Take(unsafe.Pointer(c))), nil
}

// SetTitle is a wrapper around gtk_header_bar_set_title().
func (v *HeaderBar) SetTitle(title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_header_bar_set_title(v.native(), (*C.gchar)(cstr))
}

// GetTitle is a wrapper around gtk_header_bar_get_title().
func (v *HeaderBar) GetTitle() string {
	cstr := C.gtk_header_bar_get_title(v.native())
	return C.GoString((*C.char)(cstr))
}

// SetSubtitle is a wrapper around gtk_header_bar_set_subtitle().
func (v *HeaderBar) SetSubtitle(subtitle string) {
	cstr := C.CString(subtitle)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_header_bar_set_subtitle(v.native(), (*C.gchar)(cstr))
}

// GetSubtitle is a wrapper around gtk_header_bar_get_subtitle().
func (v *HeaderBar) GetSubtitle() string {
	cstr := C.gtk_header_bar_get_subtitle(v.native())
	return C.GoString((*C.char)(cstr))
}

// SetCustomTitle is a wrapper around gtk_header_bar_set_custom_title().
func (v *HeaderBar) SetCustomTitle(titleWidget IWidget) {
	C.gtk_header_bar_set_custom_title(v.native(), titleWidget.toWidget())
}

// GetCustomTitle is a wrapper around gtk_header_bar_get_custom_title().
func (v *HeaderBar) GetCustomTitle() (IWidget, error) {
	c := C.gtk_header_bar_get_custom_title(v.native())
	if c == nil {
		return nil, nil
	}
	return castWidget(c)
}

// PackStart is a wrapper around gtk_header_bar_pack_start().
func (v *HeaderBar) PackStart(child IWidget) {
	C.gtk_header_bar_pack_start(v.native(), child.toWidget())
}

// PackEnd is a wrapper around gtk_header_bar_pack_end().
func (v *HeaderBar) PackEnd(child IWidget) {
	C.gtk_header_bar_pack_end(v.native(), child.toWidget())
}

// SetShowCloseButton is a wrapper around gtk_header_bar_set_show_close_button().
func (v *HeaderBar) SetShowCloseButton(setting bool) {
	C.gtk_header_bar_set_show_close_button(v.native(), gbool(setting))
}

// GetShowCloseButton is a wrapper around gtk_header_bar_get_show_close_button().
func (v *HeaderBar) GetShowCloseButton() bool {
	c := C.gtk_header_bar_get_show_close_button(v.native())
	return gobool(c)
}

/*
 * GtkLabel
 */

// GetLines is a wrapper around gtk_label_get_lines().
func (v *Label) GetLines() int {
	c := C.gtk_label_get_lines(v.native())
	return int(c)
}

// SetLines is a wrapper around gtk_label_set_lines().
func (v *Label) SetLines(lines int) {
	C.gtk_label_set_lines(v.native(), C.gint(lines))
}

/*
 * GtkListBox
 */

// ListBox is a representation of GTK's GtkListBox.
type ListBox struct {
	Container
}

// native returns a pointer to the underlying GtkListBox.
func (v *ListBox) native() *C.GtkListBox {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkListBox(p)
}

func marshalListBox(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapListBox(obj), nil
}

func wrapListBox(obj *glib.Object) *ListBox {
	if obj == nil {
		return nil
	}

	return &ListBox{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// ListBoxNew is a wrapper around gtk_list_box_new().
func ListBoxNew() (*ListBox, error) {
	c := C.gtk_list_box_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapListBox(glib.Take(unsafe.Pointer(c))), nil
}

// Prepend is a wrapper around gtk_list_box_prepend().
func (v *ListBox) Prepend(child IWidget) {
	C.gtk_list_box_prepend(v.native(), child.toWidget())
}

// Insert is a wrapper around gtk_list_box_insert().
func (v *ListBox) Insert(child IWidget, position int) {
	C.gtk_list_box_insert(v.native(), child.toWidget(), C.gint(position))
}

// SelectRow is a wrapper around gtk_list_box_select_row().
func (v *ListBox) SelectRow(row *ListBoxRow) {
	C.gtk_list_box_select_row(v.native(), row.native())
}

// GetSelectedRow is a wrapper around gtk_list_box_get_selected_row().
func (v *ListBox) GetSelectedRow() *ListBoxRow {
	c := C.gtk_list_box_get_selected_row(v.native())
	if c == nil {
		return nil
	}
	return wrapListBoxRow(glib.Take(unsafe.Pointer(c)))
}

// SetSelectionMode is a wrapper around gtk_list_box_set_selection_mode().
func (v *ListBox) SetSelectionMode(mode SelectionMode) {
	C.gtk_list_box_set_selection_mode(v.native(), C.GtkSelectionMode(mode))
}

// GetSelectionMode is a wrapper around gtk_list_box_get_selection_mode()
func (v *ListBox) GetSelectionMode() SelectionMode {
	c := C.gtk_list_box_get_selection_mode(v.native())
	return SelectionMode(c)
}

// SetActivateOnSingleClick is a wrapper around gtk_list_box_set_activate_on_single_click().
func (v *ListBox) SetActivateOnSingleClick(single bool) {
	C.gtk_list_box_set_activate_on_single_click(v.native(), gbool(single))
}

// GetActivateOnSingleClick is a wrapper around gtk_list_box_get_activate_on_single_click().
func (v *ListBox) GetActivateOnSingleClick() bool {
	c := C.gtk_list_box_get_activate_on_single_click(v.native())
	return gobool(c)
}

// GetAdjustment is a wrapper around gtk_list_box_get_adjustment().
func (v *ListBox) GetAdjustment() *Adjustment {
	c := C.gtk_list_box_get_adjustment(v.native())
	obj := glib.Take(unsafe.Pointer(c))
	return &Adjustment{glib.InitiallyUnowned{obj}}
}

// SetAdjustment is a wrapper around gtk_list_box_set_adjustment().
func (v *ListBox) SetAdjustment(adjustment *Adjustment) {
	C.gtk_list_box_set_adjustment(v.native(), adjustment.native())
}

// SetPlaceholder is a wrapper around gtk_list_box_set_placeholder().
func (v *ListBox) SetPlaceholder(placeholder IWidget) {
	C.gtk_list_box_set_placeholder(v.native(), placeholder.toWidget())
}

// GetRowAtIndex is a wrapper around gtk_list_box_get_row_at_index().
func (v *ListBox) GetRowAtIndex(index int) *ListBoxRow {
	c := C.gtk_list_box_get_row_at_index(v.native(), C.gint(index))
	if c == nil {
		return nil
	}
	return wrapListBoxRow(glib.Take(unsafe.Pointer(c)))
}

// GetRowAtY is a wrapper around gtk_list_box_get_row_at_y().
func (v *ListBox) GetRowAtY(y int) *ListBoxRow {
	c := C.gtk_list_box_get_row_at_y(v.native(), C.gint(y))
	if c == nil {
		return nil
	}
	return wrapListBoxRow(glib.Take(unsafe.Pointer(c)))
}

// InvalidateFilter is a wrapper around gtk_list_box_invalidate_filter().
func (v *ListBox) InvalidateFilter() {
	C.gtk_list_box_invalidate_filter(v.native())
}

// InvalidateHeaders is a wrapper around gtk_list_box_invalidate_headers().
func (v *ListBox) InvalidateHeaders() {
	C.gtk_list_box_invalidate_headers(v.native())
}

// InvalidateSort is a wrapper around gtk_list_box_invalidate_sort().
func (v *ListBox) InvalidateSort() {
	C.gtk_list_box_invalidate_sort(v.native())
}

// ListBoxFilterFunc is a representation of GtkListBoxFilterFunc
type ListBoxFilterFunc func(row *ListBoxRow, userData ...interface{}) bool

type listBoxFilterFuncData struct {
	fn       ListBoxFilterFunc
	userData []interface{}
}

var (
	listBoxFilterFuncRegistry = struct {
		sync.RWMutex
		next int
		m    map[int]listBoxFilterFuncData
	}{
		next: 1,
		m:    make(map[int]listBoxFilterFuncData),
	}
)

// SetFilterFunc is a wrapper around gtk_list_box_set_filter_func
func (v *ListBox) SetFilterFunc(fn ListBoxFilterFunc, userData ...interface{}) {
	// TODO: figure out a way to determine when we can clean up
	listBoxFilterFuncRegistry.Lock()
	id := listBoxFilterFuncRegistry.next
	listBoxFilterFuncRegistry.next++
	listBoxFilterFuncRegistry.m[id] = listBoxFilterFuncData{fn: fn, userData: userData}
	listBoxFilterFuncRegistry.Unlock()

	C._gtk_list_box_set_filter_func(v.native(), C.gpointer(uintptr(id)))
}

// ListBoxHeaderFunc is a representation of GtkListBoxUpdateHeaderFunc
type ListBoxHeaderFunc func(row *ListBoxRow, before *ListBoxRow, userData ...interface{})

type listBoxHeaderFuncData struct {
	fn       ListBoxHeaderFunc
	userData []interface{}
}

var (
	listBoxHeaderFuncRegistry = struct {
		sync.RWMutex
		next int
		m    map[int]listBoxHeaderFuncData
	}{
		next: 1,
		m:    make(map[int]listBoxHeaderFuncData),
	}
)

// SetHeaderFunc is a wrapper around gtk_list_box_set_header_func
func (v *ListBox) SetHeaderFunc(fn ListBoxHeaderFunc, userData ...interface{}) {
	// TODO: figure out a way to determine when we can clean up
	listBoxHeaderFuncRegistry.Lock()
	id := listBoxHeaderFuncRegistry.next
	listBoxHeaderFuncRegistry.next++
	listBoxHeaderFuncRegistry.m[id] = listBoxHeaderFuncData{fn: fn, userData: userData}
	listBoxHeaderFuncRegistry.Unlock()

	C._gtk_list_box_set_header_func(v.native(), C.gpointer(uintptr(id)))
}

// ListBoxSortFunc is a representation of GtkListBoxSortFunc
type ListBoxSortFunc func(row1 *ListBoxRow, row2 *ListBoxRow, userData ...interface{}) int

type listBoxSortFuncData struct {
	fn       ListBoxSortFunc
	userData []interface{}
}

var (
	listBoxSortFuncRegistry = struct {
		sync.RWMutex
		next int
		m    map[int]listBoxSortFuncData
	}{
		next: 1,
		m:    make(map[int]listBoxSortFuncData),
	}
)

// SetSortFunc is a wrapper around gtk_list_box_set_sort_func
func (v *ListBox) SetSortFunc(fn ListBoxSortFunc, userData ...interface{}) {
	// TODO: figure out a way to determine when we can clean up
	listBoxSortFuncRegistry.Lock()
	id := listBoxSortFuncRegistry.next
	listBoxSortFuncRegistry.next++
	listBoxSortFuncRegistry.m[id] = listBoxSortFuncData{fn: fn, userData: userData}
	listBoxSortFuncRegistry.Unlock()

	C._gtk_list_box_set_sort_func(v.native(), C.gpointer(uintptr(id)))
}

// DragHighlightRow is a wrapper around gtk_list_box_drag_highlight_row()
func (v *ListBox) DragHighlightRow(row *ListBoxRow) {
	C.gtk_list_box_drag_highlight_row(v.native(), row.native())
}

// DragUnhighlightRow is a wrapper around gtk_list_box_drag_unhighlight_row().
func (v *ListBox) DragUnhighlightRow() {
	C.gtk_list_box_drag_unhighlight_row(v.native())
}

/*
 * GtkListBoxRow
 */

// ListBoxRow is a representation of GTK's GtkListBoxRow.
type ListBoxRow struct {
	Bin
}

// native returns a pointer to the underlying GtkListBoxRow.
func (v *ListBoxRow) native() *C.GtkListBoxRow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkListBoxRow(p)
}

func marshalListBoxRow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapListBoxRow(obj), nil
}

func wrapListBoxRow(obj *glib.Object) *ListBoxRow {
	if obj == nil {
		return nil
	}

	return &ListBoxRow{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// ListBoxRowNew is a wrapper around gtk_list_box_row_new().
func ListBoxRowNew() (*ListBoxRow, error) {
	c := C.gtk_list_box_row_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapListBoxRow(glib.Take(unsafe.Pointer(c))), nil
}

// Changed is a wrapper around gtk_list_box_row_changed().
func (v *ListBoxRow) Changed() {
	C.gtk_list_box_row_changed(v.native())
}

// GetHeader is a wrapper around gtk_list_box_row_get_header().
func (v *ListBoxRow) GetHeader() (IWidget, error) {
	c := C.gtk_list_box_row_get_header(v.native())
	if c == nil {
		return nil, nil
	}
	return castWidget(c)
}

// SetHeader is a wrapper around gtk_list_box_row_set_header().
func (v *ListBoxRow) SetHeader(header IWidget) {
	C.gtk_list_box_row_set_header(v.native(), header.toWidget())
}

// GetIndex is a wrapper around gtk_list_box_row_get_index()
func (v *ListBoxRow) GetIndex() int {
	c := C.gtk_list_box_row_get_index(v.native())
	return int(c)
}

/*
 * GtkPlacesSidebar
 */

// TODO:
// gtk_places_sidebar_new().
// gtk_places_sidebar_set_open_flags().
// gtk_places_sidebar_get_open_flags().
// gtk_places_sidebar_set_location().
// gtk_places_sidebar_get_location().
// gtk_places_sidebar_set_show_desktop().
// gtk_places_sidebar_get_show_desktop().
// gtk_places_sidebar_add_shortcut().
// gtk_places_sidebar_remove_shortcut().
// gtk_places_sidebar_list_shortcuts().
// gtk_places_sidebar_get_nth_bookmark().
// enum GtkPlacesOpenFlags

/*
 * GtkRevealer
 */

// Revealer is a representation of GTK's GtkRevealer
type Revealer struct {
	Bin
}

// native returns a pointer to the underlying GtkRevealer.
func (v *Revealer) native() *C.GtkRevealer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkRevealer(p)
}

func marshalRevealer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapRevealer(obj), nil
}

func wrapRevealer(obj *glib.Object) *Revealer {
	if obj == nil {
		return nil
	}

	return &Revealer{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// RevealerNew is a wrapper around gtk_revealer_new()
func RevealerNew() (*Revealer, error) {
	c := C.gtk_revealer_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapRevealer(glib.Take(unsafe.Pointer(c))), nil
}

// GetRevealChild is a wrapper around gtk_revealer_get_reveal_child().
func (v *Revealer) GetRevealChild() bool {
	c := C.gtk_revealer_get_reveal_child(v.native())
	return gobool(c)
}

// SetRevealChild is a wrapper around gtk_revealer_set_reveal_child().
func (v *Revealer) SetRevealChild(revealChild bool) {
	C.gtk_revealer_set_reveal_child(v.native(), gbool(revealChild))
}

// GetChildRevealed is a wrapper around gtk_revealer_get_child_revealed().
func (v *Revealer) GetChildRevealed() bool {
	c := C.gtk_revealer_get_child_revealed(v.native())
	return gobool(c)
}

// GetTransitionDuration is a wrapper around gtk_revealer_get_transition_duration()
func (v *Revealer) GetTransitionDuration() uint {
	c := C.gtk_revealer_get_transition_duration(v.native())
	return uint(c)
}

// SetTransitionDuration is a wrapper around gtk_revealer_set_transition_duration().
func (v *Revealer) SetTransitionDuration(duration uint) {
	C.gtk_revealer_set_transition_duration(v.native(), C.guint(duration))
}

// GetTransitionType is a wrapper around gtk_revealer_get_transition_type()
func (v *Revealer) GetTransitionType() RevealerTransitionType {
	c := C.gtk_revealer_get_transition_type(v.native())
	return RevealerTransitionType(c)
}

// SetTransitionType is a wrapper around gtk_revealer_set_transition_type()
func (v *Revealer) SetTransitionType(transition RevealerTransitionType) {
	t := C.GtkRevealerTransitionType(transition)
	C.gtk_revealer_set_transition_type(v.native(), t)
}

/*
 * GtkSearchBar
 */

// SearchBar is a representation of GTK's GtkSearchBar.
type SearchBar struct {
	Bin
}

// native returns a pointer to the underlying GtkSearchBar.
func (v *SearchBar) native() *C.GtkSearchBar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSearchBar(p)
}

func marshalSearchBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSearchBar(obj), nil
}

func wrapSearchBar(obj *glib.Object) *SearchBar {
	if obj == nil {
		return nil
	}

	return &SearchBar{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// SearchBarNew is a wrapper around gtk_search_bar_new()
func SearchBarNew() (*SearchBar, error) {
	c := C.gtk_search_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSearchBar(glib.Take(unsafe.Pointer(c))), nil
}

// ConnectEntry is a wrapper around gtk_search_bar_connect_entry().
func (v *SearchBar) ConnectEntry(entry IEntry) {
	C.gtk_search_bar_connect_entry(v.native(), entry.toEntry())
}

// GetSearchMode is a wrapper around gtk_search_bar_get_search_mode().
func (v *SearchBar) GetSearchMode() bool {
	c := C.gtk_search_bar_get_search_mode(v.native())
	return gobool(c)
}

// SetSearchMode is a wrapper around gtk_search_bar_set_search_mode().
func (v *SearchBar) SetSearchMode(searchMode bool) {
	C.gtk_search_bar_set_search_mode(v.native(), gbool(searchMode))
}

// GetShowCloseButton is a wrapper arounb gtk_search_bar_get_show_close_button().
func (v *SearchBar) GetShowCloseButton() bool {
	c := C.gtk_search_bar_get_show_close_button(v.native())
	return gobool(c)
}

// SetShowCloseButton is a wrapper around gtk_search_bar_set_show_close_button()
func (v *SearchBar) SetShowCloseButton(visible bool) {
	C.gtk_search_bar_set_show_close_button(v.native(), gbool(visible))
}

// HandleEvent is a wrapper around gtk_search_bar_handle_event()
func (v *SearchBar) HandleEvent(event *gdk.Event) {
	e := (*C.GdkEvent)(unsafe.Pointer(event.Native()))
	C.gtk_search_bar_handle_event(v.native(), e)
}

/*
 * GtkStack
 */

// Stack is a representation of GTK's GtkStack.
type Stack struct {
	Container
}

// native returns a pointer to the underlying GtkStack.
func (v *Stack) native() *C.GtkStack {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkStack(p)
}

func marshalStack(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapStack(obj), nil
}

func wrapStack(obj *glib.Object) *Stack {
	if obj == nil {
		return nil
	}

	return &Stack{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// StackNew is a wrapper around gtk_stack_new().
func StackNew() (*Stack, error) {
	c := C.gtk_stack_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapStack(glib.Take(unsafe.Pointer(c))), nil
}

// AddNamed is a wrapper around gtk_stack_add_named().
func (v *Stack) AddNamed(child IWidget, name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_stack_add_named(v.native(), child.toWidget(), (*C.gchar)(cstr))
}

// AddTitled is a wrapper around gtk_stack_add_titled().
func (v *Stack) AddTitled(child IWidget, name, title string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.gtk_stack_add_titled(v.native(), child.toWidget(), (*C.gchar)(cName),
		(*C.gchar)(cTitle))
}

// SetVisibleChild is a wrapper around gtk_stack_set_visible_child().
func (v *Stack) SetVisibleChild(child IWidget) {
	C.gtk_stack_set_visible_child(v.native(), child.toWidget())
}

// GetVisibleChild is a wrapper around gtk_stack_get_visible_child().
func (v *Stack) GetVisibleChild() (IWidget, error) {
	c := C.gtk_stack_get_visible_child(v.native())
	if c == nil {
		return nil, nil
	}
	return castWidget(c)
}

// SetVisibleChildName is a wrapper around gtk_stack_set_visible_child_name().
func (v *Stack) SetVisibleChildName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_stack_set_visible_child_name(v.native(), (*C.gchar)(cstr))
}

// GetVisibleChildName is a wrapper around gtk_stack_get_visible_child_name().
func (v *Stack) GetVisibleChildName() string {
	c := C.gtk_stack_get_visible_child_name(v.native())
	return C.GoString((*C.char)(c))
}

// SetVisibleChildFull is a wrapper around gtk_stack_set_visible_child_full().
func (v *Stack) SetVisibleChildFull(name string, transaction StackTransitionType) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_stack_set_visible_child_full(v.native(), (*C.gchar)(cstr),
		C.GtkStackTransitionType(transaction))
}

// SetHomogeneous is a wrapper around gtk_stack_set_homogeneous().
func (v *Stack) SetHomogeneous(homogeneous bool) {
	C.gtk_stack_set_homogeneous(v.native(), gbool(homogeneous))
}

// GetHomogeneous is a wrapper around gtk_stack_get_homogeneous().
func (v *Stack) GetHomogeneous() bool {
	c := C.gtk_stack_get_homogeneous(v.native())
	return gobool(c)
}

// SetTransitionDuration is a wrapper around gtk_stack_set_transition_duration().
func (v *Stack) SetTransitionDuration(duration uint) {
	C.gtk_stack_set_transition_duration(v.native(), C.guint(duration))
}

// GetTransitionDuration is a wrapper around gtk_stack_get_transition_duration().
func (v *Stack) GetTransitionDuration() uint {
	c := C.gtk_stack_get_transition_duration(v.native())
	return uint(c)
}

// SetTransitionType is a wrapper around gtk_stack_set_transition_type().
func (v *Stack) SetTransitionType(transition StackTransitionType) {
	C.gtk_stack_set_transition_type(v.native(), C.GtkStackTransitionType(transition))
}

// GetTransitionType is a wrapper around gtk_stack_get_transition_type().
func (v *Stack) GetTransitionType() StackTransitionType {
	c := C.gtk_stack_get_transition_type(v.native())
	return StackTransitionType(c)
}

/*
 * GtkBuilder
 */

// BuilderNewFromFile is a wrapper around gtk_builder_new_from_file().
func BuilderNewFromFile(filePath string) (*Builder, error) {
	cstr := C.CString(filePath)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_builder_new_from_file((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return &Builder{obj}, nil
}

// BuilderNewFromResource is a wrapper around gtk_builder_new_from_resource().
func BuilderNewFromResource(resourcePath string) (*Builder, error) {
	cstr := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_builder_new_from_resource((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return &Builder{obj}, nil
}

// BuilderNewFromString is a wrapper around gtk_builder_new_from_string().
func BuilderNewFromString(resource string) (*Builder, error) {
	cstr := C.CString(resource)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_builder_new_from_string((*C.gchar)(cstr), C.gssize(len(resource)))
	if c == nil {
		return nil, nilPtrErr
	}

	obj := glib.Take(unsafe.Pointer(c))
	return &Builder{obj}, nil
}

// TODO:
// gtk_builder_add_callback_symbol
// gtk_builder_add_callback_symbols
// gtk_builder_lookup_callback_symbol
// gtk_builder_set_application
// gtk_builder_get_application

// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

// +build !gtk_3_6,!gtk_3_8
// not use this: go build -tags gtk_3_8'. Otherwise, if no build tags are used, GTK 3.10

package gtkf

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
// #include "gtk_since_3_10.go.h"
import "C"
import (
	"unsafe"

	gdk_impl "github.com/gotk3/gotk3/gdkf"
	"github.com/gotk3/gotk3/glib"
	glib_impl "github.com/gotk3/gotk3/glibf"
	"github.com/gotk3/gotk3/gtk"
)

func init() {
	tm := []glib_impl.TypeMarshaler{
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
	glib_impl.RegisterGValueMarshalers(tm)

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

	gtk.ALIGN_BASELINE = C.GTK_ALIGN_BASELINE

	gtk.REVEALER_TRANSITION_TYPE_NONE = C.GTK_REVEALER_TRANSITION_TYPE_NONE
	gtk.REVEALER_TRANSITION_TYPE_CROSSFADE = C.GTK_REVEALER_TRANSITION_TYPE_CROSSFADE
	gtk.REVEALER_TRANSITION_TYPE_SLIDE_RIGHT = C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_RIGHT
	gtk.REVEALER_TRANSITION_TYPE_SLIDE_LEFT = C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_LEFT
	gtk.REVEALER_TRANSITION_TYPE_SLIDE_UP = C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_UP
	gtk.REVEALER_TRANSITION_TYPE_SLIDE_DOWN = C.GTK_REVEALER_TRANSITION_TYPE_SLIDE_DOWN

	gtk.STACK_TRANSITION_TYPE_NONE = C.GTK_STACK_TRANSITION_TYPE_NONE
	gtk.STACK_TRANSITION_TYPE_CROSSFADE = C.GTK_STACK_TRANSITION_TYPE_CROSSFADE
	gtk.STACK_TRANSITION_TYPE_SLIDE_RIGHT = C.GTK_STACK_TRANSITION_TYPE_SLIDE_RIGHT
	gtk.STACK_TRANSITION_TYPE_SLIDE_LEFT = C.GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT
	gtk.STACK_TRANSITION_TYPE_SLIDE_UP = C.GTK_STACK_TRANSITION_TYPE_SLIDE_UP
	gtk.STACK_TRANSITION_TYPE_SLIDE_DOWN = C.GTK_STACK_TRANSITION_TYPE_SLIDE_DOWN
	gtk.STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT = C.GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT
	gtk.STACK_TRANSITION_TYPE_SLIDE_UP_DOWN = C.GTK_STACK_TRANSITION_TYPE_SLIDE_UP_DOWN
}

func marshalRevealerTransitionType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.RevealerTransitionType(c), nil
}

func marshalStackTransitionType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.StackTransitionType(c), nil
}

/*
 * GtkButton
 */

// ButtonNewFromIconName is a wrapper around gtk_button_new_from_icon_name().
func ButtonNewFromIconName(iconName string, size gtk.IconSize) (*button, error) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_button_new_from_icon_name((*C.gchar)(cstr),
		C.GtkIconSize(size))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapButton(wrapObject(unsafe.Pointer(c))), nil
}

/*
 * GtkHeaderBar
 */

type headerBar struct {
	container
}

// native returns a pointer to the underlying GtkHeaderBar.
func (v *headerBar) native() *C.GtkHeaderBar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkHeaderBar(p)
}

func marshalHeaderBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapHeaderBar(obj), nil
}

func wrapHeaderBar(obj *glib_impl.Object) *headerBar {
	return &headerBar{container{widget{glib_impl.InitiallyUnowned{obj}}}}
}

// HeaderBarNew is a wrapper around gtk_header_bar_new().
func HeaderBarNew() (*headerBar, error) {
	c := C.gtk_header_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapHeaderBar(wrapObject(unsafe.Pointer(c))), nil
}

// SetTitle is a wrapper around gtk_header_bar_set_title().
func (v *headerBar) SetTitle(title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_header_bar_set_title(v.native(), (*C.gchar)(cstr))
}

// GetTitle is a wrapper around gtk_header_bar_get_title().
func (v *headerBar) GetTitle() string {
	cstr := C.gtk_header_bar_get_title(v.native())
	return C.GoString((*C.char)(cstr))
}

// SetSubtitle is a wrapper around gtk_header_bar_set_subtitle().
func (v *headerBar) SetSubtitle(subtitle string) {
	cstr := C.CString(subtitle)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_header_bar_set_subtitle(v.native(), (*C.gchar)(cstr))
}

// GetSubtitle is a wrapper around gtk_header_bar_get_subtitle().
func (v *headerBar) GetSubtitle() string {
	cstr := C.gtk_header_bar_get_subtitle(v.native())
	return C.GoString((*C.char)(cstr))
}

// SetCustomTitle is a wrapper around gtk_header_bar_set_custom_title().
func (v *headerBar) SetCustomTitle(titleWidget gtk.Widget) {
	C.gtk_header_bar_set_custom_title(v.native(), titleWidget.(IWidget).toWidget())
}

// GetCustomTitle is a wrapper around gtk_header_bar_get_custom_title().
func (v *headerBar) GetCustomTitle() (*widget, error) {
	c := C.gtk_header_bar_get_custom_title(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWidget(wrapObject(unsafe.Pointer(c))), nil
}

// PackStart is a wrapper around gtk_header_bar_pack_start().
func (v *headerBar) PackStart(child gtk.Widget) {
	C.gtk_header_bar_pack_start(v.native(), child.(IWidget).toWidget())
}

// PackEnd is a wrapper around gtk_header_bar_pack_end().
func (v *headerBar) PackEnd(child gtk.Widget) {
	C.gtk_header_bar_pack_end(v.native(), child.(IWidget).toWidget())
}

// SetShowCloseButton is a wrapper around gtk_header_bar_set_show_close_button().
func (v *headerBar) SetShowCloseButton(setting bool) {
	C.gtk_header_bar_set_show_close_button(v.native(), gbool(setting))
}

// GetShowCloseButton is a wrapper around gtk_header_bar_get_show_close_button().
func (v *headerBar) GetShowCloseButton() bool {
	c := C.gtk_header_bar_get_show_close_button(v.native())
	return gobool(c)
}

/*
 * GtkLabel
 */

// GetLines() is a wrapper around gtk_label_get_lines().
func (v *label) GetLines() int {
	c := C.gtk_label_get_lines(v.native())
	return int(c)
}

// SetLines() is a wrapper around gtk_label_set_lines().
func (v *label) SetLines(lines int) {
	C.gtk_label_set_lines(v.native(), C.gint(lines))
}

/*
 * GtkListBox
 */

// ListBox is a representation of GTK's GtkListBox.
type listBox struct {
	container
}

// native returns a pointer to the underlying GtkListBox.
func (v *listBox) native() *C.GtkListBox {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkListBox(p)
}

func marshalListBox(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapListBox(obj), nil
}

func wrapListBox(obj *glib_impl.Object) *listBox {
	return &listBox{container{widget{glib_impl.InitiallyUnowned{obj}}}}
}

// ListBoxNew is a wrapper around gtk_list_box_new().
func ListBoxNew() (*listBox, error) {
	c := C.gtk_list_box_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapListBox(wrapObject(unsafe.Pointer(c))), nil
}

// Prepend is a wrapper around gtk_list_box_prepend().
func (v *listBox) Prepend(child gtk.Widget) {
	C.gtk_list_box_prepend(v.native(), child.(IWidget).toWidget())
}

// Insert is a wrapper around gtk_list_box_insert().
func (v *listBox) Insert(child gtk.Widget, position int) {
	C.gtk_list_box_insert(v.native(), child.(IWidget).toWidget(), C.gint(position))
}

// SelectRow is a wrapper around gtk_list_box_select_row().
func (v *listBox) SelectRow(row *listBoxRow) {
	C.gtk_list_box_select_row(v.native(), row.native())
}

// GetSelectedRow is a wrapper around gtk_list_box_get_selected_row().
func (v *listBox) GetSelectedRow() *listBoxRow {
	c := C.gtk_list_box_get_selected_row(v.native())
	if c == nil {
		return nil
	}
	return wrapListBoxRow(wrapObject(unsafe.Pointer(c)))
}

// SetSelectionMode is a wrapper around gtk_list_box_set_selection_mode().
func (v *listBox) SetSelectionMode(mode gtk.SelectionMode) {
	C.gtk_list_box_set_selection_mode(v.native(), C.GtkSelectionMode(mode))
}

// GetSelectionMode is a wrapper around gtk_list_box_get_selection_mode()
func (v *listBox) GetSelectionMode() gtk.SelectionMode {
	c := C.gtk_list_box_get_selection_mode(v.native())
	return gtk.SelectionMode(c)
}

// SetActivateOnSingleClick is a wrapper around gtk_list_box_set_activate_on_single_click().
func (v *listBox) SetActivateOnSingleClick(single bool) {
	C.gtk_list_box_set_activate_on_single_click(v.native(), gbool(single))
}

// GetActivateOnSingleClick is a wrapper around gtk_list_box_get_activate_on_single_click().
func (v *listBox) GetActivateOnSingleClick() bool {
	c := C.gtk_list_box_get_activate_on_single_click(v.native())
	return gobool(c)
}

// GetAdjustment is a wrapper around gtk_list_box_get_adjustment().
func (v *listBox) GetAdjustment() *adjustment {
	c := C.gtk_list_box_get_adjustment(v.native())
	obj := wrapObject(unsafe.Pointer(c))
	return &adjustment{glib_impl.InitiallyUnowned{obj}}
}

// SetAdjustment is a wrapper around gtk_list_box_set_adjustment().
func (v *listBox) SetAdjuctment(adjustment *adjustment) {
	C.gtk_list_box_set_adjustment(v.native(), adjustment.native())
}

// SetPlaceholder is a wrapper around gtk_list_box_set_placeholder().
func (v *listBox) SetPlaceholder(placeholder gtk.Widget) {
	C.gtk_list_box_set_placeholder(v.native(), placeholder.(IWidget).toWidget())
}

// GetRowAtIndex is a wrapper around gtk_list_box_get_row_at_index().
func (v *listBox) GetRowAtIndex(index int) *listBoxRow {
	c := C.gtk_list_box_get_row_at_index(v.native(), C.gint(index))
	if c == nil {
		return nil
	}
	return wrapListBoxRow(wrapObject(unsafe.Pointer(c)))
}

// GetRowAtY is a wrapper around gtk_list_box_get_row_at_y().
func (v *listBox) GetRowAtY(y int) *listBoxRow {
	c := C.gtk_list_box_get_row_at_y(v.native(), C.gint(y))
	if c == nil {
		return nil
	}
	return wrapListBoxRow(wrapObject(unsafe.Pointer(c)))
}

// InvalidateFilter is a wrapper around gtk_list_box_invalidate_filter().
func (v *listBox) InvalidateFilter() {
	C.gtk_list_box_invalidate_filter(v.native())
}

// InvalidateHeaders is a wrapper around gtk_list_box_invalidate_headers().
func (v *listBox) InvalidateHeaders() {
	C.gtk_list_box_invalidate_headers(v.native())
}

// InvalidateSort is a wrapper around gtk_list_box_invalidate_sort().
func (v *listBox) InvalidateSort() {
	C.gtk_list_box_invalidate_sort(v.native())
}

// TODO: SetFilterFunc
// TODO: SetHeaderFunc
// TODO: SetSortFunc

// DragHighlightRow is a wrapper around gtk_list_box_drag_highlight_row()
func (v *listBox) DragHighlightRow(row *listBoxRow) {
	C.gtk_list_box_drag_highlight_row(v.native(), row.native())
}

/*
 * GtkListBoxRow
 */

// ListBoxRow is a representation of GTK's GtkListBoxRow.
type listBoxRow struct {
	bin
}

// native returns a pointer to the underlying GtkListBoxRow.
func (v *listBoxRow) native() *C.GtkListBoxRow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkListBoxRow(p)
}

func marshalListBoxRow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapListBoxRow(obj), nil
}

func wrapListBoxRow(obj *glib_impl.Object) *listBoxRow {
	return &listBoxRow{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}
}

func ListBoxRowNew() (*listBoxRow, error) {
	c := C.gtk_list_box_row_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapListBoxRow(wrapObject(unsafe.Pointer(c))), nil
}

// Changed is a wrapper around gtk_list_box_row_changed().
func (v *listBoxRow) Changed() {
	C.gtk_list_box_row_changed(v.native())
}

// GetHeader is a wrapper around gtk_list_box_row_get_header().
func (v *listBoxRow) GetHeader() *widget {
	c := C.gtk_list_box_row_get_header(v.native())
	if c == nil {
		return nil
	}
	return wrapWidget(wrapObject(unsafe.Pointer(c)))
}

// SetHeader is a wrapper around gtk_list_box_row_get_header().
func (v *listBoxRow) SetHeader(header gtk.Widget) {
	C.gtk_list_box_row_set_header(v.native(), header.(IWidget).toWidget())
}

// GetIndex is a wrapper around gtk_list_box_row_get_index()
func (v *listBoxRow) GetIndex() int {
	c := C.gtk_list_box_row_get_index(v.native())
	return int(c)
}

/*
 * GtkRevealer
 */

// Revealer is a representation of GTK's GtkRevealer
type revealer struct {
	bin
}

// native returns a pointer to the underlying GtkRevealer.
func (v *revealer) native() *C.GtkRevealer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkRevealer(p)
}

func marshalRevealer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapRevealer(obj), nil
}

func wrapRevealer(obj *glib_impl.Object) *revealer {
	return &revealer{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}
}

// RevealerNew is a wrapper around gtk_revealer_new()
func RevealerNew() (*revealer, error) {
	c := C.gtk_revealer_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapRevealer(wrapObject(unsafe.Pointer(c))), nil
}

// GetRevealChild is a wrapper around gtk_revealer_get_reveal_child().
func (v *revealer) GetRevealChild() bool {
	c := C.gtk_revealer_get_reveal_child(v.native())
	return gobool(c)
}

// SetRevealChild is a wrapper around gtk_revealer_set_reveal_child().
func (v *revealer) SetRevealChild(revealChild bool) {
	C.gtk_revealer_set_reveal_child(v.native(), gbool(revealChild))
}

// GetChildRevealed is a wrapper around gtk_revealer_get_child_revealed().
func (v *revealer) GetChildRevealed() bool {
	c := C.gtk_revealer_get_child_revealed(v.native())
	return gobool(c)
}

// GetTransitionDuration is a wrapper around gtk_revealer_get_transition_duration()
func (v *revealer) GetTransitionDuration() uint {
	c := C.gtk_revealer_get_transition_duration(v.native())
	return uint(c)
}

// SetTransitionDuration is a wrapper around gtk_revealer_set_transition_duration().
func (v *revealer) SetTransitionDuration(duration uint) {
	C.gtk_revealer_set_transition_duration(v.native(), C.guint(duration))
}

// GetTransitionType is a wrapper around gtk_revealer_get_transition_type()
func (v *revealer) GetTransitionType() gtk.RevealerTransitionType {
	c := C.gtk_revealer_get_transition_type(v.native())
	return gtk.RevealerTransitionType(c)
}

// SetTransitionType is a wrapper around gtk_revealer_set_transition_type()
func (v *revealer) SetTransitionType(transition gtk.RevealerTransitionType) {
	t := C.GtkRevealerTransitionType(transition)
	C.gtk_revealer_set_transition_type(v.native(), t)
}

/*
 * GtkSearchBar
 */

// SearchBar is a representation of GTK's GtkSearchBar.
type searchBar struct {
	bin
}

// native returns a pointer to the underlying GtkSearchBar.
func (v *searchBar) native() *C.GtkSearchBar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSearchBar(p)
}

func marshalSearchBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapSearchBar(obj), nil
}

func wrapSearchBar(obj *glib_impl.Object) *searchBar {
	return &searchBar{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}
}

// SearchBarNew is a wrapper around gtk_search_bar_new()
func SearchBarNew() (*searchBar, error) {
	c := C.gtk_search_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSearchBar(wrapObject(unsafe.Pointer(c))), nil
}

// ConnectEntry is a wrapper around gtk_search_bar_connect_entry().
func (v *searchBar) ConnectEntry(entry IEntry) {
	C.gtk_search_bar_connect_entry(v.native(), entry.toEntry())
}

// GetSearchMode is a wrapper around gtk_search_bar_get_search_mode().
func (v *searchBar) GetSearchMode() bool {
	c := C.gtk_search_bar_get_search_mode(v.native())
	return gobool(c)
}

// SetSearchMode is a wrapper around gtk_search_bar_set_search_mode().
func (v *searchBar) SetSearchMode(searchMode bool) {
	C.gtk_search_bar_set_search_mode(v.native(), gbool(searchMode))
}

// GetShowCloseButton is a wrapper arounb gtk_search_bar_get_show_close_button().
func (v *searchBar) GetShowCloseButton() bool {
	c := C.gtk_search_bar_get_show_close_button(v.native())
	return gobool(c)
}

// SetShowCloseButton is a wrapper around gtk_search_bar_set_show_close_button()
func (v *searchBar) SetShowCloseButton(visible bool) {
	C.gtk_search_bar_set_show_close_button(v.native(), gbool(visible))
}

// HandleEvent is a wrapper around gtk_search_bar_handle_event()
func (v *searchBar) HandleEvent(event *gdk_impl.Event) {
	e := (*C.GdkEvent)(unsafe.Pointer(event.Native()))
	C.gtk_search_bar_handle_event(v.native(), e)
}

/*
 * GtkStack
 */

// Stack is a representation of GTK's GtkStack.
type stack struct {
	container
}

// native returns a pointer to the underlying GtkStack.
func (v *stack) native() *C.GtkStack {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkStack(p)
}

func marshalStack(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapStack(obj), nil
}

func wrapStack(obj *glib_impl.Object) *stack {
	return &stack{container{widget{glib_impl.InitiallyUnowned{obj}}}}
}

// StackNew is a wrapper around gtk_stack_new().
func StackNew() (*stack, error) {
	c := C.gtk_stack_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapStack(wrapObject(unsafe.Pointer(c))), nil
}

// AddNamed is a wrapper around gtk_stack_add_named().
func (v *stack) AddNamed(child gtk.Widget, name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_stack_add_named(v.native(), child.(IWidget).toWidget(), (*C.gchar)(cstr))
}

// AddTitled is a wrapper around gtk_stack_add_titled().
func (v *stack) AddTitled(child gtk.Widget, name, title string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.gtk_stack_add_titled(v.native(), child.(IWidget).toWidget(), (*C.gchar)(cName),
		(*C.gchar)(cTitle))
}

// SetVisibleChild is a wrapper around gtk_stack_set_visible_child().
func (v *stack) SetVisibleChild(child gtk.Widget) {
	C.gtk_stack_set_visible_child(v.native(), child.(IWidget).toWidget())
}

// GetVisibleChild is a wrapper around gtk_stack_get_visible_child().
func (v *stack) GetVisibleChild() *widget {
	c := C.gtk_stack_get_visible_child(v.native())
	if c == nil {
		return nil
	}
	return wrapWidget(wrapObject(unsafe.Pointer(c)))
}

// SetVisibleChildName is a wrapper around gtk_stack_set_visible_child_name().
func (v *stack) SetVisibleChildName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_stack_set_visible_child_name(v.native(), (*C.gchar)(cstr))
}

// GetVisibleChildName is a wrapper around gtk_stack_get_visible_child_name().
func (v *stack) GetVisibleChildName() string {
	c := C.gtk_stack_get_visible_child_name(v.native())
	return C.GoString((*C.char)(c))
}

// SetVisibleChildFull is a wrapper around gtk_stack_set_visible_child_full().
func (v *stack) SetVisibleChildFull(name string, transaction gtk.StackTransitionType) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_stack_set_visible_child_full(v.native(), (*C.gchar)(cstr),
		C.GtkStackTransitionType(transaction))
}

// SetHomogeneous is a wrapper around gtk_stack_set_homogeneous().
func (v *stack) SetHomogeneous(homogeneous bool) {
	C.gtk_stack_set_homogeneous(v.native(), gbool(homogeneous))
}

// GetHomogeneous is a wrapper around gtk_stack_get_homogeneous().
func (v *stack) GetHomogeneous() bool {
	c := C.gtk_stack_get_homogeneous(v.native())
	return gobool(c)
}

// SetTransitionDuration is a wrapper around gtk_stack_set_transition_duration().
func (v *stack) SetTransitionDuration(duration uint) {
	C.gtk_stack_set_transition_duration(v.native(), C.guint(duration))
}

// GetTransitionDuration is a wrapper around gtk_stack_get_transition_duration().
func (v *stack) GetTransitionDuration() uint {
	c := C.gtk_stack_get_transition_duration(v.native())
	return uint(c)
}

// SetTransitionType is a wrapper around gtk_stack_set_transition_type().
func (v *stack) SetTransitionType(transition gtk.StackTransitionType) {
	C.gtk_stack_set_transition_type(v.native(), C.GtkStackTransitionType(transition))
}

// GetTransitionType is a wrapper around gtk_stack_get_transition_type().
func (v *stack) GetTransitionType() gtk.StackTransitionType {
	c := C.gtk_stack_get_transition_type(v.native())
	return gtk.StackTransitionType(c)
}

// Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
//
// This file originated from: http://opensource.conformal.com/
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// This file includes wrapers for symbols included since GTK 3.10, and
// and should not be included in a build intended to target any older GTK
// versions.  To target an older build, such as 3.8, use
// 'go build -tags gtk_3_8'.  Otherwise, if no build tags are used, GTK 3.10
// is assumed and this file is built.
// +build !gtk_3_6,!gtk_3_8

package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
// #include "gtk_3_10-12.go.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/andre-hub/gotk3/gdk"
	"github.com/andre-hub/gotk3/glib"
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
		{glib.Type(C.gtk_alignment_get_type()), marshalAlignment},
		{glib.Type(C.gtk_arrow_get_type()), marshalArrow},
		{glib.Type(C.gtk_misc_get_type()), marshalMisc},
		{glib.Type(C.gtk_status_icon_get_type()), marshalStatusIcon},
	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * Constants
 */

const (
	ALIGN_BASELINE Align = C.GTK_ALIGN_BASELINE
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
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return b, nil
}

/*
 * GtkHeaderBar
 */

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
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapHeaderBar(obj), nil
}

func wrapHeaderBar(obj *glib.Object) *HeaderBar {
	return &HeaderBar{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// HeaderBarNew is a wrapper around gtk_header_bar_new().
func HeaderBarNew() (*HeaderBar, error) {
	c := C.gtk_header_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	h := wrapHeaderBar(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return h, nil
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
func (v *HeaderBar) GetCustomTitle() (*Widget, error) {
	c := C.gtk_header_bar_get_custom_title(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
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

// GetLines() is a wrapper around gtk_label_get_lines().
func (v *Label) GetLines() int {
	c := C.gtk_label_get_lines(v.native())
	return int(c)
}

// SetLines() is a wrapper around gtk_label_set_lines().
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
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapListBox(obj), nil
}

func wrapListBox(obj *glib.Object) *ListBox {
	return &ListBox{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// ListBoxNew is a wrapper around gtk_list_box_new().
func ListBoxNew() (*ListBox, error) {
	c := C.gtk_list_box_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	l := wrapListBox(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return l, nil
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
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	l := wrapListBoxRow(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return l
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
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return &Adjustment{glib.InitiallyUnowned{obj}}
}

// SetAdjustment is a wrapper around gtk_list_box_set_adjustment().
func (v *ListBox) SetAdjuctment(adjustment *Adjustment) {
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
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	l := wrapListBoxRow(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return l
}

// GetRowAtY is a wrapper around gtk_list_box_get_row_at_y().
func (v *ListBox) GetRowAtY(y int) *ListBoxRow {
	c := C.gtk_list_box_get_row_at_y(v.native(), C.gint(y))
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	l := wrapListBoxRow(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return l
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

// TODO: SetFilterFunc
// TODO: SetHeaderFunc
// TODO: SetSortFunc

// DragHighlightRow is a wrapper around gtk_list_box_drag_highlight_row()
func (v *ListBox) DragHighlightRow(row *ListBoxRow) {
	C.gtk_list_box_drag_highlight_row(v.native(), row.native())
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
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapListBoxRow(obj), nil
}

func wrapListBoxRow(obj *glib.Object) *ListBoxRow {
	return &ListBoxRow{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

func ListBoxRowNew() (*ListBoxRow, error) {
	c := C.gtk_list_box_row_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapListBoxRow(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// Changed is a wrapper around gtk_list_box_row_changed().
func (v *ListBoxRow) Changed() {
	C.gtk_list_box_row_changed(v.native())
}

// GetHeader is a wrapper around gtk_list_box_row_get_header().
func (v *ListBoxRow) GetHeader() *Widget {
	c := C.gtk_list_box_row_get_header(v.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w
}

// SetHeader is a wrapper around gtk_list_box_row_get_header().
func (v *ListBoxRow) SetHeader(header IWidget) {
	C.gtk_list_box_row_set_header(v.native(), header.toWidget())
}

// GetIndex is a wrapper around gtk_list_box_row_get_index()
func (v *ListBoxRow) GetIndex() int {
	c := C.gtk_list_box_row_get_index(v.native())
	return int(c)
}

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
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapRevealer(obj), nil
}

func wrapRevealer(obj *glib.Object) *Revealer {
	return &Revealer{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// RevealerNew is a wrapper around gtk_revealer_new()
func RevealerNew() (*Revealer, error) {
	c := C.gtk_revealer_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	r := wrapRevealer(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return r, nil
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
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapSearchBar(obj), nil
}

func wrapSearchBar(obj *glib.Object) *SearchBar {
	return &SearchBar{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// SearchBarNew is a wrapper around gtk_search_bar_new()
func SearchBarNew() (*SearchBar, error) {
	c := C.gtk_search_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapSearchBar(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
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
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapStack(obj), nil
}

func wrapStack(obj *glib.Object) *Stack {
	return &Stack{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// StackNew is a wrapper around gtk_stack_new().
func StackNew() (*Stack, error) {
	c := C.gtk_stack_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapStack(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
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
func (v *Stack) GetVisibleChild() *Widget {
	c := C.gtk_stack_get_visible_child(v.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s
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
 * GtkStackSwitcher
 */

// StackSwitcher is a representation of GTK's GtkStackSwitcher
type StackSwitcher struct {
	Box
}

// native returns a pointer to the underlying GtkStackSwitcher.
func (v *StackSwitcher) native() *C.GtkStackSwitcher {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkStackSwitcher(p)
}

func marshalStackSwitcher(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapStackSwitcher(obj), nil
}

func wrapStackSwitcher(obj *glib.Object) *StackSwitcher {
	return &StackSwitcher{Box{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// StackSwitcherNew is a wrapper around gtk_stack_switcher_new().
func StackSwitcherNew() (*StackSwitcher, error) {
	c := C.gtk_stack_switcher_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapStackSwitcher(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s, nil
}

// SetStack is a wrapper around gtk_stack_switcher_set_stack().
func (v *StackSwitcher) SetStack(stack *Stack) {
	C.gtk_stack_switcher_set_stack(v.native(), stack.native())
}

// GetStack is a wrapper around gtk_stack_switcher_get_stack().
func (v *StackSwitcher) GetStack() *Stack {
	c := C.gtk_stack_switcher_get_stack(v.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	s := wrapStack(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return s
}

/*
 * GtkWindow
 */

// SetTitlebar is a wrapper around gtk_window_set_titlebar().
func (v *Window) SetTitlebar(titlebar IWidget) {
	C.gtk_window_set_titlebar(v.native(), titlebar.toWidget())
}

// Close is a wrapper around gtk_window_close().
func (v *Window) Close() {
	C.gtk_window_close(v.native())
}

func cast_3_10(class string, o *glib.Object) glib.IObject {
	var g glib.IObject
	switch class {
	case "GtkListBox":
		g = wrapListBox(o)
	case "GtkListBoxRow":
		g = wrapListBoxRow(o)
	case "GtkRevealer":
		g = wrapRevealer(o)
	case "GtkSearchBar":
		g = wrapSearchBar(o)
	case "GtkStack":
		g = wrapStack(o)
	case "GtkStackSwitcher":
		g = wrapStackSwitcher(o)
	case "GtkAlignment":
		g = wrapAlignment(o)
	case "GtkArrow":
		g = wrapArrow(o)
	}
	return g
}

func init() {
	cast_3_10_func = cast_3_10
}

/*
 * deprecated since version 3.14 and should not be used in newly-written code
 */

// ResizeGripIsVisible is a wrapper around
// gtk_window_resize_grip_is_visible().
func (v *Window) ResizeGripIsVisible() bool {
	c := C.gtk_window_resize_grip_is_visible(v.native())
	return gobool(c)
}

// SetHasResizeGrip is a wrapper around gtk_window_set_has_resize_grip().
func (v *Window) SetHasResizeGrip(setting bool) {
	C.gtk_window_set_has_resize_grip(v.native(), gbool(setting))
}

// GetHasResizeGrip is a wrapper around gtk_window_get_has_resize_grip().
func (v *Window) GetHasResizeGrip() bool {
	c := C.gtk_window_get_has_resize_grip(v.native())
	return gobool(c)
}

// Reparent() is a wrapper around gtk_widget_reparent().
func (v *Widget) Reparent(newParent IWidget) {
	C.gtk_widget_reparent(v.native(), newParent.toWidget())
}

// GetPadding is a wrapper around gtk_alignment_get_padding().
func (v *Alignment) GetPadding() (top, bottom, left, right uint) {
	var ctop, cbottom, cleft, cright C.guint
	C.gtk_alignment_get_padding(v.native(), &ctop, &cbottom, &cleft,
		&cright)
	return uint(ctop), uint(cbottom), uint(cleft), uint(cright)
}

// SetPadding is a wrapper around gtk_alignment_set_padding().
func (v *Alignment) SetPadding(top, bottom, left, right uint) {
	C.gtk_alignment_set_padding(v.native(), C.guint(top), C.guint(bottom),
		C.guint(left), C.guint(right))
}

// AlignmentNew is a wrapper around gtk_alignment_new().
func AlignmentNew(xalign, yalign, xscale, yscale float32) (*Alignment, error) {
	c := C.gtk_alignment_new(C.gfloat(xalign), C.gfloat(yalign), C.gfloat(xscale),
		C.gfloat(yscale))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAlignment(obj), nil
}

// Set is a wrapper around gtk_alignment_set().
func (v *Alignment) Set(xalign, yalign, xscale, yscale float32) {
	C.gtk_alignment_set(v.native(), C.gfloat(xalign), C.gfloat(yalign),
		C.gfloat(xscale), C.gfloat(yscale))
}

/*
 * GtkArrow
 */

// Arrow is a representation of GTK's GtkArrow.
type Arrow struct {
	Misc
}

// ArrowNew is a wrapper around gtk_arrow_new().
func ArrowNew(arrowType ArrowType, shadowType ShadowType) (*Arrow, error) {
	c := C.gtk_arrow_new(C.GtkArrowType(arrowType),
		C.GtkShadowType(shadowType))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapArrow(obj), nil
}

// Set is a wrapper around gtk_arrow_set().
func (v *Arrow) Set(arrowType ArrowType, shadowType ShadowType) {
	C.gtk_arrow_set(v.native(), C.GtkArrowType(arrowType), C.GtkShadowType(shadowType))
}

// SetAlignment() is a wrapper around gtk_button_set_alignment().
func (v *Button) SetAlignment(xalign, yalign float32) {
	C.gtk_button_set_alignment(v.native(), (C.gfloat)(xalign),
		(C.gfloat)(yalign))
}

// GetAlignment() is a wrapper around gtk_button_get_alignment().
func (v *Button) GetAlignment() (xalign, yalign float32) {
	var x, y C.gfloat
	C.gtk_button_get_alignment(v.native(), &x, &y)
	return float32(x), float32(y)
}

// SetReallocateRedraws is a wrapper around
// gtk_container_set_reallocate_redraws().
func (v *Container) SetReallocateRedraws(needsRedraws bool) {
	C.gtk_container_set_reallocate_redraws(v.native(), gbool(needsRedraws))
}

// GetAlignment is a wrapper around gtk_misc_get_alignment().
func (v *Misc) GetAlignment() (xAlign, yAlign float32) {
	var x, y C.gfloat
	C.gtk_misc_get_alignment(v.native(), &x, &y)
	return float32(x), float32(y)
}

// SetAlignment is a wrapper around gtk_misc_set_alignment().
func (v *Misc) SetAlignment(xAlign, yAlign float32) {
	C.gtk_misc_set_alignment(v.native(), C.gfloat(xAlign), C.gfloat(yAlign))
}

// GetPadding is a wrapper around gtk_misc_get_padding().
func (v *Misc) GetPadding() (xpad, ypad int) {
	var x, y C.gint
	C.gtk_misc_get_padding(v.native(), &x, &y)
	return int(x), int(y)
}

// SetPadding is a wrapper around gtk_misc_set_padding().
func (v *Misc) SetPadding(xPad, yPad int) {
	C.gtk_misc_set_padding(v.native(), C.gint(xPad), C.gint(yPad))
}

// SetDoubleBuffered is a wrapper around gtk_widget_set_double_buffered().
func (v *Widget) SetDoubleBuffered(doubleBuffered bool) {
	C.gtk_widget_set_double_buffered(v.native(), gbool(doubleBuffered))
}

// GetDoubleBuffered is a wrapper around gtk_widget_get_double_buffered().
func (v *Widget) GetDoubleBuffered() bool {
	c := C.gtk_widget_get_double_buffered(v.native())
	return gobool(c)
}

/*
 * GtkArrow
 * deprecated since version 3.14
 */
// native returns a pointer to the underlying GtkButton.
func (v *Arrow) native() *C.GtkArrow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkArrow(p)
}

func marshalArrow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapArrow(obj), nil
}

func wrapArrow(obj *glib.Object) *Arrow {
	return &Arrow{Misc{Widget{glib.InitiallyUnowned{obj}}}}
}

/*
 * GtkAlignment
 * deprecated since version 3.14
 */

type Alignment struct {
	Bin
}

// native returns a pointer to the underlying GtkAlignment.
func (v *Alignment) native() *C.GtkAlignment {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkAlignment(p)
}

func marshalAlignment(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapAlignment(obj), nil
}

func wrapAlignment(obj *glib.Object) *Alignment {
	return &Alignment{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

/*
 * GtkStatusIcon
 * deprecated since version 3.14
 */

// StatusIcon is a representation of GTK's GtkStatusIcon
type StatusIcon struct {
	*glib.Object
}

func marshalStatusIcon(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapStatusIcon(obj), nil
}

func wrapStatusIcon(obj *glib.Object) *StatusIcon {
	return &StatusIcon{obj}
}

func (v *StatusIcon) native() *C.GtkStatusIcon {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkStatusIcon(p)
}

// StatusIconNew is a wrapper around gtk_status_icon_new()
func StatusIconNew() (*StatusIcon, error) {
	c := C.gtk_status_icon_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	e := wrapStatusIcon(obj)
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return e, nil
}

// StatusIconNewFromFile is a wrapper around gtk_status_icon_new_from_file()
func StatusIconNewFromFile(filename string) (*StatusIcon, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_status_icon_new_from_file((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	obj.RefSink()
	e := wrapStatusIcon(obj)
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return e, nil
}

// StatusIconNewFromIconName is a wrapper around gtk_status_icon_new_from_name()
func StatusIconNewFromIconName(iconName string) (*StatusIcon, error) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	s := C.gtk_status_icon_new_from_icon_name((*C.gchar)(cstr))
	if s == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(s))}
	obj.RefSink()
	e := wrapStatusIcon(obj)
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return e, nil
}

// SetFromFile is a wrapper around gtk_status_icon_set_from_file()
func (v *StatusIcon) SetFromFile(filename string) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_from_file(v.native(), (*C.gchar)(cstr))
}

// SetFromIconName is a wrapper around gtk_status_icon_set_from_icon_name()
func (v *StatusIcon) SetFromIconName(iconName string) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_from_icon_name(v.native(), (*C.gchar)(cstr))
}

// GetStorageType is a wrapper around gtk_status_icon_get_storage_type()
func (v *StatusIcon) GetStorageType() ImageType {
	return (ImageType)(C.gtk_status_icon_get_storage_type(v.native()))
}

// SetTooltipText is a wrapper around gtk_status_icon_set_tooltip_text()
func (v *StatusIcon) SetTooltipText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_tooltip_text(v.native(), (*C.gchar)(cstr))
}

// GetTooltipText is a wrapper around gtk_status_icon_get_tooltip_text()
func (v *StatusIcon) GetTooltipText() string {
	cstr := (*C.char)(C.gtk_status_icon_get_tooltip_text(v.native()))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// SetTooltipMarkup is a wrapper around gtk_status_icon_set_tooltip_markup()
func (v *StatusIcon) SetTooltipMarkup(markup string) {
	cstr := (*C.gchar)(C.CString(markup))
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_tooltip_markup(v.native(), cstr)
}

// GetTooltipMarkup is a wrapper around gtk_status_icon_get_tooltip_markup()
func (v *StatusIcon) GetTooltipMarkup() string {
	cstr := (*C.char)(C.gtk_status_icon_get_tooltip_markup(v.native()))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// SetHasTooltip is a wrapper around gtk_status_icon_set_has_tooltip()
func (v *StatusIcon) SetHasTooltip(hasTooltip bool) {
	C.gtk_status_icon_set_has_tooltip(v.native(), gbool(hasTooltip))
}

// GetTitle is a wrapper around gtk_status_icon_get_title()
func (v *StatusIcon) GetTitle() string {
	cstr := (*C.char)(C.gtk_status_icon_get_title(v.native()))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// SetName is a wrapper around gtk_status_icon_set_name()
func (v *StatusIcon) SetName(name string) {
	cstr := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_name(v.native(), cstr)
}

// SetVisible is a wrapper around gtk_status_icon_set_visible()
func (v *StatusIcon) SetVisible(visible bool) {
	C.gtk_status_icon_set_visible(v.native(), gbool(visible))
}

// GetVisible is a wrapper around gtk_status_icon_get_visible()
func (v *StatusIcon) GetVisible() bool {
	return gobool(C.gtk_status_icon_get_visible(v.native()))
}

// IsEmbedded is a wrapper around gtk_status_icon_is_embedded()
func (v *StatusIcon) IsEmbedded() bool {
	return gobool(C.gtk_status_icon_is_embedded(v.native()))
}

// GetX11WindowID is a wrapper around gtk_status_icon_get_x11_window_id()
func (v *StatusIcon) GetX11WindowID() int {
	return int(C.gtk_status_icon_get_x11_window_id(v.native()))
}

// GetHasTooltip is a wrapper around gtk_status_icon_get_has_tooltip()
func (v *StatusIcon) GetHasTooltip() bool {
	return gobool(C.gtk_status_icon_get_has_tooltip(v.native()))
}

// SetTitle is a wrapper around gtk_status_icon_set_title()
func (v *StatusIcon) SetTitle(title string) {
	cstr := (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_title(v.native(), cstr)
}

// GetIconName is a wrapper around gtk_status_icon_get_icon_name()
func (v *StatusIcon) GetIconName() string {
	cstr := (*C.char)(C.gtk_status_icon_get_icon_name(v.native()))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// GetSize is a wrapper around gtk_status_icon_get_size()
func (v *StatusIcon) GetSize() int {
	return int(C.gtk_status_icon_get_size(v.native()))
}

/*
 * End deprecated since version 3.14
 */

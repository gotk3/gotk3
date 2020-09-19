// Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
//
// Wrapper for GtkPopover originated from: http://opensource.conformal.com/
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

// +build !gtk_3_6,!gtk_3_8,!gtk_3_10
// not use this: go build -tags gtk_3_8'. Otherwise, if no build tags are used, GTK 3.10

package gtk

// #include <stdlib.h>
// #include <gtk/gtk.h>
// #include "gtk_since_3_12.go.h"
import "C"

import (
	"errors"
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
)

/*
 * Constants
 */

const (
	DIALOG_USE_HEADER_BAR DialogFlags = C.GTK_DIALOG_USE_HEADER_BAR
)

const (
	STATE_FLAG_LINK    StateFlags = C.GTK_STATE_FLAG_LINK
	STATE_FLAG_VISITED StateFlags = C.GTK_STATE_FLAG_VISITED
)

const (
	BUTTONBOX_EXPAND ButtonBoxStyle = C.GTK_BUTTONBOX_EXPAND
)

func init() {
	tm := []glib.TypeMarshaler{
		// Objects/Interfaces
		{glib.Type(C.gtk_flow_box_get_type()), marshalFlowBox},
		{glib.Type(C.gtk_flow_box_child_get_type()), marshalFlowBoxChild},
		{glib.Type(C.gtk_popover_get_type()), marshalPopover},
	}
	glib.RegisterGValueMarshalers(tm)

	//Contribute to casting
	for k, v := range map[string]WrapFn{
		"GtkFlowBox":      wrapFlowBox,
		"GtkFlowBoxChild": wrapFlowBoxChild,
		"GtkPopover":      wrapPopover,
	} {
		WrapMap[k] = v
	}
}

// GetLocaleDirection is a wrapper around gtk_get_locale_direction().
func GetLocaleDirection() TextDirection {
	c := C.gtk_get_locale_direction()
	return TextDirection(c)
}

/*
 * GtkStack
 */

const (
	STACK_TRANSITION_TYPE_OVER_UP      StackTransitionType = C.GTK_STACK_TRANSITION_TYPE_OVER_UP
	STACK_TRANSITION_TYPE_OVER_DOWN                        = C.GTK_STACK_TRANSITION_TYPE_OVER_DOWN
	STACK_TRANSITION_TYPE_OVER_LEFT                        = C.GTK_STACK_TRANSITION_TYPE_OVER_LEFT
	STACK_TRANSITION_TYPE_OVER_RIGHT                       = C.GTK_STACK_TRANSITION_TYPE_OVER_RIGHT
	STACK_TRANSITION_TYPE_UNDER_UP                         = C.GTK_STACK_TRANSITION_TYPE_UNDER_UP
	STACK_TRANSITION_TYPE_UNDER_DOWN                       = C.GTK_STACK_TRANSITION_TYPE_UNDER_DOWN
	STACK_TRANSITION_TYPE_UNDER_LEFT                       = C.GTK_STACK_TRANSITION_TYPE_UNDER_LEFT
	STACK_TRANSITION_TYPE_UNDER_RIGHT                      = C.GTK_STACK_TRANSITION_TYPE_UNDER_RIGHT
	STACK_TRANSITION_TYPE_OVER_UP_DOWN                     = C.GTK_STACK_TRANSITION_TYPE_OVER_UP_DOWN
)

/*
 * Dialog
 */

// GetHeaderBar is a wrapper around gtk_dialog_get_header_bar().
func (v *Dialog) GetHeaderBar() (IWidget, error) {
	c := C.gtk_dialog_get_header_bar(v.native())
	if c == nil {
		return nil, nil
	}
	return castWidget(c)
}

/*
 * Entry
 */

// SetMaxWidthChars is a wrapper around gtk_entry_set_max_width_chars().
func (v *Entry) SetMaxWidthChars(nChars int) {
	C.gtk_entry_set_max_width_chars(v.native(), C.gint(nChars))
}

// GetMaxWidthChars is a wrapper around gtk_entry_get_max_width_chars().
func (v *Entry) GetMaxWidthChars() int {
	c := C.gtk_entry_get_max_width_chars(v.native())
	return int(c)
}

/*
 * HeaderBar
 */

// GetDecorationLayout is a wrapper around gtk_header_bar_get_decoration_layout().
func (v *HeaderBar) GetDecorationLayout() string {
	c := C.gtk_header_bar_get_decoration_layout(v.native())
	return C.GoString((*C.char)(c))
}

// SetDecorationLayout is a wrapper around gtk_header_bar_set_decoration_layout().
func (v *HeaderBar) SetDecorationLayout(layout string) {
	cstr := C.CString(layout)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_header_bar_set_decoration_layout(v.native(), (*C.gchar)(cstr))
}

// GetHasSubtitle is a wrapper around gtk_header_bar_get_has_subtitle().
func (v *HeaderBar) GetHasSubtitle() bool {
	c := C.gtk_header_bar_get_has_subtitle(v.native())
	return gobool(c)
}

// SetHasSubtitle is a wrapper around gtk_header_bar_set_has_subtitle().
func (v *HeaderBar) SetHasSubtitle(setting bool) {
	C.gtk_header_bar_set_has_subtitle(v.native(), gbool(setting))
}

/*
 * MenuButton
 */

// SetPopover is a wrapper around gtk_menu_button_set_popover().
func (v *MenuButton) SetPopover(popover *Popover) {
	C.gtk_menu_button_set_popover(v.native(), popover.toWidget())
}

// GetPopover is a wrapper around gtk_menu_button_get_popover().
func (v *MenuButton) GetPopover() *Popover {
	c := C.gtk_menu_button_get_popover(v.native())
	if c == nil {
		return nil
	}
	return wrapPopover(glib.Take(unsafe.Pointer(c)))
}

// GetUsePopover is a wrapper around gtk_menu_button_get_use_popover().
func (v *MenuButton) GetUsePopover() bool {
	c := C.gtk_menu_button_get_use_popover(v.native())
	return gobool(c)
}

// SetUsePopover is a wrapper around gtk_menu_button_set_use_popover().
func (v *MenuButton) SetUsePopover(setting bool) {
	C.gtk_menu_button_set_use_popover(v.native(), gbool(setting))
}

/*
 * FlowBox
 */

// FlowBox is a representation of GtkFlowBox
type FlowBox struct {
	Container
}

func (fb *FlowBox) native() *C.GtkFlowBox {
	if fb == nil || fb.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(fb.GObject)
	return C.toGtkFlowBox(p)
}

func marshalFlowBox(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFlowBox(obj), nil
}

func wrapFlowBox(obj *glib.Object) *FlowBox {
	return &FlowBox{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// FlowBoxNew is a wrapper around gtk_flow_box_new()
func FlowBoxNew() (*FlowBox, error) {
	c := C.gtk_flow_box_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapFlowBox(glib.Take(unsafe.Pointer(c))), nil
}

// Insert is a wrapper around gtk_flow_box_insert()
func (fb *FlowBox) Insert(widget IWidget, position int) {
	C.gtk_flow_box_insert(fb.native(), widget.toWidget(), C.gint(position))
}

// GetChildAtIndex is a wrapper around gtk_flow_box_get_child_at_index()
func (fb *FlowBox) GetChildAtIndex(idx int) *FlowBoxChild {
	c := C.gtk_flow_box_get_child_at_index(fb.native(), C.gint(idx))
	if c == nil {
		return nil
	}
	return wrapFlowBoxChild(glib.Take(unsafe.Pointer(c)))
}

// TODO 3.22.6 gtk_flow_box_get_child_at_pos()

// SetHAdjustment is a wrapper around gtk_flow_box_set_hadjustment()
func (fb *FlowBox) SetHAdjustment(adjustment *Adjustment) {
	C.gtk_flow_box_set_hadjustment(fb.native(), adjustment.native())
}

// SetVAdjustment is a wrapper around gtk_flow_box_set_vadjustment()
func (fb *FlowBox) SetVAdjustment(adjustment *Adjustment) {
	C.gtk_flow_box_set_vadjustment(fb.native(), adjustment.native())
}

// SetHomogeneous is a wrapper around gtk_flow_box_set_homogeneous()
func (fb *FlowBox) SetHomogeneous(homogeneous bool) {
	C.gtk_flow_box_set_homogeneous(fb.native(), gbool(homogeneous))
}

// GetHomogeneous is a wrapper around gtk_flow_box_get_homogeneous()
func (fb *FlowBox) GetHomogeneous() bool {
	c := C.gtk_flow_box_get_homogeneous(fb.native())
	return gobool(c)
}

// SetRowSpacing is a wrapper around gtk_flow_box_set_row_spacing()
func (fb *FlowBox) SetRowSpacing(spacing uint) {
	C.gtk_flow_box_set_row_spacing(fb.native(), C.guint(spacing))
}

// GetRowSpacing is a wrapper around gtk_flow_box_get_row_spacing()
func (fb *FlowBox) GetRowSpacing() uint {
	c := C.gtk_flow_box_get_row_spacing(fb.native())
	return uint(c)
}

// SetColumnSpacing is a wrapper around gtk_flow_box_set_column_spacing()
func (fb *FlowBox) SetColumnSpacing(spacing uint) {
	C.gtk_flow_box_set_column_spacing(fb.native(), C.guint(spacing))
}

// GetColumnSpacing is a wrapper around gtk_flow_box_get_column_spacing()
func (fb *FlowBox) GetColumnSpacing() uint {
	c := C.gtk_flow_box_get_column_spacing(fb.native())
	return uint(c)
}

// SetMinChildrenPerLine is a wrapper around gtk_flow_box_set_min_children_per_line()
func (fb *FlowBox) SetMinChildrenPerLine(n_children uint) {
	C.gtk_flow_box_set_min_children_per_line(fb.native(), C.guint(n_children))
}

// GetMinChildrenPerLine is a wrapper around gtk_flow_box_get_min_children_per_line()
func (fb *FlowBox) GetMinChildrenPerLine() uint {
	c := C.gtk_flow_box_get_min_children_per_line(fb.native())
	return uint(c)
}

// SetMaxChildrenPerLine is a wrapper around gtk_flow_box_set_max_children_per_line()
func (fb *FlowBox) SetMaxChildrenPerLine(n_children uint) {
	C.gtk_flow_box_set_max_children_per_line(fb.native(), C.guint(n_children))
}

// GetMaxChildrenPerLine is a wrapper around gtk_flow_box_get_max_children_per_line()
func (fb *FlowBox) GetMaxChildrenPerLine() uint {
	c := C.gtk_flow_box_get_max_children_per_line(fb.native())
	return uint(c)
}

// SetActivateOnSingleClick is a wrapper around gtk_flow_box_set_activate_on_single_click()
func (fb *FlowBox) SetActivateOnSingleClick(single bool) {
	C.gtk_flow_box_set_activate_on_single_click(fb.native(), gbool(single))
}

// GetActivateOnSingleClick gtk_flow_box_get_activate_on_single_click()
func (fb *FlowBox) GetActivateOnSingleClick() bool {
	c := C.gtk_flow_box_get_activate_on_single_click(fb.native())
	return gobool(c)
}

// TODO: gtk_flow_box_selected_foreach()

// GetSelectedChildren is a wrapper around gtk_flow_box_get_selected_children()
func (fb *FlowBox) GetSelectedChildren() (rv []*FlowBoxChild) {
	c := C.gtk_flow_box_get_selected_children(fb.native())
	if c == nil {
		return
	}
	list := glib.WrapList(uintptr(unsafe.Pointer(c)))
	for l := list; l != nil; l = l.Next() {
		o := wrapFlowBoxChild(glib.Take(l.Data().(unsafe.Pointer)))
		rv = append(rv, o)
	}
	// We got a transfer container, so we must free the list.
	list.Free()

	return
}

// SelectChild is a wrapper around gtk_flow_box_select_child()
func (fb *FlowBox) SelectChild(child *FlowBoxChild) {
	C.gtk_flow_box_select_child(fb.native(), child.native())
}

// UnselectChild is a wrapper around gtk_flow_box_unselect_child()
func (fb *FlowBox) UnselectChild(child *FlowBoxChild) {
	C.gtk_flow_box_unselect_child(fb.native(), child.native())
}

// SelectAll is a wrapper around gtk_flow_box_select_all()
func (fb *FlowBox) SelectAll() {
	C.gtk_flow_box_select_all(fb.native())
}

// UnselectAll is a wrapper around gtk_flow_box_unselect_all()
func (fb *FlowBox) UnselectAll() {
	C.gtk_flow_box_unselect_all(fb.native())
}

// SetSelectionMode is a wrapper around gtk_flow_box_set_selection_mode()
func (fb *FlowBox) SetSelectionMode(mode SelectionMode) {
	C.gtk_flow_box_set_selection_mode(fb.native(), C.GtkSelectionMode(mode))
}

// GetSelectionMode is a wrapper around gtk_flow_box_get_selection_mode()
func (fb *FlowBox) GetSelectionMode() SelectionMode {
	c := C.gtk_flow_box_get_selection_mode(fb.native())
	return SelectionMode(c)
}

// TODO gtk_flow_box_set_filter_func()
// TODO gtk_flow_box_invalidate_filter()
// TODO gtk_flow_box_set_sort_func()
// TODO gtk_flow_box_invalidate_sort()
// TODO 3.18 gtk_flow_box_bind_model()

/*
 * FlowBoxChild
 */

// FlowBoxChild is a representation of GtkFlowBoxChild
type FlowBoxChild struct {
	Bin
}

func (fbc *FlowBoxChild) native() *C.GtkFlowBoxChild {
	if fbc == nil || fbc.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(fbc.GObject)
	return C.toGtkFlowBoxChild(p)
}

func marshalFlowBoxChild(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFlowBoxChild(obj), nil
}

func wrapFlowBoxChild(obj *glib.Object) *FlowBoxChild {
	return &FlowBoxChild{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// FlowBoxChildNew is a wrapper around gtk_flow_box_child_new()
func FlowBoxChildNew() (*FlowBoxChild, error) {
	c := C.gtk_flow_box_child_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapFlowBoxChild(glib.Take(unsafe.Pointer(c))), nil
}

// GetIndex is a wrapper around gtk_flow_box_child_get_index()
func (fbc *FlowBoxChild) GetIndex() int {
	c := C.gtk_flow_box_child_get_index(fbc.native())
	return int(c)
}

// IsSelected is a wrapper around gtk_flow_box_child_is_selected()
func (fbc *FlowBoxChild) IsSelected() bool {
	c := C.gtk_flow_box_child_is_selected(fbc.native())
	return gobool(c)
}

// Changed is a wrapper around gtk_flow_box_child_changed()
func (fbc *FlowBoxChild) Changed() {
	C.gtk_flow_box_child_changed(fbc.native())
}

/*
 * GtkPlacesSidebar
 */

// TODO:
// gtk_places_sidebar_get_local_only().
// gtk_places_sidebar_set_local_only().

/*
 * GtkPopover
 */

// Popover is a representation of GTK's GtkPopover.
type Popover struct {
	Bin
}

func (v *Popover) native() *C.GtkPopover {
	if v == nil || v.GObject == nil {
		return nil
	}

	p := unsafe.Pointer(v.GObject)
	return C.toGtkPopover(p)
}

func marshalPopover(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapPopover(glib.Take(unsafe.Pointer(c))), nil
}

func wrapPopover(obj *glib.Object) *Popover {
	return &Popover{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// PopoverNew is a wrapper around gtk_popover_new().
func PopoverNew(relative IWidget) (*Popover, error) {
	//Takes relative to widget
	var c *C.struct__GtkWidget
	if relative == nil {
		c = C.gtk_popover_new(nil)
	} else {
		c = C.gtk_popover_new(relative.toWidget())
	}
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapPopover(glib.Take(unsafe.Pointer(c))), nil
}

// PopoverNewFromModel is a wrapper around gtk_popover_new_from_model().
func PopoverNewFromModel(relative IWidget, model *glib.MenuModel) (*Popover, error) {
	//Takes relative to widget
	var c *C.struct__GtkWidget

	mptr := C.toGMenuModel(unsafe.Pointer(model.Native()))

	if relative == nil {
		c = C.gtk_popover_new_from_model(nil, mptr)
	} else {
		c = C.gtk_popover_new_from_model(relative.toWidget(), mptr)
	}
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapPopover(glib.Take(unsafe.Pointer(c))), nil
}

// BindModel is a wrapper around gtk_popover_bind_model().
func (v *Popover) BindModel(menuModel *glib.MenuModel, actionNamespace string) {
	cstr := C.CString(actionNamespace)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_popover_bind_model(v.native(), C.toGMenuModel(unsafe.Pointer(menuModel.Native())), (*C.gchar)(cstr))
}

// SetRelativeTo is a wrapper around gtk_popover_set_relative_to().
func (v *Popover) SetRelativeTo(relative IWidget) {
	C.gtk_popover_set_relative_to(v.native(), relative.toWidget())
}

// GetRelativeTo is a wrapper around gtk_popover_get_relative_to().
func (v *Popover) GetRelativeTo() (IWidget, error) {
	c := C.gtk_popover_get_relative_to(v.native())
	if c == nil {
		return nil, nil
	}
	return castWidget(c)
}

// SetPointingTo is a wrapper around gtk_popover_set_pointing_to().
func (v *Popover) SetPointingTo(rect gdk.Rectangle) {
	C.gtk_popover_set_pointing_to(v.native(), nativeGdkRectangle(rect))
}

// GetPointingTo is a wrapper around gtk_popover_get_pointing_to().
func (v *Popover) GetPointingTo() (*gdk.Rectangle, bool) {
	var cRect *C.GdkRectangle
	isSet := C.gtk_popover_get_pointing_to(v.native(), cRect)
	rect := gdk.WrapRectangle(uintptr(unsafe.Pointer(cRect)))
	return rect, gobool(isSet)
}

// SetPosition is a wrapper around gtk_popover_set_position().
func (v *Popover) SetPosition(position PositionType) {
	C.gtk_popover_set_position(v.native(), C.GtkPositionType(position))
}

// GetPosition is a wrapper around gtk_popover_get_position().
func (v *Popover) GetPosition() PositionType {
	c := C.gtk_popover_get_position(v.native())
	return PositionType(c)
}

// SetModal is a wrapper around gtk_popover_set_modal().
func (v *Popover) SetModal(modal bool) {
	C.gtk_popover_set_modal(v.native(), gbool(modal))
}

// GetModal is a wrapper around gtk_popover_get_modal().
func (v *Popover) GetModal() bool {
	return gobool(C.gtk_popover_get_modal(v.native()))
}

/*
 * TreePath
 */

// TreePathNewFromIndicesv is a wrapper around gtk_tree_path_new_from_indicesv().
func TreePathNewFromIndicesv(indices []int) (*TreePath, error) {
	if len(indices) == 0 {
		return nil, errors.New("no indice")
	}

	var cIndices []C.gint
	for i := 0; i < len(indices); i++ {
		cIndices = append(cIndices, C.gint(indices[i]))
	}

	var cIndicesPointer *C.gint = &cIndices[0]
	c := C.gtk_tree_path_new_from_indicesv(cIndicesPointer, C.gsize(len(indices)))
	if c == nil {
		return nil, nilPtrErr
	}
	t := &TreePath{c}
	runtime.SetFinalizer(t, (*TreePath).free)
	return t, nil
}

// Same copyright and license as the rest of the files in this project

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
)

/*
 * GtkCellArea
 */

// TODO: macro
// GTK_CELL_AREA_WARN_INVALID_CELL_PROPERTY_ID(object, property_id, pspec)
// object - the GObject on which set_cell_property() or get_cell_property() was called
// property_id - the numeric id of the property
// pspec - the GParamSpec of the property
// C.GTK_CELL_AREA_WARN_INVALID_CELL_PROPERTY_ID

// CellArea is a representation of GTK's GtkCellArea.
type CellArea struct {
	glib.InitiallyUnowned
}

type ICellArea interface {
	toCellArea() *C.GtkCellArea
	ToCellArea() *C.GtkCellArea
}

// native returns a pointer to the underlying GtkCellArea.
func (v *CellArea) native() *C.GtkCellArea {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCellArea(p)
}

func (v *CellArea) toCellArea() *C.GtkCellArea {
	if v == nil {
		return nil
	}
	return v.native()
}

// ToCellArea is a helper getter, in case you use the interface gtk.ICellArea in your program.
// It returns e.g. *gtk.CellAreaBox as a *gtk.CellArea.
func (v *CellArea) ToCellArea() *CellArea {
	return v
}

func marshalCellArea(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCellArea(obj), nil
}

func wrapCellArea(obj *glib.Object) *CellArea {
	return &CellArea{glib.InitiallyUnowned{obj}}
}

// Add is a wrapper around gtk_cell_area_add().
func (v *CellArea) Add(renderer *CellRenderer) {
	C.gtk_cell_area_add(v.native(), renderer.native())
}

// Remove is a wrapper around gtk_cell_area_remove().
func (v *CellArea) Remove(renderer *CellRenderer) {
	C.gtk_cell_area_remove(v.native(), renderer.native())
}

// HasRenderer is a wrapper around gtk_cell_area_has_renderer().
func (v *CellArea) HasRenderer(renderer *CellRenderer) bool {
	return gobool(C.gtk_cell_area_has_renderer(v.native(), renderer.native()))
}

// TODO:
// depends on GtkCellCallback
// Foreach is a wrapper around gtk_cell_area_foreach().
// func (v *CellArea) Foreach(cb CellCallback, callbackData interface{}) {
// }

// TODO:
// depends on GtkCellAllocCallback
// ForeachAlloc is a wrapper around gtk_cell_area_foreach_alloc().
// func (v *CellArea) ForeachAlloc(context *CellAreaContext, widget IWidget, cellArea, backgroundArea *gdk.Rectangle, cb CellAllocCallback, callbackData interface{}) {
// }

// AreaEvent is a wrapper around gtk_cell_area_event().
func (v *CellArea) AreaEvent(context *CellAreaContext, widget IWidget,
	event *gdk.Event, cellArea *gdk.Rectangle, flags CellRendererState) int {

	e := (*C.GdkEvent)(unsafe.Pointer(event.Native()))
	c := C.gtk_cell_area_event(v.native(), context.native(), widget.toWidget(),
		e, nativeGdkRectangle(*cellArea), C.GtkCellRendererState(flags))

	return int(c)
}

// Render is a wrapper around gtk_cell_area_render().
func (v *CellArea) Render(context *CellAreaContext, widget IWidget, cr *cairo.Context,
	backgroundArea, cellArea *gdk.Rectangle, flags CellRendererState, paintFocus bool) {

	cairoContext := (*C.cairo_t)(unsafe.Pointer(cr.Native()))

	C.gtk_cell_area_render(v.native(), context.native(), widget.toWidget(), cairoContext,
		nativeGdkRectangle(*backgroundArea), nativeGdkRectangle(*cellArea),
		C.GtkCellRendererState(flags), gbool(paintFocus))
}

// GetCellAllocation is a wrapper around gtk_cell_area_get_cell_allocation().
func (v *CellArea) GetCellAllocation(context *CellAreaContext, widget IWidget,
	renderer *CellRenderer, cellArea *gdk.Rectangle) *gdk.Rectangle {

	var cRect *C.GdkRectangle
	C.gtk_cell_area_get_cell_allocation(v.native(), context.native(), widget.toWidget(), renderer.native(),
		nativeGdkRectangle(*cellArea), cRect)
	allocation := gdk.WrapRectangle(uintptr(unsafe.Pointer(cRect)))
	return allocation

}

// GetCellAtPosition is a wrapper around gtk_cell_area_get_cell_at_position().
func (v *CellArea) GetCellAtPosition(context *CellAreaContext, widget IWidget,
	cellArea *gdk.Rectangle, x, y int) (*CellRenderer, *gdk.Rectangle) {

	var cRect *C.GdkRectangle

	renderer := C.gtk_cell_area_get_cell_at_position(v.native(), context.native(), widget.toWidget(),
		nativeGdkRectangle(*cellArea), C.gint(x), C.gint(y), cRect)

	var allocation *gdk.Rectangle

	if cRect != nil {
		allocation = gdk.WrapRectangle(uintptr(unsafe.Pointer(cRect)))
	}

	return wrapCellRenderer(glib.Take(unsafe.Pointer(renderer))), allocation
}

// CreateContext is a wrapper around gtk_cell_area_create_context().
func (v *CellArea) CreateContext(renderer *CellRenderer) (*CellAreaContext, error) {
	c := C.gtk_cell_area_create_context(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapCellAreaContext(glib.Take(unsafe.Pointer(c))), nil
}

// CopyContext is a wrapper around gtk_cell_area_copy_context().
func (v *CellArea) CopyContext(context *CellAreaContext) (*CellAreaContext, error) {
	c := C.gtk_cell_area_copy_context(v.native(), context.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapCellAreaContext(glib.Take(unsafe.Pointer(c))), nil
}

// TODO:
// depends on GtkSizeRequestMode
// gtk_cell_area_get_request_mode

// GetPreferredWidth is a wrapper around gtk_cell_area_get_preferred_width().
func (v *CellArea) GetPreferredWidth(context *CellAreaContext, widget IWidget) (int, int) {
	var minWidth C.gint
	var naturalWidth C.gint
	C.gtk_cell_area_get_preferred_width(v.native(), context.native(), widget.toWidget(),
		&minWidth, &naturalWidth)

	return int(minWidth), int(naturalWidth)
}

// GetPreferredHeightForWidth is a wrapper around gtk_cell_area_get_preferred_height_for_width().
func (v *CellArea) GetPreferredHeightForWidth(context *CellAreaContext, widget IWidget, width int) (int, int) {
	var minHeight C.gint
	var naturalHeight C.gint
	C.gtk_cell_area_get_preferred_height_for_width(v.native(), context.native(), widget.toWidget(),
		C.gint(width), &minHeight, &naturalHeight)

	return int(minHeight), int(naturalHeight)
}

// GetPreferredHeight is a wrapper around gtk_cell_area_get_preferred_height().
func (v *CellArea) GetPreferredHeight(context *CellAreaContext, widget IWidget) (int, int) {
	var minHeight C.gint
	var naturalHeight C.gint
	C.gtk_cell_area_get_preferred_height(v.native(), context.native(), widget.toWidget(),
		&minHeight, &naturalHeight)

	return int(minHeight), int(naturalHeight)
}

// GetPreferredWidthForHeight is a wrapper around gtk_cell_area_get_preferred_width_for_height().
func (v *CellArea) GetPreferredWidthForHeight(context *CellAreaContext, widget IWidget, height int) (int, int) {
	var minWidth C.gint
	var naturalWidth C.gint
	C.gtk_cell_area_get_preferred_width_for_height(v.native(), context.native(), widget.toWidget(),
		C.gint(height), &minWidth, &naturalWidth)

	return int(minWidth), int(naturalWidth)
}

// GetCurrentPathString is a wrapper around gtk_cell_area_get_current_path_string().
func (v *CellArea) GetCurrentPathString() string {
	c := C.gtk_cell_area_get_current_path_string(v.native())
	// This string belongs to the area and should not be freed.
	return goString(c)
}

// ApplyAttributes is a wrapper around gtk_cell_area_apply_attributes().
func (v *CellArea) ApplyAttributes(model ITreeModel, iter *TreeIter, isExpander, isExpanded bool) {
	C.gtk_cell_area_apply_attributes(v.native(), model.toTreeModel(), iter.native(),
		gbool(isExpander), gbool(isExpanded))
}

// AttributeConnect is a wrapper around gtk_cell_area_attribute_connect().
func (v *CellArea) AttributeConnect(renderer *CellRenderer, attribute string, column int) {
	cstr := C.CString(attribute)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_cell_area_attribute_connect(v.native(), renderer.native(), (*C.gchar)(cstr), C.gint(column))
}

// AttributeDisconnect is a wrapper around gtk_cell_area_attribute_disconnect().
func (v *CellArea) AttributeDisconnect(renderer *CellRenderer, attribute string) {
	cstr := C.CString(attribute)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_cell_area_attribute_disconnect(v.native(), renderer.native(), (*C.gchar)(cstr))
}

// TODO:
// gtk_cell_area_class_install_cell_property // depends on GParamSpec
// gtk_cell_area_class_find_cell_property // depends on GParamSpec
// gtk_cell_area_class_list_cell_properties // depends on GParamSpec
// gtk_cell_area_add_with_properties
// gtk_cell_area_cell_set
// gtk_cell_area_cell_get
// gtk_cell_area_cell_set_valist
// gtk_cell_area_cell_get_valist

// CellSetProperty is a wrapper around gtk_cell_area_cell_set_property().
func (v *CellArea) CellSetProperty(renderer *CellRenderer, propertyName string, value interface{}) error {
	gval, err := glib.GValue(value)
	if err != nil {
		return err
	}
	cstr := C.CString(propertyName)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_cell_area_cell_set_property(v.native(), renderer.native(), (*C.gchar)(cstr),
		(*C.GValue)(unsafe.Pointer(gval.Native())))
	return nil
}

// CellGetProperty is a wrapper around gtk_cell_area_cell_get_property().
func (v *CellArea) CellGetProperty(renderer *CellRenderer, propertyName string) (interface{}, error) {
	cstr := C.CString(propertyName)
	defer C.free(unsafe.Pointer(cstr))

	var gval C.GValue
	C.gtk_cell_area_cell_get_property(v.native(), renderer.native(), (*C.gchar)(cstr), &gval)
	value := glib.ValueFromNative(unsafe.Pointer(&gval))
	return value.GoValue()
}

// IsActivatable is a wrapper around gtk_cell_area_is_activatable().
func (v *CellArea) IsActivatable() bool {
	return gobool(C.gtk_cell_area_is_activatable(v.native()))
}

// Activate is a wrapper around gtk_cell_area_activate().
func (v *CellArea) Activate(context *CellAreaContext, widget IWidget,
	cellArea *gdk.Rectangle, flags CellRendererState, editOnly bool) {

	C.gtk_cell_area_activate(v.native(), context.native(), widget.toWidget(),
		nativeGdkRectangle(*cellArea), C.GtkCellRendererState(flags), gbool(editOnly))
}

// Focus is a wrapper around gtk_cell_area_focus().
func (v *CellArea) Focus(direction DirectionType) bool {
	return gobool(C.gtk_cell_area_focus(v.native(), C.GtkDirectionType(direction)))
}

// SetFocusCell is a wrapper around gtk_cell_area_set_focus_cell().
func (v *CellArea) SetFocusCell(renderer *CellRenderer) {
	C.gtk_cell_area_set_focus_cell(v.native(), renderer.native())
}

// GetFocusCell is a wrapper around gtk_cell_area_get_focus_cell().
func (v *CellArea) GetFocusCell() *CellRenderer {
	c := C.gtk_cell_area_get_focus_cell(v.native())
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCellRenderer(obj)
}

// AddFocusSibling is a wrapper around gtk_cell_area_add_focus_sibling().
func (v *CellArea) AddFocusSibling(renderer, sibling *CellRenderer) {
	C.gtk_cell_area_add_focus_sibling(v.native(), renderer.native(), sibling.native())
}

// RemoveFocusSibling is a wrapper around gtk_cell_area_remove_focus_sibling().
func (v *CellArea) RemoveFocusSibling(renderer, sibling *CellRenderer) {
	C.gtk_cell_area_remove_focus_sibling(v.native(), renderer.native(), sibling.native())
}

// IsFocusSibling is a wrapper around gtk_cell_area_is_focus_sibling().
func (v *CellArea) IsFocusSibling(renderer, sibling *CellRenderer) bool {
	return gobool(C.gtk_cell_area_is_focus_sibling(v.native(), renderer.native(), sibling.native()))
}

// GetFocusSiblings is a wrapper around gtk_cell_area_get_focus_siblings().
func (v *CellArea) GetFocusSiblings(renderer *CellRenderer) *glib.List {
	clist := C.gtk_cell_area_get_focus_siblings(v.native(), renderer.native())
	if clist == nil {
		return nil
	}

	// The returned list is internal and should not be freed.
	glist := glib.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		return wrapCellRenderer(glib.Take(ptr))
	})

	return glist
}

// GetFocusFromSibling is a wrapper around gtk_cell_area_get_focus_from_sibling().
func (v *CellArea) GetFocusFromSibling(renderer *CellRenderer) *CellRenderer {
	c := C.gtk_cell_area_get_focus_from_sibling(v.native(), renderer.native())
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCellRenderer(obj)
}

// GetEditedCell is a wrapper around gtk_cell_area_get_edited_cell().
func (v *CellArea) GetEditedCell() *CellRenderer {
	c := C.gtk_cell_area_get_edited_cell(v.native())
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCellRenderer(obj)
}

// TODO:
// gtk_cell_area_get_edit_widget // depends on GtkCellEditable

// ActivateCell is a wrapper around gtk_cell_area_activate_cell().
func (v *CellArea) ActivateCell(widget IWidget, renderer *CellRenderer,
	event *gdk.Event, cellArea *gdk.Rectangle, flags CellRendererState) bool {

	e := (*C.GdkEvent)(unsafe.Pointer(event.Native()))
	c := C.gtk_cell_area_activate_cell(v.native(), widget.toWidget(), renderer.native(),
		e, nativeGdkRectangle(*cellArea), C.GtkCellRendererState(flags))

	return gobool(c)
}

// StopEditing is a wrapper around gtk_cell_area_stop_editing().
func (v *CellArea) StopEditing(cancelled bool) {
	C.gtk_cell_area_stop_editing(v.native(), gbool(cancelled))
}

// InnerCellArea is a wrapper around gtk_cell_area_inner_cell_area().
func (v *CellArea) InnerCellArea(widget IWidget, cellArea *gdk.Rectangle) *gdk.Rectangle {
	var cRect *C.GdkRectangle
	C.gtk_cell_area_inner_cell_area(v.native(), widget.toWidget(), nativeGdkRectangle(*cellArea), cRect)
	innerArea := gdk.WrapRectangle(uintptr(unsafe.Pointer(cRect)))
	return innerArea
}

// RequestRenderer is a wrapper around gtk_cell_area_request_renderer().
func (v *CellArea) RequestRenderer(renderer *CellRenderer, orientation Orientation,
	widget IWidget, forSize int) (int, int) {

	var minSize C.gint
	var naturalSize C.gint

	C.gtk_cell_area_request_renderer(v.native(), renderer.native(), C.GtkOrientation(orientation),
		widget.toWidget(), C.gint(forSize), &minSize, &naturalSize)

	return int(minSize), int(naturalSize)
}

/*
 * GtkCellAreaContext
 */

type CellAreaContext struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkCellAreaContext.
func (v *CellAreaContext) native() *C.GtkCellAreaContext {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCellAreaContext(p)
}

func (v *CellAreaContext) toCellAreaContext() *C.GtkCellAreaContext {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalCellAreaContext(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCellAreaContext(obj), nil
}

func wrapCellAreaContext(obj *glib.Object) *CellAreaContext {
	return &CellAreaContext{obj}
}

// GetArea is a wrapper around gtk_cell_area_context_get_area().
func (v *CellAreaContext) GetArea() (*CellArea, error) {
	c := C.gtk_cell_area_context_get_area(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapCellArea(glib.Take(unsafe.Pointer(c))), nil
}

// Allocate is a wrapper around gtk_cell_area_context_allocate().
func (v *CellAreaContext) Allocate(width, height int) {
	C.gtk_cell_area_context_allocate(v.native(), C.gint(width), C.gint(height))
}

// Reset is a wrapper around gtk_cell_area_context_reset().
func (v *CellAreaContext) Reset(width, height int) {
	C.gtk_cell_area_context_reset(v.native())
}

// GetPreferredWidth is a wrapper around gtk_cell_area_context_get_preferred_width().
func (v *CellAreaContext) GetPreferredWidth() (int, int) {
	var minWidth C.gint
	var naturalWidth C.gint
	C.gtk_cell_area_context_get_preferred_width(v.native(), &minWidth, &naturalWidth)
	return int(minWidth), int(naturalWidth)
}

// GetPreferredHeight is a wrapper around gtk_cell_area_context_get_preferred_height().
func (v *CellAreaContext) GetPreferredHeight() (int, int) {
	var minHeight C.gint
	var naturalHeight C.gint
	C.gtk_cell_area_context_get_preferred_height(v.native(), &minHeight, &naturalHeight)
	return int(minHeight), int(naturalHeight)
}

// GetPreferredHeightForWidth is a wrapper around gtk_cell_area_context_get_preferred_height_for_width().
func (v *CellAreaContext) GetPreferredHeightForWidth(width int) (int, int) {
	var minHeight C.gint
	var naturalHeight C.gint
	C.gtk_cell_area_context_get_preferred_height_for_width(v.native(), C.gint(width), &minHeight, &naturalHeight)
	return int(minHeight), int(naturalHeight)
}

// GetPreferredWidthForHeight is a wrapper around gtk_cell_area_context_get_preferred_width_for_height().
func (v *CellAreaContext) GetPreferredWidthForHeight(height int) (int, int) {
	var minWidth C.gint
	var naturalWidth C.gint
	C.gtk_cell_area_context_get_preferred_width_for_height(v.native(), C.gint(height), &minWidth, &naturalWidth)
	return int(minWidth), int(naturalWidth)
}

// GetAllocation is a wrapper around gtk_cell_area_context_get_allocation().
func (v *CellAreaContext) GetAllocation() (int, int) {
	var height C.gint
	var width C.gint
	C.gtk_cell_area_context_get_allocation(v.native(), &height, &width)
	return int(height), int(width)
}

// PushPreferredWidth is a wrapper around gtk_cell_area_context_push_preferred_width().
func (v *CellAreaContext) PushPreferredWidth(minWidth, naturalWidth int) {
	C.gtk_cell_area_context_push_preferred_width(v.native(), C.gint(minWidth), C.gint(naturalWidth))
}

// PushPreferredHeight is a wrapper around gtk_cell_area_context_push_preferred_height().
func (v *CellAreaContext) PushPreferredHeight(minHeight, naturalHeight int) {
	C.gtk_cell_area_context_push_preferred_height(v.native(), C.gint(minHeight), C.gint(naturalHeight))
}

/*
 * GtkCellAreaBox
 */

// CellAreaBox is a representation of GTK's GtkCellAreaBox.
type CellAreaBox struct {
	CellArea
}

// native returns a pointer to the underlying GtkCellAreaBox.
func (v *CellAreaBox) native() *C.GtkCellAreaBox {
	if v == nil || v.CellArea.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.CellArea.GObject)
	return C.toGtkCellAreaBox(p)
}

func marshalCellAreaBox(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapCellAreaBox(obj), nil
}

func wrapCellAreaBox(obj *glib.Object) *CellAreaBox {
	return &CellAreaBox{CellArea{glib.InitiallyUnowned{obj}}}
}

// CellAreaBoxNew is a wrapper around gtk_cell_area_box_new().
func CellAreaBoxNew() (*CellAreaBox, error) {
	c := C.gtk_cell_area_box_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapCellAreaBox(glib.Take(unsafe.Pointer(c))), nil
}

// PackStart is a wrapper around gtk_cell_area_box_pack_start().
func (v *CellAreaBox) PackStart(renderer *CellRenderer, expand, align, fixed bool) {
	C.gtk_cell_area_box_pack_start(v.native(), renderer.native(), gbool(expand), gbool(align), gbool(fixed))
}

// PackEnd is a wrapper around gtk_cell_area_box_pack_end().
func (v *CellAreaBox) PackEnd(renderer *CellRenderer, expand, align, fixed bool) {
	C.gtk_cell_area_box_pack_end(v.native(), renderer.native(), gbool(expand), gbool(align), gbool(fixed))
}

// GetSpacing is a wrapper around gtk_cell_area_box_get_spacing().
func (v *CellAreaBox) GetSpacing(renderer *CellRenderer) int {
	return int(C.gtk_cell_area_box_get_spacing(v.native()))
}

// SetSpacing is a wrapper around gtk_cell_area_box_set_spacing().
func (v *CellAreaBox) SetSpacing(renderer *CellRenderer, spacing int) {
	C.gtk_cell_area_box_set_spacing(v.native(), C.gint(spacing))
}

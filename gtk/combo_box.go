package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_combo_box_get_type()), marshalComboBox},
		{glib.Type(C.gtk_combo_box_text_get_type()), marshalComboBoxText},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkComboBox"] = wrapComboBox
	WrapMap["GtkComboBoxText"] = wrapComboBoxText
}

/*
 * GtkComboBox
 */

// ComboBox is a representation of GTK's GtkComboBox.
type ComboBox struct {
	Bin

	// Interfaces
	CellLayout
	CellEditable
}

// native returns a pointer to the underlying GtkComboBox.
func (v *ComboBox) native() *C.GtkComboBox {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkComboBox(p)
}

func (v *ComboBox) toCellLayout() *C.GtkCellLayout {
	if v == nil {
		return nil
	}
	return C.toGtkCellLayout(unsafe.Pointer(v.GObject))
}

func marshalComboBox(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapComboBox(obj), nil
}

func wrapComboBox(obj *glib.Object) *ComboBox {
	cl := wrapCellLayout(obj)
	ce := wrapCellEditable(obj)
	return &ComboBox{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}, *cl, *ce}
}

// ComboBoxNew is a wrapper around gtk_combo_box_new().
func ComboBoxNew() (*ComboBox, error) {
	c := C.gtk_combo_box_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapComboBox(obj), nil
}

// ComboBoxNewWithEntry is a wrapper around gtk_combo_box_new_with_entry().
func ComboBoxNewWithEntry() (*ComboBox, error) {
	c := C.gtk_combo_box_new_with_entry()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapComboBox(obj), nil
}

// ComboBoxNewWithModel is a wrapper around gtk_combo_box_new_with_model().
func ComboBoxNewWithModel(model ITreeModel) (*ComboBox, error) {
	c := C.gtk_combo_box_new_with_model(model.toTreeModel())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapComboBox(obj), nil
}

// is a wrapper around gtk_combo_box_new_with_model_and_entry().
func ComboBoxNewWithModelAndEntry(model ITreeModel) (*ComboBox, error) {
	c := C.gtk_combo_box_new_with_model_and_entry(model.toTreeModel())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapComboBox(obj), nil
}

// GetWrapWidth is a wrapper around gtk_combo_box_get_wrap_width().
func (v *ComboBox) GetWrapWidth() int {
	c := C.gtk_combo_box_get_wrap_width(v.native())
	return int(c)
}

// SetWrapWidth is a wrapper around gtk_combo_box_set_wrap_width().
func (v *ComboBox) SetWrapWidth(wrapWidth int) {
	C.gtk_combo_box_set_wrap_width(v.native(), C.gint(wrapWidth))
}

// GetRowSpanColumn is a wrapper around gtk_combo_box_get_row_span_column().
func (v *ComboBox) GetRowSpanColumn() int {
	c := C.gtk_combo_box_get_row_span_column(v.native())
	return int(c)
}

// SetRowSpanColumn is a wrapper around gtk_combo_box_set_row_span_column().
func (v *ComboBox) SetRowSpanColumn(rowSpan int) {
	C.gtk_combo_box_set_row_span_column(v.native(), C.gint(rowSpan))
}

// GetColumnSpanColumn is a wrapper around gtk_combo_box_get_column_span_column().
func (v *ComboBox) GetColumnSpanColumn() int {
	c := C.gtk_combo_box_get_column_span_column(v.native())
	return int(c)
}

// SetColumnSpanColumn is a wrapper around gtk_combo_box_set_column_span_column().
func (v *ComboBox) SetColumnSpanColumn(wrapWidth int) {
	C.gtk_combo_box_set_column_span_column(v.native(), C.gint(wrapWidth))
}

// GetActive is a wrapper around gtk_combo_box_get_active().
func (v *ComboBox) GetActive() int {
	c := C.gtk_combo_box_get_active(v.native())
	return int(c)
}

// SetActive is a wrapper around gtk_combo_box_set_active().
func (v *ComboBox) SetActive(index int) {
	C.gtk_combo_box_set_active(v.native(), C.gint(index))
}

// GetActiveIter is a wrapper around gtk_combo_box_get_active_iter().
func (v *ComboBox) GetActiveIter() (*TreeIter, error) {
	var cIter C.GtkTreeIter
	c := C.gtk_combo_box_get_active_iter(v.native(), &cIter)
	if !gobool(c) {
		return nil, errors.New("unable to get active iter")
	}
	return &TreeIter{cIter}, nil
}

// SetActiveIter is a wrapper around gtk_combo_box_set_active_iter().
func (v *ComboBox) SetActiveIter(iter *TreeIter) {
	var cIter *C.GtkTreeIter
	if iter != nil {
		cIter = &iter.GtkTreeIter
	}
	C.gtk_combo_box_set_active_iter(v.native(), cIter)
}

// GetIDColumn is a wrapper around gtk_combo_box_get_id_column()
func (v *ComboBox) GetIDColumn() int {
	c := C.gtk_combo_box_get_id_column(v.native())
	return int(c)
}

// SetIDColumn is a wrapper around gtk_combo_box_set_id_column()
func (v *ComboBox) SetIDColumn(idColumn int) {
	C.gtk_combo_box_set_id_column(v.native(), C.gint(idColumn))
}

// GetActiveID is a wrapper around gtk_combo_box_get_active_id().
func (v *ComboBox) GetActiveID() string {
	c := C.gtk_combo_box_get_active_id(v.native())
	return C.GoString((*C.char)(c))
}

// SetActiveID is a wrapper around gtk_combo_box_set_active_id().
func (v *ComboBox) SetActiveID(id string) bool {
	cid := C.CString(id)
	defer C.free(unsafe.Pointer(cid))
	c := C.gtk_combo_box_set_active_id(v.native(), (*C.gchar)(cid))
	return gobool(c)
}

// GetModel is a wrapper around gtk_combo_box_get_model().
func (v *ComboBox) GetModel() (ITreeModel, error) {
	c := C.gtk_combo_box_get_model(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return castTreeModel(c)
}

// SetModel is a wrapper around gtk_combo_box_set_model().
func (v *ComboBox) SetModel(model ITreeModel) {
	var mptr *C.GtkTreeModel
	if model != nil {
		mptr = model.toTreeModel()
	}
	C.gtk_combo_box_set_model(v.native(), mptr)
}

// PopupForDevice is a wrapper around gtk_combo_box_popup_for_device()
func (v *ComboBox) PopupForDevice(device *gdk.Device) {
	var devicePtr *C.GdkDevice
	if device != nil {
		devicePtr = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	}
	C.gtk_combo_box_popup_for_device(v.native(), devicePtr)
}

// Popup is a wrapper around gtk_combo_box_popup().
func (v *ComboBox) Popup() {
	C.gtk_combo_box_popup(v.native())
}

// Popdown is a wrapper around gtk_combo_box_popdown().
func (v *ComboBox) Popdown() {
	C.gtk_combo_box_popdown(v.native())
}

// GetHasEntry is a wrapper around gtk_combo_box_get_has_entry().
func (v *ComboBox) GetHasEntry() bool {
	c := C.gtk_combo_box_get_has_entry(v.native())
	return gobool(c)
}

// SetEntryTextColumn is a wrapper around gtk_combo_box_set_entry_text_column().
func (v *ComboBox) SetEntryTextColumn(textColumn int) {
	C.gtk_combo_box_set_entry_text_column(v.native(), C.gint(textColumn))
}

// GetEntryTextColumn is a wrapper around gtk_combo_box_get_entry_text_column().
func (v *ComboBox) GetEntryTextColumn() int {
	c := C.gtk_combo_box_get_entry_text_column(v.native())
	return int(c)
}

// SetPopupFixedWidth is a wrapper around gtk_combo_box_set_popup_fixed_width
func (v *ComboBox) SetPopupFixedWidth(fixedWidth bool) {
	C.gtk_combo_box_set_popup_fixed_width(v.native(), gbool(fixedWidth))
}

// GetPopupFixedWidth is a wrapper around gtk_combo_box_get_popup_fixed_width
func (v *ComboBox) GetPopupFixedWidth() bool {
	c := C.gtk_combo_box_get_popup_fixed_width(v.native())
	return gobool(c)
}

// GetEntry is a convenience func to get the Entry within the ComboBox.
// If the Combobox does not contain an Entry, an error is returned.
func (v *ComboBox) GetEntry() (*Entry, error) {
	hasEntry := C.gtk_combo_box_get_has_entry(v.native())
	if hasEntry == C.FALSE {
		return nil, errors.New("combobox has no entry")
	}
	bin := &v.Bin
	widget, err := bin.GetChild()
	if err != nil {
		return nil, err
	}
	entry, ok := widget.(*Entry)
	if !ok {
		return nil, fmt.Errorf("expected child to be of type *gtk.Entry, got %T", widget)
	}
	return entry, nil
}

/*
 * GtkComboBoxText
 */

// ComboBoxText is a representation of GTK's GtkComboBoxText.
type ComboBoxText struct {
	ComboBox
}

// native returns a pointer to the underlying GtkComboBoxText.
func (v *ComboBoxText) native() *C.GtkComboBoxText {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkComboBoxText(p)
}

func marshalComboBoxText(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapComboBoxText(obj), nil
}

func wrapComboBoxText(obj *glib.Object) *ComboBoxText {
	return &ComboBoxText{*wrapComboBox(obj)}
}

// ComboBoxTextNew is a wrapper around gtk_combo_box_text_new().
func ComboBoxTextNew() (*ComboBoxText, error) {
	c := C.gtk_combo_box_text_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapComboBoxText(obj), nil
}

// ComboBoxTextNewWithEntry is a wrapper around gtk_combo_box_text_new_with_entry().
func ComboBoxTextNewWithEntry() (*ComboBoxText, error) {
	c := C.gtk_combo_box_text_new_with_entry()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapComboBoxText(obj), nil
}

// Append is a wrapper around gtk_combo_box_text_append().
func (v *ComboBoxText) Append(id, text string) {
	cid := C.CString(id)
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(cid))
	defer C.free(unsafe.Pointer(ctext))
	C.gtk_combo_box_text_append(v.native(), (*C.gchar)(cid), (*C.gchar)(ctext))
}

// Prepend is a wrapper around gtk_combo_box_text_prepend().
func (v *ComboBoxText) Prepend(id, text string) {
	cid := C.CString(id)
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(cid))
	defer C.free(unsafe.Pointer(ctext))
	C.gtk_combo_box_text_prepend(v.native(), (*C.gchar)(cid), (*C.gchar)(ctext))
}

// Insert is a wrapper around gtk_combo_box_text_insert().
func (v *ComboBoxText) Insert(position int, id, text string) {
	cid := C.CString(id)
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(cid))
	defer C.free(unsafe.Pointer(ctext))
	C.gtk_combo_box_text_insert(v.native(), C.gint(position), (*C.gchar)(cid), (*C.gchar)(ctext))
}

// AppendText is a wrapper around gtk_combo_box_text_append_text().
func (v *ComboBoxText) AppendText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_combo_box_text_append_text(v.native(), (*C.gchar)(cstr))
}

// PrependText is a wrapper around gtk_combo_box_text_prepend_text().
func (v *ComboBoxText) PrependText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_combo_box_text_prepend_text(v.native(), (*C.gchar)(cstr))
}

// InsertText is a wrapper around gtk_combo_box_text_insert_text().
func (v *ComboBoxText) InsertText(position int, text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_combo_box_text_insert_text(v.native(), C.gint(position), (*C.gchar)(cstr))
}

// Remove is a wrapper around gtk_combo_box_text_remove().
func (v *ComboBoxText) Remove(position int) {
	C.gtk_combo_box_text_remove(v.native(), C.gint(position))
}

// RemoveAll is a wrapper around gtk_combo_box_text_remove_all().
func (v *ComboBoxText) RemoveAll() {
	C.gtk_combo_box_text_remove_all(v.native())
}

// GetActiveText is a wrapper around gtk_combo_box_text_get_active_text().
func (v *ComboBoxText) GetActiveText() string {
	c := (*C.char)(C.gtk_combo_box_text_get_active_text(v.native()))
	defer C.free(unsafe.Pointer(c))
	return C.GoString(c)
}

// gtk_combo_box_new_with_area, requires GtkCellArea
// gtk_combo_box_new_with_area_and_entry, requires GtkCellArea
// gtk_combo_box_get_row_separator_func, requires GtkTreeViewRowSeparatorFunc
// gtk_combo_box_set_row_separator_func, requires GtkTreeViewRowSeparatorFunc
// gtk_combo_box_get_popup_accessible, requires AtkObject
// gtk_combo_box_set_add_tearoffs, deprecated since 3.10
// gtk_combo_box_get_add_tearoffs, deprecated since 3.10
// gtk_combo_box_set_title, deprecated since 3.10
// gtk_combo_box_get_title, deprecated since 3.10
// gtk_combo_box_set_focus_on_click, deprecated since 3.20
// gtk_combo_box_get_focus_on_click, deprecated since 3.20
// gtk_combo_box_set_button_sensitivity, requires GtkSensitivityType
// gtk_combo_box_get_button_sensitivity, requires GtkSensitivityType

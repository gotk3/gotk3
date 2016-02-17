package impl

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"errors"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	glib_impl "github.com/gotk3/gotk3/glib/impl"
	"github.com/gotk3/gotk3/gtk"
)

func init() {
	tm := []glib_impl.TypeMarshaler{
		{glib.Type(C.gtk_combo_box_get_type()), marshalComboBox},
		{glib.Type(C.gtk_combo_box_text_get_type()), marshalComboBoxText},
	}

	glib_impl.RegisterGValueMarshalers(tm)

	WrapMap["GtkComboBox"] = wrapComboBox
	WrapMap["GtkComboBoxText"] = wrapComboBoxText
}

/*
 * GtkComboBox
 */

// ComboBox is a representation of GTK's GtkComboBox.
type comboBox struct {
	bin

	// Interfaces
	cellLayout
}

// native returns a pointer to the underlying GtkComboBox.
func (v *comboBox) native() *C.GtkComboBox {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkComboBox(p)
}

func (v *comboBox) toCellLayout() *C.GtkCellLayout {
	if v == nil {
		return nil
	}
	return C.toGtkCellLayout(unsafe.Pointer(v.GObject))
}

func marshalComboBox(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapComboBox(obj), nil
}

func wrapComboBox(obj *glib_impl.Object) *comboBox {
	cl := wrapCellLayout(obj)
	return &comboBox{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}, *cl}
}

// ComboBoxNew() is a wrapper around gtk_combo_box_new().
func ComboBoxNew() (*comboBox, error) {
	c := C.gtk_combo_box_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapComboBox(obj), nil
}

// ComboBoxNewWithEntry() is a wrapper around gtk_combo_box_new_with_entry().
func ComboBoxNewWithEntry() (*comboBox, error) {
	c := C.gtk_combo_box_new_with_entry()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapComboBox(obj), nil
}

// ComboBoxNewWithModel() is a wrapper around gtk_combo_box_new_with_model().
func ComboBoxNewWithModel(model ITreeModel) (*comboBox, error) {
	c := C.gtk_combo_box_new_with_model(model.toTreeModel())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapComboBox(obj), nil
}

// GetActive() is a wrapper around gtk_combo_box_get_active().
func (v *comboBox) GetActive() int {
	c := C.gtk_combo_box_get_active(v.native())
	return int(c)
}

// SetActive() is a wrapper around gtk_combo_box_set_active().
func (v *comboBox) SetActive(index int) {
	C.gtk_combo_box_set_active(v.native(), C.gint(index))
}

// GetActiveIter is a wrapper around gtk_combo_box_get_active_iter().
func (v *comboBox) GetActiveIter() (gtk.TreeIter, error) {
	var cIter C.GtkTreeIter
	c := C.gtk_combo_box_get_active_iter(v.native(), &cIter)
	if !gobool(c) {
		return nil, errors.New("unable to get active iter")
	}
	return &treeIter{cIter}, nil
}

// SetActiveIter is a wrapper around gtk_combo_box_set_active_iter().
func (v *comboBox) SetActiveIter(iter gtk.TreeIter) {
	var cIter *C.GtkTreeIter
	if iter != nil {
		cIter = &castToTreeIter(iter).GtkTreeIter
	}
	C.gtk_combo_box_set_active_iter(v.native(), cIter)
}

// GetActiveID is a wrapper around gtk_combo_box_get_active_id().
func (v *comboBox) GetActiveID() string {
	c := C.gtk_combo_box_get_active_id(v.native())
	return C.GoString((*C.char)(c))
}

// SetActiveID is a wrapper around gtk_combo_box_set_active_id().
func (v *comboBox) SetActiveID(id string) bool {
	cid := C.CString(id)
	defer C.free(unsafe.Pointer(cid))
	c := C.gtk_combo_box_set_active_id(v.native(), (*C.gchar)(cid))
	return gobool(c)
}

// GetModel is a wrapper around gtk_combo_box_get_model().
func (v *comboBox) GetModel() (gtk.TreeModel, error) {
	c := C.gtk_combo_box_get_model(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapTreeModel(obj), nil
}

// SetModel is a wrapper around gtk_combo_box_set_model().
func (v *comboBox) SetModel(model gtk.TreeModel) {
	C.gtk_combo_box_set_model(v.native(), model.(ITreeModel).toTreeModel())
}

/*
 * GtkComboBoxText
 */

// ComboBoxText is a representation of GTK's GtkComboBoxText.
type comboBoxText struct {
	comboBox
}

// native returns a pointer to the underlying GtkComboBoxText.
func (v *comboBoxText) native() *C.GtkComboBoxText {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkComboBoxText(p)
}

func marshalComboBoxText(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapComboBoxText(obj), nil
}

func wrapComboBoxText(obj *glib_impl.Object) *comboBoxText {
	return &comboBoxText{*wrapComboBox(obj)}
}

// ComboBoxTextNew is a wrapper around gtk_combo_box_text_new().
func ComboBoxTextNew() (*comboBoxText, error) {
	c := C.gtk_combo_box_text_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapComboBoxText(obj), nil
}

// ComboBoxTextNewWithEntry is a wrapper around gtk_combo_box_text_new_with_entry().
func ComboBoxTextNewWithEntry() (*comboBoxText, error) {
	c := C.gtk_combo_box_text_new_with_entry()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapComboBoxText(obj), nil
}

// Append is a wrapper around gtk_combo_box_text_append().
func (v *comboBoxText) Append(id, text string) {
	cid := C.CString(id)
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(cid))
	defer C.free(unsafe.Pointer(ctext))
	C.gtk_combo_box_text_append(v.native(), (*C.gchar)(cid), (*C.gchar)(ctext))
}

// Prepend is a wrapper around gtk_combo_box_text_prepend().
func (v *comboBoxText) Prepend(id, text string) {
	cid := C.CString(id)
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(cid))
	defer C.free(unsafe.Pointer(ctext))
	C.gtk_combo_box_text_prepend(v.native(), (*C.gchar)(cid), (*C.gchar)(ctext))
}

// Insert is a wrapper around gtk_combo_box_text_insert().
func (v *comboBoxText) Insert(position int, id, text string) {
	cid := C.CString(id)
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(cid))
	defer C.free(unsafe.Pointer(ctext))
	C.gtk_combo_box_text_insert(v.native(), C.gint(position), (*C.gchar)(cid), (*C.gchar)(ctext))
}

// AppendText is a wrapper around gtk_combo_box_text_append_text().
func (v *comboBoxText) AppendText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_combo_box_text_append_text(v.native(), (*C.gchar)(cstr))
}

// PrependText is a wrapper around gtk_combo_box_text_prepend_text().
func (v *comboBoxText) PrependText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_combo_box_text_prepend_text(v.native(), (*C.gchar)(cstr))
}

// InsertText is a wrapper around gtk_combo_box_text_insert_text().
func (v *comboBoxText) InsertText(position int, text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_combo_box_text_insert_text(v.native(), C.gint(position), (*C.gchar)(cstr))
}

// Remove is a wrapper around gtk_combo_box_text_remove().
func (v *comboBoxText) Remove2(position int) {
	C.gtk_combo_box_text_remove(v.native(), C.gint(position))
}

// RemoveAll is a wrapper around gtk_combo_box_text_remove_all().
func (v *comboBoxText) RemoveAll() {
	C.gtk_combo_box_text_remove_all(v.native())
}

// GetActiveText is a wrapper around gtk_combo_box_text_get_active_text().
func (v *comboBoxText) GetActiveText() string {
	c := (*C.char)(C.gtk_combo_box_text_get_active_text(v.native()))
	defer C.free(unsafe.Pointer(c))
	return C.GoString(c)
}

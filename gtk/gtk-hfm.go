

package gtk

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"errors"
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
)

// HasSelection() is a wrapper around gtk_text_buffer_get_has_selection().
// 1st methode create an issue: "panic: runtime error: cgo argument has Go pointer to Go pointer ..."
// func (v *TextBuffer) HasSelectionN() bool {
// 	return gobool(C.gtk_text_buffer_get_has_selection(v.native()))
// }

// HasSelection() is a variant solution around gtk_text_buffer_get_has_selection().
// So I use property methode to do the job it's more stable.
func (v *TextBuffer) HasSelection() bool {
	value, _ := v.GetProperty("has-selection")
	return value.(bool)
}

// DeleteSelection() is a wrapper around gtk_text_buffer_delete_selection().
func (v *TextBuffer) DeleteSelection(interactive, defaultEditable bool) bool {
	return gobool(C.gtk_text_buffer_delete_selection(v.native(), gbool(interactive), gbool(defaultEditable)))
}

// GetSelectionBound() is a wrapper around gtk_text_buffer_get_selection_bound().
func (v *TextBuffer) GetSelectionBound() *TextMark {
	ret := C.gtk_text_buffer_get_selection_bound(v.native())
	return (*TextMark)(ret)
}

// GetSelectionBounds() is a wrapper around gtk_text_buffer_get_selection_bounds().
func (v *TextBuffer) GetSelectionBounds() (start, end *TextIter) {
	start, end = new(TextIter), new(TextIter)
	C.gtk_text_buffer_get_selection_bounds(v.native(), (*C.GtkTextIter)(start), (*C.GtkTextIter)(end))
	return
}

// GetIterAtLineOffset() is a wrapper around gtk_text_buffer_get_iter_at_line_offset().
func (v *TextBuffer) GetIterAtLineOffset(lineNumber, charOffset int) (iter *TextIter) {
	iter = new(TextIter)
	C.gtk_text_buffer_get_iter_at_line_offset(v.native(), (*C.GtkTextIter)(iter), (C.gint)(lineNumber), (C.gint)(charOffset))
	return
}

// CreateTag() is a variant solution around gtk_text_buffer_create_tag().
func (v *TextBuffer) CreateTag(name string, props map[string]interface{}) (tag *TextTag, err error) {
	if tag, err = TextTagNew(name); err == nil {
		if tagTable, err := v.GetTagTable(); err == nil {
			tagTable.Add(tag)
			for n, p := range props {
				err = tag.SetProperty(n, p)
			}
		}
	}
	return
}

// // CreateTagN() is a wrapper around gtk_text_buffer_create_tag(). I got same error as HasSelection()
// func (v *TextBuffer) CreateTagN(name string, props map[string]interface{}) (tag *TextTag) {
// 	cname = C.CString(name)
// 	defer C.free(unsafe.Pointer(cname))
// 	cstr := C.CString("")
// 	defer C.g_free(C.gpointer(unsafe.Pointer(cstr)))
// 	c := C.gtk_text_buffer_create_tag(v.native(), (*C.gchar)(cname), unsafe.Pointer(&cstr))
// 	if tag = wrapTextTag(glib.Take(unsafe.Pointer(c))); tag != nil {
// 		for n, p := range props { // Useless to handle error (my opinion ...)
// 			tag.SetProperty(n, p)
// 		}
// 	}
// 	return
// }

// RemoveTagByName() is a wrapper around  gtk_text_buffer_remove_tag_by_name()
func (v *TextBuffer) RemoveTagByName(name string, start, end *TextIter) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_remove_tag_by_name(v.native(), (*C.gchar)(cstr), (*C.GtkTextIter)(start), (*C.GtkTextIter)(end))
}

// InsertMarkup() is a wrapper around  gtk_text_buffer_insert_markup()
func (v *TextBuffer) InsertMarkup(start *TextIter, text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_insert_markup(v.native(), (*C.GtkTextIter)(start), (*C.gchar)(cstr), C.gint(len(text)))
}

// InsertMarkup() is a wrapper around  gtk_text_buffer_register_serialize_tagset()
func (v *TextBuffer) RegisterSerializeTagset(tagsetName string) gdk.Atom {
	cstr := C.CString(tagsetName)
	if len(tagsetName) == 0 {
		cstr = nil
	}
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_text_buffer_register_serialize_tagset(v.native(), (*C.gchar)(cstr))
	return gdk.Atom(uintptr(unsafe.Pointer(c)))
}

// RegisterDeserializeTagset() is a wrapper around  gtk_text_buffer_register_deserialize_tagset()
func (v *TextBuffer) RegisterDeserializeTagset(tagsetName string) gdk.Atom {
	cstr := C.CString(tagsetName)
	if len(tagsetName) == 0 {
		cstr = nil
	}
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_text_buffer_register_deserialize_tagset(v.native(), (*C.gchar)(cstr))
	return gdk.Atom(uintptr(unsafe.Pointer(c)))
}

// Serialize() is a wrapper around  gtk_text_buffer_serialize()
func (v *TextBuffer) Serialize(contentBuffer *TextBuffer, format gdk.Atom, start, end *TextIter) string {
	var length = new(C.ulong)
	ptr := C.gtk_text_buffer_serialize(v.native(), contentBuffer.native(), C.GdkAtom(unsafe.Pointer(format)),
		(*C.GtkTextIter)(start), (*C.GtkTextIter)(end), length)
	return C.GoStringN((*C.char)(unsafe.Pointer(ptr)), (C.int)(*length))
}

// Deserialize() is a wrapper around  gtk_text_buffer_deserialize()
func (v *TextBuffer) Deserialize(contentBuffer *TextBuffer, format gdk.Atom, iter *TextIter, data []byte) (ok bool, err error) {
	var length = (C.ulong)(len(data))
	var cerr *C.GError = nil
	cbool := C.gtk_text_buffer_deserialize(v.native(), contentBuffer.native(), C.GdkAtom(unsafe.Pointer(format)),
		(*C.GtkTextIter)(iter), (*C.guchar)(unsafe.Pointer(&data[0])), length, &cerr)
	if !gobool(cbool) {
		defer C.g_error_free(cerr)
		return false, errors.New(goString(cerr.message))
	}
	return gobool(cbool), nil
}

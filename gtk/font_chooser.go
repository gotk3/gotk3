package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
"unsafe"

"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_font_chooser_get_type()), marshalFontChooser},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkFontChooser"] = wrapFontChooser
}

/*
 * GtkFontChooser
 */

// FontChooser is a representation of GTK's GtkFontChooser GInterface.
type FontChooser struct {
	*glib.Object
}

// IFontChooser is an interface type implemented by all structs
// embedding an FontChooser. It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkFontChooser.
type IFontChooser interface {
	toFontChooser() *C.GtkFontChooser
}

// native returns a pointer to the underlying GtkFontChooser.
func (v *FontChooser) native() *C.GtkFontChooser {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFontChooser(p)
}

func marshalFontChooser(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFontChooser(obj), nil
}

func wrapFontChooser(obj *glib.Object) *FontChooser {
	return &FontChooser{obj}
}

func (v *FontChooser) toFontChooser() *C.GtkFontChooser {
	if v == nil {
		return nil
	}
	return v.native()
}

// GetFont is a wrapper around gtk_font_chooser_get_font().
func (v *FontChooser) GetFont() string {
	c := C.gtk_font_chooser_get_font(v.native())
	return goString(c)
}

// SetFont is a wrapper around gtk_font_chooser_set_font().
func (v *FontChooser) SetFont(font string) {
	cstr := C.CString(font)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_font_chooser_set_font(v.native(), (*C.gchar)(cstr))
}


//PangoFontFamily *	gtk_font_chooser_get_font_family ()
//PangoFontFace *	gtk_font_chooser_get_font_face ()
//gint	gtk_font_chooser_get_font_size ()
//PangoFontDescription *	gtk_font_chooser_get_font_desc ()
//void	gtk_font_chooser_set_font_desc ()
//gchar *	gtk_font_chooser_get_preview_text ()
//void	gtk_font_chooser_set_preview_text ()
//gboolean	gtk_font_chooser_get_show_preview_entry ()
//void	gtk_font_chooser_set_show_preview_entry ()
//gboolean	(*GtkFontFilterFunc) ()
//void	gtk_font_chooser_set_filter_func ()
//void	gtk_font_chooser_set_font_map ()
//PangoFontMap *	gtk_font_chooser_get_font_map ()
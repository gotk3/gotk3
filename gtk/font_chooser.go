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
		{glib.Type(C.gtk_font_button_get_type()), marshalFontButton},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkFontChooser"] = wrapFontChooser
	WrapMap["GtkFontButton"] = wrapFontButton

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
	if obj == nil {
		return nil
	}

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

/*
 * GtkFontButton
 */

// FontButton is a representation of GTK's GtkFontButton.
type FontButton struct {
	Button

	// Interfaces
	FontChooser
}

// native returns a pointer to the underlying GtkFontButton.
func (v *FontButton) native() *C.GtkFontButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFontButton(p)
}

func marshalFontButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFontButton(obj), nil
}

func wrapFontButton(obj *glib.Object) *FontButton {
	if obj == nil {
		return nil
	}

	button := wrapButton(obj)
	fc := wrapFontChooser(obj)
	return &FontButton{*button, *fc}
}

// FontButtonNew is a wrapper around gtk_font_button_new().
func FontButtonNew() (*FontButton, error) {
	c := C.gtk_font_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFontButton(obj), nil
}

// FontButtonNewWithFont is a wrapper around gtk_font_button_new_with_font().
func FontButtonNewWithFont(fontname string) (*FontButton, error) {
	cstr := C.CString(fontname)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_font_button_new_with_font((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFontButton(obj), nil
}

// SetShowStyle is a wrapper around gtk_font_button_set_show_style().
func (v *FontButton) SetShowStyle(showStyle bool) {
	C.gtk_font_button_set_show_style(v.native(), gbool(showStyle))
}

// GetShowStyle is a wrapper around gtk_font_button_get_show_style().
func (v *FontButton) GetShowStyle() bool {
	c := C.gtk_font_button_get_show_style(v.native())
	return gobool(c)
}

// SetShowSize is a wrapper around gtk_font_button_set_show_size().
func (v *FontButton) SetShowSize(showSize bool) {
	C.gtk_font_button_set_show_size(v.native(), gbool(showSize))
}

// GetShowSize is a wrapper around gtk_font_button_get_show_size().
func (v *FontButton) GetShowSize() bool {
	c := C.gtk_font_button_get_show_size(v.native())
	return gobool(c)
}

// SetUseFont is a wrapper around gtk_font_button_set_use_font().
func (v *FontButton) SetUseFont(useFont bool) {
	C.gtk_font_button_set_use_font(v.native(), gbool(useFont))
}

// GetUseFont is a wrapper around gtk_font_button_get_use_font().
func (v *FontButton) GetUseFont() bool {
	c := C.gtk_font_button_get_use_font(v.native())
	return gobool(c)
}

// SetUseSize is a wrapper around gtk_font_button_set_use_size().
func (v *FontButton) SetUseSize(useSize bool) {
	C.gtk_font_button_set_use_size(v.native(), gbool(useSize))
}

// GetUseSize is a wrapper around gtk_font_button_get_use_size().
func (v *FontButton) GetUseSize() bool {
	c := C.gtk_font_button_get_use_size(v.native())
	return gobool(c)
}

// SetTitle is a wrapper around gtk_font_button_set_title().
func (v *FontButton) SetTitle(title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_font_button_set_title(v.native(), (*C.gchar)(cstr))
}

// GetTitle is a wrapper around gtk_font_button_get_title().
func (v *FontButton) GetTitle() string {
	c := C.gtk_font_button_get_title(v.native())
	defer C.free(unsafe.Pointer(c))
	return goString(c)
}

/*
 * GtkFontChooserWidget
 */

// TODO:
// gtk_font_chooser_widget_new().

package impl

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	gdk_impl "github.com/gotk3/gotk3/gdk/impl"
	"github.com/gotk3/gotk3/glib"
	glib_impl "github.com/gotk3/gotk3/glib/impl"
	"github.com/gotk3/gotk3/gtk"
)

func init() {
	tm := []glib_impl.TypeMarshaler{
		{glib.Type(C.gtk_color_chooser_get_type()), marshalColorChooser},
		{glib.Type(C.gtk_color_chooser_dialog_get_type()), marshalColorChooserDialog},
	}

	glib_impl.RegisterGValueMarshalers(tm)

	WrapMap["GtkColorChooser"] = wrapColorChooser
	WrapMap["GtkColorChooserDialog"] = wrapColorChooserDialog
}

/*
 * GtkColorChooser
 */

// ColorChooser is a representation of GTK's GtkColorChooser GInterface.
type colorChooser struct {
	*glib_impl.Object
}

// IColorChooser is an interface type implemented by all structs
// embedding an ColorChooser. It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkColorChooser.
type IColorChooser interface {
	toColorChooser() *C.GtkColorChooser
}

// native returns a pointer to the underlying GtkAppChooser.
func (v *colorChooser) native() *C.GtkColorChooser {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkColorChooser(p)
}

func marshalColorChooser(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapColorChooser(obj), nil
}

func wrapColorChooser(obj *glib_impl.Object) *colorChooser {
	return &colorChooser{obj}
}

func (v *colorChooser) toColorChooser() *C.GtkColorChooser {
	if v == nil {
		return nil
	}
	return v.native()
}

// GetRGBA() is a wrapper around gtk_color_chooser_get_rgba().
func (v *colorChooser) GetRGBA() gdk.RGBA {
	gdkColor := gdk_impl.NewRGBA()
	C.gtk_color_chooser_get_rgba(v.native(), (*C.GdkRGBA)(unsafe.Pointer(gdkColor.Native())))
	return gdkColor
}

// SetRGBA() is a wrapper around gtk_color_chooser_set_rgba().
func (v *colorChooser) SetRGBA(gdkColor gdk.RGBA) {
	C.gtk_color_chooser_set_rgba(v.native(), (*C.GdkRGBA)(unsafe.Pointer(gdk_impl.CastToRGBA(gdkColor).Native())))
}

// GetUseAlpha() is a wrapper around gtk_color_chooser_get_use_alpha().
func (v *colorChooser) GetUseAlpha() bool {
	return gobool(C.gtk_color_chooser_get_use_alpha(v.native()))
}

// SetUseAlpha() is a wrapper around gtk_color_chooser_set_use_alpha().
func (v *colorChooser) SetUseAlpha(use_alpha bool) {
	C.gtk_color_chooser_set_use_alpha(v.native(), gbool(use_alpha))
}

// AddPalette() is a wrapper around gtk_color_chooser_add_palette().
func (v *colorChooser) AddPalette(orientation gtk.Orientation, colors_per_line int, colors []gdk.RGBA) {
	n_colors := len(colors)
	var c_colors []C.GdkRGBA
	for _, c := range colors {
		c_colors = append(c_colors, *(*C.GdkRGBA)(unsafe.Pointer(gdk_impl.CastToRGBA(c).Native())))
	}
	C.gtk_color_chooser_add_palette(
		v.native(),
		C.GtkOrientation(orientation),
		C.gint(colors_per_line),
		C.gint(n_colors),
		&c_colors[0],
	)
}

/*
 * GtkColorChooserDialog
 */

// ColorChooserDialog is a representation of GTK's GtkColorChooserDialog.
type colorChooserDialog struct {
	dialog

	// Interfaces
	colorChooser
}

// native returns a pointer to the underlying GtkColorChooserButton.
func (v *colorChooserDialog) native() *C.GtkColorChooserDialog {
	if v == nil || v.GObject == nil {
		return nil
	}

	p := unsafe.Pointer(v.GObject)
	return C.toGtkColorChooserDialog(p)
}

func marshalColorChooserDialog(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapColorChooserDialog(wrapObject(unsafe.Pointer(c))), nil
}

func wrapColorChooserDialog(obj *glib_impl.Object) *colorChooserDialog {
	dialog := wrapDialog(obj)
	cc := wrapColorChooser(obj)
	return &colorChooserDialog{*dialog, *cc}
}

// ColorChooserDialogNew() is a wrapper around gtk_color_chooser_dialog_new().
func ColorChooserDialogNew(title string, parent *window) (*colorChooserDialog, error) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_color_chooser_dialog_new((*C.gchar)(cstr), parent.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapColorChooserDialog(wrapObject(unsafe.Pointer(c))), nil
}

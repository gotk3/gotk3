package gtk

/*
#cgo pkg-config: gtk+-3.0
#include <gtk/gtk.h>
#include <stdlib.h> // free

static GtkColorButton *
toGtkColorButton(void *p)
{
	return (GTK_COLOR_BUTTON(p));
}

static GtkColorChooser *
toGtkColorChooser(void *p)
{
	return (GTK_COLOR_CHOOSER(p));
}

*/
import "C"

import (
	"github.com/conformal/gotk3/gdk"
	"github.com/conformal/gotk3/glib"
	"runtime"
	"unsafe"
)

// OverrideColor is a wrapper around gtk_widget_override_color().
func (v *Widget) OverrideColor(state StateFlags, color *gdk.RGBA) {
	var cColor *C.GdkRGBA
	if color != nil {
		cColor = (*C.GdkRGBA)(unsafe.Pointer((&color.RGBA)))
	}
	C.gtk_widget_override_color(v.native(), C.GtkStateFlags(state), cColor)
}

//
//-------------------------------------------------------------[ ColorButton ]--

/*
 * GtkColorButton
 */

// ColorButton is a representation of GTK's GtkColorButton.
type ColorButton struct {
	Button
}

// Native returns a pointer to the underlying GtkColorButton.
func (v *ColorButton) native() *C.GtkColorButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkColorButton(p)
}

func wrapColorButton(obj *glib.Object) *ColorButton {
	return &ColorButton{Button{Bin{Container{Widget{
		glib.InitiallyUnowned{obj}}}}}}
}

// ColorButtonNew is a wrapper around gtk_color_button_new().
func ColorButtonNew() (*ColorButton, error) {
	c := C.gtk_color_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	tb := wrapColorButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return tb, nil
}

// ColorButtonNewWithRGBA is a wrapper around gtk_color_button_new_with_rgba().
func ColorButtonNewWithRGBA(gdkColor *gdk.RGBA) (*ColorButton, error) {
	c := C.gtk_color_button_new_with_rgba((*C.GdkRGBA)(unsafe.Pointer((&gdkColor.RGBA))))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	tb := wrapColorButton(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return tb, nil
}

func (v *ColorButton) GetRGBA() *gdk.RGBA {
	gdkColor := gdk.NewRGBA()
	C.gtk_color_chooser_get_rgba(C.toGtkColorChooser(unsafe.Pointer(v.native())), (*C.GdkRGBA)(unsafe.Pointer((&gdkColor.RGBA))))
	return gdkColor
}

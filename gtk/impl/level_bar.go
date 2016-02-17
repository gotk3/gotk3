package impl

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	glib_impl "github.com/gotk3/gotk3/glib/impl"
	"github.com/gotk3/gotk3/gtk"
)

func init() {
	tm := []glib_impl.TypeMarshaler{
		{glib.Type(C.gtk_level_bar_mode_get_type()), marshalLevelBarMode},

		{glib.Type(C.gtk_level_bar_get_type()), marshalLevelBar},
	}

	glib_impl.RegisterGValueMarshalers(tm)

	WrapMap["GtkLevelBar"] = wrapLevelBar

	gtk.LEVEL_BAR_MODE_CONTINUOUS = C.GTK_LEVEL_BAR_MODE_CONTINUOUS
	gtk.LEVEL_BAR_MODE_DISCRETE = C.GTK_LEVEL_BAR_MODE_DISCRETE

	gtk.LEVEL_BAR_OFFSET_LOW = C.GTK_LEVEL_BAR_OFFSET_LOW
	gtk.LEVEL_BAR_OFFSET_HIGH = C.GTK_LEVEL_BAR_OFFSET_HIGH
}

func marshalLevelBarMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.LevelBarMode(c), nil
}

/*
 * GtkLevelBar
 */

type levelBar struct {
	widget
}

// native returns a pointer to the underlying GtkLevelBar.
func (v *levelBar) native() *C.GtkLevelBar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkLevelBar(p)
}

func marshalLevelBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapLevelBar(obj), nil
}

func wrapLevelBar(obj *glib_impl.Object) *levelBar {
	return &levelBar{widget{glib_impl.InitiallyUnowned{obj}}}
}

// LevelBarNew() is a wrapper around gtk_level_bar_new().
func LevelBarNew() (*levelBar, error) {
	c := C.gtk_level_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapLevelBar(wrapObject(unsafe.Pointer(c))), nil
}

// LevelBarNewForInterval() is a wrapper around gtk_level_bar_new_for_interval().
func LevelBarNewForInterval(min_value, max_value float64) (*levelBar, error) {
	c := C.gtk_level_bar_new_for_interval(C.gdouble(min_value), C.gdouble(max_value))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapLevelBar(wrapObject(unsafe.Pointer(c))), nil
}

// SetMode() is a wrapper around gtk_level_bar_set_mode().
func (v *levelBar) SetMode(m gtk.LevelBarMode) {
	C.gtk_level_bar_set_mode(v.native(), C.GtkLevelBarMode(m))
}

// GetMode() is a wrapper around gtk_level_bar_get_mode().
func (v *levelBar) GetMode() gtk.LevelBarMode {
	return gtk.LevelBarMode(C.gtk_level_bar_get_mode(v.native()))
}

// SetValue() is a wrapper around gtk_level_bar_set_value().
func (v *levelBar) SetValue(value float64) {
	C.gtk_level_bar_set_value(v.native(), C.gdouble(value))
}

// GetValue() is a wrapper around gtk_level_bar_get_value().
func (v *levelBar) GetValue() float64 {
	c := C.gtk_level_bar_get_value(v.native())
	return float64(c)
}

// SetMinValue() is a wrapper around gtk_level_bar_set_min_value().
func (v *levelBar) SetMinValue(value float64) {
	C.gtk_level_bar_set_min_value(v.native(), C.gdouble(value))
}

// GetMinValue() is a wrapper around gtk_level_bar_get_min_value().
func (v *levelBar) GetMinValue() float64 {
	c := C.gtk_level_bar_get_min_value(v.native())
	return float64(c)
}

// SetMaxValue() is a wrapper around gtk_level_bar_set_max_value().
func (v *levelBar) SetMaxValue(value float64) {
	C.gtk_level_bar_set_max_value(v.native(), C.gdouble(value))
}

// GetMaxValue() is a wrapper around gtk_level_bar_get_max_value().
func (v *levelBar) GetMaxValue() float64 {
	c := C.gtk_level_bar_get_max_value(v.native())
	return float64(c)
}

// AddOffsetValue() is a wrapper around gtk_level_bar_add_offset_value().
func (v *levelBar) AddOffsetValue(name string, value float64) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_level_bar_add_offset_value(v.native(), (*C.gchar)(cstr), C.gdouble(value))
}

// RemoveOffsetValue() is a wrapper around gtk_level_bar_remove_offset_value().
func (v *levelBar) RemoveOffsetValue(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_level_bar_remove_offset_value(v.native(), (*C.gchar)(cstr))
}

// GetOffsetValue() is a wrapper around gtk_level_bar_get_offset_value().
func (v *levelBar) GetOffsetValue(name string) (float64, bool) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	var value C.gdouble
	c := C.gtk_level_bar_get_offset_value(v.native(), (*C.gchar)(cstr), &value)
	return float64(value), gobool(c)
}
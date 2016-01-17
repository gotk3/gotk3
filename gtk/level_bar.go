package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import (
	"C"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_level_bar_mode_get_type()), marshalLevelBarMode},

		{glib.Type(C.gtk_level_bar_get_type()), marshalLevelBar},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkLevelBar"] = wrapLevelBar
}

// LevelBarMode is a representation of GTK's GtkLevelBarMode.
type LevelBarMode int

const (
	LEVEL_BAR_MODE_CONTINUOUS LevelBarMode = C.GTK_LEVEL_BAR_MODE_CONTINUOUS
	LEVEL_BAR_MODE_DISCRETE   LevelBarMode = C.GTK_LEVEL_BAR_MODE_DISCRETE
)

func marshalLevelBarMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return LevelBarMode(c), nil
}

/*
 * GtkLevelBar
 */

type LevelBar struct {
	Widget
}

// native returns a pointer to the underlying GtkLevelBar.
func (v *LevelBar) native() *C.GtkLevelBar {
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

func wrapLevelBar(obj *glib.Object) *LevelBar {
	return &LevelBar{Widget{glib.InitiallyUnowned{obj}}}
}

// LevelBarNew() is a wrapper around gtk_level_bar_new().
func LevelBarNew() (*LevelBar, error) {
	c := C.gtk_level_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapLevelBar(wrapObject(unsafe.Pointer(c))), nil
}

// SetMode() is a wrapper around gtk_level_bar_set_mode().
func (v *LevelBar) SetMode(m LevelBarMode) {
	C.gtk_level_bar_set_mode(v.native(), C.GtkLevelBarMode(m))
}

// GetMode() is a wrapper around gtk_level_bar_get_mode().
func (v *LevelBar) GetMode() LevelBarMode {
	return LevelBarMode(C.gtk_level_bar_get_mode(v.native()))
}

// SetValue() is a wrapper around gtk_level_bar_set_value().
func (v *LevelBar) SetValue(value float64) {
	C.gtk_level_bar_set_value(v.native(), C.gdouble(value))
}

// GetValue() is a wrapper around gtk_level_bar_get_value().
func (v *LevelBar) GetValue() float64 {
	c := C.gtk_level_bar_get_value(v.native())
	return float64(c)
}

// SetMinValue() is a wrapper around gtk_level_bar_set_min_value().
func (v *LevelBar) SetMinValue(value float64) {
	C.gtk_level_bar_set_min_value(v.native(), C.gdouble(value))
}

// GetMinValue() is a wrapper around gtk_level_bar_get_min_value().
func (v *LevelBar) GetMinValue() float64 {
	c := C.gtk_level_bar_get_min_value(v.native())
	return float64(c)
}

// SetMaxValue() is a wrapper around gtk_level_bar_set_max_value().
func (v *LevelBar) SetMaxValue(value float64) {
	C.gtk_level_bar_set_max_value(v.native(), C.gdouble(value))
}

// GetMaxValue() is a wrapper around gtk_level_bar_get_max_value().
func (v *LevelBar) GetMaxValue() float64 {
	c := C.gtk_level_bar_get_max_value(v.native())
	return float64(c)
}

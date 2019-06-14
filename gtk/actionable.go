package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
// #include "actionable.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_actionable_get_type()), marshalActionable},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkActionable"] = marshalActionable
}

// IActionable is a representation of the GtkActionable GInterface,
// used to avoid duplication when embedding the type in a wrapper of another GObject-based type.
// The non-Interface version should only be used Actionable is used if the concrete type is not known.
type IActionable interface {
	Native() uintptr
	toActionable() *C.GtkActionable

	SetActionName(name string)
	GetActionName() (string, error)
	// SetActionTargetValue(value *glib.Variant)
	// GetActionTargetValue() (*glib.Variant, error)
	// SetActionTarget(string, params...)
	SetDetailedActionName(name string)
}

// Actionable is a representation of the GtkActionable GInterface.
// Do not embed this concrete type in implementing structs but rather use IActionable
// (see Button wrapper for an example)
type Actionable struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkActionable.
func (v *Actionable) native() *C.GtkActionable {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkActionable(p)
}

func marshalActionable(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapActionable(obj), nil
}

func wrapActionable(obj *glib.Object) *Actionable {
	return &Actionable{obj}
}

func (v *Actionable) toActionable() *C.GtkActionable {
	if v == nil {
		return nil
	}
	return v.native()
}

// SetActionName is a wrapper around gtk_actionable_set_action_name().
// Since 3.4
func (v *Actionable) SetActionName(action_name string) {
	cstr := C.CString(action_name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_actionable_set_action_name(v.native(), (*C.gchar)(cstr))
}

// GetActionName is a wrapper around gtk_actionable_set_action_name().
// Since 3.4
func (v *Actionable) GetActionName() (string, error) {
	c := C.gtk_actionable_get_action_name(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// SetDetailedActionName is a wrapper around gtk_actionable_set_detailed_action_name().
// Since 3.4
func (v *Actionable) SetDetailedActionName(detailed_action_name string) {
	cstr := C.CString(detailed_action_name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_actionable_set_detailed_action_name(v.native(), (*C.gchar)(cstr))
}

// SetActionTargetValue is a wrapper around gtk_actionable_set_action_target_value().
// Since 3.4
/*func (v *Actionable) SetActionTargetValue(value *glib.Variant) {
	// FIXME ToGVariant does not work here
	C.gtk_actionable_set_action_target_value(v.native(), value.ToGVariant())
}*/

// GetActionTargetValue is a wrapper around gtk_actionable_get_action_target_value().
// Since 3.4
/*func (v *Actionable) GetActionTargetValue() (*glib.Variant, error) {
	// FIXME: newVariant is not exported from glib
	return newVariant(C.gtk_actionable_get_action_target_value(v.native(), cstr))
}*/

/*
// Since 3.4
void
gtk_actionable_set_action_target (GtkActionable *actionable,
                                  const gchar *format_string,
								  ...);
*/

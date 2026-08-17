// +build !glib_2_34

package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
// #include "glib_since_2_40.go.h"
import "C"
import "unsafe"

func init() {
	tm := []TypeMarshaler{
		// Objects/Interfaces
		{Type(C.g_property_action_get_type()), marshalPropertyAction},
	}
	RegisterGValueMarshalers(tm)
}

/*
 * Action
 */

// ActionPrintDetailedName is a wrapper around g_action_print_detailed_name().
func ActionPrintDetailedName(action_name string, target_value *Variant) string {
	cstr := C.CString(action_name)
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString((*C.char)(C.g_action_print_detailed_name((*C.gchar)(cstr), target_value.native())))
}

// ActionNameIsValid is a wrapper around g_action_name_is_valid
func ActionNameIsValid(actionName string) bool {
	cstr := (*C.gchar)(C.CString(actionName))
	return gobool(C.g_action_name_is_valid(cstr))
}

/*
 * PropertyAction
 */

// PropertyAction is a representation of GPropertyAction
type PropertyAction struct {
	Action
}

func (v *PropertyAction) native() *C.GPropertyAction {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGPropertyAction(unsafe.Pointer(v.GObject))
}

func (v *PropertyAction) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalPropertyAction(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapPropertyAction(wrapObject(unsafe.Pointer(c))), nil
}

func wrapPropertyAction(obj *Object) *PropertyAction {
	return &PropertyAction{Action{obj}}
}

// PropertyActionNew is a wrapper around g_property_action_new
func PropertyActionNew(name string, object *Object, propertyName string) *PropertyAction {
	c := C.g_property_action_new((*C.gchar)(C.CString(name)), C.gpointer(unsafe.Pointer(object.native())), (*C.gchar)(C.CString(propertyName)))
	if c == nil {
		return nil
	}
	return wrapPropertyAction(wrapObject(unsafe.Pointer(c)))
}


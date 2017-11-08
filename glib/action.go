package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

/*
type ActionType int

const (
	ACTION_TYPE_ ActionType = C.GTK_
)
*/

// ActionGroup is a representation of GActionGroup.
type ActionGroup struct {
	*Object
}

// native() returns a pointer to the underlying GActionGroup.
func (v *ActionGroup) native() *C.GActionGroup {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGActionGroup(unsafe.Pointer(v.GObject))
}

func (v *ActionGroup) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalActionGroup(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapActionGroup(wrapObject(unsafe.Pointer(c))), nil
}

func wrapActionGroup(obj *Object) *ActionGroup {
	return &ActionGroup{obj}
}

// ActionGroup is a representation of GActionGroup.
type Action struct {
	*Object
}

// native() returns a pointer to the underlying GActionGroup.
func (v *Action) native() *C.GAction {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGAction(unsafe.Pointer(v.GObject))
}

func (v *Action) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalAction(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapAction(wrapObject(unsafe.Pointer(c))), nil
}

func wrapAction(obj *Object) *Action {
	return &Action{obj}
}

func (v *Action) NameIsValid(name string) bool {
	cstr := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(cstr))
	return gobool(C.g_action_name_is_valid(cstr))
}

func (v *Action) GetName() string {
	c := C.g_action_get_name(v.native())

	return C.GoString((*C.char)(c))
}

// XXX: g_action_get_parameter_type
// XXX: g_action_get_state_type
// XXX: g_action_get_state_hint

func (v *Action) GetEnabled() bool {
	return gobool(C.g_action_get_enabled(v.native()))
}

// XXX: g_action_get_state
// XXX: g_action_change_state
// XXX: g_action_activate
// XXX: g_action_parse_detailed_name
// XXX: g_action_print_detailed_name



// ActionGroup is a representation of GActionGroup.
type SimpleAction struct {
	Action
}

// native() returns a pointer to the underlying GActionGroup.
func (v *SimpleAction) native() *C.GSimpleAction {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGSimpleAction(unsafe.Pointer(v.GObject))
}

func (v *SimpleAction) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalSimpleAction(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapSimpleAction(wrapObject(unsafe.Pointer(c))), nil
}

func wrapSimpleAction(obj *Object) *SimpleAction {
	return &SimpleAction{*wrapAction(obj)}
}

func SimpleActionNew(name string, t *VariantType) *SimpleAction {
	cstr := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(cstr))
	c := C.g_simple_action_new(cstr, t.native())
	if c == nil {
		return nil
	}
	return wrapSimpleAction(wrapObject(unsafe.Pointer(c)))
}

func (v *SimpleAction) NewStateful(name string, t *VariantType, state *Variant) *SimpleAction {
	cstr := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(cstr))
	c := C.g_simple_action_new_stateful(cstr, t.native(), state.native())
	if c == nil {
		return nil
	}
	return wrapSimpleAction(wrapObject(unsafe.Pointer(c)))
}

func (v *SimpleAction) SetEnabled(state bool) {
	C.g_simple_action_set_enabled(v.native(), gbool(state))
}

func (v *SimpleAction) SetState(state Variant) {
	C.g_simple_action_set_state(v.native(), state.ToGVariant())
}

// XXX: g_simple_action_set_state_hint


// ActionGroup is a representation of GActionGroup.
type ActionMap struct {
	*Object
}

// native() returns a pointer to the underlying GActionGroup.
func (v *ActionMap) native() *C.GActionMap {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGActionMap(unsafe.Pointer(v.GObject))
}

func (v *ActionMap) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalActionMap(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapActionMap(wrapObject(unsafe.Pointer(c))), nil
}

func wrapActionMap(obj *Object) *ActionMap {
	return &ActionMap{obj}
}

func (v *ActionMap) LookupAction(name string) *Action {
	cstr := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(cstr))
	c := C.g_action_map_lookup_action(v.native(), cstr)
	if c == nil {
		return nil
	}
	return wrapAction(wrapObject(unsafe.Pointer(c)))
}
/*
// Requires GActionEntry
func (v *ActionMap) AddActionEntries() {
	
}
*/
func (v *ActionMap) AddAction(action *Action) {
	C.g_action_map_add_action(v.native(), action.native())
}

func (v *ActionMap) RemoveAction(name string) {
	cstr := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(cstr))
	C.g_action_map_remove_action(v.native(), cstr)
}


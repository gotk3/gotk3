// Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
//
// This file originated from: http://opensource.conformal.com/
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// Package glib provides Go bindings for GLib 2. It supports version 2.36 and
// later.
package glib

// #cgo pkg-config: gio-2.0 glib-2.0 gobject-2.0
// #include <gio/gio.h>
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/internal/callback"
	"github.com/gotk3/gotk3/internal/closure"
)

/*
 * Type conversions
 */

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

func gobool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

/*
 * Unexported vars
 */

var nilPtrErr = errors.New("cgo returned unexpected nil pointer")

/*
 * Constants
 */

// Type is a representation of GLib's GType.
type Type uint

const (
	TYPE_INVALID   Type = C.G_TYPE_INVALID
	TYPE_NONE      Type = C.G_TYPE_NONE
	TYPE_INTERFACE Type = C.G_TYPE_INTERFACE
	TYPE_CHAR      Type = C.G_TYPE_CHAR
	TYPE_UCHAR     Type = C.G_TYPE_UCHAR
	TYPE_BOOLEAN   Type = C.G_TYPE_BOOLEAN
	TYPE_INT       Type = C.G_TYPE_INT
	TYPE_UINT      Type = C.G_TYPE_UINT
	TYPE_LONG      Type = C.G_TYPE_LONG
	TYPE_ULONG     Type = C.G_TYPE_ULONG
	TYPE_INT64     Type = C.G_TYPE_INT64
	TYPE_UINT64    Type = C.G_TYPE_UINT64
	TYPE_ENUM      Type = C.G_TYPE_ENUM
	TYPE_FLAGS     Type = C.G_TYPE_FLAGS
	TYPE_FLOAT     Type = C.G_TYPE_FLOAT
	TYPE_DOUBLE    Type = C.G_TYPE_DOUBLE
	TYPE_STRING    Type = C.G_TYPE_STRING
	TYPE_POINTER   Type = C.G_TYPE_POINTER
	TYPE_BOXED     Type = C.G_TYPE_BOXED
	TYPE_PARAM     Type = C.G_TYPE_PARAM
	TYPE_OBJECT    Type = C.G_TYPE_OBJECT
	TYPE_VARIANT   Type = C.G_TYPE_VARIANT
)

// IsValue checks whether the passed in type can be used for g_value_init().
func (t Type) IsValue() bool {
	return gobool(C._g_type_is_value(C.GType(t)))
}

// Name is a wrapper around g_type_name().
func (t Type) Name() string {
	return C.GoString((*C.char)(C.g_type_name(C.GType(t))))
}

// Depth is a wrapper around g_type_depth().
func (t Type) Depth() uint {
	return uint(C.g_type_depth(C.GType(t)))
}

// Parent is a wrapper around g_type_parent().
func (t Type) Parent() Type {
	return Type(C.g_type_parent(C.GType(t)))
}

// IsA is a wrapper around g_type_is_a().
func (t Type) IsA(isAType Type) bool {
	return gobool(C.g_type_is_a(C.GType(t), C.GType(isAType)))
}

// TypeFromName is a wrapper around g_type_from_name
func TypeFromName(typeName string) Type {
	cstr := (*C.gchar)(C.CString(typeName))
	defer C.free(unsafe.Pointer(cstr))
	return Type(C.g_type_from_name(cstr))
}

//TypeNextBase is a wrapper around g_type_next_base
func TypeNextBase(leafType, rootType Type) Type {
	return Type(C.g_type_next_base(C.GType(leafType), C.GType(rootType)))
}

// SettingsBindFlags is a representation of GLib's GSettingsBindFlags.
type SettingsBindFlags int

const (
	SETTINGS_BIND_DEFAULT        SettingsBindFlags = C.G_SETTINGS_BIND_DEFAULT
	SETTINGS_BIND_GET            SettingsBindFlags = C.G_SETTINGS_BIND_GET
	SETTINGS_BIND_SET            SettingsBindFlags = C.G_SETTINGS_BIND_SET
	SETTINGS_BIND_NO_SENSITIVITY SettingsBindFlags = C.G_SETTINGS_BIND_NO_SENSITIVITY
	SETTINGS_BIND_GET_NO_CHANGES SettingsBindFlags = C.G_SETTINGS_BIND_GET_NO_CHANGES
	SETTINGS_BIND_INVERT_BOOLEAN SettingsBindFlags = C.G_SETTINGS_BIND_INVERT_BOOLEAN
)

// UserDirectory is a representation of GLib's GUserDirectory.
type UserDirectory int

const (
	USER_DIRECTORY_DESKTOP      UserDirectory = C.G_USER_DIRECTORY_DESKTOP
	USER_DIRECTORY_DOCUMENTS    UserDirectory = C.G_USER_DIRECTORY_DOCUMENTS
	USER_DIRECTORY_DOWNLOAD     UserDirectory = C.G_USER_DIRECTORY_DOWNLOAD
	USER_DIRECTORY_MUSIC        UserDirectory = C.G_USER_DIRECTORY_MUSIC
	USER_DIRECTORY_PICTURES     UserDirectory = C.G_USER_DIRECTORY_PICTURES
	USER_DIRECTORY_PUBLIC_SHARE UserDirectory = C.G_USER_DIRECTORY_PUBLIC_SHARE
	USER_DIRECTORY_TEMPLATES    UserDirectory = C.G_USER_DIRECTORY_TEMPLATES
	USER_DIRECTORY_VIDEOS       UserDirectory = C.G_USER_DIRECTORY_VIDEOS
)

const USER_N_DIRECTORIES int = C.G_USER_N_DIRECTORIES

/*
 * GApplicationFlags
 */

type ApplicationFlags int

const (
	APPLICATION_FLAGS_NONE           ApplicationFlags = C.G_APPLICATION_FLAGS_NONE
	APPLICATION_IS_SERVICE           ApplicationFlags = C.G_APPLICATION_IS_SERVICE
	APPLICATION_HANDLES_OPEN         ApplicationFlags = C.G_APPLICATION_HANDLES_OPEN
	APPLICATION_HANDLES_COMMAND_LINE ApplicationFlags = C.G_APPLICATION_HANDLES_COMMAND_LINE
	APPLICATION_SEND_ENVIRONMENT     ApplicationFlags = C.G_APPLICATION_SEND_ENVIRONMENT
	APPLICATION_NON_UNIQUE           ApplicationFlags = C.G_APPLICATION_NON_UNIQUE
)

// goMarshal is called by the GLib runtime when a closure needs to be invoked.
// The closure will be invoked with as many arguments as it can take, from 0 to
// the full amount provided by the call. If the closure asks for more parameters
// than there are to give, then a runtime panic will occur.
//
//export goMarshal
func goMarshal(
	gclosure *C.GClosure,
	retValue *C.GValue,
	nParams C.guint,
	params *C.GValue,
	invocationHint C.gpointer,
	marshalData *C.GValue) {

	// Get the function value associated with this callback closure.
	fs := closure.Get(unsafe.Pointer(gclosure))
	if !fs.IsValid() {
		// Possible data race, bail.
		return
	}

	fsType := fs.Func.Type()

	// Get number of parameters passed in.
	nGLibParams := int(nParams)
	nTotalParams := nGLibParams

	// Reflect may panic, so we defer recover here to re-panic with our trace.
	defer fs.TryRepanic()

	// Get number of parameters from the callback closure. If this exceeds
	// the total number of marshaled parameters, trigger a runtime panic.
	nCbParams := fsType.NumIn()
	if nCbParams > nTotalParams {
		fs.Panicf("too many closure args: have %d, max %d", nCbParams, nTotalParams)
	}

	// Create a slice of reflect.Values as arguments to call the function.
	gValues := gValueSlice(params, nCbParams)
	args := make([]reflect.Value, 0, nCbParams)

	// Fill beginning of args, up to the minimum of the total number of callback
	// parameters and parameters from the glib runtime.
	for i := 0; i < nCbParams && i < nGLibParams; i++ {
		v := Value{&gValues[i]}

		val, err := v.GoValue()
		if err != nil {
			fs.Panicf("no suitable Go value for arg %d: %v", i, err)
		}

		// Parameters that are descendants of GObject come wrapped in another
		// GObject. For C applications, the default marshaller
		// (g_cclosure_marshal_VOID__VOID in gmarshal.c in the GTK glib library)
		// 'peeks' into the enclosing object and passes the wrapped object to
		// the handler. Use the *Object.goValue function to emulate that for Go
		// signal handlers.
		switch objVal := val.(type) {
		case *Object:
			if innerVal, err := objVal.goValue(); err == nil {
				val = innerVal
			}

		case *Variant:
			switch ts := objVal.TypeString(); ts {
			case "s":
				val = objVal.GetString()
			case "b":
				val = gobool(C.g_variant_get_boolean(objVal.native()))
			case "d":
				val = float64(C.g_variant_get_double(objVal.native()))
			case "n":
				val = int16(C.g_variant_get_int16(objVal.native()))
			case "i":
				val = int32(C.g_variant_get_int32(objVal.native()))
			case "x":
				val = int64(C.g_variant_get_int64(objVal.native()))
			case "y":
				val = uint8(C.g_variant_get_byte(objVal.native()))
			case "q":
				val = uint16(C.g_variant_get_uint16(objVal.native()))
			case "u":
				val = uint32(C.g_variant_get_uint32(objVal.native()))
			case "t":
				val = uint64(C.g_variant_get_uint64(objVal.native()))
			default:
				fs.Panicf("variant conversion not yet implemented for type %s", ts)
			}
		}

		args = append(args, reflect.ValueOf(val).Convert(fsType.In(i)))
	}

	// Call closure with args. If the callback returns one or more values, save
	// the GValue equivalent of the first.
	rv := fs.Func.Call(args)
	if retValue != nil && len(rv) > 0 {
		g, err := GValue(rv[0].Interface())
		if err != nil {
			fs.Panicf("cannot save callback return value: %v", err)
		}

		t, _, err := g.Type()
		if err != nil {
			fs.Panicf("cannot determine callback return value: %v", err)
		}

		// Explicitly copy the return value as it may point to go-owned memory.
		C.g_value_unset(retValue)
		C.g_value_init(retValue, C.GType(t))
		C.g_value_copy(g.native(), retValue)
	}
}

// gValueSlice converts a C array of GValues to a Go slice.
func gValueSlice(values *C.GValue, nValues int) (slice []C.GValue) {
	header := (*reflect.SliceHeader)((unsafe.Pointer(&slice)))
	header.Cap = nValues
	header.Len = nValues
	header.Data = uintptr(unsafe.Pointer(values))
	return
}

/*
 * Main event loop
 */

// Priority is the enumerated type for GLib priority event sources.
type Priority int

const (
	PRIORITY_HIGH         Priority = C.G_PRIORITY_HIGH
	PRIORITY_DEFAULT      Priority = C.G_PRIORITY_DEFAULT // TimeoutAdd
	PRIORITY_HIGH_IDLE    Priority = C.G_PRIORITY_HIGH_IDLE
	PRIORITY_DEFAULT_IDLE Priority = C.G_PRIORITY_DEFAULT_IDLE // IdleAdd
	PRIORITY_LOW          Priority = C.G_PRIORITY_LOW
)

type SourceHandle uint

// sourceFunc is the callback for g_idle_add_full and g_timeout_add_full that
// replaces the GClosure API.
//
//export sourceFunc
func sourceFunc(data C.gpointer) C.gboolean {
	v := callback.Get(uintptr(data))
	fs := v.(closure.FuncStack)

	rv := fs.Func.Call(nil)
	if len(rv) == 1 && rv[0].Bool() {
		return C.TRUE
	}

	return C.FALSE
}

//export removeSourceFunc
func removeSourceFunc(data C.gpointer) {
	callback.Delete(uintptr(data))
}

var (
	_sourceFunc       = (*[0]byte)(C.sourceFunc)
	_removeSourceFunc = (*[0]byte)(C.removeSourceFunc)
)

// IdleAdd adds an idle source to the default main event loop context with the
// DefaultIdle priority. If f is not a function with no parameter, then IdleAdd
// will panic.
//
// After running once, the source func will be removed from the main event loop,
// unless f returns a single bool true.
func IdleAdd(f interface{}) SourceHandle {
	return idleAdd(PRIORITY_DEFAULT_IDLE, f)
}

// IdleAddPriority adds an idle source to the default main event loop context
// with the given priority. Its behavior is the same as IdleAdd.
func IdleAddPriority(priority Priority, f interface{}) SourceHandle {
	return idleAdd(priority, f)
}

func idleAdd(priority Priority, f interface{}) SourceHandle {
	fs := closure.NewIdleFuncStack(f, 2)
	id := C.gpointer(callback.Assign(fs))
	h := C.g_idle_add_full(C.gint(priority), _sourceFunc, id, _removeSourceFunc)

	return SourceHandle(h)
}

// TimeoutAdd adds an timeout source to the default main event loop context.
// Timeout is in milliseconds. If f is not a function with no parameter, then it
// will panic.
//
// After running once, the source func will be removed from the main event loop,
// unless f returns a single bool true.
func TimeoutAdd(milliseconds uint, f interface{}) SourceHandle {
	return timeoutAdd(milliseconds, false, PRIORITY_DEFAULT, f)
}

// TimeoutAddPriority is similar to TimeoutAdd with the given priority. Refer to
// TimeoutAdd for more information.
func TimeoutAddPriority(milliseconds uint, priority Priority, f interface{}) SourceHandle {
	return timeoutAdd(milliseconds, false, priority, f)
}

// TimeoutSecondsAdd is similar to TimeoutAdd, except with seconds granularity.
func TimeoutSecondsAdd(seconds uint, f interface{}) SourceHandle {
	return timeoutAdd(seconds, true, PRIORITY_DEFAULT, f)
}

// TimeoutSecondsAddPriority adds a timeout source with the given priority.
// Refer to TimeoutSecondsAdd for more information.
func TimeoutSecondsAddPriority(seconds uint, priority Priority, f interface{}) SourceHandle {
	return timeoutAdd(seconds, true, priority, f)
}

func timeoutAdd(time uint, sec bool, priority Priority, f interface{}) SourceHandle {
	fs := closure.NewIdleFuncStack(f, 2)
	id := C.gpointer(callback.Assign(fs))

	var h C.guint
	if sec {
		h = C.g_timeout_add_seconds_full(C.gint(priority), C.guint(time), _sourceFunc, id, _removeSourceFunc)
	} else {
		h = C.g_timeout_add_full(C.gint(priority), C.guint(time), _sourceFunc, id, _removeSourceFunc)
	}

	return SourceHandle(h)
}

// Destroy is a wrapper around g_source_destroy()
func (v *Source) Destroy() {
	C.g_source_destroy(v.native())
}

// IsDestroyed is a wrapper around g_source_is_destroyed()
func (v *Source) IsDestroyed() bool {
	return gobool(C.g_source_is_destroyed(v.native()))
}

// Unref is a wrapper around g_source_unref()
func (v *Source) Unref() {
	C.g_source_unref(v.native())
}

// Ref is a wrapper around g_source_ref()
func (v *Source) Ref() *Source {
	c := C.g_source_ref(v.native())
	if c == nil {
		return nil
	}
	return (*Source)(c)
}

// SourceRemove is a wrapper around g_source_remove()
func SourceRemove(src SourceHandle) bool {
	return gobool(C.g_source_remove(C.guint(src)))
}

/*
 * Miscellaneous Utility Functions
 */

// GetHomeDir is a wrapper around g_get_home_dir().
func GetHomeDir() string {
	c := C.g_get_home_dir()
	return C.GoString((*C.char)(c))
}

// GetUserCacheDir is a wrapper around g_get_user_cache_dir().
func GetUserCacheDir() string {
	c := C.g_get_user_cache_dir()
	return C.GoString((*C.char)(c))
}

// GetUserDataDir is a wrapper around g_get_user_data_dir().
func GetUserDataDir() string {
	c := C.g_get_user_data_dir()
	return C.GoString((*C.char)(c))
}

// GetUserConfigDir is a wrapper around g_get_user_config_dir().
func GetUserConfigDir() string {
	c := C.g_get_user_config_dir()
	return C.GoString((*C.char)(c))
}

// GetUserRuntimeDir is a wrapper around g_get_user_runtime_dir().
func GetUserRuntimeDir() string {
	c := C.g_get_user_runtime_dir()
	return C.GoString((*C.char)(c))
}

// GetUserSpecialDir is a wrapper around g_get_user_special_dir().  A
// non-nil error is returned in the case that g_get_user_special_dir()
// returns NULL to differentiate between NULL and an empty string.
func GetUserSpecialDir(directory UserDirectory) (string, error) {
	c := C.g_get_user_special_dir(C.GUserDirectory(directory))
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// FormatSize is a wrapper around g_format_size().
func FormatSize(size uint64) string {
	char := C.g_format_size(C.guint64(size))
	defer C.free(unsafe.Pointer(char))

	return C.GoString(char)
}

// FormatSizeFlags are flags to modify the format of the string returned by
// FormatSizeFull.
type FormatSizeFlags int

const (
	FORMAT_SIZE_DEFAULT     FormatSizeFlags = C.G_FORMAT_SIZE_DEFAULT
	FORMAT_SIZE_LONG_FORMAT FormatSizeFlags = C.G_FORMAT_SIZE_LONG_FORMAT
	FORMAT_SIZE_IEC_UNITS   FormatSizeFlags = C.G_FORMAT_SIZE_IEC_UNITS
	FORMAT_SIZE_BITS        FormatSizeFlags = C.G_FORMAT_SIZE_BITS
)

// FormatSizeFull is a wrapper around g_format_size_full().
func FormatSizeFull(size uint64, flags FormatSizeFlags) string {
	char := C.g_format_size_full(C.guint64(size), C.GFormatSizeFlags(flags))
	defer C.free(unsafe.Pointer(char))

	return C.GoString(char)
}

// SpacedPrimesClosest is a wrapper around g_spaced_primes_closest().
func SpacedPrimesClosest(num uint) uint {
	return uint(C.g_spaced_primes_closest(C.guint(num)))
}

/*
 * GObject
 */

// IObject is an interface type implemented by Object and all types which embed
// an Object. It is meant to be used as a type for function arguments which
// require GObjects or any subclasses thereof.
type IObject interface {
	toGObject() *C.GObject
	toObject() *Object
}

// Object is a representation of GLib's GObject.
type Object struct {
	GObject *C.GObject
}

func (v *Object) toGObject() *C.GObject {
	if v == nil {
		return nil
	}
	return v.native()
}

func (v *Object) toObject() *Object {
	return v
}

// newObject creates a new Object from a GObject pointer.
func newObject(p *C.GObject) *Object {
	if p == nil {
		return nil
	}
	return &Object{GObject: p}
}

// native returns a pointer to the underlying GObject.
func (v *Object) native() *C.GObject {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGObject(p)
}

// goValue converts a *Object to a Go type (e.g. *Object => *gtk.Entry).
// It is used in goMarshal to convert generic GObject parameters to
// signal handlers to the actual types expected by the signal handler.
func (v *Object) goValue() (interface{}, error) {
	objType := Type(C._g_type_from_instance(C.gpointer(v.native())))
	f, err := gValueMarshalers.lookupType(objType)
	if err != nil {
		return nil, err
	}

	// The marshalers expect Values, not Objects
	val, err := ValueInit(objType)
	if err != nil {
		return nil, err
	}
	val.SetInstance(uintptr(unsafe.Pointer(v.GObject)))
	rv, err := f(uintptr(unsafe.Pointer(val.native())))
	return rv, err
}

// Take wraps a unsafe.Pointer as a glib.Object, taking ownership of it.
// This function is exported for visibility in other gotk3 packages and
// is not meant to be used by applications.
//
// To be clear, this should mostly be used when Gtk says "transfer none". Refer
// to AssumeOwnership for more details.
func Take(ptr unsafe.Pointer) *Object {
	obj := newObject(ToGObject(ptr))
	if obj == nil {
		return nil
	}

	obj.RefSink()
	runtime.SetFinalizer(obj, (*Object).Unref)
	return obj
}

// AssumeOwnership is similar to Take, except the function does not take a
// reference. This is usually used for newly constructed objects that for some
// reason does not have an initial floating reference.
//
// To be clear, this should often be used when Gtk says "transfer full", as it
// means the ownership is transferred to the caller, so we can assume that much.
// This is in contrary to Take, which is used when Gtk says "transfer none", as
// we're now referencing an object that might possibly be kept, so we should
// mark as such.
func AssumeOwnership(ptr unsafe.Pointer) *Object {
	obj := newObject(ToGObject(ptr))
	runtime.SetFinalizer(obj, (*Object).Unref)
	return obj
}

// Native returns a pointer to the underlying GObject.
func (v *Object) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

// IsA is a wrapper around g_type_is_a().
func (v *Object) IsA(typ Type) bool {
	return gobool(C.g_type_is_a(C.GType(v.TypeFromInstance()), C.GType(typ)))
}

// TypeFromInstance is a wrapper around g_type_from_instance().
func (v *Object) TypeFromInstance() Type {
	c := C._g_type_from_instance(C.gpointer(unsafe.Pointer(v.native())))
	return Type(c)
}

// ToGObject type converts an unsafe.Pointer as a native C GObject.
// This function is exported for visibility in other gotk3 packages and
// is not meant to be used by applications.
func ToGObject(p unsafe.Pointer) *C.GObject {
	return (*C.GObject)(p)
	// return C.toGObject(p)
}

// Ref is a wrapper around g_object_ref().
func (v *Object) Ref() {
	C.g_object_ref(C.gpointer(v.GObject))
}

// Unref is a wrapper around g_object_unref().
func (v *Object) Unref() {
	C.g_object_unref(C.gpointer(v.GObject))
}

// RefSink is a wrapper around g_object_ref_sink().
func (v *Object) RefSink() {
	C.g_object_ref_sink(C.gpointer(v.GObject))
}

// IsFloating is a wrapper around g_object_is_floating().
func (v *Object) IsFloating() bool {
	c := C.g_object_is_floating(C.gpointer(v.GObject))
	return gobool(c)
}

// ForceFloating is a wrapper around g_object_force_floating().
func (v *Object) ForceFloating() {
	C.g_object_force_floating(v.GObject)
}

// StopEmission is a wrapper around g_signal_stop_emission_by_name().
func (v *Object) StopEmission(s string) {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.g_signal_stop_emission_by_name((C.gpointer)(v.GObject),
		(*C.gchar)(cstr))
}

// Set calls SetProperty.
func (v *Object) Set(name string, value interface{}) error {
	return v.SetProperty(name, value)
}

// GetPropertyType returns the Type of a property of the underlying GObject.
// If the property is missing it will return TYPE_INVALID and an error.
func (v *Object) GetPropertyType(name string) (Type, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	paramSpec := C.g_object_class_find_property(C._g_object_get_class(v.native()), (*C.gchar)(cstr))
	if paramSpec == nil {
		return TYPE_INVALID, errors.New("couldn't find Property")
	}
	return Type(paramSpec.value_type), nil
}

// GetProperty is a wrapper around g_object_get_property().
func (v *Object) GetProperty(name string) (interface{}, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	t, err := v.GetPropertyType(name)
	if err != nil {
		return nil, err
	}

	p, err := ValueInit(t)
	if err != nil {
		return nil, errors.New("unable to allocate value")
	}
	C.g_object_get_property(v.GObject, (*C.gchar)(cstr), p.native())
	return p.GoValue()
}

// SetProperty is a wrapper around g_object_set_property().
func (v *Object) SetProperty(name string, value interface{}) error {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	if _, ok := value.(Object); ok {
		value = value.(Object).GObject
	}

	p, err := GValue(value)
	if err != nil {
		return errors.New("Unable to perform type conversion")
	}
	C.g_object_set_property(v.GObject, (*C.gchar)(cstr), p.native())
	return nil
}

/*
 * GObject Signals
 */

// Emit is a wrapper around g_signal_emitv() and emits the signal
// specified by the string s to an Object.  Arguments to callback
// functions connected to this signal must be specified in args.  Emit()
// returns an interface{} which must be type asserted as the Go
// equivalent type to the return value for native C callback.
//
// Note that this code is unsafe in that the types of values in args are
// not checked against whether they are suitable for the callback.
func (v *Object) Emit(s string, args ...interface{}) (interface{}, error) {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))

	// Create array of this instance and arguments
	valv := C.alloc_gvalue_list(C.int(len(args)) + 1)
	defer C.free(unsafe.Pointer(valv))

	// Add args and valv
	val, err := GValue(v)
	if err != nil {
		return nil, errors.New("Error converting Object to GValue: " + err.Error())
	}
	C.val_list_insert(valv, C.int(0), val.native())
	for i := range args {
		val, err := GValue(args[i])
		if err != nil {
			return nil, fmt.Errorf("Error converting arg %d to GValue: %s", i, err.Error())
		}
		C.val_list_insert(valv, C.int(i+1), val.native())
	}

	t := v.TypeFromInstance()
	// TODO: use just the signal name
	id := C.g_signal_lookup((*C.gchar)(cstr), C.GType(t))

	ret, err := ValueAlloc()
	if err != nil {
		return nil, errors.New("Error creating Value for return value")
	}
	C.g_signal_emitv(valv, id, C.GQuark(0), ret.native())

	return ret.GoValue()
}

// HandlerBlock is a wrapper around g_signal_handler_block().
func (v *Object) HandlerBlock(handle SignalHandle) {
	C.g_signal_handler_block(C.gpointer(v.GObject), C.gulong(handle))
}

// HandlerUnblock is a wrapper around g_signal_handler_unblock().
func (v *Object) HandlerUnblock(handle SignalHandle) {
	C.g_signal_handler_unblock(C.gpointer(v.GObject), C.gulong(handle))
}

// HandlerDisconnect is a wrapper around g_signal_handler_disconnect().
func (v *Object) HandlerDisconnect(handle SignalHandle) {
	// Ensure that Gtk will not use the closure beforehand.
	C.g_signal_handler_disconnect(C.gpointer(v.GObject), C.gulong(handle))
	closure.DisconnectSignal(uint(handle))
}

// Wrapper function for new objects with reference management.
func wrapObject(ptr unsafe.Pointer) *Object {
	return Take(ptr)
}

/*
 * GInitiallyUnowned
 */

// InitiallyUnowned is a representation of GLib's GInitiallyUnowned.
type InitiallyUnowned struct {
	// This must be a pointer so copies of the ref-sinked object
	// do not outlive the original object, causing an unref
	// finalizer to prematurely run.
	*Object
}

// Native returns a pointer to the underlying GObject.  This is implemented
// here rather than calling Native on the embedded Object to prevent a nil
// pointer dereference.
func (v *InitiallyUnowned) Native() uintptr {
	if v == nil || v.Object == nil {
		return uintptr(unsafe.Pointer(nil))
	}
	return v.Object.Native()
}

/*
 * GValue
 */

// Value is a representation of GLib's GValue.
//
// Don't allocate Values on the stack or heap manually as they may not
// be properly unset when going out of scope. Instead, use ValueAlloc(),
// which will set the runtime finalizer to unset the Value after it has
// left scope.
type Value struct {
	GValue *C.GValue
}

// native returns a pointer to the underlying GValue.
func (v *Value) native() *C.GValue {
	return v.GValue
}

// Native returns a pointer to the underlying GValue.
func (v *Value) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

// IsValue checks if value is a valid and initialized GValue structure.
func (v *Value) IsValue() bool {
	return gobool(C._g_is_value(v.native()))
}

// TypeName gets the type name of value.
func (v *Value) TypeName() string {
	return C.GoString((*C.char)(C._g_value_type_name(v.native())))
}

// ValueAlloc allocates a Value and sets a runtime finalizer to call
// g_value_unset() on the underlying GValue after leaving scope.
// ValueAlloc() returns a non-nil error if the allocation failed.
func ValueAlloc() (*Value, error) {
	c := C._g_value_alloc()
	if c == nil {
		return nil, nilPtrErr
	}

	v := &Value{c}

	//An allocated GValue is not guaranteed to hold a value that can be unset
	//We need to double check before unsetting, to prevent:
	//`g_value_unset: assertion 'G_IS_VALUE (value)' failed`
	runtime.SetFinalizer(v, func(f *Value) {

		if !f.IsValue() {
			C.g_free(C.gpointer(f.native()))
			return
		}

		f.unset()
	})

	return v, nil
}

// ValueInit is a wrapper around g_value_init() and allocates and
// initializes a new Value with the Type t.  A runtime finalizer is set
// to call g_value_unset() on the underlying GValue after leaving scope.
// ValueInit() returns a non-nil error if the allocation failed.
func ValueInit(t Type) (*Value, error) {
	c := C._g_value_init(C.GType(t))
	if c == nil {
		return nil, nilPtrErr
	}

	v := &Value{c}

	runtime.SetFinalizer(v, (*Value).unset)
	return v, nil
}

// ValueFromNative returns a type-asserted pointer to the Value.
func ValueFromNative(l unsafe.Pointer) *Value {
	//TODO why it does not add finalizer to the value?
	return &Value{(*C.GValue)(l)}
}

func (v *Value) unset() {
	C.g_value_unset(v.native())
}

// Unset is wrapper for g_value_unset
func (v *Value) Unset() {
	v.unset()
}

// Type is a wrapper around the G_VALUE_HOLDS_GTYPE() macro and
// the g_value_get_gtype() function.  GetType() returns TYPE_INVALID if v
// does not hold a Type, or otherwise returns the Type of v.
func (v *Value) Type() (actual Type, fundamental Type, err error) {
	if !v.IsValue() {
		return actual, fundamental, errors.New("invalid GValue")
	}
	cActual := C._g_value_type(v.native())
	cFundamental := C._g_value_fundamental(cActual)
	return Type(cActual), Type(cFundamental), nil
}

// GValue converts a Go type to a comparable GValue.  GValue()
// returns a non-nil error if the conversion was unsuccessful.
func GValue(v interface{}) (gvalue *Value, err error) {
	if v == nil {
		val, err := ValueInit(TYPE_POINTER)
		if err != nil {
			return nil, err
		}
		val.SetPointer(uintptr(unsafe.Pointer(nil)))
		return val, nil
	}

	switch e := v.(type) {
	case bool:
		val, err := ValueInit(TYPE_BOOLEAN)
		if err != nil {
			return nil, err
		}
		val.SetBool(e)
		return val, nil

	case int8:
		val, err := ValueInit(TYPE_CHAR)
		if err != nil {
			return nil, err
		}
		val.SetSChar(e)
		return val, nil

	case int64:
		val, err := ValueInit(TYPE_INT64)
		if err != nil {
			return nil, err
		}
		val.SetInt64(e)
		return val, nil

	case int:
		val, err := ValueInit(TYPE_INT)
		if err != nil {
			return nil, err
		}
		val.SetInt(e)
		return val, nil

	case uint8:
		val, err := ValueInit(TYPE_UCHAR)
		if err != nil {
			return nil, err
		}
		val.SetUChar(e)
		return val, nil

	case uint64:
		val, err := ValueInit(TYPE_UINT64)
		if err != nil {
			return nil, err
		}
		val.SetUInt64(e)
		return val, nil

	case uint:
		val, err := ValueInit(TYPE_UINT)
		if err != nil {
			return nil, err
		}
		val.SetUInt(e)
		return val, nil

	case float32:
		val, err := ValueInit(TYPE_FLOAT)
		if err != nil {
			return nil, err
		}
		val.SetFloat(e)
		return val, nil

	case float64:
		val, err := ValueInit(TYPE_DOUBLE)
		if err != nil {
			return nil, err
		}
		val.SetDouble(e)
		return val, nil

	case string:
		val, err := ValueInit(TYPE_STRING)
		if err != nil {
			return nil, err
		}
		val.SetString(e)
		return val, nil

	case *Object:
		val, err := ValueInit(TYPE_OBJECT)
		if err != nil {
			return nil, err
		}
		val.SetInstance(uintptr(unsafe.Pointer(e.GObject)))
		return val, nil

	default:
		/* Try this since above doesn't catch constants under other types */
		rval := reflect.ValueOf(v)
		switch rval.Kind() {
		case reflect.Int8:
			val, err := ValueInit(TYPE_CHAR)
			if err != nil {
				return nil, err
			}
			val.SetSChar(int8(rval.Int()))
			return val, nil

		case reflect.Int16:
			return nil, errors.New("Type not implemented")

		case reflect.Int32:
			return nil, errors.New("Type not implemented")

		case reflect.Int64:
			val, err := ValueInit(TYPE_INT64)
			if err != nil {
				return nil, err
			}
			val.SetInt64(rval.Int())
			return val, nil

		case reflect.Int:
			val, err := ValueInit(TYPE_INT)
			if err != nil {
				return nil, err
			}
			val.SetInt(int(rval.Int()))
			return val, nil

		case reflect.Uintptr, reflect.Ptr:
			val, err := ValueInit(TYPE_POINTER)
			if err != nil {
				return nil, err
			}
			val.SetPointer(rval.Pointer())
			return val, nil
		}
	}

	return nil, errors.New("Type not implemented")
}

// GValueMarshaler is a marshal function to convert a GValue into an
// appropriate Go type.  The uintptr parameter is a *C.GValue.
type GValueMarshaler func(uintptr) (interface{}, error)

// TypeMarshaler represents an actual type and it's associated marshaler.
type TypeMarshaler struct {
	T Type
	F GValueMarshaler
}

// RegisterGValueMarshalers adds marshalers for several types to the
// internal marshalers map. Once registered, calling GoValue on any
// Value with a registered type will return the data returned by the
// marshaler.
func RegisterGValueMarshalers(tm []TypeMarshaler) {
	gValueMarshalers.register(tm)
}

type marshalMap map[Type]GValueMarshaler

// gValueMarshalers is a map of Glib types to functions to marshal a
// GValue to a native Go type.
var gValueMarshalers = marshalMap{
	TYPE_INVALID:   marshalInvalid,
	TYPE_NONE:      marshalNone,
	TYPE_INTERFACE: marshalInterface,
	TYPE_CHAR:      marshalChar,
	TYPE_UCHAR:     marshalUchar,
	TYPE_BOOLEAN:   marshalBoolean,
	TYPE_INT:       marshalInt,
	TYPE_LONG:      marshalLong,
	TYPE_ENUM:      marshalEnum,
	TYPE_INT64:     marshalInt64,
	TYPE_UINT:      marshalUint,
	TYPE_ULONG:     marshalUlong,
	TYPE_FLAGS:     marshalFlags,
	TYPE_UINT64:    marshalUint64,
	TYPE_FLOAT:     marshalFloat,
	TYPE_DOUBLE:    marshalDouble,
	TYPE_STRING:    marshalString,
	TYPE_POINTER:   marshalPointer,
	TYPE_BOXED:     marshalBoxed,
	TYPE_OBJECT:    marshalObject,
	TYPE_VARIANT:   marshalVariant,
}

func (m marshalMap) register(tm []TypeMarshaler) {
	for i := range tm {
		m[tm[i].T] = tm[i].F
	}
}

func (m marshalMap) lookup(v *Value) (GValueMarshaler, error) {
	actual, fundamental, err := v.Type()
	if err != nil {
		return nil, err
	}

	if f, ok := m[actual]; ok {
		return f, nil
	}
	if f, ok := m[fundamental]; ok {
		return f, nil
	}
	return nil, errors.New("missing marshaler for type")
}

func (m marshalMap) lookupType(t Type) (GValueMarshaler, error) {
	if f, ok := m[Type(t)]; ok {
		return f, nil
	}
	return nil, errors.New("missing marshaler for type")
}

func marshalInvalid(uintptr) (interface{}, error) {
	return nil, errors.New("invalid type")
}

func marshalNone(uintptr) (interface{}, error) {
	return nil, nil
}

func marshalInterface(uintptr) (interface{}, error) {
	return nil, errors.New("interface conversion not yet implemented")
}

func marshalChar(p uintptr) (interface{}, error) {
	c := C.g_value_get_schar((*C.GValue)(unsafe.Pointer(p)))
	return int8(c), nil
}

func marshalUchar(p uintptr) (interface{}, error) {
	c := C.g_value_get_uchar((*C.GValue)(unsafe.Pointer(p)))
	return uint8(c), nil
}

func marshalBoolean(p uintptr) (interface{}, error) {
	c := C.g_value_get_boolean((*C.GValue)(unsafe.Pointer(p)))
	return gobool(c), nil
}

func marshalInt(p uintptr) (interface{}, error) {
	c := C.g_value_get_int((*C.GValue)(unsafe.Pointer(p)))
	return int(c), nil
}

func marshalLong(p uintptr) (interface{}, error) {
	c := C.g_value_get_long((*C.GValue)(unsafe.Pointer(p)))
	return int(c), nil
}

func marshalEnum(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return int(c), nil
}

func marshalInt64(p uintptr) (interface{}, error) {
	c := C.g_value_get_int64((*C.GValue)(unsafe.Pointer(p)))
	return int64(c), nil
}

func marshalUint(p uintptr) (interface{}, error) {
	c := C.g_value_get_uint((*C.GValue)(unsafe.Pointer(p)))
	return uint(c), nil
}

func marshalUlong(p uintptr) (interface{}, error) {
	c := C.g_value_get_ulong((*C.GValue)(unsafe.Pointer(p)))
	return uint(c), nil
}

func marshalFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_flags((*C.GValue)(unsafe.Pointer(p)))
	return uint(c), nil
}

func marshalUint64(p uintptr) (interface{}, error) {
	c := C.g_value_get_uint64((*C.GValue)(unsafe.Pointer(p)))
	return uint64(c), nil
}

func marshalFloat(p uintptr) (interface{}, error) {
	c := C.g_value_get_float((*C.GValue)(unsafe.Pointer(p)))
	return float32(c), nil
}

func marshalDouble(p uintptr) (interface{}, error) {
	c := C.g_value_get_double((*C.GValue)(unsafe.Pointer(p)))
	return float64(c), nil
}

func marshalString(p uintptr) (interface{}, error) {
	c := C.g_value_get_string((*C.GValue)(unsafe.Pointer(p)))
	return C.GoString((*C.char)(c)), nil
}

func marshalBoxed(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return uintptr(unsafe.Pointer(c)), nil
}

func marshalPointer(p uintptr) (interface{}, error) {
	c := C.g_value_get_pointer((*C.GValue)(unsafe.Pointer(p)))
	return unsafe.Pointer(c), nil
}

func marshalObject(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return newObject((*C.GObject)(c)), nil
}

func marshalVariant(p uintptr) (interface{}, error) {
	c := C.g_value_get_variant((*C.GValue)(unsafe.Pointer(p)))
	return newVariant((*C.GVariant)(c)), nil
}

// GoValue converts a Value to comparable Go type.  GoValue()
// returns a non-nil error if the conversion was unsuccessful.  The
// returned interface{} must be type asserted as the actual Go
// representation of the Value.
//
// This function is a wrapper around the many g_value_get_*()
// functions, depending on the type of the Value.
func (v *Value) GoValue() (interface{}, error) {
	f, err := gValueMarshalers.lookup(v)
	if err != nil {
		return nil, err
	}

	//No need to add finalizer because it is already done by ValueAlloc and ValueInit
	rv, err := f(uintptr(unsafe.Pointer(v.native())))
	return rv, err
}

// SetBool is a wrapper around g_value_set_boolean().
func (v *Value) SetBool(val bool) {
	C.g_value_set_boolean(v.native(), gbool(val))
}

// SetSChar is a wrapper around g_value_set_schar().
func (v *Value) SetSChar(val int8) {
	C.g_value_set_schar(v.native(), C.gint8(val))
}

// SetInt64 is a wrapper around g_value_set_int64().
func (v *Value) SetInt64(val int64) {
	C.g_value_set_int64(v.native(), C.gint64(val))
}

// SetInt is a wrapper around g_value_set_int().
func (v *Value) SetInt(val int) {
	C.g_value_set_int(v.native(), C.gint(val))
}

// SetUChar is a wrapper around g_value_set_uchar().
func (v *Value) SetUChar(val uint8) {
	C.g_value_set_uchar(v.native(), C.guchar(val))
}

// SetUInt64 is a wrapper around g_value_set_uint64().
func (v *Value) SetUInt64(val uint64) {
	C.g_value_set_uint64(v.native(), C.guint64(val))
}

// SetUInt is a wrapper around g_value_set_uint().
func (v *Value) SetUInt(val uint) {
	C.g_value_set_uint(v.native(), C.guint(val))
}

// SetFloat is a wrapper around g_value_set_float().
func (v *Value) SetFloat(val float32) {
	C.g_value_set_float(v.native(), C.gfloat(val))
}

// SetDouble is a wrapper around g_value_set_double().
func (v *Value) SetDouble(val float64) {
	C.g_value_set_double(v.native(), C.gdouble(val))
}

// SetString is a wrapper around g_value_set_string().
func (v *Value) SetString(val string) {
	cstr := C.CString(val)
	defer C.free(unsafe.Pointer(cstr))
	C.g_value_set_string(v.native(), (*C.gchar)(cstr))
}

// SetInstance is a wrapper around g_value_set_instance().
func (v *Value) SetInstance(instance uintptr) {
	C.g_value_set_instance(v.native(), C.gpointer(instance))
}

// SetPointer is a wrapper around g_value_set_pointer().
func (v *Value) SetPointer(p uintptr) {
	C.g_value_set_pointer(v.native(), C.gpointer(p))
}

// GetPointer is a wrapper around g_value_get_pointer().
func (v *Value) GetPointer() unsafe.Pointer {
	return unsafe.Pointer(C.g_value_get_pointer(v.native()))
}

// GetString is a wrapper around g_value_get_string().  GetString()
// returns a non-nil error if g_value_get_string() returned a NULL
// pointer to distinguish between returning a NULL pointer and returning
// an empty string.
func (v *Value) GetString() (string, error) {
	c := C.g_value_get_string(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

type Signal struct {
	name     string
	signalId C.guint
}

func SignalNew(s string) (*Signal, error) {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))

	signalId := C._g_signal_new((*C.gchar)(cstr))

	if signalId == 0 {
		return nil, fmt.Errorf("invalid signal name: %s", s)
	}

	return &Signal{
		name:     s,
		signalId: signalId,
	}, nil
}

func (s *Signal) String() string {
	return s.name
}

type Quark uint32

// GetPrgname is a wrapper around g_get_prgname().
func GetPrgname() string {
	c := C.g_get_prgname()

	return C.GoString((*C.char)(c))
}

// SetPrgname is a wrapper around g_set_prgname().
func SetPrgname(name string) {
	cstr := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(cstr))

	C.g_set_prgname(cstr)
}

// GetApplicationName is a wrapper around g_get_application_name().
func GetApplicationName() string {
	c := C.g_get_application_name()

	return C.GoString((*C.char)(c))
}

// SetApplicationName is a wrapper around g_set_application_name().
func SetApplicationName(name string) {
	cstr := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(cstr))

	C.g_set_application_name(cstr)
}

// InitI18n initializes the i18n subsystem.
func InitI18n(domain string, dir string) {
	domainStr := C.CString(domain)
	defer C.free(unsafe.Pointer(domainStr))

	dirStr := C.CString(dir)
	defer C.free(unsafe.Pointer(dirStr))

	C.init_i18n(domainStr, dirStr)
}

// Local localizes a string using gettext
func Local(input string) string {
	cstr := C.CString(input)
	defer C.free(unsafe.Pointer(cstr))

	return C.GoString(C.localize(cstr))
}

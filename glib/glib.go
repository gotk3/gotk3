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

// Go bindings for GLib 2.  Supports version 2.36 and later.
package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"sync"
	"unsafe"
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

type closureContext struct {
	rf reflect.Value

	// Never used but must be kept in go scope while the closure is valid
	// or the finalizer will run.
	userData *Value
}

var (
	nilPtrErr = errors.New("cgo returned unexpected nil pointer")

	closures = struct {
		sync.RWMutex
		m map[*C.GClosure]closureContext
	}{
		m: make(map[*C.GClosure]closureContext),
	}

	signals = make(map[SignalHandle]*C.GClosure)
)

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
 * Events
 */

type SignalHandle uint

// Connect is a wrapper around g_signal_connect_closure().  f must be
// a function with a signaure matching the callback signature for
// detailedSignal.  userData must either 0 or 1 elements which can
// be optionally passed to f.  If f takes less arguments than it is
// passed from the GLib runtime, the extra arguments are ignored.
//
// Currently, the non userdata arguments to f must be a *Object if
// the callback calls for any GObject.  Trying to use other GObject
// types will cause a panic when the callback is run.  This is a bug.
// The type of the optional user data argument in f may be any type
// that Go can type convert userData[0] as.  Non-GObject types may
// always be used as their Go equivalent types (for example, *C.gchar
// as a Go string).
func (v *Object) Connect(detailedSignal string, f interface{}, userData ...interface{}) (SignalHandle, error) {
	if len(userData) > 1 {
		return 0, errors.New("userData len must be 0 or 1")
	}

	cstr := C.CString(detailedSignal)
	defer C.free(unsafe.Pointer(cstr))

	closure, err := ClosureNew(f, userData...)
	if err != nil {
		return 0, err
	}

	C._g_closure_add_finalize_notifier(closure)

	c := C.g_signal_connect_closure(C.gpointer(v.Native()),
		(*C.gchar)(cstr), closure, gbool(false))
	handle := SignalHandle(c)

	// Map the signal handle to the closure.
	signals[handle] = closure

	return handle, nil
}

// ClosureNew creates a new GClosure and adds its callback function
// to the internally-maintained map. It's exported for visibility to other
// gotk3 packages and shouldn't be used in application code.
func ClosureNew(f interface{}, marshalData ...interface{}) (*C.GClosure, error) {
	// Create a reflect.Value from f.  This is called when the
	// returned GClosure runs.
	rf := reflect.ValueOf(f)

	// Create closure context which points to the reflected func.
	cc := closureContext{rf: rf}

	// Closures can only be created from funcs.
	if rf.Type().Kind() != reflect.Func {
		return nil, errors.New("value is not a func")
	}

	var c *C.GClosure
	if len(marshalData) == 0 {
		c = C._g_closure_new()
	} else {
		p, err := reflectGValue(marshalData[0])
		if err != nil {
			return nil, err
		}
		cc.userData = p
		c = C._g_closure_new_with_data((C.gpointer)(p))
	}

	// Associate the GClosure with rf.  rf will be looked up in this
	// map by the closure when the closure runs.
	closures.Lock()
	closures.m[c] = cc
	closures.Unlock()

	return c, nil
}

// removeClosure removes a closure from the internal closures map.  This is
// needed to prevent a leak where Go code can access the closure context
// (along with rf and userdata) even after an object has been destroyed and
// the GClosure is invalidated and will never run.
//
//export removeClosure
func removeClosure(_ C.gpointer, closure *C.GClosure) {
	closures.Lock()
	delete(closures.m, closure)
	closures.Unlock()
}

// goMarshal is called by the GLib runtime when a closure needs to be invoked.
// The closure will be invoked with as many arguments as it can take, from 0 to
// the full amount provided by the call. If the closure asks for more parameters
// than there are to give, a warning is printed to stderr and the closure is
// not run.
//
//export goMarshal
func goMarshal(closure *C.GClosure, retValue *C.GValue,
	nParams C.guint, params *C.GValue,
	invocationHint C.gpointer, marshalData *C.GValue) {

	// Get number of parameters passed in. If marshaled data was passed
	// in, increment the total number of parameters.
	nGLibParams := int(nParams)
	nTotalParams := nGLibParams
	if marshalData != nil {
		nTotalParams++
	}

	// Get the context associated with this callback closure.
	closures.RLock()
	cc := closures.m[closure]
	closures.RUnlock()

	// Get number of parameters from the callback closure.  If this exceeds
	// the total number of marshaled parameters, a warning will be printed
	// to stderr, and the callback will not be run.
	nCbParams := cc.rf.Type().NumIn()
	if nCbParams > nTotalParams {
		fmt.Fprintf(os.Stderr,
			"too many closure args: have %d, max allowed %d\n",
			nCbParams, nTotalParams)
		return
	}

	// Create a slice of reflect.Values as arguments to call the function.
	gValues := gValueSlice(params, nCbParams)
	args := make([]reflect.Value, 0, nCbParams)

	// Fill beginning of args, up to the minimum of the total number of callback
	// parameters and parameters from the glib runtime.
	for i := 0; i < nCbParams && i < nGLibParams; i++ {
		v := &Value{gValues[i]}
		val, err := v.GoValue()
		if err != nil {
			fmt.Fprintf(os.Stderr,
				"no suitable Go value for arg %d: %v\n", i, err)
			return
		}
		rv := reflect.ValueOf(val)
		args = append(args, rv.Convert(cc.rf.Type().In(i)))
	}

	// If non-nil marshaled user data was passed in and not all
	// args have been set, get and set the reflect.Value directly
	// from the GValue.
	if marshalData != nil && len(args) < cap(args) {
		v := &Value{*marshalData}
		rv := *v.reflectGoValue()
		args = append(args, rv.Convert(cc.rf.Type().In(nCbParams-1)))
	}

	// Call closure with args. If the callback returns one or more
	// values, save the GValue equivalent of the first.
	rv := cc.rf.Call(args)
	if retValue != nil && len(rv) > 0 {
		if g, err := GValue(rv[0].Interface()); err != nil {
			fmt.Fprintf(os.Stderr,
				"cannot save callback return value: %v", err)
		} else {
			*retValue = *g.Native()
		}
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

type Source struct {
	ptr unsafe.Pointer
}

type SourceHandle uint

// Native returns a pointer to the underlying GSource.
func (v *Source) Native() *C.GSource {
	if v == nil || v.ptr == nil {
		return nil
	}
	return (*C.GSource)(v.ptr)
}

// IdleAdd adds an idle source to the default main event loop
// context.  After running once, the source func will be removed
// from the main event loop, unless f returns a single bool true.
//
// This function will cause a panic when f eventually runs if the
// types of args do not match those of f.
func IdleAdd(f interface{}, args ...interface{}) (SourceHandle, error) {
	// f must be a func with no parameters.
	rf := reflect.ValueOf(f)
	if rf.Type().Kind() != reflect.Func {
		return 0, errors.New("f is not a function")
	}

	// Create an idle source func for a main loop context.
	idleSrc := C.g_idle_source_new()
	if idleSrc == nil {
		return 0, nilPtrErr
	}

	// Create a new GClosure from f that invalidates itself when
	// f returns false.  The error is ignored here, as this will
	// always be a function.
	var closure *C.GClosure
	closure, _ = ClosureNew(func() {
		// Create a slice of reflect.Values arguments to call the func.
		rargs := make([]reflect.Value, len(args))
		for i := range args {
			rargs[i] = reflect.ValueOf(args[i])
		}

		// Call func with args. The callback will be removed, unless
		// it returns exactly one return value of true.
		rv := rf.Call(rargs)
		if len(rv) == 1 {
			if rv[0].Kind() == reflect.Bool {
				if rv[0].Bool() {
					return
				}
			}
		}
		C.g_closure_invalidate(closure)
		C.g_source_destroy(idleSrc)
	})

	// Remove closure context when closure is finalized.
	C._g_closure_add_finalize_notifier(closure)

	// Set closure to run as a callback when the idle source runs.
	C.g_source_set_closure(idleSrc, closure)

	// Attach the idle source func to the default main event loop
	// context.
	cid := C.g_source_attach(idleSrc, nil)
	return SourceHandle(cid), nil
}

/*
 * Miscellaneous Utility Functions
 */

// GetUserSpecialDir() is a wrapper around g_get_user_special_dir().  A
// non-nil error is returned in the case that g_get_user_special_dir()
// returns NULL to differentiate between NULL and an empty string.
func GetUserSpecialDir(directory UserDirectory) (string, error) {
	c := C.g_get_user_special_dir(C.GUserDirectory(directory))
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

/*
 * GObject
 */

// IObject is an interface type implemented by Object and all types which embed
// an Object.  It is meant to be used as a type for function arguments which
// require GObjects or any subclasses thereof.
type IObject interface {
	toGObject() *C.GObject
}

// Object is a representation of GLib's GObject.
type Object struct {
	GObject *C.GObject
}

// ObjectNew creates a new Object from a Pointer.
func ObjectNew(p unsafe.Pointer) *Object {
	return &Object{GObject: (*C.GObject)(p)}
}

// Native() returns a pointer to the underlying GObject.
func (v *Object) Native() *C.GObject {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGObject(p)
}

// IsA is a wrapper around g_type_is_a().
func (v *Object) IsA(typ Type) bool {
	return gobool(C.g_type_is_a(C.GType(v.TypeFromInstance()), C.GType(typ)))
}

func (v *Object) toGObject() *C.GObject {
	if v == nil {
		return nil
	}
	return v.Native()
}

// TypeFromInstance is a wrapper around g_type_from_instance().
func (v *Object) TypeFromInstance() Type {
	c := C._g_type_from_instance(C.gpointer(unsafe.Pointer(v.Native())))
	return Type(c)
}

// ToGObject() type converts an unsafe.Pointer as a native C GObject.
// This function is exported for visibility in other gotk3 packages and
// is not meant to be used by applications.
func ToGObject(p unsafe.Pointer) *C.GObject {
	return C.toGObject(p)
}

// Ref() is a wrapper around g_object_ref().
func (v *Object) Ref() {
	C.g_object_ref(C.gpointer(v.GObject))
}

// Unref() is a wrapper around g_object_unref().
func (v *Object) Unref() {
	C.g_object_unref(C.gpointer(v.GObject))
}

// RefSink() is a wrapper around g_object_ref_sink().
func (v *Object) RefSink() {
	C.g_object_ref_sink(C.gpointer(v.GObject))
}

// IsFloating() is a wrapper around g_object_is_floating().
func (v *Object) IsFloating() bool {
	c := C.g_object_is_floating(C.gpointer(v.GObject))
	return gobool(c)
}

// ForceFloating() is a wrapper around g_object_force_floating().
func (v *Object) ForceFloating() {
	C.g_object_force_floating(v.GObject)
}

// StopEmission() is a wrapper around g_signal_stop_emission_by_name().
func (v *Object) StopEmission(s string) {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.g_signal_stop_emission_by_name((C.gpointer)(v.GObject),
		(*C.gchar)(cstr))
}

// Set() is a wrapper around g_object_set().  However, unlike
// g_object_set(), this function only sets one name value pair.  Make
// multiple calls to this function to set multiple properties.
func (v *Object) Set(name string, value interface{}) error {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	if _, ok := value.(Object); ok {
		value = value.(Object).GObject
	}

	// Can't call g_object_set() as it uses a variable arg list, use a
	// wrapper instead
	if p := pointerVal(value); p != nil {
		C._g_object_set_one(C.gpointer(v.GObject), (*C.gchar)(cstr), p)
		return nil
	} else {
		return errors.New("Unable to perform type conversion")
	}
}

// pointerVal attempts to return an unsafe.Pointer for value.
// Not all types are understood, in which case a nil Pointer
// is returned.
func pointerVal(value interface{}) unsafe.Pointer {
	var p unsafe.Pointer = nil
	switch v := value.(type) {
	case bool:
		c := gbool(v)
		p = unsafe.Pointer(&c)

	case int8:
		c := C.gint8(v)
		p = unsafe.Pointer(&c)

	case int16:
		c := C.gint16(v)
		p = unsafe.Pointer(&c)

	case int32:
		c := C.gint32(v)
		p = unsafe.Pointer(&c)

	case int64:
		c := C.gint64(v)
		p = unsafe.Pointer(&c)

	case int:
		c := C.gint(v)
		p = unsafe.Pointer(&c)

	case uint8:
		c := C.guchar(v)
		p = unsafe.Pointer(&c)

	case uint16:
		c := C.guint16(v)
		p = unsafe.Pointer(&c)

	case uint32:
		c := C.guint32(v)
		p = unsafe.Pointer(&c)

	case uint64:
		c := C.guint64(v)
		p = unsafe.Pointer(&c)

	case uint:
		c := C.guint(v)
		p = unsafe.Pointer(&c)

	case uintptr:
		p = unsafe.Pointer(C.gpointer(v))

	case float32:
		c := C.gfloat(v)
		p = unsafe.Pointer(&c)

	case float64:
		c := C.gdouble(v)
		p = unsafe.Pointer(&c)

	case string:
		cstr := C.CString(v)
		defer C.free(unsafe.Pointer(cstr))
		p = unsafe.Pointer(cstr)

	default:
		if pv, ok := value.(unsafe.Pointer); ok {
			p = pv
		} else {
			val := reflect.ValueOf(value)
			switch val.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16,
				reflect.Int32, reflect.Int64:
				c := C.int(val.Int())
				p = unsafe.Pointer(&c)

			case reflect.Uintptr, reflect.Ptr, reflect.UnsafePointer:
				p = unsafe.Pointer(C.gpointer(val.Pointer()))
			}
		}
	}

	return p
}

/*
 * GObject Signals
 */

// Emit() is a wrapper around g_signal_emitv() and emits the signal
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
	C.val_list_insert(valv, C.int(0), val.Native())
	for i := range args {
		val, err := GValue(args[i])
		if err != nil {
			return nil, fmt.Errorf("Error converting arg %d to GValue: %s", i, err.Error())
		}
		C.val_list_insert(valv, C.int(i+1), val.Native())
	}

	t := v.TypeFromInstance()
	// TODO: use just the signal name
	id := C.g_signal_lookup((*C.gchar)(cstr), C.GType(t))

	ret, err := ValueAlloc()
	if err != nil {
		return nil, errors.New("Error creating Value for return value")
	}
	C.g_signal_emitv(valv, id, C.GQuark(0), ret.Native())

	return ret.GoValue()
}

// HandlerBlock() is a wrapper around g_signal_handler_block().
func (v *Object) HandlerBlock(handle SignalHandle) {
	C.g_signal_handler_block(C.gpointer(v.GObject), C.gulong(handle))
}

// HandlerUnblock() is a wrapper around g_signal_handler_unblock().
func (v *Object) HandlerUnblock(handle SignalHandle) {
	C.g_signal_handler_unblock(C.gpointer(v.GObject), C.gulong(handle))
}

// HandlerDisconnect() is a wrapper around g_signal_handler_disconnect().
func (v *Object) HandlerDisconnect(handle SignalHandle) {
	C.g_signal_handler_disconnect(C.gpointer(v.GObject), C.gulong(handle))
	C.g_closure_invalidate(signals[handle])
	delete(closures.m, signals[handle])
	delete(signals, handle)
}

/*
 * GInitiallyUnowned
 */

// InitiallyUnowned is a representation of GLib's GInitiallyUnowned.
type InitiallyUnowned struct {
	*Object
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
	GValue C.GValue
}

// Native() returns a pointer to the underlying GValue.
func (v *Value) Native() *C.GValue {
	return &v.GValue
}

// ValueAlloc() allocates a Value and sets a runtime finalizer to call
// g_value_unset() on the underlying GValue after leaving scope.
// ValueAlloc() returns a non-nil error if the allocation failed.
func ValueAlloc() (*Value, error) {
	c := C._g_value_alloc()
	if c == nil {
		return nil, nilPtrErr
	}
	v := &Value{*c}
	runtime.SetFinalizer(v, (*Value).unset)
	return v, nil
}

// ValueInit() is a wrapper around g_value_init() and allocates and
// initializes a new Value with the Type t.  A runtime finalizer is set
// to call g_value_unset() on the underlying GValue after leaving scope.
// ValueInit() returns a non-nil error if the allocation failed.
func ValueInit(t Type) (*Value, error) {
	c := C._g_value_init(C.GType(t))
	if c == nil {
		return nil, nilPtrErr
	}
	v := &Value{*c}
	runtime.SetFinalizer(v, (*Value).unset)
	return v, nil
}

func (v *Value) unset() {
	C.g_value_unset(v.Native())
}

// Type() is a wrapper around the G_VALUE_HOLDS_GTYPE() macro and
// the g_value_get_gtype() function.  GetType() returns TYPE_INVALID if v
// does not hold a Type, or otherwise returns the Type of v.
func (v *Value) Type() (actual Type, fundamental Type, err error) {
	if !gobool(C._g_is_value(v.Native())) {
		return actual, fundamental, errors.New("invalid GValue")
	}
	c_actual := C._g_value_type(v.Native())
	c_fundamental := C._g_value_fundamental(c_actual)
	return Type(c_actual), Type(c_fundamental), nil
}

// GValue() converts a Go type to a comparable GValue.  GValue()
// returns a non-nil error if the conversion was unsuccessful.
func GValue(v interface{}) (gvalue *Value, err error) {
	if v == nil {
		val, err := ValueInit(TYPE_POINTER)
		if err != nil {
			return nil, err
		}
		val.SetPointer(uintptr(0)) // technically not portable
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

// GoValue() converts a Value to comparable Go type.  GoValue()
// returns a non-nil error if the conversion was unsuccessful.  The
// returned interface{} must be type asserted as the actual Go
// representation of the Value.
//
// This function is a wrapper around the many g_value_get_*()
// functions, depending on the type of the Value.
func (v *Value) GoValue() (interface{}, error) {
	_, fundamental, err := v.Type()
	if err != nil {
		return nil, err
	}

	// TODO: verify that all of these cases are indeed fundamental types
	switch fundamental {
	case TYPE_INVALID:
		return nil, errors.New("invalid type")

	case TYPE_NONE:
		return nil, nil

	case TYPE_INTERFACE:
		return nil, errors.New("interface conversion not yet implemented")

	case TYPE_CHAR:
		c := C.g_value_get_schar(v.Native())
		return int8(c), nil

	case TYPE_UCHAR:
		c := C.g_value_get_uchar(v.Native())
		return uint8(c), nil

	case TYPE_BOOLEAN:
		c := C.g_value_get_boolean(v.Native())
		return gobool(c), nil

	// TODO: TYPE_INT should probably be a Go int32.
	case TYPE_INT, TYPE_LONG, TYPE_ENUM:
		c := C.g_value_get_int(v.Native())
		return int(c), nil

	case TYPE_INT64:
		c := C.g_value_get_int64(v.Native())
		return int64(c), nil

	// TODO: TYPE_UINT should probably be a Go uint32.
	case TYPE_UINT, TYPE_ULONG, TYPE_FLAGS:
		c := C.g_value_get_uint(v.Native())
		return uint(c), nil

	case TYPE_UINT64:
		c := C.g_value_get_uint64(v.Native())
		return uint64(c), nil

	case TYPE_FLOAT:
		c := C.g_value_get_float(v.Native())
		return float32(c), nil

	case TYPE_DOUBLE:
		c := C.g_value_get_double(v.Native())
		return float64(c), nil

	case TYPE_STRING:
		c := C.g_value_get_string(v.Native())
		return C.GoString((*C.char)(c)), nil

	case TYPE_POINTER:
		c := C.g_value_get_pointer(v.Native())
		return unsafe.Pointer(c), nil

	case TYPE_OBJECT:
		c := C.g_value_get_object(v.Native())
		// TODO: need to try and return an actual pointer to the correct object type
		// this may require an additional cast()-like method for each module
		return ObjectNew(unsafe.Pointer(c)), nil

	case TYPE_VARIANT:
		return nil, errors.New("variant conversion not yet implemented")

	default:
		return nil, errors.New("type conversion not supported")
	}
}

// reflectGValue returns a Value but attempts to create one with a
// Pointer GType that points to a reflect.Value created from v,
// instead of checking the type of v for a pointer and setting the
// GValue with that.
func reflectGValue(v interface{}) (gvalue *Value, err error) {
	rv := reflect.ValueOf(v)
	val, err := ValueInit(TYPE_POINTER)
	if err != nil {
		return nil, err
	}
	val.SetPointer(uintptr(unsafe.Pointer(&rv)))
	return val, nil
}

// reflectGoValue returns a pointer to a reflect.Value created as
// a Pointer GValue with reflectGValue.
func (v *Value) reflectGoValue() *reflect.Value {
	rp := unsafe.Pointer(C.g_value_get_pointer(v.Native()))
	return (*reflect.Value)(rp)
}

// SetBool() is a wrapper around g_value_set_boolean().
func (v *Value) SetBool(val bool) {
	C.g_value_set_boolean(v.Native(), gbool(val))
}

// SetSChar() is a wrapper around g_value_set_schar().
func (v *Value) SetSChar(val int8) {
	C.g_value_set_schar(v.Native(), C.gint8(val))
}

// SetInt64() is a wrapper around g_value_set_int64().
func (v *Value) SetInt64(val int64) {
	C.g_value_set_int64(v.Native(), C.gint64(val))
}

// SetInt() is a wrapper around g_value_set_int().
func (v *Value) SetInt(val int) {
	C.g_value_set_int(v.Native(), C.gint(val))
}

// SetUChar() is a wrapper around g_value_set_uchar().
func (v *Value) SetUChar(val uint8) {
	C.g_value_set_uchar(v.Native(), C.guchar(val))
}

// SetUInt64() is a wrapper around g_value_set_uint64().
func (v *Value) SetUInt64(val uint64) {
	C.g_value_set_uint64(v.Native(), C.guint64(val))
}

// SetUInt() is a wrapper around g_value_set_uint().
func (v *Value) SetUInt(val uint) {
	C.g_value_set_uint(v.Native(), C.guint(val))
}

// SetFloat() is a wrapper around g_value_set_float().
func (v *Value) SetFloat(val float32) {
	C.g_value_set_float(v.Native(), C.gfloat(val))
}

// SetDouble() is a wrapper around g_value_set_double().
func (v *Value) SetDouble(val float64) {
	C.g_value_set_double(v.Native(), C.gdouble(val))
}

// SetString() is a wrapper around g_value_set_string().
func (v *Value) SetString(val string) {
	cstr := C.CString(val)
	defer C.free(unsafe.Pointer(cstr))
	C.g_value_set_string(v.Native(), (*C.gchar)(cstr))
}

// SetInstance() is a wrapper around g_value_set_instance().
func (v *Value) SetInstance(instance uintptr) {
	C.g_value_set_instance(v.Native(), C.gpointer(instance))
}

// SetPointer() is a wrapper around g_value_set_pointer().
func (v *Value) SetPointer(p uintptr) {
	C.g_value_set_pointer(v.Native(), C.gpointer(p))
}

// GetString() is a wrapper around g_value_get_string().  GetString()
// returns a non-nil error if g_value_get_string() returned a NULL
// pointer to distinguish between returning a NULL pointer and returning
// an empty string.
func (v *Value) GetString() (string, error) {
	c := C.g_value_get_string(v.Native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

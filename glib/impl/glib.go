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

// Package glib provides Go bindings for GLib 2.  Supports version 2.36
// and later.
package impl

// #cgo pkg-config: glib-2.0 gobject-2.0 gio-2.0
// #include <gio/gio.h>
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

	"github.com/gotk3/gotk3/glib"
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
	rf       reflect.Value
	userData reflect.Value
}

var (
	errNilPtr = errors.New("cgo returned unexpected nil pointer")

	closures = struct {
		sync.RWMutex
		m map[*C.GClosure]closureContext
	}{
		m: make(map[*C.GClosure]closureContext),
	}

	signals = make(map[glib.SignalHandle]*C.GClosure)
)

func init() {
	glib.USER_DIRECTORY_DESKTOP = C.G_USER_DIRECTORY_DESKTOP
	glib.USER_DIRECTORY_DOCUMENTS = C.G_USER_DIRECTORY_DOCUMENTS
	glib.USER_DIRECTORY_DOWNLOAD = C.G_USER_DIRECTORY_DOWNLOAD
	glib.USER_DIRECTORY_MUSIC = C.G_USER_DIRECTORY_MUSIC
	glib.USER_DIRECTORY_PICTURES = C.G_USER_DIRECTORY_PICTURES
	glib.USER_DIRECTORY_PUBLIC_SHARE = C.G_USER_DIRECTORY_PUBLIC_SHARE
	glib.USER_DIRECTORY_TEMPLATES = C.G_USER_DIRECTORY_TEMPLATES
	glib.USER_DIRECTORY_VIDEOS = C.G_USER_DIRECTORY_VIDEOS

	glib.USER_N_DIRECTORIES = C.G_USER_N_DIRECTORIES

	glib.APPLICATION_FLAGS_NONE = C.G_APPLICATION_FLAGS_NONE
	glib.APPLICATION_IS_SERVICE = C.G_APPLICATION_IS_SERVICE
	glib.APPLICATION_HANDLES_OPEN = C.G_APPLICATION_HANDLES_OPEN
	glib.APPLICATION_HANDLES_COMMAND_LINE = C.G_APPLICATION_HANDLES_COMMAND_LINE
	glib.APPLICATION_SEND_ENVIRONMENT = C.G_APPLICATION_SEND_ENVIRONMENT
	glib.APPLICATION_NON_UNIQUE = C.G_APPLICATION_NON_UNIQUE

	glib.TYPE_INVALID = C.G_TYPE_INVALID
	glib.TYPE_NONE = C.G_TYPE_NONE
	glib.TYPE_INTERFACE = C.G_TYPE_INTERFACE
	glib.TYPE_CHAR = C.G_TYPE_CHAR
	glib.TYPE_UCHAR = C.G_TYPE_UCHAR
	glib.TYPE_BOOLEAN = C.G_TYPE_BOOLEAN
	glib.TYPE_INT = C.G_TYPE_INT
	glib.TYPE_UINT = C.G_TYPE_UINT
	glib.TYPE_LONG = C.G_TYPE_LONG
	glib.TYPE_ULONG = C.G_TYPE_ULONG
	glib.TYPE_INT64 = C.G_TYPE_INT64
	glib.TYPE_UINT64 = C.G_TYPE_UINT64
	glib.TYPE_ENUM = C.G_TYPE_ENUM
	glib.TYPE_FLAGS = C.G_TYPE_FLAGS
	glib.TYPE_FLOAT = C.G_TYPE_FLOAT
	glib.TYPE_DOUBLE = C.G_TYPE_DOUBLE
	glib.TYPE_STRING = C.G_TYPE_STRING
	glib.TYPE_POINTER = C.G_TYPE_POINTER
	glib.TYPE_BOXED = C.G_TYPE_BOXED
	glib.TYPE_PARAM = C.G_TYPE_PARAM
	glib.TYPE_OBJECT = C.G_TYPE_OBJECT
	glib.TYPE_VARIANT = C.G_TYPE_VARIANT
}

// TypeName is a wrapper around g_type_name().
func TypeName(t glib.Type) string {
	return C.GoString((*C.char)(C.g_type_name(C.GType(t))))
}

// TypeDepth is a wrapper around g_type_depth().
func TypeDepth(t glib.Type) uint {
	return uint(C.g_type_depth(C.GType(t)))
}

// TypeParent is a wrapper around g_type_parent().
func TypeParent(t glib.Type) glib.Type {
	return glib.Type(C.g_type_parent(C.GType(t)))
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

	// Get the context associated with this callback closure.
	closures.RLock()
	cc := closures.m[closure]
	closures.RUnlock()

	// Get number of parameters passed in.  If user data was saved with the
	// closure context, increment the total number of parameters.
	nGLibParams := int(nParams)
	nTotalParams := nGLibParams
	if cc.userData.IsValid() {
		nTotalParams++
	}

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
		v := &value{&gValues[i]}
		val, err := v.GoValue()
		if err != nil {
			fmt.Fprintf(os.Stderr,
				"no suitable Go value for arg %d: %v\n", i, err)
			return
		}
		rv := reflect.ValueOf(val)
		args = append(args, rv.Convert(cc.rf.Type().In(i)))
	}

	// If non-nil user data was passed in and not all args have been set,
	// get and set the reflect.Value directly from the GValue.
	if cc.userData.IsValid() && len(args) < cap(args) {
		args = append(args, cc.userData.Convert(cc.rf.Type().In(nCbParams-1)))
	}

	// Call closure with args. If the callback returns one or more
	// values, save the GValue equivalent of the first.
	rv := cc.rf.Call(args)
	if retValue != nil && len(rv) > 0 {
		if g, err := GValue(rv[0].Interface()); err != nil {
			fmt.Fprintf(os.Stderr,
				"cannot save callback return value: %v", err)
		} else {
			*retValue = *g.native()
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

// IdleAdd adds an idle source to the default main event loop
// context.  After running once, the source func will be removed
// from the main event loop, unless f returns a single bool true.
//
// This function will cause a panic when f eventually runs if the
// types of args do not match those of f.
func IdleAdd(f interface{}, args ...interface{}) (glib.SourceHandle, error) {
	// f must be a func with no parameters.
	rf := reflect.ValueOf(f)
	if rf.Type().Kind() != reflect.Func {
		return 0, errors.New("f is not a function")
	}

	// Create an idle source func to be added to the main loop context.
	idleSrc := C.g_idle_source_new()
	if idleSrc == nil {
		return 0, errNilPtr
	}
	return sourceAttach(idleSrc, rf, args...)
}

// TimeoutAdd adds an timeout source to the default main event loop
// context.  After running once, the source func will be removed
// from the main event loop, unless f returns a single bool true.
//
// This function will cause a panic when f eventually runs if the
// types of args do not match those of f.
// timeout is in milliseconds
func TimeoutAdd(timeout uint, f interface{}, args ...interface{}) (glib.SourceHandle, error) {
	// f must be a func with no parameters.
	rf := reflect.ValueOf(f)
	if rf.Type().Kind() != reflect.Func {
		return 0, errors.New("f is not a function")
	}

	// Create a timeout source func to be added to the main loop context.
	timeoutSrc := C.g_timeout_source_new(C.guint(timeout))
	if timeoutSrc == nil {
		return 0, errNilPtr
	}

	return sourceAttach(timeoutSrc, rf, args...)
}

// sourceAttach attaches a source to the default main loop context.
func sourceAttach(src *C.struct__GSource, rf reflect.Value, args ...interface{}) (glib.SourceHandle, error) {
	if src == nil {
		return 0, errNilPtr
	}

	// rf must be a func with no parameters.
	if rf.Type().Kind() != reflect.Func {
		C.g_source_destroy(src)
		return 0, errors.New("rf is not a function")
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
		C.g_source_destroy(src)
	})

	// Remove closure context when closure is finalized.
	C._g_closure_add_finalize_notifier(closure)

	// Set closure to run as a callback when the idle source runs.
	C.g_source_set_closure(src, closure)

	// Attach the idle source func to the default main event loop
	// context.
	cid := C.g_source_attach(src, nil)
	return glib.SourceHandle(cid), nil
}

/*
 * Miscellaneous Utility Functions
 */

// GetUserSpecialDir is a wrapper around g_get_user_special_dir().  A
// non-nil error is returned in the case that g_get_user_special_dir()
// returns NULL to differentiate between NULL and an empty string.
func GetUserSpecialDir(directory glib.UserDirectory) (string, error) {
	c := C.g_get_user_special_dir(C.GUserDirectory(directory))
	if c == nil {
		return "", errNilPtr
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

// Native returns a pointer to the underlying GObject.
func (v *Object) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

// IsA is a wrapper around g_type_is_a().
func (v *Object) IsA(typ glib.Type) bool {
	return gobool(C.g_type_is_a(C.GType(v.TypeFromInstance()), C.GType(typ)))
}

// TypeFromInstance is a wrapper around g_type_from_instance().
func (v *Object) TypeFromInstance() glib.Type {
	c := C._g_type_from_instance(C.gpointer(unsafe.Pointer(v.native())))
	return glib.Type(c)
}

// ToGObject type converts an unsafe.Pointer as a native C GObject.
// This function is exported for visibility in other gotk3 packages and
// is not meant to be used by applications.
func ToGObject(p unsafe.Pointer) *C.GObject {
	return C.toGObject(p)
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

// Set is a wrapper around g_object_set().  However, unlike
// g_object_set(), this function only sets one name value pair.  Make
// multiple calls to this function to set multiple properties.
func (v *Object) Set(name string, value interface{}) error {
	return v.SetProperty(name, value)
	/*
		cstr := C.CString(name)
		defer C.free(unsafe.Pointer(cstr))

		if _, ok := value.(Object); ok {
			value = value.(Object).GObject
		}

		// Can't call g_object_set() as it uses a variable arg list, use a
		// wrapper instead
		var p unsafe.Pointer
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
			defer C.g_free(C.gpointer(unsafe.Pointer(cstr)))
			p = unsafe.Pointer(&cstr)

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
		if p == nil {
			return errors.New("Unable to perform type conversion")
		}
		C._g_object_set_one(C.gpointer(v.GObject), (*C.gchar)(cstr), p)
		return nil*/
}

// GetPropertyType returns the Type of a property of the underlying GObject.
// If the property is missing it will return TYPE_INVALID and an error.
func (v *Object) GetPropertyType(name string) (glib.Type, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	paramSpec := C.g_object_class_find_property(C._g_object_get_class(v.native()), (*C.gchar)(cstr))
	if paramSpec == nil {
		return glib.TYPE_INVALID, errors.New("couldn't find Property")
	}
	return glib.Type(paramSpec.value_type), nil
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

// pointerVal attempts to return an unsafe.Pointer for value.
// Not all types are understood, in which case a nil Pointer
// is returned.
/*func pointerVal(value interface{}) unsafe.Pointer {
	var p unsafe.Pointer
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
}*/

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
func (v *Object) HandlerBlock(handle glib.SignalHandle) {
	C.g_signal_handler_block(C.gpointer(v.GObject), C.gulong(handle))
}

// HandlerUnblock is a wrapper around g_signal_handler_unblock().
func (v *Object) HandlerUnblock(handle glib.SignalHandle) {
	C.g_signal_handler_unblock(C.gpointer(v.GObject), C.gulong(handle))
}

// HandlerDisconnect is a wrapper around g_signal_handler_disconnect().
func (v *Object) HandlerDisconnect(handle glib.SignalHandle) {
	C.g_signal_handler_disconnect(C.gpointer(v.GObject), C.gulong(handle))
	C.g_closure_invalidate(signals[handle])
	delete(closures.m, signals[handle])
	delete(signals, handle)
}

// Wrapper function for new objects with reference management.
func wrapObject(ptr unsafe.Pointer) *Object {
	obj := &Object{ToGObject(ptr)}

	if obj.IsFloating() {
		obj.RefSink()
	} else {
		obj.Ref()
	}

	runtime.SetFinalizer(obj, (*Object).Unref)
	return obj
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
type value struct {
	GValue *C.GValue
}

// native returns a pointer to the underlying GValue.
func (v *value) native() *C.GValue {
	return v.GValue
}

// Native returns a pointer to the underlying GValue.
func (v *value) Native() unsafe.Pointer {
	return unsafe.Pointer(v.native())
}

// ValueAlloc allocates a Value and sets a runtime finalizer to call
// g_value_unset() on the underlying GValue after leaving scope.
// ValueAlloc() returns a non-nil error if the allocation failed.
func ValueAlloc() (*value, error) {
	c := C._g_value_alloc()
	if c == nil {
		return nil, errNilPtr
	}

	v := &value{c}

	//An allocated GValue is not guaranteed to hold a value that can be unset
	//We need to double check before unsetting, to prevent:
	//`g_value_unset: assertion 'G_IS_VALUE (value)' failed`
	runtime.SetFinalizer(v, func(f *value) {
		if t, _, err := f.Type(); err != nil || t == glib.TYPE_INVALID || t == glib.TYPE_NONE {
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
func ValueInit(t glib.Type) (*value, error) {
	c := C._g_value_init(C.GType(t))
	if c == nil {
		return nil, errNilPtr
	}

	v := &value{c}

	runtime.SetFinalizer(v, (*value).unset)
	return v, nil
}

// ValueFromNative returns a type-asserted pointer to the Value.
func ValueFromNative(l unsafe.Pointer) *value {
	//TODO why it does not add finalizer to the value?
	return &value{(*C.GValue)(l)}
}

func (v *value) unset() {
	C.g_value_unset(v.native())
}

// Type is a wrapper around the G_VALUE_HOLDS_GTYPE() macro and
// the g_value_get_gtype() function.  GetType() returns TYPE_INVALID if v
// does not hold a Type, or otherwise returns the Type of v.
func (v *value) Type() (actual glib.Type, fundamental glib.Type, err error) {
	if !gobool(C._g_is_value(v.native())) {
		return actual, fundamental, errors.New("invalid GValue")
	}
	cActual := C._g_value_type(v.native())
	cFundamental := C._g_value_fundamental(cActual)
	return glib.Type(cActual), glib.Type(cFundamental), nil
}

// GValue converts a Go type to a comparable GValue.  GValue()
// returns a non-nil error if the conversion was unsuccessful.
func GValue(v interface{}) (gvalue *value, err error) {
	if v == nil {
		val, err := ValueInit(glib.TYPE_POINTER)
		if err != nil {
			return nil, err
		}
		val.SetPointer(uintptr(unsafe.Pointer(nil)))
		return val, nil
	}

	switch e := v.(type) {
	case bool:
		val, err := ValueInit(glib.TYPE_BOOLEAN)
		if err != nil {
			return nil, err
		}
		val.SetBool(e)
		return val, nil

	case int8:
		val, err := ValueInit(glib.TYPE_CHAR)
		if err != nil {
			return nil, err
		}
		val.SetSChar(e)
		return val, nil

	case int64:
		val, err := ValueInit(glib.TYPE_INT64)
		if err != nil {
			return nil, err
		}
		val.SetInt64(e)
		return val, nil

	case int:
		val, err := ValueInit(glib.TYPE_INT)
		if err != nil {
			return nil, err
		}
		val.SetInt(e)
		return val, nil

	case uint8:
		val, err := ValueInit(glib.TYPE_UCHAR)
		if err != nil {
			return nil, err
		}
		val.SetUChar(e)
		return val, nil

	case uint64:
		val, err := ValueInit(glib.TYPE_UINT64)
		if err != nil {
			return nil, err
		}
		val.SetUInt64(e)
		return val, nil

	case uint:
		val, err := ValueInit(glib.TYPE_UINT)
		if err != nil {
			return nil, err
		}
		val.SetUInt(e)
		return val, nil

	case float32:
		val, err := ValueInit(glib.TYPE_FLOAT)
		if err != nil {
			return nil, err
		}
		val.SetFloat(e)
		return val, nil

	case float64:
		val, err := ValueInit(glib.TYPE_DOUBLE)
		if err != nil {
			return nil, err
		}
		val.SetDouble(e)
		return val, nil

	case string:
		val, err := ValueInit(glib.TYPE_STRING)
		if err != nil {
			return nil, err
		}
		val.SetString(e)
		return val, nil

	case *Object:
		val, err := ValueInit(glib.TYPE_OBJECT)
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
			val, err := ValueInit(glib.TYPE_CHAR)
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
			val, err := ValueInit(glib.TYPE_INT64)
			if err != nil {
				return nil, err
			}
			val.SetInt64(rval.Int())
			return val, nil

		case reflect.Int:
			val, err := ValueInit(glib.TYPE_INT)
			if err != nil {
				return nil, err
			}
			val.SetInt(int(rval.Int()))
			return val, nil

		case reflect.Uintptr, reflect.Ptr:
			val, err := ValueInit(glib.TYPE_POINTER)
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
// appropiate Go type.  The uintptr parameter is a *C.GValue.
type GValueMarshaler func(uintptr) (interface{}, error)

// TypeMarshaler represents an actual type and it's associated marshaler.
type TypeMarshaler struct {
	T glib.Type
	F GValueMarshaler
}

// RegisterGValueMarshalers adds marshalers for several types to the
// internal marshalers map.  Once registered, calling GoValue on any
// Value witha registered type will return the data returned by the
// marshaler.
func RegisterGValueMarshalers(tm []TypeMarshaler) {
	gValueMarshalers.register(tm)
}

type marshalMap map[glib.Type]GValueMarshaler

// gValueMarshalers is a map of Glib types to functions to marshal a
// GValue to a native Go type.
var gValueMarshalers = marshalMap{
	glib.TYPE_INVALID:   marshalInvalid,
	glib.TYPE_NONE:      marshalNone,
	glib.TYPE_INTERFACE: marshalInterface,
	glib.TYPE_CHAR:      marshalChar,
	glib.TYPE_UCHAR:     marshalUchar,
	glib.TYPE_BOOLEAN:   marshalBoolean,
	glib.TYPE_INT:       marshalInt,
	glib.TYPE_LONG:      marshalLong,
	glib.TYPE_ENUM:      marshalEnum,
	glib.TYPE_INT64:     marshalInt64,
	glib.TYPE_UINT:      marshalUint,
	glib.TYPE_ULONG:     marshalUlong,
	glib.TYPE_FLAGS:     marshalFlags,
	glib.TYPE_UINT64:    marshalUint64,
	glib.TYPE_FLOAT:     marshalFloat,
	glib.TYPE_DOUBLE:    marshalDouble,
	glib.TYPE_STRING:    marshalString,
	glib.TYPE_POINTER:   marshalPointer,
	glib.TYPE_BOXED:     marshalBoxed,
	glib.TYPE_OBJECT:    marshalObject,
	glib.TYPE_VARIANT:   marshalVariant,
}

func (m marshalMap) register(tm []TypeMarshaler) {
	for i := range tm {
		m[tm[i].T] = tm[i].F
	}
}

func (m marshalMap) lookup(v *value) (GValueMarshaler, error) {
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
	return nil, errors.New("variant conversion not yet implemented")
}

// GoValue converts a Value to comparable Go type.  GoValue()
// returns a non-nil error if the conversion was unsuccessful.  The
// returned interface{} must be type asserted as the actual Go
// representation of the Value.
//
// This function is a wrapper around the many g_value_get_*()
// functions, depending on the type of the Value.
func (v *value) GoValue() (interface{}, error) {
	f, err := gValueMarshalers.lookup(v)
	if err != nil {
		return nil, err
	}

	//No need to add finalizer because it is already done by ValueAlloc and ValueInit
	rv, err := f(uintptr(unsafe.Pointer(v.native())))
	return rv, err
}

// SetBool is a wrapper around g_value_set_boolean().
func (v *value) SetBool(val bool) {
	C.g_value_set_boolean(v.native(), gbool(val))
}

// SetSChar is a wrapper around g_value_set_schar().
func (v *value) SetSChar(val int8) {
	C.g_value_set_schar(v.native(), C.gint8(val))
}

// SetInt64 is a wrapper around g_value_set_int64().
func (v *value) SetInt64(val int64) {
	C.g_value_set_int64(v.native(), C.gint64(val))
}

// SetInt is a wrapper around g_value_set_int().
func (v *value) SetInt(val int) {
	C.g_value_set_int(v.native(), C.gint(val))
}

// SetUChar is a wrapper around g_value_set_uchar().
func (v *value) SetUChar(val uint8) {
	C.g_value_set_uchar(v.native(), C.guchar(val))
}

// SetUInt64 is a wrapper around g_value_set_uint64().
func (v *value) SetUInt64(val uint64) {
	C.g_value_set_uint64(v.native(), C.guint64(val))
}

// SetUInt is a wrapper around g_value_set_uint().
func (v *value) SetUInt(val uint) {
	C.g_value_set_uint(v.native(), C.guint(val))
}

// SetFloat is a wrapper around g_value_set_float().
func (v *value) SetFloat(val float32) {
	C.g_value_set_float(v.native(), C.gfloat(val))
}

// SetDouble is a wrapper around g_value_set_double().
func (v *value) SetDouble(val float64) {
	C.g_value_set_double(v.native(), C.gdouble(val))
}

// SetString is a wrapper around g_value_set_string().
func (v *value) SetString(val string) {
	cstr := C.CString(val)
	defer C.free(unsafe.Pointer(cstr))
	C.g_value_set_string(v.native(), (*C.gchar)(cstr))
}

// SetInstance is a wrapper around g_value_set_instance().
func (v *value) SetInstance(instance uintptr) {
	C.g_value_set_instance(v.native(), C.gpointer(instance))
}

// SetPointer is a wrapper around g_value_set_pointer().
func (v *value) SetPointer(p uintptr) {
	C.g_value_set_pointer(v.native(), C.gpointer(p))
}

// GetPointer is a wrapper around g_value_get_pointer().
func (v *value) GetPointer() unsafe.Pointer {
	return unsafe.Pointer(C.g_value_get_pointer(v.native()))
}

// GetString is a wrapper around g_value_get_string().  GetString()
// returns a non-nil error if g_value_get_string() returned a NULL
// pointer to distinguish between returning a NULL pointer and returning
// an empty string.
func (v *value) GetString() (string, error) {
	c := C.g_value_get_string(v.native())
	if c == nil {
		return "", errNilPtr
	}
	return C.GoString((*C.char)(c)), nil
}

type signal struct {
	name     string
	signalId C.guint
}

func SignalNew(s string) (*signal, error) {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))

	signalId := C._g_signal_new((*C.gchar)(cstr))

	if signalId == 0 {
		return nil, fmt.Errorf("invalid signal name: %s", s)
	}

	return &signal{
		name:     s,
		signalId: signalId,
	}, nil
}

func (s *signal) String() string {
	return s.name
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
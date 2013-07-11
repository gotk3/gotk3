/*
 * Copyright (c) 2013 Conformal Systems <info@conformal.com>
 *
 * This file originated from: http://opensource.conformal.com/
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

/*
Go bindings for GLib 2.  Supports version 2.36 and later.
*/
package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"unsafe"
)

var (
	callbackContexts = struct {
		sync.RWMutex
		s []*CallbackContext
	}{}
	idleFnContexts = struct {
		sync.RWMutex
		s []*idleFnContext
	}{}
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

type Type uint

const _TYPE_FUNDAMENTAL_SHIFT = 2

const (
	TYPE_INVALID Type = iota << _TYPE_FUNDAMENTAL_SHIFT
	TYPE_NONE
	TYPE_INTERFACE
	TYPE_CHAR
	TYPE_UCHAR
	TYPE_BOOLEAN
	TYPE_INT
	TYPE_UINT
	TYPE_LONG
	TYPE_ULONG
	TYPE_INT64
	TYPE_UINT64
	TYPE_ENUM
	TYPE_FLAGS
	TYPE_FLOAT
	TYPE_DOUBLE
	TYPE_STRING
	TYPE_POINTER
	TYPE_BOXED
	TYPE_PARAM
	TYPE_OBJECT
	TYPE_VARIANT
)

/*
 * Events
 */

type CallbackContext struct {
	f      interface{}
	cbi    unsafe.Pointer
	target reflect.Value
	data   reflect.Value
}

type CallbackArg uintptr

func (c *CallbackContext) Target() interface{} {
	return c.target.Interface()
}

func (c *CallbackContext) Data() interface{} {
	return c.data.Interface()
}

func (c *CallbackContext) Arg(n int) CallbackArg {
	return CallbackArg(C.cbinfo_get_arg((*C.cbinfo)(c.cbi), C.int(n)))
}

func (c CallbackArg) String() string {
	return C.GoString((*C.char)(unsafe.Pointer(c)))
}

func (c CallbackArg) Int() int {
	return int(C.int(C.uintptr_t(c)))
}

func (c CallbackArg) UInt() uint {
	return uint(C.uint(C.uintptr_t(c)))
}

//export _go_glib_callback
func _go_glib_callback(cbi *C.cbinfo) {
	callbackContexts.RLock()
	ctx := callbackContexts.s[int(cbi.func_n)]
	rf := reflect.ValueOf(ctx.f)
	t := rf.Type()
	fargs := make([]reflect.Value, t.NumIn())
	if len(fargs) > 0 {
		fargs[0] = reflect.ValueOf(ctx)
	}
	callbackContexts.RUnlock()
	ret := rf.Call(fargs)
	if len(ret) > 0 {
		bret, _ := ret[0].Interface().(bool)
		cbi.ret = gbool(bret)
	}
}

/*
 * Main event loop
 */

type idleFnContext struct {
	f    interface{}
	args []reflect.Value
	idl  *C.idleinfo
}

func IdleAdd(f interface{}, datas ...interface{}) (uint, error) {
	rf := reflect.ValueOf(f)
	if rf.Kind() != reflect.Func {
		return 0, errors.New("f is not a function")
	}
	t := rf.Type()
	if t.NumIn() != len(datas) {
		return 0, errors.New("Number of arguments do not match")
	}

	var vals []reflect.Value
	for i := range datas {
		ntharg := t.In(i)
		val := reflect.ValueOf(datas[i])
		if ntharg.Kind() != val.Kind() {
			s := fmt.Sprint("Types of arg", i, "do not match")
			return 0, errors.New(s)
		}
		vals = append(vals, val)
	}

	ctx := &idleFnContext{}
	ctx.f = f
	ctx.args = vals

	idleFnContexts.Lock()
	idleFnContexts.s = append(idleFnContexts.s, ctx)
	idleFnContexts.Unlock()

	idleFnContexts.RLock()
	nIdleFns := len(idleFnContexts.s)
	idleFnContexts.RUnlock()
	idl := C._g_idle_add(C.int(nIdleFns) - 1)

	ctx.idl = idl

	return uint(idl.id), nil
}

//export _go_glib_idle_fn
func _go_glib_idle_fn(idl *C.idleinfo) {
	idleFnContexts.RLock()
	ctx := idleFnContexts.s[int(idl.func_n)]
	idleFnContexts.RUnlock()
	rf := reflect.ValueOf(ctx.f)
	rv := rf.Call(ctx.args)
	if len(rv) == 1 {
		if rv[0].Kind() == reflect.Bool {
			idl.ret = gbool(rv[0].Bool())
			return
		}
	}
	idl.ret = gbool(false)
}

//export _go_nil_unused_idle_ctx
func _go_nil_unused_idle_ctx(n C.int) {
	idleFnContexts.Lock()
	idleFnContexts.s[int(n)] = nil
	idleFnContexts.Unlock()
}

/*
 * GObject
 */

type IObject interface {
	toGObject() *C.GObject
}

type Object struct {
	GObject *C.GObject
}

func (v *Object) Native() *C.GObject {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGObject(p)
}

func (v *Object) toGObject() *C.GObject {
	if v == nil {
		return nil
	}
	return v.Native()
}

func (v *Object) typeFromInstance() Type {
	c := C._g_type_from_instance(C.gpointer(unsafe.Pointer(v.Native())))
	return Type(c)
}

func ToGObject(p unsafe.Pointer) *C.GObject {
	return C.toGObject(p)
}

func (v *Object) Ref() {
	C.g_object_ref(C.gpointer(v.GObject))
}

func (v *Object) Unref() {
	C.g_object_unref(C.gpointer(v.GObject))
}

func (v *Object) RefSink() {
	C.g_object_ref_sink(C.gpointer(v.GObject))
}

func (v *Object) IsFloating() bool {
	c := C.g_object_is_floating(C.gpointer(v.GObject))
	return gobool(c)
}

func (v *Object) ForceFloating() {
	C.g_object_force_floating(v.GObject)
}

func (v *Object) StopEmission(s string) {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.g_signal_stop_emission_by_name((C.gpointer)(v.GObject),
		(*C.gchar)(cstr))
}

func (v *Object) connectCtx(ctx *CallbackContext, s string) int {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	callbackContexts.RLock()
	nCbCtxs := len(callbackContexts.s)
	callbackContexts.RUnlock()
	ctx.cbi = unsafe.Pointer(C._g_signal_connect(unsafe.Pointer(v.GObject),
		(*C.gchar)(cstr), C.int(nCbCtxs)))
	callbackContexts.Lock()
	callbackContexts.s = append(callbackContexts.s, ctx)
	callbackContexts.Unlock()
	return nCbCtxs
}

func (v *Object) Connect(s string, f interface{}) int {
	ctx := &CallbackContext{f, nil, reflect.ValueOf(v),
		reflect.ValueOf(nil)}
	return v.connectCtx(ctx, s)
}

func (v *Object) ConnectWithData(s string, f interface{}, data interface{}) int {
	ctx := &CallbackContext{f, nil, reflect.ValueOf(v),
		reflect.ValueOf(data)}
	return v.connectCtx(ctx, s)
}

// Unlike g_object_set(), this function only sets one name value pair.
// Make multiple calls to set multiple properties.
func (v *Object) Set(name string, value interface{}) error {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	if _, ok := value.(Object); ok {
		value = value.(Object).GObject
	}

	var p unsafe.Pointer = nil
	switch value.(type) {
	case bool:
		c := gbool(value.(bool))
		p = unsafe.Pointer(&c)
	case int8:
		c := C.gint8(value.(int8))
		p = unsafe.Pointer(&c)
	case int16:
		c := C.gint16(value.(int16))
		p = unsafe.Pointer(&c)
	case int32:
		c := C.gint32(value.(int32))
		p = unsafe.Pointer(&c)
	case int64:
		c := C.gint64(value.(int64))
		p = unsafe.Pointer(&c)
	case int:
		c := C.gint(value.(int))
		p = unsafe.Pointer(&c)
	case uint8:
		c := C.guchar(value.(uint8))
		p = unsafe.Pointer(&c)
	case uint16:
		c := C.guint16(value.(uint16))
		p = unsafe.Pointer(&c)
	case uint32:
		c := C.guint32(value.(uint32))
		p = unsafe.Pointer(&c)
	case uint64:
		c := C.guint64(value.(uint64))
		p = unsafe.Pointer(&c)
	case uint:
		c := C.guint(value.(uint))
		p = unsafe.Pointer(&c)
	case uintptr:
		p = unsafe.Pointer(C.gpointer(value.(uintptr)))
	case float32:
		c := C.gfloat(value.(float32))
		p = unsafe.Pointer(&c)
	case float64:
		c := C.gdouble(value.(float64))
		p = unsafe.Pointer(&c)
	case string:
		cstr := C.CString(value.(string))
		defer C.free(unsafe.Pointer(cstr))
		p = unsafe.Pointer(cstr)
	default:
		if pv, ok := value.(unsafe.Pointer); ok {
			p = pv
		} else {
			// Constants with separate types are not type asserted
			// above, so do a runtime check here instead.
			val := reflect.ValueOf(value)
			switch val.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16,
				reflect.Int32, reflect.Int64:
				c := C.int(val.Int())
				p = unsafe.Pointer(&c)
			case reflect.Uintptr:
				p = unsafe.Pointer(C.gpointer(val.Pointer()))
			}
		}
	}
	// Can't call g_object_set() as it uses a variable arg list, use a
	// wrapper instead
	if p != nil {
		C._g_object_set_one(C.gpointer(v.GObject), (*C.gchar)(cstr), p)
		return nil
	} else {
		return errors.New("Unable to perform type conversion")
	}
}

/*
 * GObject Signals
 */

// Emit() emits the signal specified by the string s to an Object.
// Arguments to callback functions connected to this signal must be
// specified in args.  Emit() returns an interface{} which must be type
// asserted as the Go equivalent type to the return value for native C
// callback.
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

	t := v.typeFromInstance()
	id := C.g_signal_lookup((*C.gchar)(cstr), C.GType(t))

	ret, err := ValueAlloc()
	if err != nil {
		return nil, errors.New("Error creating Value for return value")
	}
	C.g_signal_emitv(valv, id, C.GQuark(0), ret.Native())

	return ret.GoValue()
}

func (v *Object) HandlerBlock(callID int) {
	callbackContexts.RLock()
	id := C.cbinfo_get_id((*C.cbinfo)(callbackContexts.s[callID].cbi))
	callbackContexts.RUnlock()
	C.g_signal_handler_block((C.gpointer)(v.GObject), id)
}

func (v *Object) HandlerUnblock(callID int) {
	callbackContexts.RLock()
	id := C.cbinfo_get_id((*C.cbinfo)(callbackContexts.s[callID].cbi))
	callbackContexts.RUnlock()
	C.g_signal_handler_unblock((C.gpointer)(v.GObject), id)
}

func (v *Object) HandlerDisconnect(callID int) {
	callbackContexts.RLock()
	id := C.cbinfo_get_id((*C.cbinfo)(callbackContexts.s[callID].cbi))
	callbackContexts.RUnlock()
	C.g_signal_handler_disconnect((C.gpointer)(v.GObject), id)
}

/*
 * GInitiallyUnowned
 */

type InitiallyUnowned struct {
	*Object
}

/*
 * GValue
 */

// Don't allocate Values on the stack or heap manually as they may not
// be properly unset when going out of scope. Instead, use ValueAlloc(),
// which will set the runtime finalizer to unset the Value.
type Value struct {
	GValue C.GValue
}

func (v *Value) Native() *C.GValue {
	return &v.GValue
}

func ValueAlloc() (*Value, error) {
	c := C._g_value_alloc()
	if c == nil {
		return nil, nilPtrErr
	}
	v := &Value{*c}
	runtime.SetFinalizer(v, (*Value).unset)
	return v, nil
}

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

func (v *Value) GetType() Type {
	c := C._g_value_holds_gtype(C.gpointer(unsafe.Pointer(v.Native())))
	if gobool(c) {
		c := C.g_value_get_gtype(v.Native())
		return Type(c)
	}
	return TYPE_INVALID
}

// Converts a Go type to a comparable GValue.
func GValue(v interface{}) (gvalue *Value, err error) {
	if v == nil {
		val, err := ValueInit(TYPE_POINTER)
		if err != nil {
			return nil, err
		}
		val.SetPointer(uintptr(0)) // technically not portable
		return val, nil
	}

	switch v.(type) {
	case bool:
		val, err := ValueInit(TYPE_BOOLEAN)
		if err != nil {
			return nil, err
		}
		val.SetBool(v.(bool))
		return val, nil
	case int8:
		val, err := ValueInit(TYPE_CHAR)
		if err != nil {
			return nil, err
		}
		val.SetSChar(v.(int8))
	case int64:
		val, err := ValueInit(TYPE_INT64)
		if err != nil {
			return nil, err
		}
		val.SetInt64(v.(int64))
		return val, nil
	case int:
		val, err := ValueInit(TYPE_INT)
		if err != nil {
			return nil, err
		}
		val.SetInt(v.(int))
		return val, nil
	case uint8:
		val, err := ValueInit(TYPE_UCHAR)
		if err != nil {
			return nil, err
		}
		val.SetUChar(v.(uint8))
	case uint64:
		val, err := ValueInit(TYPE_UINT64)
		if err != nil {
			return nil, err
		}
		val.SetUInt64(v.(uint64))
	case uint:
		val, err := ValueInit(TYPE_UINT)
		if err != nil {
			return nil, err
		}
		val.SetUInt(v.(uint))
	case float32:
		val, err := ValueInit(TYPE_FLOAT)
		if err != nil {
			return nil, err
		}
		val.SetFloat(v.(float32))
	case float64:
		val, err := ValueInit(TYPE_DOUBLE)
		if err != nil {
			return nil, err
		}
		val.SetDouble(v.(float64))
	case string:
		val, err := ValueInit(TYPE_STRING)
		if err != nil {
			return nil, err
		}
		val.SetString(v.(string))
		return val, nil
	default:
		if obj, ok := v.(*Object); ok {
			val, err := ValueInit(TYPE_OBJECT)
			if err != nil {
				return nil, err
			}
			val.SetInstance(uintptr(unsafe.Pointer(obj.GObject)))
			return val, nil
		}

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
		case reflect.Uintptr:
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

// Converts a GValue to comparable Go type
func (v *Value) GoValue() (interface{}, error) {
	switch v.GetType() {
	case TYPE_INVALID:
		return nil, errors.New("Invalid type")
	case TYPE_NONE:
		return nil, nil
	case TYPE_BOOLEAN:
		c := C.g_value_get_boolean(v.Native())
		return gobool(c), nil
	case TYPE_CHAR:
		c := C.g_value_get_schar(v.Native())
		return int8(c), nil
	case TYPE_UCHAR:
		c := C.g_value_get_uchar(v.Native())
		return uint8(c), nil
	case TYPE_INT64:
		c := C.g_value_get_int64(v.Native())
		return int64(c), nil
	case TYPE_INT:
		c := C.g_value_get_int(v.Native())
		return int(c), nil
	case TYPE_UINT64:
		c := C.g_value_get_uint64(v.Native())
		return uint64(c), nil
	case TYPE_UINT:
		c := C.g_value_get_uint(v.Native())
		return uint(c), nil
	case TYPE_FLOAT:
		c := C.g_value_get_float(v.Native())
		return float32(c), nil
	case TYPE_DOUBLE:
		c := C.g_value_get_double(v.Native())
		return float64(c), nil
	case TYPE_STRING:
		c := C.g_value_get_string(v.Native())
		return C.GoString((*C.char)(c)), nil
	default:
		return nil, errors.New("Type conversion not supported")
	}
}

func (v *Value) SetBool(val bool) {
	C.g_value_set_boolean(v.Native(), gbool(val))
}

func (v *Value) SetSChar(val int8) {
	C.g_value_set_schar(v.Native(), C.gint8(val))
}

func (v *Value) SetInt64(val int64) {
	C.g_value_set_int64(v.Native(), C.gint64(val))
}

func (v *Value) SetInt(val int) {
	C.g_value_set_int(v.Native(), C.gint(val))
}

func (v *Value) SetUChar(val uint8) {
	C.g_value_set_uchar(v.Native(), C.guchar(val))
}

func (v *Value) SetUInt64(val uint64) {
	C.g_value_set_uint64(v.Native(), C.guint64(val))
}

func (v *Value) SetUInt(val uint) {
	C.g_value_set_uint(v.Native(), C.guint(val))
}

func (v *Value) SetFloat(val float32) {
	C.g_value_set_float(v.Native(), C.gfloat(val))
}

func (v *Value) SetDouble(val float64) {
	C.g_value_set_double(v.Native(), C.gdouble(val))
}

func (v *Value) SetString(val string) {
	cstr := C.CString(val)
	defer C.free(unsafe.Pointer(cstr))
	C.g_value_set_string(v.Native(), (*C.gchar)(cstr))
}

func (v *Value) SetInstance(instance uintptr) {
	C.g_value_set_instance(v.Native(), C.gpointer(instance))
}

func (v *Value) SetPointer(p uintptr) {
	C.g_value_set_pointer(v.Native(), C.gpointer(p))
}

func (v *Value) GetString() (string, error) {
	c := C.g_value_get_string(v.Native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

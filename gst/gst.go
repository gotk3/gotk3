package gst

// #cgo pkg-config: gstreamer-1.0
// #include <gst/gst.h>
// #include "gst.go.h"
import "C"

import (
	"unsafe"
	"errors"
	"runtime"
	
	"github.com/MovingtoMars/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		//{glib.Type(C.gtk_align_get_type()), marshalAlign},

		// Objects/Interfaces

		// Boxed
	}
	glib.RegisterGValueMarshalers(tm)
}

/*
Init() is a wrapper around gst_init() and must be called before any
other gstreamer calls and is used to initialize everything necessary.

In addition to setting up gstreamer for usage, a pointer to a slice of
strings may be passed in to parse standard GTK command line arguments.
args will be modified to remove any flags that were handled.
Alternatively, nil may be passed in to not perform any command line
parsing.
*/
func Init(args *[]string) {
	if args != nil {
		argc := C.int(len(*args))
		argv := make([]*C.char, argc)
		for i, arg := range *args {
			argv[i] = C.CString(arg)
		}
		C.gst_init((*C.int)(unsafe.Pointer(&argc)),
			(***C.char)(unsafe.Pointer(&argv)))
		unhandled := make([]string, argc)
		for i := 0; i < int(argc); i++ {
			unhandled[i] = C.GoString(argv[i])
			C.free(unsafe.Pointer(argv[i]))
		}
		*args = unhandled
	} else {
		C.gst_init(nil, nil)
	}
}

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

// Unref is a wrapper around gst_object_unref().
func (v *Object) Unref() {
	C.gst_object_unref(C.gpointer(v.toObject()))
}

// RefSink is a wrapper around g_object_ref_sink().
func (v *Object) RefSink() {
	C.gst_object_ref_sink(C.gpointer(v.toObject()))
}

/*
 * Constants
 */

const (
	// Infinite timeout (unsigned representation of -1)
	CLOCK_TIME_NONE uint64 = 18446744073709551615
)

// State is a representation of GstState
type State int

const (
	STATE_VOID_PENDING State = C.GST_STATE_VOID_PENDING
	STATE_NULL State = C.GST_STATE_NULL
	STATE_READY State = C.GST_STATE_READY
	STATE_PAUSED State = C.GST_STATE_PAUSED
	STATE_PLAYING State = C.GST_STATE_PLAYING
)

// StateChangeReturn is a representation of GstStateChangeReturn
type StateChangeReturn int

const (
	STATE_CHANGE_FAILURE StateChangeReturn = C.GST_STATE_CHANGE_FAILURE
	STATE_CHANGE_SUCCESS StateChangeReturn = C.GST_STATE_CHANGE_SUCCESS
	STATE_CHANGE_ASYNC StateChangeReturn = C.GST_STATE_CHANGE_ASYNC
	STATE_CHANGE_NO_PREROLL StateChangeReturn = C.GST_STATE_CHANGE_NO_PREROLL
)

/*
 * GstElement
 */

type Element struct {
	Object
}

// native returns a pointer to the underlying GstElement.
func (v *Element) native() *C.GstElement {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGstElement(p)
}

func marshalElement(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapElement(obj), nil
}

func wrapElement(obj *glib.Object) *Element {
	return &Element{Object{glib.InitiallyUnowned{obj}}}
}

// GetState() is a wrapper around gst_element_get_state().
func (v *Element) GetState(timeout uint64) (state, pending State, change StateChangeReturn) {
	var cstate, cpending C.GstState
	c := C.gst_element_get_state(v.native(), &cstate, &cpending, C.GstClockTime(timeout))
	return State(cstate), State(cpending), StateChangeReturn(c)
}

// SetState() is a wrapper around gst_element_set_state().
func (v *Element) SetState(state State) StateChangeReturn {
	c := C.gst_element_set_state(v.native(), C.GstState(state))
	return StateChangeReturn(c)
}

/*
 * GstElementFactory
 */

// ElementFactoryMake() is a wrapper around gst_element_factory_make().
func ElementFactoryMake(factoryName, name string) (*Element, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cfactoryName := C.CString(factoryName)
	defer C.free(unsafe.Pointer(cfactoryName))
	c := C.gst_element_factory_make((*C.gchar)(cfactoryName), (*C.gchar)(cname))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapElement(obj)
	b.RefSink()
	runtime.SetFinalizer(&b.Object, (*Object).Unref)
	return b, nil
}

/*
 * GstObject
 */

// Object is a representation of gst's GstObject.
type Object struct {
	glib.InitiallyUnowned
}

// native returns a pointer to the underlying GstObject.
func (v *Object) native() *C.GstObject {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGstObject(p)
}

func (v *Object) toObject() *C.GstObject {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalObject(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapObject(obj), nil
}

func wrapObject(obj *glib.Object) *Object {
	return &Object{glib.InitiallyUnowned{obj}}
}

//gst_object_get_name (GstObject *object)
func (v *Object) GetName() string {
	c := C.gst_object_get_name(v.native())
	defer C.g_free(C.gpointer(c))
	return C.GoString((*C.char)(c))
}

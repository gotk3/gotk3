// Go bindings for gstreamer.  Supports version 1.0 and later.
//
// Functions use the same names as the native C function calls, but use
// CamelCase.  In cases where native gstreamer uses pointers to values to
// simulate multiple return values, Go's native multiple return values
// are used instead.  Whenever a native gstreamer call could return an
// unexpected NULL pointer, an additonal error is returned in the Go
// binding.
//
// gstreamers's C API documentation can be very useful for understanding how the
// functions in this package work and what each type is for.  This
// documentation can be found at http://gstreamer.freedesktop.org/data/doc/gstreamer/head/gstreamer/html/.
//
// In addition to Go versions of the C gstreamer functions, every struct type
// includes a method named Native (either by direct implementation, or
// by means of struct embedding).  These methods return a uintptr of the
// native C object the binding type represents.  These pointers may be
// type switched to a native C pointer using unsafe and used with cgo
// function calls outside this package.
//
// Memory management is handled in proper Go fashion, using runtime
// finalizers to properly free memory when it is no longer needed.  Each
// time a Go type is created with a pointer to a GObject, a reference is
// added for Go, sinking the floating reference when necessary.  After
// going out of scope and the next time Go's garbage collector is run, a
// finalizer is run to remove Go's reference to the GObject.  When this
// reference count hits zero (when neither Go nor gstreamer holds ownership)
// the object will be freed internally by gstreamer.
package gst

// #cgo pkg-config: gstreamer-1.0
// #include <gst/gst.h>
// #include "gst.go.h"
import "C"

import (
	"errors"
	"runtime"
	"unsafe"

	"github.com/MovingtoMars/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.gst_format_get_type()), marshalFormat},
		{glib.Type(C.gst_message_type_get_type()), marshalMessageType},
		{glib.Type(C.gst_state_get_type()), marshalState},
		{glib.Type(C.gst_state_change_return_get_type()), marshalStateChangeReturn},

		// Objects/Interfaces
		{glib.Type(C.gst_bus_get_type()), marshalBus},
		{glib.Type(C.gst_element_get_type()), marshalElement},
		{glib.Type(C.gst_object_get_type()), marshalObject},

		// Boxed
		{glib.Type(C.gst_message_get_type()), marshalMessage},
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

// Format is a representation of GstFormat.
type Format int

const (
	FORMAT_UNDEFINED Format = C.GST_FORMAT_UNDEFINED
	FORMAT_DEFAULT   Format = C.GST_FORMAT_DEFAULT
	FORMAT_BYTES     Format = C.GST_FORMAT_BYTES
	FORMAT_TIME      Format = C.GST_FORMAT_TIME
	FORMAT_BUFFERS   Format = C.GST_FORMAT_BUFFERS
	FORMAT_PERCENT   Format = C.GST_FORMAT_PERCENT
)

func marshalFormat(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Format(c), nil
}

// MessageType is a representation of GstMessageType.
type MessageType int

const (
	MESSAGE_UNKNOWN          MessageType = C.GST_MESSAGE_UNKNOWN
	MESSAGE_EOS              MessageType = C.GST_MESSAGE_EOS
	MESSAGE_ERROR            MessageType = C.GST_MESSAGE_ERROR
	MESSAGE_WARNING          MessageType = C.GST_MESSAGE_WARNING
	MESSAGE_INFO             MessageType = C.GST_MESSAGE_INFO
	MESSAGE_TAG              MessageType = C.GST_MESSAGE_TAG
	MESSAGE_BUFFERING        MessageType = C.GST_MESSAGE_BUFFERING
	MESSAGE_STATE_CHANGED    MessageType = C.GST_MESSAGE_STATE_CHANGED
	MESSAGE_STATE_DIRTY      MessageType = C.GST_MESSAGE_STATE_DIRTY
	MESSAGE_STEP_DONE        MessageType = C.GST_MESSAGE_STEP_DONE
	MESSAGE_CLOCK_LOST       MessageType = C.GST_MESSAGE_CLOCK_LOST
	MESSAGE_NEW_CLOCK        MessageType = C.GST_MESSAGE_NEW_CLOCK
	MESSAGE_STRUCTURE_CHANGE MessageType = C.GST_MESSAGE_STRUCTURE_CHANGE
	MESSAGE_STREAM_STATUS    MessageType = C.GST_MESSAGE_STREAM_STATUS
	MESSAGE_APPLICATION      MessageType = C.GST_MESSAGE_APPLICATION
	MESSAGE_ELEMENT          MessageType = C.GST_MESSAGE_ELEMENT
	MESSAGE_SEGMENT_START    MessageType = C.GST_MESSAGE_SEGMENT_START
	MESSAGE_SEGMENT_DONE     MessageType = C.GST_MESSAGE_SEGMENT_DONE
	MESSAGE_DURATION_CHANGED MessageType = C.GST_MESSAGE_DURATION_CHANGED
	MESSAGE_LATENCY          MessageType = C.GST_MESSAGE_LATENCY
	MESSAGE_ASYNC_START      MessageType = C.GST_MESSAGE_ASYNC_START
	MESSAGE_ASYNC_DONE       MessageType = C.GST_MESSAGE_ASYNC_DONE
	MESSAGE_REQUEST_STATE    MessageType = C.GST_MESSAGE_REQUEST_STATE
	MESSAGE_STEP_START       MessageType = C.GST_MESSAGE_STEP_START
	MESSAGE_QOS              MessageType = C.GST_MESSAGE_QOS
	MESSAGE_PROGRESS         MessageType = C.GST_MESSAGE_PROGRESS
	MESSAGE_TOC              MessageType = C.GST_MESSAGE_TOC
	MESSAGE_RESET_TIME       MessageType = C.GST_MESSAGE_RESET_TIME
	MESSAGE_STREAM_START     MessageType = C.GST_MESSAGE_STREAM_START
	MESSAGE_ANY              MessageType = C.GST_MESSAGE_ANY
)

func marshalMessageType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return MessageType(c), nil
}

// State is a representation of GstState.
type State int

const (
	STATE_VOID_PENDING State = C.GST_STATE_VOID_PENDING
	STATE_NULL         State = C.GST_STATE_NULL
	STATE_READY        State = C.GST_STATE_READY
	STATE_PAUSED       State = C.GST_STATE_PAUSED
	STATE_PLAYING      State = C.GST_STATE_PLAYING
)

func marshalState(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return State(c), nil
}

// StateChangeReturn is a representation of GstStateChangeReturn.
type StateChangeReturn int

const (
	STATE_CHANGE_FAILURE    StateChangeReturn = C.GST_STATE_CHANGE_FAILURE
	STATE_CHANGE_SUCCESS    StateChangeReturn = C.GST_STATE_CHANGE_SUCCESS
	STATE_CHANGE_ASYNC      StateChangeReturn = C.GST_STATE_CHANGE_ASYNC
	STATE_CHANGE_NO_PREROLL StateChangeReturn = C.GST_STATE_CHANGE_NO_PREROLL
)

func marshalStateChangeReturn(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return StateChangeReturn(c), nil
}

/*
 * GstBus
 */

type Bus struct {
	Object
}

// native returns a pointer to the underlying GstBus.
func (v *Bus) native() *C.GstBus {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGstBus(p)
}

func marshalBus(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapBus(obj), nil
}

func wrapBus(obj *glib.Object) *Bus {
	return &Bus{Object{glib.InitiallyUnowned{obj}}}
}

// AddSignalWatch() is a wrapper around gst_bus_add_signal_watch().
func (v *Bus) AddSignalWatch() {
	C.gst_bus_add_signal_watch(v.native())
}

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

// GetBus() is a wrapper around gst_element_get_bus().
func (v *Element) GetBus() (*Bus, error) {
	c := C.gst_element_get_bus(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapBus(obj)
	b.RefSink()
	runtime.SetFinalizer(&b.Object, (*Object).Unref)
	return b, nil
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

// QueryDuration() is a wrapper around gst_element_query_duration().
func (v *Element) QueryDuration(format Format) (cur int64, success bool) {
	var ccur C.gint64
	c := C.gst_element_query_duration(v.native(), C.GstFormat(format), &ccur)
	return int64(ccur), gobool(c)
}

// QueryPosition() is a wrapper around gst_element_query_position().
func (v *Element) QueryPosition(format Format) (cur int64, success bool) {
	var ccur C.gint64
	c := C.gst_element_query_position(v.native(), C.GstFormat(format), &ccur)
	return int64(ccur), gobool(c)
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
 * GstMessage
 */

// Message is a representation of GDK's GstMessage.
type Message struct {
	GstMessage *C.GstMessage
}

// native returns a pointer to the underlying GstMessage.
func (v *Message) native() *C.GstMessage {
	if v == nil {
		return nil
	}
	return v.GstMessage
}

// Native returns a pointer to the underlying GdkEvent.
func (v *Message) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalMessage(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &Message{(*C.GstMessage)(unsafe.Pointer(c))}, nil
}

// GetSeqnum() is a wrapper around GST_MESSAGE_SEQNUM().
func (v *Message) GetSeqnum() uint64 {
	c := C.messageSeqnum(unsafe.Pointer(v.native()))
	return uint64(c)
}

// GetType() is a wrapper around GST_MESSAGE_TYPE().
func (v *Message) GetType() MessageType {
	c := C.toGstMessageType(unsafe.Pointer(v.native()))
	return MessageType(c)
}

// Timestamp() is a wrapper around GST_MESSAGE_TIMESTAMP().
func (v *Message) GetTimestamp() uint64 {
	c := C.messageTimestamp(unsafe.Pointer(v.native()))
	return uint64(c)
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

// GetName() is a wrapper around gst_object_get_name().
func (v *Object) GetName() string {
	c := C.gst_object_get_name(v.native())
	defer C.g_free(C.gpointer(c))
	return C.GoString((*C.char)(c))
}

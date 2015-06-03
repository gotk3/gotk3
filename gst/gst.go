// Go bindings for gstreamer1.0.  Supports version 1.0 and later.
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

	"github.com/andre-hub/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.gst_buffering_mode_get_type()), marshalBufferingMode},
		{glib.Type(C.gst_format_get_type()), marshalFormat},
		{glib.Type(C.gst_message_type_get_type()), marshalMessageType},
		{glib.Type(C.gst_pad_link_return_get_type()), marshalPadLinkReturn},
		{glib.Type(C.gst_state_get_type()), marshalState},
		{glib.Type(C.gst_seek_flags_get_type()), marshalSeekFlags},
		{glib.Type(C.gst_seek_type_get_type()), marshalSeekType},
		{glib.Type(C.gst_state_change_return_get_type()), marshalStateChangeReturn},

		// Objects/Interfaces
		{glib.Type(C.gst_bin_get_type()), marshalBin},
		{glib.Type(C.gst_bus_get_type()), marshalBus},
		{glib.Type(C.gst_element_get_type()), marshalElement},
		{glib.Type(C.gst_element_factory_get_type()), marshalElementFactory},
		{glib.Type(C.gst_ghost_pad_get_type()), marshalGhostPad},
		{glib.Type(C.gst_object_get_type()), marshalObject},
		{glib.Type(C.gst_pad_get_type()), marshalPad},
		{glib.Type(C.gst_plugin_feature_get_type()), marshalPluginFeature},

		// Boxed
		{glib.Type(C.gst_message_get_type()), marshalMessage},
	}
	glib.RegisterGValueMarshalers(tm)
}

/*
Init() is a wrapper around gst_init() and must be called before any
other gstreamer calls and is used to initialize everything necessary.

In addition to setting up gstreamer for usage, a pointer to a slice of
strings may be passed in to parse standard gst command line arguments.
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

// Unref is a wrapper around gst_object_unref(). This is for internal use.
func (v *Object) Unref() {
	C.gst_object_unref(C.gpointer(v.toObject()))
}

// RefSink is a wrapper around gst_object_ref_sink(). This is for internal use.
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

// BufferingMode is a representation of GstBufferingMode
type BufferingMode int

const (
	BUFFERING_STREAM    BufferingMode = C.GST_BUFFERING_STREAM
	BUFFERING_DOWNLOAD  BufferingMode = C.GST_BUFFERING_DOWNLOAD
	BUFFERING_TIMESHIFT BufferingMode = C.GST_BUFFERING_TIMESHIFT
	BUFFERING_LIVE      BufferingMode = C.GST_BUFFERING_LIVE
)

func marshalBufferingMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return BufferingMode(c), nil
}

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

// PadLinkReturn os a representation of GstPadLinkReturn.
type PadLinkReturn int

const (
	PAD_LINK_OK              PadLinkReturn = C.GST_PAD_LINK_OK
	PAD_LINK_WRONG_HIERARCHY PadLinkReturn = C.GST_PAD_LINK_WRONG_HIERARCHY
	PAD_LINK_WAS_LINKED      PadLinkReturn = C.GST_PAD_LINK_WAS_LINKED
	PAD_LINK_WRONG_DIRECTION PadLinkReturn = C.GST_PAD_LINK_WRONG_DIRECTION
	PAD_LINK_NOFORMAT        PadLinkReturn = C.GST_PAD_LINK_NOFORMAT
	PAD_LINK_NOSCHED         PadLinkReturn = C.GST_PAD_LINK_NOSCHED
	PAD_LINK_REFUSED         PadLinkReturn = C.GST_PAD_LINK_REFUSED
)

func marshalPadLinkReturn(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return PadLinkReturn(c), nil
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

// GetName() is a wrapper around gst_element_state_get_name().
func (s *State) GetName() string {
	c := C.gst_element_state_get_name(C.GstState(*s))
	return C.GoString((*C.char)(c))
}

// SeekFlags is a representation of GstSeekFlags.
type SeekFlags int

const (
	SEEK_FLAG_NONE         SeekFlags = C.GST_SEEK_FLAG_NONE
	SEEK_FLAG_FLUSH        SeekFlags = C.GST_SEEK_FLAG_FLUSH
	SEEK_FLAG_ACCURATE     SeekFlags = C.GST_SEEK_FLAG_ACCURATE
	SEEK_FLAG_KEY_UNIT     SeekFlags = C.GST_SEEK_FLAG_KEY_UNIT
	SEEK_FLAG_SEGMENT      SeekFlags = C.GST_SEEK_FLAG_SEGMENT
	SEEK_FLAG_SKIP         SeekFlags = C.GST_SEEK_FLAG_SKIP
	SEEK_FLAG_SNAP_BEFORE  SeekFlags = C.GST_SEEK_FLAG_SNAP_BEFORE
	SEEK_FLAG_SNAP_AFTER   SeekFlags = C.GST_SEEK_FLAG_SNAP_AFTER
	SEEK_FLAG_SNAP_NEAREST SeekFlags = C.GST_SEEK_FLAG_SNAP_NEAREST
)

func marshalSeekFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SeekFlags(c), nil
}

// SeekType is a representation of GstSeekType.
type SeekType int

const (
	SEEK_TYPE_NONE SeekType = C.GST_SEEK_TYPE_NONE
	SEEK_TYPE_SET  SeekType = C.GST_SEEK_TYPE_SET
	SEEK_TYPE_END  SeekType = C.GST_SEEK_TYPE_END
)

func marshalSeekType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SeekType(c), nil
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
 * GstBin
 */

type Bin struct {
	Element
}

// native returns a pointer to the underlying GstBin.
func (v *Bin) native() *C.GstBin {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGstBin(p)
}

func marshalBin(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapBin(obj), nil
}

func wrapBin(obj *glib.Object) *Bin {
	return &Bin{Element{Object{glib.InitiallyUnowned{obj}}}}
}

// BinNew() is a wrapper around gst_bin_new().
func BinNew(name string) (*Bin, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	c := C.gst_bin_new((*C.gchar)(cname))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapBin(obj)
	b.RefSink()
	runtime.SetFinalizer(&b.Object, (*Object).Unref)
	return b, nil
}

// Add() is a wrapper around gst_bin_add().
func (v *Bin) Add(e IElement) bool {
	c := C.gst_bin_add(v.native(), e.toElement())
	return gobool(c)
}

// Remove() is a wrapper around gst_bin_remove().
func (v *Bin) Remove(e IElement) bool {
	c := C.gst_bin_remove(v.native(), e.toElement())
	return gobool(c)
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

// IElement is an interface type implemented by all structs
// embedding a Element.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C gst function taking a
// GstElement.
type IElement interface {
	toElement() *C.GstElement
}

func (v *Element) toElement() *C.GstElement {
	if v == nil {
		return nil
	}
	return v.native()
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

// AddPad() is a wrapper around gst_element_add_pad().
func (v *Element) AddPad(pad IPad) bool {
	c := C.gst_element_add_pad(v.native(), pad.toPad())
	return gobool(c)
}

// GetBus() is a wrapper around gst_element_get_bus().
func (v *Element) GetBus() (*Bus, error) {
	c := C.gst_element_get_bus(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapBus(obj)
	//b.RefSink()
	runtime.SetFinalizer(&b.Object, (*Object).Unref)
	return b, nil
}

// GetState() is a wrapper around gst_element_get_state().
func (v *Element) GetState(timeout uint64) (state, pending State, change StateChangeReturn) {
	var cstate, cpending C.GstState
	c := C.gst_element_get_state(v.native(), &cstate, &cpending, C.GstClockTime(timeout))
	return State(cstate), State(cpending), StateChangeReturn(c)
}

// GetStaticPad() is a wrapper around gst_element_get_static_pad().
func (v *Element) GetStaticPad(name string) (*Pad, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	c := C.gst_element_get_static_pad(v.native(), (*C.gchar)(cname))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapPad(obj)
	//b.RefSink()
	runtime.SetFinalizer(&b.Object, (*Object).Unref)
	return b, nil
}

// RemovePad() is a wrapper around gst_element_remove_pad().
func (v *Element) RemovePad(pad IPad) bool {
	c := C.gst_element_remove_pad(v.native(), pad.toPad())
	return gobool(c)
}

// Seek() is a wrapper around gst_element_seek().
func (v *Element) Seek(rate float64, format Format, flags SeekFlags, startType SeekType, start int64, stopType SeekType, stop int64) bool {
	c := C.gst_element_seek(v.native(),
		C.gdouble(rate),
		C.GstFormat(format),
		C.GstSeekFlags(flags),
		C.GstSeekType(startType),
		C.gint64(start),
		C.GstSeekType(stopType),
		C.gint64(stop))
	return gobool(c)
}

// SeekSimple() is a wrapper around gst_element_seek_simple().
func (v *Element) SeekSimple(format Format, flags SeekFlags, pos int64) bool {
	c := C.gst_element_seek_simple(v.native(), C.GstFormat(format), C.GstSeekFlags(flags), C.gint64(pos))
	return gobool(c)
}

// SetState() is a wrapper around gst_element_set_state().
func (v *Element) SetState(state State) StateChangeReturn {
	c := C.gst_element_set_state(v.native(), C.GstState(state))
	return StateChangeReturn(c)
}

// SyncStateWithParent() is a wrapper around gst_element_sync_state_with_parent().
func (v *Element) SyncStateWithParent() bool {
	c := C.gst_element_sync_state_with_parent(v.native())
	return gobool(c)
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

type ElementFactory struct {
	PluginFeature
}

// native returns a pointer to the underlying GstElement.
func (v *ElementFactory) native() *C.GstElementFactory {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGstElementFactory(p)
}

func marshalElementFactory(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapElementFactory(obj), nil
}

func wrapElementFactory(obj *glib.Object) *ElementFactory {
	return &ElementFactory{PluginFeature{Object{glib.InitiallyUnowned{obj}}}}
}

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

// ElementFactoryFind() is a wrapper around gst_element_factory_find().
func ElementFactoryFind(name string) (*ElementFactory, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	c := C.gst_element_factory_find((*C.gchar)(cname))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapElementFactory(obj)
	b.RefSink()
	runtime.SetFinalizer(&b.Object, (*Object).Unref)
	return b, nil
}

/*
 * GstGhostPad
 */

type GhostPad struct {
	Pad
}

// native returns a pointer to the underlying GstPad.
func (v *GhostPad) native() *C.GstGhostPad {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGstGhostPad(p)
}

func marshalGhostPad(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapGhostPad(obj), nil
}

func wrapGhostPad(obj *glib.Object) *GhostPad {
	return &GhostPad{Pad{Object{glib.InitiallyUnowned{obj}}}}
}

// GhostPadNew() is a wrapper around gst_ghost_pad_new().
func GhostPadNew(name string, pad IPad) (*GhostPad, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	c := C.gst_ghost_pad_new((*C.gchar)(cname), pad.toPad())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapGhostPad(obj)
	b.RefSink()
	runtime.SetFinalizer(&b.Object, (*Object).Unref)
	return b, nil
}

/*
 * GstMessage
 */

// Message is a representation of Gst's GstMessage.
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

// GetTypeName() is a wrapper around GST_MESSAGE_TYPE_NAME().
func (v *Message) GetTypeName() string {
	c := C.messageTypeName(unsafe.Pointer(v.native()))
	return C.GoString(c)
}

// ParseBuffering() is a wrapper around gst_message_parse_buffering().
func (v *Message) ParseBuffering() (percent int) {
	var per C.gint
	C.gst_message_parse_buffering(v.native(), &per)
	return int(per)
}

// ParseBufferingStats() is a wrapper around gst_message_parse_buffering_stats().
func (v *Message) ParseBufferingStats() (mode BufferingMode, avgIn, avgOut int, bufferingLeft int64) {
	var cavgIn, cavgOut C.gint
	var cbmode C.GstBufferingMode
	var cbleft C.gint64
	C.gst_message_parse_buffering_stats(v.native(), &cbmode, &cavgIn, &cavgOut, &cbleft)
	return BufferingMode(cbmode), int(cavgIn), int(cavgOut), int64(cbleft)
}

// ParseError() is a wrapper around gst_message_parse_error().
func (v *Message) ParseError() (err error, debug string) {
	var e *C.GError
	var d *C.gchar
	C.gst_message_parse_error(v.native(), &e, &d)
	eMessage := C.GoString((*C.char)(e.message))
	defer C.g_error_free(e)
	debug = C.GoString((*C.char)(d))
	defer C.g_free((C.gpointer)(d))
	return errors.New(eMessage), debug
}

// ParseStateChanged() is a wrapper around gst_message_parse_state_changed().
func (v *Message) ParseStateChanged() (oldState, newState, pending State) {
	var o, n, p C.GstState
	C.gst_message_parse_state_changed(v.native(), &o, &n, &p)
	return State(o), State(n), State(p)
}

// ParseStepDone() is a wrapper around gst_message_parse_step_done().
func (v *Message) ParseStepDone() (format Format, amount uint64, rate float64, flush, intermediate bool, duration uint64, eos bool) {
	var cformat C.GstFormat
	var camount, cduration C.guint64
	var crate C.gdouble
	var cflush, cintermediate, ceos C.gboolean
	C.gst_message_parse_step_done(v.native(), &cformat, &camount, &crate, &cflush, &cintermediate, &cduration, &ceos)
	return Format(cformat), uint64(camount), float64(crate), gobool(cflush), gobool(cintermediate), uint64(cduration), gobool(ceos)
}

// ParseSegmentStart() is a wrapper around gst_message_parse_segment_start().
func (v *Message) ParseSegmentStart() (format Format, position int64) {
	var cformat C.GstFormat
	var cposition C.gint64
	C.gst_message_parse_segment_start(v.native(), &cformat, &cposition)
	return Format(cformat), int64(cposition)
}

// ParseSegmentDone() is a wrapper around gst_message_parse_segment_done().
func (v *Message) ParseSegmentDone() (format Format, position int64) {
	var cformat C.GstFormat
	var cposition C.gint64
	C.gst_message_parse_segment_done(v.native(), &cformat, &cposition)
	return Format(cformat), int64(cposition)
}

// ParseQosValues() is a wrapper around gst_message_parse_qos_values().
func (v *Message) ParseQosValues() (jitter int64, proportion float64, quality int) {
	var cjitter C.gint64
	var cproportion C.gdouble
	var cquality C.gint
	C.gst_message_parse_qos_values(v.native(), &cjitter, &cproportion, &cquality)
	return int64(cjitter), float64(cproportion), int(cquality)
}

// ParseQosStats() is a wrapper around gst_message_parse_qos_stats().
func (v *Message) ParseQosStats() (format Format, processed, dropped uint64) {
	var cformat C.GstFormat
	var cprocessed, cdropped C.guint64
	C.gst_message_parse_qos_stats(v.native(), &cformat, &cprocessed, &cdropped)
	return Format(cformat), uint64(cprocessed), uint64(cdropped)
}

// ParseWarning() is a wrapper around gst_message_parse_warning().
func (v *Message) ParseWarning() (err error, debug string) {
	var e *C.GError
	var d *C.gchar
	C.gst_message_parse_warning(v.native(), &e, &d)
	eMessage := C.GoString((*C.char)(e.message))
	defer C.g_error_free(e)
	debug = C.GoString((*C.char)(d))
	defer C.g_free((C.gpointer)(d))
	return errors.New(eMessage), debug
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

/*
 * GstPad
 */

type Pad struct {
	Object
}

type IPad interface {
	toPad() *C.GstPad
}

// native returns a pointer to the underlying GstPad.
func (v *Pad) native() *C.GstPad {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGstPad(p)
}

func (v *Pad) toPad() *C.GstPad {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalPad(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapPad(obj), nil
}

func wrapPad(obj *glib.Object) *Pad {
	return &Pad{Object{glib.InitiallyUnowned{obj}}}
}

// GetPeer() is a wrapper around gst_pad_get_peer().
func (v *Pad) GetPeer() (*Pad, error) {
	c := C.gst_pad_get_peer(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapPad(obj)
	//b.RefSink()
	runtime.SetFinalizer(&b.Object, (*Object).Unref)
	return b, nil
}

// Link is a wrapper around gst_pad_link().
func (v *Pad) Link(sinkPad IPad) PadLinkReturn {
	c := C.gst_pad_link(v.native(), sinkPad.toPad())
	return PadLinkReturn(c)
}

// SetActive() is a wrapper around gst_pad_set_active().
func (v *Pad) SetActive(active bool) bool {
	c := C.gst_pad_set_active(v.native(), gbool(active))
	return gobool(c)
}

// Unlink() is a wrapper around gst_pad_unlink().
func (v *Pad) Unlink(sinkPad IPad) bool {
	c := C.gst_pad_unlink(v.native(), sinkPad.toPad())
	return gobool(c)
}

/*
 * GstPluginFeature
 */

type PluginFeature struct {
	Object
}

// native returns a pointer to the underlying GstPluginFeature.
func (v *PluginFeature) native() *C.GstPluginFeature {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGstPluginFeature(p)
}

func marshalPluginFeature(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapPluginFeature(obj), nil
}

func wrapPluginFeature(obj *glib.Object) *PluginFeature {
	return &PluginFeature{Object{glib.InitiallyUnowned{obj}}}
}

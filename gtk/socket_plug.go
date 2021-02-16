// +build linux

package gtk

// #include <gtk/gtkx.h>
// #include <gtk/gtksocket.h>
// #include <gtk/gtkplug.h>
// #include <gdk/gdkdisplay.h>
// #include "socket_plug.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_socket_get_type()), marshalSocket},
		{glib.Type(C.gtk_plug_get_type()), marshalPlug},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkSocket"] = wrapSocket
	WrapMap["GtkPlug"] = wrapPlug
}

/*
 * GtkSocket
 */

// Socket is a representation of GTK's GtkSocket
type Socket struct {
	Container
}

// native returns a pointer to the underlying GtkSocket
func (v *Socket) native() *C.GtkSocket {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSocket(p)
}

func marshalSocket(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSocket(obj), nil
}

func wrapSocket(obj *glib.Object) *Socket {
	if obj == nil {
		return nil
	}

	return &Socket{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// SocketNew is a wrapper around gtk_socket_new().
// Create a new empty GtkSocket.
func SocketNew() (*Socket, error) {
	c := C.gtk_socket_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSocket(glib.Take(unsafe.Pointer(c))), nil
}

// AddId is a wrapper around gtk_socket_add_id().
// Adds an XEMBED client, such as a GtkPlug, to the GtkSocket. The client may be in the same process or in a different
// process.
// To embed a GtkPlug in a GtkSocket, you can either create the GtkPlug with gtk_plug_new (0), call gtk_plug_get_id()
// to get the window ID of the plug, and then pass that to the gtk_socket_add_id(), or you can call gtk_socket_get_id()
// to get the window ID for the socket, and call gtk_plug_new() passing in that ID.
// The GtkSocket must have already be added into a toplevel window before you can make this call.
func (v *Socket) AddId(window uint) {
	C.gtk_socket_add_id(v.native(), C.Window(window))
}

// GetId is a wrapper around gtk_socket_get_id().
// Gets the window ID of a GtkSocket widget, which can then be used to create a client embedded inside the socket,
// for instance with gtk_plug_new().
// The GtkSocket must have already be added into a toplevel window before you can make this call.
func (v *Socket) GetId() uint {
	c := C.gtk_socket_get_id(v.native())
	return uint(c)
}

// GetPlugWindow is a wrapper around gtk_socket_get_plug_window().
// Retrieves the window of the plug. Use this to check if the plug has been created inside of the socket.
func (v *Socket) GetPlugWindow() (*Window, error) {
	c := C.gtk_socket_get_plug_window(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWindow(glib.Take(unsafe.Pointer(c))), nil
}

/*
 * GtkSocket
 */

// Plug is a representation of GTK's GtkPlug
type Plug struct {
	Window
}

// native returns a pointer to the underlying GtkSocket
func (v *Plug) native() *C.GtkPlug {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkPlug(p)
}

// native returns a C pointer to the underlying GdkDisplay.
func native(v *gdk.Display) *C.GdkDisplay {
	// I'd love to not have to copy native() from gdk/gdk.go, but it appears you can't just
	// pass C structs from package to package
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkDisplay(p)
}

func marshalPlug(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapPlug(obj), nil
}

func wrapPlug(obj *glib.Object) *Plug {
	if obj == nil {
		return nil
	}

	return &Plug{Window{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}}
}

// PlugNew is a wrapper around gtk_plug_new().
// Creates a new plug widget inside the GtkSocket identified by socket_id.
// If socket_id is 0, the plug is left “unplugged” and can later be plugged into a GtkSocket by gtk_socket_add_id().
func PlugNew(socketId uint) (*Plug, error) {
	c := C.gtk_plug_new(C.Window(socketId))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapPlug(glib.Take(unsafe.Pointer(c))), nil
}

// PlugNewForDisplay is a wrapper around gtk_plug_new_for_display().
// Creates a new plug widget inside the GtkSocket identified by socket_id.
// If socket_id is 0, the plug is left “unplugged” and can later be plugged into a GtkSocket by gtk_socket_add_id().
func PlugNewForDisplay(display *gdk.Display, socketId uint) (*Plug, error) {
	c := C.gtk_plug_new_for_display(native(display), C.Window(socketId))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapPlug(glib.Take(unsafe.Pointer(c))), nil
}

// Construct is a wrapper around gtk_plug_construct().
// Finish the initialization of plug for a given GtkSocket identified by socket_id.
// This function will generally only be used by classes deriving from GtkPlug.
func (v *Plug) Construct(socketId uint) {
	C.gtk_plug_construct(v.native(), C.Window(socketId))
}

// ConstructForDisplay is a wrapper around gtk_plug_construct_for_display().
// Finish the initialization of plug for a given GtkSocket identified by socket_id which is currently
// displayed on display.
// This function will generally only be used by classes deriving from GtkPlug.
func (v *Plug) ConstructForDisplay(display *gdk.Display, socketId uint) {
	C.gtk_plug_construct_for_display(v.native(), native(display), C.Window(socketId))
}

// GetId is a wrapper around gtk_plug_get_id().
// Gets the window ID of a GtkPlug widget, which can then be used to embed this window inside another window,
// for instance with gtk_socket_add_id().
func (v *Plug) GetId() uint {
	c := C.gtk_plug_get_id(v.native())
	return uint(c)
}

// GetEmbedded is a wrapper around gtk_plug_get_embedded().
// Determines whether the plug is embedded in a socket.
func (v *Plug) GetEmbedded() bool {
	c := C.gtk_plug_get_embedded(v.native())
	return gobool(c)
}

// GetSocketWindow is a wrapper around gtk_plug_get_socket_window().
// Retrieves the socket the plug is embedded in.
func (v *Plug) GetSocketWindow() (*Window, error) {
	c := C.gtk_plug_get_socket_window(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWindow(glib.Take(unsafe.Pointer(c))), nil
}

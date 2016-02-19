package gtkf

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	glib_impl "github.com/gotk3/gotk3/glibf"
	"github.com/gotk3/gotk3/gtk"
)

func init() {
	tm := []glib_impl.TypeMarshaler{
		{glib.Type(C.gtk_info_bar_get_type()), marshalInfoBar},
	}

	glib_impl.RegisterGValueMarshalers(tm)

	WrapMap["GtkInfoBar"] = wrapInfoBar
}

type infoBar struct {
	box
}

func (v *infoBar) native() *C.GtkInfoBar {
	if v == nil || v.GObject == nil {
		return nil
	}

	p := unsafe.Pointer(v.GObject)
	return C.toGtkInfoBar(p)
}

func marshalInfoBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapInfoBar(wrapObject(unsafe.Pointer(c))), nil
}

func wrapInfoBar(obj *glib_impl.Object) *infoBar {
	return &infoBar{box{container{widget{glib_impl.InitiallyUnowned{obj}}}}}
}

func InfoBarNew() (*infoBar, error) {
	c := C.gtk_info_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}

	return wrapInfoBar(wrapObject(unsafe.Pointer(c))), nil
}

func (v *infoBar) AddActionWidget(w gtk.Widget, responseId gtk.ResponseType) {
	C.gtk_info_bar_add_action_widget(v.native(), w.(IWidget).toWidget(), C.gint(responseId))
}

func (v *infoBar) AddButton(buttonText string, responseId gtk.ResponseType) {
	cstr := C.CString(buttonText)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_info_bar_add_button(v.native(), (*C.gchar)(cstr), C.gint(responseId))
}

func (v *infoBar) SetResponseSensitive(responseId gtk.ResponseType, setting bool) {
	C.gtk_info_bar_set_response_sensitive(v.native(), C.gint(responseId), gbool(setting))
}

func (v *infoBar) SetDefaultResponse(responseId gtk.ResponseType) {
	C.gtk_info_bar_set_default_response(v.native(), C.gint(responseId))
}

func (v *infoBar) SetMessageType(messageType gtk.MessageType) {
	C.gtk_info_bar_set_message_type(v.native(), C.GtkMessageType(messageType))
}

func (v *infoBar) GetMessageType() gtk.MessageType {
	messageType := C.gtk_info_bar_get_message_type(v.native())
	return gtk.MessageType(messageType)
}

func (v *infoBar) GetActionArea() (gtk.Widget, error) {
	c := C.gtk_info_bar_get_action_area(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return wrapWidget(wrapObject(unsafe.Pointer(c))), nil
}

func (v *infoBar) GetContentArea() (gtk.Box, error) {
	c := C.gtk_info_bar_get_content_area(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return wrapBox(wrapObject(unsafe.Pointer(c))), nil
}

func (v *infoBar) GetShowCloseButton() bool {
	b := C.gtk_info_bar_get_show_close_button(v.native())
	return gobool(b)
}

func (v *infoBar) SetShowCloseButton(setting bool) {
	C.gtk_info_bar_set_show_close_button(v.native(), gbool(setting))
}

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_info_bar_get_type()), marshalInfoBar},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkInfoBar"] = wrapInfoBar
}

type InfoBar struct {
	Box
}

func (v *InfoBar) native() *C.GtkInfoBar {
	if v == nil || v.GObject == nil {
		return nil
	}

	p := unsafe.Pointer(v.GObject)
	return C.toGtkInfoBar(p)
}

func marshalInfoBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapInfoBar(glib.Take(unsafe.Pointer(c))), nil
}

func wrapInfoBar(obj *glib.Object) *InfoBar {
	if obj == nil {
		return nil
	}

	return &InfoBar{Box{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// InfoBarNew is a wrapper around gtk_info_bar_new().
func InfoBarNew() (*InfoBar, error) {
	c := C.gtk_info_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}

	return wrapInfoBar(glib.Take(unsafe.Pointer(c))), nil
}

// TODO:
// gtk_info_bar_new_with_buttons().

// AddActionWidget is a wrapper around gtk_info_bar_add_action_widget().
func (v *InfoBar) AddActionWidget(w IWidget, responseId ResponseType) {
	C.gtk_info_bar_add_action_widget(v.native(), w.toWidget(), C.gint(responseId))
}

// AddButton is a wrapper around gtk_info_bar_add_button().
func (v *InfoBar) AddButton(buttonText string, responseId ResponseType) {
	cstr := C.CString(buttonText)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_info_bar_add_button(v.native(), (*C.gchar)(cstr), C.gint(responseId))
}

// TODO:
// gtk_info_bar_add_buttons().

// SetResponseSensitive is a wrapper around gtk_info_bar_set_response_sensitive().
func (v *InfoBar) SetResponseSensitive(responseId ResponseType, setting bool) {
	C.gtk_info_bar_set_response_sensitive(v.native(), C.gint(responseId), gbool(setting))
}

// SetDefaultResponse is a wrapper around gtk_info_bar_set_default_response().
func (v *InfoBar) SetDefaultResponse(responseId ResponseType) {
	C.gtk_info_bar_set_default_response(v.native(), C.gint(responseId))
}

// TODO:
// gtk_info_bar_response().

// SetMessageType is a wrapper around gtk_info_bar_set_message_type().
func (v *InfoBar) SetMessageType(messageType MessageType) {
	C.gtk_info_bar_set_message_type(v.native(), C.GtkMessageType(messageType))
}

// GetMessageType is a wrapper around gtk_info_bar_get_message_type().
func (v *InfoBar) GetMessageType() MessageType {
	messageType := C.gtk_info_bar_get_message_type(v.native())
	return MessageType(messageType)
}

// GetActionArea is a wrapper around gtk_info_bar_get_action_area().
func (v *InfoBar) GetActionArea() (IWidget, error) {
	c := C.gtk_info_bar_get_action_area(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return castWidget(c)
}

// GetContentArea is a wrapper around gtk_info_bar_get_content_area().
func (v *InfoBar) GetContentArea() (*Box, error) {
	c := C.gtk_info_bar_get_content_area(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return wrapBox(glib.Take(unsafe.Pointer(c))), nil
}

// GetShowCloseButton is a wrapper around gtk_info_bar_get_show_close_button().
func (v *InfoBar) GetShowCloseButton() bool {
	b := C.gtk_info_bar_get_show_close_button(v.native())
	return gobool(b)
}

// SetShowCloseButton is a wrapper around gtk_info_bar_set_show_close_button().
func (v *InfoBar) SetShowCloseButton(setting bool) {
	C.gtk_info_bar_set_show_close_button(v.native(), gbool(setting))
}

// TODO: for GTK+ 3.22.29
// gtk_info_bar_get_revealed().
// gtk_info_bar_set_revealed().

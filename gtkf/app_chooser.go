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
		{glib.Type(C.gtk_app_chooser_get_type()), marshalAppChooser},
		{glib.Type(C.gtk_app_chooser_button_get_type()), marshalAppChooserButton},
		{glib.Type(C.gtk_app_chooser_widget_get_type()), marshalAppChooserWidget},
		{glib.Type(C.gtk_app_chooser_dialog_get_type()), marshalAppChooserDialog},
	}

	glib_impl.RegisterGValueMarshalers(tm)

	WrapMap["GtkAppChooser"] = wrapAppChooser
	WrapMap["GtkAppChooserButton"] = wrapAppChooserButton
	WrapMap["GtkAppChooserWidget"] = wrapAppChooserWidget
	WrapMap["GtkAppChooserDialog"] = wrapAppChooserDialog
}

/*
 * GtkAppChooser
 */

// AppChooser is a representation of GTK's GtkAppChooser GInterface.
type appChooser struct {
	*glib_impl.Object
}

// IAppChooser is an interface type implemented by all structs
// embedding an AppChooser. It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkAppChooser.
type IAppChooser interface {
	toAppChooser() *C.GtkAppChooser
}

// native returns a pointer to the underlying GtkAppChooser.
func (v *appChooser) native() *C.GtkAppChooser {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkAppChooser(p)
}

func marshalAppChooser(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAppChooser(obj), nil
}

func wrapAppChooser(obj *glib_impl.Object) *appChooser {
	return &appChooser{obj}
}

func (v *appChooser) toAppChooser() *C.GtkAppChooser {
	if v == nil {
		return nil
	}
	return v.native()
}

// TODO: Needs gio/GAppInfo implementation first
// gtk_app_chooser_get_app_info ()

// GetContentType is a wrapper around gtk_app_chooser_get_content_type().
func (v *appChooser) GetContentType() string {
	cstr := C.gtk_app_chooser_get_content_type(v.native())
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString((*C.char)(cstr))
}

// Refresh is a wrapper around gtk_app_chooser_refresh().
func (v *appChooser) Refresh() {
	C.gtk_app_chooser_refresh(v.native())
}

/*
 * GtkAppChooserButton
 */

// AppChooserButton is a representation of GTK's GtkAppChooserButton.
type appChooserButton struct {
	comboBox

	// Interfaces
	appChooser
}

// native returns a pointer to the underlying GtkAppChooserButton.
func (v *appChooserButton) native() *C.GtkAppChooserButton {
	if v == nil || v.GObject == nil {
		return nil
	}

	p := unsafe.Pointer(v.GObject)
	return C.toGtkAppChooserButton(p)
}

func marshalAppChooserButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapAppChooserButton(wrapObject(unsafe.Pointer(c))), nil
}

func wrapAppChooserButton(obj *glib_impl.Object) *appChooserButton {
	cl := wrapCellLayout(obj)
	ac := wrapAppChooser(obj)
	return &appChooserButton{comboBox{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}, *cl}, *ac}
}

// AppChooserButtonNew() is a wrapper around gtk_app_chooser_button_new().
func AppChooserButtonNew(content_type string) (*appChooserButton, error) {
	cstr := C.CString(content_type)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_app_chooser_button_new((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapAppChooserButton(wrapObject(unsafe.Pointer(c))), nil
}

// TODO: Needs gio/GIcon implemented first
// gtk_app_chooser_button_append_custom_item ()

// AppendSeparator() is a wrapper around gtk_app_chooser_button_append_separator().
func (v *appChooserButton) AppendSeparator() {
	C.gtk_app_chooser_button_append_separator(v.native())
}

// SetActiveCustomItem() is a wrapper around gtk_app_chooser_button_set_active_custom_item().
func (v *appChooserButton) SetActiveCustomItem(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_app_chooser_button_set_active_custom_item(v.native(), (*C.gchar)(cstr))
}

// GetShowDefaultItem() is a wrapper around gtk_app_chooser_button_get_show_default_item().
func (v *appChooserButton) GetShowDefaultItem() bool {
	return gobool(C.gtk_app_chooser_button_get_show_default_item(v.native()))
}

// SetShowDefaultItem() is a wrapper around gtk_app_chooser_button_set_show_default_item().
func (v *appChooserButton) SetShowDefaultItem(setting bool) {
	C.gtk_app_chooser_button_set_show_default_item(v.native(), gbool(setting))
}

// GetShowDialogItem() is a wrapper around gtk_app_chooser_button_get_show_dialog_item().
func (v *appChooserButton) GetShowDialogItem() bool {
	return gobool(C.gtk_app_chooser_button_get_show_dialog_item(v.native()))
}

// SetShowDialogItem() is a wrapper around gtk_app_chooser_button_set_show_dialog_item().
func (v *appChooserButton) SetShowDialogItem(setting bool) {
	C.gtk_app_chooser_button_set_show_dialog_item(v.native(), gbool(setting))
}

// GetHeading() is a wrapper around gtk_app_chooser_button_get_heading().
// In case when gtk_app_chooser_button_get_heading() returns a nil string,
// GetHeading() returns a non-nil error.
func (v *appChooserButton) GetHeading() (string, error) {
	cstr := C.gtk_app_chooser_button_get_heading(v.native())
	if cstr == nil {
		return "", nilPtrErr
	}
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString((*C.char)(cstr)), nil
}

// SetHeading() is a wrapper around gtk_app_chooser_button_set_heading().
func (v *appChooserButton) SetHeading(heading string) {
	cstr := C.CString(heading)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_app_chooser_button_set_heading(v.native(), (*C.gchar)(cstr))
}

/*
 * GtkAppChooserWidget
 */

// AppChooserWidget is a representation of GTK's GtkAppChooserWidget.
type appChooserWidget struct {
	box

	// Interfaces
	appChooser
}

// native returns a pointer to the underlying GtkAppChooserWidget.
func (v *appChooserWidget) native() *C.GtkAppChooserWidget {
	if v == nil || v.GObject == nil {
		return nil
	}

	p := unsafe.Pointer(v.GObject)
	return C.toGtkAppChooserWidget(p)
}

func marshalAppChooserWidget(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapAppChooserWidget(wrapObject(unsafe.Pointer(c))), nil
}

func wrapAppChooserWidget(obj *glib_impl.Object) *appChooserWidget {
	box := wrapBox(obj)
	ac := wrapAppChooser(obj)
	return &appChooserWidget{*box, *ac}
}

// AppChooserWidgetNew() is a wrapper around gtk_app_chooser_widget_new().
func AppChooserWidgetNew(content_type string) (*appChooserWidget, error) {
	cstr := C.CString(content_type)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_app_chooser_widget_new((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapAppChooserWidget(wrapObject(unsafe.Pointer(c))), nil
}

// GetShowDefault() is a wrapper around gtk_app_chooser_widget_get_show_default().
func (v *appChooserWidget) GetShowDefault() bool {
	return gobool(C.gtk_app_chooser_widget_get_show_default(v.native()))
}

// SetShowDefault() is a wrapper around gtk_app_chooser_widget_set_show_default().
func (v *appChooserWidget) SetShowDefault(setting bool) {
	C.gtk_app_chooser_widget_set_show_default(v.native(), gbool(setting))
}

// GetShowRecommended() is a wrapper around gtk_app_chooser_widget_get_show_recommended().
func (v *appChooserWidget) GetShowRecommended() bool {
	return gobool(C.gtk_app_chooser_widget_get_show_recommended(v.native()))
}

// SetShowRecommended() is a wrapper around gtk_app_chooser_widget_set_show_recommended().
func (v *appChooserWidget) SetShowRecommended(setting bool) {
	C.gtk_app_chooser_widget_set_show_recommended(v.native(), gbool(setting))
}

// GetShowFallback() is a wrapper around gtk_app_chooser_widget_get_show_fallback().
func (v *appChooserWidget) GetShowFallback() bool {
	return gobool(C.gtk_app_chooser_widget_get_show_fallback(v.native()))
}

// SetShowFallback() is a wrapper around gtk_app_chooser_widget_set_show_fallback().
func (v *appChooserWidget) SetShowFallback(setting bool) {
	C.gtk_app_chooser_widget_set_show_fallback(v.native(), gbool(setting))
}

// GetShowOther() is a wrapper around gtk_app_chooser_widget_get_show_other().
func (v *appChooserWidget) GetShowOther() bool {
	return gobool(C.gtk_app_chooser_widget_get_show_other(v.native()))
}

// SetShowOther() is a wrapper around gtk_app_chooser_widget_set_show_other().
func (v *appChooserWidget) SetShowOther(setting bool) {
	C.gtk_app_chooser_widget_set_show_other(v.native(), gbool(setting))
}

// GetShowAll() is a wrapper around gtk_app_chooser_widget_get_show_all().
func (v *appChooserWidget) GetShowAll() bool {
	return gobool(C.gtk_app_chooser_widget_get_show_all(v.native()))
}

// SetShowAll() is a wrapper around gtk_app_chooser_widget_set_show_all().
func (v *appChooserWidget) SetShowAll(setting bool) {
	C.gtk_app_chooser_widget_set_show_all(v.native(), gbool(setting))
}

// GetDefaultText() is a wrapper around gtk_app_chooser_widget_get_default_text().
// In case when gtk_app_chooser_widget_get_default_text() returns a nil string,
// GetDefaultText() returns a non-nil error.
func (v *appChooserWidget) GetDefaultText() (string, error) {
	cstr := C.gtk_app_chooser_widget_get_default_text(v.native())
	if cstr == nil {
		return "", nilPtrErr
	}
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString((*C.char)(cstr)), nil
}

// SetDefaultText() is a wrapper around gtk_app_chooser_widget_set_default_text().
func (v *appChooserWidget) SetDefaultText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_app_chooser_widget_set_default_text(v.native(), (*C.gchar)(cstr))
}

/*
 * GtkAppChooserDialog
 */

// AppChooserDialog is a representation of GTK's GtkAppChooserDialog.
type appChooserDialog struct {
	dialog

	// Interfaces
	appChooser
}

// native returns a pointer to the underlying GtkAppChooserButton.
func (v *appChooserDialog) native() *C.GtkAppChooserDialog {
	if v == nil || v.GObject == nil {
		return nil
	}

	p := unsafe.Pointer(v.GObject)
	return C.toGtkAppChooserDialog(p)
}

func marshalAppChooserDialog(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapAppChooserDialog(wrapObject(unsafe.Pointer(c))), nil
}

func wrapAppChooserDialog(obj *glib_impl.Object) *appChooserDialog {
	dialog := wrapDialog(obj)
	ac := wrapAppChooser(obj)
	return &appChooserDialog{*dialog, *ac}
}

// TODO: Uncomment when gio builds successfully
// AppChooserDialogNew() is a wrapper around gtk_app_chooser_dialog_new().
// func AppChooserDialogNew(parent *Window, flags DialogFlags, file *gio.File) (*appChooserDialog, error) {
// 	var gfile *C.GFile
// 	if file != nil {
// 		gfile = (*C.GFile)(unsafe.Pointer(file.Native()))
// 	}
// 	c := C.gtk_app_chooser_dialog_new(parent.native(), C.GtkDialogFlags(flags), gfile)
// 	if c == nil {
// 		return nil, nilPtrErr
// 	}
// 	return wrapAppChooserDialog(wrapObject(unsafe.Pointer(c))), nil
// }

// AppChooserDialogNewForContentType() is a wrapper around gtk_app_chooser_dialog_new_for_content_type().
func AppChooserDialogNewForContentType(parent *window, flags gtk.DialogFlags, content_type string) (*appChooserDialog, error) {
	cstr := C.CString(content_type)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_app_chooser_dialog_new_for_content_type(parent.native(), C.GtkDialogFlags(flags), (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapAppChooserDialog(wrapObject(unsafe.Pointer(c))), nil
}

// GetWidget() is a wrapper around gtk_app_chooser_dialog_get_widget().
func (v *appChooserDialog) GetWidget() gtk.AppChooserWidget {
	c := C.gtk_app_chooser_dialog_get_widget(v.native())
	return wrapAppChooserWidget(wrapObject(unsafe.Pointer(c)))
}

// GetHeading() is a wrapper around gtk_app_chooser_dialog_get_heading().
// In case when gtk_app_chooser_dialog_get_heading() returns a nil string,
// GetHeading() returns a non-nil error.
func (v *appChooserDialog) GetHeading() (string, error) {
	cstr := C.gtk_app_chooser_dialog_get_heading(v.native())
	if cstr == nil {
		return "", nilPtrErr
	}
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString((*C.char)(cstr)), nil
}

// SetHeading() is a wrapper around gtk_app_chooser_dialog_set_heading().
func (v *appChooserDialog) SetHeading(heading string) {
	cstr := C.CString(heading)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_app_chooser_dialog_set_heading(v.native(), (*C.gchar)(cstr))
}

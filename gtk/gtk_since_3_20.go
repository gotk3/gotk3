// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18

// See: https://developer.gnome.org/gtk3/3.20/api-index-3-20.html

package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "gtk_since_3_20.go.h"
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

/*
 * GtkNativeDialog
 */

// NativeDialog is a representation of GTK's GtkNativeDialog.
type NativeDialog struct {
	glib.InitiallyUnowned
}

// native returns a pointer to the underlying GObject as a GtkNativeDialog.
func (v *NativeDialog) native() *C.GtkNativeDialog {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkNativeDialog(p)
}

func wrapNativeDialog(obj *glib.Object) *NativeDialog {
	return &NativeDialog{glib.InitiallyUnowned{obj}}
}

// Run() is a wrapper around gtk_native_dialog_run().
func (v *NativeDialog) Run() int {
	c := C.gtk_native_dialog_run(v.native())
	return int(c)
}

// Destroy() is a wrapper around gtk_native_dialog_destroy().
func (v *NativeDialog) Destroy() {
	C.gtk_native_dialog_destroy(v.native())
}

// SetModal is a wrapper around gtk_native_dialog_set_modal().
func (v *NativeDialog) SetModal(modal bool) {
	C.gtk_native_dialog_set_modal(v.native(), gbool(modal))
}

// GetModal() is a wrapper around gtk_native_dialog_get_modal().
func (v *NativeDialog) GetModal() bool {
	c := C.gtk_native_dialog_get_modal(v.native())
	return gobool(c)
}

// SetTitle is a wrapper around gtk_native_dialog_set_title().
func (v *NativeDialog) SetTitle(title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_native_dialog_set_title(v.native(), (*C.char)(cstr))
}

// GetTitle() is a wrapper around gtk_native_dialog_get_title().
func (v *NativeDialog) GetTitle() (string, error) {
	return stringReturn((*C.gchar)(C.gtk_native_dialog_get_title(v.native())))
}

// SetTransientFor() is a wrapper around gtk_native_dialog_set_transient_for().
func (v *NativeDialog) SetTransientFor(parent IWindow) {
	var pw *C.GtkWindow = nil
	if parent != nil {
		pw = parent.toWindow()
	}
	C.gtk_native_dialog_set_transient_for(v.native(), pw)
}

// GetTransientFor() is a wrapper around gtk_native_dialog_get_transient_for().
func (v *NativeDialog) GetTransientFor() (*Window, error) {
	c := C.gtk_native_dialog_get_transient_for(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapWindow(glib.Take(unsafe.Pointer(c))), nil
}

// GetVisible() is a wrapper around gtk_native_dialog_get_visible().
func (v *NativeDialog) GetVisible() bool {
	c := C.gtk_native_dialog_get_visible(v.native())
	return gobool(c)
}

// Show() is a wrapper around gtk_native_dialog_show().
func (v *NativeDialog) Show() {
	C.gtk_native_dialog_show(v.native())
}

// Hide() is a wrapper around gtk_native_dialog_hide().
func (v *NativeDialog) Hide() {
	C.gtk_native_dialog_hide(v.native())
}

/*
 * GtkFileChooserNative
 */

// FileChooserNativeDialog is a representation of GTK's GtkFileChooserNative.
type FileChooserNativeDialog struct {
	NativeDialog

	// Interfaces
	FileChooser
}

// native returns a pointer to the underlying GObject as a GtkNativeDialog.
func (v *FileChooserNativeDialog) native() *C.GtkFileChooserNative {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkFileChooserNative(p)
}

func wrapFileChooserNativeDialog(obj *glib.Object) *FileChooserNativeDialog {
	fc := wrapFileChooser(obj)
	return &FileChooserNativeDialog{NativeDialog{glib.InitiallyUnowned{obj}}, *fc}
}

// FileChooserNativeDialogNew is a wrapper around gtk_file_chooser_native_new().
func FileChooserNativeDialogNew(
	title string,
	parent *Window,
	action FileChooserAction,
	accept_label string,
	cancel_label string) (*FileChooserNativeDialog, error) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))
	c_accept_label := C.CString(accept_label)
	defer C.free(unsafe.Pointer(c_accept_label))
	c_cancel_label := C.CString(cancel_label)
	defer C.free(unsafe.Pointer(c_cancel_label))
	c := C.gtk_file_chooser_native_new(
		(*C.gchar)(c_title), parent.native(), C.GtkFileChooserAction(action),
		(*C.gchar)(c_accept_label), (*C.gchar)(c_cancel_label))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapFileChooserNativeDialog(obj), nil
}

/*
 * FileChooserNative
 */
func OpenFileChooserNative(title string, parent_window *Window) *string {
	c_title := C.CString(title)

	var native *C.GtkFileChooserNative

	native = C.gtk_file_chooser_native_new((*C.gchar)(c_title),
		parent_window.native(),
		C.GtkFileChooserAction(FILE_CHOOSER_ACTION_OPEN),
		(*C.gchar)(C.CString("_Open")),
		(*C.gchar)(C.CString("_Cancel")))

	p := unsafe.Pointer(unsafe.Pointer(native))
	dlg := C.toGtkNativeDialog(p)
	res := C.gtk_native_dialog_run(dlg)

	if res == C.GTK_RESPONSE_ACCEPT {
		c := C.gtk_file_chooser_get_filename(C.toGtkFileChooser(p))
		s := goString(c)
		defer C.g_free((C.gpointer)(c))

		return &s
	}

	return nil
}

// SetAcceptLabel is a wrapper around gtk_file_chooser_native_set_accept_label().
func (v *FileChooserNativeDialog) SetAcceptLabel(accept_label string) {
	cstr := C.CString(accept_label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_file_chooser_native_set_accept_label(v.native(), (*C.char)(cstr))
}

// GetAcceptLabel() is a wrapper around gtk_file_chooser_native_get_accept_label().
func (v *FileChooserNativeDialog) GetAcceptLabel() (string, error) {
	return stringReturn((*C.gchar)(C.gtk_file_chooser_native_get_accept_label(v.native())))
}

// SetCancelLabel is a wrapper around gtk_file_chooser_native_set_cancel_label().
func (v *FileChooserNativeDialog) SetCancelLabel(cancel_label string) {
	cstr := C.CString(cancel_label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_file_chooser_native_set_cancel_label(v.native(), (*C.char)(cstr))
}

// GetCancelLabel() is a wrapper around gtk_file_chooser_native_get_cancel_label().
func (v *FileChooserNativeDialog) GetCancelLabel() (string, error) {
	return stringReturn((*C.gchar)(C.gtk_file_chooser_native_get_cancel_label(v.native())))
}

func (v *Button) SetColor(color string) {
	rgba := C.GdkRGBA{}
	C.gdk_rgba_parse(&rgba, (*C.gchar)(C.CString(color)))
	C.gtk_widget_override_background_color(v.toWidget(), C.GTK_STATE_FLAG_NORMAL, &rgba)
}

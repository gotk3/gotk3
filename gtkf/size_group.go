package gtkf

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/glibf"
	"github.com/gotk3/gotk3/gtk"
)

func init() {
	gtk.SIZE_GROUP_NONE = C.GTK_SIZE_GROUP_NONE
	gtk.SIZE_GROUP_HORIZONTAL = C.GTK_SIZE_GROUP_HORIZONTAL
	gtk.SIZE_GROUP_VERTICAL = C.GTK_SIZE_GROUP_VERTICAL
	gtk.SIZE_GROUP_BOTH = C.GTK_SIZE_GROUP_BOTH
}

func marshalSizeGroupMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.SizeGroupMode(c), nil
}

/*
 * GtkSizeGroup
 */

// SizeGroup is a representation of GTK's GtkSizeGroup
type sizeGroup struct {
	*glibf.Object
}

// native() returns a pointer to the underlying GtkSizeGroup
func (v *sizeGroup) native() *C.GtkSizeGroup {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSizeGroup(p)
}

func marshalSizeGroup(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapSizeGroup(wrapObject(unsafe.Pointer(c))), nil
}

func wrapSizeGroup(obj *glibf.Object) *sizeGroup {
	return &sizeGroup{obj}
}

// SizeGroupNew is a wrapper around gtk_size_group_new().
func SizeGroupNew(mode gtk.SizeGroupMode) (gtk.SizeGroup, error) {
	c := C.gtk_size_group_new(C.GtkSizeGroupMode(mode))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSizeGroup(wrapObject(unsafe.Pointer(c))), nil
}

func (v *sizeGroup) SetMode(mode gtk.SizeGroupMode) {
	C.gtk_size_group_set_mode(v.native(), C.GtkSizeGroupMode(mode))
}

func (v *sizeGroup) GetMode() gtk.SizeGroupMode {
	return gtk.SizeGroupMode(C.gtk_size_group_get_mode(v.native()))
}
func (v *sizeGroup) SetIgnoreHidden(ignoreHidden bool) {
	C.gtk_size_group_set_ignore_hidden(v.native(), gbool(ignoreHidden))
}

func (v *sizeGroup) GetIgnoreHidden() bool {
	c := C.gtk_size_group_get_ignore_hidden(v.native())
	return gobool(c)
}

func (v *sizeGroup) AddWidget(widget gtk.Widget) {
	C.gtk_size_group_add_widget(v.native(), castToWidget(widget))
}

func (v *sizeGroup) RemoveWidget(widget gtk.Widget) {
	C.gtk_size_group_remove_widget(v.native(), castToWidget(widget))
}

func (v *sizeGroup) GetWidgets() glib.SList {
	c := C.gtk_size_group_get_widgets(v.native())
	if c == nil {
		return nil
	}
	return glibf.WrapSList(uintptr(unsafe.Pointer(c)))
}

// Same copyright and license as the rest of the files in this project
// This file contains code related to GtkStatusIcon

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)
import "github.com/gotk3/gotk3/gdk"

// TODO: GtkStatusIcon * gtk_status_icon_new_from_gicon (GIcon *icon);
// TODO: void gtk_status_icon_set_from_gicon (GtkStatusIcon *status_icon, GIcon *icon);

// TODO: GIcon * gtk_status_icon_get_gicon (GtkStatusIcon *status_icon);

// TODO: void gtk_status_icon_set_screen (GtkStatusIcon *status_icon, GdkScreen *screen);
// TODO: GdkScreen * gtk_status_icon_get_screen (GtkStatusIcon *status_icon);

// TODO: GdkPixbuf * gtk_status_icon_get_pixbuf (GtkStatusIcon *status_icon);

// TODO: void gtk_status_icon_position_menu (GtkMenu *menu, gint *x, gint *y, gboolean *push_in, gpointer user_data);
// TODO: gboolean gtk_status_icon_get_geometry (GtkStatusIcon *status_icon, GdkScreen **screen, GdkRectangle *area, GtkOrientation *orientation);

// StatusIcon is a representation of GTK's GtkStatusIcon.
// Deprecated since 3.14 in favor of notifications
// (no replacement, see https://stackoverflow.com/questions/41917903/gtk-3-statusicon-replacement)
type StatusIcon struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkStatusIcon.
func (v *StatusIcon) native() *C.GtkStatusIcon {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkStatusIcon(p)
}

func marshalStatusIcon(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapStatusIcon(obj), nil
}

func wrapStatusIcon(obj *glib.Object) *StatusIcon {
	return &StatusIcon{obj}
}

// StatusIconNew is a wrapper around gtk_status_icon_new().
func StatusIconNew(str string) (*StatusIcon, error) {
	c := C.gtk_status_icon_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapStatusIcon(obj), nil
}

// StatusIconNewFromFile is a wrapper around gtk_status_icon_new_from_file().
func StatusIconNewFromFile(filename string) (*StatusIcon, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_status_icon_new_from_file((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapStatusIcon(obj), nil
}

// StatusIconNewFromStock is a wrapper around gtk_status_icon_new_from_stock().
// Deprecated since 3.10, use StatusIconNewFromIconName (gtk_status_icon_new_from_icon_name) instead.
func StatusIconNewFromStock(stockId string) (*StatusIcon, error) {
	cstr := C.CString(stockId)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_status_icon_new_from_file((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapStatusIcon(obj), nil
}

// StatusIconNewFromIconName is a wrapper around gtk_status_icon_new_from_icon_name().
func StatusIconNewFromIconName(iconName string) (*StatusIcon, error) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_status_icon_new_from_icon_name((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapStatusIcon(obj), nil
}

// StatusIconNewFromPixbuf is a wrapper around gtk_status_icon_new_from_pixbuf().
func StatusIconNewFromPixbuf(pixbuf *gdk.Pixbuf) (*StatusIcon, error) {
	c := C.gtk_status_icon_new_from_pixbuf(C.toGdkPixbuf(unsafe.Pointer(pixbuf.Native())))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapStatusIcon(obj), nil
}

// SetFromPixbuf is a wrapper around gtk_status_icon_set_from_pixbuf()
func (v *StatusIcon) SetFromPixbuf(pixbuf *gdk.Pixbuf) {
	C.gtk_status_icon_set_from_pixbuf(v.native(), C.toGdkPixbuf(unsafe.Pointer(pixbuf.Native())))
}

// SetFromFile is a wrapper around gtk_status_icon_set_from_file()
func (v *StatusIcon) SetFromFile(filename string) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_from_file(v.native(), (*C.gchar)(cstr))
}

// SetFromStock is a wrapper around gtk_status_icon_set_from_stock()
// Deprecated since 3.10, use SetFromIconName (gtk_status_icon_set_from_icon_name) instead.
func (v *StatusIcon) SetFromStock(stockID string) {
	cstr := C.CString(stockID)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_from_stock(v.native(), (*C.gchar)(cstr))
}

// SetFromIconName is a wrapper around gtk_status_icon_set_from_icon_name()
func (v *StatusIcon) SetFromIconName(iconName string) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_from_icon_name(v.native(), (*C.gchar)(cstr))
}

// GetStorageType is a wrapper around gtk_status_icon_get_storage_type
func (v *StatusIcon) GetStorageType() ImageType {
	c := C.gtk_status_icon_get_storage_type(v.native())
	return ImageType(c)
}

// GetStock is a wrapper around gtk_status_icon_get_stock()
// Deprecated since 3.10, use GetIconName (gtk_status_icon_get_icon_name) instead
func (v *StatusIcon) GetStock() string {
	c := C.gtk_status_icon_get_stock(v.native())
	if c == nil {
		return ""
	}
	return C.GoString((*C.char)(c))
}

// GetIconName is a wrapper around gtk_status_icon_get_icon_name()
func (v *StatusIcon) GetIconName() string {
	c := C.gtk_status_icon_get_icon_name(v.native())
	if c == nil {
		return ""
	}
	return C.GoString((*C.char)(c))
}

// GetSize is a wrapper around gtk_status_icon_get_size()
func (v *StatusIcon) GetSize() int {
	c := C.gtk_status_icon_get_size(v.native())
	return int(c)
}

// SetTooltipText is a wrapper around gtk_status_icon_set_tooltip_text()
func (v *StatusIcon) SetTooltipText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_tooltip_text(v.native(), (*C.gchar)(cstr))
}

// GetTooltipText is a wrapper around gtk_status_icon_get_tooltip_text()
func (v *StatusIcon) GetTooltipText() string {
	c := C.gtk_status_icon_get_tooltip_text(v.native())
	if c == nil {
		return ""
	}
	return C.GoString((*C.char)(c))
}

// SetTooltipMarkup is a wrapper around gtk_status_icon_set_tooltip_markup()
func (v *StatusIcon) SetTooltipMarkup(markup string) {
	cstr := C.CString(markup)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_tooltip_markup(v.native(), (*C.gchar)(cstr))
}

// GetTooltipMarkup is a wrapper around gtk_status_icon_get_tooltip_markup()
func (v *StatusIcon) GetTooltipMarkup() string {
	c := C.gtk_status_icon_get_tooltip_markup(v.native())
	if c == nil {
		return ""
	}
	return C.GoString((*C.char)(c))
}

// SetHasTooltip is a wrapper around gtk_status_icon_set_has_tooltip()
func (v *StatusIcon) SetHasTooltip(hasTooltip bool) {
	C.gtk_status_icon_set_has_tooltip(v.native(), gbool(hasTooltip))
}

// GetHasTooltip is a wrapper around gtk_status_icon_get_has_tooltip()
func (v *StatusIcon) GetHasTooltip() bool {
	c := C.gtk_status_icon_get_has_tooltip(v.native())
	return gobool(c)
}

// SetTitle is a wrapper around gtk_status_icon_set_title()
func (v *StatusIcon) SetTitle(title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_title(v.native(), (*C.gchar)(cstr))
}

// GetTitle is a wrapper around gtk_status_icon_get_title()
func (v *StatusIcon) GetTitle() string {
	c := C.gtk_status_icon_get_title(v.native())
	if c == nil {
		return ""
	}
	return C.GoString((*C.char)(c))
}

// SetName is a wrapper around gtk_status_icon_set_name()
func (v *StatusIcon) SetName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_name(v.native(), (*C.gchar)(cstr))
}

// SetVisible is a wrapper around gtk_status_icon_set_visible()
func (v *StatusIcon) SetVisible(visible bool) {
	C.gtk_status_icon_set_visible(v.native(), gbool(visible))
}

// GetVisible is a wrapper around gtk_status_icon_get_visible()
func (v *StatusIcon) GetVisible() bool {
	c := C.gtk_status_icon_get_visible(v.native())
	return gobool(c)
}

// IsEmbedded is a wrapper around gtk_status_icon_is_embedded()
func (v *StatusIcon) IsEmbedded() bool {
	c := C.gtk_status_icon_is_embedded(v.native())
	return gobool(c)
}

// GetX11WindowID is a wrapper around gtk_status_icon_get_x11_window_id
func (v *StatusIcon) GetX11WindowID() uint32 {
	return uint32(C.gtk_status_icon_get_x11_window_id(v.native()))
}

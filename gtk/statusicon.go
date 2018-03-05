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

// TODO:
// void gtk_status_icon_set_from_pixbuf (GtkStatusIcon *status_icon, GdkPixbuf *pixbuf);
// void gtk_status_icon_set_from_file (GtkStatusIcon *status_icon, const gchar *filename);
// void gtk_status_icon_set_from_stock (GtkStatusIcon *status_icon, const gchar *stock_id);
//    Deprecated since 3.10 Use gtk_status_icon_set_from_icon_name() instead.
// void gtk_status_icon_set_from_icon_name (GtkStatusIcon *status_icon, const gchar *icon_name);

// GtkImageType gtk_status_icon_get_storage_type (GtkStatusIcon *status_icon);

// GdkPixbuf * gtk_status_icon_get_pixbuf (GtkStatusIcon *status_icon);
// const gchar * gtk_status_icon_get_stock (GtkStatusIcon *status_icon);
// const gchar * gtk_status_icon_get_icon_name (GtkStatusIcon *status_icon);
// GIcon * gtk_status_icon_get_gicon (GtkStatusIcon *status_icon);

// gint gtk_status_icon_get_size (GtkStatusIcon *status_icon);

// void gtk_status_icon_set_screen (GtkStatusIcon *status_icon, GdkScreen *screen);
// GdkScreen * gtk_status_icon_get_screen (GtkStatusIcon *status_icon);

// void gtk_status_icon_set_tooltip_text (GtkStatusIcon *status_icon, const gchar *text);
// gchar * gtk_status_icon_get_tooltip_text (GtkStatusIcon *status_icon);

// void gtk_status_icon_set_tooltip_markup (GtkStatusIcon *status_icon, const gchar *markup);
// gchar * gtk_status_icon_get_tooltip_markup (GtkStatusIcon *status_icon);

// void gtk_status_icon_set_has_tooltip (GtkStatusIcon *status_icon, gboolean has_tooltip);
// gboolean gtk_status_icon_get_has_tooltip (GtkStatusIcon *status_icon);

// void gtk_status_icon_set_title (GtkStatusIcon *status_icon, const gchar *title);
// const gchar * gtk_status_icon_get_title (GtkStatusIcon *status_icon);

// void gtk_status_icon_set_name (GtkStatusIcon *status_icon, const gchar *name);

// SetVisible is a wrapper around gtk_status_icon_set_visible()
func (v *StatusIcon) SetVisible(visible bool) {
	C.gtk_status_icon_set_visible(v.native(), gbool(visible))
}

// GetVisible is a wrapper around gtk_status_icon_get_visible()
func (v *StatusIcon) GetVisible() bool {
	c := C.gtk_status_icon_get_visible(v.native())
	return gobool(c)
}

// gboolean gtk_status_icon_is_embedded (GtkStatusIcon *status_icon);

// void gtk_status_icon_position_menu (GtkMenu *menu, gint *x, gint *y, gboolean *push_in, gpointer user_data);
// gboolean gtk_status_icon_get_geometry (GtkStatusIcon *status_icon, GdkScreen **screen, GdkRectangle *area, GtkOrientation *orientation);
// guint32 gtk_status_icon_get_x11_window_id (GtkStatusIcon *status_icon);

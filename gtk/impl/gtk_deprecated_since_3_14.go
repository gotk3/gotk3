// +build gtk_3_6 gtk_3_8 gtk_3_10 gtk_3_12

package impl

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
// #include "gtk_deprecated_since_3_14.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	glib_impl "github.com/gotk3/gotk3/glib/impl"
	"github.com/gotk3/gotk3/gtk"
)

func init() {
	tm := []glib_impl.TypeMarshaler{
		{glib.Type(C.gtk_alignment_get_type()), marshalAlignment},
		{glib.Type(C.gtk_arrow_get_type()), marshalArrow},
		{glib.Type(C.gtk_misc_get_type()), marshalMisc},
		{glib.Type(C.gtk_status_icon_get_type()), marshalStatusIcon},
	}
	glib_impl.RegisterGValueMarshalers(tm)

	//Contribute to casting
	for k, v := range map[string]WrapFn{
		"GtkAlignment":  wrapAlignment,
		"GtkArrow":      wrapArrow,
		"GtkMisc":       wrapMisc,
		"GtkStatusIcon": wrapStatusIcon,
	} {
		WrapMap[k] = v
	}
}

/*
 * deprecated since version 3.14 and should not be used in newly-written code
 */

// ResizeGripIsVisible is a wrapper around
// gtk_window_resize_grip_is_visible().
func (v *window) ResizeGripIsVisible() bool {
	c := C.gtk_window_resize_grip_is_visible(v.native())
	return gobool(c)
}

// SetHasResizeGrip is a wrapper around gtk_window_set_has_resize_grip().
func (v *window) SetHasResizeGrip(setting bool) {
	C.gtk_window_set_has_resize_grip(v.native(), gbool(setting))
}

// GetHasResizeGrip is a wrapper around gtk_window_get_has_resize_grip().
func (v *window) GetHasResizeGrip() bool {
	c := C.gtk_window_get_has_resize_grip(v.native())
	return gobool(c)
}

// Reparent() is a wrapper around gtk_widget_reparent().
func (v *widget) Reparent(newParent gtk.Widget) {
	C.gtk_widget_reparent(v.native(), newParent.(IWidget).toWidget())
}

// GetPadding is a wrapper around gtk_alignment_get_padding().
func (v *alignment) GetPadding() (top, bottom, left, right uint) {
	var ctop, cbottom, cleft, cright C.guint
	C.gtk_alignment_get_padding(v.native(), &ctop, &cbottom, &cleft,
		&cright)
	return uint(ctop), uint(cbottom), uint(cleft), uint(cright)
}

// SetPadding is a wrapper around gtk_alignment_set_padding().
func (v *alignment) SetPadding(top, bottom, left, right uint) {
	C.gtk_alignment_set_padding(v.native(), C.guint(top), C.guint(bottom),
		C.guint(left), C.guint(right))
}

// AlignmentNew is a wrapper around gtk_alignment_new().
func AlignmentNew(xalign, yalign, xscale, yscale float32) (*alignment, error) {
	c := C.gtk_alignment_new(C.gfloat(xalign), C.gfloat(yalign), C.gfloat(xscale),
		C.gfloat(yscale))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAlignment(obj), nil
}

// Set is a wrapper around gtk_alignment_set().
func (v *alignment) Set(xalign, yalign, xscale, yscale float32) {
	C.gtk_alignment_set(v.native(), C.gfloat(xalign), C.gfloat(yalign),
		C.gfloat(xscale), C.gfloat(yscale))
}

/*
 * GtkArrow
 */

// Arrow is a representation of GTK's GtkArrow.
type arrow struct {
	misc
}

// ArrowNew is a wrapper around gtk_arrow_new().
func ArrowNew(arrowType gtk.ArrowType, shadowType gtk.ShadowType) (*arrow, error) {
	c := C.gtk_arrow_new(C.GtkArrowType(arrowType),
		C.GtkShadowType(shadowType))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapArrow(obj), nil
}

// Set is a wrapper around gtk_arrow_set().
func (v *arrow) Set(arrowType gtk.ArrowType, shadowType gtk.ShadowType) {
	C.gtk_arrow_set(v.native(), C.GtkArrowType(arrowType), C.GtkShadowType(shadowType))
}

// SetAlignment() is a wrapper around gtk_button_set_alignment().
func (v *button) SetAlignment(xalign, yalign float32) {
	C.gtk_button_set_alignment(v.native(), (C.gfloat)(xalign),
		(C.gfloat)(yalign))
}

// GetAlignment() is a wrapper around gtk_button_get_alignment().
func (v *button) GetAlignment() (xalign, yalign float32) {
	var x, y C.gfloat
	C.gtk_button_get_alignment(v.native(), &x, &y)
	return float32(x), float32(y)
}

// SetReallocateRedraws is a wrapper around
// gtk_container_set_reallocate_redraws().
func (v *container) SetReallocateRedraws(needsRedraws bool) {
	C.gtk_container_set_reallocate_redraws(v.native(), gbool(needsRedraws))
}

// GetAlignment is a wrapper around gtk_misc_get_alignment().
func (v *misc) GetAlignment() (xAlign, yAlign float32) {
	var x, y C.gfloat
	C.gtk_misc_get_alignment(v.native(), &x, &y)
	return float32(x), float32(y)
}

// SetAlignment is a wrapper around gtk_misc_set_alignment().
func (v *misc) SetAlignment(xAlign, yAlign float32) {
	C.gtk_misc_set_alignment(v.native(), C.gfloat(xAlign), C.gfloat(yAlign))
}

// GetPadding is a wrapper around gtk_misc_get_padding().
func (v *misc) GetPadding() (xpad, ypad int) {
	var x, y C.gint
	C.gtk_misc_get_padding(v.native(), &x, &y)
	return int(x), int(y)
}

// SetPadding is a wrapper around gtk_misc_set_padding().
func (v *misc) SetPadding(xPad, yPad int) {
	C.gtk_misc_set_padding(v.native(), C.gint(xPad), C.gint(yPad))
}

// SetDoubleBuffered is a wrapper around gtk_widget_set_double_buffered().
func (v *widget) SetDoubleBuffered(doubleBuffered bool) {
	C.gtk_widget_set_double_buffered(v.native(), gbool(doubleBuffered))
}

// GetDoubleBuffered is a wrapper around gtk_widget_get_double_buffered().
func (v *widget) GetDoubleBuffered() bool {
	c := C.gtk_widget_get_double_buffered(v.native())
	return gobool(c)
}

/*
 * GtkArrow
 * deprecated since version 3.14
 */
// native returns a pointer to the underlying GtkButton.
func (v *arrow) native() *C.GtkArrow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkArrow(p)
}

func marshalArrow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapArrow(obj), nil
}

func wrapArrow(obj *glib_impl.Object) *arrow {
	return &arrow{misc{widget{glib_impl.InitiallyUnowned{obj}}}}
}

/*
 * GtkAlignment
 * deprecated since version 3.14
 */

type alignment struct {
	bin
}

// native returns a pointer to the underlying GtkAlignment.
func (v *alignment) native() *C.GtkAlignment {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkAlignment(p)
}

func marshalAlignment(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAlignment(obj), nil
}

func wrapAlignment(obj *glib_impl.Object) *alignment {
	return &alignment{bin{container{widget{glib_impl.InitiallyUnowned{obj}}}}}
}

/*
 * GtkStatusIcon
 * deprecated since version 3.14
 */

// StatusIcon is a representation of GTK's GtkStatusIcon
type statusIcon struct {
	*glib_impl.Object
}

func marshalStatusIcon(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapStatusIcon(obj), nil
}

func wrapStatusIcon(obj *glib_impl.Object) *statusIcon {
	return &statusIcon{obj}
}

func (v *statusIcon) native() *C.GtkStatusIcon {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkStatusIcon(p)
}

// StatusIconNew is a wrapper around gtk_status_icon_new()
func StatusIconNew() (*statusIcon, error) {
	c := C.gtk_status_icon_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapStatusIcon(wrapObject(unsafe.Pointer(c))), nil
}

// StatusIconNewFromFile is a wrapper around gtk_status_icon_new_from_file()
func StatusIconNewFromFile(filename string) (*statusIcon, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_status_icon_new_from_file((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapStatusIcon(wrapObject(unsafe.Pointer(c))), nil
}

// StatusIconNewFromIconName is a wrapper around gtk_status_icon_new_from_name()
func StatusIconNewFromIconName(iconName string) (*statusIcon, error) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_status_icon_new_from_icon_name((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapStatusIcon(wrapObject(unsafe.Pointer(c))), nil
}

// SetFromFile is a wrapper around gtk_status_icon_set_from_file()
func (v *statusIcon) SetFromFile(filename string) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_from_file(v.native(), (*C.gchar)(cstr))
}

// SetFromIconName is a wrapper around gtk_status_icon_set_from_icon_name()
func (v *statusIcon) SetFromIconName(iconName string) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_from_icon_name(v.native(), (*C.gchar)(cstr))
}

// GetStorageType is a wrapper around gtk_status_icon_get_storage_type()
func (v *statusIcon) GetStorageType() gtk.ImageType {
	return (ImageType)(C.gtk_status_icon_get_storage_type(v.native()))
}

// SetTooltipText is a wrapper around gtk_status_icon_set_tooltip_text()
func (v *statusIcon) SetTooltipText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_tooltip_text(v.native(), (*C.gchar)(cstr))
}

// GetTooltipText is a wrapper around gtk_status_icon_get_tooltip_text()
func (v *statusIcon) GetTooltipText() string {
	cstr := (*C.char)(C.gtk_status_icon_get_tooltip_text(v.native()))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// SetTooltipMarkup is a wrapper around gtk_status_icon_set_tooltip_markup()
func (v *statusIcon) SetTooltipMarkup(markup string) {
	cstr := (*C.gchar)(C.CString(markup))
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_tooltip_markup(v.native(), cstr)
}

// GetTooltipMarkup is a wrapper around gtk_status_icon_get_tooltip_markup()
func (v *statusIcon) GetTooltipMarkup() string {
	cstr := (*C.char)(C.gtk_status_icon_get_tooltip_markup(v.native()))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// SetHasTooltip is a wrapper around gtk_status_icon_set_has_tooltip()
func (v *statusIcon) SetHasTooltip(hasTooltip bool) {
	C.gtk_status_icon_set_has_tooltip(v.native(), gbool(hasTooltip))
}

// GetTitle is a wrapper around gtk_status_icon_get_title()
func (v *statusIcon) GetTitle() string {
	cstr := (*C.char)(C.gtk_status_icon_get_title(v.native()))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// SetName is a wrapper around gtk_status_icon_set_name()
func (v *statusIcon) SetName(name string) {
	cstr := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_name(v.native(), cstr)
}

// SetVisible is a wrapper around gtk_status_icon_set_visible()
func (v *statusIcon) SetVisible(visible bool) {
	C.gtk_status_icon_set_visible(v.native(), gbool(visible))
}

// GetVisible is a wrapper around gtk_status_icon_get_visible()
func (v *statusIcon) GetVisible() bool {
	return gobool(C.gtk_status_icon_get_visible(v.native()))
}

// IsEmbedded is a wrapper around gtk_status_icon_is_embedded()
func (v *statusIcon) IsEmbedded() bool {
	return gobool(C.gtk_status_icon_is_embedded(v.native()))
}

// GetX11WindowID is a wrapper around gtk_status_icon_get_x11_window_id()
func (v *statusIcon) GetX11WindowID() int {
	return int(C.gtk_status_icon_get_x11_window_id(v.native()))
}

// GetHasTooltip is a wrapper around gtk_status_icon_get_has_tooltip()
func (v *statusIcon) GetHasTooltip() bool {
	return gobool(C.gtk_status_icon_get_has_tooltip(v.native()))
}

// SetTitle is a wrapper around gtk_status_icon_set_title()
func (v *statusIcon) SetTitle(title string) {
	cstr := (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_title(v.native(), cstr)
}

// GetIconName is a wrapper around gtk_status_icon_get_icon_name()
func (v *statusIcon) GetIconName() string {
	cstr := (*C.char)(C.gtk_status_icon_get_icon_name(v.native()))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// GetSize is a wrapper around gtk_status_icon_get_size()
func (v *statusIcon) GetSize() int {
	return int(C.gtk_status_icon_get_size(v.native()))
}

/*
 * GtkMisc
 */

// Misc is a representation of GTK's GtkMisc.
type misc struct {
	widget
}

// native returns a pointer to the underlying GtkMisc.
func (v *misc) native() *C.GtkMisc {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMisc(p)
}

func marshalMisc(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapMisc(obj), nil
}

func wrapMisc(obj *glib_impl.Object) *misc {
	return &misc{widget{glib_impl.InitiallyUnowned{obj}}}
}

/*
 * End deprecated since version 3.14
 */

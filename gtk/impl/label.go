// Same copyright and license as the rest of the files in this project
// This file contains style related functions and structures

package impl

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/pango"

	glib_impl "github.com/gotk3/gotk3/glib/impl"
)

/*
 * GtkLabel
 */

// Label is a representation of GTK's GtkLabel.
type label struct {
	widget
}

// native returns a pointer to the underlying GtkLabel.
func (v *label) native() *C.GtkLabel {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkLabel(p)
}

func marshalLabel(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapLabel(obj), nil
}

func wrapLabel(obj *glib_impl.Object) *label {
	return &label{widget{glib_impl.InitiallyUnowned{obj}}}
}

// LabelNew is a wrapper around gtk_label_new().
func LabelNew(str string) (*label, error) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_label_new((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapLabel(obj), nil
}

// SetText is a wrapper around gtk_label_set_text().
func (v *label) SetText(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_text(v.native(), (*C.gchar)(cstr))
}

// SetMarkup is a wrapper around gtk_label_set_markup().
func (v *label) SetMarkup(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_markup(v.native(), (*C.gchar)(cstr))
}

// SetMarkupWithMnemonic is a wrapper around
// gtk_label_set_markup_with_mnemonic().
func (v *label) SetMarkupWithMnemonic(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_markup_with_mnemonic(v.native(), (*C.gchar)(cstr))
}

// SetPattern is a wrapper around gtk_label_set_pattern().
func (v *label) SetPattern(patern string) {
	cstr := C.CString(patern)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_pattern(v.native(), (*C.gchar)(cstr))
}

// SetJustify is a wrapper around gtk_label_set_justify().
func (v *label) SetJustify(jtype gtk.Justification) {
	C.gtk_label_set_justify(v.native(), C.GtkJustification(jtype))
}

// SetEllipsize is a wrapper around gtk_label_set_ellipsize().
func (v *label) SetEllipsize(mode pango.EllipsizeMode) {
	C.gtk_label_set_ellipsize(v.native(), C.PangoEllipsizeMode(mode))
}

// GetWidthChars is a wrapper around gtk_label_get_width_chars().
func (v *label) GetWidthChars() int {
	c := C.gtk_label_get_width_chars(v.native())
	return int(c)
}

// SetWidthChars is a wrapper around gtk_label_set_width_chars().
func (v *label) SetWidthChars(nChars int) {
	C.gtk_label_set_width_chars(v.native(), C.gint(nChars))
}

// GetMaxWidthChars is a wrapper around gtk_label_get_max_width_chars().
func (v *label) GetMaxWidthChars() int {
	c := C.gtk_label_get_max_width_chars(v.native())
	return int(c)
}

// SetMaxWidthChars is a wrapper around gtk_label_set_max_width_chars().
func (v *label) SetMaxWidthChars(nChars int) {
	C.gtk_label_set_max_width_chars(v.native(), C.gint(nChars))
}

// GetLineWrap is a wrapper around gtk_label_get_line_wrap().
func (v *label) GetLineWrap() bool {
	c := C.gtk_label_get_line_wrap(v.native())
	return gobool(c)
}

// SetLineWrap is a wrapper around gtk_label_set_line_wrap().
func (v *label) SetLineWrap(wrap bool) {
	C.gtk_label_set_line_wrap(v.native(), gbool(wrap))
}

// SetLineWrapMode is a wrapper around gtk_label_set_line_wrap_mode().
func (v *label) SetLineWrapMode(wrapMode pango.WrapMode) {
	C.gtk_label_set_line_wrap_mode(v.native(), C.PangoWrapMode(wrapMode))
}

// GetSelectable is a wrapper around gtk_label_get_selectable().
func (v *label) GetSelectable() bool {
	c := C.gtk_label_get_selectable(v.native())
	return gobool(c)
}

// GetText is a wrapper around gtk_label_get_text().
func (v *label) GetText() (string, error) {
	c := C.gtk_label_get_text(v.native())
	if c == nil {
		return "", nilPtrErr
	}
	return C.GoString((*C.char)(c)), nil
}

// GetJustify is a wrapper around gtk_label_get_justify().
func (v *label) GetJustify() gtk.Justification {
	c := C.gtk_label_get_justify(v.native())
	return gtk.Justification(c)
}

// GetEllipsize is a wrapper around gtk_label_get_ellipsize().
func (v *label) GetEllipsize() pango.EllipsizeMode {
	c := C.gtk_label_get_ellipsize(v.native())
	return pango.EllipsizeMode(c)
}

// GetCurrentUri is a wrapper around gtk_label_get_current_uri().
func (v *label) GetCurrentUri() string {
	c := C.gtk_label_get_current_uri(v.native())
	return C.GoString((*C.char)(c))
}

// GetTrackVisitedLinks is a wrapper around gtk_label_get_track_visited_links().
func (v *label) GetTrackVisitedLinks() bool {
	c := C.gtk_label_get_track_visited_links(v.native())
	return gobool(c)
}

// SetTrackVisitedLinks is a wrapper around gtk_label_set_track_visited_links().
func (v *label) SetTrackVisitedLinks(trackLinks bool) {
	C.gtk_label_set_track_visited_links(v.native(), gbool(trackLinks))
}

// GetAngle is a wrapper around gtk_label_get_angle().
func (v *label) GetAngle() float64 {
	c := C.gtk_label_get_angle(v.native())
	return float64(c)
}

// SetAngle is a wrapper around gtk_label_set_angle().
func (v *label) SetAngle(angle float64) {
	C.gtk_label_set_angle(v.native(), C.gdouble(angle))
}

// GetSelectionBounds is a wrapper around gtk_label_get_selection_bounds().
func (v *label) GetSelectionBounds() (start, end int, nonEmpty bool) {
	var cstart, cend C.gint
	c := C.gtk_label_get_selection_bounds(v.native(), &cstart, &cend)
	return int(cstart), int(cend), gobool(c)
}

// GetSingleLineMode is a wrapper around gtk_label_get_single_line_mode().
func (v *label) GetSingleLineMode() bool {
	c := C.gtk_label_get_single_line_mode(v.native())
	return gobool(c)
}

// SetSingleLineMode is a wrapper around gtk_label_set_single_line_mode().
func (v *label) SetSingleLineMode(mode bool) {
	C.gtk_label_set_single_line_mode(v.native(), gbool(mode))
}

// GetUseMarkup is a wrapper around gtk_label_get_use_markup().
func (v *label) GetUseMarkup() bool {
	c := C.gtk_label_get_use_markup(v.native())
	return gobool(c)
}

// SetUseMarkup is a wrapper around gtk_label_set_use_markup().
func (v *label) SetUseMarkup(use bool) {
	C.gtk_label_set_use_markup(v.native(), gbool(use))
}

// GetUseUnderline is a wrapper around gtk_label_get_use_underline().
func (v *label) GetUseUnderline() bool {
	c := C.gtk_label_get_use_underline(v.native())
	return gobool(c)
}

// SetUseUnderline is a wrapper around gtk_label_set_use_underline().
func (v *label) SetUseUnderline(use bool) {
	C.gtk_label_set_use_underline(v.native(), gbool(use))
}

// LabelNewWithMnemonic is a wrapper around gtk_label_new_with_mnemonic().
func LabelNewWithMnemonic(str string) (*label, error) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_label_new_with_mnemonic((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapLabel(obj), nil
}

// SelectRegion is a wrapper around gtk_label_select_region().
func (v *label) SelectRegion(startOffset, endOffset int) {
	C.gtk_label_select_region(v.native(), C.gint(startOffset),
		C.gint(endOffset))
}

// SetSelectable is a wrapper around gtk_label_set_selectable().
func (v *label) SetSelectable(setting bool) {
	C.gtk_label_set_selectable(v.native(), gbool(setting))
}

// SetLabel is a wrapper around gtk_label_set_label().
func (v *label) SetLabel(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_label(v.native(), (*C.gchar)(cstr))
}

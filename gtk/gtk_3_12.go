// Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
//
// This file originated from: http://opensource.conformal.com/
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// This file includes wrapers for symbols included since GTK 3.12, and
// and should not be included in a build intended to target any older GTK
// versions.  To target an older build, such as 3.10, use
// 'go build -tags gtk_3_10'.  Otherwise, if no build tags are used, GTK 3.12
// is assumed and this file is built.
// +build !gtk_3_6,!gtk_3_8,!gtk_3_10

package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
// #include "gtk_3_12.go.h"
import "C"
import (
	"github.com/conformal/gotk3/glib"
	"runtime"
	"unsafe"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Objects/Interfaces
		{glib.Type(C.gtk_popover_get_type()), marshalPopover},
	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * GtkPopover
 */

// Popover is a representation of GTK's GtkPopover.
type Popover struct {
	Bin
}

// native returns a pointer to the underlying GtkPopover.
func (v *Popover) native() *C.GtkPopover {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkPopover(p)
}

func marshalPopover(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapPopover(obj), nil
}

func wrapPopover(obj *glib.Object) *Popover {
	return &Popover{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// PopoverNew is a wrapper around gtk_popover_new().
func PopoverNew(relativeTo IWidget) (*Popover, error) {
	c := C.gtk_popover_new(relativeTo.toWidget())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	a := wrapPopover(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return a, nil
}

// TODO: gtk_popover_new_from_model
// TODO: gtk_popover_bind_model

// SetRelativeTo is a wrapper around gtk_popover_set_relative_to().
func (v *Popover) SetRelativeTo(relativeTo IWidget) {
	C.gtk_popover_set_relative_to(v.native(), relativeTo.toWidget())
}

// GetRelativeTo is a wrapper around gtk_popover_get_relative_to().
func (v *Popover) GetRelativeTo() (*Widget, error) {
	c := C.gtk_popover_get_relative_to(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// TODO: gtk_popover_set_pointing_to
// TODO: gtk_popover_get_pointing_to

// SetPosition is a wrapper around gtk_popover_set_position().
func (v *Popover) SetPosition(position PositionType) {
	C.gtk_popover_set_position(v.native(), C.GtkPositionType(position))
}

// GetPosition is a wrapper around gtk_popover_get_position().
func (v *Popover) GetPosition() PositionType {
	c := C.gtk_popover_get_position(v.native())
	return PositionType(c)
}

// SetModal is a wrapper around gtk_popover_set_modal().
func (v *Popover) SetModal(modal bool) {
	C.gtk_popover_set_modal(v.native(), gbool(modal))
}

// GetModal is a wrapper around gtk_popover_get_modal().
func (v *Popover) GetModal() bool {
	c := C.gtk_popover_get_modal(v.native())
	return gobool(c)
}

/*
 * GtkWidget
 */

// GetMarginStart is a wrapper around gtk_widget_get_margin_start().
func (v *Widget) GetMarginStart() int {
	c := C.gtk_widget_get_margin_start(v.native())
	return int(c)
}

// SetMarginStart is a wrapper around gtk_widget_set_margin_start().
func (v *Widget) SetMarginStart(margin int) {
	C.gtk_widget_set_margin_start(v.native(), C.gint(margin))
}

// GetMarginEnd is a wrapper around gtk_widget_get_margin_end().
func (v *Widget) GetMarginEnd() int {
	c := C.gtk_widget_get_margin_end(v.native())
	return int(c)
}

// SetMarginEnd is a wrapper around gtk_widget_set_margin_end().
func (v *Widget) SetMarginEnd(margin int) {
	C.gtk_widget_set_margin_end(v.native(), C.gint(margin))
}

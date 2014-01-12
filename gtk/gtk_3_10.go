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

// This file includes wrapers for symbols included since GTK 3.10, and
// and should not be included in a build intended to target any older GTK
// versions.  To target an older build, such as 3.8, use
// 'go build -tags gtk_3_8'.  Otherwise, if no build tags are used, GTK 3.10
// is assumed and this file is built.
// +build !gtk_3_6,!gtk_3_8

package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
// #include "gtk_3_10.go.h"
import "C"
import (
	"github.com/conformal/gotk3/glib"
	"runtime"
	"unsafe"
)

/*
 * Constants
 */

const (
        ALIGN_BASELINE Align = C.GTK_ALIGN_BASELINE
)

/*
 * GtkHeaderBar
 */

type HeaderBar struct {
	Container
}

// Native returns a pointer to the underlying GtkHeaderBar.
func (v *HeaderBar) Native() *C.GtkHeaderBar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkHeaderBar(p)
}

func wrapHeaderBar(obj *glib.Object) *HeaderBar {
	return &HeaderBar{Container{Widget{glib.InitiallyUnowned{obj}}}}
}

// HeaderBarNew is a wrapper around gtk_header_bar_new().
func HeaderBarNew() (*HeaderBar, error) {
	c := C.gtk_header_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	h := wrapHeaderBar(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return h, nil
}

// SetTitle is a wrapper around gtk_header_bar_set_title().
func (v *HeaderBar) SetTitle(title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_header_bar_set_title(v.Native(), (*C.gchar)(cstr))
}

// GetTitle is a wrapper around gtk_header_bar_get_title().
func (v *HeaderBar) GetTitle() string {
	cstr := C.gtk_header_bar_get_title(v.Native())
	return C.GoString((*C.char)(cstr))
}

// SetSubtitle is a wrapper around gtk_header_bar_set_subtitle().
func (v *HeaderBar) SetSubtitle(subtitle string) {
	cstr := C.CString(subtitle)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_header_bar_set_subtitle(v.Native(), (*C.gchar)(cstr))
}

// GetSubtitle is a wrapper around gtk_header_bar_get_subtitle().
func (v *HeaderBar) GetSubtitle() string {
	cstr := C.gtk_header_bar_get_subtitle(v.Native())
	return C.GoString((*C.char)(cstr))
}

// SetCustomTitle is a wrapper around gtk_header_bar_set_custom_title().
func (v *HeaderBar) SetCustomTitle(titleWidget IWidget) {
	C.gtk_header_bar_set_custom_title(v.Native(), titleWidget.toWidget())
}

// GetCustomTitle is a wrapper around gtk_header_bar_get_custom_title().
func (v *HeaderBar) GetCustomTitle() (*Widget, error) {
	c := C.gtk_header_bar_get_custom_title(v.Native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWidget(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w, nil
}

// PackStart is a wrapper around gtk_header_bar_pack_start().
func (v *HeaderBar) PackStart(child IWidget) {
	C.gtk_header_bar_pack_start(v.Native(), child.toWidget())
}

// PackEnd is a wrapper around gtk_header_bar_pack_end().
func (v *HeaderBar) PackEnd(child IWidget) {
	C.gtk_header_bar_pack_end(v.Native(), child.toWidget())
}

// SetShowCloseButton is a wrapper around gtk_header_bar_set_show_close_button().
func (v *HeaderBar) SetShowCloseButton(setting bool) {
	C.gtk_header_bar_set_show_close_button(v.Native(), gbool(setting))
}

// GetShowCloseButton is a wrapper around gtk_header_bar_get_show_close_button().
func (v *HeaderBar) GetShowCloseButton() bool {
	c := C.gtk_header_bar_get_show_close_button(v.Native())
	return gobool(c)
}

/*
 * GtkWindow
 */

// Close is a wrapper around gtk_window_close().
func (v *Window) Close() {
	C.gtk_window_close(v.Native())
}

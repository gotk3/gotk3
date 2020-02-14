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

// This file includes wrapers for symbols included since GTK 3.8, and
// and should not be included in a build intended to target any older GTK
// versions.  To target an older build, such as 3.8, use
// 'go build -tags gtk_3_8'.  Otherwise, if no build tags are used, GTK 3.18
// is assumed and this file is built.
// +build !gtk_3_6

package gtk

// #include <gtk/gtk.h>
// #include "widget_since_3_8.go.h"
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
)

/*
 * GtkWidget
 */

// IsVisible is a wrapper around gtk_widget_is_visible().
func (v *Widget) IsVisible() bool {
	c := C.gtk_widget_is_visible(v.native())
	return gobool(c)
}

// SetOpacity is a wrapper around gtk_widget_set_opacity()
func (v *Widget) SetOpacity(opacity float64) {
	C.gtk_widget_set_opacity(v.native(), C.double(opacity))
}

// GetOpacity is a wrapper around gtk_widget_get_opacity()
func (v *Widget) GetOpacity() float64 {
	c := C.gtk_widget_get_opacity(v.native())
	return float64(c)
}

// GetFrameClock is a wrapper around gtk_widget_get_frame_clock().
func (v *Widget) GetFrameClock() *gdk.FrameClock {
	c := C.gtk_widget_get_frame_clock(v.native())
	return gdk.WrapFrameClock(unsafe.Pointer(c))
}

// AddTickCallback is a wrapper around gtk_widget_add_tick_callback().
func (v *Widget) AddTickCallback(fn TickCallback, userData ...interface{}) int {
	tickCallbackRegistry.Lock()
	id := tickCallbackRegistry.next
	tickCallbackRegistry.next++
	tickCallbackRegistry.m[id] = tickCallbackData{fn: fn, userData: userData}
	tickCallbackRegistry.Unlock()

	return int(C._gtk_widget_add_tick_callback(v.native(), C.gpointer(uintptr(id))))

	// This callback is cleaned up when calling RemoveTickCallback()
}

// RemoveTickCallback is a wrapper around gtk_widget_remove_tick_callback().
func (v *Widget) RemoveTickCallback(id int) {
	C.gtk_widget_remove_tick_callback(v.native(), C.guint(id))

	tickCallbackRegistry.Lock()
	delete(tickCallbackRegistry.m, id)
	tickCallbackRegistry.Unlock()
}

// TODO:
// gtk_widget_register_window().
// gtk_widget_unregister_window().

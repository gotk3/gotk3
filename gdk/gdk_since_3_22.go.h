/*
 * Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
 *
 * This file originated from: http://opensource.conformal.com/
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

//#include <stdlib.h>

extern void goGdkSeatGrabPrepareFunc(GdkSeat *seat, GdkWindow *window,
                                     gpointer user_data);

static GdkGrabStatus _gdk_seat_grab(GdkSeat *seat, GdkWindow *window,
                                    GdkSeatCapabilities caps,
                                    gboolean owner_events, GdkCursor *cursor,
                                    const GdkEvent *event,
                                    gboolean has_prepare_func,
                                    gpointer prepare_func_data) {
  GdkSeatGrabPrepareFunc fn =
      has_prepare_func ? (GdkSeatGrabPrepareFunc)goGdkSeatGrabPrepareFunc
                       : NULL;
  return gdk_seat_grab(seat, window, caps, owner_events, cursor, event, fn,
                       prepare_func_data);
}

// Type Casting
static GdkMonitor *toGdkMonitor(void *p) { return (GDK_MONITOR(p)); }

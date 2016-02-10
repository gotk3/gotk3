/*
 * Copyright (c) 2015- terrak <terrak1975@gmail.com>
 *
 * This file originated from: http://www.terrak.net/
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

package pango

// #cgo pkg-config: pango
// #include <pango/pango.h>
// #include "pango.go.h"
import "C"
import "github.com/gotk3/gotk3/pango/iface"

//	"github.com/andre-hub/gotk3/glib"
//	"github.com/andre-hub/gotk3/cairo"
//	"unsafe"

func init() {
	iface.GRAVITY_SOUTH = C.PANGO_GRAVITY_SOUTH
	iface.GRAVITY_EAST = C.PANGO_GRAVITY_EAST
	iface.GRAVITY_NORTH = C.PANGO_GRAVITY_NORTH
	iface.GRAVITY_WEST = C.PANGO_GRAVITY_WEST
	iface.GRAVITY_AUTO = C.PANGO_GRAVITY_AUTO

	iface.GRAVITY_HINT_NATURAL = C.PANGO_GRAVITY_HINT_NATURAL
	iface.GRAVITY_HINT_STRONG = C.PANGO_GRAVITY_HINT_STRONG
	iface.GRAVITY_HINT_LINE = C.PANGO_GRAVITY_HINT_LINE
}

//double       pango_gravity_to_rotation    (PangoGravity       gravity) G_GNUC_CONST;
func GravityToRotation(gravity iface.Gravity) float64 {
	c := C.pango_gravity_to_rotation((C.PangoGravity)(gravity))
	return float64(c)
}

//PangoGravity pango_gravity_get_for_matrix (const PangoMatrix *matrix) G_GNUC_PURE;

//PangoGravity pango_gravity_get_for_script (PangoScript        script,
//					   PangoGravity       base_gravity,
//					   PangoGravityHint   hint) G_GNUC_CONST;

//PangoGravity pango_gravity_get_for_script_and_width
//					  (PangoScript        script,
//					   gboolean           wide,
//					   PangoGravity       base_gravity,
//					   PangoGravityHint   hint) G_GNUC_CONST;

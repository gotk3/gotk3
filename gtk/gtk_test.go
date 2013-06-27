/*
 * Copyright (c) 2013 Conformal Systems <info@conformal.com>
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

package gtk

import (
	"testing"
)

// TestBoolConvs tests the conversion between Go bools and gboolean
// types.
func TestBoolConvs(t *testing.T) {
	if err := testBoolConvs(); err != nil {
		t.Error(err)
	}
}

// TestConsts tests Go constants against the GTK equivalents.
func TestConsts(t *testing.T) {
	if err := testConsts(); err != nil {
		t.Error(err)
	}
}

// TestBox tests creating and adding widgets to a Box
func TestBox(t *testing.T) {
	vbox, err := BoxNew(ORIENTATION_VERTICAL, 0)
	if err != nil {
		t.Error("Unable to create box")
	}

	vbox.Set("homogeneous", true)
	if vbox.GetHomogeneous() != true {
		t.Error("Could not set or get Box homogeneous property")
	}

	vbox.SetHomogeneous(false)
	if vbox.GetHomogeneous() != false {
		t.Error("Could not set or get Box homogeneous property")
	}

	vbox.Set("spacing", 1)
	if vbox.GetSpacing() != 1 {
		t.Error("Could not set or get Box spacing")
	}

	vbox.SetSpacing(2)
	if vbox.GetSpacing() != 2 {
		t.Error("Could not set or get Box spacing")
	}

	// add a child to start and end
	start, err := LabelNew("Start")
	if err != nil {
		t.Error("Unable to create label")
	}

	end, err := LabelNew("End")
	if err != nil {
		t.Error("Unable to create label")
	}

	vbox.PackStart(start, true, true, 3)
	vbox.PackEnd(end, true, true, 3)
}

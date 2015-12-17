/*
 * Copyright (c) 2015 Axel von Blomberg
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

// Minimalistic mapping of golang slices to GDK GBytes.
// toGBytes constructs a GBytes object holding the identical data of a
// given []byte slice. A reference to the slice is stored in a map until
// the destroy callback of GBytes is called, which removes that reference.

// +build gdk_3_32, !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16
// Exported by gdk_inmem.go
extern void gdkDestroyGBytes(void *data);

static GBytes *toGBytes(void *data, int size){
    if (size <= 0) return 0;
    return g_bytes_new_with_free_func(data, size, gdkDestroyGBytes, data);
}

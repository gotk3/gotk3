// Same copyright and license as the rest of the files in this project

// +build !glib_2_40,!glib_2_42,!glib_2_44

package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
// #include "glib_since_2_44.go.h"
// #include "glib_since_2_46.go.h"
import "C"
import "github.com/gotk3/gotk3/internal/callback"

/*
 * GListStore
 */

// Sort is a wrapper around g_list_store_sort().
func (v *ListStore) Sort(compareFunc CompareDataFunc) {
	C._g_list_store_sort(v.native(), C.gpointer(callback.Assign(compareFunc)))
}

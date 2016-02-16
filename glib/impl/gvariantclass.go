// Same copyright and license as the rest of the files in this project

//GVariant : GVariant — strongly typed value datatype
// https://developer.gnome.org/glib/2.26/glib-GVariant.html

package impl

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "github.com/gotk3/gotk3/glib"

/*
 * GVariantClass
 */

func init() {
	glib.VARIANT_CLASS_BOOLEAN = C.G_VARIANT_CLASS_BOOLEAN         //The GVariant is a boolean.
	glib.VARIANT_CLASS_BYTE = C.G_VARIANT_CLASS_BYTE               //The GVariant is a byte.
	glib.VARIANT_CLASS_INT16 = C.G_VARIANT_CLASS_INT16             //The GVariant is a signed 16 bit integer.
	glib.VARIANT_CLASS_UINT16 = C.G_VARIANT_CLASS_UINT16           //The GVariant is an unsigned 16 bit integer.
	glib.VARIANT_CLASS_INT32 = C.G_VARIANT_CLASS_INT32             //The GVariant is a signed 32 bit integer.
	glib.VARIANT_CLASS_UINT32 = C.G_VARIANT_CLASS_UINT32           //The GVariant is an unsigned 32 bit integer.
	glib.VARIANT_CLASS_INT64 = C.G_VARIANT_CLASS_INT64             //The GVariant is a signed 64 bit integer.
	glib.VARIANT_CLASS_UINT64 = C.G_VARIANT_CLASS_UINT64           //The GVariant is an unsigned 64 bit integer.
	glib.VARIANT_CLASS_HANDLE = C.G_VARIANT_CLASS_HANDLE           //The GVariant is a file handle index.
	glib.VARIANT_CLASS_DOUBLE = C.G_VARIANT_CLASS_DOUBLE           //The GVariant is a double precision floating point value.
	glib.VARIANT_CLASS_STRING = C.G_VARIANT_CLASS_STRING           //The GVariant is a normal string.
	glib.VARIANT_CLASS_OBJECT_PATH = C.G_VARIANT_CLASS_OBJECT_PATH //The GVariant is a D-Bus object path string.
	glib.VARIANT_CLASS_SIGNATURE = C.G_VARIANT_CLASS_SIGNATURE     //The GVariant is a D-Bus signature string.
	glib.VARIANT_CLASS_VARIANT = C.G_VARIANT_CLASS_VARIANT         //The GVariant is a variant.
	glib.VARIANT_CLASS_MAYBE = C.G_VARIANT_CLASS_MAYBE             //The GVariant is a maybe-typed value.
	glib.VARIANT_CLASS_ARRAY = C.G_VARIANT_CLASS_ARRAY             //The GVariant is an array.
	glib.VARIANT_CLASS_TUPLE = C.G_VARIANT_CLASS_TUPLE             //The GVariant is a tuple.
	glib.VARIANT_CLASS_DICT_ENTRY = C.G_VARIANT_CLASS_DICT_ENTRY   //The GVariant is a dictionary entry.
}

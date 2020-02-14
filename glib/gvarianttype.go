// Same copyright and license as the rest of the files in this project

//GVariant : GVariant — strongly typed value datatype
// https://developer.gnome.org/glib/2.26/glib-GVariant.html

package glib

// #include <stdlib.h>
// #include <glib.h>
// #include "gvarianttype.go.h"
import "C"

import (
	"runtime"
	"unsafe"
)

// A VariantType is a wrapper for the GVariantType, which encodes type
// information for GVariants.
type VariantType struct {
	GVariantType *C.GVariantType
}

func (v *VariantType) native() *C.GVariantType {
	if v == nil {
		return nil
	}
	return v.GVariantType
}

// String returns a copy of this VariantType's type string.
func (v *VariantType) String() string {
	ch := C.g_variant_type_dup_string(v.native())
	defer C.g_free(C.gpointer(ch))
	return C.GoString((*C.char)(ch))
}

// newVariantType wraps a native GVariantType.
// Does not create a finalizer.
// Use takeVariantType for instances which need to be freed after use.
func newVariantType(v *C.GVariantType) *VariantType {
	if v == nil {
		return nil
	}
	return &VariantType{v}
}

// takeVariantType wraps a native GVariantType
// and sets up a finalizer to free the instance during GC.
func takeVariantType(v *C.GVariantType) *VariantType {
	if v == nil {
		return nil
	}
	obj := &VariantType{v}
	runtime.SetFinalizer(obj, (*VariantType).Free)
	return obj
}

// Variant types for comparing between them.  Cannot be const because
// they are pointers.
// Note that variant types cannot be compared by value, use VariantTypeEqual() instead.
var (
	VARIANT_TYPE_BOOLEAN           = newVariantType(C._G_VARIANT_TYPE_BOOLEAN)
	VARIANT_TYPE_BYTE              = newVariantType(C._G_VARIANT_TYPE_BYTE)
	VARIANT_TYPE_INT16             = newVariantType(C._G_VARIANT_TYPE_INT16)
	VARIANT_TYPE_UINT16            = newVariantType(C._G_VARIANT_TYPE_UINT16)
	VARIANT_TYPE_INT32             = newVariantType(C._G_VARIANT_TYPE_INT32)
	VARIANT_TYPE_UINT32            = newVariantType(C._G_VARIANT_TYPE_UINT32)
	VARIANT_TYPE_INT64             = newVariantType(C._G_VARIANT_TYPE_INT64)
	VARIANT_TYPE_UINT64            = newVariantType(C._G_VARIANT_TYPE_UINT64)
	VARIANT_TYPE_HANDLE            = newVariantType(C._G_VARIANT_TYPE_HANDLE)
	VARIANT_TYPE_DOUBLE            = newVariantType(C._G_VARIANT_TYPE_DOUBLE)
	VARIANT_TYPE_STRING            = newVariantType(C._G_VARIANT_TYPE_STRING)
	VARIANT_TYPE_OBJECT_PATH       = newVariantType(C._G_VARIANT_TYPE_OBJECT_PATH)
	VARIANT_TYPE_SIGNATURE         = newVariantType(C._G_VARIANT_TYPE_SIGNATURE)
	VARIANT_TYPE_VARIANT           = newVariantType(C._G_VARIANT_TYPE_VARIANT)
	VARIANT_TYPE_ANY               = newVariantType(C._G_VARIANT_TYPE_ANY)
	VARIANT_TYPE_BASIC             = newVariantType(C._G_VARIANT_TYPE_BASIC)
	VARIANT_TYPE_MAYBE             = newVariantType(C._G_VARIANT_TYPE_MAYBE)
	VARIANT_TYPE_ARRAY             = newVariantType(C._G_VARIANT_TYPE_ARRAY)
	VARIANT_TYPE_TUPLE             = newVariantType(C._G_VARIANT_TYPE_TUPLE)
	VARIANT_TYPE_UNIT              = newVariantType(C._G_VARIANT_TYPE_UNIT)
	VARIANT_TYPE_DICT_ENTRY        = newVariantType(C._G_VARIANT_TYPE_DICT_ENTRY)
	VARIANT_TYPE_DICTIONARY        = newVariantType(C._G_VARIANT_TYPE_DICTIONARY)
	VARIANT_TYPE_STRING_ARRAY      = newVariantType(C._G_VARIANT_TYPE_STRING_ARRAY)
	VARIANT_TYPE_OBJECT_PATH_ARRAY = newVariantType(C._G_VARIANT_TYPE_OBJECT_PATH_ARRAY)
	VARIANT_TYPE_BYTESTRING        = newVariantType(C._G_VARIANT_TYPE_BYTESTRING)
	VARIANT_TYPE_BYTESTRING_ARRAY  = newVariantType(C._G_VARIANT_TYPE_BYTESTRING_ARRAY)
	VARIANT_TYPE_VARDICT           = newVariantType(C._G_VARIANT_TYPE_VARDICT)
)

// Free is a wrapper around g_variant_type_free.
// Reference counting is usually handled in the gotk layer,
// most applications should not call this.
func (v *VariantType) Free() {
	C.g_variant_type_free(v.native())
}

// VariantTypeNew is a wrapper around g_variant_type_new.
func VariantTypeNew(typeString string) *VariantType {
	cstr := (*C.gchar)(C.CString(typeString))
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_variant_type_new(cstr)
	return takeVariantType(c)
}

// VariantTypeStringIsValid is a wrapper around g_variant_type_string_is_valid.
func VariantTypeStringIsValid(typeString string) bool {
	cstr := (*C.gchar)(C.CString(typeString))
	defer C.free(unsafe.Pointer(cstr))

	return gobool(C.g_variant_type_string_is_valid(cstr))
}

// VariantTypeEqual is a wrapper around g_variant_type_equal
func VariantTypeEqual(type1, type2 *VariantType) bool {
	return gobool(C.g_variant_type_equal(C.gconstpointer(type1.native()), C.gconstpointer(type2.native())))
}

// IsSubtypeOf is a wrapper around g_variant_type_is_subtype_of
func (v *VariantType) IsSubtypeOf(supertype *VariantType) bool {
	return gobool(C.g_variant_type_is_subtype_of(v.native(), supertype.native()))
}

// TODO:
// g_variant_type_copy
// g_variant_type_string_scan
// g_variant_type_is_definite
// g_variant_type_is_container
// g_variant_type_is_basic
// g_variant_type_is_maybe
// g_variant_type_is_array
// g_variant_type_is_tuple
// g_variant_type_is_dict_entry
// g_variant_type_is_variant
// g_variant_type_hash
// g_variant_type_new_maybe
// g_variant_type_new_array
// g_variant_type_new_tuple
// g_variant_type_new_dict_entry
// g_variant_type_element
// g_variant_type_n_items
// g_variant_type_first
// g_variant_type_next
// g_variant_type_key
// g_variant_type_value

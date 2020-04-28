//GVariant : GVariant — strongly typed value datatype
// https://developer.gnome.org/glib/2.26/glib-GVariant.html

package glib

// #include "gvariant.go.h"
// #include "glib.go.h"
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

/*
 * GVariant
 */

// IVariant is an interface type implemented by Variant and all types which embed
// an Variant.  It is meant to be used as a type for function arguments which
// require GVariants or any subclasses thereof.
type IVariant interface {
	ToGVariant() *C.GVariant
	ToVariant() *Variant
}

// A Variant is a representation of GLib's GVariant.
type Variant struct {
	GVariant *C.GVariant
}

// ToGVariant exposes the underlying *C.GVariant type for this Variant,
// necessary to implement IVariant.
func (v *Variant) ToGVariant() *C.GVariant {
	if v == nil {
		return nil
	}
	return v.native()
}

// ToVariant returns this Variant, necessary to implement IVariant.
func (v *Variant) ToVariant() *Variant {
	return v
}

// native returns a pointer to the underlying GVariant.
func (v *Variant) native() *C.GVariant {
	if v == nil || v.GVariant == nil {
		return nil
	}
	return v.GVariant
}

// Native returns a pointer to the underlying GVariant.
func (v *Variant) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

// newVariant wraps a native GVariant.
// Does NOT handle reference counting! Use takeVariant() to take ownership of values.
func newVariant(p *C.GVariant) *Variant {
	if p == nil {
		return nil
	}
	return &Variant{GVariant: p}
}

// TakeVariant wraps a unsafe.Pointer as a glib.Variant, taking ownership of it.
// This function is exported for visibility in other gotk3 packages and
// is not meant to be used by applications.
func TakeVariant(ptr unsafe.Pointer) *Variant {
	return takeVariant(C.toGVariant(ptr))
}

// takeVariant wraps a native GVariant,
// takes ownership and sets up a finalizer to free the instance during GC.
func takeVariant(p *C.GVariant) *Variant {
	if p == nil {
		return nil
	}
	obj := &Variant{GVariant: p}

	if obj.IsFloating() {
		obj.RefSink()
	} else {
		obj.Ref()
	}

	runtime.SetFinalizer(obj, (*Variant).Unref)
	return obj
}

// IsFloating returns true if the variant has a floating reference count.
// Reference counting is usually handled in the gotk layer,
// most applications should not call this.
func (v *Variant) IsFloating() bool {
	return gobool(C.g_variant_is_floating(v.native()))
}

// Ref is a wrapper around g_variant_ref.
// Reference counting is usually handled in the gotk layer,
// most applications should not need to call this.
func (v *Variant) Ref() {
	C.g_variant_ref(v.native())
}

// RefSink is a wrapper around g_variant_ref_sink.
// Reference counting is usually handled in the gotk layer,
// most applications should not need to call this.
func (v *Variant) RefSink() {
	C.g_variant_ref_sink(v.native())
}

// TakeRef is a wrapper around g_variant_take_ref.
// Reference counting is usually handled in the gotk layer,
// most applications should not need to call this.
func (v *Variant) TakeRef() {
	C.g_variant_take_ref(v.native())
}

// Unref is a wrapper around g_variant_unref.
// Reference counting is usually handled in the gotk layer,
// most applications should not need to call this.
func (v *Variant) Unref() {
	C.g_variant_unref(v.native())
}

// VariantFromInt16 is a wrapper around g_variant_new_int16
func VariantFromInt16(value int16) *Variant {
	return takeVariant(C.g_variant_new_int16(C.gint16(value)))
}

// VariantFromInt32 is a wrapper around g_variant_new_int32
func VariantFromInt32(value int32) *Variant {
	return takeVariant(C.g_variant_new_int32(C.gint32(value)))
}

// VariantFromInt64 is a wrapper around g_variant_new_int64
func VariantFromInt64(value int64) *Variant {
	return takeVariant(C.g_variant_new_int64(C.gint64(value)))
}

// VariantFromByte is a wrapper around g_variant_new_byte
func VariantFromByte(value uint8) *Variant {
	return takeVariant(C.g_variant_new_byte(C.guint8(value)))
}

// VariantFromUint16 is a wrapper around g_variant_new_uint16
func VariantFromUint16(value uint16) *Variant {
	return takeVariant(C.g_variant_new_uint16(C.guint16(value)))
}

// VariantFromUint32 is a wrapper around g_variant_new_uint32
func VariantFromUint32(value uint32) *Variant {
	return takeVariant(C.g_variant_new_uint32(C.guint32(value)))
}

// VariantFromUint64 is a wrapper around g_variant_new_uint64
func VariantFromUint64(value uint64) *Variant {
	return takeVariant(C.g_variant_new_uint64(C.guint64(value)))
}

// VariantFromBoolean is a wrapper around g_variant_new_boolean
func VariantFromBoolean(value bool) *Variant {
	return takeVariant(C.g_variant_new_boolean(gbool(value)))
}

// VariantFromString is a wrapper around g_variant_new_string/g_variant_new_take_string.
// Uses g_variant_new_take_string to reduce memory allocations if possible.
func VariantFromString(value string) *Variant {
	cstr := (*C.gchar)(C.CString(value))
	// g_variant_new_take_string takes owhership of the cstring and will call free() on it when done.
	// Do NOT free this string in this function!
	return takeVariant(C.g_variant_new_take_string(cstr))
}

// VariantFromVariant is a wrapper around g_variant_new_variant.
func VariantFromVariant(value *Variant) *Variant {
	return takeVariant(C.g_variant_new_variant(value.native()))
}

// TypeString returns the g variant type string for this variant.
func (v *Variant) TypeString() string {
	// the string returned from this belongs to GVariant and must not be freed.
	return C.GoString((*C.char)(C.g_variant_get_type_string(v.native())))
}

// IsContainer returns true if the variant is a container and false otherwise.
func (v *Variant) IsContainer() bool {
	return gobool(C.g_variant_is_container(v.native()))
}

// GetBoolean returns the bool value of this variant.
func (v *Variant) GetBoolean() bool {
	return gobool(C.g_variant_get_boolean(v.native()))
}

// GetString is a wrapper around g_variant_get_string.
// It returns the string value of the variant.
func (v *Variant) GetString() string {

	// The string value remains valid as long as the GVariant exists, do NOT free the cstring in this function.
	var len C.gsize
	gc := C.g_variant_get_string(v.native(), &len)

	// This is opposed to g_variant_dup_string, which copies the string.
	// g_variant_dup_string is not implemented,
	// as we copy the string value anyways when converting to a go string.

	return C.GoStringN((*C.char)(gc), (C.int)(len))
}

// GetVariant is a wrapper around g_variant_get_variant.
// It unboxes a nested GVariant.
func (v *Variant) GetVariant() *Variant {
	c := C.g_variant_get_variant(v.native())
	if c == nil {
		return nil
	}
	// The returned value is returned with full ownership transfer,
	// only Unref(), don't Ref().
	obj := newVariant(c)
	runtime.SetFinalizer(obj, (*Variant).Unref)
	return obj
}

// GetStrv returns a slice of strings from this variant.  It wraps
// g_variant_get_strv, but returns copies of the strings instead.
func (v *Variant) GetStrv() []string {
	gstrv := C.g_variant_get_strv(v.native(), nil)
	// we do not own the memory for these strings, so we must not use strfreev
	// but we must free the actual pointer we receive (transfer container).
	// We don't implement g_variant_dup_strv which copies the strings,
	// as we need to copy anyways when converting to go strings.
	c := gstrv
	defer C.g_free(C.gpointer(gstrv))
	var strs []string

	for *c != nil {
		strs = append(strs, C.GoString((*C.char)(*c)))
		c = C.next_gcharptr(c)
	}
	return strs
}

// GetObjv returns a slice of object paths from this variant.  It wraps
// g_variant_get_objv, but returns copies of the strings instead.
func (v *Variant) GetObjv() []string {
	gstrv := C.g_variant_get_objv(v.native(), nil)
	// we do not own the memory for these strings, so we must not use strfreev
	// but we must free the actual pointer we receive (transfer container).
	// We don't implement g_variant_dup_objv which copies the strings,
	// as we need to copy anyways when converting to go strings.
	c := gstrv
	defer C.g_free(C.gpointer(gstrv))
	var strs []string

	for *c != nil {
		strs = append(strs, C.GoString((*C.char)(*c)))
		c = C.next_gcharptr(c)
	}
	return strs
}

// GetInt returns the int64 value of the variant if it is an integer type, and
// an error otherwise.  It wraps variouns `g_variant_get_*` functions dealing
// with integers of different sizes.
func (v *Variant) GetInt() (int64, error) {
	t := v.TypeString()
	var i int64
	switch t {
	case "n":
		i = int64(C.g_variant_get_int16(v.native()))
	case "i":
		i = int64(C.g_variant_get_int32(v.native()))
	case "x":
		i = int64(C.g_variant_get_int64(v.native()))
	default:
		return 0, fmt.Errorf("variant type %s not a signed integer type", t)
	}
	return i, nil
}

// GetUint returns the uint64 value of the variant if it is an integer type, and
// an error otherwise.  It wraps variouns `g_variant_get_*` functions dealing
// with integers of different sizes.
func (v *Variant) GetUint() (uint64, error) {
	t := v.TypeString()
	var i uint64
	switch t {
	case "y":
		i = uint64(C.g_variant_get_byte(v.native()))
	case "q":
		i = uint64(C.g_variant_get_uint16(v.native()))
	case "u":
		i = uint64(C.g_variant_get_uint32(v.native()))
	case "t":
		i = uint64(C.g_variant_get_uint64(v.native()))
	default:
		return 0, fmt.Errorf("variant type %s not an unsigned integer type", t)
	}
	return i, nil
}

// Type returns the VariantType for this variant.
func (v *Variant) Type() *VariantType {
	// The return value is valid for the lifetime of value and must not be freed.
	return newVariantType(C.g_variant_get_type(v.native()))
}

// IsType returns true if the variant's type matches t.
func (v *Variant) IsType(t *VariantType) bool {
	return gobool(C.g_variant_is_of_type(v.native(), t.native()))
}

// String wraps g_variant_print().  It returns a string understood
// by g_variant_parse().
func (v *Variant) String() string {
	gc := C.g_variant_print(v.native(), gbool(false))
	defer C.g_free(C.gpointer(gc))
	return C.GoString((*C.char)(gc))
}

// AnnotatedString wraps g_variant_print(), but returns a type-annotated
// string.
func (v *Variant) AnnotatedString() string {
	gc := C.g_variant_print(v.native(), gbool(true))
	defer C.g_free(C.gpointer(gc))
	return C.GoString((*C.char)(gc))
}

// TODO:
//gint	g_variant_compare ()
//GVariantClass	g_variant_classify ()
//gboolean	g_variant_check_format_string ()
//void	g_variant_get ()
//void	g_variant_get_va ()
//GVariant *	g_variant_new ()
//GVariant *	g_variant_new_va ()
//GVariant *	g_variant_new_handle ()
//GVariant *	g_variant_new_double ()
//GVariant *	g_variant_new_printf ()
//GVariant *	g_variant_new_object_path ()
//gboolean	g_variant_is_object_path ()
//GVariant *	g_variant_new_signature ()
//gboolean	g_variant_is_signature ()
//GVariant *	g_variant_new_strv ()
//GVariant *	g_variant_new_objv ()
//GVariant *	g_variant_new_bytestring ()
//GVariant *	g_variant_new_bytestring_array ()
//guchar	g_variant_get_byte ()
//gint16	g_variant_get_int16 ()
//guint16	g_variant_get_uint16 ()
//gint32	g_variant_get_int32 ()
//guint32	g_variant_get_uint32 ()
//gint64	g_variant_get_int64 ()
//guint64	g_variant_get_uint64 ()
//gint32	g_variant_get_handle ()
//gdouble	g_variant_get_double ()
//const gchar *	g_variant_get_bytestring ()
//gchar *	g_variant_dup_bytestring ()
//const gchar **	g_variant_get_bytestring_array ()
//gchar **	g_variant_dup_bytestring_array ()
//GVariant *	g_variant_new_maybe ()
//GVariant *	g_variant_new_array ()
//GVariant *	g_variant_new_tuple ()
//GVariant *	g_variant_new_dict_entry ()
//GVariant *	g_variant_new_fixed_array ()
//GVariant *	g_variant_get_maybe ()
//gsize	g_variant_n_children ()
//GVariant *	g_variant_get_child_value ()
//void	g_variant_get_child ()
//GVariant *	g_variant_lookup_value ()
//gboolean	g_variant_lookup ()
//gconstpointer	g_variant_get_fixed_array ()
//gsize	g_variant_get_size ()
//gconstpointer	g_variant_get_data ()
//GBytes *	g_variant_get_data_as_bytes ()
//void	g_variant_store ()
//GVariant *	g_variant_new_from_data ()
//GVariant *	g_variant_new_from_bytes ()
//GVariant *	g_variant_byteswap ()
//GVariant *	g_variant_get_normal_form ()
//gboolean	g_variant_is_normal_form ()
//guint	g_variant_hash ()
//gboolean	g_variant_equal ()
//gchar *	g_variant_print ()
//GString *	g_variant_print_string ()
//GVariantIter *	g_variant_iter_copy ()
//void	g_variant_iter_free ()
//gsize	g_variant_iter_init ()
//gsize	g_variant_iter_n_children ()
//GVariantIter *	g_variant_iter_new ()
//GVariant *	g_variant_iter_next_value ()
//gboolean	g_variant_iter_next ()
//gboolean	g_variant_iter_loop ()
//void	g_variant_builder_unref ()
//GVariantBuilder *	g_variant_builder_ref ()
//GVariantBuilder *	g_variant_builder_new ()
//void	g_variant_builder_init ()
//void	g_variant_builder_clear ()
//void	g_variant_builder_add_value ()
//void	g_variant_builder_add ()
//void	g_variant_builder_add_parsed ()
//GVariant *	g_variant_builder_end ()
//void	g_variant_builder_open ()
//void	g_variant_builder_close ()
//void	g_variant_dict_unref ()
//GVariantDict *	g_variant_dict_ref ()
//GVariantDict *	g_variant_dict_new ()
//void	g_variant_dict_init ()
//void	g_variant_dict_clear ()
//gboolean	g_variant_dict_contains ()
//gboolean	g_variant_dict_lookup ()
//GVariant *	g_variant_dict_lookup_value ()
//void	g_variant_dict_insert ()
//void	g_variant_dict_insert_value ()
//gboolean	g_variant_dict_remove ()
//GVariant *	g_variant_dict_end ()
//#define	G_VARIANT_PARSE_ERROR
//GVariant *	g_variant_parse ()
//GVariant *	g_variant_new_parsed_va ()
//GVariant *	g_variant_new_parsed ()
//gchar *	g_variant_parse_error_print_context ()

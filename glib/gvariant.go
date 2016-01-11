//GVariant : GVariant — strongly typed value datatype
// https://developer.gnome.org/glib/2.26/glib-GVariant.html

package glib

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

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

// Variant is a representation of GLib's GVariant.
type Variant struct {
	GVariant *C.GVariant
}

func (v *Variant) ToGVariant() *C.GVariant {
	if v == nil {
		return nil
	}
	return v.native()
}

func (v *Variant) ToVariant() *Variant {
	return v
}

// newVariant creates a new Variant from a GVariant pointer.
func newVariant(p *C.GVariant) *Variant {
	return &Variant{GVariant: p}
}

func VariantFromUnsafePointer(p unsafe.Pointer) *Variant {
	return &Variant{C.toGVariant(p)}
}

// native returns a pointer to the underlying GVariant.
func (v *Variant) native() *C.GVariant {
	if v == nil || v.GVariant == nil {
		return nil
	}
	p := unsafe.Pointer(v.GVariant)
	return C.toGVariant(p)
}

// Native returns a pointer to the underlying GVariant.
func (v *Variant) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

//void	g_variant_unref ()
//GVariant *	g_variant_ref ()
//GVariant *	g_variant_ref_sink ()
//gboolean	g_variant_is_floating ()
//GVariant *	g_variant_take_ref ()
//const GVariantType *	g_variant_get_type ()
//const gchar *	g_variant_get_type_string ()
//gboolean	g_variant_is_of_type ()
//gboolean	g_variant_is_container ()
//gint	g_variant_compare ()
//GVariantClass	g_variant_classify ()
//gboolean	g_variant_check_format_string ()
//void	g_variant_get ()
//void	g_variant_get_va ()
//GVariant *	g_variant_new ()
//GVariant *	g_variant_new_va ()
//GVariant *	g_variant_new_boolean ()
//GVariant *	g_variant_new_byte ()
//GVariant *	g_variant_new_int16 ()
//GVariant *	g_variant_new_uint16 ()
//GVariant *	g_variant_new_int32 ()
//GVariant *	g_variant_new_uint32 ()
//GVariant *	g_variant_new_int64 ()
//GVariant *	g_variant_new_uint64 ()
//GVariant *	g_variant_new_handle ()
//GVariant *	g_variant_new_double ()
//GVariant *	g_variant_new_string ()
//GVariant *	g_variant_new_take_string ()
//GVariant *	g_variant_new_printf ()
//GVariant *	g_variant_new_object_path ()
//gboolean	g_variant_is_object_path ()
//GVariant *	g_variant_new_signature ()
//gboolean	g_variant_is_signature ()
//GVariant *	g_variant_new_variant ()
//GVariant *	g_variant_new_strv ()
//GVariant *	g_variant_new_objv ()
//GVariant *	g_variant_new_bytestring ()
//GVariant *	g_variant_new_bytestring_array ()
//gboolean	g_variant_get_boolean ()
//guchar	g_variant_get_byte ()
//gint16	g_variant_get_int16 ()
//guint16	g_variant_get_uint16 ()
//gint32	g_variant_get_int32 ()
//guint32	g_variant_get_uint32 ()
//gint64	g_variant_get_int64 ()
//guint64	g_variant_get_uint64 ()
//gint32	g_variant_get_handle ()
//gdouble	g_variant_get_double ()
//const gchar *	g_variant_get_string ()
//gchar *	g_variant_dup_string ()
//GVariant *	g_variant_get_variant ()
//const gchar **	g_variant_get_strv ()
//gchar **	g_variant_dup_strv ()
//const gchar **	g_variant_get_objv ()
//gchar **	g_variant_dup_objv ()
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

/*
 * GVariantClass
 */

type VariantClass int

const (
	VARIANT_CLASS_BOOLEAN     VariantClass = C.G_VARIANT_CLASS_BOOLEAN     //The GVariant is a boolean.
	VARIANT_CLASS_BYTE        VariantClass = C.G_VARIANT_CLASS_BYTE        //The GVariant is a byte.
	VARIANT_CLASS_INT16       VariantClass = C.G_VARIANT_CLASS_INT16       //The GVariant is a signed 16 bit integer.
	VARIANT_CLASS_UINT16      VariantClass = C.G_VARIANT_CLASS_UINT16      //The GVariant is an unsigned 16 bit integer.
	VARIANT_CLASS_INT32       VariantClass = C.G_VARIANT_CLASS_INT32       //The GVariant is a signed 32 bit integer.
	VARIANT_CLASS_UINT32      VariantClass = C.G_VARIANT_CLASS_UINT32      //The GVariant is an unsigned 32 bit integer.
	VARIANT_CLASS_INT64       VariantClass = C.G_VARIANT_CLASS_INT64       //The GVariant is a signed 64 bit integer.
	VARIANT_CLASS_UINT64      VariantClass = C.G_VARIANT_CLASS_UINT64      //The GVariant is an unsigned 64 bit integer.
	VARIANT_CLASS_HANDLE      VariantClass = C.G_VARIANT_CLASS_HANDLE      //The GVariant is a file handle index.
	VARIANT_CLASS_DOUBLE      VariantClass = C.G_VARIANT_CLASS_DOUBLE      //The GVariant is a double precision floating point value.
	VARIANT_CLASS_STRING      VariantClass = C.G_VARIANT_CLASS_STRING      //The GVariant is a normal string.
	VARIANT_CLASS_OBJECT_PATH VariantClass = C.G_VARIANT_CLASS_OBJECT_PATH //The GVariant is a D-Bus object path string.
	VARIANT_CLASS_SIGNATURE   VariantClass = C.G_VARIANT_CLASS_SIGNATURE   //The GVariant is a D-Bus signature string.
	VARIANT_CLASS_VARIANT     VariantClass = C.G_VARIANT_CLASS_VARIANT     //The GVariant is a variant.
	VARIANT_CLASS_MAYBE       VariantClass = C.G_VARIANT_CLASS_MAYBE       //The GVariant is a maybe-typed value.
	VARIANT_CLASS_ARRAY       VariantClass = C.G_VARIANT_CLASS_ARRAY       //The GVariant is an array.
	VARIANT_CLASS_TUPLE       VariantClass = C.G_VARIANT_CLASS_TUPLE       //The GVariant is a tuple.
	VARIANT_CLASS_DICT_ENTRY  VariantClass = C.G_VARIANT_CLASS_DICT_ENTRY  //The GVariant is a dictionary entry.
)

/*
 * GVariantIter
 */

// VariantIter is a representation of GLib's GVariantIter.
type VariantIter struct {
	GVariantIter *C.GVariantIter
}

func (v *VariantIter) toGVariantIter() *C.GVariantIter {
	if v == nil {
		return nil
	}
	return v.native()
}

func (v *VariantIter) toVariantIter() *VariantIter {
	return v
}

// newVariantIter creates a new VariantIter from a GVariantIter pointer.
func newVariantIter(p *C.GVariantIter) *VariantIter {
	return &VariantIter{GVariantIter: p}
}

// native returns a pointer to the underlying GVariantIter.
func (v *VariantIter) native() *C.GVariantIter {
	if v == nil || v.GVariantIter == nil {
		return nil
	}
	p := unsafe.Pointer(v.GVariantIter)
	return C.toGVariantIter(p)
}

// Native returns a pointer to the underlying GVariantIter.
func (v *VariantIter) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

/*
 * GVariantBuilder
 */

// VariantBuilder is a representation of GLib's VariantBuilder.
type VariantBuilder struct {
	GVariantBuilder *C.GVariantBuilder
}

func (v *VariantBuilder) toGVariantBuilder() *C.GVariantBuilder {
	if v == nil {
		return nil
	}
	return v.native()
}

func (v *VariantBuilder) toVariantBuilder() *VariantBuilder {
	return v
}

// newVariantBuilder creates a new VariantBuilder from a GVariantBuilder pointer.
func newVariantBuilder(p *C.GVariantBuilder) *VariantBuilder {
	return &VariantBuilder{GVariantBuilder: p}
}

// native returns a pointer to the underlying GVariantBuilder.
func (v *VariantBuilder) native() *C.GVariantBuilder {
	if v == nil || v.GVariantBuilder == nil {
		return nil
	}
	p := unsafe.Pointer(v.GVariantBuilder)
	return C.toGVariantBuilder(p)
}

// Native returns a pointer to the underlying GVariantBuilder.
func (v *VariantBuilder) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

/*
 * GVariantDict
 */

// VariantDict is a representation of GLib's VariantDict.
type VariantDict struct {
	GVariantDict *C.GVariantDict
}

func (v *VariantDict) toGVariantDict() *C.GVariantDict {
	if v == nil {
		return nil
	}
	return v.native()
}

func (v *VariantDict) toVariantDict() *VariantDict {
	return v
}

// newVariantDict creates a new VariantDict from a GVariantDict pointer.
func newVariantDict(p *C.GVariantDict) *VariantDict {
	return &VariantDict{GVariantDict: p}
}

// native returns a pointer to the underlying GVariantDict.
func (v *VariantDict) native() *C.GVariantDict {
	if v == nil || v.GVariantDict == nil {
		return nil
	}
	p := unsafe.Pointer(v.GVariantDict)
	return C.toGVariantDict(p)
}

// Native returns a pointer to the underlying GVariantDict.
func (v *VariantDict) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

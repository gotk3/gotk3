package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

// SettingsSchemaSource is a representation of GSettingsSchemaSource.
type SettingsSchemaSource struct {
	source *C.GSettingsSchemaSource
}

func wrapSettingsSchemaSource(obj *C.GSettingsSchemaSource) *SettingsSchemaSource {
	if obj == nil {
		return nil
	}
	return &SettingsSchemaSource{obj}
}

func (v *SettingsSchemaSource) Native() uintptr {
	return uintptr(unsafe.Pointer(v.source))
}

func (v *SettingsSchemaSource) native() *C.GSettingsSchemaSource {
	if v == nil || v.source == nil {
		return nil
	}
	return v.source
}

// SettingsSchemaSourceGetDefault is a wrapper around g_settings_schema_source_get_default().
func SettingsSchemaSourceGetDefault() *SettingsSchemaSource {
	// transfer none
	return wrapSettingsSchemaSource(C.g_settings_schema_source_get_default())
}

// Ref() is a wrapper around g_settings_schema_source_ref().
func (v *SettingsSchemaSource) Ref() *SettingsSchemaSource {
	v = wrapSettingsSchemaSource(C.g_settings_schema_source_ref(v.native()))
}

// Unref() is a wrapper around g_settings_schema_source_unref().
func (v *SettingsSchemaSource) Unref() {
	C.g_settings_schema_source_unref(v.native())
}

// SettingsSchemaSourceNewFromDirectory() is a wrapper around g_settings_schema_source_new_from_directory().
func SettingsSchemaSourceNewFromDirectory(dir string, parent *SettingsSchemaSource, trusted bool) *SettingsSchemaSource {
	cstr := (*C.gchar)(C.CString(dir))
	defer C.free(unsafe.Pointer(cstr))

	schemaSource := wrapSettingsSchemaSource(C.g_settings_schema_source_new_from_directory(cstr, parent.native(), gbool(trusted), nil))
	if schemaSource == nil {
		return nil
	}

	// g_settings_schema_source_new_from_directory sets ref counter to 1. So only Unref() via finalizer.
	runtime.SetFinalizer(schemaSource, (*SettingsSchemaSource).Unref)

	return schemaSource
}

// Lookup() is a wrapper around g_settings_schema_source_lookup().
func (v *SettingsSchemaSource) Lookup(schema string, recursive bool) *SettingsSchema {
	cstr := (*C.gchar)(C.CString(schema))
	defer C.free(unsafe.Pointer(cstr))

	schema := wrapSettingsSchema(C.g_settings_schema_source_lookup(v.native(), cstr, gbool(recursive)))
	if schema == nil {
		return nil
	}

	// transfer full -> i.e. don't Ref(), but ensure Unref() via finalizer
	runtime.SetFinalizer(schema, (*SettingsSchema).Unref)

	return schema
}

// ListSchemas is a wrapper around 	g_settings_schema_source_list_schemas().
func (v *SettingsSchemaSource) ListSchemas(recursive bool) (nonReolcatable, relocatable []string) {
	var nonRel, rel **C.gchar
	C.g_settings_schema_source_list_schemas(v.native(), gbool(recursive), &nonRel, &rel)
	return toGoStringArray(nonRel), toGoStringArray(rel)
}

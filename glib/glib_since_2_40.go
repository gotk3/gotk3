// Same copyright and license as the rest of the files in this project

// +build !glib_2_34

package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
// #include "glib_since_2_40.go.h"
import "C"
import "unsafe"

/*
 * Application
 */

// MarkBusy is a wrapper around g_application_mark_busy().
func (v *Application) MarkBusy() {
	C.g_application_mark_busy(v.native())
}

// UnmarkBusy is a wrapper around g_application_unmark_busy().
func (v *Application) UnmarkBusy() {
	C.g_application_unmark_busy(v.native())
}

// SendNotification is a wrapper around g_application_send_notification().
func (v *Application) SendNotification(id string, notification *Notification) {
	cstr1 := (*C.gchar)(C.CString(id))
	defer C.free(unsafe.Pointer(cstr1))

	C.g_application_send_notification(v.native(), cstr1, notification.native())
}

// WithdrawNotification is a wrapper around g_application_withdraw_notification().
func (v *Application) WithdrawNotification(id string) {
	cstr1 := (*C.gchar)(C.CString(id))
	defer C.free(unsafe.Pointer(cstr1))

	C.g_application_withdraw_notification(v.native(), cstr1)
}

/*
 * Variant
 */
// VariantFromString is a wrapper around g_variant_new_string/g_variant_new_take_string.
// Uses g_variant_new_take_string to reduce memory allocations if possible.
func VariantFromString(value string) *Variant {
	cstr := (*C.gchar)(C.CString(value))
	// g_variant_new_take_string takes owhership of the cstring and will call free() on it when done.
	// Do NOT free this string in this function!
	return takeVariant(C.g_variant_new_take_string(cstr))
}

/*
 * Binding
 */
// Explicitly releases the binding between the source and the target property
// expressed by Binding
func (v *Binding) Unbind() {
	C.g_binding_unbind(v.native())
}

/*
 * Menu
 */
// Predefined attribute names for GMenu
var (
	MENU_ATTRIBUTE_ACTION_NAMESPACE string = C.G_MENU_ATTRIBUTE_ACTION_NAMESPACE
	MENU_ATTRIBUTE_ICON             string = C.G_MENU_ATTRIBUTE_ICON
)

// RemoveAll is a wrapper around g_menu_remove_all().
func (v *Menu) RemoveAll() {
	C.g_menu_remove_all(v.native())
}

/*
 * SettingsSchema
 */
// HasKey() is a wrapper around g_settings_schema_has_key().
func (v *SettingsSchema) HasKey(v1 string) bool {
	cstr := (*C.gchar)(C.CString(v1))
	defer C.free(unsafe.Pointer(cstr))

	return gobool(C.g_settings_schema_has_key(v.native(), cstr))
}

/*
 * SettingsSchemaSource
 */
// ListSchemas is a wrapper around 	g_settings_schema_source_list_schemas().
func (v *SettingsSchemaSource) ListSchemas(recursive bool) (nonReolcatable, relocatable []string) {
	var nonRel, rel **C.gchar
	C.g_settings_schema_source_list_schemas(v.native(), gbool(recursive), &nonRel, &rel)
	return toGoStringArray(nonRel), toGoStringArray(rel)
}

// Same copyright and license as the rest of the files in this project

// +build !glib_2_40

package glib

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

/*
 * Notification
 */

// NotificationPriority is a representation of GLib's GNotificationPriority.
type NotificationPriority int

const (
	NOTIFICATION_PRIORITY_NORMAL NotificationPriority = C.G_NOTIFICATION_PRIORITY_NORMAL
	NOTIFICATION_PRIORITY_LOW    NotificationPriority = C.G_NOTIFICATION_PRIORITY_LOW
	NOTIFICATION_PRIORITY_HIGH   NotificationPriority = C.G_NOTIFICATION_PRIORITY_HIGH
	NOTIFICATION_PRIORITY_URGENT NotificationPriority = C.G_NOTIFICATION_PRIORITY_URGENT
)

// SetPriority is a wrapper around g_notification_set_priority().
func (v *Notification) SetPriority(prio NotificationPriority) {
	C.g_notification_set_priority(v.native(), C.GNotificationPriority(prio))
}

/*
 * Application
 */

// GetResourceBasePath is a wrapper around g_application_get_resource_base_path().
func (v *Application) GetResourceBasePath() string {
	c := C.g_application_get_resource_base_path(v.native())

	return C.GoString((*C.char)(c))
}

// SetResourceBasePath is a wrapper around g_application_set_resource_base_path().
func (v *Application) SetResourceBasePath(bp string) {
	cstr1 := (*C.gchar)(C.CString(bp))
	defer C.free(unsafe.Pointer(cstr1))

	C.g_application_set_resource_base_path(v.native(), cstr1)
}

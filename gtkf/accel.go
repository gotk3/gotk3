// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

package gtkf

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	glib_impl "github.com/gotk3/gotk3/glibf"
	"github.com/gotk3/gotk3/gtk"
)

func init() {
	gtk.ACCEL_VISIBLE = C.GTK_ACCEL_VISIBLE
	gtk.ACCEL_LOCKED = C.GTK_ACCEL_LOCKED
	gtk.ACCEL_MASK = C.GTK_ACCEL_MASK
}

func marshalAccelFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return gtk.AccelFlags(c), nil
}

// AcceleratorName is a wrapper around gtk_accelerator_name().
func AcceleratorName(key uint, mods gdk.ModifierType) string {
	c := C.gtk_accelerator_name(C.guint(key), C.GdkModifierType(mods))
	defer C.free(unsafe.Pointer(c))
	return C.GoString((*C.char)(c))
}

// AcceleratorValid is a wrapper around gtk_accelerator_valid().
func AcceleratorValid(key uint, mods gdk.ModifierType) bool {
	return gobool(C.gtk_accelerator_valid(C.guint(key), C.GdkModifierType(mods)))
}

// AcceleratorGetDefaultModMask is a wrapper around gtk_accelerator_get_default_mod_mask().
func AcceleratorGetDefaultModMask() gdk.ModifierType {
	return gdk.ModifierType(C.gtk_accelerator_get_default_mod_mask())
}

// AcceleratorParse is a wrapper around gtk_accelerator_parse().
func AcceleratorParse(acc string) (key uint, mods gdk.ModifierType) {
	cstr := C.CString(acc)
	defer C.free(unsafe.Pointer(cstr))

	k := C.guint(0)
	m := C.GdkModifierType(0)

	C.gtk_accelerator_parse((*C.gchar)(cstr), &k, &m)
	return uint(k), gdk.ModifierType(m)
}

// AcceleratorGetLabel is a wrapper around gtk_accelerator_get_label().
func AcceleratorGetLabel(key uint, mods gdk.ModifierType) string {
	c := C.gtk_accelerator_get_label(C.guint(key), C.GdkModifierType(mods))
	defer C.free(unsafe.Pointer(c))
	return C.GoString((*C.char)(c))
}

// AcceleratorSetDefaultModMask is a wrapper around gtk_accelerator_set_default_mod_mask().
func AcceleratorSetDefaultModMask(mods gdk.ModifierType) {
	C.gtk_accelerator_set_default_mod_mask(C.GdkModifierType(mods))
}

/*
 * GtkAccelGroup
 */

// AccelGroup is a representation of GTK's GtkAccelGroup.
type accelGroup struct {
	*glib_impl.Object
}

// native returns a pointer to the underlying GtkAccelGroup.
func (v *accelGroup) native() *C.GtkAccelGroup {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkAccelGroup(p)
}

func marshalAccelGroup(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAccelGroup(obj), nil
}

func wrapAccelGroup(obj *glib_impl.Object) *accelGroup {
	return &accelGroup{obj}
}

// AccelGroup is a wrapper around gtk_accel_group_new().
func AccelGroupNew() (*accelGroup, error) {
	c := C.gtk_accel_group_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAccelGroup(obj), nil
}

// Connect is a wrapper around gtk_accel_group_connect().
func (v *accelGroup) Connect2(key uint, mods gdk.ModifierType, flags gtk.AccelFlags, f interface{}) {
	closure, _ := glib_impl.ClosureNew(f)
	cl := (*C.struct__GClosure)(unsafe.Pointer(closure))
	C.gtk_accel_group_connect(
		v.native(),
		C.guint(key),
		C.GdkModifierType(mods),
		C.GtkAccelFlags(flags),
		cl)
}

// ConnectByPath is a wrapper around gtk_accel_group_connect_by_path().
func (v *accelGroup) ConnectByPath(path string, f interface{}) {
	closure, _ := glib_impl.ClosureNew(f)
	cl := (*C.struct__GClosure)(unsafe.Pointer(closure))

	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_accel_group_connect_by_path(
		v.native(),
		(*C.gchar)(cstr),
		cl)
}

// Disconnect is a wrapper around gtk_accel_group_disconnect().
func (v *accelGroup) Disconnect(f interface{}) {
	closure, _ := glib_impl.ClosureNew(f)
	cl := (*C.struct__GClosure)(unsafe.Pointer(closure))
	C.gtk_accel_group_disconnect(v.native(), cl)
}

// DisconnectKey is a wrapper around gtk_accel_group_disconnect_key().
func (v *accelGroup) DisconnectKey(key uint, mods gdk.ModifierType) {
	C.gtk_accel_group_disconnect_key(v.native(), C.guint(key), C.GdkModifierType(mods))
}

// Lock is a wrapper around gtk_accel_group_lock().
func (v *accelGroup) Lock() {
	C.gtk_accel_group_lock(v.native())
}

// Unlock is a wrapper around gtk_accel_group_unlock().
func (v *accelGroup) Unlock() {
	C.gtk_accel_group_unlock(v.native())
}

// IsLocked is a wrapper around gtk_accel_group_get_is_locked().
func (v *accelGroup) IsLocked() bool {
	return gobool(C.gtk_accel_group_get_is_locked(v.native()))
}

// AccelGroupFromClosure is a wrapper around gtk_accel_group_from_accel_closure().
func AccelGroupFromClosure(f interface{}) *accelGroup {
	closure, _ := glib_impl.ClosureNew(f)
	cl := (*C.struct__GClosure)(unsafe.Pointer(closure))
	c := C.gtk_accel_group_from_accel_closure(cl)
	if c == nil {
		return nil
	}
	return wrapAccelGroup(wrapObject(unsafe.Pointer(c)))
}

// GetModifierMask is a wrapper around gtk_accel_group_get_modifier_mask().
func (v *accelGroup) GetModifierMask() gdk.ModifierType {
	return gdk.ModifierType(C.gtk_accel_group_get_modifier_mask(v.native()))
}

// AccelGroupsActivate is a wrapper around gtk_accel_groups_activate().
func AccelGroupsActivate(obj *glib_impl.Object, key uint, mods gdk.ModifierType) bool {
	return gobool(C.gtk_accel_groups_activate((*C.GObject)(unsafe.Pointer(obj.Native())), C.guint(key), C.GdkModifierType(mods)))
}

// Activate is a wrapper around gtk_accel_group_activate().
func (v *accelGroup) Activate(quark glib.Quark, acceleratable glib.Object, key uint, mods gdk.ModifierType) bool {
	return gobool(C.gtk_accel_group_activate(v.native(), C.GQuark(quark), (*C.GObject)(unsafe.Pointer(glib_impl.CastToObject(acceleratable).Native())), C.guint(key), C.GdkModifierType(mods)))
}

// AccelGroupsFromObject is a wrapper around gtk_accel_groups_from_object().
func AccelGroupsFromObject(obj *glib_impl.Object) *glib_impl.SList {
	res := C.gtk_accel_groups_from_object((*C.GObject)(unsafe.Pointer(obj.Native())))
	if res == nil {
		return nil
	}
	return (*glib_impl.SList)(unsafe.Pointer(res))
}

/*
 * GtkAccelMap
 */

// AccelMap is a representation of GTK's GtkAccelMap.
type accelMap struct {
	*glib_impl.Object
}

// native returns a pointer to the underlying GtkAccelMap.
func (v *accelMap) native() *C.GtkAccelMap {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkAccelMap(p)
}

func marshalAccelMap(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapAccelMap(obj), nil
}

func wrapAccelMap(obj *glib_impl.Object) *accelMap {
	return &accelMap{obj}
}

// AccelMapAddEntry is a wrapper around gtk_accel_map_add_entry().
func AccelMapAddEntry(path string, key uint, mods gdk.ModifierType) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_accel_map_add_entry((*C.gchar)(cstr), C.guint(key), C.GdkModifierType(mods))
}

type accelKey struct {
	key   uint
	mods  gdk.ModifierType
	flags uint16
}

func (v *accelKey) native() *C.struct__GtkAccelKey {
	if v == nil {
		return nil
	}

	var val C.struct__GtkAccelKey
	val.accel_key = C.guint(v.key)
	val.accel_mods = C.GdkModifierType(v.mods)
	val.accel_flags = v.flags
	return &val
}

func wrapAccelKey(obj *C.struct__GtkAccelKey) *accelKey {
	var v accelKey

	v.key = uint(obj.accel_key)
	v.mods = gdk.ModifierType(obj.accel_mods)
	v.flags = uint16(obj.accel_flags)

	return &v
}

// AccelMapLookupEntry is a wrapper around gtk_accel_map_lookup_entry().
func AccelMapLookupEntry(path string) *accelKey {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	var v *C.struct__GtkAccelKey

	C.gtk_accel_map_lookup_entry((*C.gchar)(cstr), v)
	return wrapAccelKey(v)
}

// AccelMapChangeEntry is a wrapper around gtk_accel_map_change_entry().
func AccelMapChangeEntry(path string, key uint, mods gdk.ModifierType, replace bool) bool {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	return gobool(C.gtk_accel_map_change_entry((*C.gchar)(cstr), C.guint(key), C.GdkModifierType(mods), gbool(replace)))
}

// AccelMapLoad is a wrapper around gtk_accel_map_load().
func AccelMapLoad(fileName string) {
	cstr := C.CString(fileName)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_accel_map_load((*C.gchar)(cstr))
}

// AccelMapSave is a wrapper around gtk_accel_map_save().
func AccelMapSave(fileName string) {
	cstr := C.CString(fileName)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_accel_map_save((*C.gchar)(cstr))
}

// AccelMapLoadFD is a wrapper around gtk_accel_map_load_fd().
func AccelMapLoadFD(fd int) {
	C.gtk_accel_map_load_fd(C.gint(fd))
}

// AccelMapSaveFD is a wrapper around gtk_accel_map_save_fd().
func AccelMapSaveFD(fd int) {
	C.gtk_accel_map_save_fd(C.gint(fd))
}

// AccelMapAddFilter is a wrapper around gtk_accel_map_add_filter().
func AccelMapAddFilter(filter string) {
	cstr := C.CString(filter)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_accel_map_add_filter((*C.gchar)(cstr))
}

// AccelMapGet is a wrapper around gtk_accel_map_get().
func AccelMapGet() *accelMap {
	c := C.gtk_accel_map_get()
	if c == nil {
		return nil
	}
	return wrapAccelMap(wrapObject(unsafe.Pointer(c)))
}

// AccelMapLockPath is a wrapper around gtk_accel_map_lock_path().
func AccelMapLockPath(path string) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_accel_map_lock_path((*C.gchar)(cstr))
}

// AccelMapUnlockPath is a wrapper around gtk_accel_map_unlock_path().
func AccelMapUnlockPath(path string) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_accel_map_unlock_path((*C.gchar)(cstr))
}

// SetAccelGroup is a wrapper around gtk_menu_set_accel_group().
func (v *menu) SetAccelGroup(accelGroup gtk.AccelGroup) {
	C.gtk_menu_set_accel_group(v.native(), castToAccelGroup(accelGroup).native())
}

// GetAccelGroup is a wrapper around gtk_menu_get_accel_group().
func (v *menu) GetAccelGroup() gtk.AccelGroup {
	c := C.gtk_menu_get_accel_group(v.native())
	if c == nil {
		return nil
	}
	return wrapAccelGroup(wrapObject(unsafe.Pointer(c)))
}

// SetAccelPath is a wrapper around gtk_menu_set_accel_path().
func (v *menu) SetAccelPath(path string) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_menu_set_accel_path(v.native(), (*C.gchar)(cstr))
}

// GetAccelPath is a wrapper around gtk_menu_get_accel_path().
func (v *menu) GetAccelPath() string {
	c := C.gtk_menu_get_accel_path(v.native())
	return C.GoString((*C.char)(c))
}

// SetAccelPath is a wrapper around gtk_menu_item_set_accel_path().
func (v *menuItem) SetAccelPath(path string) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_menu_item_set_accel_path(v.native(), (*C.gchar)(cstr))
}

// GetAccelPath is a wrapper around gtk_menu_item_get_accel_path().
func (v *menuItem) GetAccelPath() string {
	c := C.gtk_menu_item_get_accel_path(v.native())
	return C.GoString((*C.char)(c))
}

// AddAccelerator is a wrapper around gtk_widget_add_accelerator().
func (v *widget) AddAccelerator(signal string, group gtk.AccelGroup, key uint, mods gdk.ModifierType, flags gtk.AccelFlags) {
	csignal := (*C.gchar)(C.CString(signal))
	defer C.free(unsafe.Pointer(csignal))

	C.gtk_widget_add_accelerator(v.native(),
		csignal,
		castToAccelGroup(group).native(),
		C.guint(key),
		C.GdkModifierType(mods),
		C.GtkAccelFlags(flags))
}

// RemoveAccelerator is a wrapper around gtk_widget_remove_accelerator().
func (v *widget) RemoveAccelerator(group gtk.AccelGroup, key uint, mods gdk.ModifierType) bool {
	return gobool(C.gtk_widget_remove_accelerator(v.native(),
		castToAccelGroup(group).native(),
		C.guint(key),
		C.GdkModifierType(mods)))
}

// SetAccelPath is a wrapper around gtk_widget_set_accel_path().
func (v *widget) SetAccelPath2(path string, group gtk.AccelGroup) {
	cstr := (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_widget_set_accel_path(v.native(), cstr, castToAccelGroup(group).native())
}

// CanActivateAccel is a wrapper around gtk_widget_can_activate_accel().
func (v *widget) CanActivateAccel(signalId uint) bool {
	return gobool(C.gtk_widget_can_activate_accel(v.native(), C.guint(signalId)))
}

// AddAccelGroup() is a wrapper around gtk_window_add_accel_group().
func (v *window) AddAccelGroup(accelGroup gtk.AccelGroup) {
	C.gtk_window_add_accel_group(v.native(), castToAccelGroup(accelGroup).native())
}

// RemoveAccelGroup() is a wrapper around gtk_window_add_accel_group().
func (v *window) RemoveAccelGroup(accelGroup gtk.AccelGroup) {
	C.gtk_window_remove_accel_group(v.native(), castToAccelGroup(accelGroup).native())
}

// These three functions are for system level access - thus not as high priority to implement
// TODO: void 	gtk_accelerator_parse_with_keycode ()
// TODO: gchar * 	gtk_accelerator_name_with_keycode ()
// TODO: gchar * 	gtk_accelerator_get_label_with_keycode ()

// TODO: GtkAccelKey * 	gtk_accel_group_find ()   - this function uses a function type - I don't know how to represent it in cgo
// TODO: gtk_accel_map_foreach_unfiltered  - can't be done without a function type
// TODO: gtk_accel_map_foreach  - can't be done without a function type

// TODO: gtk_accel_map_load_scanner
// TODO: gtk_widget_list_accel_closures

package webkit

// #cgo pkg-config: webkit2gtk-4.0
// #include <webkit2/webkit2.h>
// #include "webkit_2_6.go.h"
import "C"

import (
	//"errors"
	"runtime"
	"unsafe"

	//"github.com/andre-hub/gotk3/gtk"
	"github.com/andre-hub/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.webkit_user_content_injected_frames_get_type()), marshalUserContentInjectedFrames},
		{glib.Type(C.webkit_user_script_injection_time_get_type()), marshalUserScriptInjectionTime},
		{glib.Type(C.webkit_user_style_level_get_type()), marshalUserStyleLevel},

		// Objects/Interfaces
		{glib.Type(C.webkit_user_content_manager_get_type()), marshalUserContentManager},

		// Boxed
		{glib.Type(C.webkit_user_script_get_type()), marshalUserScript},
		{glib.Type(C.webkit_user_style_sheet_get_type()), marshalUserStyleSheet},
	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * Constants
 */

// UserContentInjectedFrames is a wrapper around WebKitUserContentInjectedFrames.
type UserContentInjectedFrames int

const (
	USER_CONTENT_INJECT_ALL_FRAMES UserContentInjectedFrames = C.WEBKIT_USER_CONTENT_INJECT_ALL_FRAMES
	USER_CONTENT_INJECT_TOP_FRAME  UserContentInjectedFrames = C.WEBKIT_USER_CONTENT_INJECT_TOP_FRAME
)

func marshalUserContentInjectedFrames(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return UserContentInjectedFrames(c), nil
}

// UserScriptInjectionTime is a wrapper around WebKitUserScriptInjectionTime.
type UserScriptInjectionTime int

const (
	USER_SCRIPT_INJECT_AT_DOCUMENT_START UserScriptInjectionTime = C.WEBKIT_USER_SCRIPT_INJECT_AT_DOCUMENT_START
	USER_SCRIPT_INJECT_AT_DOCUMENT_END   UserScriptInjectionTime = C.WEBKIT_USER_SCRIPT_INJECT_AT_DOCUMENT_END
)

func marshalUserScriptInjectionTime(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return UserScriptInjectionTime(c), nil
}

// UserStyleLevel is a wrapper around WebKitUserStyleLevel.
type UserStyleLevel int

const (
	USER_STYLE_LEVEL_USER   UserStyleLevel = C.WEBKIT_USER_STYLE_LEVEL_USER
	USER_STYLE_LEVEL_AUTHOR UserStyleLevel = C.WEBKIT_USER_STYLE_LEVEL_AUTHOR
)

func marshalUserStyleLevel(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return UserStyleLevel(c), nil
}

/*
 * WebKitUserContentManager
 */

type UserContentManager struct {
	*glib.Object
}

// native returns a pointer to the underlying WebKitUserContentManager.
func (v *UserContentManager) native() *C.WebKitUserContentManager {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toWebKitUserContentManager(p)
}

func marshalUserContentManager(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapUserContentManager(obj), nil
}

func wrapUserContentManager(obj *glib.Object) *UserContentManager {
	return &UserContentManager{obj}
}

// UserContentManagerNew() is a wrapper around webkit_user_content_manager_new().
func UserContentManagerNew() (*UserContentManager, error) {
	c := C.webkit_user_content_manager_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapUserContentManager(obj)
	b.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return b, nil
}

// AddScript() is a wrapper around webkit_user_content_manager_add_script_sheet().
func (v *UserContentManager) AddScript(script *UserScript) {
	C.webkit_user_content_manager_add_script(v.native(), script.native())
}

// AddStyleSheet() is a wrapper around webkit_user_content_manager_add_style_sheet().
func (v *UserContentManager) AddStyleSheet(sheet *UserStyleSheet) {
	C.webkit_user_content_manager_add_style_sheet(v.native(), sheet.native())
}

// RemoveAllScripts() is a wrapper around webkit_user_content_manager_remove_all_scripts().
func (v *UserContentManager) RemoveAllScripts() {
	C.webkit_user_content_manager_remove_all_scripts(v.native())
}

// RemoveAllStyleSheets() is a wrapper around webkit_user_content_manager_remove_all_style_sheets().
func (v *UserContentManager) RemoveAllStyleSheets() {
	C.webkit_user_content_manager_remove_all_style_sheets(v.native())
}

/*
 * WebKitWebView
 */

// WebViewNewWithSettings() is a wrapper around webkit_web_view_new_with_settings().
func WebViewNewWithSettings(settings *Settings) (*WebView, error) {
	c := C.webkit_web_view_new_with_settings(settings.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapWebView(obj)
	b.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return b, nil
}

// WebViewNewWithUserContentManager() is a wrapper around webkit_web_view_new_with_user_content_manager().
func WebViewNewWithUserContentManager(manager *UserContentManager) (*WebView, error) {
	c := C.webkit_web_view_new_with_user_content_manager(manager.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapWebView(obj)
	b.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return b, nil
}

// LoadBytes() is a wrapper around webkit_web_view_load_bytes().
func (v *WebView) LoadBytes(bytes []byte, mimeType, encoding, baseUri string) {
	gbytes := C.g_bytes_new(C.gconstpointer(&bytes[0]), C.gsize(len(bytes)))
	defer C.g_bytes_unref(gbytes)
	cmimeType := C.CString(mimeType)
	defer C.free(unsafe.Pointer(cmimeType))
	cencoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(cencoding))
	cbaseUri := C.CString(baseUri)
	defer C.free(unsafe.Pointer(cbaseUri))
	C.webkit_web_view_load_bytes(v.native(), gbytes, (*C.gchar)(cmimeType), (*C.gchar)(cencoding), (*C.gchar)(cbaseUri))
}

// GetUserContentManager() is a wrapper around webkit_web_view_get_user_content_manager().
func (v *WebView) GetUserContentManager() *UserContentManager {
	c := C.webkit_web_view_get_user_content_manager(v.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapUserContentManager(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w
}

/*
 * WebKitUserScript
 */

type UserScript struct {
	WebKitUserScript *C.WebKitUserScript
}

// native returns a pointer to the underlying WebKitUserScript.
func (v *UserScript) native() *C.WebKitUserScript {
	if v == nil {
		return nil
	}
	return v.WebKitUserScript
}

func marshalUserScript(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*UserScript)(unsafe.Pointer(c)), nil
}

// UserScriptNew() is a wrapper around webkit_user_script_new().
func UserScriptNew(source string, injectedFrames UserContentInjectedFrames, injectionTime UserScriptInjectionTime, whitelist, blacklist []string) (*UserScript, error) {
	csource := C.CString(source)
	defer C.free(unsafe.Pointer(csource))
	cwhite, cwhiteFree := gstringArray(whitelist)
	for _, i := range cwhiteFree {
		defer C.free(unsafe.Pointer(i))
	}
	cblack, cblackFree := gstringArray(blacklist)
	for _, i := range cblackFree {
		defer C.free(unsafe.Pointer(i))
	}

	c := C.webkit_user_script_new((*C.gchar)(csource), C.WebKitUserContentInjectedFrames(injectedFrames), C.WebKitUserScriptInjectionTime(injectionTime), cwhite, cblack)
	if c == nil {
		return nil, nilPtrErr
	}
	b := &UserScript{c}
	runtime.SetFinalizer(b, (*UserScript).Unref)
	return b, nil
}

// Unref is a wrapper around webkit_user_script_unref(). This is for internal use.
func (v *UserScript) Unref() {
	C.webkit_user_script_unref(v.native())
}

// RefSink is a wrapper around webkit_user_script_ref(). This is for internal use.
func (v *UserScript) Ref() {
	C.webkit_user_script_ref(v.native())
}

/*
 * WebKitUserStyleSheet
 */

type UserStyleSheet struct {
	WebKitUserStyleSheet *C.WebKitUserStyleSheet
}

// native returns a pointer to the underlying WebKitUserStyleSheet.
func (v *UserStyleSheet) native() *C.WebKitUserStyleSheet {
	if v == nil {
		return nil
	}
	return v.WebKitUserStyleSheet
}

func marshalUserStyleSheet(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*UserStyleSheet)(unsafe.Pointer(c)), nil
}

// UserStyleSheetNew() is a wrapper around webkit_user_style_sheet_new().
func UserStyleSheetNew(source string, injectedFrames UserContentInjectedFrames, level UserStyleLevel, whitelist, blacklist []string) (*UserStyleSheet, error) {
	csource := C.CString(source)
	defer C.free(unsafe.Pointer(csource))
	cwhite, _ := gstringArray(whitelist)
	/*for _, i := range cwhiteFree {
		defer C.free(unsafe.Pointer(i))
	}*/
	cblack, _ := gstringArray(blacklist)
	/*for _, i := range cblackFree {
		defer C.free(unsafe.Pointer(i))
	}*/

	c := C.webkit_user_style_sheet_new((*C.gchar)(csource), C.WebKitUserContentInjectedFrames(injectedFrames), C.WebKitUserStyleLevel(level), cwhite, cblack)
	if c == nil {
		return nil, nilPtrErr
	}
	b := &UserStyleSheet{c}
	runtime.SetFinalizer(b, (*UserStyleSheet).Unref)
	return b, nil
}

// Unref is a wrapper around webkit_user_script_unref(). This is for internal use.
func (v *UserStyleSheet) Unref() {
	C.webkit_user_style_sheet_unref(v.native())
}

// RefSink is a wrapper around webkit_user_script_ref(). This is for internal use.
func (v *UserStyleSheet) Ref() {
	C.webkit_user_style_sheet_ref(v.native())
}

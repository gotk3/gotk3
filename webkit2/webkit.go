package webkit

// #cgo pkg-config: webkit2gtk-4.0
// #include <webkit2/webkit2.h>
// #include "webkit.go.h"
import "C"

import (
	"errors"
	"runtime"
	"unsafe"

	"github.com/andre-hub/gotk3/glib"
	"github.com/andre-hub/gotk3/gtk"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.webkit_tls_errors_policy_get_type()), marshalTLSErrorsPolicy},
		{glib.Type(C.webkit_load_event_get_type()), marshalLoadEvent},

		// Objects/Interfaces
		{glib.Type(C.webkit_user_content_manager_get_type()), marshalUserContentManager},
		{glib.Type(C.webkit_settings_get_type()), marshalSettings},
		{glib.Type(C.webkit_web_context_get_type()), marshalWebContext},
		{glib.Type(C.webkit_web_view_get_type()), marshalWebView},
		{glib.Type(C.webkit_window_properties_get_type()), marshalWindowProperties},

		// Boxed
	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * Type conversions
 */

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}
func gobool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

func gstringArray(lang []string) (gstringArray **C.gchar, toBeFreed []*C.gchar) {
	clang := make([]*C.gchar, len(lang))
	for i, l := range lang {
		cstr := C.CString(l)
		clang[i] = (*C.gchar)(cstr)
	}
	var t *C.gchar
	cclang := append(clang, t)
	return &cclang[0], clang
}

/*
 * Unexported vars
 */

var nilPtrErr = errors.New("cgo returned unexpected nil pointer")

/*
 * Constants
 */

// Misc constants
const (
	// Version numbers from headers used at compile time, not the library linked at runtime.
	MAJOR_VERSION = C._WEBKIT_MAJOR_VERSION
	MINOR_VERSION = C._WEBKIT_MINOR_VERSION
	MICRO_VERSION = C._WEBKIT_MICRO_VERSION
)

// LoadEvent is a representation of WebKitLoadEvent
type LoadEvent int

const (
	LOAD_STARTED    LoadEvent = C.WEBKIT_LOAD_STARTED
	LOAD_REDIRECTED LoadEvent = C.WEBKIT_LOAD_REDIRECTED
	LOAD_COMMITTED  LoadEvent = C.WEBKIT_LOAD_COMMITTED
	LOAD_FINISHED   LoadEvent = C.WEBKIT_LOAD_FINISHED
)

func marshalLoadEvent(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return LoadEvent(c), nil
}

// TLSErrorsPolicy is a representation of WebKitTLSErrorsPolicy
type TLSErrorsPolicy int

const (
	TLS_ERRORS_POLICY_IGNORE TLSErrorsPolicy = C.WEBKIT_TLS_ERRORS_POLICY_IGNORE
	TLS_ERRORS_POLICY_FAIL   TLSErrorsPolicy = C.WEBKIT_TLS_ERRORS_POLICY_FAIL
)

func marshalTLSErrorsPolicy(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return TLSErrorsPolicy(c), nil
}

/*
 * Misc
 */

// CheckVersion() is a wrapper around WEBKIT_CHECK_VERSION().
func CheckVersion(major, minor, micro uint) bool {
	return gobool(C._WEBKIT_CHECK_VERSION(C.uint(major), C.uint(minor), C.uint(micro)))
}

// GetMajorVersion() is a wrapper around webkit_get_major_version().
func GetMajorVersion() uint {
	return uint(C.webkit_get_major_version())
}

// GetMinorVersion() is a wrapper around webkit_get_minor_version().
func GetMinorVersion() uint {
	return uint(C.webkit_get_minor_version())
}

// GetMicroVersion() is a wrapper around webkit_get_micro_version().
func GetMicroVersion() uint {
	return uint(C.webkit_get_micro_version())
}

/*
 * WebKitSettings
 */

type Settings struct {
	*glib.Object
}

// native returns a pointer to the underlying WebKitSettings.
func (v *Settings) native() *C.WebKitSettings {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toWebKitSettings(p)
}

func marshalSettings(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapSettings(obj), nil
}

func wrapSettings(obj *glib.Object) *Settings {
	return &Settings{obj}
}

// SettingsNew() is a wrapper around webkit_web_view_new().
func SettingsNew() (*Settings, error) {
	c := C.webkit_settings_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapSettings(obj)
	b.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return b, nil
}

/*
 * WebKitWebContext
 */

type WebContext struct {
	*glib.Object
}

// native returns a pointer to the underlying WebKitWebContext.
func (v *WebContext) native() *C.WebKitWebContext {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toWebKitWebContext(p)
}

func marshalWebContext(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapWebContext(obj), nil
}

func wrapWebContext(obj *glib.Object) *WebContext {
	return &WebContext{obj}
}

// WebContextGetDefault() is a wrapper around webkit_web_context_get_default().
func WebContextGetDefault() (*WebContext, error) {
	c := C.webkit_web_context_get_default()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapWebContext(obj)
	b.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return b, nil
}

// ClearCache() is a wrapper around webkit_web_context_clear_cache().
func (v *WebContext) ClearCache() {
	C.webkit_web_context_clear_cache(v.native())
}

// GetFaviconDatabaseDirectory() is a wrapper around webkit_web_context_get_favicon_database_directory().
func (v *WebContext) GetFaviconDatabaseDirectory() string {
	c := C.webkit_web_context_get_favicon_database_directory(v.native())
	return C.GoString((*C.char)(c))
}

// GetTLSErrorsPolicy() is a wrapper around webkit_web_context_get_tls_errors_policy().
func (v *WebContext) GetTLSErrorsPolicy() TLSErrorsPolicy {
	return TLSErrorsPolicy(C.webkit_web_context_get_tls_errors_policy(v.native()))
}

// PrefetchDNS() is a wrapper around webkit_web_context_prefetch_dns().
func (v *WebContext) PrefetchDNS(hostname string) {
	c := C.CString(hostname)
	defer C.free(unsafe.Pointer(c))
	C.webkit_web_context_prefetch_dns(v.native(), (*C.gchar)(c))
}

/*// GetSpellCheckingEnabled() is a wrapper around webkit_web_context_get_spell_checking_enabled().
func (v *WebContext) GetSpellCheckingEnabled() bool {
	return gobool(C.webkit_web_context_get_spell_checking_enabled(v.native()))
}

// GetSpellCheckingLanguages() is a wrapper around webkit_web_context_get_spell_checking_languages().
func (v *WebContext) GetSpellCheckingLanguages() []string {
	c := C.webkit_web_context_get_spell_checking_languages(v.native())
	fmt.Println("c", c)
	//lang := make([]string, 0)

	charTmp := C.gchar(0)
	charPtrSize := unsafe.Sizeof(&charTmp)
	carr := make([]*C.gchar, 0)
	for i := uintptr(0); i < charPtrSize; i += charPtrSize {
		fmt.Println("PTR:", uintptr(unsafe.Pointer(c)) + i)
		ptr := *(**C.gchar)(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + i))
		carr = append(carr, ptr)
	}

	lang := make([]string, len(carr))
	for i, t := range carr {
		lang[i] = C.GoString((*C.char)(t))
	}
	fmt.Print()

	return lang
}*/

// SetDiskCacheDirectory() is a wrapper around webkit_web_context_set_disk_cache_directory().
func (v *WebContext) SetDiskCacheDirectory(dir string) {
	c := C.CString(dir)
	defer C.free(unsafe.Pointer(c))
	C.webkit_web_context_set_disk_cache_directory(v.native(), (*C.gchar)(c))
}

// SetFaviconDatabaseDirectory() is a wrapper around webkit_web_context_set_favicon_database_directory().
func (v *WebContext) SetFaviconDatabaseDirectory(path string) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	C.webkit_web_context_set_favicon_database_directory(v.native(), (*C.gchar)(cpath))
}

// SetTLSErrorsPolicy() is a wrapper around webkit_web_context_set_tls_errors_policy().
func (v *WebContext) SetTLSErrorsPolicy(policy TLSErrorsPolicy) {
	C.webkit_web_context_set_tls_errors_policy(v.native(), C.WebKitTLSErrorsPolicy(policy))
}

/*// SetSpellCheckingEnabled() is a wrapper around webkit_web_context_set_spell_checking_enabled().
func (v *WebContext) SetSpellCheckingEnabled(enabled bool) {
	C.webkit_web_context_set_spell_checking_enabled(v.native(), gbool(enabled))
}

// SetSpellCheckingLanguages() is a wrapper around webkit_web_context_set_spell_checking_languages().
func (v *WebContext) SetSpellCheckingLanguages(lang []string) {
	clang := make([]*C.gchar, len(lang))
	for i, l := range lang {
		cstr := C.CString(l)
		defer C.free(unsafe.Pointer(cstr))
		clang[i] = (*C.gchar)(cstr)
	}
	var t *C.gchar
	clang = append(clang, t)
	C.webkit_web_context_set_spell_checking_languages(v.native(), &clang[0])
}*/

/*
 * WebKitWebView
 */

type WebView struct {
	gtk.Container
}

// native returns a pointer to the underlying WebKitWebView.
func (v *WebView) native() *C.WebKitWebView {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toWebKitWebView(p)
}

func marshalWebView(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapWebView(obj), nil
}

func wrapWebView(obj *glib.Object) *WebView {
	return &WebView{gtk.Container{gtk.Widget{glib.InitiallyUnowned{obj}}}}
}

// WebViewNew() is a wrapper around webkit_web_view_new().
func WebViewNew() (*WebView, error) {
	c := C.webkit_web_view_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapWebView(obj)
	b.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return b, nil
}

// WebViewNewWithContext() is a wrapper around webkit_web_view_new_with_context().
func WebViewNewWithContext(context *WebContext) (*WebView, error) {
	c := C.webkit_web_view_new_with_context(context.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapWebView(obj)
	b.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return b, nil
}

// CanGoBack() is a wrapper around webkit_web_view_can_go_back().
func (v *WebView) CanGoBack() bool {
	return gobool(C.webkit_web_view_can_go_back(v.native()))
}

// CanGoForward() is a wrapper around webkit_web_view_can_go_forward().
func (v *WebView) CanGoForward() bool {
	return gobool(C.webkit_web_view_can_go_forward(v.native()))
}

// CanShowMimeType is a wrapper around webkit_web_view_can_show_mime_type().
func (v *WebView) CanShowMimeType(mime string) bool {
	cmime := C.CString(mime)
	defer C.free(unsafe.Pointer(cmime))
	c := C.webkit_web_view_can_show_mime_type(v.native(), (*C.gchar)(cmime))
	return gobool(c)
}

// GetContext() is a wrapper around webkit_web_view_get_context().
func (v *WebView) GetContext() *WebContext {
	c := C.webkit_web_view_get_context(v.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWebContext(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w
}

// GetCustomCharset() is a wrapper around webkit_web_view_get_custom_charset().
func (v *WebView) GetCustomCharset() string {
	c := C.webkit_web_view_get_custom_charset(v.native())
	return C.GoString((*C.char)(c))
}

// GetEstimatedProgress() is a wrapper around webkit_web_view_get_estimated_load_progress().
func (v *WebView) GetEstimatedProgress() float64 {
	return float64(C.webkit_web_view_get_estimated_load_progress(v.native()))
}

// GetPageId() is a wrapper around webkit_web_view_get_page_id().
func (v *WebView) GetPageID() uint64 {
	return uint64(C.webkit_web_view_get_page_id(v.native()))
}

// GetSettings() is a wrapper around webkit_web_view_get_settings().
func (v *WebView) GetSettings() *Settings {
	c := C.webkit_web_view_get_settings(v.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapSettings(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w
}

// GetTitle() is a wrapper around webkit_web_view_get_title().
func (v *WebView) GetTitle() string {
	c := C.webkit_web_view_get_title(v.native())
	return C.GoString((*C.char)(c))
}

// GetUri() is a wrapper around webkit_web_view_get_uri().
func (v *WebView) GetURI() string {
	c := C.webkit_web_view_get_uri(v.native())
	return C.GoString((*C.char)(c))
}

// GetWindowProperties() is a wrapper around webkit_web_view_get_window_properties().
func (v *WebView) GetWindowProperties() *WindowProperties {
	c := C.webkit_web_view_get_window_properties(v.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapWindowProperties(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w
}

// GetZoomLevel() is a wrapper around webkit_web_view_get_zoom_level().
func (v *WebView) GetZoomLevel() float64 {
	return float64(C.webkit_web_view_get_zoom_level(v.native()))
}

// GoBack() is a wrapper around webkit_web_view_go_back().
func (v *WebView) GoBack() {
	C.webkit_web_view_go_back(v.native())
}

// GoForward() is a wrapper around webkit_web_view_go_forward().
func (v *WebView) GoForward() {
	C.webkit_web_view_go_forward(v.native())
}

// IsLoading() is a wrapper around webkit_web_view_is_loading().
func (v *WebView) IsLoading() bool {
	return gobool(C.webkit_web_view_is_loading(v.native()))
}

// LoadAlternateHtml() is a wrapper around webkit_web_view_load_alternate_html().
func (v *WebView) LoadAlternateHtml(content, contentUri, baseUri string) {
	ccontent := C.CString(content)
	defer C.free(unsafe.Pointer(ccontent))
	ccontentUri := C.CString(contentUri)
	defer C.free(unsafe.Pointer(ccontentUri))
	cbaseUri := C.CString(baseUri)
	defer C.free(unsafe.Pointer(cbaseUri))
	C.webkit_web_view_load_alternate_html(v.native(), (*C.gchar)(ccontent), (*C.gchar)(ccontentUri), (*C.gchar)(cbaseUri))
}

// LoadHtml() is a wrapper around webkit_web_view_load_html().
func (v *WebView) LoadHTML(content, baseUri string) {
	ccontent := C.CString(content)
	defer C.free(unsafe.Pointer(ccontent))
	cbaseUri := C.CString(baseUri)
	defer C.free(unsafe.Pointer(cbaseUri))
	C.webkit_web_view_load_html(v.native(), (*C.gchar)(ccontent), (*C.gchar)(cbaseUri))
}

// LoadPlainText() is a wrapper around webkit_web_view_load_plain_text().
func (v *WebView) LoadPlainText(text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.webkit_web_view_load_plain_text(v.native(), (*C.gchar)(ctext))
}

// LoadUri() is a wrapper around webkit_web_view_load_uri().
func (v *WebView) LoadURI(uri string) {
	curi := C.CString(uri)
	defer C.free(unsafe.Pointer(curi))
	C.webkit_web_view_load_uri(v.native(), (*C.gchar)(curi))
}

// Reload() is a wrapper around webkit_web_view_reload(),
func (v *WebView) Reload() {
	C.webkit_web_view_reload(v.native())
}

// ReloadBypassCache() is a wrapper around webkit_web_view_reload_bypass_cache().
func (v *WebView) ReloadBypassCache() {
	C.webkit_web_view_reload_bypass_cache(v.native())
}

// SetCustomCharset() is a wrapper around webkit_web_view_set_custom_charset().
func (v *WebView) SetCustomCharset(charset string) {
	ccharset := C.CString(charset)
	defer C.free(unsafe.Pointer(ccharset))
	C.webkit_web_view_set_custom_charset(v.native(), (*C.gchar)(ccharset))
}

// SetSettings() is a wrapper around webkit_web_view_set_settings().
func (v *WebView) SetSettings(settings *Settings) {
	C.webkit_web_view_set_settings(v.native(), settings.native())
}

func (v *WebView) SetZoomLevel(zoom float64) {
	C.webkit_web_view_set_zoom_level(v.native(), C.gdouble(zoom))
}

// StopLoading() is a wrapper around webkit_web_view_stop_loading().
func (v *WebView) StopLoading() {
	C.webkit_web_view_stop_loading(v.native())
}

/*
 * WebKitWindowProperties
 */

type WindowProperties struct {
	*glib.Object
}

// native returns a pointer to the underlying WebKitWindowProperties.
func (v *WindowProperties) native() *C.WebKitWindowProperties {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toWebKitWindowProperties(p)
}

func marshalWindowProperties(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapWindowProperties(obj), nil
}

func wrapWindowProperties(obj *glib.Object) *WindowProperties {
	return &WindowProperties{obj}
}

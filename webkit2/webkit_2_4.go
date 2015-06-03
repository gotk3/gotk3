package webkit

// #cgo pkg-config: webkit2gtk-4.0
// #include <webkit2/webkit2.h>
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
		{glib.Type(C.webkit_process_model_get_type()), marshalProcessModel},

		// Objects/Interfaces
		//{glib.Type(C.webkit_user_content_manager_get_type()), marshalUserContentManager},

		// Boxed
	}
	glib.RegisterGValueMarshalers(tm)
}

/*
 * Constants
 */

type ProcessModel int

const (
	PROCESS_MODEL_SHARED_SECONDARY_PROCESS     ProcessModel = C.WEBKIT_PROCESS_MODEL_SHARED_SECONDARY_PROCESS
	PROCESS_MODEL_MULTIPLE_SECONDARY_PROCESSES ProcessModel = C.WEBKIT_PROCESS_MODEL_MULTIPLE_SECONDARY_PROCESSES
)

func marshalProcessModel(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return ProcessModel(c), nil
}

/*
 * WebKitWebContext
 */

// GetProcessModel() is a wrapper around webkit_web_context_get_process_model().
func (v *WebContext) GetProcessModel() ProcessModel {
	return ProcessModel(C.webkit_web_context_get_process_model(v.native()))
}

// SetProcessModel() is a wrapper around webkit_web_context_set_process_model().
func (v *WebContext) SetProcessModel(model ProcessModel) {
	C.webkit_web_context_set_process_model(v.native(), C.WebKitProcessModel(model))
}

/*
 * WebKitWebView
 */

// WebViewNewWithRelatedView() is a wrapper around webkit_web_view_new_with_related_view().
func WebViewNewWithRelatedView(view *WebView) (*WebView, error) {
	c := C.webkit_web_view_new_with_related_view(view.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapWebView(obj)
	b.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return b, nil
}
